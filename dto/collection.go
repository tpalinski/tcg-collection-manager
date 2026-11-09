package dto

import "time"

type CreateCollectionDTO struct {
	Name        string `json:"name"`
	Description string `json:"description"`
}

type CollectionDTO struct {
	ID          uint      `json:"id"`
	Name        string    `json:"name"`
	Description string    `json:"description"`
	CreatedAt   time.Time `json:"createdAt"`
	UpdatedAt   time.Time `json:"updatedAt"`
}

type CollectionIdDTO struct {
	ID uint `json:"id"`
}

type UpdateCollectionDTO struct {
	Name        string `json:"name"`
	Description string `json:"description"`
}

type UpdateCollectionCardDTO struct {
	CardId string `json:"cardId"`
	Amount int16  `json:"amount"`
}

type CollectionCardDTO struct {
	Id       uint   `json:"internalId"`
	Amount   int16  `json:"amount"`
	Name     string `json:"name"`
	CardId   string `json:"cardId"`
	ImageUrl string `json:"imageUrl"`
	Category string `json:"category"`
}
