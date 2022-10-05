package validation

import (
	"HeadHunter/internal/entity"
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
		user     entity.User
		expected error
	}{
		{
			name: "valid user",
			user: entity.User{
				Email:    "example@mail.com",
				Password: "123456!a",
			},
			expected: nil,
		},
		{
			name: "invalid email format",
			user: entity.User{
				Email:    "examplemail.com",
				Password: "123456!a",
			},
			expected: errorHandler.InvalidEmailFormat,
		},
		{
			name: "incorrect email length",
			user: entity.User{
				Email:    "xmp@ml",
				Password: "123456!a",
			},
			expected: errorHandler.IncorrectEmailLength,
		},
		{
			name: "invalid password format",
			user: entity.User{
				Email:    "example@mail.com",
				Password: "123456a",
			},
			expected: errorHandler.InvalidPasswordFormat,
		},
		{
			name: "incorrect password length",
			user: entity.User{
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
			result := IsValidateAuthData(tc.user)
			assert.Equal(t, tc.expected, result)
		})

	}
}

func Test_IsValidate(t *testing.T) {
	testTable := []struct {
		name     string
		user     entity.User
		expected error
	}{
		{
			name: "valid user",
			user: entity.User{
				Name:     "Zakhar",
				Surname:  "Urvancev",
				Email:    "example@mail.com",
				Password: "123456!a",
				Role:     "employer",
			},
			expected: nil,
		},
		{
			name: "valid user 2",
			user: entity.User{
				Name:     "Zakhar",
				Surname:  "Urvancev",
				Email:    "example@mail.com",
				Password: "123456!a",
				Role:     "applicant",
			},
			expected: nil,
		},
		{
			name: "invalid user role",
			user: entity.User{
				Name:     "Zakhar",
				Surname:  "Urvancev",
				Email:    "example@mail.com",
				Password: "123456!a",
				Role:     "some_role)",
			},
			expected: errorHandler.InvalidUserRole,
		},
		{
			name: "incorrect name length 1",
			user: entity.User{
				Name:     "Z",
				Surname:  "Urvancev",
				Email:    "example@mail.com",
				Password: "123456!a",
			},
			expected: errorHandler.IncorrectNameLength,
		},
		{
			name: "incorrect name length 2",
			user: entity.User{
				Name:     "Zaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaahar",
				Surname:  "Urvancev",
				Email:    "example@mail.com",
				Password: "123456!a",
			},
			expected: errorHandler.IncorrectNameLength,
		},
		{
			name: "incorrect surname length 1",
			user: entity.User{
				Name:     "Zakhar",
				Surname:  "U",
				Email:    "example@mail.com",
				Password: "123456!a",
			},
			expected: errorHandler.IncorrectSurnameLength,
		},
		{
			name: "incorrect surname length 2",
			user: entity.User{
				Name:     "Zakhar",
				Surname:  "Urvaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaancev",
				Email:    "example@mail.com",
				Password: "123456!a",
			},
			expected: errorHandler.IncorrectSurnameLength,
		},
	}
	for _, testCase := range testTable {
		tc := testCase
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()
			result := IsValidate(tc.user)
			assert.Equal(t, tc.expected, result)
		})

	}
}
