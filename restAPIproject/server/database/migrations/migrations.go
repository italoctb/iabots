package migrations

import (
	"github.com/italoctb/rest-api-project-go/server/models"
	"gorm.io/gorm"
)

func RunMigrations(db *gorm.DB) {
	db.AutoMigrate(models.Book{})
}
