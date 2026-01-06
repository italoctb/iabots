package database

import (
	"fmt"

	"iabots-server/internal/domain/entities"

	"gorm.io/gorm"
)

func AutoMigrate(db *gorm.DB) error {
	if err := db.AutoMigrate(
		&entities.User{},
		// &entities.Customer{},
		// &entities.AssistantBot{},
		// &entities.Faq{},
		// &entities.Credits{},
		// &entities.CreditTransaction{},
		// &entities.SessionLog{},
		// &entities.Plan{},
		// &entities.ModelProvider{},
	); err != nil {
		return fmt.Errorf("auto-migrate failed: %w", err)
	}
	return nil
}
