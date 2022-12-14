package impl

import (
	"HeadHunter/common/session/mocks"
	"HeadHunter/configs"
	"HeadHunter/internal/entity/models"
	mock_repository "HeadHunter/internal/repository/mocks"
	mock_usecase "HeadHunter/mail_microservice/usecase/mocks"
	"HeadHunter/pkg/errorHandler"
	"fmt"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestUserService_GetUserByEmail(t *testing.T) {
	type mockBehavior func(r *mock_repository.MockUserRepository, email string)

	testTable := []struct {
		name         string
		email        string
		mockBehavior mockBehavior
		expected     *models.UserAccount
		expectedErr  error
	}{
		{
			name:  "ok",
			email: "test@gmail.com",
			expected: &models.UserAccount{
				Email:       "test@gmail.com",
				Password:    "123456A!",
				CompanyName: "VK",
				UserType:    "employer",
			},
			mockBehavior: func(r *mock_repository.MockUserRepository, email string) {
				expected := &models.UserAccount{
					Email:       "test@gmail.com",
					Password:    "123456A!",
					CompanyName: "VK",
					UserType:    "employer",
				}
				r.EXPECT().GetUserByEmail(email).Return(expected, nil)
			},
			expectedErr: nil,
		},
		{
			name:     "error",
			email:    "test@gmail.com",
			expected: &models.UserAccount{},
			mockBehavior: func(r *mock_repository.MockUserRepository, email string) {
				expected := &models.UserAccount{}
				r.EXPECT().GetUserByEmail(email).Return(expected, errorHandler.ErrUserNotExists)
			},
			expectedErr: errorHandler.ErrUserNotExists,
		},
	}

	for _, test := range testTable {
		testCase := test
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()
			c := gomock.NewController(t)
			defer c.Finish()

			mockUserRepository := mock_repository.NewMockUserRepository(c)
			testCase.mockBehavior(mockUserRepository, test.email)
			userService := UserService{userRep: mockUserRepository}
			result, err := userService.GetUserByEmail(testCase.email)

			assert.Equal(t, testCase.expected, result)
			assert.Equal(t, testCase.expectedErr, err)
		})
	}
} //100%

