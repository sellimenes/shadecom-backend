package models

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Name  string `gorm:"not null"`
    Email string `gorm:"unique;not null"`
	Password string `gorm:"nullable"`
	Phone string `gorm:"nullable"`
	Address string `gorm:"nullable"`
	RoleID uint `gorm:"default:2"`
}