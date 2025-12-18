package entities

import (
	"time"

	"github.com/google/uuid"
)

type Credits struct {
	ID           uuid.UUID `gorm:"primaryKey"`
	CustomerID   *string
	ConsultantID *string
	AmountBits   int
	CreatedAt    time.Time
	UpdatedAt    time.Time
}