func TestUserService_SignIn(t *testing.T) { //%100
	type mockBehavior func(r *mock_repository.MockUserRepository, email string)
	type sessionRepBehavior func(r *mock_session.MockRepository, email string)
	type mailBehavior func(r *mock_usecase.MockMail, email string)
	testTable := []struct {
		name               string
		inputUser          *models.UserAccount
		mockBehavior       mockBehavior
		sessionRepBehavior sessionRepBehavior
		mailBehavior       mailBehavior
		expectedToken      string
		expectedUser       *models.UserAccount
		expectedErr        error
		confirmMode        bool
	}{
		{
			name: "ok",
			inputUser: &models.UserAccount{
				Email:    "test@gmail.com",
				Password: "123456A!",
			},
			expectedToken: "valid_token",
			expectedUser: &models.UserAccount{
				Email:       "test@gmail.com",
				Password:    "123456A!",
				CompanyName: "VK",
				UserType:    "employer",
			},
			mockBehavior: func(r *mock_repository.MockUserRepository, email string) {
				expected := &models.UserAccount{
					Email:       "test@gmail.com",
					Password:    "$2a$10$AN0zyrlPjIZvO12mLf8PierYf579fxzfqh6j9aZn3LWetXMjJjYb.",
					CompanyName: "VK",
					UserType:    "employer",
				}
				r.EXPECT().GetUserByEmail(email).Return(expected, nil)
			},
			sessionRepBehavior: func(r *mock_session.MockRepository, email string) {
				r.EXPECT().NewSession(email).Return("valid_token", nil)
			},
			mailBehavior: func(r *mock_usecase.MockMail, email string) {},
			expectedErr:  nil,
		},
		{
			name: "ok (confirmation mode)",
			inputUser: &models.UserAccount{
				Email:    "test@gmail.com",
				Password: "123456A!",
			},
			expectedToken: "valid_token",
			expectedUser: &models.UserAccount{
				Email:       "test@gmail.com",
				Password:    "123456A!",
				CompanyName: "VK",
				UserType:    "employer",
				IsConfirmed: true,
			},
			mockBehavior: func(r *mock_repository.MockUserRepository, email string) {
				expected := &models.UserAccount{
					Email:       "test@gmail.com",
					Password:    "$2a$10$AN0zyrlPjIZvO12mLf8PierYf579fxzfqh6j9aZn3LWetXMjJjYb.",
					CompanyName: "VK",
					UserType:    "employer",
					IsConfirmed: true,
				}
				r.EXPECT().GetUserByEmail(email).Return(expected, nil)
			},
			sessionRepBehavior: func(r *mock_session.MockRepository, email string) {
				r.EXPECT().NewSession(email).Return("valid_token", nil)
			},
			mailBehavior: func(r *mock_usecase.MockMail, email string) {},
			confirmMode:  true,
			expectedErr:  nil,
		},
		{
			name: "forbidden (two-factor)",
			inputUser: &models.UserAccount{
				Email:           "test@gmail.com",
				Password:        "123456A!",
				TwoFactorSignIn: true,
				IsConfirmed:     true,
			},
			confirmMode:   true,
			expectedToken: "valid_token",
			expectedUser: &models.UserAccount{
				Email:           "test@gmail.com",
				Password:        "123456A!",
				CompanyName:     "VK",
				UserType:        "employer",
				TwoFactorSignIn: true,
				IsConfirmed:     true,
			},
			mockBehavior: func(r *mock_repository.MockUserRepository, email string) {
				expected := &models.UserAccount{
					Email:           "test@gmail.com",
					Password:        "$2a$10$AN0zyrlPjIZvO12mLf8PierYf579fxzfqh6j9aZn3LWetXMjJjYb.",
					CompanyName:     "VK",
					UserType:        "employer",
					TwoFactorSignIn: true,
					IsConfirmed:     true,
				}
				r.EXPECT().GetUserByEmail(email).Return(expected, nil)
			},
			sessionRepBehavior: func(r *mock_session.MockRepository, email string) {
			},
			mailBehavior: func(r *mock_usecase.MockMail, email string) {
				r.EXPECT().SendConfirmCode(email).Return(nil)
			},
			expectedErr: nil,
		},
		{
			name: "send error (two-factor)",
			inputUser: &models.UserAccount{
				Email:           "test@gmail.com",
				Password:        "123456A!",
				TwoFactorSignIn: true,
				IsConfirmed:     true,
			},
			confirmMode:   true,
			expectedToken: "valid_token",
			expectedUser: &models.UserAccount{
				Email:           "test@gmail.com",
				Password:        "123456A!",
				CompanyName:     "VK",
				UserType:        "employer",
				TwoFactorSignIn: true,
				IsConfirmed:     true,
			},
			mockBehavior: func(r *mock_repository.MockUserRepository, email string) {
				expected := &models.UserAccount{
					Email:           "test@gmail.com",
					Password:        "$2a$10$AN0zyrlPjIZvO12mLf8PierYf579fxzfqh6j9aZn3LWetXMjJjYb.",
					CompanyName:     "VK",
					UserType:        "employer",
					TwoFactorSignIn: true,
					IsConfirmed:     true,
				}
				r.EXPECT().GetUserByEmail(email).Return(expected, nil)
			},
			sessionRepBehavior: func(r *mock_session.MockRepository, email string) {
			},
			mailBehavior: func(r *mock_usecase.MockMail, email string) {
				r.EXPECT().SendConfirmCode(email).Return(errorHandler.ErrBadRequest)
			},
			expectedErr: errorHandler.ErrBadRequest,
		},
		{
			name: "user is not confirmed",
			inputUser: &models.UserAccount{
				Email:    "test@gmail.com",
				Password: "123456A!",
			},
			expectedToken: "valid_token",
			expectedUser: &models.UserAccount{
				Email:       "test@gmail.com",
				Password:    "123456A!",
				CompanyName: "VK",
				UserType:    "employer",
			},
			mockBehavior: func(r *mock_repository.MockUserRepository, email string) {
				expected := &models.UserAccount{
					Email:       "test@gmail.com",
					Password:    "$2a$10$AN0zyrlPjIZvO12mLf8PierYf579fxzfqh6j9aZn3LWetXMjJjYb.",
					CompanyName: "VK",
					UserType:    "employer",
				}
				r.EXPECT().GetUserByEmail(email).Return(expected, nil)
			},
			sessionRepBehavior: func(r *mock_session.MockRepository, email string) {},
			mailBehavior:       func(r *mock_usecase.MockMail, email string) {},
			expectedErr:        errorHandler.ErrIsNotConfirmed,
			confirmMode:        true,
		},
		{
			name: "wrong answer",
			inputUser: &models.UserAccount{
				Email:    "test@gmail.com",
				Password: "123456A!",
			},
			expectedToken: "valid_token",
			expectedUser:  nil,
			mockBehavior: func(r *mock_repository.MockUserRepository, email string) {
				expected := &models.UserAccount{
					Email:       "test@gmail.com",
					Password:    "123456A!",
					CompanyName: "VK",
					UserType:    "employer",
				}
				r.EXPECT().GetUserByEmail(email).Return(expected, nil)
			},
			sessionRepBehavior: func(r *mock_session.MockRepository, email string) {},
			mailBehavior:       func(r *mock_usecase.MockMail, email string) {},
			expectedErr:        errorHandler.ErrWrongPassword,
		},
		{
			name: "user not found",
			inputUser: &models.UserAccount{
				Email:    "test@gmail.com",
				Password: "123456A!",
			},
			expectedToken: "valid_token",
			expectedUser:  nil,
			mockBehavior: func(r *mock_repository.MockUserRepository, email string) {
				r.EXPECT().GetUserByEmail(email).Return(nil, errorHandler.ErrUserNotExists)
			},
			sessionRepBehavior: func(r *mock_session.MockRepository, email string) {},
			mailBehavior:       func(r *mock_usecase.MockMail, email string) {},
			expectedErr:        errorHandler.ErrUserNotExists,
		},
		{
			name: "cannot create session",
			inputUser: &models.UserAccount{
				Email:    "test@gmail.com",
				Password: "123456A!",
			},
			expectedToken: "valid_token",
			expectedUser:  nil,
			mockBehavior: func(r *mock_repository.MockUserRepository, email string) {
				expected := &models.UserAccount{
					Email:       "test@gmail.com",
					Password:    "$2a$10$AN0zyrlPjIZvO12mLf8PierYf579fxzfqh6j9aZn3LWetXMjJjYb.",
					CompanyName: "VK",
					UserType:    "employer",
				}
				r.EXPECT().GetUserByEmail(email).Return(expected, nil)
			},
			sessionRepBehavior: func(r *mock_session.MockRepository, email string) {
				r.EXPECT().NewSession(email).Return("", errorHandler.ErrBadRequest)
			},
			mailBehavior: func(r *mock_usecase.MockMail, email string) {},
			expectedErr:  errorHandler.ErrBadRequest,
		},
		{
			name: "cannot create session",
			inputUser: &models.UserAccount{
				Email:    "test@gmail.com",
				Password: "123456A!",
			},
			expectedToken: "valid_token",
			expectedUser:  nil,
			mockBehavior: func(r *mock_repository.MockUserRepository, email string) {
				expected := &models.UserAccount{
					Email:       "test@gmail.com",
					Password:    "$2a$10$AN0zyrlPjIZvO12mLf8PierYf579fxzfqh6j9aZn3LWetXMjJjYb.",
					CompanyName: "VK",
					UserType:    "asfafafa",
				}
				r.EXPECT().GetUserByEmail(email).Return(expected, nil)
			},
			sessionRepBehavior: func(r *mock_session.MockRepository, email string) {
				r.EXPECT().NewSession(email).Return("valid_token", nil)
			},
			mailBehavior: func(r *mock_usecase.MockMail, email string) {},
			expectedErr:  errorHandler.InvalidUserType,
		},
		{
			name: "invalid auth data",
			inputUser: &models.UserAccount{
				Email:    "test@gmail.com",
				Password: "123456",
			},
			expectedToken: "valid_token",
			expectedUser:  nil,
			mockBehavior: func(r *mock_repository.MockUserRepository, email string) {
			},
			sessionRepBehavior: func(r *mock_session.MockRepository, email string) {
			},
			mailBehavior: func(r *mock_usecase.MockMail, email string) {},
			expectedErr:  errorHandler.InvalidPasswordFormat,
		},
	}

	for _, test := range testTable {
		testCase := test
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()
			c := gomock.NewController(t)
			defer c.Finish()

			mockUserRepository := mock_repository.NewMockUserRepository(c)
			mockSessionRep := mock_session.NewMockRepository(c)
			mailRep := mock_usecase.NewMockMail(c)
			testCase.sessionRepBehavior(mockSessionRep, testCase.inputUser.Email)
			testCase.mockBehavior(mockUserRepository, testCase.inputUser.Email)
			testCase.mailBehavior(mailRep, testCase.inputUser.Email)
			cfg := &configs.Config{
				Validation: configs.ValidationConfig{
					MaxEmailLength:    30,
					MinSurnameLength:  2,
					MaxSurnameLength:  30,
					MinNameLength:     2,
					MaxNameLength:     30,
					MaxPasswordLength: 20,
					MinPasswordLength: 8,
					MinEmailLength:    8,
				}}
			if testCase.confirmMode {
				cfg.Security.ConfirmAccountMode = true
			}
			userService := UserService{
				userRep: mockUserRepository, sessionRepo: mockSessionRep, cfg: cfg, mail: mailRep,
			}
			_, err := userService.SignIn(testCase.inputUser)

			assert.Equal(t, testCase.expectedErr, err)
		})
	}
} //100%

