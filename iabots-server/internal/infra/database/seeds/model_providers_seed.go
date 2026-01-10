package seeds

import (
	"log"
	"time"

	. "iabots-server/internal/domain/entities"

	"github.com/google/uuid"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

func SeedModelProviders(db *gorm.DB) error {
	now := time.Now()

	models := []ModelProvider{
		{
			ID:              uuid.New(),
			Provider:        "openai",
			Model:           "gpt-4o-mini",
			CostPromptPer1M: 0.15,
			CostOutputPer1M: 0.60,
			Active:          true,
			CreatedAt:       now,
			UpdatedAt:       now,
		},
		{
			ID:              uuid.New(),
			Provider:        "openai",
			Model:           "gpt-5-nano",
			CostPromptPer1M: 0.05,
			CostOutputPer1M: 0.40,
			Active:          true,
			CreatedAt:       now,
			UpdatedAt:       now,
		},
		{
			ID:              uuid.New(),
			Provider:        "openai",
			Model:           "gpt-5-mini",
			CostPromptPer1M: 0.25,
			CostOutputPer1M: 2.00,
			Active:          true, // premium
			CreatedAt:       now,
			UpdatedAt:       now,
		},
	}

	for _, model := range models {
		err := db.Clauses(clause.OnConflict{
			Columns: []clause.Column{
				{Name: "provider"},
				{Name: "model"},
			},
			DoUpdates: clause.AssignmentColumns([]string{
				"cost_prompt_per1_m",
				"cost_output_per1_m",
				"active",
				"updated_at",
			}),
		}).Create(&model).Error

		if err != nil {
			return err
		}
	}

	log.Println("✅ Model providers seeded successfully")
	return nil
}
