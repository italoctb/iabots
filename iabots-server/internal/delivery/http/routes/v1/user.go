package v1

import (
	"iabots-server/internal/delivery/http/handlers"

	"github.com/gin-gonic/gin"
)

func RegisterUserRoutes(rg *gin.RouterGroup, h *handlers.UserHandler) {
	users := rg.Group("/users")
	users.POST("/", h.CreateUser)
}
