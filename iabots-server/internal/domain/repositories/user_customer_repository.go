package repositories

import (
	. "iabots-server/internal/domain/entities"

	"github.com/google/uuid"
)

type UserCustomerRepository interface {
	Link(userCustomer *UserCustomer) error
	Unlink(userID uuid.UUID, customerID uuid.UUID) error
	FindCustomersByUserID(userID uuid.UUID) ([]Customer, error)
	FindUsersByCustomerID(customerID uuid.UUID) ([]User, error)
	Exists(userID uuid.UUID, customerID uuid.UUID) (bool, error)
}
