package app

import (
	"net/http"
	"tcgcollectionmanager/controller"
)

type Server struct {
	CollectionController *controller.CollectionController
}

func NewServer() *Server {
	return &Server{
		CollectionController: ResolveCollectionController(&App),
	}
}

func (s *Server) RegisterRoutes(mux *http.ServeMux) {
	mux.HandleFunc("GET /collection/{id}/cards", s.CollectionController.GetCollectionCard)
	mux.HandleFunc("PUT /collection/{id}/cards", s.CollectionController.PutCollectionCard)
	mux.HandleFunc("GET /collection/{id}", s.CollectionController.GetCollection)
	mux.HandleFunc("DELETE /collection/{id}", s.CollectionController.DeleteCollection)
	mux.HandleFunc("PUT /collection/{id}", s.CollectionController.PutCollection)
	mux.HandleFunc("POST /collection", s.CollectionController.PostCollection)
	mux.HandleFunc("GET /collection", s.CollectionController.Get)
}
