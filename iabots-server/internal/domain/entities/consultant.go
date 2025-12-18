package entities

import "github.com/google/uuid"

type Consultant struct {
	ID             uuid.UUID `gorm:"primaryKey"`
	UserID         string
	CustomerID     string
	CommissionRate float64 // Ex: 0.10 (10%)
}