func TestUserService_SignUp(t *testing.T) { //91%
	type mockBehavior func(r *mock_repository.MockUserRepository, email string, user *models.UserAccount)
	type sessionRepBehavior func(r *mock_session.MockRepository, email string)
	testTable := []struct {
		name               string
		inputUser          *models.UserAccount
		mockBehavior       mockBehavior
		sessionRepBehavior sessionRepBehavior
		expectedUser       *models.UserAccount
		expectedErr        error
		isUserValid        bool
	}{
		{
			name: "ok",
			inputUser: &models.UserAccount{
				Email:            "test@gmail.com",
				Password:         "123456A!",
				ApplicantSurname: "Kozirev",
				ApplicantName:    "Zakhar",
				UserType:         "applicant",
			},
			isUserValid:  true,
			expectedUser: nil,
			mockBehavior: func(r *mock_repository.MockUserRepository, email string, user *models.UserAccount) {
				r.EXPECT().GetUserByEmail(email).Return(nil, errorHandler.ErrUserNotExists)
				r.EXPECT().CreateUser(user).Return(nil)
			},
			sessionRepBehavior: func(r *mock_session.MockRepository, email string) {
				r.EXPECT().NewSession(email).Return("valid_token", nil)
			},
			expectedErr: nil,
		},
		{
			name: "cannot create user",
			inputUser: &models.UserAccount{
				Email:            "test@gmail.com",
				Password:         "123456A!",
				ApplicantSurname: "Kozirev",
				ApplicantName:    "Zakhar",
				UserType:         "applicant",
			},
			isUserValid:  true,
			expectedUser: nil,
			mockBehavior: func(r *mock_repository.MockUserRepository, email string, user *models.UserAccount) {
				r.EXPECT().GetUserByEmail(email).Return(nil, errorHandler.ErrUserNotExists)
				r.EXPECT().CreateUser(user).Return(errorHandler.ErrCannotCreateUser)
			},
			sessionRepBehavior: func(r *mock_session.MockRepository, email string) {
			},
			expectedErr: fmt.Errorf("creating session user: %w", errorHandler.ErrCannotCreateUser),
		},
		{
			name: "cannot create session",
			inputUser: &models.UserAccount{
				Email:            "test@gmail.com",
				Password:         "123456A!",
				ApplicantSurname: "Kozirev",
				ApplicantName:    "Zakhar",
				UserType:         "applicant",
			},
			isUserValid:  true,
			expectedUser: nil,
			mockBehavior: func(r *mock_repository.MockUserRepository, email string, user *models.UserAccount) {
				r.EXPECT().GetUserByEmail(email).Return(nil, errorHandler.ErrUserNotExists)
				r.EXPECT().CreateUser(user).Return(nil)
			},
			sessionRepBehavior: func(r *mock_session.MockRepository, email string) {
				r.EXPECT().NewSession(email).Return("", errorHandler.ErrBadRequest)
			},
			expectedErr: errorHandler.ErrBadRequest,
		},
		{
			name: "user already exist",
			inputUser: &models.UserAccount{
				Email:            "test@gmail.com",
				Password:         "123456A!",
				ApplicantSurname: "Kozirev",
				ApplicantName:    "Zakhar",
				UserType:         "applicant",
			},
			isUserValid:  true,
			expectedUser: nil,
			mockBehavior: func(r *mock_repository.MockUserRepository, email string, user *models.UserAccount) {
				r.EXPECT().GetUserByEmail(email).Return(&models.UserAccount{}, nil)
			},
			sessionRepBehavior: func(r *mock_session.MockRepository, email string) {
			},
			expectedErr: errorHandler.ErrUserExists,
		},
		{
			name: "invalid input 1",
			inputUser: &models.UserAccount{
				Email:            "test@gmail.com",
				Password:         "123456A!",
				ApplicantSurname: "Kozirev",
				ApplicantName:    "Zakhar",
				UserType:         "applicnt",
			},
			isUserValid:  false,
			expectedUser: nil,
			mockBehavior: func(r *mock_repository.MockUserRepository, email string, user *models.UserAccount) {
			},
			sessionRepBehavior: func(r *mock_session.MockRepository, email string) {
			},
			expectedErr: errorHandler.InvalidUserType,
		},
		{
			name: "invalid input 2",
			inputUser: &models.UserAccount{
				Email:            "testgmail.com",
				Password:         "123456A!",
				ApplicantSurname: "Kozirev",
				ApplicantName:    "Zakhar",
				UserType:         "applicant",
			},
			isUserValid:  false,
			expectedUser: nil,
			mockBehavior: func(r *mock_repository.MockUserRepository, email string, user *models.UserAccount) {
			},
			sessionRepBehavior: func(r *mock_session.MockRepository, email string) {
			},
			expectedErr: errorHandler.InvalidEmailFormat,
		},
	}

	for _, test := range testTable {
		testCase := test
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()
			c := gomock.NewController(t)
			defer c.Finish()

			mockUserRepository := mock_repository.NewMockUserRepository(c)
			mockSessionRep := mock_session.NewMockRepository(c)

			testCase.sessionRepBehavior(mockSessionRep, testCase.inputUser.Email)
			testCase.mockBehavior(mockUserRepository, testCase.inputUser.Email, testCase.inputUser)

			userService := UserService{userRep: mockUserRepository, sessionRepo: mockSessionRep, cfg: &configs.Config{
				Validation: configs.ValidationConfig{
					MaxEmailLength:    30,
					MinSurnameLength:  2,
					MaxSurnameLength:  30,
					MinNameLength:     2,
					MaxNameLength:     30,
					MaxPasswordLength: 20,
					MinPasswordLength: 8,
					MinEmailLength:    8,
				},
			}}
			_, err := userService.SignUp(testCase.inputUser)

			assert.Equal(t, testCase.expectedErr, err)
		})
	}
} //91%

