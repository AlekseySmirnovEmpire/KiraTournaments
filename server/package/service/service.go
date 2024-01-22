package service

import (
	cardmarketplace "github.com/AlekseySmirnovEmpire/CardMarketplace"
	"github.com/AlekseySmirnovEmpire/CardMarketplace/package/repository"
)

type Authorization interface {
	CreateUser(user cardmarketplace.User) (int, error)
	GenerateToken(username, password string) (string, error)
	ParseToken(token string) (int, error)
}

type Service struct {
	Authorization
}

func NewService(repos *repository.Repository) *Service {
	return &Service{
		Authorization: NewAuthService(repos.Authorization),
	}
}
