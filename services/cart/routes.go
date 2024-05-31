package cart

import (
	"net/http"

	"github.com/illiakornyk/e-commerce/types"
	"github.com/illiakornyk/e-commerce/utils"
)

type Handler struct {
	store      types.ProductStore
	orderStore types.OrdersStore
}

func NewHandler(store types.ProductStore, orderStore types.OrdersStore) *Handler {
	return &Handler{
		store: store,
		orderStore: orderStore,
	}
}

func (h *Handler) RegisterRoutes(router *http.ServeMux) {

	router.HandleFunc("/cart/checkout", h.handleCheckout)
}

func (h *Handler) handleCheckout(w http.ResponseWriter, r *http.Request) {
	// userID := auth.GetUserIDFromContext(r.Context())
	userID := 0

	var cart types.CartCheckoutPayload
	if err := utils.ParseJSON(r, &cart); err != nil {
		utils.WriteError(w, http.StatusBadRequest, err)
		return
	}

	// if err := utils.Validate.Struct(cart); err != nil {
	// 	errors := err.(validator.ValidationErrors)
	// 	utils.WriteError(w, http.StatusBadRequest, fmt.Errorf("invalid payload: %v", errors))
	// 	return
	// }

	productIds, err := getCartItemsIDs(cart.Items)
	if err != nil {
		utils.WriteError(w, http.StatusBadRequest, err)
		return
	}

	// get products
	products, err := h.store.GetProductsByID(productIds)
	if err != nil {
		utils.WriteError(w, http.StatusInternalServerError, err)
		return
	}

	orderID, totalPrice, err := h.createOrder(products, cart.Items, userID)
	if err != nil {
		utils.WriteError(w, http.StatusBadRequest, err)
		return
	}

	utils.WriteJSON(w, http.StatusOK, map[string]interface{}{
		"total_price": totalPrice,
		"order_id":    orderID,
	})
}
