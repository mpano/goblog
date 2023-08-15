package models

import (
	"goblog/internal/modules/user/models"
	"gorm.io/gorm"
)

type Article struct {
	gorm.Model
	Title   string `gorm:"varchar:70"`
	Content string `gorm:"text"`
	UserId  uint
	User    models.User
}
