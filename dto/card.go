package dto

type CardDTO struct {
	ID       uint   `json:"id"`
	Name     string `json:"name"`
	CardId   string `json:"cardId"`
	ImageUrl string `json:"imageUrl"`
	Category string `json:"category"`
}
