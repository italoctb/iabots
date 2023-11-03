package models

import (
	"time"

	"gorm.io/gorm"
)

type Customer struct {
	ID              uint           `json:"id" gorm:"primaryKey"`
	Wid             string         `json:"wid"`
	Name            string         `json:"name"`
	FallbackMessage string         `json:"fallback"`
	EndMessage      string         `json:"end_message"`
	RateTemplateID  int            `json:"rate_template_id"`
	CreatedAt       time.Time      `json:"created"`
	UpdateAt        time.Time      `json:"updated"`
	DeleteAt        gorm.DeletedAt `gorm:"index" json:"deleted"`
}
