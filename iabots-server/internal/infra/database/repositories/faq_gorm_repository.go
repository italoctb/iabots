package repositories

import (
	"fmt"

	"iabots-server/internal/domain/entities"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type FaqGormRepository struct {
	db *gorm.DB
}

func NewFaqGormRepository(db *gorm.DB) *FaqGormRepository {
	return &FaqGormRepository{db: db}
}

func (r *FaqGormRepository) Create(faq *entities.Faq) error {
	return r.db.Create(faq).Error
}

func (r *FaqGormRepository) Update(faq *entities.Faq) error {
	return r.db.Save(faq).Error
}

func (r *FaqGormRepository) Delete(id uuid.UUID) error {
	return r.db.Delete(&entities.Faq{}, "id = ?", id).Error
}

func (r *FaqGormRepository) FindByID(id uuid.UUID) (*entities.Faq, error) {
	var faq entities.Faq
	if err := r.db.First(&faq, "id = ?", id).Error; err != nil {
		return nil, err
	}
	return &faq, nil
}

func (r *FaqGormRepository) FindByCustomerID(customerID uuid.UUID) ([]entities.Faq, error) {
	var faqs []entities.Faq
	if err := r.db.Find(&faqs, "customer_id = ?", customerID).Error; err != nil {
		return nil, err
	}
	return faqs, nil
}

func (r *FaqGormRepository) SearchByEmbeddings(customerID uuid.UUID, embedding []float32, limit int) ([]entities.Faq, error) {
	if len(embedding) == 0 {
		return nil, fmt.Errorf("embedding is empty")
	}

	vector := entities.NewVector(embedding).ToString()

	query := fmt.Sprintf(`
		SELECT * FROM faqs
		WHERE customer_id = ?
		ORDER BY cosine_similarity(vector, %s) DESC
		LIMIT ?
	`, vector)

	var faqs []entities.Faq
	if err := r.db.Raw(query, customerID, limit).Scan(&faqs).Error; err != nil {
		return nil, err
	}
	return faqs, nil
}
