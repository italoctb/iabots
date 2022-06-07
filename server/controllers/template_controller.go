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

func UpdateTemplate(c *gin.Context) {
	db := database.GetDatabase()
	var Template models.Template
	err := c.ShouldBindJSON(&Template)
	if err != nil {
		c.JSON(400, gin.H{"error": "Cannot bind the JSON"})
		return
	}
	err = db.Save(&Template).Error
	if err != nil {
		c.JSON(400, gin.H{"error": "cannot update Template: " + err.Error()})
	}

	c.JSON(200, Template)
}

func UpdateOption(c *gin.Context) {
	db := database.GetDatabase()
	var Option models.Option
	err := c.ShouldBindJSON(&Option)
	if err != nil {
		c.JSON(400, gin.H{"error": "Cannot bind the JSON"})
		return
	}
	err = db.Save(&Option).Error
	if err != nil {
		c.JSON(400, gin.H{"error": "cannot update Option: " + err.Error()})
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
	Template.TemplateMessage = Template.GetMessage()
	if err != nil {
		c.JSON(400, gin.H{
			"error": "cannot find Template: " + err.Error(),
		})

		return
	}

	c.JSON(200, Template)
}

func ShowTemplates(c *gin.Context) {
	db := database.GetDatabase()
	var Templates []models.Template
	err := db.Preload("Options").Find(&Templates).Error

	if err != nil {
		c.JSON(400, gin.H{
			"error": "cannot find Template: " + err.Error(),
		})
		return
	}

	c.JSON(200, Templates)
}

func DeleteTemplate(c *gin.Context) {
	id := c.Param("id")
	newid, err := strconv.Atoi(id)
	if err != nil {
		c.JSON(400, gin.H{
			"error": "ID has to be a integer",
		})
		return
	}

	db := database.GetDatabase()
	err = db.Delete(&models.Template{}, newid).Error

	if err != nil {
		c.JSON(400, gin.H{"error": "Cannot find the ID: " + err.Error()})
		return
	}

	c.Status(204)
}

func DeleteAllTemplates(c *gin.Context) {

	db := database.GetDatabase()
	err := db.Where("1 = 1").Delete(&models.Template{}).Error

	if err != nil {
		c.JSON(400, gin.H{"error": "Cannot Delete, show error: " + err.Error()})
		return
	}

	c.Status(204)
}

func DeleteAllOptions(c *gin.Context) {

	db := database.GetDatabase()
	err := db.Where("1 = 1").Delete(&models.Option{}).Error

	if err != nil {
		c.JSON(400, gin.H{"error": "Cannot Delete, show error: " + err.Error()})
		return
	}

	c.Status(204)
}
