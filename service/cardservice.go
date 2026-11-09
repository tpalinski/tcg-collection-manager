package service

import (
	"errors"
	"tcgcollectionmanager/client"
	"tcgcollectionmanager/dto"
	"tcgcollectionmanager/persistence/entity"
	"tcgcollectionmanager/persistence/repo"

	"gorm.io/gorm"
)

type CardService struct {
	Repo   *repo.CardRepo
	Client *client.TCGDexClient
}

func (s *CardService) GetCard(cardId string) (*dto.CardDTO, error) {
	card, err := s.Repo.GetByCardId(cardId)
	var res dto.CardDTO
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			external, err := s.Client.FindCard(cardId)
			if err != nil {
				return nil, errors.New("No such card exists")
			}
			cacheTarget := entity.Card{
				CardId:   cardId,
				Name:     external.Name,
				Category: external.Category,
				ImageUrl: external.Image,
			}
			s.Repo.Create(&cacheTarget)
			// just create the dto on the fly not to query the db twice
			res.CardId = cacheTarget.CardId
			res.ID = cacheTarget.ID
			res.ImageUrl = cacheTarget.ImageUrl
			res.Name = cacheTarget.Name
			res.Category = cacheTarget.Category
		}
	} else {
		res.CardId = card.CardId
		res.ID = card.ID
		res.Name = card.Name
		res.ImageUrl = card.ImageUrl
		res.Category = card.Category
	}
	return &res, nil
}
