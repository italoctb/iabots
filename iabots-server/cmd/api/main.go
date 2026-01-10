package main

import (
	"log"

	. "iabots-server/configs"
	. "iabots-server/internal/app"
	. "iabots-server/internal/delivery/http/routes"
	. "iabots-server/internal/infra/database"

	"github.com/gin-gonic/gin"
)

func main() {

	LoadEnv()

	db, err := NewDatabase()
	if err != nil {
		log.Fatalf("database error: %v", err)
	}

	appModules := NewAppModules(db)

	router := gin.New()
	router.Use(gin.Logger(), gin.Recovery())

	RegisterRoutes(router, RouteDependencies{
		UserHandler:         appModules.User.Handler,
		CustomerHandler:     appModules.Customer.Handler,
		AssistantBotHandler: appModules.AssistantBot.Handler,
		FaqHandler:          appModules.Faq.Handler,
	})

	log.Println("🚀 API running on port", Env.AppPort, "| env =", Env.AppEnv)
	_ = router.Run(":" + Env.AppPort)
}
