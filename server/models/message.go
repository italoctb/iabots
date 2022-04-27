package models

import (
	"time"

	"gorm.io/gorm"
)

type Message struct {
	ID          uint           `json:"id" gorm:"primaryKey"`
	WID         string         `json:"wid"`
	Message     string         `json:"message"`
	ProcessedAt bool           `json:"processed"`
	CreatedAt   time.Time      `json:"created"`
	UpdateAt    time.Time      `json:"updated"`
	DeleteAt    gorm.DeletedAt `gorm:"index" json:"deleted"`
}
