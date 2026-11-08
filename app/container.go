package app

import (
	"log"
	"tcgcollectionmanager/persistence/entity"
	"tcgcollectionmanager/persistence/repo"
	"tcgcollectionmanager/service"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

const dbName = "tcgmanager.sqlite"

type Container struct {
	// Register everything here (just the types i guess)
	DB                *gorm.DB
	CardRepo          *repo.CardRepo
	CollectionRepo    *repo.CollectionRepo
	CollectionService *service.CollectionService
}

func ResolveDB(c *Container) *gorm.DB {
	if c.DB == nil {
		log.Println("Initializing new db connection")
		db, err := gorm.Open(sqlite.Open(dbName), &gorm.Config{})
		if err != nil {
			log.Fatalf("Could not open the database under %s", dbName)
			panic("Could not open db, exiting")
		}
		db.SetupJoinTable(&entity.Collection{}, "Cards", &entity.CollectionCard{})
		c.DB = db
	}
	return c.DB
}

func ResolveCardRepo(c *Container) *repo.CardRepo {
	if c.CardRepo == nil {
		c.CardRepo = &repo.CardRepo{DB: ResolveDB(c)}
	}
	return c.CardRepo
}

func ResolveCollectionRepo(c *Container) *repo.CollectionRepo {
	if c.CollectionRepo == nil {
		c.CollectionRepo = &repo.CollectionRepo{DB: ResolveDB(c)}
	}
	return c.CollectionRepo
}

func ResolveCollectionService(c *Container) *service.CollectionService {
	if c.CollectionService == nil {
		c.CollectionService = &service.CollectionService{Repo: ResolveCollectionRepo(c)}
	}
	return c.CollectionService
}

var App Container
