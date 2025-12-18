package repositories

import (
	"iabots-server/internal/domain/entities"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type AssistantBotGormRepository struct {
	db *gorm.DB
}

func NewAssistantBotGormRepository(db *gorm.DB) *AssistantBotGormRepository {
	return &AssistantBotGormRepository{db: db}
}

func (r *AssistantBotGormRepository) Create(bot *entities.AssistantBot) error {
	return r.db.Create(bot).Error
}

func (r *AssistantBotGormRepository) Update(bot *entities.AssistantBot) error {
	return r.db.Save(bot).Error
}

func (r *AssistantBotGormRepository) FindByID(id uuid.UUID) (*entities.AssistantBot, error) {
	var bot entities.AssistantBot
	if err := r.db.First(&bot, "id = ?", id).Error; err != nil {
		return nil, err
	}
	return &bot, nil
}

func (r *AssistantBotGormRepository) FindByCustomerID(customerID uuid.UUID) ([]entities.AssistantBot, error) {
	var bots []entities.AssistantBot
	if err := r.db.Where("customer_id = ?", customerID).Find(&bots).Error; err != nil {
		return nil, err
	}
	return bots, nil
}

func (r *AssistantBotGormRepository) Delete(id uuid.UUID) error {
	return r.db.Delete(&entities.AssistantBot{}, "id = ?", id).Error
}
