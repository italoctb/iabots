package repositories

import (
	"iabots-server/internal/domain/entities"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type ModelProviderGormRepository struct {
	db *gorm.DB
}

func NewModelProviderGormRepository(db *gorm.DB) *ModelProviderGormRepository {
	return &ModelProviderGormRepository{db: db}
}

func (r *ModelProviderGormRepository) Create(provider *entities.ModelProvider) error {
	return r.db.Create(provider).Error
}

func (r *ModelProviderGormRepository) Update(provider *entities.ModelProvider) error {
	return r.db.Save(provider).Error
}

func (r *ModelProviderGormRepository) FindByID(id uuid.UUID) (*entities.ModelProvider, error) {
	var provider entities.ModelProvider
	if err := r.db.First(&provider, "id = ?", id).Error; err != nil {
		return nil, err
	}
	return &provider, nil
}

func (r *ModelProviderGormRepository) FindByName(name string) (*entities.ModelProvider, error) {
	var provider entities.ModelProvider
	if err := r.db.First(&provider, "name = ?", name).Error; err != nil {
		return nil, err
	}
	return &provider, nil
}

func (r *ModelProviderGormRepository) FindAll() ([]entities.ModelProvider, error) {
	var providers []entities.ModelProvider
	if err := r.db.Find(&providers).Error; err != nil {
		return nil, err
	}
	return providers, nil
}

func (r *ModelProviderGormRepository) Delete(id uuid.UUID) error {
	return r.db.Delete(&entities.ModelProvider{}, "id = ?", id).Error
}
