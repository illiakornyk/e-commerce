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

func (s *Store) GetCustomers() ([]*types.Customer, error) {
    rows, err := s.db.Query("SELECT id, first_name, last_name, email, phone_number, address, created_at, updated_at FROM customers")
    if err != nil {
        return nil, err
    }
    defer rows.Close()

    customers := make([]*types.Customer, 0)
    for rows.Next() {
        c, err := scanRowsIntoCustomer(rows)
        if err != nil {
            return nil, err
        }

        customers = append(customers, c)
    }

    if err = rows.Err(); err != nil {
        return nil, err
    }

    return customers, nil
}


func (s *Store) GetCustomerByID(customerID int) (*types.Customer, error) {
    query := "SELECT id, first_name, last_name, email, phone_number, address, created_at, updated_at FROM customers WHERE id = $1"
    row := s.db.QueryRow(query, customerID)

    c := new(types.Customer)
    err := row.Scan(&c.ID, &c.FirstName, &c.LastName, &c.Email, &c.PhoneNumber, &c.Address, &c.CreatedAt, &c.UpdatedAt)
    if err != nil {
        return nil, err
    }

    return c, nil
}


func scanRowsIntoCustomer(rows *sql.Rows) (*types.Customer, error) {
    customer := new(types.Customer)

    err := rows.Scan(
        &customer.ID,
        &customer.FirstName,
        &customer.LastName,
        &customer.Email,
        &customer.PhoneNumber,
        &customer.Address,
        &customer.CreatedAt,
        &customer.UpdatedAt,
    )
    if err != nil {
        return nil, err
    }

    return customer, nil
}
