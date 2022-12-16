package models

import "time"

type Notification struct {
	ID          uint      `json:"id" gorm:"primaryKey"`
	Type        string    `json:"type"`
	UserToID    uint      `json:"user_to_id"`
	UserFromID  uint      `json:"user_from_id"`
	ObjectId    uint      `json:"object_id"` //id вакансии либо резюме
	IsViewed    bool      `json:"is_viewed"`
	CreatedTime time.Time `json:"created_time" gorm:"autoCreateTime"`
}

var AllowedNotificationTypes = []string{"apply", "download resume"}

type NotificationPreview struct {
	ID            uint      `json:"id"`
	Type          string    `json:"type"`
	UserToID      uint      `json:"user_to_id"`
	UserFromID    uint      `json:"user_from_id"`
	ObjectId      uint      `json:"object_id"` //id вакансии либо резюме
	ApplicantName string    `json:"applicant_name"`
	Title         string    `json:"title"`
	CompanyName   string    `json:"company_name"`
	IsViewed      bool      `json:"is_viewed"`
	CreatedTime   time.Time `json:"created_time" gorm:"autoCreateTime"`
}