func TestUserService_Logout(t *testing.T) {
	type sessionRepBehavior func(r *mock_session.MockRepository, email string)
	testTable := []struct {
		name               string
		token              string
		sessionRepBehavior sessionRepBehavior
		expectedErr        error
	}{
		{
			name:  "ok",
			token: "valid_token",
			sessionRepBehavior: func(r *mock_session.MockRepository, token string) {
				r.EXPECT().Delete(token).Return(nil)
			},
			expectedErr: nil,
		},
		{
			name:  "error",
			token: "valid_token",
			sessionRepBehavior: func(r *mock_session.MockRepository, token string) {
				r.EXPECT().Delete(token).Return(errorHandler.ErrCannotDeleteSession)
			},
			expectedErr: errorHandler.ErrCannotDeleteSession,
		},
	}
	for _, test := range testTable {
		testCase := test
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()
			c := gomock.NewController(t)
			defer c.Finish()
			mockSessionRep := mock_session.NewMockRepository(c)
			userService := UserService{sessionRepo: mockSessionRep, cfg: &configs.Config{
				Validation: configs.ValidationConfig{
					MaxEmailLength:    30,
					MinSurnameLength:  2,
					MaxSurnameLength:  30,
					MinNameLength:     2,
					MaxNameLength:     30,
					MaxPasswordLength: 20,
					MinPasswordLength: 8,
					MinEmailLength:    8,
				},
			}}
			testCase.sessionRepBehavior(mockSessionRep, testCase.token)
			err := userService.Logout(testCase.token)
			assert.Equal(t, testCase.expectedErr, err)
		})
	}
} //100%

