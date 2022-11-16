package impl

import (
	"HeadHunter/configs"
	"HeadHunter/internal/entity/models"
	"HeadHunter/internal/errorHandler"
	mock_repository "HeadHunter/internal/repository/mocks"
	mock_session "HeadHunter/internal/repository/session/mocks"
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
}

func TestUserService_SignIn(t *testing.T) {
	type mockBehavior func(r *mock_repository.MockUserRepository, email string)
	type sessionRepBehavior func(r *mock_session.MockRepository, email string)
	testTable := []struct {
		name               string
		inputUser          *models.UserAccount
		mockBehavior       mockBehavior
		sessionRepBehavior sessionRepBehavior
		expectedToken      string
		expectedUser       *models.UserAccount
		expectedErr        error
	}{
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
			sessionRepBehavior: func(r *mock_session.MockRepository, email string) {
				r.EXPECT().NewSession(email).Return("valid_token", nil)
			},
			expectedErr: errorHandler.ErrWrongPassword,
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
			sessionRepBehavior: func(r *mock_session.MockRepository, email string) {
				r.EXPECT().NewSession(email).Return("valid_token", nil)
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
			mockSessionRep := mock_session.NewMockRepository(c)
			if testCase.expectedUser != nil {
				testCase.sessionRepBehavior(mockSessionRep, testCase.inputUser.Email)
			}
			testCase.mockBehavior(mockUserRepository, testCase.inputUser.Email)
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
			_, err := userService.SignIn(testCase.inputUser)

			assert.Equal(t, testCase.expectedErr, err)
		})
	}
}

func TestUserService_SignUp(t *testing.T) {
	type mockBehavior func(r *mock_repository.MockUserRepository, email string)
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
			mockBehavior: func(r *mock_repository.MockUserRepository, email string) {
				r.EXPECT().IsUserExist(email).Return(true, nil)
			},
			sessionRepBehavior: func(r *mock_session.MockRepository, email string) {
				r.EXPECT().NewSession(email).Return("valid_token", nil)
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
			mockBehavior: func(r *mock_repository.MockUserRepository, email string) {
				r.EXPECT().IsUserExist(email).Return(true, nil)
			},
			sessionRepBehavior: func(r *mock_session.MockRepository, email string) {
				r.EXPECT().NewSession(email).Return("valid_token", nil)
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
			mockBehavior: func(r *mock_repository.MockUserRepository, email string) {
				r.EXPECT().IsUserExist(email).Return(true, nil)
			},
			sessionRepBehavior: func(r *mock_session.MockRepository, email string) {
				r.EXPECT().NewSession(email).Return("valid_token", nil)
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
			if testCase.isUserValid {
				if testCase.expectedUser != nil {
					testCase.sessionRepBehavior(mockSessionRep, testCase.inputUser.Email)
				}
				testCase.mockBehavior(mockUserRepository, testCase.inputUser.Email)
			}
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
}
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
				r.EXPECT().DeleteSession(token).Return(nil)
			},
			expectedErr: nil,
		},
		{
			name:  "error",
			token: "valid_token",
			sessionRepBehavior: func(r *mock_session.MockRepository, token string) {
				r.EXPECT().DeleteSession(token).Return(errorHandler.ErrCannotDeleteSession)
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
}

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
}

func TestUserService_GetUser(t *testing.T) {
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
}

func TestUserService_GetUserSafety(t *testing.T) {
	type mockBehavior func(r *mock_repository.MockUserRepository, id uint, slice []string)
	testTable := []struct {
		name         string
		id           uint
		mockBehavior mockBehavior
		expectedUser *models.UserAccount
		expectedErr  error
	}{
		{
			name: "ok",
			mockBehavior: func(r *mock_repository.MockUserRepository, id uint, slice []string) {
				expected := &models.UserAccount{
					Email: "test@gmail.com",
				}
				r.EXPECT().GetUserSafety(id, slice).Return(expected, nil)
			},
			expectedUser: &models.UserAccount{
				Email: "test@gmail.com",
			},
			expectedErr: nil,
		}, {
			name: "error",
			mockBehavior: func(r *mock_repository.MockUserRepository, id uint, slice []string) {
				expected := &models.UserAccount{
					Email: "test@gmail.com",
				}
				r.EXPECT().GetUserSafety(id, slice).Return(expected, errorHandler.ErrUserNotExists)
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
			testCase.mockBehavior(mockUserRepository, testCase.id, models.PrivateUserFields)
			userService := UserService{userRep: mockUserRepository}
			user, err := userService.GetUserSafety(testCase.id)
			if testCase.expectedErr == nil {
				assert.Equal(t, *testCase.expectedUser, *user)
			}
			assert.Equal(t, testCase.expectedErr, err)
		})
	}
}

func TestUserService_UpdateUser(t *testing.T) {
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
				r.EXPECT().UpdateUser(oldUser, newUser).Return(nil)
			},
			expectedErr: nil,
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
				r.EXPECT().UpdateUser(oldUser, newUser).Return(nil)
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
				r.EXPECT().UpdateUser(oldUser, newUser).Return(nil)
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
				r.EXPECT().UpdateUser(oldUser, newUser).Return(errorHandler.ErrBadRequest)
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
			if testCase.oldUser != nil && testCase.oldUser.UserType == testCase.inputUser.UserType {
				testCase.updateMockBehavior(mockUserRepository, testCase.oldUser, testCase.inputUser)
			}
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
}
