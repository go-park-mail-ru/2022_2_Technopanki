package complexModels

type ResumePreview struct {
	Image            string `json:"image"`
	ApplicantName    string `json:"applicant_name"`
	ApplicantSurname string `json:"applicant_surname"`
	Id               int    `json:"id"`
	Title            string `json:"title"`
}
