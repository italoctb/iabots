package repositories

import (
	"iabots-server/internal/domain/entities"

	"github.com/google/uuid"
)

type SessionLogRepository interface {
	Create(log *entities.SessionLog) error
	FindByCustomerID(customerID uuid.UUID) ([]entities.SessionLog, error)
	FindByAssistantID(assistantID uuid.UUID) ([]entities.SessionLog, error)
}
