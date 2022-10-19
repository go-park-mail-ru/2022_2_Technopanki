package Models

import "time"

type Vacancy struct {
	ID                uint              `json:"id" gorm:"primaryKey;autoIncrement;"`
	PostedByUserId    uint              `json:"postedByUserId" gorm:"not null;"`
	JobType           string            `json:"jobType" gorm:"not null;"`
	CreatedDate       time.Time         `json:"createdDate" gorm:"not null;"`
	JobDescription    string            `json:"jobDescription" gorm:"not null;"`
	JobLocationId     uint              `json:"jobLocationId" gorm:"not null;"`
	IsActive          bool              `json:"isActive" gorm:"not null;"`
	VacancyActivities []VacancyActivity `json:"vacancyActivities" gorm:"foreignKey:VacancyId;constraint:OnDelete:CASCADE;"`
	Skills            []Skill           `json:"skills" gorm:"many2many:vacancy_skills;"`
}
