package repository

import (
	"alnshine/CRUD_FOR_BAL"
	"fmt"
	"github.com/jmoiron/sqlx"
	"github.com/sirupsen/logrus"
	"strings"
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
func (r *VacancyPostgres) Delete(userId, vacId int) error {
	query := fmt.Sprintf("DELETE FROM %s v USING %s ul WHERE v.id = ul.vacancy_id AND ul.user_id=$1 AND ul.vacancy_id=$2",
		vacanciesTable, usersListsTable)
	_, err := r.db.Exec(query, userId, vacId)
	return err
}
func (r *VacancyPostgres) Update(userId, VacId int, input CRUD_FOR_BAL.UpdateVac) error {
	setValues := make([]string, 0)
	args := make([]interface{}, 0)
	argId := 1
	if input.Title != nil {
		setValues = append(setValues, fmt.Sprintf("title=$%d", argId))
		args = append(args, *input.Title)
		argId++
	}
	if input.Description != nil {
		setValues = append(setValues, fmt.Sprintf("description=$%d", argId))
		args = append(args, *input.Description)
		argId++
	}
	setQuery := strings.Join(setValues, ", ")
	query := fmt.Sprintf("UPDATE %s v SET %s FROM %s users_lists WHERE v.id = users_lists.vacancy_id AND users_lists.vacancy_id=$%d AND users_lists.user_id=$%d",
		vacanciesTable, setQuery, usersTable, argId, argId+1)
	args = append(args, VacId, userId)
	logrus.Debugf("updateQuery %s", query)
	logrus.Debugf("args: %s", args)
	_, err := r.db.Exec(query, args...)
	return err
}
