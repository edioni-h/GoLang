package main

import (
	"context"
	"embed"
	"fmt"
	"log"

	"github.com/uptrace/bun/migrate"
)

var sqlMigrations embed.FS

var Migrations = migrate.NewMigrations()

func init() {
	if err := Migrations.Discover(sqlMigrations); err != nil {
		panic(err)
	}
}

// Migrate runs all migrations
func Migrate() {
	db, err := dbconnection()
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	migrator := migrate.NewMigrator(db, Migrations)

	err = migrator.Init(context.Background())
	if err != nil {
		log.Fatal(err)
	}

	group, err := migrator.Migrate(context.Background())
	if err != nil {
		log.Fatal(err)
	}

	if group.ID == 0 {
		fmt.Printf("there are no new migrations to run\n")
	} else {
		fmt.Printf("migrated to %s\n", group)
	}
}
