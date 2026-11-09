package controller

import (
	"encoding/json"
	"errors"
	"net/http"
	"strconv"
	"tcgcollectionmanager/dto"
	"tcgcollectionmanager/service"
)

type DeckController struct {
	Service *service.DeckService
}

func parseDeckId(r *http.Request) (uint, error) {
	id, err := strconv.ParseUint(r.PathValue("deckId"), 10, 64)
	if err != nil {
		return 0, errors.New("Invalid deck id")
	}
	return uint(id), nil
}

func (c *DeckController) PostDeck(w http.ResponseWriter, r *http.Request) {
	id, err := parseId(r)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	var body dto.CreateDeckRequest
	if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	data := dto.CreateDeckDTO{
		Name:         body.Name,
		Description:  body.Description,
		CollectionId: id,
	}
	res, err := c.Service.Create(&data)
	if err != nil {
		http.Error(w, "Could not create deck. No such parent collection exists", http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)

	// We will just assume this shit is infallible for my own sanity uwu
	json.NewEncoder(w).Encode(res)
}

func (c *DeckController) PutDeck(w http.ResponseWriter, r *http.Request) {
	id, err := parseDeckId(r)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	var body dto.UpdateDeckDTO
	if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}
	err = c.Service.Update(&dto.DeckIdDTO{ID: id}, &body)
	if err != nil {
		http.Error(w, "No deck found", http.StatusNotFound)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
}

func (c *DeckController) GetDeck(w http.ResponseWriter, r *http.Request) {
	id, err := parseDeckId(r)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	data := dto.DeckIdDTO{ID: id}
	collection, err := c.Service.Get(&data)
	if err != nil {
		http.Error(w, "No deck found", http.StatusNotFound)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	// We will just assume this shit is infallible for my own sanity uwu
	json.NewEncoder(w).Encode(collection)
}

func (c *DeckController) Get(w http.ResponseWriter, r *http.Request) {
	id, err := parseId(r)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	data := dto.CollectionIdDTO{ID: id}
	decks, err := c.Service.GetAllForCollection(data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	// We will just assume this shit is infallible for my own sanity uwu
	json.NewEncoder(w).Encode(decks)
}

func (c *DeckController) DeleteDeck(w http.ResponseWriter, r *http.Request) {
	id, err := parseDeckId(r)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	data := dto.DeckIdDTO{ID: id}
	err = c.Service.Delete(&data)
	if err != nil {
		http.Error(w, "No deck found", http.StatusNotFound)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusNoContent)
}

func (c *DeckController) GetDeckCard(w http.ResponseWriter, r *http.Request) {
	id, err := parseId(r)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	res, err := c.Service.GetDeckCards(&dto.DeckIdDTO{ID: id})
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	// We will just assume this shit is infallible for my own sanity uwu
	json.NewEncoder(w).Encode(res)
}

func (c *DeckController) PutDeckCard(w http.ResponseWriter, r *http.Request) {
	id, err := parseId(r)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	var body dto.UpdateDeckCardDTO
	if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	err = c.Service.UpdateDeckCard(&dto.DeckIdDTO{ID: id}, &body)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
}
