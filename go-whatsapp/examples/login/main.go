package main

import (
	"fmt"
	"os"
	"time"

	qrcodeTerminal "github.com/Baozisoftware/qrcode-terminal-go"
	"github.com/Rhymen/go-whatsapp"
)

func main() {
	wac, err := whatsapp.NewConn(3 * time.Minute)

	if err != nil {
		panic(err)
	}

	qr := make(chan string)
	go func() {
		terminal := qrcodeTerminal.New()
		terminal.Get(<-qr).Print()
	}()

	session, err := wac.Login(qr)
	wac.SetClientVersion(2, 2123, 7)
	if err != nil {
		fmt.Fprintf(os.Stderr, "error during login: %v\n", err)
		return
	}
	fmt.Printf("login successful, session: %v\n", session)
}
