package service

import (
	"errors"
	"tcgcollectionmanager/dto"
	"tcgcollectionmanager/persistence/entity"
	"tcgcollectionmanager/persistence/repo"
)

type DeckService struct {
	Repo           *repo.DeckRepo
	CollectionRepo *repo.CollectionRepo
	CardService    *CardService
}

func toDeckDTO(deck *entity.Deck) *dto.DeckDTO {
	return &dto.DeckDTO{
		Name:           deck.Name,
		Description:    deck.Description,
		ID:             deck.ID,
		CreatedAt:      deck.CreatedAt,
		UpdatedAt:      deck.UpdatedAt,
		CollectionId:   deck.CollectionID,
		CollectionName: deck.Collection.Name,
	}
}

func (c *DeckService) Create(data *dto.CreateDeckDTO) (*dto.DeckDTO, error) {
	_, err := c.CollectionRepo.Get(data.CollectionId)
	if err != nil {
		return nil, errors.New("No such collection exists")
	}
	var created entity.Deck
	created.Name = data.Name
	created.Description = data.Description
	created.CollectionID = data.CollectionId
	c.Repo.Create(&created)
	return toDeckDTO(&created), nil
}

func (c *DeckService) GetAll() []dto.DeckDTO {
	res := c.Repo.All()
	out := make([]dto.DeckDTO, 0)
	for _, e := range *res {
		data := toDeckDTO(&e)
		out = append(out, *data)
	}
	return out
}

func (c *DeckService) GetAllForCollection(collection dto.CollectionIdDTO) ([]dto.DeckDTO, error) {
	col, err := c.CollectionRepo.Get(collection.ID)
	if err != nil {
		return nil, errors.New("No such collection exists")
	}
	res := c.Repo.CollectionAll(col)
	out := make([]dto.DeckDTO, 0)
	for _, e := range *res {
		data := toDeckDTO(&e)
		out = append(out, *data)
	}
	return out, nil
}

func (c *DeckService) Get(id *dto.DeckIdDTO) (*dto.DeckDTO, error) {
	fetched, err := c.Repo.Get(id.ID)
	if err != nil {
		return nil, err
	}
	return toDeckDTO(fetched), nil
}

func (c *DeckService) Delete(id *dto.DeckIdDTO) error {
	_, err := c.Repo.Get(id.ID)
	if err != nil {
		return err
	}
	return c.Repo.DeleteById(id.ID)
}

func (c *DeckService) Update(id *dto.DeckIdDTO, data *dto.UpdateDeckDTO) error {
	deck, err := c.Repo.Get(id.ID)
	if err != nil {
		return err
	}
	deck.Name = data.Name
	deck.Description = data.Description
	c.Repo.Update(deck)
	return nil
}

func (c *DeckService) GetDeckCards(id *dto.DeckIdDTO) ([]dto.DeckCardDTO, error) {
	deck, err := c.Repo.Get(id.ID)
	if err != nil {
		return nil, errors.New("No such deck exists")
	}
	res := make([]dto.DeckCardDTO, 0)
	for _, card := range deck.Cards {
		deckCard, _ := c.Repo.GetDeckCard(deck, card.ID)
		collectionCard, err := c.CollectionRepo.GetCollectionCard(&deck.Collection, card.ID)
		var inCollection int16
		if err != nil {
			inCollection = 0
		} else {
			inCollection = collectionCard.Amount
		}
		res = append(res, dto.DeckCardDTO{
			CardId:       card.CardId,
			ImageUrl:     card.ImageUrl,
			Category:     card.Category,
			Amount:       deckCard.Amount,
			Id:           card.ID,
			Name:         card.Name,
			InCollection: inCollection,
		})
	}
	return res, nil
}

func (c *DeckService) UpdateDeckCard(id *dto.DeckIdDTO, data *dto.UpdateDeckCardDTO) error {
	deck, err := c.Repo.Get(id.ID)
	if err != nil {
		return errors.New("No such deck exists")
	}
	card, err := c.CardService.GetCard(data.CardId)
	if err != nil {
		return errors.New("No such card exists")
	}
	if data.Amount == 0 {
		c.Repo.DeleteCard(deck, card.ID)
	} else {
		c.Repo.UpdateCard(deck, card.ID, data.Amount)
	}
	return nil
}
