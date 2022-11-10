package models

import "time"

type Resume struct {
	ID               uint             `json:"resume_id" gorm:"primaryKey;"`
	UserAccountId    uint             `json:"user_account_id" gorm:"not null;"`
	Title            string           `json:"title" gorm:"not null"`
	Description      string           `json:"description" gorm:"not null;"`
	UserName         string           `json:"user_name"`
	UserSurname      string           `json:"user_surname"`
	ImgSrc           string           `json:"imgSrc"`
	CreatedTime      time.Time        `json:"created_date" gorm:"autoCreateTime"`
	EducationDetail  EducationDetail  `json:"education_detail" gorm:"foreignKey:ResumeId;constraint:OnDelete:CASCADE;"`
	ExperienceDetail ExperienceDetail `json:"experience_detail" gorm:"foreignKey:ResumeId;constraint:OnDelete:CASCADE;"`
	ApplicantSkills  []Skill          `json:"applicant_skills" gorm:"many2many:resume_skills;"`
}
