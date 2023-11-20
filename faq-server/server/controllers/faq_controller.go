package controllers

import (
	"app/server/adapters"
	"app/server/database"
	"app/server/models"
	"database/sql"
	"fmt"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/gofrs/uuid"
)

func ShowFaq(c *gin.Context) {
	id := c.Param("id")
	newID, err := strconv.Atoi(id)
	if err != nil {
		c.JSON(400, gin.H{
			"error": "ID has to be an integer",
		})
		return
	}

	db := database.GetDatabaseSql()

	var faq models.Faq
	err = db.QueryRow("SELECT id, customer_id, question, answer, prompt, embedding FROM faqs WHERE id = $1", newID).Scan(&faq.ID, &faq.CustomerId, &faq.Question, &faq.Answer)

	if err != nil {
		if err == sql.ErrNoRows {
			c.JSON(404, gin.H{
				"error": "FAQ not found",
			})
			return
		}
		c.JSON(500, gin.H{
			"error": "Internal Server Error: " + err.Error(),
		})
		return
	}

	c.JSON(200, faq)
}

func CreateFaq(c *gin.Context) {
	payload := models.Faq{}
	err := c.ShouldBindJSON(&payload)
	if err != nil {
		c.JSON(400, gin.H{
			"error": "Cannot bind JSON: " + err.Error(),
		})
		return
	}

	emb, err := adapters.FetchEmbedingFromGPT4(fmt.Sprintf("p: %s \n r: %s \n", payload.Question, payload.Answer))
	if err != nil {
		c.JSON(400, gin.H{
			"error": "Cannot create FAQ: " + err.Error(),
		})
		return
	}

	payload.ID, _ = uuid.NewV7()
	//payload.Embedding = emb

	sql := fmt.Sprintln(`INSERT INTO faqs
	(id, customer_id, question, answer, vector)
	VALUES (?, ?, ?, ?, ?)`)
	params := []interface{}{
		payload.ID.String(),
		payload.CustomerId,
		payload.Question,
		payload.Answer,
		models.NewVector(emb),
	}

	db := database.GetDatabase()
	err = db.Exec(sql, params...).Error
	if err != nil {
		c.JSON(400, gin.H{
			"error": "Cannot create FAQ: " + err.Error(),
		})
		return
	}

	c.JSON(200, payload)
}

type SearchFaq struct {
	CustomerId int    `json:"customer_id"`
	Question   string `json:"question"`
	Limit      int    `json:"limit"`
}

func SearchFaqByEmbedding(c *gin.Context) {
	payload := SearchFaq{}
	err := c.ShouldBindJSON(&payload)
	if err != nil {
		c.JSON(400, gin.H{
			"error": "Cannot bind JSON: " + err.Error(),
		})
		return
	}

	emb, err := adapters.FetchEmbedingFromGPT4(payload.Question)
	if err != nil {
		c.JSON(400, gin.H{
			"error": "Cannot create FAQ: " + err.Error(),
		})
		return
	}

	db := database.GetDatabase()
	faqRepo := models.NewFAQRepo(db)
	faq, err := faqRepo.SearchByEmbeddings(payload.CustomerId, emb, payload.Limit)

	if err != nil {
		if err == sql.ErrNoRows {
			c.JSON(404, gin.H{
				"error": "FAQ not found",
			})
			return
		}
		c.JSON(500, gin.H{
			"error": "Internal Server Error: " + err.Error(),
		})
		return
	}

	c.JSON(200, faq)
}

type CustomerChatPayload struct {
	CustomerId int                    `json:"customer_id"`
	Messages   []adapters.RoleMessage `json:"messages"`
}

