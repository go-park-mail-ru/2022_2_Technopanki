package models

type JobLocation struct {
	ID            uint      `json:"id" gorm:"primaryKey;"`
	StreetAddress string    `json:"streetAddress" gorm:"not null;"`
	City          string    `json:"city" gorm:"not null;"`
	Country       string    `json:"country" gorm:"not null;"`
	Vacancies     []Vacancy `json:"vacancies" gorm:"foreignKey:JobLocationId;constraint:OnDelete:CASCADE;"`
}
