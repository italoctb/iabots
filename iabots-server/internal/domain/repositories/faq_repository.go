package repositories

import (
	. "iabots-server/internal/domain/entities"

	"github.com/google/uuid"
)

type FaqRepository interface {
	Create(faq *Faq) error
	Update(faq *Faq) error
	Delete(id uuid.UUID) error
	FindByID(id uuid.UUID) (*Faq, error)
	FindByBotID(botID uuid.UUID) ([]Faq, error)
	SearchByEmbeddings(botID uuid.UUID, embedding []float32, limit int) ([]Faq, error)
}
