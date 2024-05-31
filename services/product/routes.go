package product

import (
	"fmt"
	"net/http"

	"github.com/illiakornyk/e-commerce/types"
	"github.com/illiakornyk/e-commerce/utils"
)

type Handler struct {
	store types.ProductStore
}

func NewHandler(store types.ProductStore) *Handler {
	return &Handler{store: store}
}

func (h *Handler) RegisterRoutes(mux *http.ServeMux) {
	    mux.HandleFunc("/products", h.handleProducts)


}

func (h *Handler) handleProducts(w http.ResponseWriter, r *http.Request) {
    switch r.Method {
    case http.MethodGet:
        h.handleGetProducts(w, r)
    case http.MethodPost:
        h.handleCreateProduct(w, r)
    default:
        utils.WriteError(w, http.StatusMethodNotAllowed, fmt.Errorf("method not allowed"))
    }
}


func (h *Handler) handleGetProducts(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		utils.WriteError(w, http.StatusMethodNotAllowed, fmt.Errorf("method not allowed"))
		return
	}
	ps, err := h.store.GetProducts()
	if err != nil {
		utils.WriteError(w, http.StatusInternalServerError, err)
		return
	}

	utils.WriteJSON(w, http.StatusOK, ps)
}


func (h *Handler) handleCreateProduct(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		utils.WriteError(w, http.StatusMethodNotAllowed, fmt.Errorf("method not allowed"))
		return
	}

	var payload types.CreateProductPayload
	if err := utils.ParseJSON(r, &payload); err != nil {
		utils.WriteError(w, http.StatusBadRequest, err)
		return
	}

	existingProduct, err := h.store.GetProductByTitle(payload.Title)
	if err == nil && existingProduct != nil {
		utils.WriteError(w, http.StatusBadRequest, fmt.Errorf("product with title %s already exists", payload.Title))
		return
	}

	err = h.store.CreateProduct(payload)
	if err != nil {
		utils.WriteError(w, http.StatusInternalServerError, err)
		return
	}

	utils.WriteJSON(w, http.StatusCreated, map[string]string{"status": "product created"})
}
