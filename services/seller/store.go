package seller

import (
	"database/sql"

	"github.com/illiakornyk/e-commerce/types"
)

type Store struct {
	db *sql.DB
}

func NewStore(db *sql.DB) *Store {
	return &Store{db: db}
}

func (s *Store) GetSellerByID(productID int) (*types.Product, error) {
    query := "SELECT id, title, description, price, seller, quantity, created_at FROM products WHERE id = $1"
    row := s.db.QueryRow(query, productID)

    p := new(types.Product)
    err := row.Scan(&p.ID, &p.Title, &p.Description, &p.Price, &p.Seller, &p.Quantity, &p.CreatedAt)
    if err != nil {
        return nil, err
    }

    return p, nil
}

func (s *Store) CreateSeller(seller types.CreateSellerPayload) error {
    _, err := s.db.Exec("INSERT INTO sellers (name, phone_number) VALUES ($1, $2) ",
		seller.Name, seller.PhoneNumber)
    if err != nil {
        return err
    }

    return nil
}
