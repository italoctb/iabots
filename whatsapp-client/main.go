package main

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"whatsapp_client/models"

	_ "github.com/lib/pq"
	_ "github.com/mattn/go-sqlite3"
	"github.com/mdp/qrterminal"
	"go.mau.fi/whatsmeow"
	waProto "go.mau.fi/whatsmeow/binary/proto"
	"go.mau.fi/whatsmeow/store/sqlstore"
	"go.mau.fi/whatsmeow/types"
	"go.mau.fi/whatsmeow/types/events"
	waLog "go.mau.fi/whatsmeow/util/log"
	"google.golang.org/protobuf/proto"
)

var client *whatsmeow.Client

func eventHandler(evt interface{}) {
	switch v := evt.(type) {
	case *events.Message:
		fmt.Println("Received a message!", v.Message.GetConversation())
		if client != nil {
			// Chamar a função para obter o texto da resposta
			//client.SendChatPresence(v.Info.Sender, types.ChatPresence(t), types.ChatPresenceMediaText)
			client.SendChatPresence(v.Info.Sender, types.ChatPresenceComposing, types.ChatPresenceMediaText)
			text, err := getResponseText(v.Message.GetConversation())
			if err != nil {
				fmt.Println("Erro ao obter o texto da resposta:", err)
				return
			}

			if text != "" {
				//msgID := client.GenerateMessageID()
				client.SendMessage(context.Background(), v.Info.Sender, &waProto.Message{
					Conversation: proto.String(text),
				})
			}
		}
	}
}

func getResponseText(message string) (string, error) {
	payload := models.IabotsPayload{
		CustomerID: 2,
		Messages: []models.RoleMessage{
			{
				Role:    "user",
				Content: message,
			},
		},
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

	return jsonResponse.Text, nil
}

func dbConn() string {
	host := os.Getenv("DB_HOST")
	port := os.Getenv("DB_PORT")
	user := os.Getenv("DB_USER")
	password := os.Getenv("DB_PASSWORD")
	dbname := os.Getenv("DB_NAME")

	return fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)
}

func main() {
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
	client = whatsmeow.NewClient(deviceStore, clientLog)
	client.AddEventHandler(eventHandler)

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
	<-c

	client.Disconnect()
}
