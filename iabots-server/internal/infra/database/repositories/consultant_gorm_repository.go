package repositories

import (
	"iabots-server/internal/domain/entities"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type ConsultantGormRepository struct {
	db *gorm.DB
}

func NewConsultantGormRepository(db *gorm.DB) *ConsultantGormRepository {
	return &ConsultantGormRepository{db: db}
}

func (r *ConsultantGormRepository) Create(consultant *entities.Consultant) error {
	return r.db.Create(consultant).Error
}

func (r *ConsultantGormRepository) FindByID(id uuid.UUID) (*entities.Consultant, error) {
	var consultant entities.Consultant
	if err := r.db.First(&consultant, "id = ?", id).Error; err != nil {
		return nil, err
	}
	return &consultant, nil
}

func (r *ConsultantGormRepository) FindCompaniesByConsultant(consultantID uuid.UUID) ([]entities.Customer, error) {
	var customers []entities.Customer
	if err := r.db.
		Model(&entities.Customer{}).
		Joins("JOIN consultant_customers ON consultant_customers.customer_id = customers.id").
		Where("consultant_customers.consultant_id = ?", consultantID).
		Find(&customers).Error; err != nil {
		return nil, err
	}
	return customers, nil
}
