package repositories

import (
	. "iabots-server/internal/domain/entities"

	"github.com/google/uuid"
)

type AssistantBotRepository interface {
	Create(bot *AssistantBot) error
	Update(bot *AssistantBot) error
	FindByID(id uuid.UUID) (*AssistantBot, error)
	FindByCustomerID(customerID uuid.UUID) ([]AssistantBot, error)
	Delete(id uuid.UUID) error
}
