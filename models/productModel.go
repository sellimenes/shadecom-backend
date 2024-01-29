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
	CategoryID 		uint
	Category 		Category `gorm:"foreignKey:CategoryID"`
}



// {
//     "ID": 26,
//     "CreatedAt": "2024-01-02T11:42:41.278724Z",
//     "UpdatedAt": "2024-01-02T11:42:41.278724Z",
//     "DeletedAt": null,
//     "Name": "Demo Ürün 5",
//     "Slug": "demo-urun-5",
//     "Description": "",
//     "Price": 24999.99,
//     "Stock": 21,
//     "Images": [
//         "https://shadecom.s3.eu-central-1.amazonaws.com/dummy-urun-5.jpeg"
//     ],
//     "CoverImage": "https://shadecom.s3.eu-central-1.amazonaws.com/dummy-urun-5.jpeg",
//     "CategoryID": 49,
//     "Category": {
//         "ID": 0,
//         "CreatedAt": "0001-01-01T00:00:00Z",
//         "UpdatedAt": "0001-01-01T00:00:00Z",
//         "DeletedAt": null,
//         "Name": "",
//         "Slug": ""
//     },
//     "IsActive": true,
//     "IsSale": false,
//     "IsFeatured": false,
//     "SaleProcent": 0
// }