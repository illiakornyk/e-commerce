package product

import (
	"database/sql"
	"errors"
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
    query := "SELECT id, title, description, price, seller_id, quantity, created_at FROM products WHERE id = $1"
    row := s.db.QueryRow(query, productID)

    p := new(types.Product)
    err := row.Scan(&p.ID, &p.Title, &p.Description, &p.Price, &p.SellerID, &p.Quantity, &p.CreatedAt)
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

    query := fmt.Sprintf("SELECT id, title, description, price, seller_id, quantity, created_at FROM products WHERE id IN (%s)", placeholderStr)

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
        err := rows.Scan(&p.ID, &p.Title, &p.Description, &p.Price, &p.SellerID, &p.Quantity, &p.CreatedAt)
        if err != nil {
            return nil, err
        }
        products = append(products, p)
    }

    return products, nil
}

func (s *Store) GetProductByTitle(title string) (*types.Product, error) {
    query := "SELECT id, title, description, price, seller_id, quantity, created_at FROM products WHERE title = $1"
    row := s.db.QueryRow(query, title)

    p := new(types.Product)
    err := row.Scan(&p.ID, &p.Title, &p.Description, &p.Price, &p.SellerID, &p.Quantity, &p.CreatedAt)
    if err != nil {
        if err == sql.ErrNoRows {
            return nil, nil
        }
        return nil, err
    }

    return p, nil
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
    exists, err := s.SellerExists(product.SellerID)
        if err != nil {
            return err
        }
        if !exists {
            return errors.New("seller does not exist")
        }


    // Insert the new product
    _, err = s.db.Exec("INSERT INTO products (title, price, description, seller_id, quantity) VALUES ($1, $2, $3, $4, $5)",
        product.Title, product.Price, product.Description, product.SellerID, product.Quantity)
    return err
}


func (s *Store) PatchProduct(product types.PatchProductPayload, productID int) error {
    query := "UPDATE products SET"
    params := []interface{}{}
    paramID := 1

	if product.SellerID >= 0 {
        exists, err := s.SellerExists(product.SellerID)
        if err != nil {
            return err
        }
        if !exists {
            return errors.New("seller does not exist")
        }

        query += fmt.Sprintf(" seller_id = $%d,", paramID)
        params = append(params, product.SellerID)
        paramID++
    }

    // Check if the pointer is not nil before appending to the query
    if product.Title != nil {
        query += fmt.Sprintf(" title = $%d,", paramID)
        params = append(params, *product.Title)
        paramID++
    }
    if product.Description != nil {
        query += fmt.Sprintf(" description = $%d,", paramID)
        params = append(params, *product.Description)
        paramID++
    }
    if product.Price != nil {
        query += fmt.Sprintf(" price = $%d,", paramID)
        params = append(params, *product.Price)
        paramID++
    }
    if product.Quantity != nil {
        query += fmt.Sprintf(" quantity = $%d,", paramID)
        params = append(params, *product.Quantity)
        paramID++
    }

    // Remove the trailing comma
    query = strings.TrimSuffix(query, ",")

    // Append the WHERE clause to update the product by ID
    query += fmt.Sprintf(" WHERE id = $%d", paramID)
    params = append(params, productID)

    // Execute the query
    _, err := s.db.Exec(query, params...)
    if err != nil {
        return err
    }

    return nil
}




func (s *Store) UpdateProduct(product types.Product) error {
    _, err := s.db.Exec("UPDATE products SET title = $1, price = $2, description = $3, seller_id = $4, quantity = $5 WHERE id = $6",
        product.Title, product.Price, product.Description, product.SellerID, product.Quantity, product.ID)
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
		&product.SellerID,
		&product.Quantity,
		&product.CreatedAt,
	)
	if err != nil {
		return nil, err
	}

	return product, nil
}


func (s *Store) DeleteProduct(productID int) error {
    _, err := s.db.Exec("DELETE FROM products WHERE id = $1", productID)
    if err != nil {
        return err
    }

    return nil
}


func (s *Store) SellerExists(sellerID int) (bool, error) {
    var exists bool
    err := s.db.QueryRow("SELECT EXISTS(SELECT 1 FROM sellers WHERE id = $1)", sellerID).Scan(&exists)
    return exists, err
}
