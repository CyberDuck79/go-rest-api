package database

import (
	"context"
	"log"

	_ "github.com/mattn/go-sqlite3"
	"rest_api/ent"
)

func Init() (client *ent.Client) {
	client, err := ent.Open("sqlite3", "file:ent?mode=memory&cache=shared&_fk=1")
	if err != nil {
		log.Fatalf("failed opening connection to sqlite: %v", err)
		return nil
	}
	log.Println("Database: connected")
	if err := client.Schema.Create(context.Background()); err != nil {
		log.Fatalf("failed creating schema resources: %v", err)
		return nil
	}
	log.Println("Database: schema created")
	return
}
