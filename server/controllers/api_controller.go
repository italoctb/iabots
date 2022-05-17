package controllers

import (
	"app/server/api"
	"app/server/database"
	"app/server/models"

	"github.com/gin-gonic/gin"
)

func SendMessageApi(c *gin.Context) {
	var Api api.Api
	Api.Adapter = models.CreatePositusObject("https://api.positus.global/v2/sandbox/whatsapp/numbers/6334ea09-d3fe-4689-8acb-684eb0d0ec78/messages", "eyJ0eXAiOiJKV1QiLCJhbGciOiJSUzI1NiJ9.eyJhdWQiOiIxIiwianRpIjoiNzRhZTc5OTBhNjhlOTdmZDVjM2YxYzViOTU0NzYxMmZmNzE5OWE1OWM1MzA5Zjg4MGU4MTNlOGU3ZDE4MzBmNTE2MmFmZTQ5MThkNWVkZWQiLCJpYXQiOjE2NTI1MDI1MjMuNjk1NTI3LCJuYmYiOjE2NTI1MDI1MjMuNjk1NTMsImV4cCI6MTY4NDAzODUyMy42OTM1MTIsInN1YiI6IjUxNTMiLCJzY29wZXMiOltdfQ.P4CNAolzIrEq7I94U5ULd-wEym3xbieRNw3iBASg4kls8umE49ujP3EJnAZSxK7ufwSbzimf_hEiP3-95j7t3bPqrxLBtZ-W5qEtuUH7WefTv3YUXTHK1HzcHajghARRuiWYA6zHjz3ozw_xOgq8HxN82g15W2V7BYTeVM8Vk0hXl6L9j4BpPhngs_6E-Ffnd2hxqdCKlC0_6sVk2iEhUaar6twNFIOGntpjUeaFa-y8ec-G_cgVMlBDuKU2oTvqG_fjeD9fqxBGTMaCWx0sPIqeU5SqV4Qib877XVqDOT5AXw5-2ClTdn55Jff00XXnEB53sPRVn4jhGtpD_El5IylS-Url2zmUD3Ir6BU2tA6rj08f6sBSXyOIIuM-LJJCcHox_h3ZXVn4_LvZliOSTfdhS1sZkFxlLBLrm3bvt1gXSFwjVZVQbe6MPuWV2sxNfhgJHPHr20juO6VDrYFShASMZlpg-EyKfnFX7pTtr8RF1dWkUhmx89V7JSi54Ui1u1fYnfAzA1LSH1WiMtF2Ncgh2yQiV7q-7QmyJet_ULwdIIZQnMgz5w1LYhRMmsBW_g4XDoBe3257_-BpgNjEi3kJaN3HZZ4hs8-9ytujut7bWKB5tQViQYIasLtBz_L-y5VMN7NsfbJl3r2qTN2PI8WHiEaDW90RhHlphfhEKI8")

	var Message models.Message

	err := c.ShouldBindJSON(&Message)

	if err != nil {
		c.JSON(400, gin.H{
			"error": "cannot bind JSON: " + err.Error(),
		})
		return
	}
	err = Api.SendTextMessage(Message.WidReceiver, Message.Message, "POST")

	if err != nil {
		c.JSON(400, gin.H{
			"error": "cannot create Message: " + err.Error(),
		})
	}
	c.JSON(200, Message)
}

func ReceiveApi(c *gin.Context) {
	db := database.GetDatabase()

	var Message models.Message

	err := c.ShouldBindJSON(&Message)

	if err != nil {
		c.JSON(400, gin.H{
			"error": "cannot bind JSON: " + err.Error(),
		})
		return
	}
	/*err = db.Create(&Message).Error

	if err != nil {
		c.JSON(400, gin.H{
			"error": "cannot create Message: " + err.Error(),
		})
	}*/
	var MessageResponse models.Message

	MessageResponse.WidReceiver = Message.WidSender
	MessageResponse.Message = "Recebido, meu querido!"
	MessageResponse.ProcessedAt = true

	err = db.Create(&MessageResponse).Error

	if err != nil {
		c.JSON(400, gin.H{
			"error": "cannot create Message: " + err.Error(),
		})
	}

	var Api api.Api
	Api.Adapter = models.CreatePositusObject("https://api.positus.global/v2/sandbox/whatsapp/numbers/6334ea09-d3fe-4689-8acb-684eb0d0ec78/messages", "eyJ0eXAiOiJKV1QiLCJhbGciOiJSUzI1NiJ9.eyJhdWQiOiIxIiwianRpIjoiNzRhZTc5OTBhNjhlOTdmZDVjM2YxYzViOTU0NzYxMmZmNzE5OWE1OWM1MzA5Zjg4MGU4MTNlOGU3ZDE4MzBmNTE2MmFmZTQ5MThkNWVkZWQiLCJpYXQiOjE2NTI1MDI1MjMuNjk1NTI3LCJuYmYiOjE2NTI1MDI1MjMuNjk1NTMsImV4cCI6MTY4NDAzODUyMy42OTM1MTIsInN1YiI6IjUxNTMiLCJzY29wZXMiOltdfQ.P4CNAolzIrEq7I94U5ULd-wEym3xbieRNw3iBASg4kls8umE49ujP3EJnAZSxK7ufwSbzimf_hEiP3-95j7t3bPqrxLBtZ-W5qEtuUH7WefTv3YUXTHK1HzcHajghARRuiWYA6zHjz3ozw_xOgq8HxN82g15W2V7BYTeVM8Vk0hXl6L9j4BpPhngs_6E-Ffnd2hxqdCKlC0_6sVk2iEhUaar6twNFIOGntpjUeaFa-y8ec-G_cgVMlBDuKU2oTvqG_fjeD9fqxBGTMaCWx0sPIqeU5SqV4Qib877XVqDOT5AXw5-2ClTdn55Jff00XXnEB53sPRVn4jhGtpD_El5IylS-Url2zmUD3Ir6BU2tA6rj08f6sBSXyOIIuM-LJJCcHox_h3ZXVn4_LvZliOSTfdhS1sZkFxlLBLrm3bvt1gXSFwjVZVQbe6MPuWV2sxNfhgJHPHr20juO6VDrYFShASMZlpg-EyKfnFX7pTtr8RF1dWkUhmx89V7JSi54Ui1u1fYnfAzA1LSH1WiMtF2Ncgh2yQiV7q-7QmyJet_ULwdIIZQnMgz5w1LYhRMmsBW_g4XDoBe3257_-BpgNjEi3kJaN3HZZ4hs8-9ytujut7bWKB5tQViQYIasLtBz_L-y5VMN7NsfbJl3r2qTN2PI8WHiEaDW90RhHlphfhEKI8")

	err = Api.SendTextMessage(MessageResponse.WidReceiver, MessageResponse.Message, "POST")

	if err != nil {
		c.JSON(400, gin.H{
			"error": "cannot create Message: " + err.Error(),
		})
	}
	c.JSON(200, MessageResponse)
}

func TesteHeroku(c *gin.Context) {
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
