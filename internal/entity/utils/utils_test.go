package utils

import (
	"HeadHunter/internal/entity/models"
	"HeadHunter/pkg/errorHandler"
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_HasStringArrayElement(t *testing.T) {
	testTable := []struct {
		name       string
		inputArray []string
		inputValue string
		expected   bool
	}{
		{
			name:       "has",
			inputArray: []string{"a", "b", "c"},
			inputValue: "b",
			expected:   true,
		},
		{
			name:       "has not",
			inputArray: []string{"a", "b", "c"},
			inputValue: "d",
			expected:   false,
		},
		{
			name:       "empty array",
			inputArray: []string{},
			inputValue: "d",
			expected:   false,
		},
	}
	for _, testCase := range testTable {
		tc := testCase
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()
			result := HasStringArrayElement(tc.inputValue, tc.inputArray)
			assert.Equal(t, tc.expected, result)
		})

	}
}

func Test_FillUser(t *testing.T) {
	testTable := []struct {
		name           string
		inputUser      *models.UserAccount
		inputReference *models.UserAccount
		expectedUser   *models.UserAccount
		expectedError  error
	}{
		{
			name:      "applicant success",
			inputUser: &models.UserAccount{},
			inputReference: &models.UserAccount{
				ID:               1,
				ApplicantName:    "Ivan",
				ApplicantSurname: "Kozirev",
				Image:            "1.webp",
				UserType:         "applicant",
			},
			expectedUser: &models.UserAccount{
				ID:               1,
				ApplicantName:    "Ivan",
				ApplicantSurname: "Kozirev",
				Image:            "1.webp",
				UserType:         "applicant",
			},
			expectedError: nil,
		},
		{
			name:      "employer success",
			inputUser: &models.UserAccount{},
			inputReference: &models.UserAccount{
				ID:          1,
				CompanyName: "VK",
				Image:       "1.webp",
				UserType:    "employer",
			},
			expectedUser: &models.UserAccount{
				ID:          1,
				CompanyName: "VK",
				Image:       "1.webp",
				UserType:    "employer",
			},
			expectedError: nil,
		},
		{
			name:      "error",
			inputUser: &models.UserAccount{},
			inputReference: &models.UserAccount{
				ID:          1,
				CompanyName: "VK",
				Image:       "1.webp",
				UserType:    "empoyer",
			},
			expectedUser: &models.UserAccount{
				ID:          1,
				CompanyName: "",
				Image:       "1.webp",
				UserType:    "empoyer",
			},
			expectedError: errorHandler.InvalidUserType,
		},
	}
	for _, testCase := range testTable {
		tc := testCase
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()
			err := FillUser(tc.inputUser, tc.inputReference)
			assert.Equal(t, *tc.expectedUser, *tc.inputUser)
			assert.Equal(t, tc.expectedError, err)
		})

	}
}
