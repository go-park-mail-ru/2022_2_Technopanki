package models

import "time"

type VacancyActivity struct {
	UserAccountId uint      `json:"user_account_id" gorm:"primaryKey"`
	ResumeId      uint      `json:"resume_id" gorm:"primaryKey"`
	VacancyId     uint      `json:"vacancy_id" gorm:"primaryKey"`
	UserName      string    `json:"user_name"`
	UserSurname   string    `json:"user_surname"`
	ResumeTitle   string    `json:"title"`
	Image         string    `json:"image"`
	ApplyDate     time.Time `json:"created_date" gorm:"autoCreateTime"`
}

type GetAllAppliesResponce struct {
	Data []*VacancyActivity `json:"data"`
}
