package user

import (
	"database/sql"
	"log"

	"fmt"

	"github.com/illiakornyk/e-commerce/types"
)

type Store struct {
	db *sql.DB
}

func NewStore(db *sql.DB) *Store {

	return &Store{db: db}

}

func (s *Store) GetUserByEmail(email string) (*types.User, error) {
    query := "SELECT id, username, email, password, created_at FROM users WHERE email = $1"
    row := s.db.QueryRow(query, email)

    u := new(types.User)
    err := row.Scan(&u.ID, &u.Username, &u.Email, &u.Password, &u.CreatedAt)
    if err != nil {
        if err == sql.ErrNoRows {
            log.Println("No user found with email:", email)
            return nil, nil
        }
        log.Println("Error querying user by email:", err)
        return nil, err
    }

    return u, nil
}




func scanRowsIntoUser(rows *sql.Rows) (*types.User, error) {

	u := new(types.User)
	err := rows.Scan(
		&u.ID,
		&u.Username,
		&u.Password,
		&u.Email,
		&u.CreatedAt,
	)

	if err != nil {
		return nil, err
	}

	return u, err
}


func (s *Store) GetUserByID(id int) (*types.User, error) {
		rows, err := s.db.Query("SELECT * FROM users WHERE id = $1", id)

	if err != nil {
		return nil, err
	}
	defer rows.Close()

	u := new(types.User)
	for rows.Next() {
		u, err = scanRowsIntoUser(rows)

		if err != nil {
			return nil, err
		}
	}

	if u.ID == 0 {
		return nil, fmt.Errorf("user not found")
	}

	return u, nil
}




func (s *Store) CreateUser(u types.User) error {
	_, err := s.db.Exec("INSERT INTO users (username, password, email) VALUES ($1, $2, $3)", u.Username, u.Password, u.Email)


	if err != nil {
		return err
	}

	return nil
}
