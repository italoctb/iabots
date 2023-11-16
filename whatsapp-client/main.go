package main

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"math"
	"net/http"
	"os"
	"os/signal"
	"strconv"
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
		msg := v.Message.GetConversation()
		if v.Message.ExtendedTextMessage != nil {
			msg += v.Message.ExtendedTextMessage.GetText()
		}
		fmt.Println("Numero do usuario:", v.Info.Sender.User)
		var user *UserManager
		if v.Info.IsFromMe {
			user = getUser(v.Info.Chat.User)
			user.lastUserInteraction = time.Now()
			user.addMessageToHistory(models.RoleMessage{
				Role:    "assistant",
				Content: msg,
			})
		} else {
			fmt.Println("Received a message!", v.Message.GetConversation())
			user = getUser(v.Info.Sender.User)
			user.addMessageToHistory(models.RoleMessage{
				Role:    "user",
				Content: msg,
			})

		}

		if user.context != nil {
			user.cancel()
			user.context, user.cancel = context.WithCancel(context.Background())
		}
		if user.context == nil {
			user.context, user.cancel = context.WithCancel(context.Background())
		}

		if user.lastUserInteraction.IsZero() || time.Since(user.lastUserInteraction) > 2*time.Hour {
			wpp.Client.SendChatPresence(v.Info.Sender, types.ChatPresenceComposing, types.ChatPresenceMediaText)
			text, err := GPTResponseText(user.historyMessages, user.context, 5)

			user.addMessageToHistory(models.RoleMessage{
				Role:    "assistant",
				Content: text,
			})
			fmt.Println("Texto da resposta:", text)
			if err != nil {
				fmt.Println("Erro ao obter o texto da resposta:", err)
			}

			if text != "" {
				go wpp.Client.SendMessage(user.context, v.Info.Sender, &waProto.Message{
					Conversation: &text},
				)
			}
			user.context = nil
		}

	}
}

// func getResponseTextWithRetry(message string, jwid string) (string, error) {
// 	for i := 1; i <= 5; i++ {
// 		fmt.Println("Tentativa: ", i)
// 		responseText, err := getResponseText(message, jwid)
// 		if err == nil && responseText != "" {
// 			return responseText, err
// 		}
// 		time.Sleep(5 * time.Second)
// 	}
// 	return "Desculpas! Poderia repetir?", nil
// }

func GPTResponseText(messages []models.RoleMessage, ctx context.Context, n int) (string, error) {

	payload := models.IabotsPayload{
		CustomerID: clientId,
		Messages:   messages,
	}

	payloadJSON, err := json.Marshal(payload)
	if err != nil {
		return "", err
	}

	var response *http.Response
	for i := 0; i < n; i++ {
		req, err := http.NewRequestWithContext(ctx, "POST", "https://whatsapp-api-pv.herokuapp.com/api/v1/faq/gpt", bytes.NewBuffer(payloadJSON))
		if err != nil {
			return "", err
		}
		req.Header.Set("Content-Type", "application/json")
		response, err = http.DefaultClient.Do(req)
		if err == nil && response.StatusCode == http.StatusOK {
			break
		}
		time.Sleep(time.Duration(math.Pow(2, float64(i))) * time.Second)
	}
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

	return text, nil
}

func dbConn() string {
	return os.Getenv("DATABASE_URL")
}

var (
	users          = make(map[string]*UserManager)
	mu             sync.Mutex
	maxNumMessages = 20
)

type UserManager struct {
	historyMessages     []models.RoleMessage
	context             context.Context
	cancel              context.CancelFunc
	lastUserInteraction time.Time
}

func getUser(jwid string) *UserManager {
	mu.Lock()
	defer mu.Unlock()
	if users[jwid] == nil {
		users[jwid] = &UserManager{}
	}
	return users[jwid]
}

func (u *UserManager) addMessageToHistory(message models.RoleMessage) {
	if u.historyMessages == nil {
		u.historyMessages = make([]models.RoleMessage, 0)
	}
	if len(u.historyMessages) >= maxNumMessages {
		u.historyMessages = u.historyMessages[1:]
	}
	u.historyMessages = append(u.historyMessages, message)
}

var clientId = 2

func main() {
	if len(os.Args) > 1 {
		arg := os.Args[1]
		clientId, _ = strconv.Atoi(arg)
	}
	godotenv.Load()
	dbLog := waLog.Stdout("Database", "DEBUG", true)
	// Make sure you add appropriate DB connector imports, e.g. github.com/mattn/go-sqlite3 for SQLite
	container, err := sqlstore.New("postgres", dbConn(), dbLog)
	if err != nil {
		panic(err)
	}
	// If you want multiple sessions, remember their JIDs and use .GetDevice(jid) or .GetAllDevices() instead.
	devices, err := container.GetAllDevices()
	if err != nil {
		panic(err)
	}
	for _, device := range devices {
		fmt.Println("Device: %+v", device)
	}
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
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM)
	disconnected := <-c // Blocks until we get a disconnect signal
	if disconnected != nil {
		client.SendPresence(types.PresenceUnavailable)
	}

	client.Disconnect()
}
