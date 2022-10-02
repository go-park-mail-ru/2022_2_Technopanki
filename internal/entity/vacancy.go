package entity

import "time"

type Vacancy struct {
	ID            int       `json:"-"`
	Title         string    `json:"title"`
	Description   string    `json:"description"`
	MinimalSalary int       `json:"minimal_salary"`
	MaximumSalary int       `json:"maximum_salary"`
	EmployerID    int       `json:"-"`
	Date          time.Time `json:"date"`
}
