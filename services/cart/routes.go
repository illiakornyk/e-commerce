package cart

import (
	"net/http"

	"github.com/illiakornyk/e-commerce/services/auth"
	"github.com/illiakornyk/e-commerce/types"
	"github.com/illiakornyk/e-commerce/utils"
)

type Handler struct {
	store      types.ProductStore
	orderStore types.OrdersStore
	userStore  types.UserStore
}

func NewHandler(store types.ProductStore, orderStore types.OrdersStore, userStore types.UserStore) *Handler {
	return &Handler{
		store: store,
		orderStore: orderStore,
		userStore: userStore,
	}
}

func (h *Handler) RegisterRoutes(router *http.ServeMux) {

	router.HandleFunc("/cart/checkout", auth.WithJWTAuth(h.handleCheckout, h.userStore))


}

func (h *Handler) handleCheckout(w http.ResponseWriter, r *http.Request) {
	// userID := auth.GetUserIDFromContext(r.Context())

	var cart types.CartCheckoutPayload
	if err := utils.ParseJSON(r, &cart); err != nil {
		utils.WriteError(w, http.StatusBadRequest, err)
		return
	}


	if err := cart.Validate(); err != nil {
		utils.WriteError(w, http.StatusBadRequest, err)
		return
	}


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


	orderID, totalPrice, err := h.createOrder(products, cart.Items, cart.CustomerID)
	if err != nil {
		utils.WriteError(w, http.StatusBadRequest, err)
		return
	}

	utils.WriteJSON(w, http.StatusOK, map[string]interface{}{
		"total_price": totalPrice,
		"order_id":    orderID,
	})
}
