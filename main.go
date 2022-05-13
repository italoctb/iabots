package main

import (
	"app/server"
	"app/server/database"

	"github.com/joho/godotenv"
)

func main() {
	godotenv.Load()
	database.StartDB()
	server := server.NewServer()
	server.Run()
}
