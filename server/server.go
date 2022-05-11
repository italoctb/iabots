package server

import (
	"app/client"
	"app/server/routes"
	"log"
	"os"

	"github.com/gin-gonic/gin"
)

type Server struct {
	port   string
	server *gin.Engine
}

func NewServer() Server {
	port := os.Getenv("PORT")
	return Server{port: port, server: gin.Default()}
}

func (s *Server) Run() {

	router := routes.ConfigRoutes(s.server)
	router = client.ServerSideHandler(router)

	log.Print("Server is running at port: " + s.port)
	log.Fatal(router.Run(":" + s.port))
}
