package entities

import (
	"time"

	"github.com/google/uuid"
)

type TokenUsage struct {
	ID               uuid.UUID `gorm:"primaryKey"`
	AssistantID      string
	ModelProviderID  string
	PromptTokens     int
	CompletionTokens int
	TotalTokens      int
	CostInUSD        float64
	CostInBits       float64
	CreatedAt        time.Time
}
