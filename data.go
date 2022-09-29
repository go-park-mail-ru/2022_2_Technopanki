package jobflow

import (
	entity2 "HeadHunter/internal/entity"
	"HeadHunter/internal/errorHandler"
)

func GetEmployer(email, password string) (entity2.Employer, error) {
	if employer, ok := Employers[email]; ok && employer.Password == password {
		return employer, nil
	}
	return errorHandler.ReturnErrorCase[entity2.Employer]("the user is not found")
}

var Employers = map[string]entity2.Employer{
	"vkjob@mail.ru": {
		ID:       "1",
		Name:     "VK",
		Email:    "vkjob@mail.ru",
		Password: "123",
	},
	"yandexjob@yandex.ru": {
		ID:       "2",
		Name:     "Yandex",
		Email:    "yandexjob@yandex.ru",
		Password: "123",
	},
}
var Applicants = []entity2.Applicant{
	{ID: "1",
		Name:     "Zakhar",
		Surname:  "Urvancev",
		Email:    "zahar@gmail.com",
		Password: "123",
	},
	{
		ID:       "2",
		Name:     "Vladislav",
		Surname:  "Kirpichov",
		Email:    "vlad@gmail.com",
		Password: "123",
	},
	{
		ID:       "3",
		Name:     "Sonya",
		Surname:  "Sitnichenko",
		Email:    "sonya@gmail.com",
		Password: "123",
	},
	{
		ID:       "4",
		Name:     "Akim",
		Surname:  "Egorov",
		Email:    "akim@gmail.com",
		Password: "123",
	},
}
var Vacancies = []entity2.Vacancy{
	{
		ID:          "1",
		Title:       "Android-разработчик",
		Description: "Мы разрабатываем сложную рекомендательную систему, делаем приложения под Android и iOS, интегрируем модули Дзена в другие сервисы и пилим свой редактор видео. Все, чтобы авторы нашли аудиторию, а пользователи — то, что им интересно.",
		Salary:      100000,
	},
	{
		ID:          "2",
		Title:       "12312",
		Description: "12412414141",
		Salary:      100000,
	},
}
