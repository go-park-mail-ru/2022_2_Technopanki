package models

type Notification struct {
	ID         uint   `json:"id"`
	Type       string `json:"type"`
	UserToID   uint   `json:"user_to_id"`
	UserFromID uint   `json:"user_from_id"`
}

var AllowedNotificationTypes = []string{"apply", "download resume"}

type NotificationPreview struct {
	ID            uint   `json:"id"`
	Type          string `json:"type"`
	UserToID      uint   `json:"user_to_id"`
	UserFromID    uint   `json:"user_from_id"`
	ApplicantName string `json:"applicant_name"`
	Title         string `json:"title"`
	CompanyName   string `json:"company_name"`
}
