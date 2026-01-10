package repositories

import (
	. "iabots-server/internal/domain/entities"
	i "iabots-server/internal/domain/repositories"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type UserGormRepository struct {
	db *gorm.DB
}

var _ i.UserRepository = (*UserGormRepository)(nil)

func NewUserGormRepository(db *gorm.DB) i.UserRepository {
	return &UserGormRepository{db: db}
}

func (r *UserGormRepository) Create(user *User) error {
	return r.db.Create(user).Error
}

func (r *UserGormRepository) Update(user *User) error {
	return r.db.Save(user).Error
}

func (r *UserGormRepository) FindByID(id uuid.UUID) (*User, error) {
	var user User
	if err := r.db.First(&user, "id = ?", id).Error; err != nil {
		return nil, err
	}
	return &user, nil
}

func (r *UserGormRepository) FindByEmail(email string) (*User, error) {
	var user User
	if err := r.db.First(&user, "email = ?", email).Error; err != nil {
		return nil, err
	}
	return &user, nil
}

func (r *UserGormRepository) Delete(id uuid.UUID) error {
	return r.db.Delete(&User{}, "id = ?", id).Error
}
