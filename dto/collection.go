package dto

type CreateCollectionDTO struct {
	Name        string
	Description string
}

type CollectionDTO struct {
	ID          uint
	Name        string
	Description string
	CreatedAt   string
	UpdatedAt   string
}

type CollectionIdDTO struct {
	ID uint
}

type UpdateCollectionDTO struct {
	Name        string
	Description string
}

type UpdateCollectionCardDTO struct {
	CardId string
	Amount int16
}
