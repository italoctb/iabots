package repositories

import (
	"iabots-server/internal/domain/entities"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type CreditTransactionGormRepository struct {
	db *gorm.DB
}

func NewCreditTransactionGormRepository(db *gorm.DB) *CreditTransactionGormRepository {
	return &CreditTransactionGormRepository{db: db}
}

func (r *CreditTransactionGormRepository) Create(tx *entities.CreditTransaction) error {
	return r.db.Create(tx).Error
}

func (r *CreditTransactionGormRepository) FindByCustomerID(customerID uuid.UUID) ([]entities.CreditTransaction, error) {
	var transactions []entities.CreditTransaction
	if err := r.db.Where("customer_id = ?", customerID).Find(&transactions).Error; err != nil {
		return nil, err
	}
	return transactions, nil
}

func (r *CreditTransactionGormRepository) FindByConsultantID(consultantID uuid.UUID) ([]entities.CreditTransaction, error) {
	var transactions []entities.CreditTransaction
	if err := r.db.Where("consultant_id = ?", consultantID).Find(&transactions).Error; err != nil {
		return nil, err
	}
	return transactions, nil
}

func (r *CreditTransactionGormRepository) FindAll() ([]entities.CreditTransaction, error) {
	var transactions []entities.CreditTransaction
	if err := r.db.Find(&transactions).Error; err != nil {
		return nil, err
	}
	return transactions, nil
}
