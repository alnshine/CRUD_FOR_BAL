package repository

import (
	"alnshine/CRUD_FOR_BAL"
	"fmt"
	"github.com/jmoiron/sqlx"
)

type VacancyPostgres struct {
	db *sqlx.DB
}

func NewVacancyPostgres(db *sqlx.DB) *VacancyPostgres {
	return &VacancyPostgres{db: db}
}
func (r *VacancyPostgres) Create(userId int, vac CRUD_FOR_BAL.Vacancy) (int, error) {
	tx, err := r.db.Begin()
	if err != nil {
		return 0, err
	}
	var id int
	createVacQuery := fmt.Sprintf("INSERT INTO %s (title,description,type,salary) VALUES ($1,$2,$3,$4) RETURNING id", vacanciesTable)
	row := tx.QueryRow(createVacQuery, vac.Title, vac.Description, vac.Type, vac.Salary)
	if err := row.Scan(&id); err != nil {
		tx.Rollback()
		return 0, err
	}
	createUsersListQuery := fmt.Sprintf("INSERT INTO %s (user_id, vacancy_id) VALUES ($1, $2)", usersListsTable)
	_, err = tx.Exec(createUsersListQuery, userId, id)
	if err != nil {
		tx.Rollback()
		return 0, err
	}

	return id, tx.Commit()
}
func (r *VacancyPostgres) GetAll(userId int) ([]CRUD_FOR_BAL.Vacancy, error) {
	var vac []CRUD_FOR_BAL.Vacancy
	query := fmt.Sprintf("SELECT v.id, v.title, v.description FROM %s v INNER JOIN %s ul on v.id = ul.vacancy_id WHERE ul.user_id = $1", vacanciesTable, usersListsTable)
	err := r.db.Select(&vac, query, userId)

	return vac, err
}
func (r *VacancyPostgres) GetById(userId, vacId int) (CRUD_FOR_BAL.Vacancy, error) {
	var vac CRUD_FOR_BAL.Vacancy

	query := fmt.Sprintf(`SELECT v.id, v.title, v.description FROM %s v
								INNER JOIN %s ul on v.id = ul.vacancy_id WHERE ul.user_id = $1 AND ul.vacancy_id = $2`,
		vacanciesTable, usersListsTable)
	err := r.db.Get(&vac, query, userId, vacId)

	return vac, err
}
