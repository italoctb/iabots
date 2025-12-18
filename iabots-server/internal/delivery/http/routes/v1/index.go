package v1

import (
	"iabots-server/internal/delivery/http/handlers"

	"github.com/gin-gonic/gin"
)

type RouteDependencies struct {
	UserHandler *handlers.UserHandler
}

func RegisterRoutes(r *gin.Engine, deps RouteDependencies) {
	v1 := r.Group("/api/v1")

	RegisterUserRoutes(v1, deps.UserHandler)
	// futuras: RegisterBotRoutes, RegisterFAQRoutes, etc.
}