func TestUserService_AuthCheck(t *testing.T) {
	type mockBehavior func(r *mock_repository.MockUserRepository, email string)
	testTable := []struct {
		name         string
		inputEmail   string
		mockBehavior mockBehavior
		expectedUser *models.UserAccount
		expectedErr  error
	}{
		{
			name:       "ok",
			inputEmail: "test@gmail.com",
			expectedUser: &models.UserAccount{
				Email: "test@gmail.com",
			},
			mockBehavior: func(r *mock_repository.MockUserRepository, email string) {
				expected := &models.UserAccount{
					Email: "test@gmail.com",
				}
				r.EXPECT().GetUserByEmail(email).Return(expected, nil)
			},
			expectedErr: nil,
		},
		{
			name:         "user not found",
			inputEmail:   "test@gmail.com",
			expectedUser: &models.UserAccount{},
			mockBehavior: func(r *mock_repository.MockUserRepository, email string) {
				r.EXPECT().GetUserByEmail(email).Return(&models.UserAccount{}, errorHandler.ErrUserNotExists)
			},
			expectedErr: errorHandler.ErrUserNotExists,
		},
	}
	for _, test := range testTable {
		testCase := test
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()
			c := gomock.NewController(t)
			defer c.Finish()

			mockUserRepository := mock_repository.NewMockUserRepository(c)
			testCase.mockBehavior(mockUserRepository, testCase.inputEmail)
			userService := UserService{userRep: mockUserRepository}
			user, err := userService.AuthCheck(testCase.inputEmail)
			if testCase.expectedErr == nil {
				assert.Equal(t, *testCase.expectedUser, *user)
			}
			assert.Equal(t, testCase.expectedErr, err)
		})
	}
} //100%

func TestUserService_GetUser(t *testing.T) { //100%
	type mockBehavior func(r *mock_repository.MockUserRepository, id uint)
	testTable := []struct {
		name         string
		id           uint
		mockBehavior mockBehavior
		expectedUser *models.UserAccount
		expectedErr  error
	}{
		{
			name: "ok",
			mockBehavior: func(r *mock_repository.MockUserRepository, id uint) {
				expected := &models.UserAccount{
					Email: "test@gmail.com",
				}
				r.EXPECT().GetUser(id).Return(expected, nil)
			},
			expectedUser: &models.UserAccount{
				Email: "test@gmail.com",
			},
			expectedErr: nil,
		}, {
			name: "error",
			mockBehavior: func(r *mock_repository.MockUserRepository, id uint) {
				expected := &models.UserAccount{
					Email: "test@gmail.com",
				}
				r.EXPECT().GetUser(id).Return(expected, errorHandler.ErrUserNotExists)
			},
			expectedUser: &models.UserAccount{
				Email: "test@gmail.com",
			},
			expectedErr: errorHandler.ErrUserNotExists,
		},
	}
	for _, test := range testTable {
		testCase := test
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()
			c := gomock.NewController(t)
			defer c.Finish()

			mockUserRepository := mock_repository.NewMockUserRepository(c)
			testCase.mockBehavior(mockUserRepository, testCase.id)
			userService := UserService{userRep: mockUserRepository}
			user, err := userService.GetUser(testCase.id)
			if testCase.expectedErr == nil {
				assert.Equal(t, *testCase.expectedUser, *user)
			}
			assert.Equal(t, testCase.expectedErr, err)
		})
	}
} //100%

