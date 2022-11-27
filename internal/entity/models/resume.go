package models

import "time"

type Resume struct {
	ID                uint             `json:"id" gorm:"primaryKey;"`
	UserAccountId     uint             `json:"user_account_id" gorm:"not null;"`
	Title             string           `json:"title" gorm:"not null"`
	Description       string           `json:"description" gorm:"not null;"`
	CreatedTime       time.Time        `json:"created_date" gorm:"autoCreateTime"`
	Location          string           `json:"location,omitempty"`
	ExperienceInYears uint             `json:"experience_in_years,omitempty"`
	Salary            uint             `json:"salary,omitempty"`
	EducationDetail   EducationDetail  `json:"education_detail" gorm:"foreignKey:ResumeId;constraint:OnDelete:CASCADE;"`
	ExperienceDetail  ExperienceDetail `json:"experience_detail" gorm:"foreignKey:ResumeId;constraint:OnDelete:CASCADE;"`
	ApplicantSkills   []Skill          `json:"applicant_skills" gorm:"many2many:resume_skills;"`
}

type ResumePreview struct {
	Image            string    `json:"image"`
	ApplicantName    string    `json:"applicant_name"`
	ApplicantSurname string    `json:"applicant_surname"`
	Id               uint      `json:"id"`
	Title            string    `json:"title"`
	CreatedTime      time.Time `json:"created_date" gorm:"autoCreateTime"`
}

type GetAllResumesResponcePointer struct {
	Data []*Resume `json:"data"`
}

type ResumeFilter struct {
	Title             string
	Location          string
	ExperienceInYears string
	FirstSalaryValue  string
	SecondSalaryValue string
}
