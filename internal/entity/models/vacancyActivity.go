package models

import "time"

type VacancyActivity struct {
	UserAccountId    uint      `json:"user_account_id" gorm:"primaryKey"`
	ResumeId         uint      `json:"id" gorm:"primaryKey"`
	VacancyId        uint      `json:"vacancy_id" gorm:"primaryKey"`
	ApplicantName    string    `json:"applicant_name"`
	ApplicantSurname string    `json:"applicant_surname"`
	ResumeTitle      string    `json:"title"`
	Image            string    `json:"image"`
	ApplyDate        time.Time `json:"created_date" gorm:"autoCreateTime"`
}

type GetAllAppliesResponce struct {
	Data []*VacancyActivity `json:"data"`
}
