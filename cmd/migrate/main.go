package main

import (
	"log"
	"os"

	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"github.com/illiakornyk/e-commerce/config"
	"github.com/illiakornyk/e-commerce/db"
)

const migrationsDir = "file://cmd/migrate/migrations"

func main() {
	pgDB, err := db.NewPostgresStorage(db.Config{
		Host:     config.Envs.Host,
		Port:     config.Envs.Port,
		User:     config.Envs.User,
		Password: config.Envs.Password,
		DBName:   config.Envs.DBName,
		SSLMode:  config.Envs.SSLMode,
	})
	if err != nil {
		log.Fatal(err)
	}
	defer pgDB.Close()

	driver, err := postgres.WithInstance(pgDB, &postgres.Config{})
	if err != nil {
		log.Fatal(err)
	}

	m, err := migrate.NewWithDatabaseInstance(
		migrationsDir,
		"postgres", driver,
	)
	if err != nil {
		log.Fatal(err)
	}

	if err := m.Up(); err != nil && err != migrate.ErrNoChange {
		log.Fatal(err)
	}

	log.Println("Migrations applied successfully")

	cmd := os.Args[(len(os.Args) - 1)]
	if cmd == "up" {
		if err := m.Up(); err != nil && err != migrate.ErrNoChange {
			log.Fatal(err)
		}
	} else if cmd == "down" {
		if err := m.Down(); err != nil && err != migrate.ErrNoChange {
			log.Fatal(err)
		}
	}
}
