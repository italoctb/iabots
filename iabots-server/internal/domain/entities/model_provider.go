package entities

import (
	"time"

	"github.com/google/uuid"
)

type ModelProvider struct {
	ID       uuid.UUID `gorm:"primaryKey"`
	Provider string    `gorm:"size:50;not null;uniqueIndex:idx_provider_model"`
	Model    string    `gorm:"size:50;not null;uniqueIndex:idx_provider_model"`

	CostPromptPer1M float64 // USD por 1M tokens de input
	CostOutputPer1M float64 // USD por 1M tokens de output

	Active    bool `gorm:"default:true"`
	CreatedAt time.Time
	UpdatedAt time.Time
}
