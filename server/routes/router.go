package routes

import (
	"app/server/controllers"
	"app/server/database"
	"app/server/models"
	"fmt"
	"os"
	"strconv"

	"github.com/gin-gonic/gin"
)

func CORSMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Header("Access-Control-Allow-Origin", os.Getenv("FRONTEND_HOST"))
		c.Header("Access-Control-Allow-Headers", "*")

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
		} else {
			fmt.Println(Session.GetActualMessage(db))
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
		}
		templates := main.Group("templates")
		{
			templates.GET("/:id", controllers.ShowTemplate)
			templates.POST("/", controllers.CreateTemplate)
			templates.POST("/:id", controllers.AddOption)
		}
	}
	return router
}
