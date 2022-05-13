package migrations

import (
	"app/server/models"

	"gorm.io/gorm"
)

func RunMigrations(db *gorm.DB) {
	db.AutoMigrate(models.Message{})
	db.AutoMigrate(models.Template{})
	db.AutoMigrate(models.Option{})
	db.AutoMigrate(models.Session{})
}
