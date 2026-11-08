package tcgcollectionmanager

import (
	"context"
	"log"
	"tcgcollectionmanager/app"
)

func main() {
	log.Println("Starting up the application");
  log.Println("Migrating the database");
  db := app.ResolveDB(&app.App);
  ctx := context.Background();
}
