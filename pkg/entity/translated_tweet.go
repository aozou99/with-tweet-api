package entity

import "gorm.io/gorm"

type TranslatedTweet struct {
	gorm.Model
	ID             string `gorm:"primarykey"`
	OriginText     string `gorm:"column:origin_text;NOT NULL"`
	TranslatedText string `gorm:"column:translated_text;NOT NULL"`
}
