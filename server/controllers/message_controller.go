package controllers

import (
	"app/server/adapters"
	"app/server/bots"
	"app/server/database"
	"app/server/models"
	"app/server/pipelines"
	"fmt"
	"strconv"

	"github.com/gin-gonic/gin"
)

func ShowMessage(c *gin.Context) {
	id := c.Param("id")
	newid, err := strconv.Atoi(id)
	if err != nil {
		c.JSON(400, gin.H{
			"error": "ID has to be a integer",
		})
	}

	db := database.GetDatabase()

	var Message models.Message
	err = db.First(&Message, newid).Error

	if err != nil {
		c.JSON(400, gin.H{
			"error": "cannot find Message: " + err.Error(),
		})

		return
	}

	c.JSON(200, Message)
}

func CreateMessage(c *gin.Context) {
	db := database.GetDatabase()

	var Message models.Message

	err := c.ShouldBindJSON(&Message)

	if err != nil {
		c.JSON(400, gin.H{
			"error": "cannot bind JSON: " + err.Error(),
		})
		return
	}
	Message.ProcessedAt = false
	err = db.Create(&Message).Error

	if err != nil {
		c.JSON(400, gin.H{
			"error": "cannot create Message: " + err.Error(),
		})
	}
	c.JSON(200, Message)
}

func ShowMessages(c *gin.Context) {
	db := database.GetDatabase()

	var Messages []models.Message

	err := db.Order("id desc").Find(&Messages).Error

	if err != nil {
		c.JSON(400, gin.H{
			"error": "cannot list Messages: " + err.Error(),
		})

		return
	}

	c.JSON(200, Messages)
}

func UpdateMessage(c *gin.Context) {
	db := database.GetDatabase()
	var Message models.Message
	err := c.ShouldBindJSON(&Message)
	if err != nil {
		c.JSON(400, gin.H{"error": "Cannot bind the JSON"})
		return
	}
	err = db.Save(&Message).Error
	if err != nil {
		c.JSON(400, gin.H{"error": "cannot update Message: " + err.Error()})
	}

	c.JSON(200, Message)
}

func DeleteMessages(c *gin.Context) {
	id := c.Param("id")
	newid, err := strconv.Atoi(id)
	if err != nil {
		c.JSON(400, gin.H{
			"error": "ID has to be a integer",
		})
		return
	}

	db := database.GetDatabase()
	err = db.Delete(&models.Message{}, newid).Error

	if err != nil {
		c.JSON(400, gin.H{"error": "Cannot find the ID: " + err.Error()})
		return
	}

	c.Status(204)
}

func DeleteAllMessages(c *gin.Context) {
	db := database.GetDatabase()
	err := db.Where("1 = 1").Delete(&models.Message{}).Error

	if err != nil {
		c.JSON(400, gin.H{"error": "Cannot find the ID: " + err.Error()})
		return
	}

	c.Status(204)
}

func ProcessMessages(c *gin.Context) {
	db := database.GetDatabase()

	var Messages []models.Message

	var Client models.Client

	err := db.Where("processed_at = ?", false).Find(&Messages).Error
	if err != nil {
		c.JSON(400, gin.H{
			"error": "cannot list Messages: " + err.Error(),
		})

		return
	}

	err = db.First(&Client).Error

	if err != nil {
		c.JSON(400, gin.H{
			"error": "cannot retrieve Client: " + err.Error(),
		})

		return
	}

	for _, Message := range Messages {
		var bot bots.ExampleBot
		pipelines.ChainProcess(bot, Client, &Message)
	}

	c.JSON(200, Messages)
}

func PositusWebhook(c *gin.Context) {
	db := database.GetDatabase()

	var PositusResponse adapters.ResposeType

	var Client models.Client

	err := c.ShouldBindJSON(&PositusResponse)

	if err != nil {
		c.JSON(400, gin.H{
			"error": "cannot bind JSON: " + err.Error(),
		})
		return
	}

	err = db.First(&Client).Error

	if err != nil {
		c.JSON(400, gin.H{
			"error": "cannot retrieve Client: " + err.Error(),
		})

		return
	}

	for _, PositusMessage := range PositusResponse.Messages {
		if PositusMessage.Type == "text" {
			Message := models.Message{
				WidSender:   PositusMessage.From,
				WidReceiver: Client.Wid,
				Message:     PositusMessage.Text.Body,
			}
			db.Create(&Message)
			Bot := bots.ExampleBot{}
			pipelines.ChainProcess(Bot, Client, &Message)
			fmt.Println(Message)
		}
	}

	if err != nil {
		c.JSON(400, gin.H{
			"error": "cannot create Message: " + err.Error(),
		})
	}
	c.JSON(200, PositusResponse)
}
