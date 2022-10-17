package entity

import (
	"gorm.io/gorm"
	"time"
)

type Vacancy struct {
	gorm.Model
	ID            int       `json:"-" gorm:"primary_key"`
	Title         string    `json:"title"`
	Description   string    `json:"description"`
	MinimalSalary int       `json:"minimal_salary"`
	MaximumSalary int       `json:"maximum_salary"`
	EmployerID    int       `json:"-"`
	EmployerName  string    `json:"employer_name"`
	City          string    `json:"city"`
	Date          time.Time `json:"date"`
}
