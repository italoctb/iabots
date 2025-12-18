package repositories

import (
	. "iabots-server/internal/domain/entities"

	"github.com/google/uuid"
)

type ConsultantRepository interface {
	Create(consultant *Consultant) error
	FindByID(id uuid.UUID) (*Consultant, error)
	FindCompaniesByConsultant(consultantID uuid.UUID) ([]Customer, error)
}
