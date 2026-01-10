package v1

import (
	"iabots-server/internal/delivery/http/handlers"

	"github.com/gin-gonic/gin"
)

type RouteDependencies struct {
	UserHandler         *handlers.UserHandler
	CustomerHandler     *handlers.CustomerHandler
	AssistantBotHandler *handlers.AssistantBotHandler
	FaqHandler          *handlers.FaqHandler
}

func RegisterRoutes(r *gin.Engine, deps RouteDependencies) {
	v1 := r.Group("/api/v1/")

	RegisterUserRoutes(v1, deps.UserHandler)

	RegisterCustomerRoutes(v1, deps.CustomerHandler, deps.AssistantBotHandler)
	// RegisterBotRoutes(protected, ...)
	RegisterFaqRoutes(v1, deps.FaqHandler)
}
