package models

import "github.com/jinzhu/gorm"

// User model
type User struct {
	gorm.Model
	Name  string `gorm:"not null" json:"name"`
	Email string `json:"email" gorm:"not null;unique"`
}
