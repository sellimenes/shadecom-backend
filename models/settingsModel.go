package models

import "gorm.io/gorm"

type Settings struct {
	gorm.Model
	websiteName 			string
	websiteDescription 		string
	websiteLogo 			string
	websiteFavicon 			string
	websiteKeywords 		string
	websiteEmail 			string
	websitePhone 			string `gorm:"nullable"`
	websiteAddress 			string `gorm:"nullable"`
	websiteFacebook 		string `gorm:"nullable"`
	websiteX 				string `gorm:"nullable"`
	websiteInstagram 		string `gorm:"nullable"`
	websiteLinkedin 		string `gorm:"nullable"`
	websiteYoutube 			string `gorm:"nullable"`
	websiteGoogleAnalytics 	string `gorm:"nullable"`
}