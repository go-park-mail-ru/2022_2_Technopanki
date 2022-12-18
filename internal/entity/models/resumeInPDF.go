package models

type ResumeInPDF struct {
	Title             string `json:"title"`
	ApplicantName     string `json:"applicant_name"`
	ApplicantSurname  string `json:"applicant_surname"`
	Location          string `json:"location"`
	ContactNumber     string `json:"contact_number"`
	Email             string `json:"email"`
	Age               uint   `json:"age"`
	ExperienceInYears string `json:"experience"`
	Description       string `json:"description"`
	Image             string `json:"image"`
}