func TestUserService_GetUserSafety(t *testing.T) { //100%
	type mockBehavior func(r *mock_repository.MockUserRepository, id uint, fields []string)
	testTable := []struct {
		name           string
		id             uint
		mockBehavior   mockBehavior
		expectedUser   *models.UserAccount
		expectedErr    error
		expectedFields []string
	}{
		{
			name: "ok",
			mockBehavior: func(r *mock_repository.MockUserRepository, id uint, fields []string) {
				expected := &models.UserAccount{
					Email:        "test@gmail.com",
					PublicFields: "email contact_number applicant_current_salary",
				}
				r.EXPECT().GetUser(id).Return(expected, nil)
				r.EXPECT().GetUserSafety(id, fields).Return(expected, nil)
			},
			expectedUser: &models.UserAccount{
				Email:        "test@gmail.com",
				PublicFields: "email contact_number applicant_current_salary",
			},
			expectedErr:    nil,
			expectedFields: models.PrivateUserFields,
		}, {
			name: "failed fields validation",
			mockBehavior: func(r *mock_repository.MockUserRepository, id uint, fields []string) {
				expected := &models.UserAccount{
					Email:        "test@gmail.com",
					PublicFields: "e e e e e e e e e e e e e e e ee e e e e e  s s s s s ssss s s s s ss s",
				}
				r.EXPECT().GetUser(id).Return(expected, nil)
			},
			expectedUser: &models.UserAccount{
				Email:        "test@gmail.com",
				PublicFields: "email contact_number applicant_current_salary",
			},
			expectedErr:    errorHandler.ErrBadRequest,
			expectedFields: models.PrivateUserFields,
		},
		{
			name: "cannot get user",
			mockBehavior: func(r *mock_repository.MockUserRepository, id uint, fields []string) {
				r.EXPECT().GetUser(id).Return(nil, errorHandler.ErrBadRequest)
			},
			expectedUser: &models.UserAccount{
				Email:        "test@gmail.com",
				PublicFields: "email contact_number applicant_current_salary",
			},
			expectedErr:    errorHandler.ErrBadRequest,
			expectedFields: models.PrivateUserFields,
		},
		{
			name: "cannot get user safety",
			mockBehavior: func(r *mock_repository.MockUserRepository, id uint, fields []string) {
				expected := &models.UserAccount{
					Email:        "test@gmail.com",
					PublicFields: "email contact_number applicant_current_salary",
				}
				r.EXPECT().GetUser(id).Return(expected, nil)
				r.EXPECT().GetUserSafety(id, fields).Return(nil, errorHandler.ErrBadRequest)
			},
			expectedUser: &models.UserAccount{
				Email:        "test@gmail.com",
				PublicFields: "email contact_number applicant_current_salary",
			},
			expectedErr:    errorHandler.ErrBadRequest,
			expectedFields: models.PrivateUserFields,
		},
		{
			name: "failed fields validation",
			mockBehavior: func(r *mock_repository.MockUserRepository, id uint, fields []string) {
				expected := &models.UserAccount{
					Email:        "test@gmail.com",
					PublicFields: "email contat_number applicant_current_salary",
				}
				r.EXPECT().GetUser(id).Return(expected, nil)
			},
			expectedUser: &models.UserAccount{
				Email:        "test@gmail.com",
				PublicFields: "email contact_number applicant_current_salary",
			},
			expectedErr:    errorHandler.ErrBadRequest,
			expectedFields: models.PrivateUserFields,
		},
		{
			name: "empty fields",
			mockBehavior: func(r *mock_repository.MockUserRepository, id uint, fields []string) {
				expected := &models.UserAccount{
					Email:        "test@gmail.com",
					PublicFields: "",
				}
				r.EXPECT().GetUser(id).Return(expected, nil)
				r.EXPECT().GetUserSafety(id, fields).Return(expected, nil)
			},
			expectedUser: &models.UserAccount{
				Email:        "test@gmail.com",
				PublicFields: "",
			},
			expectedErr:    nil,
			expectedFields: []string{},
		},
	}
	for _, test := range testTable {
		testCase := test
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()
			c := gomock.NewController(t)
			defer c.Finish()
			mockUserRepository := mock_repository.NewMockUserRepository(c)
			testCase.mockBehavior(mockUserRepository, testCase.id, testCase.expectedFields)
			userService := UserService{userRep: mockUserRepository}
			user, err := userService.GetUserSafety(testCase.id)
			if testCase.expectedErr == nil {
				assert.Equal(t, *testCase.expectedUser, *user)
			}
			assert.Equal(t, testCase.expectedErr, err)
		})
	}
} //100%

