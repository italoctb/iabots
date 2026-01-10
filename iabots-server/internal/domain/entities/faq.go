package entities

import "github.com/google/uuid"

type Faq struct {
	ID       uuid.UUID `gorm:"type:uuid;primaryKey"`
	BotID    uuid.UUID `gorm:"type:uuid;not null;index"`
	Question string    `gorm:"not null"`
	Answer   string    `gorm:"not null"`
	Vector   Vector    `gorm:"type:double precision[]"`
	IsActive bool      `gorm:"default:true"`
}
