package customer

import (
	"fmt"
	"net/http"
	"strconv"
	"strings"

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
	mux.HandleFunc("/customers/", h.handleCustomers)
}

func (h *Handler) handleCustomers(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path == "/customers" {
    	switch r.Method {
        	case http.MethodGet:
            	h.handleGetCustomers(w, r)
        	case http.MethodPost:
            	h.handleCreateCustomer(w, r)
			default:
				utils.WriteError(w, http.StatusMethodNotAllowed, fmt.Errorf("method not allowed"))
			}
			return
    }


	pathSegments := strings.Split(r.URL.Path, "/")
    if len(pathSegments) == 3 {
        switch r.Method {
        case http.MethodGet:
            h.handleGetCustomerByID(w, r)
        // case http.MethodPatch:
        //     h.handlePatchCustomer(w, r)
		// case http.MethodDelete:
		// 	h.handleDeleteCustomer(w, r)
        default:
            utils.WriteError(w, http.StatusMethodNotAllowed, fmt.Errorf("method not allowed"))
        }
        return
    }

    utils.WriteError(w, http.StatusNotFound, fmt.Errorf("not found"))
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


	if err := payload.Validate(); err != nil {
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



func (h *Handler) handleGetCustomerByID(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		utils.WriteError(w, http.StatusMethodNotAllowed, fmt.Errorf("method not allowed"))
		return
	}

	// Extract the product ID from the URL path
	pathSegments := strings.Split(r.URL.Path, "/")
	if len(pathSegments) < 3 {
		utils.WriteError(w, http.StatusBadRequest, fmt.Errorf("customer ID is required"))
		return
	}
	productIDStr := pathSegments[2]
	customerID, err := strconv.Atoi(productIDStr)
	if err != nil {
		utils.WriteError(w, http.StatusBadRequest, fmt.Errorf("invalid customer ID"))
		return
	}

	// Retrieve the product by ID
	customer, err := h.store.GetCustomerByID(customerID)
	if err != nil {
		utils.WriteError(w, http.StatusInternalServerError, err)
		return
	}
	if customer == nil {
		utils.WriteError(w, http.StatusNotFound, fmt.Errorf("customer with ID %d not found", customerID))
		return
	}

	utils.WriteJSON(w, http.StatusOK, customer)
}


func (h *Handler) handleGetCustomers(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		utils.WriteError(w, http.StatusMethodNotAllowed, fmt.Errorf("method not allowed"))
		return
	}
	customers, err := h.store.GetCustomers()
	if err != nil {
		utils.WriteError(w, http.StatusInternalServerError, err)
		return
	}

	utils.WriteJSON(w, http.StatusOK, customers)
}
