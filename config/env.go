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