func TestUserService_UpdateUser(t *testing.T) { //93%
	type getMockBehavior func(r *mock_repository.MockUserRepository, email string)
	type updateMockBehavior func(r *mock_repository.MockUserRepository, oldUser, newUser *models.UserAccount)
	testTable := []struct {
		name               string
		inputUser          *models.UserAccount
		oldUser            *models.UserAccount
		getMockBehavior    getMockBehavior
		updateMockBehavior updateMockBehavior
		expectedErr        error
	}{
		{
			name: "ok",
			inputUser: &models.UserAccount{
				Email:            "test@gmail.com",
				ApplicantSurname: "Kozirev",
				ApplicantName:    "Zakhar",
				UserType:         "applicant",
			},
			oldUser: &models.UserAccount{
				Email:            "test@gmail.com",
				ApplicantSurname: "Urvancev",
				ApplicantName:    "Zakhar",
				UserType:         "applicant",
			},
			getMockBehavior: func(r *mock_repository.MockUserRepository, email string) {
				old := &models.UserAccount{
					Email:            "test@gmail.com",
					ApplicantSurname: "Urvancev",
					ApplicantName:    "Zakhar",
					UserType:         "applicant",
				}
				r.EXPECT().GetUserByEmail(email).Return(old, nil)
			},
			updateMockBehavior: func(r *mock_repository.MockUserRepository, oldUser, newUser *models.UserAccount) {
				r.EXPECT().UpdateUser(newUser).Return(nil)
			},
			expectedErr: nil,
		},
		{
			name: "ok with password",
			inputUser: &models.UserAccount{
				Email:            "test@gmail.com",
				ApplicantSurname: "Kozirev",
				ApplicantName:    "Zakhar",
				UserType:         "applicant",
				Password:         "123456z!",
			},
			oldUser: &models.UserAccount{
				Email:            "test@gmail.com",
				ApplicantSurname: "Urvancev",
				ApplicantName:    "Zakhar",
				UserType:         "applicant",
			},
			getMockBehavior: func(r *mock_repository.MockUserRepository, email string) {
				old := &models.UserAccount{
					Email:            "test@gmail.com",
					ApplicantSurname: "Urvancev",
					ApplicantName:    "Zakhar",
					UserType:         "applicant",
				}
				r.EXPECT().GetUserByEmail(email).Return(old, nil)
			},
			updateMockBehavior: func(r *mock_repository.MockUserRepository, oldUser, newUser *models.UserAccount) {
				r.EXPECT().UpdateUser(newUser).Return(nil)
			},
			expectedErr: nil,
		},
		{
			name: "user is not valid",
			inputUser: &models.UserAccount{
				Email:            "testgmail.com",
				ApplicantSurname: "K",
				ApplicantName:    "Zakhar",
				UserType:         "applicant",
			},
			oldUser: &models.UserAccount{
				Email:            "test@gmail.com",
				ApplicantSurname: "Urvancev",
				ApplicantName:    "Zakhar",
				UserType:         "applicant",
			},
			getMockBehavior: func(r *mock_repository.MockUserRepository, email string) {
			},
			updateMockBehavior: func(r *mock_repository.MockUserRepository, oldUser, newUser *models.UserAccount) {
			},
			expectedErr: errorHandler.IncorrectSurnameLength,
		},
		{
			name: "user not exists",
			inputUser: &models.UserAccount{
				Email:            "test@gmail.com",
				ApplicantSurname: "Kozirev",
				ApplicantName:    "Zakhar",
				UserType:         "applicant",
			},
			oldUser: nil,
			getMockBehavior: func(r *mock_repository.MockUserRepository, email string) {
				r.EXPECT().GetUserByEmail(email).Return(nil, errorHandler.ErrUserNotExists)
			},
			updateMockBehavior: func(r *mock_repository.MockUserRepository, oldUser, newUser *models.UserAccount) {
			},
			expectedErr: errorHandler.ErrUserNotExists,
		},
		{
			name: "invalid user type",
			inputUser: &models.UserAccount{
				Email:            "test@gmail.com",
				ApplicantSurname: "Kozirev",
				ApplicantName:    "Zakhar",
				UserType:         "applicant",
			},
			oldUser: &models.UserAccount{
				Email:            "test@gmail.com",
				ApplicantSurname: "Urvancev",
				ApplicantName:    "Zakhar",
				UserType:         "employer",
			},
			getMockBehavior: func(r *mock_repository.MockUserRepository, email string) {
				old := &models.UserAccount{
					Email:            "test@gmail.com",
					ApplicantSurname: "Urvancev",
					ApplicantName:    "Zakhar",
					UserType:         "employer",
				}
				r.EXPECT().GetUserByEmail(email).Return(old, nil)
			},
			updateMockBehavior: func(r *mock_repository.MockUserRepository, oldUser, newUser *models.UserAccount) {
			},
			expectedErr: errorHandler.ErrBadRequest,
		},
		{
			name: "error with rep",
			inputUser: &models.UserAccount{
				Email:            "test@gmail.com",
				ApplicantSurname: "Kozirev",
				ApplicantName:    "Zakhar",
				UserType:         "applicant",
			},
			oldUser: &models.UserAccount{
				Email:            "test@gmail.com",
				ApplicantSurname: "Urvancev",
				ApplicantName:    "Zakhar",
				UserType:         "applicant",
			},
			getMockBehavior: func(r *mock_repository.MockUserRepository, email string) {
				old := &models.UserAccount{
					Email:            "test@gmail.com",
					ApplicantSurname: "Urvancev",
					ApplicantName:    "Zakhar",
					UserType:         "applicant",
				}
				r.EXPECT().GetUserByEmail(email).Return(old, nil)
			},
			updateMockBehavior: func(r *mock_repository.MockUserRepository, oldUser, newUser *models.UserAccount) {
				r.EXPECT().UpdateUser(newUser).Return(errorHandler.ErrBadRequest)
			},
			expectedErr: errorHandler.ErrBadRequest,
		},
	}

	for _, test := range testTable {
		testCase := test
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()
			c := gomock.NewController(t)
			defer c.Finish()

			mockUserRepository := mock_repository.NewMockUserRepository(c)
			mockSessionRep := mock_session.NewMockRepository(c)

			testCase.getMockBehavior(mockUserRepository, testCase.inputUser.Email)
			testCase.updateMockBehavior(mockUserRepository, testCase.oldUser, testCase.inputUser)

			userService := UserService{userRep: mockUserRepository, sessionRepo: mockSessionRep, cfg: &configs.Config{
				Validation: configs.ValidationConfig{
					MaxEmailLength:    30,
					MinSurnameLength:  2,
					MaxSurnameLength:  30,
					MinNameLength:     2,
					MaxNameLength:     30,
					MaxPasswordLength: 20,
					MinPasswordLength: 8,
					MinEmailLength:    8,
				},
			}}

			err := userService.UpdateUser(testCase.inputUser)
			assert.Equal(t, testCase.expectedErr, err)
		})
	}
} //93%

