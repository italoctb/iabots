package models

import "gorm.io/gorm"

type Template struct {
	gorm.Model
	ID              uint     `json:"id" gorm:"primaryKey"`
	TemplateMessage string   `json:"template_message"`
	Options         []Option `json:"options" gorm:"ForeignKey:TemplateID"`
}

type Option struct {
	gorm.Model
	ID         uint   `json:"id" gorm:"primaryKey"`
	Label      string `json:"label"`
	TemplateID uint   `json:"template_id"`
}
