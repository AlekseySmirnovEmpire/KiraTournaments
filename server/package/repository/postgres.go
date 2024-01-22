package repository

import "github.com/jmoiron/sqlx"

const (
	userTable = "users"
)

type Config struct {
	ConnectionString string
}

func NewPostgresDB(cfg Config) (*sqlx.DB, error) {
	db, err := sqlx.Open("postgres", cfg.ConnectionString)
	if err != nil {
		return nil, err
	}

	err = db.Ping()
	if err != nil {
		return nil, err
	}

	return db, nil
}
