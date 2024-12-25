package user

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Name  string `json:"name" gorm:"size:100;not null;index"`
	Email string `json:"email" gorm:"size:255;not null;unique"`
}
