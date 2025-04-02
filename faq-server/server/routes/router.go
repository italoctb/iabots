package routes

import (
	"app/server/controllers"

	"github.com/gin-gonic/gin"
)

func CORSMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Header("Access-Control-Allow-Origin", "*")
		c.Header("Access-Control-Allow-Headers", "*")
		c.Header("Access-Control-Allow-Methods", "GET,POST,PUT,DELETE")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}

		c.Next()
	}
}

func ConfigRoutes(router *gin.Engine) *gin.Engine {

	router.Use(CORSMiddleware())
	main := router.Group("api/v1")
	{
		messages := main.Group("messages")
		{
			messages.GET("/:id", controllers.ShowMessage)
			messages.GET("/", controllers.ShowMessages)
			messages.GET("/process", controllers.ProcessMessages)
			messages.POST("/", controllers.CreateMessage)
			messages.PUT("/", controllers.UpdateMessage)
			messages.DELETE("/:id", controllers.DeleteMessages)
			messages.DELETE("/deleteall", controllers.DeleteAllMessages)
			messages.POST("/positus", controllers.PositusWebhook)
		}
		templates := main.Group("templates")
		{
			templates.GET("/", controllers.ShowTemplates)
			templates.GET("/:id", controllers.ShowTemplate)
			templates.POST("/", controllers.CreateTemplate)
			templates.POST("/:id", controllers.AddOption)
			templates.PUT("/:id", controllers.UpdateTemplate)
			templates.PUT("/option/:id", controllers.UpdateOption)
			templates.DELETE("/:id", controllers.DeleteTemplate)
			templates.DELETE("deleteall", controllers.DeleteAllTemplates)
			templates.DELETE("/option/deleteall", controllers.DeleteAllOptions)
		}
		sessions := main.Group("sessions")
		{
			sessions.GET("/", controllers.ShowSessions)
			sessions.GET("/:id", controllers.ShowSession)
			sessions.POST("/", controllers.CreateSession)
			sessions.PUT("/:id", controllers.UpdateSession)
			sessions.DELETE("/:id", controllers.DeleteSession)
			sessions.DELETE("/deleteall", controllers.DeleteAllSessions)
		}
		Customer := main.Group("customer")
		{
			Customer.GET("/:id", controllers.ShowCustomer)
			Customer.GET("/", controllers.ShowCustomers)
			Customer.POST("/", controllers.CreateCustomer)
			Customer.PUT("/:id", controllers.UpdateCustomer)
			Customer.DELETE("/:id", controllers.DeleteCustomer)
			Customer.DELETE("/deleteall", controllers.DeleteAllCustomers)
			Customer.POST("/gpt", controllers.CreateCustomerGPTConfig)
			Customer.PUT("/gpt", controllers.UpdateCustomerGPTConfig)
		}
		Faq := main.Group("faq")
		{
			Faq.GET("/:id", controllers.ShowFaqByCustomer)
			Faq.GET("/", controllers.ShowFaqs)
			Faq.POST("/", controllers.CreateFaq)
			Faq.POST("/search", controllers.SearchFaqByEmbedding)
			Faq.POST("/gpt", controllers.GPTWithFaqs)
			Faq.GET("/gpt-configs", controllers.GetChatGptConfigs)
			Faq.GET("/gpt-config/:id", controllers.GetChatGptConfig)
			Faq.PUT("/", controllers.UpdateFaq)
			Faq.DELETE("/:id", controllers.DeleteFaq)
			Faq.DELETE("/deleteall", controllers.DeleteAllFaqs)
		}

		gpt := main.Group("gpt")
		{
			gpt.POST("/", controllers.MetaGPTHandler)
			gpt.GET("/", controllers.ValidateWebhook)
		}
		/*whatsapp := main.Group("whatsapp")
		{
			whatsapp.POST("/", controllers.SendMessageApi)
			whatsapp.POST("/receive", controllers.ReceiveApi)
			whatsapp.GET("/testeheroku", controllers.TesteHeroku)
			/*whatsapp.GET("/", controllers.ShowSessions)
			whatsapp.GET("/:id", controllers.ShowSession)
			whatsapp.DELETE("/:id", controllers.DeleteSession)
		}*/
	}
	return router
}
