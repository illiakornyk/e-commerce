package main

import (
	"log"

	"github.com/illiakornyk/e-commerce/cmd/api"
)

func main() {
	server := api.NewApiServerInstance(":8080", nil)
	if err := server.Run(); err != nil {
		log.Fatal(err)
	}
}
