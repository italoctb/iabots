package repositories

import (
	"iabots-server/internal/domain/entities"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type CreditsGormRepository struct {
	db *gorm.DB
}

func NewCreditsGormRepository(db *gorm.DB) *CreditsGormRepository {
	return &CreditsGormRepository{db: db}
}

func (r *CreditsGormRepository) GetBalanceByCustomerID(customerID uuid.UUID) (*entities.Credits, error) {
	var credits entities.Credits
	if err := r.db.First(&credits, "customer_id = ?", customerID).Error; err != nil {
		return nil, err
	}
	return &credits, nil
}

func (r *CreditsGormRepository) GetBalanceByConsultantID(consultantID uuid.UUID) (*entities.Credits, error) {
	var credits entities.Credits
	if err := r.db.First(&credits, "consultant_id = ?", consultantID).Error; err != nil {
		return nil, err
	}
	return &credits, nil
}

func (r *CreditsGormRepository) UpdateBalance(customerID uuid.UUID, amountBits int) error {
	return r.db.Model(&entities.Credits{}).
		Where("customer_id = ?", customerID).
		Update("amount_bits", amountBits).Error
}

func (r *CreditsGormRepository) UpdateConsultantBalance(consultantID uuid.UUID, amountBits int) error {
	return r.db.Model(&entities.Credits{}).
		Where("consultant_id = ?", consultantID).
		Update("amount_bits", amountBits).Error
}

func (r *CreditsGormRepository) Create(credits *entities.Credits) error {
	return r.db.Create(credits).Error
}
