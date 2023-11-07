package main

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"sync"
	"syscall"
	"time"

	"whatsapp_client/models"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
	_ "github.com/mattn/go-sqlite3"
	"github.com/mdp/qrterminal"
	"go.mau.fi/whatsmeow"
	waProto "go.mau.fi/whatsmeow/binary/proto"
	"go.mau.fi/whatsmeow/store/sqlstore"
	"go.mau.fi/whatsmeow/types"
	"go.mau.fi/whatsmeow/types/events"
	waLog "go.mau.fi/whatsmeow/util/log"
)

type Wpp struct {
	Client *whatsmeow.Client
}

func NewWpp(client *whatsmeow.Client) *Wpp {
	app := &Wpp{Client: client}
	client.AddEventHandler(app.EventHandler)
	return app
}

func (wpp *Wpp) EventHandler(evt interface{}) {
	switch v := evt.(type) {
	case *events.Message:
		fmt.Println("Received a message!", v.Message.GetConversation())
		wpp.Client.SendChatPresence(v.Info.Sender, types.ChatPresenceComposing, types.ChatPresenceMediaText)
		msg := v.Message.GetConversation()
		if v.Message.ExtendedTextMessage != nil {
			msg += v.Message.ExtendedTextMessage.GetText()
		}
		fmt.Println("Numero do usuario:", v.Info.Sender.User)
		text, err := getResponseTextWithRetry(msg, v.Info.Sender.User)
		fmt.Println("Texto da resposta:", text)
		if err != nil {
			fmt.Println("Erro ao obter o texto da resposta:", err)
		}

		if text != "" {
			wpp.Client.SendMessage(context.Background(), v.Info.Sender, &waProto.Message{
				Conversation: &text},
			)
		}
	}
}

func getResponseTextWithRetry(message string, jwid string) (string, error) {
	for i := 1; i <= 5; i++ {
		fmt.Println("Tentativa: ", i)
		responseText, err := getResponseText(message, jwid)
		if err == nil && responseText != "" {
			return responseText, err
		}
		time.Sleep(5 * time.Second)
	}
	return "Desculpas! Poderia repetir?", nil
}

func getResponseText(message string, jwid string) (string, error) {
	currentRoleMessage := models.RoleMessage{
		Role:    "user",
		Content: message,
	}
	addMessageToHistory(jwid, currentRoleMessage)

	payload := models.IabotsPayload{
		CustomerID: 2,
		Messages:   getMessagesFromHistory(jwid),
	}

	payloadJSON, err := json.Marshal(payload)
	if err != nil {
		return "", err
	}

	response, err := http.Post("https://whatsapp-api-pv.herokuapp.com/api/v1/faq/gpt", "application/json", bytes.NewBuffer(payloadJSON))
	if err != nil {
		return "", err
	}
	defer response.Body.Close()

	var jsonResponse struct {
		Text string `json:"text"`
	}

	decoder := json.NewDecoder(response.Body)
	if err := decoder.Decode(&jsonResponse); err != nil {
		return "", err
	}

	text := jsonResponse.Text

	newRoleMessage := models.RoleMessage{
		Role:    "assistant",
		Content: text,
	}

	addMessageToHistory(jwid, newRoleMessage)

	return text, nil
}

func dbConn() string {
	return os.Getenv("DATABASE_URL")
}

var (
	historyMessages = make(map[string][]models.RoleMessage)
	mu              sync.Mutex
	maxNumMessages  = 20
)

func addMessageToHistory(jwid string, message models.RoleMessage) {
	mu.Lock()
	defer mu.Unlock()
	if historyMessages[jwid] == nil {
		historyMessages[jwid] = make([]models.RoleMessage, 0)
	}
	if len(historyMessages[jwid]) >= maxNumMessages {
		historyMessages[jwid] = historyMessages[jwid][1:]
	}
	historyMessages[jwid] = append(historyMessages[jwid], message)
}

func getMessagesFromHistory(jwid string) []models.RoleMessage {
	return historyMessages[jwid]
}

func main() {
	godotenv.Load()
	dbLog := waLog.Stdout("Database", "DEBUG", true)
	// Make sure you add appropriate DB connector imports, e.g. github.com/mattn/go-sqlite3 for SQLite
	container, err := sqlstore.New("postgres", dbConn(), dbLog)
	if err != nil {
		panic(err)
	}
	// If you want multiple sessions, remember their JIDs and use .GetDevice(jid) or .GetAllDevices() instead.
	deviceStore, err := container.GetFirstDevice()
	if err != nil {
		panic(err)
	}
	clientLog := waLog.Stdout("Client", "DEBUG", true)
	client := whatsmeow.NewClient(deviceStore, clientLog)
	wpp := NewWpp(client)
	wpp.Client.SendPresence(types.PresenceAvailable)

	if client.Store.ID == nil {
		// No ID stored, new login
		qrChan, _ := client.GetQRChannel(context.Background())
		err = client.Connect()
		if err != nil {
			panic(err)
		}
		for evt := range qrChan {
			if evt.Event == "code" {
				// Render the QR code here
				qrterminal.GenerateHalfBlock(evt.Code, qrterminal.L, os.Stdout)
				// or just manually `echo 2@... | qrencode -t ansiutf8` in a terminal
				fmt.Println("QR code:", evt.Code)
			} else {
				fmt.Println("Login event:", evt.Event)
			}
		}
	} else {
		// Already logged in, just connect
		err = client.Connect()
		if err != nil {
			panic(err)
		}
	}

	// Listen to Ctrl+C (you can also do something else that prevents the program from exiting)
	c := make(chan os.Signal)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM)
	disconnected := <-c // Blocks until we get a disconnect signal
	if disconnected != nil {
		client.SendPresence(types.PresenceUnavailable)
	}

	client.Disconnect()
}
