package persistence

import (
	"gorm.io/gorm"
	"tcgcollectionmanager/persistence/entity"
)

func Migrate(db *gorm.DB) error {
	err := db.AutoMigrate(&entity.Card{}, &entity.Collection{}, &entity.Deck{})
	if err != nil {
		return err
	}
	return nil
}
