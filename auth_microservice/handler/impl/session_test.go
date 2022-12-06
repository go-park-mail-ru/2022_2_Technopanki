package impl

import (
	"HeadHunter/auth_microservice/handler"
	mock_usecase "HeadHunter/auth_microservice/usecase/mocks"
	"HeadHunter/pkg/errorHandler"
	"context"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSessionHandler_NewSession(t *testing.T) {
	ctx := context.Background()
	type mockBehavior func(s *mock_usecase.MockRepository, in *handler.Email)
	testTable := []struct {
		name         string
		in           *handler.Email
		mockBehavior mockBehavior
		expected     *handler.Token
		expectedErr  error
	}{
		{
			name: "ok",
			in:   &handler.Email{Value: "email"},
			mockBehavior: func(s *mock_usecase.MockRepository, in *handler.Email) {
				s.EXPECT().NewSession(in.Value).Return("hello", nil)
			},
			expected:    &handler.Token{Value: "hello"},
			expectedErr: nil,
		},
		{
			name: "wrong",
			in:   &handler.Email{Value: "email"},
			mockBehavior: func(s *mock_usecase.MockRepository, in *handler.Email) {
				s.EXPECT().NewSession(in.Value).Return("", errorHandler.ErrCannotCreateSession)
			},
			expected:    &handler.Token{},
			expectedErr: errorHandler.ErrCannotCreateSession,
		},
		{
			name: "empty in",
			in:   nil,
			mockBehavior: func(s *mock_usecase.MockRepository, in *handler.Email) {
			},
			expected:    &handler.Token{},
			expectedErr: errorHandler.ErrBadRequest,
		},
	}
	for _, test := range testTable {
		testCase := test
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()
			c := gomock.NewController(t)
			defer c.Finish()

			usecase := mock_usecase.NewMockRepository(c)

			testCase.mockBehavior(usecase, testCase.in)
			handler := SessionHandler{sessionUseCase: usecase}
			token, err := handler.NewSession(ctx, testCase.in)
			assert.Equal(t, testCase.expectedErr, err)
			assert.Equal(t, testCase.expected.Value, token.Value)
		})
	}
}

func TestSessionHandler_GetSession(t *testing.T) {
	ctx := context.Background()
	type mockBehavior func(s *mock_usecase.MockRepository, in *handler.Token)
	testTable := []struct {
		name         string
		in           *handler.Token
		mockBehavior mockBehavior
		expected     *handler.Token
		expectedErr  error
	}{
		{
			name: "ok",
			in:   &handler.Token{Value: "email"},
			mockBehavior: func(s *mock_usecase.MockRepository, in *handler.Token) {
				s.EXPECT().GetSession(in.Value).Return("hello", nil)
			},
			expected:    &handler.Token{Value: "hello"},
			expectedErr: nil,
		},
		{
			name: "wrong",
			in:   &handler.Token{Value: "email"},
			mockBehavior: func(s *mock_usecase.MockRepository, in *handler.Token) {
				s.EXPECT().GetSession(in.Value).Return("", errorHandler.ErrCannotCreateSession)
			},
			expected:    &handler.Token{},
			expectedErr: errorHandler.ErrCannotCreateSession,
		},
		{
			name: "empty in",
			in:   nil,
			mockBehavior: func(s *mock_usecase.MockRepository, in *handler.Token) {
			},
			expected:    &handler.Token{},
			expectedErr: errorHandler.ErrBadRequest,
		},
	}
	for _, test := range testTable {
		testCase := test
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()
			c := gomock.NewController(t)
			defer c.Finish()

			usecase := mock_usecase.NewMockRepository(c)

			testCase.mockBehavior(usecase, testCase.in)
			handler := SessionHandler{sessionUseCase: usecase}
			email, err := handler.GetSession(ctx, testCase.in)
			assert.Equal(t, testCase.expectedErr, err)
			assert.Equal(t, testCase.expected.Value, email.Value)
		})
	}
}

