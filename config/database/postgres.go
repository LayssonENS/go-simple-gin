package database

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/LayssonENS/go-simple-gin/config"
	_ "github.com/lib/pq"
)

func NewPostgresConnection() (*sql.DB, error) {
	dataSourceName := fmt.Sprintf(
		"postgres://%s:%s@%s:%s/%s?sslmode=disable",
		config.GetEnv().DbConfig.User,
		config.GetEnv().DbConfig.Password,
		config.GetEnv().DbConfig.Host,
		config.GetEnv().DbConfig.Port,
		config.GetEnv().DbConfig.Name,
	)

	db, err := sql.Open("postgres", dataSourceName)
	if err != nil {
		return nil, fmt.Errorf("failed to open database: %v", err)
	}

	err = db.Ping()
	if err != nil {
		return nil, fmt.Errorf("failed to ping database: %v", err)
	}

	log.Println("Successfully connected to the PostgreSQL database.")
	return db, nil
}
