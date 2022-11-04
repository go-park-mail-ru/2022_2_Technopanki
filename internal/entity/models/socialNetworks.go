package models

type SocialNetworks struct {
	ID            string `json:"id" gorm:"primaryKey;"`
	UserAccountId string `json:"user_account_id" gorm:"not null;"`
	VK            string `json:"vk,omitempty"`
	Telegram      string `json:"telegram,omitempty"`
	Facebook      string `json:"facebook,omitempty"`
	YouTube       string `json:"youtube,omitempty"`
	Twitter       string `json:"twitter,omitempty"`
	Instagram     string `json:"instagram,omitempty"`
}
