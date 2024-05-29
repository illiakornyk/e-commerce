package types

import "time"

type UserStore interface {
	GetUserByEmail(email string) (*User, error)
	GetUserByID(id int) (*User, error)
	CreateUser(u User) error
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
