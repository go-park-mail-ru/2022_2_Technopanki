package models

import "time"

type UserAccount struct {
	ID                     uint              `json:"id" gorm:"primaryKey;"`
	UserType               string            `json:"user_type" gorm:"not null;"`
	Email                  string            `json:"email" gorm:"not null;"`
	Password               string            `json:"password" gorm:"not null;"`
	ContactNumber          string            `json:"contact_number" gorm:"not null;"`
	Status                 string            `json:"status" gorm:"not null"`
	Description            string            `json:"description" gorm:"not null;"`
	Image                  string            `json:"image"`
	DateOfBirth            time.Time         `json:"date_of_birth" gorm:"not null"`
	ApplicantName          string            `json:"applicant_name,omitempty"`
	ApplicantSurname       string            `json:"applicant_surname,omitempty"`
	ApplicantCurrentSalary uint              `json:"applicant_current_salary,omitempty"`
	CompanyName            string            `json:"company_name,omitempty"`
	BusinessType           []BusinessType    `json:"business_type,omitempty" gorm:"many2many:business_type;"`
	CompanyWebsiteUrl      string            `json:"company_website_url,omitempty"`
	Location               string            `json:"location,omitempty"`
	CompanySize            uint              `json:"company_size"`
	PublicFields           string            `json:"public_fields"`
	Resumes                []Resume          `json:"resumes" gorm:"foreignKey:UserAccountId;constraint:OnDelete:CASCADE;"`
	Vacancies              []Vacancy         `json:"vacancies" gorm:"foreignKey:PostedByUserId;constraint:OnDelete:CASCADE;"`
	VacancyActivities      []VacancyActivity `json:"vacancy_activities" gorm:"foreignKey:UserAccountId;constraint:OnDelete:CASCADE;"`
}

// PrivateUserFields Поля, доступ к которым по умолчанию не доступен
var PrivateUserFields = []string{"email", "contact_number",
	"applicant_current_salary"}

// SafeUserFields Поля, доступ к которым доступен всегда
var SafeUserFields = []string{"id", "user_type", "description", "status", "date_of_birth", "image",
	"applicant_name", "applicant_surname", "company_name",
	"location", "company_size", "company_website_url"}
