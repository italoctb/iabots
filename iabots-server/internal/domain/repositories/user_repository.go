package repositories

import (
	"iabots-server/internal/domain/entities"

	"github.com/google/uuid"
)

type UserRepository interface {
	Create(user *entities.User) error
	Update(user *entities.User) error
	FindByID(id uuid.UUID) (*entities.User, error)
	FindByEmail(email string) (*entities.User, error)
	Delete(id uuid.UUID) error
}
