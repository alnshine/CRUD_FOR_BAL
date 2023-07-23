package CRUD_FOR_BAL

type Vacancy struct {
	Id          int     `json:"id"`
	Title       string  `json:"title"`
	Description string  `json:"description"`
	Type        string  `json:"type"`
	Salary      float64 `json:"salary"`
}

type UsersList struct {
	Id        int
	UserId    int
	VacancyId int
}
