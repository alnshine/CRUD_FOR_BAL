package CRUD_FOR_BAL

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
