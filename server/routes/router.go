package routes

import (
	"github/italoctb/whatsapp-api/server/controllers"
	"github.com/gin-gonic/gin"
)

func ConfigRoutes(router *gin.Engine) *gin.Engine {
	main := router.Group("api/v1")
	{
		whatsapp := main.Group("whatsapp")
		{
			{
				whatsapp.GET("/", controllers.JustLogin)
				whatsapp.POST("/", controllers.SendMessage)
			}
		}
	}
	return router
}
