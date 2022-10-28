package validation

import (
	"HeadHunter/internal/entity/models"
	"HeadHunter/internal/errorHandler"
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_VerifyPassword(t *testing.T) {
	testTable := []struct {
		name     string
		password string
		expected bool
	}{
		{
			name:     "verified password",
			password: "123456ab!",
			expected: true,
		},
		{
			name:     "password without latin symbols",
			password: "123456!",
			expected: false,
		},
		{
			name:     "password without numbers",
			password: "ab!",
			expected: false,
		},
		{
			name:     "password without spec-symbols",
			password: "abasd123",
			expected: false,
		},
		{
			name:     "password with incorrect symbols",
			password: "??????",
			expected: false,
		},
	}
	for _, testCase := range testTable {
		tc := testCase
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()
			result := verifyPassword(tc.password)
			assert.Equal(t, tc.expected, result)
		})

	}
}

func Test_IsValidateAuthData(t *testing.T) {
	testTable := []struct {
		name     string
		user     models.UserAccount
		expected error
	}{
		{
			name: "valid user",
			user: models.UserAccount{
				Email:    "example@mail.com",
				Password: "123456!a",
			},
			expected: nil,
		},
		{
			name: "invalid email format",
			user: models.UserAccount{
				Email:    "examplemail.com",
				Password: "123456!a",
			},
			expected: errorHandler.InvalidEmailFormat,
		},
		{
			name: "incorrect email length",
			user: models.UserAccount{
				Email:    "xmp@ml",
				Password: "123456!a",
			},
			expected: errorHandler.IncorrectEmailLength,
		},
		{
			name: "invalid password format",
			user: models.UserAccount{
				Email:    "example@mail.com",
				Password: "123456a",
			},
			expected: errorHandler.InvalidPasswordFormat,
		},
		{
			name: "incorrect password length",
			user: models.UserAccount{
				Email:    "example@mail.com",
				Password: "1a!",
			},
			expected: errorHandler.IncorrectPasswordLength,
		},
	}
	for _, testCase := range testTable {
		tc := testCase
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()
			result := IsAuthDataValid(tc.user)
			assert.Equal(t, tc.expected, result)
		})

	}
}

func Test_IsValidate(t *testing.T) {
	testTable := []struct {
		name     string
		user     models.UserAccount
		expected error
	}{
		{
			name: "valid user",
			user: models.UserAccount{
				CompanyName: "Mail.ru",
				Email:       "example@mail.com",
				Password:    "123456!a",
				UserType:    "employer",
			},
			expected: nil,
		},
		{
			name: "valid user 2",
			user: models.UserAccount{
				ApplicantName:    "Zakhar",
				ApplicantSurname: "Urvancev",
				Email:            "example@mail.com",
				Password:         "123456!a",
				UserType:         "applicant",
			},
			expected: nil,
		},
		{
			name: "invalid user role",
			user: models.UserAccount{
				ApplicantName:    "Zakhar",
				ApplicantSurname: "Urvancev",
				Email:            "example@mail.com",
				Password:         "123456!a",
				UserType:         "some_role)",
			},
			expected: errorHandler.InvalidUserType,
		},
		{
			name: "incorrect name length 1",
			user: models.UserAccount{
				ApplicantName:    "Z",
				ApplicantSurname: "Urvancev",
				Email:            "example@mail.com",
				Password:         "123456!a",
				UserType:         "applicant",
			},
			expected: errorHandler.IncorrectNameLength,
		},
		{
			name: "incorrect name length 2",
			user: models.UserAccount{
				ApplicantName:    "Zaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaahar",
				ApplicantSurname: "Urvancev",
				Email:            "example@mail.com",
				Password:         "123456!a",
				UserType:         "applicant",
			},
			expected: errorHandler.IncorrectNameLength,
		},
		{
			name: "incorrect surname length 1",
			user: models.UserAccount{
				ApplicantName:    "Zakhar",
				ApplicantSurname: "U",
				Email:            "example@mail.com",
				Password:         "123456!a",
				UserType:         "applicant",
			},
			expected: errorHandler.IncorrectSurnameLength,
		},
		{
			name: "incorrect surname length 2",
			user: models.UserAccount{
				ApplicantName:    "Zakhar",
				ApplicantSurname: "Urvaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaancev",
				Email:            "example@mail.com",
				Password:         "123456!a",
				UserType:         "applicant",
			},
			expected: errorHandler.IncorrectSurnameLength,
		},
	}
	for _, testCase := range testTable {
		tc := testCase
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()
			result := IsUserValid(tc.user)
			assert.Equal(t, tc.expected, result)
		})

	}
}
