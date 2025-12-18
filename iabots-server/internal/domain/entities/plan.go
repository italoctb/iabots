package entities

import "github.com/google/uuid"

type Plan struct {
	ID          uuid.UUID `gorm:"primaryKey"`
	Name        string
	Description string
	Markup      float64 // Ex: 5.0
}
