package models

import "gorm.io/gorm"

type User struct {
	gorm.Model
	name string
	email string `gorm:"unique"`
	password string
	phone string `gorm:"nullable"`
	address string `gorm:"nullable"`
	roleID uint `gorm:"default:2"`
}