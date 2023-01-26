package model

import (
	"gorm.io/gorm"
)

type Todo struct {
	gorm.Model
	Title       string `gorm:"not null,type:varchar(100)" json:"title"`
	Description string `gorm:"not null,type:varchar(1000)" json:"description"`
	ParentID    int    `gorm:"index;default:null" json:"parent_id"`
	File        string `json:"file_url"`
}
