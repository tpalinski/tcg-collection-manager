package app

import (
	"context"
	"log"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

const dbName = "tcgmanager.sqlite";

type Container struct {
  // Register everything here (just the types i guess)
  DB *gorm.DB;
}

func ResolveDB(c *Container) *gorm.DB {
  if (c.DB == nil) {
    db, err := gorm.Open(sqlite.Open(dbName), &gorm.Config{});
    if (err != nil) {
      log.Fatalf("Could not open the database under %s", dbName);
      panic("Could not open db, exiting");
    }
    c.DB = db;
  }
  return c.DB;
}

var App Container;



