package models

import "time"

type Vacancy struct {
	ID                uint              `json:"id" gorm:"primaryKey;"`
	PostedByUserId    uint              `json:"postedByUserId" gorm:"not null;"`
	Title             string            `json:"title" gorm:"not null;"`
	Description       string            `json:"description,omitempty"`
	Tasks             string            `json:"tasks,omitempty"`
	Requirements      string            `json:"requirements,omitempty"`
	Extra             string            `json:"extra,omitempty"`
	CreatedDate       time.Time         `json:"createdDate" gorm:"autoCreateTime"`
	Salary            string            `json:"salary,omitempty"`
	Location          string            `json:"location,omitempty"`
	IsActive          bool              `json:"isActive,omitempty"`
	Experience        string            `json:"experience,omitempty"`
	Format            string            `json:"format,omitempty"`
	Hours             string            `json:"hours,omitempty"`
	VacancyActivities []VacancyActivity `json:"vacancyActivities" gorm:"foreignKey:VacancyId;constraint:OnDelete:CASCADE;"`
	Skills            []Skill           `json:"skills" gorm:"many2many:vacancy_skills;"`
}

//type UpdateVacancy struct {
//	JobType      string `json:"jobType"`
//	Title        string `json:"title"`
//	Description  string `json:"description"`
//	Tasks        string `json:"tasks"`
//	Requirements string `json:"requirements"`
//	Extra        string `json:"extra"`
//	Salary       string `json:"salary"`
//	Location     string `json:"location"`
//	IsActive     bool   `json:"isActive"`
//	Experience   string `json:"experience"`
//	Format       string `json:"format"`
//	Hours        string `json:"hours"`
//}

//func (u UpdateVacancy) Validate() error {
//	if u.JobType == "" && u.Title == "" && u.Description =="" && u.Tasks == "" && u.Requirements == "" && u.Extra == "" && u.Salary == "" && u.Location == "" && u.Experience == "" && u.Format == "" && u.Hours == ""{
//		return errorHandler.ErrUpdateStructHasNoValues
//	}
//	return nil
//}
