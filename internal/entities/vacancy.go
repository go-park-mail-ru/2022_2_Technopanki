package entities

import "time"

type Vacancy struct {
	ID          string    `json:"id"`
	Title       string    `json:"title"`
	Description string    `json:"description"`
	Salary      string    `json:"salary "`
	EmployerID  int       `json:"employerID"`
	Date        time.Time `json:"date"`
}
