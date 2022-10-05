package sessions

import (
	"HeadHunter/internal/errorHandler"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"testing"
)

func Test_GetSession(t *testing.T) {
	testTable := []struct {
		name          string
		token         Token
		expected      Session
		expectedError error
	}{
		{
			name:  "contained case",
			token: "someToken",
			expected: Session{
				Email: "test@mail.ru",
			},
			expectedError: nil,
		},
		{
			name:          "not contained case",
			token:         "someToken2",
			expected:      Session{},
			expectedError: errorHandler.ErrSessionNotFound,
		},
	}
	testStore := Store{
		Values: map[Token]Session{
			"someToken": {
				Email: "test@mail.ru",
			},
		},
	}
	for _, testCase := range testTable {
		tc := testCase
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()
			result, errResult := testStore.GetSession(tc.token)
			assert.Equal(t, tc.expected, result)
			assert.Equal(t, tc.expectedError, errResult)
		})

	}
}

func Test_NewSession(t *testing.T) {
	testTable := []struct {
		name  string
		email string
	}{
		{
			name:  "default case",
			email: "test@mail.ru",
		},
	}
	testStore := Store{
		Values: map[Token]Session{
			"someToken": {
				Email: "test@mail.ru",
			},
		},
	}
	for _, testCase := range testTable {
		tc := testCase
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()
			token := testStore.NewSession(tc.email)
			result, errResult := testStore.GetSession(Token(token))
			require.NoError(t, errResult)
			assert.Equal(t, tc.email, result.Email)
		})

	}
}

func Test_DeleteSession(t *testing.T) {
	testTable := []struct {
		name     string
		token    Token
		expected error
	}{
		{
			name:     "can delete",
			token:    "someToken",
			expected: nil,
		},
		{
			name:     "can't delete",
			token:    "someToken2",
			expected: errorHandler.ErrCannotDeleteSession,
		},
	}
	testStore := Store{
		Values: map[Token]Session{
			"someToken": {
				Email: "test@mail.ru",
			},
		},
	}
	for _, testCase := range testTable {
		tc := testCase
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()
			result := testStore.DeleteSession(tc.token)
			assert.Equal(t, tc.expected, result)
		})

	}
}
