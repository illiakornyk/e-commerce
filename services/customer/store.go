package customer

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

func (s *Store) CreateCustomer(customer types.CreateCustomerPayload) error {
    _, err := s.db.Exec("INSERT INTO customers (first_name, last_name, email, phone_number, address) VALUES ($1, $2, $3, $4, $5)",
        customer.FirstName, customer.LastName, customer.Email, customer.PhoneNumber, customer.Address)
    if err != nil {
        return err
    }

    return nil
}
