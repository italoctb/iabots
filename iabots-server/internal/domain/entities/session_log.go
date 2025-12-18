package entities

import (
	"time"

	"github.com/google/uuid"
)

type SessionLog struct {
	ID          uuid.UUID `gorm:"primaryKey"`
	CustomerID  string
	AssistantID string
	UserInput   string
	BotResponse string
	TokensUsed  int
	CreatedAt   time.Time
}
