package Models

import "time"

type ExperienceDetail struct {
	ResumeId        uint      `json:"resume_id" gorm:"primaryKey"`
	IsCurrentJob    string    `json:"is_current_job" gorm:"not null;"`
	StartDate       time.Time `json:"start_date" gorm:"not null;"`
	EndDate         time.Time `json:"end_date"`
	JobTitle        string    `json:"job_title" gorm:"not null;"`
	CompanyName     string    `json:"company_name" gorm:"not null;"`
	JobLocationCity string    `json:"job_location_city" gorm:"not null;"`
	Description     string    `json:"description" gorm:"not null;"`
}
