package models

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Name string
	Email string `gorm:"unique"`
	Password string `gorm:"nullable"`
	Phone string `gorm:"nullable"`
	Address string `gorm:"nullable"`
	RoleID uint `gorm:"default:2"`
}