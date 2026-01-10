package routes

import (
	. "iabots-server/internal/delivery/http/handlers"
	v1 "iabots-server/internal/delivery/http/routes/v1"

	"github.com/gin-gonic/gin"
)

type RouteDependencies struct {
	UserHandler         *UserHandler
	CustomerHandler     *CustomerHandler
	AssistantBotHandler *AssistantBotHandler
	FaqHandler          *FaqHandler
}

func RegisterRoutes(r *gin.Engine, deps RouteDependencies) {
	v1.RegisterRoutes(r, v1.RouteDependencies{
		UserHandler:         deps.UserHandler,
		CustomerHandler:     deps.CustomerHandler,
		AssistantBotHandler: deps.AssistantBotHandler,
		FaqHandler:          deps.FaqHandler,
	})
}
