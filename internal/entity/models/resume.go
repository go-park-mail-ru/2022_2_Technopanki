package models

type Resume struct {
	ID               uint             `json:"id" gorm:"primaryKey;"`
	UserAccountId    uint             `json:"user_account_id" gorm:"not null;"`
	Description      string           `json:"description" gorm:"not null;"`
	EducationDetail  EducationDetail  `json:"education_detail" gorm:"foreignKey:ResumeId;constraint:OnDelete:CASCADE;"`
	ExperienceDetail ExperienceDetail `json:"experience_detail" gorm:"foreignKey:ResumeId;constraint:OnDelete:CASCADE;"`
	ApplicantSkills  []Skill          `json:"applicant_skills" gorm:"many2many:resume_skills;"`
}
