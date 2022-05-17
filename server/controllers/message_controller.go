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
	Message.Step = CheckStep(&Message)
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

func ProcessMessages(c *gin.Context) {
	db := database.GetDatabase()

	var Messages []models.Message

	err := db.Where("processed_at = ?", false).Find(&Messages).Error
	if err != nil {
		c.JSON(400, gin.H{
			"error": "cannot list Messages: " + err.Error(),
		})

		return
	}
	for _, Message := range Messages {
		var bot bots.ExampleBot
		pipelines.ChainProcess(bot, &Message)
	}

	c.JSON(200, Messages)
}

func CheckStep(Message *models.Message) int {
	var LastMessage models.Message
	db := database.GetDatabase()
	err := db.Where("wid_sender = ?", Message.WidSender).Last(&LastMessage).Error
	if err != nil {
		return 0
	}
	return LastMessage.Step
}

func PositusWebhook(c *gin.Context) {
	db := database.GetDatabase()

	var PositusResponse adapters.ResposeType

	err := c.ShouldBindJSON(&PositusResponse)

	if err != nil {
		c.JSON(400, gin.H{
			"error": "cannot bind JSON: " + err.Error(),
		})
		return
	}

	for _, PositusMessage := range PositusResponse.Messages {
		if PositusMessage.Type == "text" {
			Message := models.Message{
				WidSender: PositusMessage.From,
				Message:   PositusMessage.Text.Body,
			}
			db.Create(&Message)
			Bot := bots.ExampleBot{}
			pipelines.ChainProcess(Bot, &Message)
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
