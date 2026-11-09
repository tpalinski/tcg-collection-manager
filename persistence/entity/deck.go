package entity

import "gorm.io/gorm"

type Deck struct {
	gorm.Model
	Name         string
	Description  string
	Cards        []Card `gorm:"many2many:deck_cards"`
	CollectionID uint
	Collection   Collection
}
