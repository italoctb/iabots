package repositories

import (
	"iabots-server/internal/domain/entities"

	"github.com/google/uuid"
)

type PlanRepository interface {
	Create(plan *entities.Plan) error
	Update(plan *entities.Plan) error
	FindByID(id uuid.UUID) (*entities.Plan, error)
	FindByCustomerID(customerID uuid.UUID) ([]entities.Plan, error)
	Delete(id uuid.UUID) error
}
