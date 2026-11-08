package repo

import (
	"gorm.io/gorm"
	"tcgcollectionmanager/persistence/entity"
)

type CollectionRepo struct {
	DB *gorm.DB
}

func (c *CollectionRepo) Create(collection *entity.Collection) {
	c.DB.Create(collection)
}

func (c *CollectionRepo) Get(id uint) (*entity.Collection, error) {
	var res entity.Collection
	result := c.DB.Model(&entity.Collection{}).First(&res, id)
	return &res, result.Error
}

func (c *CollectionRepo) All() *[]entity.Collection {
	var collections []entity.Collection
	c.DB.Find(&collections)
	return &collections
}

func (c *CollectionRepo) DeleteById(id uint) error {
	result := c.DB.Delete(&entity.Collection{}, id)
	return result.Error
}

func (c *CollectionRepo) Update(collection *entity.Collection) {
	c.DB.Save(collection)
}

func (c *CollectionRepo) UpdateCard(collection *entity.Collection, cardId uint, newAmount int16) {
	var mapping entity.CollectionCard
	c.DB.FirstOrCreate(&mapping, "card_id = ? AND collection_id = ?", cardId, collection.ID)
	mapping.Amount = newAmount
	c.DB.Save(&mapping)
}
