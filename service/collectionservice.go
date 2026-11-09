package service

import (
	"errors"
	"tcgcollectionmanager/dto"
	"tcgcollectionmanager/persistence/entity"
	"tcgcollectionmanager/persistence/repo"
)

type CollectionService struct {
	Repo        *repo.CollectionRepo
	CardService *CardService
}

func toCollectionDTO(collection *entity.Collection) *dto.CollectionDTO {
	return &dto.CollectionDTO{
		Name:        collection.Name,
		Description: collection.Description,
		ID:          collection.ID,
		CreatedAt:   collection.CreatedAt,
		UpdatedAt:   collection.UpdatedAt,
	}
}

func (c *CollectionService) Create(data *dto.CreateCollectionDTO) (*dto.CollectionDTO, error) {
	var created entity.Collection
	created.Name = data.Name
	created.Description = data.Description
	c.Repo.Create(&created)
	return toCollectionDTO(&created), nil
}

func (c *CollectionService) GetAll() []dto.CollectionDTO {
	res := c.Repo.All()
	out := make([]dto.CollectionDTO, 0)
	for _, e := range *res {
		data := toCollectionDTO(&e)
		out = append(out, *data)
	}
	return out
}

func (c *CollectionService) Get(id *dto.CollectionIdDTO) (*dto.CollectionDTO, error) {
	fetched, err := c.Repo.Get(id.ID)
	if err != nil {
		return nil, err
	}
	return toCollectionDTO(fetched), nil
}

func (c *CollectionService) Delete(id *dto.CollectionIdDTO) error {
	_, err := c.Repo.Get(id.ID)
	if err != nil {
		return err
	}
	return c.Repo.DeleteById(id.ID)
}

func (c *CollectionService) Update(id *dto.CollectionIdDTO, data *dto.UpdateCollectionDTO) error {
	collection, err := c.Repo.Get(id.ID)
	if err != nil {
		return err
	}
	collection.Name = data.Name
	collection.Description = data.Description
	c.Repo.Update(collection)
	return nil
}

func (c *CollectionService) GetCollectionCards(id *dto.CollectionIdDTO) ([]dto.CollectionCardDTO, error) {
	collection, err := c.Repo.Get(id.ID)
	if err != nil {
		return nil, errors.New("No such collection exists")
	}
	res := make([]dto.CollectionCardDTO, 0)
	for _, card := range collection.Cards {
		collectionCard, _ := c.Repo.GetCollectionCard(collection, card.ID)
		res = append(res, dto.CollectionCardDTO{
			CardId:   card.CardId,
			ImageUrl: card.ImageUrl,
			Category: card.Category,
			Amount:   collectionCard.Amount,
			Id:       card.ID,
			Name:     card.Name,
		})
	}
	return res, nil
}

func (c *CollectionService) UpdateCollectionCard(id *dto.CollectionIdDTO, data *dto.UpdateCollectionCardDTO) error {
	collection, err := c.Repo.Get(id.ID)
	if err != nil {
		return errors.New("No such collection exists")
	}
	card, err := c.CardService.GetCard(data.CardId)
	if err != nil {
		return errors.New("No such card exists")
	}
	if data.Amount == 0 {
		c.Repo.DeleteCard(collection, card.ID)
	} else {
		c.Repo.UpdateCard(collection, card.ID, data.Amount)
	}
	return nil
}
