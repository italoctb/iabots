package repositories

import (
	. "iabots-server/internal/domain/entities"

	"github.com/google/uuid"
)

type CustomerRepository interface {
	Create(customer *Customer) error
	FindByID(id uuid.UUID) (*Customer, error)
	FindByWhatsappNumber(number string) (*Customer, error)
	Update(customer *Customer) error
	Delete(id uuid.UUID) error
}
