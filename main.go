package main

import (
	//"fmt"
	"app/server"

	"app/server/database"
)

func main() {
	database.StartDB()
	server := server.NewServer()
	server.Run()
}
