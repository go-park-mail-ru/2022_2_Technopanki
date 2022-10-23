package Models

import "time"

type UserAccount struct {
	ID                     uint              `json:"id" gorm:"primaryKey;"`
	UUID                   string            `json:"-"` //TODO убрать
	UserType               string            `json:"user_type" gorm:"not null;"`
	Email                  string            `json:"email" gorm:"not null;"`
	Password               string            `json:"password" gorm:"not null;"`
	ContactNumber          uint              `json:"contact_number" gorm:"not null;"'`
	Description            string            `json:"description" gorm:"not null;"`
	Image                  string            `json:"image"`
	DateOfBirth            time.Time         `json:"date_of_birth" gorm:"not null;"`
	ApplicantName          string            `json:"applicant_name,omitempty"`
	ApplicantSurname       string            `json:"applicant_surname,omitempty"`
	ApplicantCurrentSalary uint              `json:"applicant_current_salary,omitempty"`
	CompanyName            string            `json:"company_name,omitempty"`
	BusinessType           string            `json:"business_type,omitempty"`
	CompanyWebsiteUrl      string            `json:"company_website_url,omitempty"`
	Resumes                []Resume          `json:"resumes" gorm:"foreignKey:UserAccountId;constraint:OnDelete:CASCADE;"`
	Vacancies              []Vacancy         `json:"vacancies" gorm:"foreignKey:PostedByUserId;constraint:OnDelete:CASCADE;"`
	VacancyActivities      []VacancyActivity `json:"vacancy_activities" gorm:"foreignKey:UserAccountId;constraint:OnDelete:CASCADE;"`
}
