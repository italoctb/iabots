package repositories

import (
	. "iabots-server/internal/domain/entities"

	"github.com/google/uuid"
)

type ConsultantWithdrawalRequestRepository interface {
	Create(request *ConsultantWithdrawalRequest) error
	FindByID(id uuid.UUID) (*ConsultantWithdrawalRequest, error)
	FindAllByConsultantID(consultantID uuid.UUID) ([]ConsultantWithdrawalRequest, error)
	UpdateStatus(id uuid.UUID, status WithdrawalStatus) error
}
