package entities

import (
	"github.com/google/uuid"
)

type User struct {
	ID       uuid.UUID `gorm:"primaryKey"`
	Email    string    `gorm:"unique"`
	Password string    // hash
	Name     string
	Role     UserRole
}
