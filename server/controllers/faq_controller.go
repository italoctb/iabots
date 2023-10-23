package controllers

import (
	"app/server/database"
	"app/server/models"
	"database/sql"
	"strconv"

	"github.com/gin-gonic/gin"
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
	err = db.QueryRow("SELECT id, customer_id, question, answer, prompt, embedding FROM faqs WHERE id = $1", newID).Scan(&faq.ID, &faq.CustomerId, &faq.Question, &faq.Answer, &faq.Embedding)

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
	db := database.GetDatabaseSql()

	var faq models.Faq

	err := c.ShouldBindJSON(&faq)

	if err != nil {
		c.JSON(400, gin.H{
			"error": "Cannot bind JSON: " + err.Error(),
		})
		return
	}

	// Executar a consulta SQL para inserir um novo FAQ
	_, err = db.Exec("INSERT INTO faqs (customer_id, question, answer, prompt, embedding) VALUES ($1, $2, $3, $4, $5)",
		faq.CustomerId, faq.Question, faq.Answer, faq.Embedding)

	if err != nil {
		c.JSON(400, gin.H{
			"error": "Cannot create FAQ: " + err.Error(),
		})
		return
	}

	c.JSON(200, faq)
}

func ShowFaqs(c *gin.Context) {
	db := database.GetDatabaseSql()

	var faqs []models.Faq

	// Executar a consulta SQL para buscar todos os FAQs
	rows, err := db.Query("SELECT * FROM faqs ORDER BY id DESC")

	if err != nil {
		c.JSON(400, gin.H{
			"error": "Cannot list FAQs: " + err.Error(),
		})
		return
	}

	defer rows.Close()
	for rows.Next() {
		var faq models.Faq
		err = rows.Scan(&faq.ID, &faq.Question, &faq.Answer, &faq.Embedding, &faq.CustomerId)
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
		err = rows.Scan(&faq.ID, &faq.CustomerId, &faq.Question, &faq.Answer, &faq.Embedding)
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
		faq.CustomerId, faq.Question, faq.Answer, faq.Embedding, faq.ID)

	if err != nil {
		c.JSON(400, gin.H{"error": "Cannot update FAQ: " + err.Error()})
		return
	}

	c.JSON(200, faq)
}

func DeleteFaq(c *gin.Context) {
	id := c.Param("id")
	newID, err := strconv.Atoi(id)
	if err != nil {
		c.JSON(400, gin.H{
			"error": "ID has to be an integer",
		})
		return
	}

	db := database.GetDatabaseSql()

	// Executar a consulta SQL para excluir o FAQ
	_, err = db.Exec("DELETE FROM faqs WHERE id = $1", newID)

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
