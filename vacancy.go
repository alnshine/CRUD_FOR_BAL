package CRUD_FOR_BAL

import "errors"

type Vacancy struct {
	Id          int     `json:"id" db:"id"`
	Title       string  `json:"title" db:"title" binding:"required"`
	Description string  `json:"description" db:"description"`
	Type        string  `json:"type" db:"type"`
	Salary      float64 `json:"salary" db:"salary"`
}

type UsersList struct {
	Id        int
	UserId    int
	VacancyId int
}
type UpdateVac struct {
	Title       *string `json:"title"`
	Description *string `json:"description"`
	Type        *string `json:"type"`
	Salary      *string `json:"salary"`
}

func (i UpdateVac) Validate() error {
	if i.Title == nil && i.Description == nil {
		return errors.New("update structure has no values")
	}
	return nil
}
