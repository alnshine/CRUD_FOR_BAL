package service

import (
	"alnshine/CRUD_FOR_BAL"
	"alnshine/CRUD_FOR_BAL/pkg/repository"
)

type Authorization interface {
	CreateUser(user CRUD_FOR_BAL.User) (int, error)
	GenerateToken(username, password string) (string, error)
	ParseToken(token string) (int, error)
}
type Vacancy interface {
	Create(userId int, list CRUD_FOR_BAL.Vacancy) (int, error)
	GetAll(userId int) ([]CRUD_FOR_BAL.Vacancy, error)
	GetById(userId, vacId int) (CRUD_FOR_BAL.Vacancy, error)
	Delete(userId, vacId int) error
	Update(userId, vacId int, input CRUD_FOR_BAL.UpdateVac) error
}
type Service struct {
	Authorization
	Vacancy
}

func NewService(repos *repository.Repository) *Service {
	return &Service{
		Authorization: NewAuthService(repos.Authorization),
		Vacancy:       NewVacancyService(repos.Vacancy),
	}
}
