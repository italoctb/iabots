package entities

import "github.com/google/uuid"

type Faq struct {
	ID         uuid.UUID `gorm:"primaryKey"`
	CustomerID uuid.UUID
	Question   string
	Answer     string
	Vector     Vector `gorm:"type:double precision[]"`
}
