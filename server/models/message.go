package models

import (
	"time"

	"gorm.io/gorm"
)

type Message struct {
	ID          uint           `json:"id" gorm:"primaryKey"`
	WidSender   string         `json:"widSender"`
	WidReceiver string         `json:"widReceiver"`
	Message     string         `json:"message"`
	SessionID   int            `json:"sessionId"`
	ProcessedAt bool           `json:"processed"`
	CreatedAt   time.Time      `json:"created"`
	UpdateAt    time.Time      `json:"updated"`
	DeleteAt    gorm.DeletedAt `gorm:"index" json:"deleted"`
}
