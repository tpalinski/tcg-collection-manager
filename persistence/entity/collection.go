package entity

import (
	"gorm.io/gorm"
)

type Collection struct {
	gorm.Model
	Name        string
	Description string
	Cards       []Card `gorm:"many2many:collection_cards"`
}
