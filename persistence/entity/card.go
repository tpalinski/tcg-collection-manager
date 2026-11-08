package entity

import "gorm.io/gorm"

type Card struct {
	gorm.Model
	CardId   string
	Name     string
	ImageUrl string
}
