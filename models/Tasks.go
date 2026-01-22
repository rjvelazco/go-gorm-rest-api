package models

import "gorm.io/gorm"

type Task struct {
	gorm.Model
	Title       string `gorm:"not null;type:varchar(100);unique_index" json:"title"`
	Description string `gorm:"type:varchar(255)" json:"description"`
	Done        bool   `gorm:"default:false" json:"done"`
	UserID      uint   `gorm:"not null" json:"user_id"`
}