func TestSessionHandler_CreateConfirmationCode(t *testing.T) {
	ctx := context.Background()
	type mockBehavior func(s *mock_usecase.MockRepository, in *handler.Email)
	testTable := []struct {
		name         string
		in           *handler.Email
		mockBehavior mockBehavior
		expected     *handler.Token
		expectedErr  error
	}{
		{
			name: "ok",
			in:   &handler.Email{Value: "email"},
			mockBehavior: func(s *mock_usecase.MockRepository, in *handler.Email) {
				s.EXPECT().CreateConfirmationCode(in.Value).Return("hello", nil)
			},
			expected:    &handler.Token{Value: "hello"},
			expectedErr: nil,
		},
		{
			name: "wrong",
			in:   &handler.Email{Value: "email"},
			mockBehavior: func(s *mock_usecase.MockRepository, in *handler.Email) {
				s.EXPECT().CreateConfirmationCode(in.Value).Return("", errorHandler.ErrCannotCreateSession)
			},
			expected:    &handler.Token{},
			expectedErr: errorHandler.ErrCannotCreateSession,
		},
		{
			name: "empty in",
			in:   nil,
			mockBehavior: func(s *mock_usecase.MockRepository, in *handler.Email) {
			},
			expected:    &handler.Token{},
			expectedErr: errorHandler.ErrBadRequest,
		},
	}
	for _, test := range testTable {
		testCase := test
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()
			c := gomock.NewController(t)
			defer c.Finish()

			usecase := mock_usecase.NewMockRepository(c)

			testCase.mockBehavior(usecase, testCase.in)
			handler := SessionHandler{sessionUseCase: usecase}
			token, err := handler.CreateConfirmationCode(ctx, testCase.in)
			assert.Equal(t, testCase.expectedErr, err)
			assert.Equal(t, testCase.expected.Value, token.Value)
		})
	}
}

func TestSessionHandler_DeleteSession(t *testing.T) {
	ctx := context.Background()
	type mockBehavior func(s *mock_usecase.MockRepository, in *handler.Token)
	testTable := []struct {
		name         string
		in           *handler.Token
		mockBehavior mockBehavior
		expected     *handler.Token
		expectedErr  error
	}{
		{
			name: "ok",
			in:   &handler.Token{Value: "email"},
			mockBehavior: func(s *mock_usecase.MockRepository, in *handler.Token) {
				s.EXPECT().DeleteSession(in.Value).Return(nil)
			},
			expected:    &handler.Token{Value: "hello"},
			expectedErr: nil,
		},
		{
			name: "wrong",
			in:   &handler.Token{Value: "email"},
			mockBehavior: func(s *mock_usecase.MockRepository, in *handler.Token) {
				s.EXPECT().DeleteSession(in.Value).Return(errorHandler.ErrCannotCreateSession)
			},
			expected:    &handler.Token{},
			expectedErr: errorHandler.ErrCannotCreateSession,
		},
		{
			name: "empty in",
			in:   nil,
			mockBehavior: func(s *mock_usecase.MockRepository, in *handler.Token) {
			},
			expected:    &handler.Token{},
			expectedErr: errorHandler.ErrBadRequest,
		},
	}
	for _, test := range testTable {
		testCase := test
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()
			c := gomock.NewController(t)
			defer c.Finish()

			usecase := mock_usecase.NewMockRepository(c)

			testCase.mockBehavior(usecase, testCase.in)
			handler := SessionHandler{sessionUseCase: usecase}
			_, err := handler.DeleteSession(ctx, testCase.in)
			assert.Equal(t, testCase.expectedErr, err)
		})
	}
}

func TestSessionHandler_GetCodeFromEmail(t *testing.T) {
	ctx := context.Background()
	type mockBehavior func(s *mock_usecase.MockRepository, in *handler.Email)
	testTable := []struct {
		name         string
		in           *handler.Email
		mockBehavior mockBehavior
		expected     *handler.Token
		expectedErr  error
	}{
		{
			name: "ok",
			in:   &handler.Email{Value: "email"},
			mockBehavior: func(s *mock_usecase.MockRepository, in *handler.Email) {
				s.EXPECT().GetCodeFromEmail(in.Value).Return("hello", nil)
			},
			expected:    &handler.Token{Value: "hello"},
			expectedErr: nil,
		},
		{
			name: "wrong",
			in:   &handler.Email{Value: "email"},
			mockBehavior: func(s *mock_usecase.MockRepository, in *handler.Email) {
				s.EXPECT().GetCodeFromEmail(in.Value).Return("", errorHandler.ErrCannotCreateSession)
			},
			expected:    &handler.Token{},
			expectedErr: errorHandler.ErrCannotCreateSession,
		},
		{
			name: "empty in",
			in:   nil,
			mockBehavior: func(s *mock_usecase.MockRepository, in *handler.Email) {
			},
			expected:    &handler.Token{},
			expectedErr: errorHandler.ErrBadRequest,
		},
	}
	for _, test := range testTable {
		testCase := test
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()
			c := gomock.NewController(t)
			defer c.Finish()

			usecase := mock_usecase.NewMockRepository(c)

			testCase.mockBehavior(usecase, testCase.in)
			handler := SessionHandler{sessionUseCase: usecase}
			email, err := handler.GetCodeFromEmail(ctx, testCase.in)
			assert.Equal(t, testCase.expectedErr, err)
			assert.Equal(t, testCase.expected.Value, email.Value)
		})
	}
}
