package repo

import (
	"tcgcollectionmanager/persistence/entity"

	"gorm.io/gorm"
)

type CardRepo struct {
	DB *gorm.DB
}

func (c *CardRepo) Create(card *entity.Card) {
	c.DB.Create(card)
}

func (c *CardRepo) GetByCardId(id string) *entity.Card {
	var res entity.Card
	c.DB.Model(&entity.Card{}).First(&res, "card_id = ?", id)
	return &res
}

func (c *CardRepo) Get(id uint) *entity.Card {
	var res entity.Card
	c.DB.Model(&entity.Card{}).First(&res, id)
	return &res
}
