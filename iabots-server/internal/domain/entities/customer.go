package entities

import "github.com/google/uuid"

type Customer struct {
	ID          uuid.UUID `gorm:"primaryKey"`
	CompanyName string
	Whatsapp    string `gorm:"unique"`
	UserID      string // FK para User (cliente comum)
	AssistantID string // FK para o bot configurado
}
