package repositories

import (
	"iabots-server/internal/domain/entities"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type UserGormRepository struct {
	db *gorm.DB
}

func NewUserGormRepository(db *gorm.DB) *UserGormRepository {
	return &UserGormRepository{db: db}
}

func (r *UserGormRepository) Create(user *entities.User) error {
	return r.db.Create(user).Error
}

func (r *UserGormRepository) Update(user *entities.User) error {
	return r.db.Save(user).Error
}

func (r *UserGormRepository) FindByID(id uuid.UUID) (*entities.User, error) {
	var user entities.User
	if err := r.db.First(&user, "id = ?", id).Error; err != nil {
		return nil, err
	}
	return &user, nil
}

func (r *UserGormRepository) FindByEmail(email string) (*entities.User, error) {
	var user entities.User
	if err := r.db.First(&user, "email = ?", email).Error; err != nil {
		return nil, err
	}
	return &user, nil
}

func (r *UserGormRepository) Delete(id uuid.UUID) error {
	return r.db.Delete(&entities.User{}, "id = ?", id).Error
}
