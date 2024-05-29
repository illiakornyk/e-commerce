package db

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"
)

type Config struct {
	Host     string
	Port     int
	User     string
	Password string
	DBName   string
	SSLMode  string
}

func (c *Config) FormatDSN() string {
	return fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=%s",
		c.Host, c.Port, c.User, c.Password, c.DBName, c.SSLMode)
}

func NewPostgresStorage(cfg Config) (*sql.DB, error) {
	dsn := cfg.FormatDSN()
	db, err := sql.Open("postgres", dsn)
	if err != nil {
		log.Fatal(err)
	}

	return db, nil
}
