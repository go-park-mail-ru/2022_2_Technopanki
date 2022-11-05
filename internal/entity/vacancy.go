package entity

import (
	"time"
)

type Vacancy struct {
	ID            int       `json:"id" gorm:"primary_key"`
	Title         string    `json:"title"`
	Description   string    `json:"description"`
	MinimalSalary int       `json:"minimal_salary"`
	MaximumSalary int       `json:"salary"`
	EmployerID    int       `json:"-"`
	EmployerName  string    `json:"image"`
	City          string    `json:"city"`
	Date          time.Time `json:"date"`
}
