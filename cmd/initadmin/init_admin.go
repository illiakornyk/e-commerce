package main

import (
	"log"

	"github.com/illiakornyk/e-commerce/config"
	"github.com/illiakornyk/e-commerce/db"
	"github.com/illiakornyk/e-commerce/services/auth"
	"github.com/illiakornyk/e-commerce/services/user"
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

	userStore := user.NewStore(db)

	result, err := auth.CreateAdminUser(userStore, config.Envs.ADMIN_USERNAME, config.Envs.ADMIN_PASSWORD, config.Envs.ADMIN_EMAIL)
	if err != nil {
		log.Fatal(err)
	}

	switch result {
	case auth.AdminCreated:
		log.Println("Admin user created successfully")
	case auth.AdminAlreadyExists:
		log.Println("Admin user already exists")
	default:
		log.Println("Unknown result from CreateAdminUser")
	}

}
