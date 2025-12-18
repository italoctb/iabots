package entities

import (
	"github.com/google/uuid"
)

type AssistantBot struct {
	ID              uuid.UUID `gorm:"primaryKey"`
	Name            string
	ModelProviderID string
	ContextMessage  string
	MaxTokens       int
	FreezeTime      int
	Status          AssistantStatus
	OnlineStartTime string // formato: "15:04" (ex: 08:30)
	OnlineEndTime   string // formato: "15:04" (ex: 18:45)
}
