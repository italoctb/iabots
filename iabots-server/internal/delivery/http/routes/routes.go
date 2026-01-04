package routes

import (
	"iabots-server/internal/delivery/http/handlers"
	v1 "iabots-server/internal/delivery/http/routes/v1"

	"github.com/gin-gonic/gin"
)

type RouteDependencies struct {
	UserHandler *handlers.UserHandler
}

func RegisterRoutes(r *gin.Engine, deps RouteDependencies) {
	v1.RegisterRoutes(r, v1.RouteDependencies{
		UserHandler: deps.UserHandler,
	})
}
