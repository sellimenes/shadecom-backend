package models

import "gorm.io/gorm"

type Basket struct {
	gorm.Model
	UserEmail 	string
	User 		User `gorm:"foreignKey:UserEmail;references:Email"`
	ProductID 	uint
	Product 	Product `gorm:"foreignKey:ProductID"`
	Quantity 	int
}