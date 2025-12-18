package repositories

import (
	"iabots-server/internal/domain/entities"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type CustomerGormRepository struct {
	db *gorm.DB
}

func NewCustomerGormRepository(db *gorm.DB) *CustomerGormRepository {
	return &CustomerGormRepository{db: db}
}

func (r *CustomerGormRepository) Create(customer *entities.Customer) error {
	return r.db.Create(customer).Error
}

func (r *CustomerGormRepository) FindByID(id uuid.UUID) (*entities.Customer, error) {
	var customer entities.Customer
	if err := r.db.First(&customer, "id = ?", id).Error; err != nil {
		return nil, err
	}
	return &customer, nil
}

func (r *CustomerGormRepository) FindByWhatsappNumber(number string) (*entities.Customer, error) {
	var customer entities.Customer
	if err := r.db.First(&customer, "whatsapp_number = ?", number).Error; err != nil {
		return nil, err
	}
	return &customer, nil
}

func (r *CustomerGormRepository) Update(customer *entities.Customer) error {
	return r.db.Save(customer).Error
}

func (r *CustomerGormRepository) Delete(id uuid.UUID) error {
	return r.db.Delete(&entities.Customer{}, "id = ?", id).Error
}
