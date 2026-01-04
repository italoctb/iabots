package main

import (
	"log"

	"iabots-server/configs"
	"iabots-server/internal/app"
	"iabots-server/internal/delivery/http/routes"
	"iabots-server/internal/infra/database"

	"github.com/gin-gonic/gin"
)

func main() {
	// 1) env
	configs.LoadEnv()

	// 2) db
	db, err := database.NewDatabase()
	if err != nil {
		log.Fatalf("database error: %v", err)
	}

	// 3) modules
	appModules := app.NewAppModules(db)

	// 4) router
	router := gin.New()
	router.Use(gin.Logger(), gin.Recovery())

	// 5) routes
	routes.RegisterRoutes(router, routes.RouteDependencies{
		UserHandler: appModules.User.Handler,
	})

	// 6) start
	log.Println("🚀 API running on port", configs.Env.AppPort, "| env =", configs.Env.AppEnv)
	_ = router.Run(":" + configs.Env.AppPort)
}
