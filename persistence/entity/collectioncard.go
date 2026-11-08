package entity

type CollectionCard struct {
	CardID       uint `gorm:"primaryKey"`
	CollectionID uint `gorm:"primaryKey"`
	Amount       int16
}
