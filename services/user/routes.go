package user

import (
	"net/http"
)

type Handler struct {
}

func NewHandler() *Handler {
	return &Handler{}
}

func (h *Handler) RegisterRoutes(mux *http.ServeMux) {
	mux.HandleFunc("/login", h.handleLogin)
	mux.HandleFunc("/register", h.handleRegister)
}

func (h *Handler) handleLogin(w http.ResponseWriter, r *http.Request) {
	// handle login logic here
}

func (h *Handler) handleRegister(w http.ResponseWriter, r *http.Request) {
	// handle register logic here

}
