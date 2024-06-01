package order

import (
	"database/sql"
	"fmt"

	"github.com/illiakornyk/e-commerce/types"
)

type Store struct {
	db *sql.DB
}

func NewStore(db *sql.DB) *Store {
	return &Store{db: db}
}
func (s *Store) CreateOrder(order types.Order) (int, error) {
	var id int
	fmt.Printf("order: %v\n", order)
	err := s.db.QueryRow("INSERT INTO orders (user_id, total, status, address) VALUES ($1, $2, $3, $4) RETURNING id",
		order.UserID, order.Total, order.Status, order.Address).Scan(&id)
	if err != nil {
		return 0, err
	}

	return id, nil
}

func (s *Store) CreateOrderItem(orderItem types.OrderItem) error {
	_, err := s.db.Exec("INSERT INTO order_items (order_id, product_id, quantity, price) VALUES ($1, $2, $3, $4)",
		orderItem.OrderID, orderItem.ProductID, orderItem.Quantity, orderItem.Price)
	return err
}
