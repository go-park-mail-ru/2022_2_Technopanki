package models

import (
	"time"
)

type Vacancy struct {
	ID                uint              `json:"id" gorm:"primaryKey;"`
	PostedByUserId    uint              `json:"postedByUserId" gorm:"not null;"`
	JobType           string            `json:"jobType" gorm:"not null;"`
	Title             string            `json:"title" gorm:"not null;"`
	Description       string            `json:"description" gorm:"not null;"`
	Tasks             string            `json:"tasks" gorm:"not null;"`
	Requirements      string            `json:"requirements" gorm:"not null"`
	Extra             string            `json:"extra" gorm:"not null"`
	CreatedDate       time.Time         `json:"createdDate" gorm:"not null;"`
	Salary            string            `json:"salary"`
	Location          string            `json:"location" gorm:"not null;"`
	IsActive          bool              `json:"isActive" gorm:"not null;"`
	Experience        string            `json:"experience" gorm:"not null;"`
	Format            string            `json:"format" gorm:"not null;"`
	Hours             string            `json:"hours" gorm:"not null;"`
	VacancyActivities []VacancyActivity `json:"vacancyActivities" gorm:"foreignKey:VacancyId;constraint:OnDelete:CASCADE;"`
	Skills            []Skill           `json:"skills" gorm:"many2many:vacancy_skills;"`
}

type UpdateVacancy struct {
	JobType      string `json:"jobType"`
	Title        string `json:"title"`
	Description  string `json:"description"`
	Tasks        string `json:"tasks"`
	Requirements string `json:"requirements"`
	Extra        string `json:"extra"`
	Salary       string `json:"salary"`
	Location     string `json:"location"`
	IsActive     bool   `json:"isActive"`
	Experience   string `json:"experience"`
	Format       string `json:"format"`
	Hours        string `json:"hours"`
}

//func (u UpdateVacancy) Validate() error {
//	if u.JobType == "" && u.Title == "" && u.Description =="" && u.Tasks == "" && u.Requirements == "" && u.Extra == "" && u.Salary == "" && u.Location == "" && u.Experience == "" && u.Format == "" && u.Hours == ""{
//		return errorHandler.ErrUpdateStructHasNoValues
//	}
//	return nil
//}
