package v1

import (
	. "iabots-server/internal/delivery/http/handlers"
	. "iabots-server/internal/delivery/http/middlewares"

	"github.com/gin-gonic/gin"
)

func RegisterCustomerRoutes(rg *gin.RouterGroup, h *CustomerHandler, b *AssistantBotHandler) {
	protected := rg.Group("/customers")
	protected.Use(DevIdentityMiddleware())
	protected.POST("/", h.CreateCustomer)
	protected.POST("/:customerId/bots", b.CreateBot)
}
