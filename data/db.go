// data/db.go
package data

import (
	"context"
	"log"

	"expense-tracker/ent"

	_ "github.com/mattn/go-sqlite3"
)

var DB *ent.Client

func InitDB() {
	var err error
	DB, err = ent.Open("sqlite3", "file:expenses.db?cache=shared&_fk=1")
	if err != nil {
		log.Fatalf("failed opening connection to sqlite: %v", err)
	}
	// Run the auto migration tool.
	if err := DB.Schema.Create(context.Background()); err != nil {
		log.Fatalf("failed creating schema resources: %v", err)
	}
}

func GetDB() *ent.Client {
	return DB
}
