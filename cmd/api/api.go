package api

import (
	"database/sql"
	"log"
	"net/http"
	"github.com/illiakornyk/e-commerce/services/user"
)

type APIServer struct {
	listenAddress string
	databaseConn *sql.DB
}
func NewApiServerInstance(add string, db *sql.DB) *APIServer {
	return &APIServer{listenAddress: add, databaseConn: db}
}


func (s *APIServer) Run() error {
	mux := http.NewServeMux()

	// Creating a subrouter for /api/v1
	apiV1Mux := http.NewServeMux()
	mux.Handle("/api/v1/", http.StripPrefix("/api/v1", apiV1Mux))

	userHandler := user.NewHandler()
	userHandler.RegisterRoutes(apiV1Mux)

	log.Println("Starting server on port", s.listenAddress)
	return http.ListenAndServe(s.listenAddress, mux)
}
