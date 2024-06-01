package seller

import (
	"fmt"
	"net/http"
	"strconv"
	"strings"

	"github.com/illiakornyk/e-commerce/services/auth"
	"github.com/illiakornyk/e-commerce/types"
	"github.com/illiakornyk/e-commerce/utils"
)

type Handler struct {
	store types.SellerStore
	userStore  types.UserStore
}

func NewHandler(store types.SellerStore, userStore types.UserStore) *Handler {
	return &Handler{store: store, userStore: userStore}
}

func (h *Handler) RegisterRoutes(mux *http.ServeMux) {


    mux.HandleFunc("/sellers", auth.WithJWTAuth(h.handleSellers, h.userStore))
    mux.HandleFunc("/sellers/", auth.WithJWTAuth(h.handleSellers, h.userStore))
}

func (h *Handler) handleSellers(w http.ResponseWriter, r *http.Request) {
    if r.URL.Path == "/sellers" {
        switch r.Method {
        case http.MethodGet:
            h.handleGetSellers(w, r)
        case http.MethodPost:
            h.handleCreateSeller(w, r)
        default:
            utils.WriteError(w, http.StatusMethodNotAllowed, fmt.Errorf("method not allowed"))
        }
        return
    }

    pathSegments := strings.Split(r.URL.Path, "/")
    if len(pathSegments) == 3 {
        switch r.Method {
        case http.MethodGet:
            h.handleGetSellerByID(w, r)
        case http.MethodPatch:
            // h.handlePatchProduct(w, r)
		case http.MethodDelete:
			// h.handleDeleteProduct(w, r)
        default:
            utils.WriteError(w, http.StatusMethodNotAllowed, fmt.Errorf("method not allowed"))
        }
        return
    }

    utils.WriteError(w, http.StatusNotFound, fmt.Errorf("not found"))
}






func (h *Handler) handleCreateSeller(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		utils.WriteError(w, http.StatusMethodNotAllowed, fmt.Errorf("method not allowed"))
		return
	}

	var payload types.CreateSellerPayload
	if err := utils.ParseJSON(r, &payload); err != nil {
		utils.WriteError(w, http.StatusBadRequest, err)
		return
	}

	if err := payload.Validate(); err != nil {
		utils.WriteError(w, http.StatusBadRequest, err)
		return
	}


	err := h.store.CreateSeller(payload)
	if err != nil {
		utils.WriteError(w, http.StatusInternalServerError, err)
		return
	}

	utils.WriteJSON(w, http.StatusCreated, map[string]string{"status": "seller created"})
}


func (h *Handler) handleGetSellerByID(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		utils.WriteError(w, http.StatusMethodNotAllowed, fmt.Errorf("method not allowed"))
		return
	}

	// Extract the product ID from the URL path
	pathSegments := strings.Split(r.URL.Path, "/")
	if len(pathSegments) < 3 {
		utils.WriteError(w, http.StatusBadRequest, fmt.Errorf("seller ID is required"))
		return
	}
	sellerIDStr := pathSegments[2]
	sellerID, err := strconv.Atoi(sellerIDStr)
	if err != nil {
		utils.WriteError(w, http.StatusBadRequest, fmt.Errorf("invalid seller ID"))
		return
	}

	// Retrieve the product by ID
	seller, err := h.store.GetSellerByID(sellerID)
	if err != nil {
		utils.WriteError(w, http.StatusInternalServerError, err)
		return
	}
	if seller == nil {
		utils.WriteError(w, http.StatusNotFound, fmt.Errorf("seller with ID %d not found", sellerID))
		return
	}

	utils.WriteJSON(w, http.StatusOK, seller)
}


func (h *Handler) handleGetSellers(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		utils.WriteError(w, http.StatusMethodNotAllowed, fmt.Errorf("method not allowed"))
		return
	}
	sellers, err := h.store.GetSellers()
	if err != nil {
		utils.WriteError(w, http.StatusInternalServerError, err)
		return
	}

	utils.WriteJSON(w, http.StatusOK, sellers)
}
