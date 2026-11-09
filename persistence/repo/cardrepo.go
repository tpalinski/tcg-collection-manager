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

func (c *CardRepo) GetByCardId(id string) (*entity.Card, error) {
	var res entity.Card
	q := c.DB.Model(&entity.Card{}).First(&res, "card_id = ?", id)
	return &res, q.Error
}

func (c *CardRepo) Get(id uint) (*entity.Card, error) {
	var res entity.Card
	q := c.DB.Model(&entity.Card{}).First(&res, id)
	return &res, q.Error
}
