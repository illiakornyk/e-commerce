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

	ADMIN_USERNAME string
	ADMIN_PASSWORD string
	ADMIN_EMAIL string
}

var Envs = initConfig()

func initConfig() Config {
	godotenv.Load()

	return Config{
		Host:    getRequiredEnv("DB_HOST"),
		Port:    getRequiredIntEnv("DB_PORT"),
		User:    getRequiredEnv("DB_USER"),
		Password: getRequiredEnv("DB_PASSWORD"),
		DBName:  getRequiredEnv("DB_NAME"),
		SSLMode: getRequiredEnv("DB_SSL_MODE"),
		JWTExpirationInSeconds: getRequiredIntEnv("JWT_EXPIRATION_IN_SECONDS"),
		JWTSecret: getRequiredEnv("JWT_SECRET"),

		ADMIN_USERNAME: getRequiredEnv("ADMIN_USERNAME"),
		ADMIN_PASSWORD: getRequiredEnv("ADMIN_PASSWORD"),
		ADMIN_EMAIL:    getRequiredEnv("ADMIN_EMAIL"),
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


func getRequiredIntEnv(key string) int {
	valueStr := getRequiredEnv(key)
	value, err := strconv.Atoi(valueStr)
	if err != nil {
		log.Fatalf("Environment variable %s must be an integer and is required", key)
	}
	return value
}

func getRequiredEnv(key string) string {
	value, ok := os.LookupEnv(key)
	if !ok {
		log.Fatalf("Environment variable %s is required and not set", key)
	}
	return value
}
