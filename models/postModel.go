package models

import "gorm.io/gorm"

type Post struct {
	gorm.Model
	Title string
	Body  string
}
type User struct {
	gorm.Model
	Username string `gorm:"type:varchar(50);unique;not null"`
	Password string `gorm:"not null"`
}
