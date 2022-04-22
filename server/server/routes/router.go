package routes

import (
	"github/italoctb/restAPIproject/server/controllers"

	"github.com/gin-gonic/gin"
)

func ConfigRoutes(router *gin.Engine) *gin.Engine {
	main := router.Group("api/v1")
	{
		messages := main.Group("messages")
		{
			messages.GET("/:id", controllers.ShowMessage)
			messages.GET("/", controllers.ShowMessages)
			messages.POST("/", controllers.CreateMessage)
			messages.PUT("/", controllers.UpdateMessage)
			messages.DELETE("/:id", controllers.DeleteMessages)
		}
	}
	return router
}
