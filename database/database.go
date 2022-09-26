package database

import (
	jobflow "HeadHunter"
	"HeadHunter/entity"
	"HeadHunter/errorHandler"
)

func GetEmployer(email, password string) (entity.Employer, error) {
	for _, elem := range jobflow.Employers {
		if elem.Email == email && elem.Password == password {
			return elem, nil
		}
	}
	return errorHandler.ReturnErrorCase[entity.Employer]("the user is not found")
}
