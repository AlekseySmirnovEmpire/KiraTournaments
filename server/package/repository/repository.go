package repository

import (
	cardmarketplace "github.com/AlekseySmirnovEmpire/CardMarketplace"
	"github.com/jmoiron/sqlx"
)

type Authorization interface {
	CreateUser(user cardmarketplace.User) (int, error)
	GetUser(username, password string) (cardmarketplace.User, error)
}

type Repository struct {
	Authorization
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{
		Authorization: NewAuthPostgres(db),
	}
}
