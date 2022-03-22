package main

import (
	//"fmt"
	"github/italoctb/restAPIproject/server"

	"github.com/italoctb/rest-api-project-go/server/database"
)

func main() {
	database.StartDB()
	server := server.NewServer()
	server.Run()
}
