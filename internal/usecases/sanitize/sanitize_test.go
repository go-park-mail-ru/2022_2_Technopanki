package sanitize

import (
	"HeadHunter/internal/entity/models"
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_SanitizeString(t *testing.T) {
	testTable := []struct {
		name        string
		input       string
		expected    string
		expectedErr error
	}{
		{
			name:        "default",
			input:       "Hello world",
			expected:    "Hello world",
			expectedErr: nil,
		},
		{
			name:        "with html",
			input:       "Hello <STYLE>.XSS{background-image:url();}</STYLE><A CLASS=XSS></A>World",
			expected:    "Hello World",
			expectedErr: nil,
		},
		{
			name:        "empty",
			input:       "",
			expected:    "",
			expectedErr: nil,
		},
	}
	for _, testCase := range testTable {
		tc := testCase
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()
			result, err := SanitizeString(tc.input)
			assert.Equal(t, tc.expected, result)
			assert.Equal(t, tc.expectedErr, err)
		})

	}
}

func Test_SanitizeObject(t *testing.T) {
	testTable := []struct {
		name        string
		input       models.UserAccount
		expected    models.UserAccount
		expectedErr error
	}{
		{
			name: "default",
			input: models.UserAccount{
				Email:  "test@gmail.com",
				Status: "HelloWorld",
			},
			expected: models.UserAccount{
				Email:  "test@gmail.com",
				Status: "HelloWorld",
			},
			expectedErr: nil,
		},
		{
			name: "with html",
			input: models.UserAccount{
				Email:  "test@gmail.com",
				Status: "Hello <STYLE>.XSS{background-image:url();}</STYLE><A CLASS=XSS></A>World",
			},
			expected: models.UserAccount{
				Email:  "test@gmail.com",
				Status: "Hello World",
			},
			expectedErr: nil,
		},
	}
	for _, testCase := range testTable {
		tc := testCase
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()
			result, err := SanitizeObject(&tc.input)
			assert.Equal(t, tc.expected, *result)
			assert.Equal(t, tc.expectedErr, err)
		})

	}
}
