package dto

type CreateDeckRequest struct {
	Name        string `json:"name"`
	Description string `json:"description"`
}
