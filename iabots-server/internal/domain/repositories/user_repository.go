package repositories

import (
	. "iabots-server/internal/domain/entities"

	"github.com/google/uuid"
)

type UserRepository interface {
	Create(user *User) error
	Update(user *User) error
	FindByID(id uuid.UUID) (*User, error)
	FindByEmail(email string) (*User, error)
	Delete(id uuid.UUID) error
}
