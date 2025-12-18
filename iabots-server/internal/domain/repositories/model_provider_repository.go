package repositories

import (
	"iabots-server/internal/domain/entities"

	"github.com/google/uuid"
)

type ModelProviderRepository interface {
	Create(provider *entities.ModelProvider) error
	Update(provider *entities.ModelProvider) error
	FindByID(id uuid.UUID) (*entities.ModelProvider, error)
	FindByName(name string) (*entities.ModelProvider, error)
	FindAll() ([]entities.ModelProvider, error)
	Delete(id uuid.UUID) error
}
