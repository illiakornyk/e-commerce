package cart

import (
	"net/http"

	"github.com/illiakornyk/e-commerce/types"
)

type Handler struct {
	store types.OrdersStore
}

func NewHandler(store types.OrdersStore) *Handler {
	return &Handler{store: store}
}

func (h *Handler) RegisterRoutes(router *http.ServeMux) {

}

func (h *Handler) handleCheckout(w http.ResponseWriter, r *http.Request) {
}
