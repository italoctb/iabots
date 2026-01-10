package database

import (
	"fmt"

	. "iabots-server/internal/domain/entities"

	"gorm.io/gorm"
)

func AutoMigrate(db *gorm.DB) error {
	if err := db.AutoMigrate(
		&User{},
		&Customer{},
		&UserCustomer{},
		&AssistantBot{},
		&ModelProvider{},
		&Faq{},
		// &e.AssistantBot{},
		// &e.Faq{},
		// &e.Credits{},
		// &e.CreditTransaction{},
		// &e.SessionLog{},
		// &e.Plan{},
		// &e.ModelProvider{},
	); err != nil {
		return fmt.Errorf("auto-migrate failed: %w", err)
	}
	return nil
}
