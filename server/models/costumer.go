package models

import (
	"time"

	"gorm.io/gorm"
)

type Costumer struct {
	ID              uint           `json:"id" gorm:"primaryKey"`
	Wid             string         `json:"wid"`
	Name            string         `json:"name"`
	FallbackMessage string         `json:"fallback"`
	EndMessage      string         `json:"endmessage"`
	RateTemplateID  int            `json:"rateTemplateID"`
	CreatedAt       time.Time      `json:"created"`
	UpdateAt        time.Time      `json:"updated"`
	DeleteAt        gorm.DeletedAt `gorm:"index" json:"deleted"`
}
