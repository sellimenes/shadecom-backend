package models

import "gorm.io/gorm"

type Basket struct {
	gorm.Model
	UserEmail 	string
	ProductID 	uint
	Product 	Product `gorm:"foreignKey:ProductID"`
	Quantity 	int `gorm:"default:1"`
}