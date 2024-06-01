package customer

import (
	"fmt"
	"net/http"

	"github.com/illiakornyk/e-commerce/types"
	"github.com/illiakornyk/e-commerce/utils"
)

type Handler struct {
	store types.CustomerStore
}

func NewHandler(store types.CustomerStore) *Handler {
	return &Handler{store: store}
}

func (h *Handler) RegisterRoutes(mux *http.ServeMux) {
    mux.HandleFunc("/customers", h.handleCustomers)
}

func (h *Handler) handleCustomers(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path == "/customers" {
    	switch r.Method {
        	case http.MethodGet:
            	// h.handleGetProducts(w, r)
        	case http.MethodPost:
            	h.handleCreateCustomer(w, r)
			default:
				utils.WriteError(w, http.StatusMethodNotAllowed, fmt.Errorf("method not allowed"))
			}
			return
    }
}


func (h *Handler) handleCreateCustomer(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		utils.WriteError(w, http.StatusMethodNotAllowed, fmt.Errorf("method not allowed"))
		return
	}

	var payload types.CreateCustomerPayload
	if err := utils.ParseJSON(r, &payload); err != nil {
		utils.WriteError(w, http.StatusBadRequest, err)
		return
	}


	err := h.store.CreateCustomer(payload)
	if err != nil {
		utils.WriteError(w, http.StatusInternalServerError, err)
		return
	}

	utils.WriteJSON(w, http.StatusCreated, map[string]string{"status": "customer created"})
}
