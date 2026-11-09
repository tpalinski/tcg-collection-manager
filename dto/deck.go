package dto

import "time"

type CreateDeckDTO struct {
	Name         string `json:"name"`
	Description  string `json:"description"`
	CollectionId uint   `json:"collectionId"`
}

type DeckDTO struct {
	ID             uint      `json:"id"`
	Name           string    `json:"name"`
	Description    string    `json:"description"`
	CollectionId   uint      `json:"collectionId"`
	CollectionName string    `json:"collectionName"`
	CreatedAt      time.Time `json:"createdAt"`
	UpdatedAt      time.Time `json:"updatedAt"`
}

type DeckIdDTO struct {
	ID uint `json:"id"`
}

type UpdateDeckDTO struct {
	Name        string `json:"name"`
	Description string `json:"description"`
}

type UpdateDeckCardDTO struct {
	CardId string `json:"cardId"`
	Amount int16  `json:"amount"`
}

type DeckCardDTO struct {
	Id           uint   `json:"internalId"`
	Amount       int16  `json:"amount"`
	Name         string `json:"name"`
	CardId       string `json:"cardId"`
	ImageUrl     string `json:"imageUrl"`
	Category     string `json:"category"`
	InCollection int16  `json:"inCollection"`
}
