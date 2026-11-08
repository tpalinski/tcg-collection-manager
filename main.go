package main

import (
	"log"
	"tcgcollectionmanager/app"
	"tcgcollectionmanager/persistence"
	"tcgcollectionmanager/persistence/entity"
)

func main() {
	log.Println("Starting up the application")
	log.Println("Migrating the database")
	db := app.ResolveDB(&app.App)
	err := persistence.Migrate(db)
	if err != nil {
		log.Panicf("Could not migrate the database: %s \n", err.Error())
	}
	repo := app.ResolveCardRepo(&app.App)
	repo.Create(&entity.Card{
		CardId:   "stuff",
		Name:     "name",
		ImageUrl: "chuj",
	})
}
