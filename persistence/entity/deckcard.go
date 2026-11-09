package entity

type DeckCard struct {
	CardID uint `gorm:"primaryKey"`
	DeckID uint `gorm:"primaryKey"`
	Amount int16
}
