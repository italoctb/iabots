package models

import (
	"time"

	"gorm.io/gorm"
)

type Message struct {
	ID          uint           `json:"id" gorm:"primaryKey"`
	WidSender   string         `json:"wid_sender"`
	WidReceiver string         `json:"wid_receiver"`
	Message     string         `json:"message"`
	SessionID   int            `json:"session_id"`
	ProcessedAt bool           `json:"processed"`
	CreatedAt   time.Time      `json:"created"`
	UpdateAt    time.Time      `json:"updated"`
	DeleteAt    gorm.DeletedAt `gorm:"index" json:"deleted"`
}
