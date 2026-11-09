package client

import (
	"encoding/json"
	"fmt"
	"net/http"
	"tcgcollectionmanager/dto"
)

type TCGDexClient struct {
	BaseUrl string
	Client  *http.Client
}

func (c *TCGDexClient) FindCard(cardId string) (*dto.PokemonCardDTO, error) {
	url := fmt.Sprintf("%s/cards/%s", c.BaseUrl, cardId)

	resp, err := c.Client.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("TCGdex returned status %d", resp.StatusCode)
	}

	var card dto.PokemonCardDTO
	if err := json.NewDecoder(resp.Body).Decode(&card); err != nil {
		return nil, err
	}

	return &card, nil
}
