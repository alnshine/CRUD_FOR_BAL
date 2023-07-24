package service

import (
	"alnshine/CRUD_FOR_BAL"
	"alnshine/CRUD_FOR_BAL/pkg/repository"
)

type VacancyService struct {
	repo repository.Vacancy
}

func NewVacancyService(repo repository.Vacancy) *VacancyService {
	return &VacancyService{repo: repo}
}
func (s *VacancyService) Create(userId int, vac CRUD_FOR_BAL.Vacancy) (int, error) {
	return s.repo.Create(userId, vac)
}
func (s *VacancyService) GetAll(userId int) ([]CRUD_FOR_BAL.Vacancy, error) {
	return s.repo.GetAll(userId)
}
func (s *VacancyService) GetById(userId, vacId int) (CRUD_FOR_BAL.Vacancy, error) {
	return s.repo.GetById(userId, vacId)
}
