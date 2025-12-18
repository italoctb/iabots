package repositories

import (
	. "iabots-server/internal/domain/entities"

	"github.com/google/uuid"
)

type CreditsRepository interface {
	GetBalanceByCustomerID(customerID uuid.UUID) (*Credits, error)
	GetBalanceByConsultantID(consultantID uuid.UUID) (*Credits, error)
	UpdateBalance(customerID uuid.UUID, amountBits int) error
	UpdateConsultantBalance(consultantID uuid.UUID, amountBits int) error
	Create(credits *Credits) error
}
