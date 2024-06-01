package auth

import (
	"github.com/illiakornyk/e-commerce/types"
)

// AdminCreationResult is used to indicate the result of the CreateAdminUser function.
type AdminCreationResult int

const (
	AdminCreated AdminCreationResult = iota // AdminCreated indicates a new admin was created.
	AdminAlreadyExists                      // AdminAlreadyExists indicates the admin already existed.
)

func CreateAdminUser(store types.UserStore, username, password, email string) (AdminCreationResult, error) {
	existingUser, err := store.GetUserByEmail(email)
	if err != nil {
		return 0, err
	}
	if existingUser != nil {
		return AdminAlreadyExists, nil
	}

	hashedPassword, err := HashPassword(password)
	if err != nil {
		return 0, err
	}

	adminUser := types.User{
		Username: username,
		Password: hashedPassword,
		Email:    email,
	}

	err = store.CreateUser(adminUser)
	if err != nil {
		return 0, err
	}

	return AdminCreated, nil
}