func TestUserService_GetUserId(t *testing.T) {
	type mockBehavior func(r *mock_repository.MockUserRepository, email string)

	testTable := []struct {
		name         string
		email        string
		mockBehavior mockBehavior
		expected     uint
		expectedErr  error
	}{
		{
			name:     "ok",
			email:    "test@gmail.com",
			expected: 1,
			mockBehavior: func(r *mock_repository.MockUserRepository, email string) {
				expected := &models.UserAccount{
					ID:          1,
					Email:       "test@gmail.com",
					Password:    "123456A!",
					CompanyName: "VK",
					UserType:    "employer",
				}
				r.EXPECT().GetUserByEmail(email).Return(expected, nil)
			},
			expectedErr: nil,
		},
		{
			name:     "error",
			email:    "test@gmail.com",
			expected: 0,
			mockBehavior: func(r *mock_repository.MockUserRepository, email string) {
				expected := &models.UserAccount{}
				r.EXPECT().GetUserByEmail(email).Return(expected, errorHandler.ErrUserNotExists)
			},
			expectedErr: errorHandler.ErrUserNotExists,
		},
	}

	for _, test := range testTable {
		testCase := test
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()
			c := gomock.NewController(t)
			defer c.Finish()

			mockUserRepository := mock_repository.NewMockUserRepository(c)
			testCase.mockBehavior(mockUserRepository, test.email)
			userService := UserService{userRep: mockUserRepository}
			resultId, err := userService.GetUserId(testCase.email)

			assert.Equal(t, testCase.expected, resultId)
			assert.Equal(t, testCase.expectedErr, err)
		})
	}
} //100%

func TestUserService_GetAllEmployers(t *testing.T) {
	type mockBehavior func(r *mock_repository.MockUserRepository, conditions interface{}, filterValues interface{}, filters models.UserFilter)
	testTable := []struct {
		name          string
		conditions    interface{}
		filterValues  interface{}
		mockBehavior  mockBehavior
		expectedUsers []*models.UserAccount
		filters       models.UserFilter
		expectedErr   error
	}{
		{
			name: "ok",
			mockBehavior: func(r *mock_repository.MockUserRepository, conditions interface{}, filterValues interface{}, filters models.UserFilter) {
				expected := []*models.UserAccount{
					{
						CompanyName: "Job",
					},
				}
				r.EXPECT().GetAllUsers(conditions, filterValues, "employer").Return(expected, nil)
			},
			expectedUsers: []*models.UserAccount{
				{
					CompanyName: "Job",
				},
			},
			expectedErr: nil,
		},
		{
			name: "cannot get employers",
			mockBehavior: func(r *mock_repository.MockUserRepository, conditions interface{}, filterValues interface{}, filters models.UserFilter) {
				r.EXPECT().GetAllUsers(conditions, filterValues, "employer").Return([]*models.UserAccount{}, errorHandler.ErrBadRequest)
			},
			expectedUsers: []*models.UserAccount{},
			expectedErr:   errorHandler.ErrBadRequest,
		},
	}
	for _, test := range testTable {
		testCase := test
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()
			c := gomock.NewController(t)
			defer c.Finish()

			userRep := mock_repository.NewMockUserRepository(c)
			testCase.mockBehavior(userRep, testCase.conditions, testCase.filterValues, testCase.filters)
			userService := UserService{userRep: userRep}
			user, err := userService.GetAllEmployers(testCase.filters)
			if testCase.expectedErr == nil {
				assert.Equal(t, testCase.expectedUsers, user)
			}
			assert.Equal(t, testCase.expectedErr, err)
		})

	}
}

func TestUserService_GetAllApplicants(t *testing.T) {
	type mockBehavior func(r *mock_repository.MockUserRepository, conditions interface{}, filterValues interface{}, filters models.UserFilter)
	testTable := []struct {
		name          string
		conditions    interface{}
		filterValues  interface{}
		mockBehavior  mockBehavior
		expectedUsers []*models.UserAccount
		filters       models.UserFilter
		expectedErr   error
	}{
		{
			name: "ok",
			mockBehavior: func(r *mock_repository.MockUserRepository, conditions interface{}, filterValues interface{}, filters models.UserFilter) {
				expected := []*models.UserAccount{
					{
						CompanyName: "Job",
					},
				}
				r.EXPECT().GetAllUsers(conditions, filterValues, "applicant").Return(expected, nil)
			},
			expectedUsers: []*models.UserAccount{
				{
					CompanyName: "Job",
				},
			},
			expectedErr: nil,
		},
		{
			name: "cannot get applicants",
			mockBehavior: func(r *mock_repository.MockUserRepository, conditions interface{}, filterValues interface{}, filters models.UserFilter) {
				r.EXPECT().GetAllUsers(conditions, filterValues, "applicant").Return([]*models.UserAccount{}, errorHandler.ErrBadRequest)
			},
			expectedUsers: []*models.UserAccount{},
			expectedErr:   errorHandler.ErrBadRequest,
		},
	}
	for _, test := range testTable {
		testCase := test
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()
			c := gomock.NewController(t)
			defer c.Finish()

			userRep := mock_repository.NewMockUserRepository(c)
			testCase.mockBehavior(userRep, testCase.conditions, testCase.filterValues, testCase.filters)
			userService := UserService{userRep: userRep}
			user, err := userService.GetAllApplicants(testCase.filters)
			if testCase.expectedErr == nil {
				assert.Equal(t, testCase.expectedUsers, user)
			}
			assert.Equal(t, testCase.expectedErr, err)
		})

	}
}
