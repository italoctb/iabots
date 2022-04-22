package pipelines

import (
	"app/server/database"
	"app/server/models"
)

func ChainProcess(Message *models.Message) error {
	Message.ProcessedAt = true
	db := database.GetDatabase()
	err := db.Save(&Message).Error
	return err
}
