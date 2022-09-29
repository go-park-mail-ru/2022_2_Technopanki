package entity

import "time"

type Vacancy struct {
	ID          string `json:"-"`
	Title       string
	Description string
	Salary      int
	EmployerID  int `json:"-"`
	Date        time.Time
}
