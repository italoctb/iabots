package controllers

import (
	"app/server/database"
	"app/server/models"
	"strconv"

	"github.com/gin-gonic/gin"
)

func ShowCustomer(c *gin.Context) {
	id := c.Param("id")
	newid, err := strconv.Atoi(id)
	if err != nil {
		c.JSON(400, gin.H{
			"error": "ID has to be a integer",
		})
	}

	db := database.GetDatabase()

	var Message models.Customer
	err = db.First(&Message, newid).Error

	if err != nil {
		c.JSON(400, gin.H{
			"error": "cannot find Message: " + err.Error(),
		})

		return
	}

	c.JSON(200, Message)
}

func CreateCustomer(c *gin.Context) {
	db := database.GetDatabase()

	var Customer models.Customer

	err := c.ShouldBindJSON(&Customer)

	if err != nil {
		c.JSON(400, gin.H{
			"error": "cannot bind JSON: " + err.Error(),
		})
		return
	}
	err = db.Create(&Customer).Error

	if err != nil {
		c.JSON(400, gin.H{
			"error": "cannot create Message: " + err.Error(),
		})
	}
	c.JSON(200, Customer)
}

func ShowCustomers(c *gin.Context) {
	db := database.GetDatabase()

	var Customers []models.Customer

	err := db.Order("id desc").Find(&Customers).Error

	if err != nil {
		c.JSON(400, gin.H{
			"error": "cannot list Messages: " + err.Error(),
		})

		return
	}

	c.JSON(200, Customers)
}

func UpdateCustomer(c *gin.Context) {
	db := database.GetDatabase()
	var Customer models.Customer
	err := c.ShouldBindJSON(&Customer)
	if err != nil {
		c.JSON(400, gin.H{"error": "Cannot bind the JSON"})
		return
	}
	err = db.Save(&Customer).Error
	if err != nil {
		c.JSON(400, gin.H{"error": "cannot update Message: " + err.Error()})
	}

	c.JSON(200, Customer)
}

func DeleteCustomer(c *gin.Context) {
	id := c.Param("id")
	newid, err := strconv.Atoi(id)
	if err != nil {
		c.JSON(400, gin.H{
			"error": "ID has to be a integer",
		})
		return
	}

	db := database.GetDatabase()
	err = db.Delete(&models.Customer{}, newid).Error

	if err != nil {
		c.JSON(400, gin.H{"error": "Cannot find the ID: " + err.Error()})
		return
	}

	c.Status(204)
}

func DeleteAllCustomers(c *gin.Context) {
	db := database.GetDatabase()
	err := db.Where("1 = 1").Delete(&models.Customer{}).Error

	if err != nil {
		c.JSON(400, gin.H{"error": "Cannot find the ID: " + err.Error()})
		return
	}

	c.Status(204)
}

func CreateCustomerGPTConfig(c *gin.Context) {
	db := database.GetDatabase()

	var config models.ChatGPTConfig

	err := c.ShouldBindJSON(&config)

	if err != nil {
		c.JSON(400, gin.H{
			"error": "cannot bind JSON: " + err.Error(),
		})
		return
	}
	err = db.Create(&config).Error

	if err != nil {
		c.JSON(400, gin.H{
			"error": "cannot create Message: " + err.Error(),
		})
	}
	c.JSON(200, config)
}

func UpdateCustomerGPTConfig(c *gin.Context) {
	db := database.GetDatabase()
	var config models.ChatGPTConfig
	err := c.ShouldBindJSON(&config)
	if err != nil {
		c.JSON(400, gin.H{"error": "Cannot bind the JSON"})
		return
	}
	err = db.Where("customer_id = ?", config.CustomerID).Save(&config).Error
	if err != nil {
		c.JSON(400, gin.H{"error": "cannot update Message: " + err.Error()})
	}

	c.JSON(200, config)
}
