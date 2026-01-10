package v1

import (
	"iabots-server/internal/delivery/http/handlers"
	"iabots-server/internal/delivery/http/middlewares"

	"github.com/gin-gonic/gin"
)

func RegisterUserRoutes(rg *gin.RouterGroup, h *handlers.UserHandler) {

	// 🔓 Rotas públicas
	public := rg.Group("/users")
	public.POST("/", h.CreateUser)

	// 🔒 Rotas protegidas
	protected := rg.Group("/users")
	protected.Use(middlewares.DevIdentityMiddleware())

	// protected.GET("/", h.ListUsers)
	// protected.GET("/:id", h.GetUserByID)
	// protected.DELETE("/:id", h.DeleteUser)
}
