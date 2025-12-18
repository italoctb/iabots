package repositories

import (
	"iabots-server/internal/domain/entities"

	"github.com/google/uuid"
)

type TokenUsageRepository interface {
	Create(usage *entities.TokenUsage) error
	FindByCustomerID(customerID uuid.UUID) ([]entities.TokenUsage, error)
	FindByAssistantID(assistantID uuid.UUID) ([]entities.TokenUsage, error)
}
