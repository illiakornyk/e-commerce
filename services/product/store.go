package product

import (
	"database/sql"
	"fmt"
	"strings"

	"github.com/illiakornyk/e-commerce/types"
)

type Store struct {
	db *sql.DB
}

func NewStore(db *sql.DB) *Store {
	return &Store{db: db}
}

func (s *Store) GetProductByID(productID int) (*types.Product, error) {
    query := "SELECT id, title, description, price, seller, created_at FROM products WHERE id = $1"
    row := s.db.QueryRow(query, productID)

    p := new(types.Product)
    err := row.Scan(&p.ID, &p.Title, &p.Description, &p.Price, &p.Seller, &p.CreatedAt)
    if err != nil {
        return nil, err
    }

    return p, nil
}


func (s *Store) GetProductsByID(productIDs []int) ([]types.Product, error) {
    placeholders := make([]string, len(productIDs))
    for i := range placeholders {
        placeholders[i] = fmt.Sprintf("$%d", i+1)
    }
    placeholderStr := strings.Join(placeholders, ",")

    query := fmt.Sprintf("SELECT id, title, description, price, seller, created_at FROM products WHERE id IN (%s)", placeholderStr)

    args := make([]interface{}, len(productIDs))
    for i, v := range productIDs {
        args[i] = v
    }

    rows, err := s.db.Query(query, args...)
    if err != nil {
        return nil, err
    }
    defer rows.Close()

    var products []types.Product
    for rows.Next() {
        var p types.Product
        // Scan the row into the Product struct.
        err := rows.Scan(&p.ID, &p.Title, &p.Description, &p.Price, &p.Seller, &p.CreatedAt)
        if err != nil {
            return nil, err
        }
        products = append(products, p)
    }

    return products, nil
}


func (s *Store) GetProducts() ([]*types.Product, error) {
	rows, err := s.db.Query("SELECT * FROM products")
	if err != nil {
		return nil, err
	}

	products := make([]*types.Product, 0)
	for rows.Next() {
		p, err := scanRowsIntoProduct(rows)
		if err != nil {
			return nil, err
		}

		products = append(products, p)
	}

	return products, nil
}

func (s *Store) CreateProduct(product types.CreateProductPayload) error {
    _, err := s.db.Exec("INSERT INTO products (title, price, description, seller) VALUES ($1, $2, $3, $4)",
        product.Title, product.Price, product.Description, product.Seller)
    if err != nil {
        return err
    }

    return nil
}

func (s *Store) UpdateProduct(product types.Product) error {
    _, err := s.db.Exec("UPDATE products SET title = $1, price = $2, description = $3, seller = $4 WHERE id = $5",
        product.Title, product.Price, product.Description, product.Seller, product.ID)
    if err != nil {
        return err
    }

    return nil
}


func scanRowsIntoProduct(rows *sql.Rows) (*types.Product, error) {
	product := new(types.Product)

	err := rows.Scan(
		&product.ID,
		&product.Title,
		&product.Description,
		&product.Price,
		&product.Seller,
		&product.CreatedAt,
	)
	if err != nil {
		return nil, err
	}

	return product, nil
}
