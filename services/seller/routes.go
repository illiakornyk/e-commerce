package seller

import (
	"fmt"
	"net/http"

	"github.com/illiakornyk/e-commerce/types"
	"github.com/illiakornyk/e-commerce/utils"
)

type Handler struct {
	store types.SellerStore
}

func NewHandler(store types.SellerStore) *Handler {
	return &Handler{store: store}
}

func (h *Handler) RegisterRoutes(mux *http.ServeMux) {
    mux.HandleFunc("/sellers", h.handleSellers)
    mux.HandleFunc("/sellers/", h.handleSellers)
}

func (h *Handler) handleSellers(w http.ResponseWriter, r *http.Request) {
    // Check if the path is exactly "/products", which means we're handling the collection of products
    if r.URL.Path == "/sellers" {
        switch r.Method {
        case http.MethodGet:
            // h.handleGetProducts(w, r)
        case http.MethodPost:
            h.handleCreateSeller(w, r)
        default:
            utils.WriteError(w, http.StatusMethodNotAllowed, fmt.Errorf("method not allowed"))
        }
        return
    }

    // pathSegments := strings.Split(r.URL.Path, "/")
    // if len(pathSegments) == 3 {
    //     switch r.Method {
    //     case http.MethodGet:
    //         h.handleGetProductByID(w, r)
    //     case http.MethodPatch:
    //         h.handlePatchProduct(w, r)
	// 	case http.MethodDelete:
	// 		h.handleDeleteProduct(w, r)
    //     default:
    //         utils.WriteError(w, http.StatusMethodNotAllowed, fmt.Errorf("method not allowed"))
    //     }
    //     return
    // }

    // utils.WriteError(w, http.StatusNotFound, fmt.Errorf("not found"))
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


	err := h.store.CreateSeller(payload)
	if err != nil {
		utils.WriteError(w, http.StatusInternalServerError, err)
		return
	}

	utils.WriteJSON(w, http.StatusCreated, map[string]string{"status": "seller created"})
}
