package controllers

import (
	"app/server/database"
	"app/server/models"
	"strconv"

	"github.com/gin-gonic/gin"
)

func CreateTemplate(c *gin.Context) {
	db := database.GetDatabase()

	var Template models.Template
	err := c.ShouldBindJSON(&Template)

	if err != nil {
		c.JSON(400, gin.H{
			"error": "cannot bind JSON: " + err.Error(),
		})
		return
	}
	err = db.Create(&Template).Error

	if err != nil {
		c.JSON(400, gin.H{
			"error": "cannot create Template: " + err.Error(),
		})
	}
	c.JSON(200, Template)
}

func AddOption(c *gin.Context) {
	id := c.Param("id")
	newid, err := strconv.Atoi(id)
	if err != nil {
		c.JSON(400, gin.H{
			"error": "ID has to be a integer",
		})
	}

	db := database.GetDatabase()

	var Option models.Option
	c.ShouldBindJSON(&Option)

	if err != nil {
		c.JSON(400, gin.H{
			"error": "cannot bind JSON: " + err.Error(),
		})
		return
	}

	Option.TemplateID = uint(newid)
	err = db.Create(&Option).Error

	if err != nil {
		c.JSON(400, gin.H{
			"error": "cannot create Option: " + err.Error(),
		})
	}
	c.JSON(200, Option)
}

func ShowTemplate(c *gin.Context) {
	id := c.Param("id")
	newid, err := strconv.Atoi(id)
	if err != nil {
		c.JSON(400, gin.H{
			"error": "ID has to be a integer",
		})
	}

	db := database.GetDatabase()

	var Template models.Template
	err = db.Preload("Options").First(&Template, newid).Error

	if err != nil {
		c.JSON(400, gin.H{
			"error": "cannot find Template: " + err.Error(),
		})

		return
	}

	c.JSON(200, Template)
}
