package models

import (
	"HeadHunter/internal/errorHandler"
	"time"
)

type Vacancy struct {
	ID                uint              `json:"id" gorm:"primaryKey;"`
	PostedByUserId    uint              `json:"postedByUserId" gorm:"not null;"`
	JobType           string            `json:"jobType" gorm:"not null;"`
	CreatedDate       time.Time         `json:"createdDate" gorm:"not null;"`
	JobDescription    string            `json:"jobDescription" gorm:"not null;"`
	JobLocationId     uint              `json:"jobLocationId" gorm:"not null;"`
	IsActive          bool              `json:"isActive" gorm:"not null;"`
	VacancyActivities []VacancyActivity `json:"vacancyActivities" gorm:"foreignKey:VacancyId;constraint:OnDelete:CASCADE;"`
	Skills            []Skill           `json:"skills" gorm:"many2many:vacancy_skills;"`
}

type UpdateVacancy struct {
	JobType        *string `json:"jobType"`
	JobDescription *string `json:"jobDescription"`
	JobLocationId  *uint   `json:"jobLocationId"`
	IsActive       *bool   `json:"isActive"`
}

func (u UpdateVacancy) Validate() error {
	if u.JobType == nil && u.JobDescription == nil && u.JobLocationId == nil && u.IsActive == nil {
		return errorHandler.ErrUpdateStructHasNoValues
	}
	return nil
}
