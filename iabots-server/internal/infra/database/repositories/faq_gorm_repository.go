package repositories

import (
	"fmt"

	. "iabots-server/internal/domain/entities"
	. "iabots-server/internal/domain/repositories"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type FaqGormRepository struct {
	db *gorm.DB
}

func NewFaqGormRepository(db *gorm.DB) FaqRepository {
	return &FaqGormRepository{db: db}
}
func (r *FaqGormRepository) Create(faq *Faq) error {
	return r.db.Create(faq).Error
}

func (r *FaqGormRepository) Update(faq *Faq) error {
	return r.db.Save(faq).Error
}

func (r *FaqGormRepository) Delete(id uuid.UUID) error {
	return r.db.Delete(&Faq{}, "id = ?", id).Error
}

func (r *FaqGormRepository) FindByID(id uuid.UUID) (*Faq, error) {
	var faq Faq
	if err := r.db.First(&faq, "id = ?", id).Error; err != nil {
		return nil, err
	}
	return &faq, nil
}

func (r *FaqGormRepository) FindByBotID(botID uuid.UUID) ([]Faq, error) {
	var faqs []Faq
	if err := r.db.Find(&faqs, "bot_id = ?", botID).Error; err != nil {
		return nil, err
	}
	return faqs, nil
}

func (r *FaqGormRepository) SearchByEmbeddings(botID uuid.UUID, embedding []float32, limit int) ([]Faq, error) {
	if len(embedding) == 0 {
		return nil, fmt.Errorf("embedding is empty")
	}

	vector := NewVector(embedding).ToString()

	query := fmt.Sprintf(`
		SELECT * FROM faqs
		WHERE bot_id = ?
		ORDER BY cosine_similarity(vector, %s) DESC
		LIMIT ?
	`, vector)

	var faqs []Faq
	if err := r.db.Raw(query, botID, limit).Scan(&faqs).Error; err != nil {
		return nil, err
	}
	return faqs, nil
}
