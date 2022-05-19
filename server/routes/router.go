package routes

import (
	"app/server/controllers"
	"context"
	"log"
	"net/http"

	oidc "github.com/coreos/go-oidc"
	"github.com/gin-gonic/gin"
	"golang.org/x/oauth2"
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

var (
	clientID     = "app"
	clientSecret = ""
)

func ConfigAuthRoutes(router *gin.Engine) {
	ctx := context.Background()
	provider, err := oidc.NewProvider(ctx, "http://localhost:8080/auth/realms/myrealm")

	if err != nil {
		log.Fatal(err)
	}
	config := oauth2.Config{
		ClientID:     clientID,
		ClientSecret: clientSecret,
		Endpoint:     provider.Endpoint(),
		RedirectURL:  "http://localhost:5000/auth/callback",
		Scopes:       []string{oidc.ScopeOpenID, "profile", "email", "roles"},
	}

	state := "exemplo"

	router.GET("/", func(c *gin.Context) {
		c.Redirect(http.StatusFound, config.AuthCodeURL(state))
	})

	router.GET("/auth/callback", func(c *gin.Context) {
		if c.Request.URL.Query().Get("state") != state {
			return
		}

		oauth2token, err := config.Exchange(ctx, c.Request.URL.Query().Get("code"))
		if err != nil {
			return
		}

		idToken, ok := oauth2token.Extra("id_token").(string)
		if !ok {
			return
		}
		c.JSON(200, struct {
			OAuth2Token *oauth2.Token
			IDToken     string
		}{
			oauth2token, idToken,
		})
	})
}

func ConfigRoutes(router *gin.Engine) *gin.Engine {
	ConfigAuthRoutes(router)
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
