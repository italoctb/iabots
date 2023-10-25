package models

import (
	"database/sql/driver"
	"fmt"
	"strconv"
	"strings"

	"github.com/gofrs/uuid"
	"gorm.io/gorm"
)

type Faq struct {
	ID         uuid.UUID `gorm:"primaryKey"`
	CustomerId int       `json:"customer_id"`
	Question   string    `json:"question"`
	Answer     string    `json:"answer"`
	Embedding  []float32 `json:"embedding,omitempty"`
	Vector     Vector    `json:"vector,omitempty" gorm:"type:double precision[]"`
}

type Vector []float32

type IVector interface {
	Scan(value interface{}) error
	Value() (driver.Value, error)
	ToString() string
}

func NewVector(values []float32) IVector {
	v := Vector(values)
	return &v
}

func (v Vector) Scan(value interface{}) error {
	if value == nil {
		v = nil
		return nil
	}

	str, ok := value.(string)
	if !ok {
		return fmt.Errorf("unable to scan Vector from: %v", value)
	}

	str = strings.TrimPrefix(str, "{")
	str = strings.TrimSuffix(str, "}")
	parts := strings.Split(str, ",")

	v = make([]float32, len(parts))
	for i, part := range parts {
		val, err := strconv.ParseFloat(part, 32)
		if err != nil {
			return fmt.Errorf("unable to parse float from: %s", part)
		}
		v[i] = float32(val)
	}

	return nil
}

func (v Vector) Value() (driver.Value, error) {
	if v == nil {
		return nil, nil
	}

	var strBuilder strings.Builder
	strBuilder.WriteByte('{')

	for i, val := range v {
		if i > 0 {
			strBuilder.WriteByte(',')
		}
		strBuilder.WriteString(strconv.FormatFloat(float64(val), 'f', -1, 32))
	}

	strBuilder.WriteByte('}')
	return strBuilder.String(), nil
}

func (v Vector) ToString() string {
	stringBuilder := strings.Builder{}
	stringBuilder.WriteString("ARRAY[")
	for i, value := range v {
		if i > 0 {
			stringBuilder.WriteString(",")
		}
		stringBuilder.WriteString(strconv.FormatFloat(float64(value), 'f', -1, 32))
	}
	stringBuilder.WriteString("]")
	return stringBuilder.String()
}

type IFaqRepo interface {
	SearchByEmbeddings(customer int, embeddings []float32, limit int) ([]Faq, error)
}

type FAQRepo struct {
	db *gorm.DB
}

func NewFAQRepo(db *gorm.DB) IFaqRepo {
	return &FAQRepo{
		db: db,
	}
}

func (r *FAQRepo) SearchByEmbeddings(customer int, embeddings []float32, limit int) ([]Faq, error) {
	if len(embeddings) == 0 {
		return nil, fmt.Errorf("embeddings is empty")
	}

	var FAQs []Faq

	sql := fmt.Sprintf("SELECT * FROM faqs where customer_id=%d ORDER BY cosine_similarity(vector, %s) desc LIMIT %d", customer, NewVector(embeddings).ToString(), limit)

	err := r.db.Raw(sql).Scan(&FAQs).Error
	if err != nil {
		return nil, err
	}

	return FAQs, nil
}

/**
CREATE OR REPLACE FUNCTION cosine_similarity(a double precision[], b double precision[])
RETURNS double precision AS $$
DECLARE
    dot_product double precision := 0;
    norm_a double precision := 0;
    norm_b double precision := 0;
    i integer;
BEGIN
    IF array_length(a, 1) <> array_length(b, 1) THEN
        RAISE EXCEPTION 'The input arrays must have the same length';
    END IF;

    FOR i IN 1..array_length(a, 1)
    LOOP
        dot_product := dot_product + (a[i] * b[i]);
        norm_a := norm_a + (a[i] * a[i]);
        norm_b := norm_b + (b[i] * b[i]);
    END LOOP;

    IF norm_a = 0 OR norm_b = 0 THEN
        RETURN 0;
    ELSE
        RETURN dot_product / (sqrt(norm_a) * sqrt(norm_b));
    END IF;
END;
$$ LANGUAGE plpgsql;
**/
