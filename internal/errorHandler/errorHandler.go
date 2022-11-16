package errorHandler

import (
	"errors"
	"net/http"
)

var (
	ErrResumeNotFound              = errors.New("Резюме не найдено")
	ErrBadRequest                  = errors.New("Некорректный запрос")
	ErrUnauthorized                = errors.New("Клиент не авторизован")
	ErrServiceUnavailable          = errors.New("Сервис недоступен")
	ErrUserExists                  = errors.New("Пользователь с таким email уже существует")
	ErrUserNotExists               = errors.New("Пользователя с таким email не существует")
	ErrInvalidParam                = errors.New("Некорректный параметр")
	ErrCannotCreateUser            = errors.New("Невозможно создать пользователя")
	ErrCannotDeleteVacancy         = errors.New("Невозможно удалить вакансию")
	ErrCannotUpdateVacancy         = errors.New("Невозможно обновить вакансию")
	ErrCannotApplyForVacancy       = errors.New("Невозможно откликнуться на вакансию")
	ErrUpdateStructHasNoValues     = errors.New("Нет значений для обновления")
	ErrCSRFTokenMismatched         = errors.New("Несоответствие CSRF-токена")
	InvalidResumeTitleLength       = errors.New("Длина заголовка резюме должна быть от 3 до 30 символов")
	InvalidResumeDescriptionLength = errors.New("Описание должно быть более подробным")

	ErrCannotCreateSession = errors.New("Невозможно создать сессию")
	ErrSessionNotFound     = errors.New("Сессия с данным токеном не найдена")
	ErrCannotDeleteSession = errors.New("Невозможно удалить сессию")

	ErrVacancyNotFound     = errors.New("Вакансия не найдена")
	ErrCannotDeleteAvatar  = errors.New("Невозможно удалить аватар")
	ErrForbidden           = errors.New("Запрещено")
	ErrWrongPassword       = errors.New("Неправильный пароль")
	ErrInvalidFileFormat   = errors.New("Некорректный формат файла")
	IncorrectNameLength    = errors.New("Длина имени должна быть между 2 и 30 символами")
	IncorrectSurnameLength = errors.New("Длина фамилии должна быть между 2 и 30 символами")

	InvalidEmailFormat   = errors.New("Email должен содержать @")
	IncorrectEmailLength = errors.New("Длина email должна быть между 8 and 30 символами")

	InvalidPasswordFormat   = errors.New("Пароль должен содержать буквы латиницы, цифры и спецсимволы(!#%^$)")
	IncorrectPasswordLength = errors.New("Длина пароля должна быть между 8 и 20 символами")

	InvalidUserType = errors.New("Некорректный входной тип пользователя")
)

var errorToCode = map[error]int{
	ErrResumeNotFound:              http.StatusNotFound,
	ErrBadRequest:                  http.StatusBadRequest,
	ErrUnauthorized:                http.StatusUnauthorized,
	ErrServiceUnavailable:          http.StatusServiceUnavailable,
	ErrUserExists:                  http.StatusBadRequest,
	ErrUserNotExists:               http.StatusUnauthorized,
	ErrInvalidParam:                http.StatusBadRequest,
	ErrCannotCreateUser:            http.StatusServiceUnavailable,
	ErrCannotDeleteVacancy:         http.StatusServiceUnavailable,
	ErrCannotUpdateVacancy:         http.StatusServiceUnavailable,
	ErrCannotCreateSession:         http.StatusInternalServerError,
	ErrSessionNotFound:             http.StatusUnauthorized,
	ErrCannotDeleteSession:         http.StatusInternalServerError,
	ErrCannotDeleteAvatar:          http.StatusBadRequest,
	ErrResumeNotFound:              http.StatusNotFound,
	ErrBadRequest:                  http.StatusBadRequest,
	ErrUnauthorized:                http.StatusUnauthorized,
	ErrServiceUnavailable:          http.StatusServiceUnavailable,
	ErrUserExists:                  http.StatusBadRequest,
	ErrUserNotExists:               http.StatusUnauthorized,
	ErrInvalidParam:                http.StatusBadRequest,
	ErrCannotCreateUser:            http.StatusServiceUnavailable,
	ErrCannotDeleteVacancy:         http.StatusServiceUnavailable,
	ErrCannotUpdateVacancy:         http.StatusServiceUnavailable,
	ErrCannotCreateSession:         http.StatusInternalServerError,
	ErrSessionNotFound:             http.StatusUnauthorized,
	ErrCannotDeleteSession:         http.StatusInternalServerError,
	ErrCannotDeleteAvatar:          http.StatusBadRequest,
	ErrCannotApplyForVacancy:       http.StatusBadRequest,
	InvalidResumeTitleLength:       http.StatusBadRequest,
	InvalidResumeDescriptionLength: http.StatusBadRequest,

	ErrCSRFTokenMismatched: http.StatusForbidden,

	ErrVacancyNotFound:         http.StatusNotFound,
	ErrUpdateStructHasNoValues: http.StatusInternalServerError,

	ErrForbidden:           http.StatusForbidden,
	ErrWrongPassword:       http.StatusBadRequest,
	ErrInvalidFileFormat:   http.StatusBadRequest,
	IncorrectNameLength:    http.StatusBadRequest,
	IncorrectSurnameLength: http.StatusBadRequest,
	InvalidUserType:        http.StatusBadRequest,

	InvalidEmailFormat:   http.StatusBadRequest,
	IncorrectEmailLength: http.StatusBadRequest,

	InvalidPasswordFormat:   http.StatusBadRequest,
	IncorrectPasswordLength: http.StatusBadRequest,
}

func ConvertError(err error) int {
	result, ok := errorToCode[err]
	if ok {
		return result
	}
	return http.StatusInternalServerError
}

var errorDescriptors = map[error]string{
	ErrUserExists:                  "email",
	ErrUserNotExists:               "email",
	ErrWrongPassword:               "password",
	IncorrectNameLength:            "name",
	IncorrectSurnameLength:         "surname",
	InvalidEmailFormat:             "email",
	IncorrectEmailLength:           "email",
	InvalidPasswordFormat:          "password",
	IncorrectPasswordLength:        "password",
	InvalidResumeTitleLength:       "resume_title",
	InvalidResumeDescriptionLength: "resume_descriptors",
}

func GetErrorDescriptors(err error) string {
	result, ok := errorDescriptors[err]
	if ok {
		return result
	}
	return ""
}
