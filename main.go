package main

import (
	"log"
	"net/http"
	"tcgcollectionmanager/app"
	"tcgcollectionmanager/persistence"
)

func main() {
	log.Println("Starting up the application")
	log.Println("Migrating the database")
	db := app.ResolveDB(&app.App)
	err := persistence.Migrate(db)
	if err != nil {
		log.Panicf("Could not migrate the database: %s \n", err.Error())
	}
	log.Println("Done")
	log.Println("Registering route handlers...")
	mux := http.NewServeMux()
	server := app.NewServer()
	server.RegisterRoutes(mux)
	handler := app.CORS(mux)
	log.Println("Done")
	log.Println("Starting to listen on port 2137")
	http.ListenAndServe("0.0.0.0:2137", handler)
}
