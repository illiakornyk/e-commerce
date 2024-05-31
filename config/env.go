package config

import (
	"log"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

type Config struct {
	Host     string
	Port     int
	User     string
	Password string
	DBName   string
	SSLMode  string
	JWTExpirationInSeconds int
	JWTSecret string
}

var Envs = initConfig()

func initConfig() Config {
	godotenv.Load()

	return Config{
		Host:    getEnv("DB_HOST", "localhost"),
		Port:     parseIntEnv("DB_PORT", 5432),
		User:     getEnv("DB_USER", "postgres"),
		Password: getEnv("DB_PASSWORD", "postgres"),
		DBName:   getEnv("DB_NAME", "ecommerce"),
		SSLMode:  getEnv("DB_SSL_MODE", "disable"),
		JWTExpirationInSeconds: parseIntEnv("JWT_EXPIRATION_IN_SECONDS", 3600*24*7),
		JWTSecret: getEnv("JWT_SECRET", "secret"),
	}
}

func parseIntEnv(key string, defaultValue int) int {
	value, ok := os.LookupEnv(key)
	if !ok {
		return defaultValue
	}

	i, err := strconv.Atoi(value)
	if err != nil {
		log.Printf("Failed to parse %s environment variable: %v", key, err)
		return defaultValue
	}

	return i
}



func getEnv(key, fallback string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}
	return fallback
}
