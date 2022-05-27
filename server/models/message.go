package models

import (
	"time"

	"gorm.io/gorm"
)

type Message struct {
	ID           uint           `json:"id" gorm:"primaryKey"`
	WidSender    string         `json:"widSender"`
	WidReceiver  string         `json:"widReceiver"`
	Message      string         `json:"message"`
	FromTemplate string         `json:"fromTemplate"`
	ProcessedAt  bool           `json:"processed"`
	Step         int            `json:"step"`
	CreatedAt    time.Time      `json:"created"`
	UpdateAt     time.Time      `json:"updated"`
	DeleteAt     gorm.DeletedAt `gorm:"index" json:"deleted"`
}
