package models

import (
	"strconv"
	"time"

	"gorm.io/gorm"
)

type Template struct {
	gorm.Model
	ID              uint           `json:"id" gorm:"primaryKey"`
	TemplateMessage string         `json:"template_message"`
	IsFirst         bool           `json:"is_first"`
	Options         []Option       `json:"options" gorm:"ForeignKey:TemplateID"`
	CreatedAt       time.Time      `json:"created"`
	UpdateAt        time.Time      `json:"updated"`
	DeleteAt        gorm.DeletedAt `gorm:"index" json:"deleted"`
}

type Option struct {
	gorm.Model
	ID         uint           `json:"id" gorm:"primaryKey"`
	Label      string         `json:"label"`
	TemplateID uint           `json:"template_id"`
	Goto       string         `json:"goto"`
	CreatedAt  time.Time      `json:"created"`
	UpdateAt   time.Time      `json:"updated"`
	DeleteAt   gorm.DeletedAt `gorm:"index" json:"deleted"`
}

func (t Template) GetMessage() string {
	msg := t.TemplateMessage + "\n"
	for index, Option := range t.Options {
		msg += string(strconv.Itoa(index)) + ". " + Option.Label + "\n"
	}
	return msg
}
