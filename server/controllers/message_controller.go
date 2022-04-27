package controllers

import (
	"app/server/database"
	"app/server/models"
	"app/server/pipelines"
	"github.com/gin-gonic/gin"
	"strconv"
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

	err := db.Find(&Messages).Error

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
		pipelines.ChainProcess(&Message)
	}

	c.JSON(200, Messages)
}
