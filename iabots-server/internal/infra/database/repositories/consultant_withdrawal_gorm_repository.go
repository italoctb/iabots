package repositories

import (
	. "iabots-server/internal/domain/entities"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type ConsultantWithdrawalRequestGormRepository struct {
	db *gorm.DB
}

func NewConsultantWithdrawalRequestGormRepository(db *gorm.DB) *ConsultantWithdrawalRequestGormRepository {
	return &ConsultantWithdrawalRequestGormRepository{db: db}
}

func (r *ConsultantWithdrawalRequestGormRepository) Create(request *ConsultantWithdrawalRequest) error {
	return r.db.Create(request).Error
}

func (r *ConsultantWithdrawalRequestGormRepository) FindByID(id uuid.UUID) (*ConsultantWithdrawalRequest, error) {
	var req ConsultantWithdrawalRequest
	if err := r.db.First(&req, "id = ?", id).Error; err != nil {
		return nil, err
	}
	return &req, nil
}

func (r *ConsultantWithdrawalRequestGormRepository) FindAllByConsultantID(consultantID uuid.UUID) ([]ConsultantWithdrawalRequest, error) {
	var requests []ConsultantWithdrawalRequest
	if err := r.db.Where("consultant_id = ?", consultantID).Find(&requests).Error; err != nil {
		return nil, err
	}
	return requests, nil
}

func (r *ConsultantWithdrawalRequestGormRepository) UpdateStatus(id uuid.UUID, status WithdrawalStatus) error {
	return r.db.Model(&ConsultantWithdrawalRequest{}).
		Where("id = ?", id).
		Update("status", status).Error
}
