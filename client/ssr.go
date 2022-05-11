package client

import (
	"app/server/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func index(ctx *gin.Context) {
	ctx.HTML(http.StatusOK, "index.html", gin.H{
		"title": "Title",
	})
}

func messages(ctx *gin.Context) {
	m := []models.Message{
		{Message: "Teste"},
		{Message: "Teste2"},
	}
	ctx.HTML(http.StatusOK, "list.html", gin.H{
		"messages": m,
	})
}

func ServerSideHandler(router *gin.Engine) *gin.Engine {
	router.LoadHTMLGlob("./client/tmpl/**/*.html")
	router.GET("/messages", messages)
	router.GET("/", index)
	return router

}
