//go:generate easyjson -all vacancy.go
package models

import "time"

//easyjson:json
type Vacancy struct {
	ID                uint              `json:"id" gorm:"primaryKey;"`
	PostedByUserId    uint              `json:"postedByUserId" gorm:"not null;"`
	Title             string            `json:"title" gorm:"not null;"`
	Description       string            `json:"description,omitempty"`
	Tasks             string            `json:"tasks,omitempty"`
	Requirements      string            `json:"requirements,omitempty"`
	Extra             string            `json:"extra,omitempty"`
	CreatedDate       time.Time         `json:"createdDate" gorm:"autoCreateTime"`
	Salary            uint              `json:"salary,omitempty"`
	Location          string            `json:"location,omitempty"`
	IsActive          bool              `json:"isActive,omitempty"`
	Experience        string            `json:"experience,omitempty"`
	Format            string            `json:"format,omitempty"`
	Hours             string            `json:"hours,omitempty"`
	Image             string            `json:"image,omitempty"`
	VacancyActivities []VacancyActivity `json:"vacancyActivities" gorm:"foreignKey:VacancyId;constraint:OnDelete:CASCADE;"`
	Skills            []Skill           `json:"skills" gorm:"many2many:vacancy_skills;"`
}

//easyjson:json
type GetAllVacanciesResponcePointer struct {
	Data []*Vacancy `json:"data"`
}

//easyjson:json
type VacancyFilter struct {
	Title             string
	Location          string
	Format            string
	Experience        string
	FirstSalaryValue  string
	SecondSalaryValue string
}

//easyjson:json
type VacancyPreview struct {
	Id          uint   `json:"id"`
	Title       string `json:"title"`
	Image       string `json:"image"`
	Salary      uint   `json:"salary"`
	Location    string `json:"location"`
	Format      string `json:"format"`
	Hours       string `json:"hours"`
	Description string `json:"description"`
}

//easyjson:json
type VacancyPreviewsResponse struct {
	Data []*VacancyPreview `json:"data"`
}
