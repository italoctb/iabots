package repositories

import (
	"errors"
	"fmt"

	. "iabots-server/internal/domain/entities"
	i "iabots-server/internal/domain/repositories"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type UserCustomerGormRepository struct {
	db *gorm.DB
}

var _ i.UserCustomerRepository = (*UserCustomerGormRepository)(nil)

func NewUserCustomerGormRepository(db *gorm.DB) i.UserCustomerRepository {
	return &UserCustomerGormRepository{db: db}
}

func (r *UserCustomerGormRepository) Link(
	userCustomer *UserCustomer,
) error {
	// idempotente: se já existe, atualiza o role
	var existing UserCustomer
	err := r.db.
		Where("user_id = ? AND customer_id = ?", userCustomer.UserID, userCustomer.CustomerID).
		First(&existing).Error

	if err == nil {
		// já existe -> update role (se mudou)
		if existing.Role != userCustomer.Role {
			return r.db.Model(&UserCustomer{}).
				Where("id = ?", existing.ID).
				Update("membership_role", userCustomer.Role).Error
		}
		return nil
	}

	if !errors.Is(err, gorm.ErrRecordNotFound) {
		return err
	}

	return r.db.Create(userCustomer).Error
}

func (r *UserCustomerGormRepository) Unlink(userID uuid.UUID, customerID uuid.UUID) error {
	return r.db.
		Where("user_id = ? AND customer_id = ?", userID, customerID).
		Delete(&UserCustomer{}).Error
}

func (r *UserCustomerGormRepository) Exists(userID uuid.UUID, customerID uuid.UUID) (bool, error) {
	var count int64
	err := r.db.Model(&UserCustomer{}).
		Where("user_id = ? AND customer_id = ?", userID, customerID).
		Count(&count).Error
	if err != nil {
		return false, err
	}
	return count > 0, nil
}

func (r *UserCustomerGormRepository) FindCustomersByUserID(userID uuid.UUID) ([]Customer, error) {
	var customers []Customer

	// JOIN user_customers -> customers
	err := r.db.
		Table("customers").
		Select("customers.*").
		Joins("JOIN user_customers ON user_customers.customer_id = customers.id").
		Where("user_customers.user_id = ?", userID).
		Find(&customers).Error

	if err != nil {
		return nil, err
	}
	return customers, nil
}

func (r *UserCustomerGormRepository) FindUsersByCustomerID(customerID uuid.UUID) ([]User, error) {
	var users []User

	// JOIN user_customers -> users
	err := r.db.
		Table("users").
		Select("users.*").
		Joins("JOIN user_customers ON user_customers.user_id = users.id").
		Where("user_customers.customer_id = ?", customerID).
		Find(&users).Error

	if err != nil {
		return nil, err
	}
	return users, nil
}

func (r *UserCustomerGormRepository) GetMembershipRole(userID uuid.UUID, customerID uuid.UUID) (MembershipRoleType, error) {
	var link UserCustomer
	err := r.db.
		Select("membership_role").
		Where("user_id = ? AND customer_id = ?", userID, customerID).
		First(&link).Error

	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return "", fmt.Errorf("membership not found")
		}
		return "", err
	}

	return link.Role, nil
}
