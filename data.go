package jobflow

import (
	"HeadHunter/internal/entity"
)

var Vacancies = map[int]entity.Vacancy{
	1: {
		ID:            1,
		Title:         "Android-разработчик",
		Description:   "Мы разрабатываем сложную рекомендательную систему, делаем приложения под Android и iOS, интегрируем модули Дзена в другие сервисы и пилим свой редактор видео. Все, чтобы авторы нашли аудиторию, а пользователи — то, что им интересно.",
		MinimalSalary: 100000,
		MaximumSalary: 150000,
		EmployerName:  "VK",
		City:          "Moscow",
	},
	2: {
		ID:            2,
		Title:         "Golang разработчик",
		Description:   "Мы приглашаем в свою команду разработчика с отличным знанием Go (golang) для разработки высоконагруженных сервисов, которые помогают тысячам клиентов в области транспорта решать их профессиональные задачи",
		MinimalSalary: 50000,
		MaximumSalary: 150000,
		EmployerName:  "JobFlow",
		City:          "Moscow",
	},
	3: {
		ID:            3,
		Title:         "Middle Golang Developer",
		Description:   "Группа Компаний «Flow inc.» открывает набор на позицию Middle Golang разработчик, в сильный и амбициозный коллектив. Мы - финтех компания, которая разработала платежный шлюз, работающий напрямую с Qiwi, разрабатываем высоконагруженные сервисы с отказоустойстойчивой архитектурой;",
		MinimalSalary: 0,
		MaximumSalary: 350000,
		EmployerName:  "Flow inc.",
		City:          "Moscow",
	},
	4: {
		ID:            4,
		Title:         "Онлайн-репетитор по математике для подготовки к ЕГЭ/ОГЭ",
		Description:   "В онлайн-школу требуется преподаватель математики по подготовке к ЕГЭ, ОГЭ, олимпиадам для проведения занятий с помощью современных образовательных технологий: онлайн-платформы и интерактивных инструментов.",
		MinimalSalary: 48000,
		MaximumSalary: 96000,
		EmployerName:  "Umschool",
		City:          "Voronezh",
	},
	5: {
		ID:            5,
		Title:         "IT-наставник для детей по разработке игр (Minecraft, Roblox, Unity, Construct 3, Unreal Engine)",
		Description:   "В онлайн-школу требуется преподаватель компьютерных курсов для детей по разработке игр",
		MinimalSalary: 0,
		MaximumSalary: 80000,
		EmployerName:  "IT-School",
		City:          "Saint-Petersburg",
	},
}
