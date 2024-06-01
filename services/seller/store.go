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

func (s *Store) GetSellerByID(sellerID int) (*types.Seller, error) {
    query := "SELECT id, name, phone_number, created_at, updated_at FROM sellers WHERE id = $1"
    row := s.db.QueryRow(query, sellerID)

    seller := new(types.Seller)
    err := row.Scan(&seller.ID, &seller.Name, &seller.PhoneNumber, &seller.CreatedAt, &seller.UpdatedAt)
    if err != nil {
        return nil, err
    }

    return seller, nil
}

func (s *Store) CreateSeller(seller types.CreateSellerPayload) error {
    _, err := s.db.Exec("INSERT INTO sellers (name, phone_number) VALUES ($1, $2) ",
		seller.Name, seller.PhoneNumber)
    if err != nil {
        return err
    }

    return nil
}


func (s *Store) GetSellers() ([]*types.Seller, error) {
	rows, err := s.db.Query("SELECT * FROM sellers")
	if err != nil {
		return nil, err
	}

	sellers := make([]*types.Seller, 0)
	for rows.Next() {
		seller, err := scanRowsIntoSeller(rows)
		if err != nil {
			return nil, err
		}

		sellers = append(sellers, seller)
	}

	return sellers, nil
}



func scanRowsIntoSeller(rows *sql.Rows) (*types.Seller, error) {
	seller := new(types.Seller)

	err := rows.Scan(
		&seller.ID,
		&seller.Name,
		&seller.PhoneNumber,
		&seller.CreatedAt,
		&seller.UpdatedAt,
	)
	if err != nil {
		return nil, err
	}

	return seller, nil
}
