package impl

import (
	mock_session "HeadHunter/common/session/mocks"
	"HeadHunter/pkg/errorHandler"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSessionUseCase_GetSession(t *testing.T) {
	type mockBehavior func(s *mock_session.MockRepository, token string)
	testTable := []struct {
		name         string
		mockBehavior mockBehavior
		token        string
		expectEmail  string
		expectError  error
	}{
		{
			name:        "ok",
			token:       "123",
			expectError: nil,
			expectEmail: "some@gmail.com",
			mockBehavior: func(s *mock_session.MockRepository, token string) {
				s.EXPECT().GetSession(token).Return("some@gmail.com", nil)
			},
		},
		{
			name:        "wrong",
			token:       "123",
			expectError: errorHandler.ErrBadRequest,
			expectEmail: "",
			mockBehavior: func(s *mock_session.MockRepository, token string) {
				s.EXPECT().GetSession(token).Return("", errorHandler.ErrBadRequest)
			},
		},
	}
	for _, test := range testTable {
		testCase := test
		t.Run(testCase.name, func(t *testing.T) {
			t.Parallel()
			c := gomock.NewController(t)
			defer c.Finish()

			sessionRep := mock_session.NewMockRepository(c)
			testCase.mockBehavior(sessionRep, testCase.token)
			suc := SessionUseCase{redisStore: sessionRep}
			email, err := suc.GetSession(testCase.token)
			assert.Equal(t, testCase.expectError, err)
			assert.Equal(t, testCase.expectEmail, email)
		})
	}
}

func TestSessionUseCase_NewSession(t *testing.T) {
	type mockBehavior func(s *mock_session.MockRepository, token string)
	testTable := []struct {
		name         string
		mockBehavior mockBehavior
		token        string
		expectEmail  string
		expectError  error
	}{
		{
			name:        "ok",
			token:       "123",
			expectError: nil,
			expectEmail: "some@gmail.com",
			mockBehavior: func(s *mock_session.MockRepository, token string) {
				s.EXPECT().NewSession(token).Return("some@gmail.com", nil)
			},
		},
		{
			name:        "wrong",
			token:       "123",
			expectError: errorHandler.ErrBadRequest,
			expectEmail: "",
			mockBehavior: func(s *mock_session.MockRepository, token string) {
				s.EXPECT().NewSession(token).Return("", errorHandler.ErrBadRequest)
			},
		},
	}
	for _, test := range testTable {
		testCase := test
		t.Run(testCase.name, func(t *testing.T) {
			t.Parallel()
			c := gomock.NewController(t)
			defer c.Finish()

			sessionRep := mock_session.NewMockRepository(c)
			testCase.mockBehavior(sessionRep, testCase.token)
			suc := SessionUseCase{redisStore: sessionRep}
			email, err := suc.NewSession(testCase.token)
			assert.Equal(t, testCase.expectError, err)
			assert.Equal(t, testCase.expectEmail, email)
		})
	}
}

func TestSessionUseCase_DeleteSession(t *testing.T) {
	type mockBehavior func(s *mock_session.MockRepository, token string)
	testTable := []struct {
		name         string
		mockBehavior mockBehavior
		token        string
		expectEmail  string
		expectError  error
	}{
		{
			name:        "ok",
			token:       "123",
			expectError: nil,
			expectEmail: "some@gmail.com",
			mockBehavior: func(s *mock_session.MockRepository, token string) {
				s.EXPECT().Delete(token).Return(nil)
			},
		},
		{
			name:        "wrong",
			token:       "123",
			expectError: errorHandler.ErrBadRequest,
			expectEmail: "",
			mockBehavior: func(s *mock_session.MockRepository, token string) {
				s.EXPECT().Delete(token).Return(errorHandler.ErrBadRequest)
			},
		},
	}
	for _, test := range testTable {
		testCase := test
		t.Run(testCase.name, func(t *testing.T) {
			t.Parallel()
			c := gomock.NewController(t)
			defer c.Finish()

			sessionRep := mock_session.NewMockRepository(c)
			testCase.mockBehavior(sessionRep, testCase.token)
			suc := SessionUseCase{redisStore: sessionRep}
			err := suc.DeleteSession(testCase.token)
			assert.Equal(t, testCase.expectError, err)
		})
	}
}

func TestSessionUseCase_CreateConfirmationCode(t *testing.T) {
	type mockBehavior func(s *mock_session.MockRepository, token string)
	testTable := []struct {
		name         string
		mockBehavior mockBehavior
		token        string
		expectEmail  string
		expectError  error
	}{
		{
			name:        "ok",
			token:       "123",
			expectError: nil,
			expectEmail: "some@gmail.com",
			mockBehavior: func(s *mock_session.MockRepository, token string) {
				s.EXPECT().CreateConfirmationCode(token).Return("some@gmail.com", nil)
			},
		},
		{
			name:        "wrong",
			token:       "123",
			expectError: errorHandler.ErrBadRequest,
			expectEmail: "",
			mockBehavior: func(s *mock_session.MockRepository, token string) {
				s.EXPECT().CreateConfirmationCode(token).Return("", errorHandler.ErrBadRequest)
			},
		},
	}
	for _, test := range testTable {
		testCase := test
		t.Run(testCase.name, func(t *testing.T) {
			t.Parallel()
			c := gomock.NewController(t)
			defer c.Finish()

			sessionRep := mock_session.NewMockRepository(c)
			testCase.mockBehavior(sessionRep, testCase.token)
			suc := SessionUseCase{redisStore: sessionRep}
			email, err := suc.CreateConfirmationCode(testCase.token)
			assert.Equal(t, testCase.expectError, err)
			assert.Equal(t, testCase.expectEmail, email)
		})
	}
}

func TestSessionUseCase_GetCodeFromEmail(t *testing.T) {
	type mockBehavior func(s *mock_session.MockRepository, token string)
	testTable := []struct {
		name         string
		mockBehavior mockBehavior
		token        string
		expectEmail  string
		expectError  error
	}{
		{
			name:        "ok",
			token:       "123",
			expectError: nil,
			expectEmail: "some@gmail.com",
			mockBehavior: func(s *mock_session.MockRepository, token string) {
				s.EXPECT().GetCodeFromEmail(token).Return("some@gmail.com", nil)
			},
		},
		{
			name:        "wrong",
			token:       "123",
			expectError: errorHandler.ErrBadRequest,
			expectEmail: "",
			mockBehavior: func(s *mock_session.MockRepository, token string) {
				s.EXPECT().GetCodeFromEmail(token).Return("", errorHandler.ErrBadRequest)
			},
		},
	}
	for _, test := range testTable {
		testCase := test
		t.Run(testCase.name, func(t *testing.T) {
			t.Parallel()
			c := gomock.NewController(t)
			defer c.Finish()

			sessionRep := mock_session.NewMockRepository(c)
			testCase.mockBehavior(sessionRep, testCase.token)
			suc := SessionUseCase{redisStore: sessionRep}
			email, err := suc.GetCodeFromEmail(testCase.token)
			assert.Equal(t, testCase.expectError, err)
			assert.Equal(t, testCase.expectEmail, email)
		})
	}
}
