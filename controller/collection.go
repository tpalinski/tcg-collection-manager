package controller

import (
	"encoding/json"
	"errors"
	"net/http"
	"strconv"
	"tcgcollectionmanager/dto"
	"tcgcollectionmanager/service"
)

type CollectionController struct {
	Service service.CollectionService
}

func parseId(r *http.Request) (uint, error) {
	id, err := strconv.ParseUint(r.PathValue("id"), 10, 64)
	if err != nil {
		return 0, errors.New("Invalid collection id")
	}
	return uint(id), nil
}

func (c *CollectionController) PostCollection(w http.ResponseWriter, r *http.Request) {
	var body dto.CreateCollectionDTO
	if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	res, err := c.Service.Create(&body)
	if err != nil {
		http.Error(w, "Could not create collection. Why? Idk man, this app is scuffed", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)

	// We will just assume this shit is infallible for my own sanity uwu
	json.NewEncoder(w).Encode(res)
}

func (c *CollectionController) PutCollection(w http.ResponseWriter, r *http.Request) {
	id, err := parseId(r)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	var body dto.UpdateCollectionDTO
	if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}
	err = c.Service.Update(&dto.CollectionIdDTO{ID: id}, &body)
	if err != nil {
		http.Error(w, "No collection found", http.StatusNotFound)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
}

func (c *CollectionController) GetCollection(w http.ResponseWriter, r *http.Request) {
	id, err := parseId(r)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	data := dto.CollectionIdDTO{ID: id}
	collection, err := c.Service.Get(&data)
	if err != nil {
		http.Error(w, "No collection found", http.StatusNotFound)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	// We will just assume this shit is infallible for my own sanity uwu
	json.NewEncoder(w).Encode(collection)
}

func (c *CollectionController) Get(w http.ResponseWriter, r *http.Request) {
	collections := c.Service.GetAll()
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	// We will just assume this shit is infallible for my own sanity uwu
	json.NewEncoder(w).Encode(collections)
}

func (c *CollectionController) DeleteCollection(w http.ResponseWriter, r *http.Request) {
	id, err := parseId(r)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	data := dto.CollectionIdDTO{ID: id}
	err = c.Service.Delete(&data)
	if err != nil {
		http.Error(w, "No collection found", http.StatusNotFound)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusNoContent)
}

func (c *CollectionController) GetCollectionCard(w http.ResponseWriter, r *http.Request) {
	id, err := parseId(r)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	res, err := c.Service.GetCollectionCards(&dto.CollectionIdDTO{ID: id})
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	// We will just assume this shit is infallible for my own sanity uwu
	json.NewEncoder(w).Encode(res)
}

func (c *CollectionController) PutCollectionCard(w http.ResponseWriter, r *http.Request) {
	id, err := parseId(r)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	var body dto.UpdateCollectionCardDTO
	if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	err = c.Service.UpdateCollectionCard(&dto.CollectionIdDTO{ID: id}, &body)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
}
