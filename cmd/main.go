package main

import (
	"database/sql"
	"log"

	"github.com/illiakornyk/e-commerce/cmd/api"
	"github.com/illiakornyk/e-commerce/config"
	"github.com/illiakornyk/e-commerce/db"
)

func main() {
	db, err := db.NewPostgresStorage(db.Config{
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

	initStorage(db)


	server := api.NewApiServerInstance(":8080", nil)
	if err := server.Run(); err != nil {
		log.Fatal(err)
	}
}

func initStorage(db *sql.DB) {
	err := db.Ping()

	if err != nil {
		log.Fatal(err)
	}

	log.Println("Connected to Postgres!")
}
