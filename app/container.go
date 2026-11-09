package app

import (
	"log"
	"net/http"
	"tcgcollectionmanager/client"
	"tcgcollectionmanager/controller"
	"tcgcollectionmanager/persistence/entity"
	"tcgcollectionmanager/persistence/repo"
	"tcgcollectionmanager/service"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

const dbName = "tcgmanager.sqlite"
const tcgDexBaseUrl = "https://api.tcgdex.net/v2/en"

// Register everything here (just the types i guess)
type Container struct {
	DB                   *gorm.DB
	Client               *client.TCGDexClient
	CardRepo             *repo.CardRepo
	CardService          *service.CardService
	CollectionRepo       *repo.CollectionRepo
	CollectionService    *service.CollectionService
	CollectionController *controller.CollectionController
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

func ResolveCardService(c *Container) *service.CardService {
	if c.CardService == nil {
		c.CardService = &service.CardService{
			Client: ResolveClient(c),
			Repo:   ResolveCardRepo(c),
		}
	}
	return c.CardService
}

func ResolveCollectionRepo(c *Container) *repo.CollectionRepo {
	if c.CollectionRepo == nil {
		c.CollectionRepo = &repo.CollectionRepo{DB: ResolveDB(c)}
	}
	return c.CollectionRepo
}

func ResolveCollectionService(c *Container) *service.CollectionService {
	if c.CollectionService == nil {
		c.CollectionService = &service.CollectionService{
			Repo:        ResolveCollectionRepo(c),
			CardService: ResolveCardService(c),
		}
	}
	return c.CollectionService
}

func ResolveCollectionController(c *Container) *controller.CollectionController {
	if c.CollectionController == nil {
		c.CollectionController = &controller.CollectionController{Service: *ResolveCollectionService(c)}
	}
	return c.CollectionController
}

func ResolveClient(c *Container) *client.TCGDexClient {
	if c.Client == nil {
		c.Client = &client.TCGDexClient{
			Client:  &http.Client{},
			BaseUrl: tcgDexBaseUrl,
		}
	}
	return c.Client
}

var App Container
