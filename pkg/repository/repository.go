package repository

import (
	"alnshine/CRUD_FOR_BAL"
	"github.com/jmoiron/sqlx"
)

type Authorization interface {
	CreateUser(user CRUD_FOR_BAL.User) (int, error)
	GetUser(username, password string) (CRUD_FOR_BAL.User, error)
}
type Vacancy interface {
	Create(userId int, vac CRUD_FOR_BAL.Vacancy) (int, error)
	GetAll(userId int) ([]CRUD_FOR_BAL.Vacancy, error)
	GetById(userId, vacId int) (CRUD_FOR_BAL.Vacancy, error)
}
type Repository struct {
	Authorization
	Vacancy
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{
		Authorization: NewAuthPostgres(db),
		Vacancy:       NewVacancyPostgres(db),
	}
}
