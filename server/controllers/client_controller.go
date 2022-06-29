package controllers

import (
	"app/server/database"
	"app/server/models"
	"strconv"

	"github.com/gin-gonic/gin"
)

func ShowCostumer(c *gin.Context) {
	id := c.Param("id")
	newid, err := strconv.Atoi(id)
	if err != nil {
		c.JSON(400, gin.H{
			"error": "ID has to be a integer",
		})
	}

	db := database.GetDatabase()

	var Message models.Costumer
	err = db.First(&Message, newid).Error

	if err != nil {
		c.JSON(400, gin.H{
			"error": "cannot find Message: " + err.Error(),
		})

		return
	}

	c.JSON(200, Message)
}

func CreateCostumer(c *gin.Context) {
	db := database.GetDatabase()

	var Costumer models.Costumer

	err := c.ShouldBindJSON(&Costumer)

	if err != nil {
		c.JSON(400, gin.H{
			"error": "cannot bind JSON: " + err.Error(),
		})
		return
	}
	err = db.Create(&Costumer).Error

	if err != nil {
		c.JSON(400, gin.H{
			"error": "cannot create Message: " + err.Error(),
		})
	}
	c.JSON(200, Costumer)
}

func ShowCostumers(c *gin.Context) {
	db := database.GetDatabase()

	var Costumers []models.Costumer

	err := db.Order("id desc").Find(&Costumers).Error

	if err != nil {
		c.JSON(400, gin.H{
			"error": "cannot list Messages: " + err.Error(),
		})

		return
	}

	c.JSON(200, Costumers)
}

func UpdateCostumer(c *gin.Context) {
	db := database.GetDatabase()
	var Costumer models.Costumer
	err := c.ShouldBindJSON(&Costumer)
	if err != nil {
		c.JSON(400, gin.H{"error": "Cannot bind the JSON"})
		return
	}
	err = db.Save(&Costumer).Error
	if err != nil {
		c.JSON(400, gin.H{"error": "cannot update Message: " + err.Error()})
	}

	c.JSON(200, Costumer)
}

func DeleteCostumer(c *gin.Context) {
	id := c.Param("id")
	newid, err := strconv.Atoi(id)
	if err != nil {
		c.JSON(400, gin.H{
			"error": "ID has to be a integer",
		})
		return
	}

	db := database.GetDatabase()
	err = db.Delete(&models.Costumer{}, newid).Error

	if err != nil {
		c.JSON(400, gin.H{"error": "Cannot find the ID: " + err.Error()})
		return
	}

	c.Status(204)
}

func DeleteAllCostumers(c *gin.Context) {
	db := database.GetDatabase()
	err := db.Where("1 = 1").Delete(&models.Costumer{}).Error

	if err != nil {
		c.JSON(400, gin.H{"error": "Cannot find the ID: " + err.Error()})
		return
	}

	c.Status(204)
}
