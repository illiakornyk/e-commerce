package user

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/illiakornyk/e-commerce/types"
)

func TestUserServiceHandlers(t *testing.T) {
	userStore := &mockUserStore{}
	handler := NewHandler(userStore)

	t.Run("should fail if the user payload is invalid", func(t *testing.T) {
		payload := types.RegisterUserPayload{
			Username: "myusername is here",
			Password: "231",
			Email:    "emailIsInvalid",
		}
		marshalled, _ := json.Marshal(payload)

		req, err := http.NewRequest(http.MethodPost, "/register", bytes.NewBuffer(marshalled))
		if err != nil {
			t.Fatal(err)
		}

		rr := httptest.NewRecorder()
		router := http.NewServeMux()
		router.HandleFunc("/register", handler.handleRegister)

		router.ServeHTTP(rr, req)
		if status := rr.Code; status != http.StatusBadRequest {
			t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusBadRequest)
		}
	})

	// t.Run("should register a new user", func(t *testing.T) {
	// 	payload := types.RegisterUserPayload{
	// 				Username: "username",
	// 				Password: "password1245",
	// 				Email:    "emailIsValid@mail.com",
	// 			}
	// 			marshalled, _ := json.Marshal(payload)

	// 			req, err := http.NewRequest(http.MethodPost, "/register", bytes.NewBuffer(marshalled))
	// 			if err != nil {
	// 				t.Fatal(err)
	// 			}

	// 			rr := httptest.NewRecorder()
	// 			router := http.NewServeMux()
	// 			router.HandleFunc("/register", handler.handleRegister)

	// 			router.ServeHTTP(rr, req)


	// 			if status := rr.Code; status != http.StatusCreated {
	// 				t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusCreated)
	// 			}
	// })
}

type mockUserStore struct{}

func (m *mockUserStore) GetUserByEmail(email string) (*types.User, error) {
	return nil, fmt.Errorf("user not found")
}

func (m *mockUserStore) GetUserByID(id int) (*types.User, error) {
	return nil, nil
}

func (m *mockUserStore) CreateUser(u types.User) error {
	return nil
}
