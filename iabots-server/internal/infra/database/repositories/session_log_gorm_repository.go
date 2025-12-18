package repositories

import (
	"iabots-server/internal/domain/entities"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type SessionLogGormRepository struct {
	db *gorm.DB
}

func NewSessionLogGormRepository(db *gorm.DB) *SessionLogGormRepository {
	return &SessionLogGormRepository{db: db}
}

func (r *SessionLogGormRepository) Create(log *entities.SessionLog) error {
	return r.db.Create(log).Error
}

func (r *SessionLogGormRepository) FindByCustomerID(customerID uuid.UUID) ([]entities.SessionLog, error) {
	var logs []entities.SessionLog
	if err := r.db.Where("customer_id = ?", customerID).Order("created_at desc").Find(&logs).Error; err != nil {
		return nil, err
	}
	return logs, nil
}

func (r *SessionLogGormRepository) FindByAssistantID(assistantID uuid.UUID) ([]entities.SessionLog, error) {
	var logs []entities.SessionLog
	if err := r.db.Where("assistant_id = ?", assistantID).Order("created_at desc").Find(&logs).Error; err != nil {
		return nil, err
	}
	return logs, nil
}
