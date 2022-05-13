package models

import (
	"time"

	"gorm.io/gorm"
)

type TemplateMessage struct {
	TemplateID               uint           `json:"templateId" gorm:"primaryKey"`
	WelcomeMessage           string         `json:"welcomeMessage"`
	FirstLayerMessage        string         `json:"firstLayerMessage"`
	FirstLayerAnswerMessage  string         `json:"firstLayerAnswerMessage"`
	GoodbyeMessage           string         `json:"goodbyeMessage"`
	GoodbyeAnswerMessage     string         `json:"goodbyeAnswerMessage"`
	DefaultFirstLayerMessage string         `json:"defaultFirstLayerMessage"`
	DefaultGoodbyeMessage    string         `json:"defaultGoodbyeMessage"`
	CreatedAt                time.Time      `json:"created"`
	UpdateAt                 time.Time      `json:"updated"`
	DeleteAt                 gorm.DeletedAt `gorm:"index" json:"deleted"`
}
