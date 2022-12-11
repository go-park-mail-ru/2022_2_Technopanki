package models

type Notification struct {
	ID               uint   `json:"id"'`
	NotificationType string `json:"notification_type"`
	UserToID         uint   `json:"user_to_id"`
	UserFromID       uint   `json:"user_from_id"`
}
