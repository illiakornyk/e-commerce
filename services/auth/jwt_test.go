package auth

import (
	"testing"
	"time"

	"github.com/golang-jwt/jwt"
	"github.com/illiakornyk/e-commerce/config"
)

func TestCreateJWT(t *testing.T) {
	secret := []byte("secret")

	userID := 1
	expiration := time.Second * time.Duration(config.Envs.JWTExpirationInSeconds)

	claims := jwt.MapClaims{
		"user_id":   userID,
		"expiredAt": time.Now().Add(expiration).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(secret)

	if err != nil {
		t.Fatalf("error creating JWT: %v", err)
	}

	if tokenString == "" {
		t.Error("expected token to be not empty")
	}
}
