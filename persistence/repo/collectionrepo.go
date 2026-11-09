package repo

import (
	"tcgcollectionmanager/persistence/entity"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type CollectionRepo struct {
	DB *gorm.DB
}

func (c *CollectionRepo) Create(collection *entity.Collection) {
	c.DB.Create(collection)
}

func (c *CollectionRepo) Get(id uint) (*entity.Collection, error) {
	var res entity.Collection
	result := c.DB.Preload("Cards").Model(&entity.Collection{}).First(&res, id)
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

func (c *CollectionRepo) UpdateCard(collection *entity.Collection, cardId uint, newAmount int16) error {
	mapping := entity.CollectionCard{
		CollectionID: collection.ID,
		CardID:       cardId,
		Amount:       newAmount,
	}

	return c.DB.Clauses(clause.OnConflict{
		Columns: []clause.Column{
			{Name: "collection_id"},
			{Name: "card_id"},
		},
		DoUpdates: clause.AssignmentColumns([]string{"amount"}),
	}).Create(&mapping).Error
}

func (c *CollectionRepo) DeleteCard(collection *entity.Collection, cardId uint) {
	c.DB.Delete(&entity.CollectionCard{}, "card_id = ? AND collection_id = ?", cardId, collection.ID)
}

func (c *CollectionRepo) GetCollectionCard(collection *entity.Collection, cardId uint) (*entity.CollectionCard, error) {
	var mapping entity.CollectionCard
	res := c.DB.First(&mapping, "card_id = ? AND collection_id = ?", cardId, collection.ID)
	return &mapping, res.Error
}
