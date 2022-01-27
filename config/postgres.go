package config

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

func NewPostgres(cfg *Config) (*sql.DB, error) {
	dbConfig := fmt.Sprintf("host=%s port=%s user=%s "+
		"password=%s dbname=%s sslmode=%s",
		cfg.Database.Host,
		cfg.Database.Port,
		cfg.Database.Username,
		cfg.Database.Password,
		cfg.Database.Name,
		cfg.Database.SSL,
	)

	db, err := sql.Open("postgres", dbConfig)

	if err != nil {
		return nil, err
	}

	return db, err
}
