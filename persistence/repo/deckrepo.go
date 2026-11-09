package repo

import (
	"tcgcollectionmanager/persistence/entity"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type DeckRepo struct {
	DB *gorm.DB
}

func (c *DeckRepo) Create(deck *entity.Deck) {
	c.DB.Create(deck)
}

func (c *DeckRepo) Get(id uint) (*entity.Deck, error) {
	var res entity.Deck
	result := c.DB.Preload("Cards").Model(&entity.Deck{}).First(&res, id)
	return &res, result.Error
}

func (c *DeckRepo) All() *[]entity.Deck {
	var decks []entity.Deck
	c.DB.Find(&decks)
	return &decks
}

func (c *DeckRepo) CollectionAll(collection *entity.Collection) *[]entity.Deck {
	var decks []entity.Deck
	c.DB.Find(&decks, "collection_id = ?", collection.ID)
	return &decks
}

func (c *DeckRepo) DeleteById(id uint) error {
	result := c.DB.Delete(&entity.Deck{}, id)
	return result.Error
}

func (c *DeckRepo) Update(deck *entity.Deck) {
	c.DB.Save(deck)
}

func (c *DeckRepo) UpdateCard(deck *entity.Deck, cardId uint, newAmount int16) error {
	mapping := entity.DeckCard{
		DeckID: deck.ID,
		CardID: cardId,
		Amount: newAmount,
	}

	return c.DB.Clauses(clause.OnConflict{
		Columns: []clause.Column{
			{Name: "deck_id"},
			{Name: "card_id"},
		},
		DoUpdates: clause.AssignmentColumns([]string{"amount"}),
	}).Create(&mapping).Error
}

func (c *DeckRepo) DeleteCard(deck *entity.Deck, cardId uint) {
	c.DB.Delete(&entity.DeckCard{}, "card_id = ? AND deck_id = ?", cardId, deck.ID)
}

func (c *DeckRepo) GetDeckCard(deck *entity.Deck, cardId uint) (*entity.DeckCard, error) {
	var mapping entity.DeckCard
	res := c.DB.First(&mapping, "card_id = ? AND deck_id = ?", cardId, deck.ID)
	return &mapping, res.Error
}
