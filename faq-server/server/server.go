package server

import (
	"app/server/routes"
	"app/server/ssr"
	"fmt"

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
	if port == "" {
		port = "5001"
	}
	fmt.Println("Porta: ", port)
	return Server{port: "8080", server: gin.Default()}
}

func (s *Server) Run() {

	router := routes.ConfigRoutes(s.server)
	router = ssr.ServerSideHandler(router)

	log.Print("Server is running at port: " + s.port)
	log.Fatal(router.Run(":" + s.port))
}