func GPTWithFaqs(c *gin.Context) {
	var payload CustomerChatPayload
	err := c.ShouldBindJSON(&payload)
	if err != nil {
		c.JSON(400, gin.H{
			"error": "Cannot bind JSON: " + err.Error(),
		})
		return
	}
	db := database.GetDatabase()
	var gptConfig models.ChatGPTConfig

	// use payload.CustomerID to find the config
	err = db.First(&gptConfig, "customer_id = ?", payload.CustomerId).Error
	if err != nil {
		c.JSON(400, gin.H{
			"error": "Cannot GPT respond: " + err.Error(),
		})
		return
	}

	accumulateMessages := ""
	for _, message := range payload.Messages {
		if message.Role == "user" {
			accumulateMessages += message.Content + ";"
		}
	}

	emb, err := adapters.FetchEmbedingFromGPT4(accumulateMessages)
	if err != nil {
		c.JSON(400, gin.H{
			"error": "Cannot GPT respond: " + err.Error(),
		})
		return
	}

	faqRepo := models.NewFAQRepo(db)
	faqs, err := faqRepo.SearchByEmbeddings(payload.CustomerId, emb, 5)
	if err != nil {
		c.JSON(400, gin.H{
			"error": "Cannot GPT respond: " + err.Error(),
		})
		return
	}

	completePrompt := gptConfig.Prompt
	for _, faq := range faqs {
		completePrompt += fmt.Sprintf("\nq: %s \na: %s", faq.Question, faq.Answer)
	}

	completePrompt += todayInfo()

	gptConfig.Prompt = completePrompt

	text, err := adapters.GPTChatCompletion(gptConfig, payload.Messages)
	if err != nil {
		c.JSON(400, gin.H{
			"error": "Cannot GPT respond: " + err.Error(),
		})
		return
	}

	c.JSON(200,
		gin.H{
			"config": gptConfig,
			"text":   text,
			"faqs":   faqs,
		},
	)
}

func todayInfo() string {
	today := time.Now()
	// deve retornar dia mes ano e dia da semana
	return fmt.Sprintf("; Hoje é %s, %d de %s de %d;", today.Weekday(), today.Day(), today.Month(), today.Year())
}

func ShowFaqs(c *gin.Context) {
	db := database.GetDatabase()

	faqs := []models.Faq{}
	err := db.Find(&faqs).Error

	if err != nil {
		c.JSON(400, gin.H{
			"error": "Cannot list FAQs: " + err.Error(),
		})
		return
	}

	c.JSON(200, faqs)
}

func ShowFaqsFromCustomer(c *gin.Context) {
	id := c.Param("id")
	db := database.GetDatabaseSql()

	var faqs []models.Faq

	// Executar a consulta SQL para buscar FAQs de um cliente específico
	rows, err := db.Query("SELECT * FROM faqs WHERE customer_id = $1 ORDER BY id ASC", id)

	if err != nil {
		c.JSON(400, gin.H{
			"error": "Cannot list FAQs: " + err.Error(),
		})
		return
	}

	defer rows.Close()
	for rows.Next() {
		var faq models.Faq
		err = rows.Scan(&faq.ID, &faq.CustomerId, &faq.Question, &faq.Answer) //&faq.Embedding)
		if err != nil {
			c.JSON(400, gin.H{
				"error": "Error scanning FAQ: " + err.Error(),
			})
			return
		}
		faqs = append(faqs, faq)
	}

	c.JSON(200, faqs)
}

func UpdateFaq(c *gin.Context) {
	db := database.GetDatabaseSql()
	var faq models.Faq

	err := c.ShouldBindJSON(&faq)
	if err != nil {
		c.JSON(400, gin.H{"error": "Cannot bind JSON"})
		return
	}

	// Executar a consulta SQL para atualizar o FAQ
	_, err = db.Exec("UPDATE faqs SET customer_id = $1, question = $2, answer = $3, prompt = $4, embedding = $5 WHERE id = $6",
		faq.CustomerId, faq.Question, faq.Answer, faq.ID)

	if err != nil {
		c.JSON(400, gin.H{"error": "Cannot update FAQ: " + err.Error()})
		return
	}

	c.JSON(200, faq)
}

func DeleteFaq(c *gin.Context) {
	id := c.Param("id")

	db := database.GetDatabaseSql()

	// Executar a consulta SQL para excluir o FAQ
	_, err := db.Exec("DELETE FROM faqs WHERE id = $1", id)

	if err != nil {
		c.JSON(400, gin.H{"error": "Cannot find the ID: " + err.Error()})
		return
	}

	c.Status(204)
}

func DeleteAllFaqs(c *gin.Context) {
	db := database.GetDatabaseSql()

	// Executar a consulta SQL para excluir todos os FAQs
	_, err := db.Exec("DELETE FROM faqs")

	if err != nil {
		c.JSON(400, gin.H{"error": "Cannot delete FAQs: " + err.Error()})
		return
	}

	c.Status(204)
}
