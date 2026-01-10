package v1

import (
	"iabots-server/internal/delivery/http/handlers"
	"iabots-server/internal/delivery/http/middlewares"

	"github.com/gin-gonic/gin"
)

func RegisterFaqRoutes(rg *gin.RouterGroup, h *handlers.FaqHandler) {
	rg.Use(middlewares.DevIdentityMiddleware())
	// 🔒 Rotas protegidas
	bots := rg.Group("/bots")
	{
		bots.POST("/:botId/faqs", h.CreateFaq)
		// bots.GET("/:botId/faqs", h.List)
		// bots.GET("/:botId/faqs/search", h.Search)
	}

	// faqs := rg.Group("/faqs")
	{
		// faqs.GET("/:faqId", h.GetByID)
		// faqs.PUT("/:faqId", h.Update)
		// faqs.DELETE("/:faqId", h.Delete)
	}
}
