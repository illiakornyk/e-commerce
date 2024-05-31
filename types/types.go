package types

import (
	"errors"
	"regexp"
	"time"
)

type UserStore interface {
	GetUserByEmail(email string) (*User, error)
	GetUserByID(id int) (*User, error)
	CreateUser(User) error
}

type User struct {
	ID        int         `json:"id"`
	Username  string      `json:"username"`
	Password  string      `json:"password,omitempty"`
	Email     string      `json:"email"`
	CreatedAt time.Time   `json:"created_at"`
}

type RegisterUserPayload struct {
	Username string `json:"username"`
	Password string `json:"password"`
	Email    string `json:"email"`
}


type LoginUserPayload struct {
	Password string `json:"password"`
	Email    string `json:"email"`
}

type ProductStore interface {
	GetProductByID(id int) (*Product, error)
	GetProductsByID(ids []int) ([]Product, error)
	GetProducts() ([]*Product, error)
	CreateProduct(CreateProductPayload) error
	UpdateProduct(Product) error
}

type Product struct {
	ID          int     `json:"id"`
	Title       string  `json:"title"`
	Description string  `json:"description"`
	Price       float64 `json:"price"`
	Seller      string  `json:"seller"`
	CreatedAt time.Time   `json:"created_at"`
}

type CreateProductPayload struct {
	Title       string  `json:"title"`
	Description string  `json:"description"`
	Price       float64 `json:"price"`
	Seller      string  `json:"seller"`
}




func (p *RegisterUserPayload) Validate() error {
	if p.Username == "" {
		return errors.New("username is required")
	}
	if p.Password == "" {
		return errors.New("password is required")
	}
	if p.Email == "" {
		return errors.New("email is required")
	}

	// Check if email format is valid
	matched, err := regexp.MatchString(`^[a-z0-9._%+\-]+@[a-z0-9.\-]+\.[a-z]{2,4}$`, p.Email)
	if err != nil {
		return err
	}
	if !matched {
		return errors.New("invalid email format")
	}

	return nil
}

func (p *LoginUserPayload) Validate() error {
	if p.Password == "" {
		return errors.New("password is required")
	}
	if p.Email == "" {
		return errors.New("email is required")
	}

	// Check if email format is valid
	matched, err := regexp.MatchString(`^[a-z0-9._%+\-]+@[a-z0-9.\-]+\.[a-z]{2,4}$`, p.Email)
	if err != nil {
		return err
	}
	if !matched {
		return errors.New("invalid email format")
	}

	return nil
}
