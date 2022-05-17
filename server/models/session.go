package models

import (
	"strconv"
	"time"

	"gorm.io/gorm"
)

type Session struct {
	ID        uint `json:"id" gorm:"primaryKey"`
	State     string
	CreatedAt time.Time      `json:"created"`
	UpdateAt  time.Time      `json:"updated"`
	DeleteAt  gorm.DeletedAt `gorm:"index" json:"deleted"`
}

func (s Session) GetActualTemplate(db *gorm.DB) (Template, error) {
	var t Template
	id, err := strconv.Atoi(s.State)
	if err != nil {
		return Template{}, err
	}
	err = db.Where("ID=?", id).Preload("Options").First(&t).Error
	return t, err
}
