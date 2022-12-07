package models

import "time"

type VacancyActivity struct {
	UserAccountId uint      `json:"user_account_id" gorm:"primaryKey"`
	ResumeId      uint      `json:"id" gorm:"primaryKey"`
	VacancyId     uint      `json:"vacancy_id" gorm:"primaryKey"`
	ApplyDate     time.Time `json:"created_date" gorm:"autoCreateTime"`
}

type GetAllAppliesResponce struct {
	Data []*VacancyActivityPreview `json:"data"`
}

type VacancyActivityPreview struct {
	UserAccountId    uint      `json:"user_account_id"`
	ResumeId         uint      `json:"id"`
	VacancyId        uint      `json:"vacancy_id"`
	ApplicantName    string    `json:"applicant_name"`
	ApplicantSurname string    `json:"applicant_surname"`
	ResumeTitle      string    `json:"title"`
	Image            string    `json:"image"`
	ApplyDate        time.Time `json:"created_date"`
}
