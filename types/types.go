package types

import (
	"errors"
	"fmt"
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
	PatchProduct(payload PatchProductPayload, productID int) error
	UpdateProduct(Product) error
	DeleteProduct(productID int) error
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

type PatchProductPayload struct {
	Title       *string  `json:"title"`
	Description *string  `json:"description"`
	Price       *float64 `json:"price"`
	Seller      *string  `json:"seller"`
	Quantity    *int     `json:"quantity"`
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


type CustomerStore interface {
	GetCustomerByID(id int) (*Customer, error)
	GetCustomers() ([]*Customer, error)
	CreateCustomer(CreateCustomerPayload) error
	// UpdateCustomer(Customer) error
	// DeleteCustomer(customerID int) error
}

type Customer struct {
	ID          int       `json:"id"`
	FirstName   string    `json:"first_name"`
	LastName    string    `json:"last_name"`
	Email       string    `json:"email"`
	PhoneNumber string    `json:"phone_number,omitempty"`
	Address     string    `json:"address"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

type CreateCustomerPayload struct {
	FirstName   string `json:"first_name"`
	LastName    string `json:"last_name"`
	Email       string `json:"email"`
	PhoneNumber string `json:"phone_number,omitempty"`
	Address     string `json:"address"`
}


type SellerStore interface {
	GetSellerByID(id int) (*Seller, error)
	GetSellers() ([]*Seller, error)
	CreateSeller(CreateSellerPayload) error
	// UpdateSeller(Seller) error
	// DeleteSeller(sellerID int) error
}

type Seller struct {
	ID          int       `json:"id"`
	Name        string    `json:"name"`
	PhoneNumber string    `json:"phone_number"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

type CreateSellerPayload struct {
	Name        string `json:"name"`
	PhoneNumber string `json:"phone_number"`
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


func (ccp *CartCheckoutPayload) Validate() error {
	if ccp.Items == nil || len(ccp.Items) == 0 {
		return errors.New("the cart cannot be empty")
	}

	for _, item := range ccp.Items {
		if item.ProductID <= 0 {
			return fmt.Errorf("invalid ProductID: %d; must be positive", item.ProductID)
		}
		if item.Quantity <= 0 {
			return fmt.Errorf("invalid Quantity for ProductID %d: %d; must be positive", item.ProductID, item.Quantity)
		}
	}

	return nil
}

func (cpp *CreateProductPayload) Validate() error {
	if cpp.Title == "" {
		return errors.New("the title cannot be empty")
	}
	if cpp.Description == "" {
		return errors.New("the description cannot be empty")
	}
	if cpp.Price <= 0 {
		return fmt.Errorf("invalid Price: %f; must be positive", cpp.Price)
	}
	if cpp.Seller == "" {
		return errors.New("the seller cannot be empty")
	}
	if cpp.Quantity < 0 {
		return fmt.Errorf("invalid Quantity: %d; must be non-negative", cpp.Quantity)
	}

	return nil
}


func (p *CreateCustomerPayload) Validate() error {
	if p.FirstName == "" {
		return errors.New("first name is required")
	}
	if p.LastName == "" {
		return errors.New("last name is required")
	}
	if p.Email == "" {
		return errors.New("email is required")
	}
	if !isValidEmail(p.Email) {
		return errors.New("email is not valid")
	}
	// PhoneNumber is optional, but if provided, it should be validated.
	if p.PhoneNumber != "" && !isValidPhoneNumber(p.PhoneNumber) {
		return errors.New("phone number is not valid")
	}
	if p.Address == "" {
		return errors.New("address is required")
	}
	return nil
}

func isValidEmail(email string) bool {
	emailRegex := regexp.MustCompile(`^[a-z0-9._%+\-]+@[a-z0-9.\-]+\.[a-z]{2,4}$`)
	return emailRegex.MatchString(email)
}

func isValidPhoneNumber(phone string) bool {
	phoneRegex := regexp.MustCompile(`^\+[1-9]\d{1,14}$`)
	return phoneRegex.MatchString(phone)
}
