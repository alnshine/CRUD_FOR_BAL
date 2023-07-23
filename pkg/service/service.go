package service

import (
	"alnshine/CRUD_FOR_BAL"
	"alnshine/CRUD_FOR_BAL/pkg/repository"
)

type Authorization interface {
	CreateUser(user CRUD_FOR_BAL.User) (int, error)
	GenerateToken(username, password string) (string, error)
}
type Vacancy interface {
}
type Service struct {
	Authorization
	Vacancy
}

func NewService(repos *repository.Repository) *Service {
	return &Service{
		Authorization: NewAuthService(repos.Authorization),
	}
}
