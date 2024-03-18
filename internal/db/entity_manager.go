package db

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
	"invoice-service/internal/config"
)

type EntityManager struct {
	DB *sql.DB
}

// NewEntityManager creates a new EntityManager with the provided database configuration.
func NewEntityManager(config *config.DBConfig) (*EntityManager, error) {
	db, err := sql.Open("postgres", "postgresql://"+config.Username+":"+config.Password+"@"+config.Host+"/postgres?sslmode=disable")
	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	err = db.Ping()
	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	return &EntityManager{DB: db}, nil
}
