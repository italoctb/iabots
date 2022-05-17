package routes

import (
	"app/server/controllers"
	"app/server/database"
	"app/server/models"
	"strconv"

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

func StartSessionMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		db := database.GetDatabase()
		var Session models.Session
		err := db.Last(&Session).Error
		if err != nil {
			var FirstTemplate models.Template
			db.First(&FirstTemplate)
			Session.State = strconv.FormatUint(uint64(FirstTemplate.ID), 10)
			db.Create(&Session)
		}
		c.Next()
	}
}

func ConfigRoutes(router *gin.Engine) *gin.Engine {

	router.Use(CORSMiddleware())
	router.Use(StartSessionMiddleware())
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
		}
		sessions := main.Group("sessions")
		{
			sessions.GET("/", controllers.ShowSessions)
			sessions.GET("/:id", controllers.ShowSession)
			sessions.POST("/", controllers.CreateSession)
			sessions.PUT("/:id", controllers.UpdateSession)
			sessions.DELETE("/:id", controllers.DeleteSession)
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
