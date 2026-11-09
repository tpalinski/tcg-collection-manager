package dto

type CardDTO struct {
	ID       uint   `json:"id"`
	Name     string `json:"name"`
	CardId   string `json:"cardId"`
	ImageUrl string `json:"imageUrl"`
	Category string `json:"category"`
}

type MissingCardDTO struct {
	Id           uint   `json:"internalId"`
	Amount       int16  `json:"amount"`
	InCollection int16  `json:"inCollection"`
	Name         string `json:"name"`
	CardId       string `json:"cardId"`
	ImageUrl     string `json:"imageUrl"`
	Category     string `json:"category"`
}

type DeckCardMissingDTO struct {
	Missing int16 `json:"amountMissing"`
	DeckID  uint  `json:"deckId"`
}
