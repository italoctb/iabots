package repositories

import (
	"iabots-server/internal/domain/entities"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type PlanGormRepository struct {
	db *gorm.DB
}

func NewPlanGormRepository(db *gorm.DB) *PlanGormRepository {
	return &PlanGormRepository{db: db}
}

func (r *PlanGormRepository) Create(plan *entities.Plan) error {
	return r.db.Create(plan).Error
}

func (r *PlanGormRepository) Update(plan *entities.Plan) error {
	return r.db.Save(plan).Error
}

func (r *PlanGormRepository) FindByID(id uuid.UUID) (*entities.Plan, error) {
	var plan entities.Plan
	if err := r.db.First(&plan, "id = ?", id).Error; err != nil {
		return nil, err
	}
	return &plan, nil
}

func (r *PlanGormRepository) FindByCustomerID(customerID uuid.UUID) ([]entities.Plan, error) {
	var plans []entities.Plan
	if err := r.db.Where("customer_id = ?", customerID).Find(&plans).Error; err != nil {
		return nil, err
	}
	return plans, nil
}

func (r *PlanGormRepository) Delete(id uuid.UUID) error {
	return r.db.Delete(&entities.Plan{}, "id = ?", id).Error
}
