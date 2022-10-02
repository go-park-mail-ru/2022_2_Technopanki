package entity

import "time"

type Vacancy struct {
	ID            string    `json:"-"`
	Title         string    `json:"title"`
	Description   string    `json:"description"`
	MinimalSalary int       `json:"minimal salary"`
	MaximumSalary int       `json:"maximum salary"`
	EmployerID    int       `json:"-"`
	Date          time.Time `json:"date"`
}
