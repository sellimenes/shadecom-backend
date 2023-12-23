package models

import "gorm.io/gorm"

type Settings struct {
	gorm.Model
	WebsiteName 			string
	WebsiteDescription 		string
	WebsiteLogo 			string
	WebsiteFavicon 			string
	WebsiteKeywords 		string
	WebsiteEmail 			string
	WebsitePhone 			string `gorm:"nullable"`
	WebsiteAddress 			string `gorm:"nullable"`
	WebsiteFacebook 		string `gorm:"nullable"`
	WebsiteX 				string `gorm:"nullable"`
	WebsiteInstagram 		string `gorm:"nullable"`
	WebsiteLinkedin 		string `gorm:"nullable"`
	WebsiteYoutube 			string `gorm:"nullable"`
	WebsiteGoogleAnalytics 	string `gorm:"nullable"`
}