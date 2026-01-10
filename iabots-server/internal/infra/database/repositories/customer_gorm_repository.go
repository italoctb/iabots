package repositories

import (
	"iabots-server/internal/domain/entities"
	. "iabots-server/internal/domain/entities"
	i "iabots-server/internal/domain/repositories"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type CustomerGormRepository struct {
	db *gorm.DB
}

var _ i.CustomerRepository = (*CustomerGormRepository)(nil)

func NewCustomerGormRepository(db *gorm.DB) i.CustomerRepository {
	return &CustomerGormRepository{db: db}
}

func (r *CustomerGormRepository) Create(customer *Customer) error {
	return r.db.Create(customer).Error
}

func (r *CustomerGormRepository) Update(customer *Customer) error {
	return r.db.Save(customer).Error
}

func (r *CustomerGormRepository) Delete(id uuid.UUID) error {
	return r.db.Delete(&Customer{}, "id = ?", id).Error
}

func (r *CustomerGormRepository) FindByID(id uuid.UUID) (*Customer, error) {
	var customer Customer
	if err := r.db.First(&customer, "id = ?", id).Error; err != nil {
		return nil, err
	}
	return &customer, nil
}

func (r *CustomerGormRepository) FindByWhatsappNumber(number string) (*Customer, error) {
	var customer Customer
	if err := r.db.First(&customer, "whatsapp = ?", number).Error; err != nil {
		return nil, err
	}
	return &customer, nil
}

// FindByUserID retorna todos os customers que o usuário tem acesso via tabela pivô user_customers
func (r *CustomerGormRepository) FindByUserID(userID uuid.UUID) ([]entities.Customer, error) {
	var customers []entities.Customer

	err := r.db.
		Table("customers").
		Joins("JOIN user_customers ON user_customers.customer_id = customers.id").
		Where("user_customers.user_id = ?", userID).
		Find(&customers).Error

	if err != nil {
		return nil, err
	}

	return customers, nil
}
