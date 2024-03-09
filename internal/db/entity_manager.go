package db

import "database/sql"
import "invoice-service/internal/config"

type EntityManager struct {
	DB *sql.DB
}

// NewEntityManager creates a new EntityManager with the provided database configuration.
func NewEntityManager(config *config.DBConfig) (*EntityManager, error) {
	db, err := sql.Open("mysql", config.Username+":"+config.Password+"@tcp("+config.Host+":"+config.Port+")/"+config.Database)
	if err != nil {
		return nil, err
	}

	err = db.Ping()
	if err != nil {
		return nil, err
	}

	return &EntityManager{DB: db}, nil
}
