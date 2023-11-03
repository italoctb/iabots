package models

import (
	"strconv"
	"time"

	"gorm.io/gorm"
)

type Template struct {
	ID              uint           `json:"id" gorm:"primaryKey"`
	WidCustomer     string         `json:"wid_customer"`
	TemplateMessage string         `json:"template_message"`
	IsFirst         bool           `json:"is_first"`
	Options         []Option       `json:"options" gorm:"ForeignKey:TemplateID"`
	CreatedAt       time.Time      `json:"created"`
	UpdateAt        time.Time      `json:"updated"`
	DeleteAt        gorm.DeletedAt `gorm:"index" json:"deleted"`
}

type Option struct {
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
		msg += string(strconv.Itoa(index+1)) + ". " + Option.Label + "\n"
	}
	return msg
}
