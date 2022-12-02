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
	Age                    uint              `json:"age,omitempty"`
	CreatedTime            time.Time         `json:"created_time" gorm:"autoCreateTime"`
	ApplicantName          string            `json:"applicant_name,omitempty"`
	ApplicantSurname       string            `json:"applicant_surname,omitempty"`
	ApplicantCurrentSalary uint              `json:"applicant_current_salary,omitempty"`
	CompanyName            string            `json:"company_name,omitempty"`
	BusinessType           string            `json:"business_type,omitempty"`
	CompanyWebsiteUrl      string            `json:"company_website_url,omitempty"`
	Location               string            `json:"location,omitempty"`
	CompanySize            uint              `json:"company_size"`
	PublicFields           string            `json:"public_fields"`
	IsConfirmed            bool              `json:"is_confirmed"`
	TwoFactorSignIn        bool              `json:"two_factor_sign_in"`
	MailingApproval        bool              `json:"mailing_approval"`
	Resumes                []Resume          `json:"resumes" gorm:"foreignKey:UserAccountId;constraint:OnDelete:CASCADE;"`
	Vacancies              []Vacancy         `json:"vacancies" gorm:"foreignKey:PostedByUserId;constraint:OnDelete:CASCADE;"`
	VacancyActivities      []VacancyActivity `json:"vacancy_activities" gorm:"foreignKey:UserAccountId;constraint:OnDelete:CASCADE;"`
}

type UserFilter struct {
	ApplicantName          string
	ApplicantSurname       string
	CompanyName            string
	Location               string
	BusinessType           string
	FirstCompanySizeValue  string
	SecondCompanySizeValue string
	DateOfBirth            string
}
type GetAllUsersResponcePointer struct {
	Data []*UserAccount `json:"data"`
}

// PrivateUserFields Поля, доступ к которым по умолчанию не доступен
var PrivateUserFields = []string{"email", "contact_number",
	"applicant_current_salary"}

// SafeUserFields Поля, доступ к которым доступен всегда
var SafeUserFields = []string{"id", "user_type", "description", "status", "date_of_birth", "image",
	"applicant_name", "applicant_surname", "company_name",
	"location", "company_size", "company_website_url", "public_fields", "is_confirmed", "business_type"}

const NoPublicFields string = "null"

type ApplicantPreview struct {
	ID               uint   `json:"id"`
	Image            string `json:"image"`
	ApplicantName    string `json:"applicant_name"`
	ApplicantSurname string `json:"applicant_surname"`
	Status           string `json:"status"`
	Location         string `json:"location"`
}
