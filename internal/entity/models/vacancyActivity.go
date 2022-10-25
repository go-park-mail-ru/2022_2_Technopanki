package models

import "time"

type VacancyActivity struct {
	UserAccountId uint      `json:"user_account_id" gorm:"primaryKey"`
	VacancyId     uint      `json:"vacancy_id" gorm:"primaryKey"`
	ApplyDate     time.Time `json:"apply_date" gorm:"not null;"`
}
