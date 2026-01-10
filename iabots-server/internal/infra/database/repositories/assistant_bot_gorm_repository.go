package repositories

import (
	. "iabots-server/internal/domain/entities"
	. "iabots-server/internal/domain/repositories"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type AssistantBotGormRepository struct {
	db *gorm.DB
}

var _ AssistantBotRepository = (*AssistantBotGormRepository)(nil)

func NewAssistantBotGormRepository(db *gorm.DB) AssistantBotRepository {
	return &AssistantBotGormRepository{db: db}
}

func (r *AssistantBotGormRepository) Create(bot *AssistantBot) error {
	return r.db.Create(bot).Error
}

func (r *AssistantBotGormRepository) Update(bot *AssistantBot) error {
	return r.db.Save(bot).Error
}

func (r *AssistantBotGormRepository) FindByID(id uuid.UUID) (*AssistantBot, error) {
	var bot AssistantBot
	if err := r.db.First(&bot, "id = ?", id).Error; err != nil {
		return nil, err
	}
	return &bot, nil
}

func (r *AssistantBotGormRepository) FindByCustomerID(customerID uuid.UUID) ([]AssistantBot, error) {
	var bots []AssistantBot
	if err := r.db.Where("customer_id = ?", customerID).Find(&bots).Error; err != nil {
		return nil, err
	}
	return bots, nil
}

func (r *AssistantBotGormRepository) Delete(id uuid.UUID) error {
	return r.db.Delete(&AssistantBot{}, "id = ?", id).Error
}
