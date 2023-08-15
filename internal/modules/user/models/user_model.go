package models

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Name     string `gorm:"varchar:70"`
	Email    string `gorm:"varchar:100"`
	Address  string `gorm:"varchar:40"`
	Password string `gorm:"varchar:100"`
}
