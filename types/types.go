package types

import (
	"errors"
	"regexp"
	"time"
)

// User
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

// Product

type ProductStore interface {
	GetProductByID(id int) (*Product, error)
	GetProductsByID(ids []int) ([]Product, error)
	GetProductByTitle(title string) (*Product, error)
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
	Quantity    int     `json:"quantity"`
	CreatedAt time.Time   `json:"created_at"`
}

type CreateProductPayload struct {
	Title       string  `json:"title"`
	Description string  `json:"description"`
	Price       float64 `json:"price"`
	Seller      string  `json:"seller"`
	Quantity    int     `json:"quantity"`
}

type OrdersStore interface {
	CreateOrder(Order) (int, error)
	CreateOrderItem(OrderItem) error
}

type Order struct {
	ID int `json:"id"`
	UserID int `json:"user_id"`
	Total float64 `json:"total"`
	Status string `json:"status"`
	Address string `json:"address"`
	CreatedAt time.Time   `json:"created_at"`
}

type OrderItem struct {
	ID int `json:"id"`
	OrderID int `json:"order_id"`
	ProductID int `json:"product_id"`
	Quantity int `json:"quantity"`
	Price float64 `json:"price"`
	CreatedAt time.Time   `json:"created_at"`
}

type CartItem struct {
	ProductID int `json:"product_id"`
	Quantity int `json:"quantity"`
}

type CartCheckoutPayload struct {
	Items []CartCheckoutItem `json:"items"`
}

type CartCheckoutItem struct {
	ProductID int `json:"product_id"`
	Quantity  int `json:"quantity"`
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
