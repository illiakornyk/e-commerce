package product

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
	store types.ProductStore

	userStore  types.UserStore
}

func NewHandler(store types.ProductStore, userStore types.UserStore) *Handler {
	return &Handler{store: store, userStore: userStore}
}

func (h *Handler) RegisterRoutes(mux *http.ServeMux) {


    mux.HandleFunc("/products", auth.WithJWTAuth(h.handleProducts, h.userStore))
    mux.HandleFunc("/products/", auth.WithJWTAuth(h.handleProducts, h.userStore))
}

func (h *Handler) handleProducts(w http.ResponseWriter, r *http.Request) {
    // Check if the path is exactly "/products", which means we're handling the collection of products
    if r.URL.Path == "/products" {
        switch r.Method {
        case http.MethodGet:
            h.handleGetProducts(w, r)
        case http.MethodPost:
            h.handleCreateProduct(w, r)
        default:
            utils.WriteError(w, http.StatusMethodNotAllowed, fmt.Errorf("method not allowed"))
        }
        return
    }

    pathSegments := strings.Split(r.URL.Path, "/")
    if len(pathSegments) == 3 {
        switch r.Method {
        case http.MethodGet:
            h.handleGetProductByID(w, r)
        case http.MethodPatch:
            h.handlePatchProduct(w, r)
		case http.MethodDelete:
			h.handleDeleteProduct(w, r)
        default:
            utils.WriteError(w, http.StatusMethodNotAllowed, fmt.Errorf("method not allowed"))
        }
        return
    }

    utils.WriteError(w, http.StatusNotFound, fmt.Errorf("not found"))
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

func (h *Handler) handleGetProductByID(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		utils.WriteError(w, http.StatusMethodNotAllowed, fmt.Errorf("method not allowed"))
		return
	}

	// Extract the product ID from the URL path
	pathSegments := strings.Split(r.URL.Path, "/")
	if len(pathSegments) < 3 {
		utils.WriteError(w, http.StatusBadRequest, fmt.Errorf("product ID is required"))
		return
	}
	productIDStr := pathSegments[2]
	productID, err := strconv.Atoi(productIDStr)
	if err != nil {
		utils.WriteError(w, http.StatusBadRequest, fmt.Errorf("invalid product ID"))
		return
	}

	// Retrieve the product by ID
	product, err := h.store.GetProductByID(productID)
	if err != nil {
		utils.WriteError(w, http.StatusInternalServerError, err)
		return
	}
	if product == nil {
		utils.WriteError(w, http.StatusNotFound, fmt.Errorf("product with ID %d not found", productID))
		return
	}

	utils.WriteJSON(w, http.StatusOK, product)
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


	err := h.store.CreateProduct(payload)
	if err != nil {
		utils.WriteError(w, http.StatusInternalServerError, err)
		return
	}

	utils.WriteJSON(w, http.StatusCreated, map[string]string{"status": "product created"})
}


func (h *Handler) handlePatchProduct(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPatch {
		utils.WriteError(w, http.StatusMethodNotAllowed, fmt.Errorf("method not allowed"))
		return
	}

	productIDStr := strings.TrimPrefix(r.URL.Path, "/products/")
	productID, err := strconv.Atoi(productIDStr)
	if err != nil {
		utils.WriteError(w, http.StatusBadRequest, fmt.Errorf("invalid product ID"))
		return
	}

	var payload types.PatchProductPayload
	if err := utils.ParseJSON(r, &payload); err != nil {
		utils.WriteError(w, http.StatusBadRequest, err)
		return
	}

	existingProduct, err := h.store.GetProductByID(productID)
	if err != nil {
		utils.WriteError(w, http.StatusInternalServerError, err)
		return
	}
	if existingProduct == nil {
		utils.WriteError(w, http.StatusNotFound, fmt.Errorf("product with ID %d does not exist", productID))
		return
	}


	err = h.store.PatchProduct(payload, productID)
	if err != nil {
		utils.WriteError(w, http.StatusInternalServerError, err)
		return
	}

	utils.WriteJSON(w, http.StatusOK, map[string]string{"status": "product updated"})
}


func (h *Handler) handleDeleteProduct(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodDelete {
		utils.WriteError(w, http.StatusMethodNotAllowed, fmt.Errorf("method not allowed"))
		return
	}


	productIDStr := strings.TrimPrefix(r.URL.Path, "/products/")
	productID, err := strconv.Atoi(productIDStr)
	if err != nil {
		utils.WriteError(w, http.StatusBadRequest, fmt.Errorf("invalid product ID"))
		return
	}


	err = h.store.DeleteProduct(productID)
	if err != nil {
		utils.WriteError(w, http.StatusInternalServerError, err)
		return
	}

	utils.WriteJSON(w, http.StatusCreated, map[string]string{"status": "product deleted"})
}
