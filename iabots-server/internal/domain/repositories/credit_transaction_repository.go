package repositories

import (
	. "iabots-server/internal/domain/entities"

	"github.com/google/uuid"
)

type CreditTransactionRepository interface {
	Create(tx *CreditTransaction) error
	FindByCustomerID(customerID uuid.UUID) ([]CreditTransaction, error)
	FindByConsultantID(consultantID uuid.UUID) ([]CreditTransaction, error)
	FindAll() ([]CreditTransaction, error)
}
