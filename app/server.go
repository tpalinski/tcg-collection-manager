package app

import (
	"net/http"
	"tcgcollectionmanager/controller"
)

type Server struct {
	CollectionController *controller.CollectionController
	DeckController       *controller.DeckController
}

func CORS(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, PATCH, DELETE, OPTIONS")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")

		if r.Method == http.MethodOptions {
			w.WriteHeader(http.StatusNoContent)
			return
		}

		next.ServeHTTP(w, r)
	})
}

func NewServer() *Server {
	return &Server{
		CollectionController: ResolveCollectionController(&App),
		DeckController:       ResolveDeckController(&App),
	}
}

func (s *Server) RegisterRoutes(mux *http.ServeMux) {
	// Collection related routes
	mux.HandleFunc("GET /collection/{id}/cards", s.CollectionController.GetCollectionCard)
	mux.HandleFunc("PUT /collection/{id}/cards", s.CollectionController.PutCollectionCard)
	mux.HandleFunc("GET /collection/{id}", s.CollectionController.GetCollection)
	mux.HandleFunc("DELETE /collection/{id}", s.CollectionController.DeleteCollection)
	mux.HandleFunc("PUT /collection/{id}", s.CollectionController.PutCollection)
	mux.HandleFunc("POST /collection", s.CollectionController.PostCollection)
	mux.HandleFunc("GET /collection", s.CollectionController.Get)
	// Deck related routes
	mux.HandleFunc("PUT /collection/{id}/deck/{deckId}/cards", s.DeckController.PutDeckCard)
	mux.HandleFunc("GET /collection/{id}/deck/{deckId}/cards", s.DeckController.GetDeckCard)
	mux.HandleFunc("GET /collection/{id}/deck/{deckId}", s.DeckController.GetDeck)
	mux.HandleFunc("PUT /collection/{id}/deck/{deckId}", s.DeckController.PutDeck)
	mux.HandleFunc("DELETE /collection/{id}/deck/{deckId}", s.DeckController.DeleteDeck)
	mux.HandleFunc("GET /collection/{id}/deck", s.DeckController.Get)
	mux.HandleFunc("POST /collection/{id}/deck", s.DeckController.PostDeck)
}
