package main

import (
	"context"
	"fmt"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
	"github.com/mdp/qrterminal"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"go.mau.fi/whatsmeow"
	"go.mau.fi/whatsmeow/store/sqlstore"
	"go.mau.fi/whatsmeow/types/events"
	waLog "go.mau.fi/whatsmeow/util/log"
)

func dbConn() string {
	host := os.Getenv("DB_HOST")
	port := os.Getenv("DB_PORT")
	user := os.Getenv("DB_USER")
	password := os.Getenv("DB_PASS")
	dbname := os.Getenv("DB_NAME")

	return fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)
}

type Whatsapp struct {
	Client    *whatsmeow.Client
	Container *sqlstore.Container
	Connected chan bool
	QRcode    chan string
}

var _client *Whatsapp

func (w *Whatsapp) Start() error {
	var walog waLog.Logger
	if os.Getenv("DEBUG") != "false" {
		walog = waLog.Stdout("Database", "DEBUG", true)
	}

	container, err := sqlstore.New("postgres", dbConn(), walog)
	if err != nil {
		return err
	}
	deviceStore, err := container.GetFirstDevice()
	if err != nil {
		return err
	}
	client := whatsmeow.NewClient(deviceStore, walog)
	client.AddEventHandler(func(evt interface{}) {
		switch v := evt.(type) {
		case *events.Message:
			fmt.Println("Message received", v)
		default:
		}
	})

	if client.Store.ID == nil {
		qrChan, _ := client.GetQRChannel(context.Background())
		err = client.Connect()
		if err != nil {
			return err
		}
		for evt := range qrChan {
			if evt.Event == "code" {
				qrterminal.GenerateHalfBlock(evt.Code, qrterminal.L, os.Stdout)
				if err != nil {
					return err
				}
			}
		}
	} else {
		err = client.Connect()
		if err != nil {
			return err
		}
	}

	w.Client = client
	w.Container = container
	return nil
}

func Logout() {
	_client.Stop()
}

func (w *Whatsapp) Stop() {
	w.Client.Disconnect()
}

func (w *Whatsapp) LogoutHandler(c *gin.Context) {
	w.Client.Logout()
	c.JSON(http.StatusOK, gin.H{"message": "logged out"})
}

type LogHook struct {
}

func (h *LogHook) Fire(msg events.Message) error {
	fmt.Println("Message received", msg)
	return nil
}

func NewWhatsapp() *Whatsapp {
	if _client == nil {
		_client = &Whatsapp{}
	}
	return _client
}

func main() {
	w := NewWhatsapp()
	err := w.Start()
	defer w.Stop()
	if err != nil {
		fmt.Println(err)
	}

	r := gin.Default()
	r.POST("/api/v1/logout", w.LogoutHandler)
	r.GET("/healthz", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"status": "ok",
		})
	})
	r.GET("/metrics", gin.WrapH(promhttp.Handler()))
	r.Run(":" + os.Getenv("PORT"))
}
