package entities

import (
	"time"

	"github.com/google/uuid"
)

type ConsultantWithdrawalRequest struct {
	ID           uuid.UUID `gorm:"primaryKey"`
	ConsultantID string
	AmountBits   float64
	Amount       float64 // valor em moeda local
	Currency     string  // Ex: "BRL"
	Status       WithdrawalStatus
	PixKey       string
	RequestedAt  time.Time
	CompletedAt  *time.Time
}
