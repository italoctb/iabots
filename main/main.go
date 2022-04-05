package main

import (
	//"fmt"
	"github/italoctb/whatsapp-api/server"

	//"github.com/italoctb/whatsapp-api/server/database"
)

func main() {
	//database.StartDB()
	server := server.NewServer()
	server.Run()
}
