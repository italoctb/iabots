package repositories

import (
	"iabots-server/internal/domain/entities"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type TokenUsageGormRepository struct {
	db *gorm.DB
}

func NewTokenUsageGormRepository(db *gorm.DB) *TokenUsageGormRepository {
	return &TokenUsageGormRepository{db: db}
}

func (r *TokenUsageGormRepository) Create(usage *entities.TokenUsage) error {
	return r.db.Create(usage).Error
}

func (r *TokenUsageGormRepository) FindByCustomerID(customerID uuid.UUID) ([]entities.TokenUsage, error) {
	var usages []entities.TokenUsage
	if err := r.db.Where("customer_id = ?", customerID).Order("created_at desc").Find(&usages).Error; err != nil {
		return nil, err
	}
	return usages, nil
}

func (r *TokenUsageGormRepository) FindByAssistantID(assistantID uuid.UUID) ([]entities.TokenUsage, error) {
	var usages []entities.TokenUsage
	if err := r.db.Where("assistant_id = ?", assistantID).Order("created_at desc").Find(&usages).Error; err != nil {
		return nil, err
	}
	return usages, nil
}
