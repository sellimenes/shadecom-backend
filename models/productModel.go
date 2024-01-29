package models

import (
	"encoding/json"

	"gorm.io/gorm"
)

type Product struct{
	gorm.Model
	Name 			string
	Slug 			string
	Description 	string			`gorm:"nullable"`
	Price 			float64
	Stock 			int
	Images      	json.RawMessage `gorm:"type:json"`
	CoverImage 		string
	IsActive 		bool			`gorm:"default:true"`
	IsSale 			bool			`gorm:"default:false"`
	IsFeatured 		bool			`gorm:"default:false"`
	SaleProcent 	int				`gorm:"default:0"`
	CategoryID 		int
	Category 		Category `gorm:"foreignKey:CategoryID"`
}