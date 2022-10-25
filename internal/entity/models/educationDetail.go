package models

import "time"

type EducationDetail struct {
	ResumeId              uint      `json:"resume_id" gorm:"primaryKey"`
	CertificateDegreeName string    `json:"certificate_degree_name" gorm:"primaryKey"`
	Major                 string    `json:"major" gorm:"primaryKey"`
	UniversityName        string    `json:"university_name" gorm:"not null;"`
	StartingDate          time.Time `json:"starting_date" gorm:"not null;"`
	CompletionDate        time.Time `json:"completion_date"`
}
