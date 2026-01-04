package repositories

import (
	"iabots-server/internal/domain/entities"
	i "iabots-server/internal/domain/repositories"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type UserGormRepository struct {
	db *gorm.DB
}

// Segurança de que UserGormRepository implementa i.UserRepository em tempo de compilação
var _ i.UserRepository = (*UserGormRepository)(nil)

func NewUserGormRepository(db *gorm.DB) i.UserRepository {
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
