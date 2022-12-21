package impl

import (
	"HeadHunter/common/session/mocks"
	"HeadHunter/configs"
	"HeadHunter/internal/entity/models"
	"HeadHunter/internal/network/middleware"
	mock_usecases "HeadHunter/internal/usecases/mocks"
	errorHandler2 "HeadHunter/pkg/errorHandler"
	"bytes"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestUserHandler_SignUp(t *testing.T) {
	type mockBehavior func(r *mock_usecases.MockUser, user *models.UserAccount)

	testTable := []struct {
		name                 string
		inputUser            models.UserAccount
		inputBody            string
		mockBehavior         mockBehavior
		expectedStatusCode   int
		expectedResponseBody string
	}{
		{
			name: "valid applicant",
			inputUser: models.UserAccount{
				Email:            "test@gmail.com",
				Password:         "123456a!",
				ApplicantSurname: "Urvancev",
				ApplicantName:    "Zakhar",
				UserType:         "applicant",
			},
			inputBody: `{
    			"email": "test@gmail.com",
    			"password": "123456a!",
                "applicant_name": "Zakhar",
                "applicant_surname": "Urvancev",
    			"user_type": "applicant"
}`,
			mockBehavior: func(r *mock_usecases.MockUser, user *models.UserAccount) {
				r.EXPECT().SignUp(user).Return(gomock.Any().String(), nil)
			},
			expectedStatusCode:   200,
			expectedResponseBody: "{\"id\":0,\"user_type\":\"applicant\",\"email\":\"test@gmail.com\",\"password\":\"\",\"contact_number\":\"\",\"status\":\"\",\"description\":\"\",\"image\":\"\",\"average_color\":\"\",\"date_of_birth\":\"0001-01-01T00:00:00Z\",\"created_time\":\"0001-01-01T00:00:00Z\",\"applicant_name\":\"Zakhar\",\"applicant_surname\":\"Urvancev\",\"company_size\":0,\"public_fields\":\"\",\"is_confirmed\":false,\"two_factor_sign_in\":false,\"mailing_approval\":false,\"resumes\":null,\"vacancies\":null,\"favourite_vacancies\":null,\"vacancy_activities\":null}",
		},
		{
			name: "confirmation mode",
			inputUser: models.UserAccount{
				Email:            "test@gmail.com",
				Password:         "123456a!",
				ApplicantSurname: "Urvancev",
				ApplicantName:    "Zakhar",
				UserType:         "applicant",
			},
			inputBody: `{
    			"email": "test@gmail.com",
    			"password": "123456a!",
                "applicant_name": "Zakhar",
                "applicant_surname": "Urvancev",
    			"user_type": "applicant"
}`,
			mockBehavior: func(r *mock_usecases.MockUser, user *models.UserAccount) {
				r.EXPECT().SignUp(user).Return(gomock.Any().String(), nil)
			},
			expectedStatusCode:   200,
			expectedResponseBody: "",
		},
		{
			name: "valid employer",
			inputUser: models.UserAccount{
				Email:       "test@gmail.com",
				Password:    "123456a!",
				CompanyName: "Some company",
				UserType:    "employer",
			},
			inputBody: `{
    			"email": "test@gmail.com",
    			"password": "123456a!",
                "company_name": "Some company",
    			"user_type": "employer"
}`,
			mockBehavior: func(r *mock_usecases.MockUser, user *models.UserAccount) {
				r.EXPECT().SignUp(user).Return(gomock.Any().String(), nil)
			},
			expectedStatusCode:   200,
			expectedResponseBody: "{\"id\":0,\"user_type\":\"employer\",\"email\":\"test@gmail.com\",\"password\":\"\",\"contact_number\":\"\",\"status\":\"\",\"description\":\"\",\"image\":\"\",\"average_color\":\"\",\"date_of_birth\":\"0001-01-01T00:00:00Z\",\"created_time\":\"0001-01-01T00:00:00Z\",\"company_name\":\"Some company\",\"company_size\":0,\"public_fields\":\"\",\"is_confirmed\":false,\"two_factor_sign_in\":false,\"mailing_approval\":false,\"resumes\":null,\"vacancies\":null,\"favourite_vacancies\":null,\"vacancy_activities\":null}",
		},
		{
			name: "user already exists",
			inputUser: models.UserAccount{
				Email:       "test@gmail.com",
				Password:    "123456a!",
				CompanyName: "Some company",
				UserType:    "employer",
			},
			inputBody: `{
    			"email": "test@gmail.com",
    			"password": "123456a!",
                "company_name": "Some company",
    			"user_type": "employer"
}`,
			mockBehavior: func(r *mock_usecases.MockUser, user *models.UserAccount) {
				r.EXPECT().SignUp(user).Return("", errorHandler2.ErrUserExists)
			},
			expectedStatusCode:   200,
			expectedResponseBody: "",
		},
	}
	for _, test := range testTable {
		testCase := test
		t.Run(testCase.name, func(t *testing.T) {
			t.Parallel()
			c := gomock.NewController(t)
			defer c.Finish()

			mockUseCase := mock_usecases.NewMockUser(c)
			testCase.mockBehavior(mockUseCase, &testCase.inputUser)
			_cfg := &configs.Config{
				DefaultExpiringSession: 100,
				Cookie: configs.CookieConfig{
					Secure:   false,
					HTTPOnly: true,
				},
			}
			if testCase.name == "confirmation mode" {
				_cfg.Security.ConfirmAccountMode = true
			}
			handler := UserHandler{
				userUseCase: mockUseCase,
				cfg:         _cfg,
			}

			r := gin.New()
			r.POST("/sign-up", handler.SignUp)

			w := httptest.NewRecorder()
			req := httptest.NewRequest("POST", "/sign-up",
				bytes.NewBufferString(testCase.inputBody))

			r.ServeHTTP(w, req)

			assert.Equal(t, testCase.expectedStatusCode, w.Code)
			assert.Equal(t, testCase.expectedResponseBody, w.Body.String())
		})
	}
}

func TestUserHandler_SignIn(t *testing.T) {
	type mockBehavior func(r *mock_usecases.MockUser, user *models.UserAccount)

	testTable := []struct {
		name                 string
		inputUser            models.UserAccount
		inputBody            string
		mockBehavior         mockBehavior
		expectedStatusCode   int
		expectedResponseBody string
	}{
		{
			name: "valid applicant",
			inputUser: models.UserAccount{
				Email:    "test@gmail.com",
				Password: "123456a!",
				UserType: "applicant",
			},
			inputBody: `{
   			"email": "test@gmail.com",
   			"password": "123456a!",
				"user_type": "applicant"
}`,
			mockBehavior: func(r *mock_usecases.MockUser, user *models.UserAccount) {
				r.EXPECT().SignIn(user).Return(gomock.Any().String(), nil)
			},
			expectedStatusCode:   200,
			expectedResponseBody: "{\"id\":0,\"user_type\":\"applicant\",\"email\":\"test@gmail.com\",\"password\":\"\",\"contact_number\":\"\",\"status\":\"\",\"description\":\"\",\"image\":\"\",\"average_color\":\"\",\"date_of_birth\":\"0001-01-01T00:00:00Z\",\"created_time\":\"0001-01-01T00:00:00Z\",\"company_size\":0,\"public_fields\":\"\",\"is_confirmed\":false,\"two_factor_sign_in\":false,\"mailing_approval\":false,\"resumes\":null,\"vacancies\":null,\"favourite_vacancies\":null,\"vacancy_activities\":null}",
		},
		{
			name: "valid applicant(two factor)",
			inputUser: models.UserAccount{
				Email:           "test@gmail.com",
				Password:        "123456a!",
				UserType:        "applicant",
				TwoFactorSignIn: true,
			},
			inputBody: `{
    			"email": "test@gmail.com",
    			"password": "123456a!",
				"user_type": "applicant",
				"two_factor_sign_in": true
}`,
			mockBehavior: func(r *mock_usecases.MockUser, user *models.UserAccount) {
				r.EXPECT().SignIn(user).Return(gomock.Any().String(), nil)
			},
			expectedStatusCode:   202,
			expectedResponseBody: "",
		},
		{
			name: "invalid body",
			inputUser: models.UserAccount{
				Email:            "test@gmail.com",
				Password:         "123456a!",
				ApplicantSurname: "Urvancev",
				ApplicantName:    "Zakhar",
				UserType:         "applicant",
			},
			inputBody: `{
   			"email": "test@gmail.com",
   	ar",
               "applicant_surname": "Urvancev",
   			"user_type": "applicant"
}`,
			mockBehavior: func(r *mock_usecases.MockUser, user *models.UserAccount) {
			},
			expectedStatusCode:   400,
			expectedResponseBody: "",
		},
		{
			name: "invalid user type",
			inputUser: models.UserAccount{
				Email:    "test@gmail.com",
				Password: "123456a!",
			},
			inputBody: `{
   			"email": "test@gmail.com",
   			"password": "123456a!"
}`,
			mockBehavior: func(r *mock_usecases.MockUser, user *models.UserAccount) {
				r.EXPECT().SignIn(user).Return(gomock.Any().String(), errorHandler2.InvalidUserType)
			},
			expectedStatusCode:   400,
			expectedResponseBody: "{\"descriptors\":\"\",\"error\":\"Некорректный входной тип пользователя\"}",
		},
		{
			name: "user not exist",
			inputUser: models.UserAccount{
				Email:    "test@gmail.com",
				Password: "123456a!",
			},
			inputBody: `{
   			"email": "test@gmail.com",
   			"password": "123456a!"
}`,
			mockBehavior: func(r *mock_usecases.MockUser, user *models.UserAccount) {
				r.EXPECT().SignIn(user).Return("", errorHandler2.ErrUserNotExists)
			},
			expectedStatusCode:   401,
			expectedResponseBody: "{\"descriptors\":\"email\",\"error\":\"Пользователя с таким email не существует\"}",
		},
	}
	for _, test := range testTable {
		testCase := test
		t.Run(testCase.name, func(t *testing.T) {
			t.Parallel()
			c := gomock.NewController(t)
			defer c.Finish()

			mockUseCase := mock_usecases.NewMockUser(c)
			testCase.mockBehavior(mockUseCase, &testCase.inputUser)

			handler := UserHandler{
				userUseCase: mockUseCase,
				cfg: &configs.Config{
					DefaultExpiringSession: 100,
					Cookie: configs.CookieConfig{
						Secure:   false,
						HTTPOnly: true,
					},
				},
			}

			r := gin.New()
			r.POST("/sign-in", handler.SignIn, errorHandler2.Middleware())

			w := httptest.NewRecorder()
			req := httptest.NewRequest("POST", "/sign-in",
				bytes.NewBufferString(testCase.inputBody))

			r.ServeHTTP(w, req)

			assert.Equal(t, testCase.expectedStatusCode, w.Code)
			assert.Equal(t, testCase.expectedResponseBody, w.Body.String())
		})
	}
}

func TestUserHandler_Logout(t *testing.T) {
	type mockBehavior func(r *mock_usecases.MockUser, token string)

	testTable := []struct {
		name                 string
		mockBehavior         mockBehavior
		inputToken           string
		expectedStatusCode   int
		expectedResponseBody string
	}{
		{
			name:       "no cookie",
			inputToken: "invalid_token",
			mockBehavior: func(r *mock_usecases.MockUser, token string) {
				r.EXPECT().Logout(token).Return(errorHandler2.ErrBadRequest)
			},
			expectedStatusCode:   400,
			expectedResponseBody: "{\"descriptors\":\"\",\"error\":\"Некорректный запрос\"}",
		},
		{
			name:       "with cookie",
			inputToken: "valid_token",
			mockBehavior: func(r *mock_usecases.MockUser, token string) {
				r.EXPECT().Logout(token).Return(errorHandler2.ErrBadRequest)
			},
			expectedStatusCode:   200,
			expectedResponseBody: "",
		},
	}
	for _, test := range testTable {
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()
			c := gomock.NewController(t)
			defer c.Finish()

			mockUseCase := mock_usecases.NewMockUser(c)
			test.mockBehavior(mockUseCase, test.inputToken)

			handler := UserHandler{
				userUseCase: mockUseCase,
				cfg: &configs.Config{
					DefaultExpiringSession: 100,
					Cookie: configs.CookieConfig{
						Secure:   false,
						HTTPOnly: true,
					},
				},
			}

			r := gin.New()
			r.POST("/logout", handler.Logout)
			w := httptest.NewRecorder()

			cookie := &http.Cookie{
				Name:  "session",
				Value: "valid_token",
			}
			req := httptest.NewRequest("POST", "/logout",
				bytes.NewBufferString(""))

			req.AddCookie(cookie)
			r.ServeHTTP(w, req)

			assert.Equal(t, test.expectedStatusCode, w.Code)
			assert.Equal(t, test.expectedResponseBody, w.Body.String())
		})
	}
}

func TestUserHandler_AuthCheck(t *testing.T) {
	type mockBehavior func(r *mock_usecases.MockUser, email string)
	type sessionRepBehavior func(r *mock_session.MockRepository, token string)
	testTable := []struct {
		name                 string
		mockBehavior         mockBehavior
		sessionRepBehavior   sessionRepBehavior
		inputToken           string
		emailFromToken       string
		expectedStatusCode   int
		expectedResponseBody string
	}{
		{
			name:           "have cookie",
			inputToken:     "valid_token",
			emailFromToken: "test@gmail.com",
			mockBehavior: func(r *mock_usecases.MockUser, email string) {
				expectedUser := &models.UserAccount{
					Email:       "test@gmail.com",
					CompanyName: "Some company",
					UserType:    "employer",
					Image:       "basic_applicant_avatar.webp",
				}
				r.EXPECT().AuthCheck(email).Return(expectedUser, nil)
			},
			sessionRepBehavior: func(sessionRep *mock_session.MockRepository, token string) {
				sessionRep.EXPECT().GetSession(token).Return("test@gmail.com", nil)
			},
			expectedStatusCode:   200,
			expectedResponseBody: "{\"id\":0,\"user_type\":\"employer\",\"email\":\"test@gmail.com\",\"password\":\"\",\"contact_number\":\"\",\"status\":\"\",\"description\":\"\",\"image\":\"basic_applicant_avatar.webp\",\"average_color\":\"\",\"date_of_birth\":\"0001-01-01T00:00:00Z\",\"created_time\":\"0001-01-01T00:00:00Z\",\"company_name\":\"Some company\",\"company_size\":0,\"public_fields\":\"\",\"is_confirmed\":false,\"two_factor_sign_in\":false,\"mailing_approval\":false,\"resumes\":null,\"vacancies\":null,\"favourite_vacancies\":null,\"vacancy_activities\":null}",
		},
		{
			name:           "invalid cookie",
			inputToken:     "invalid_token",
			emailFromToken: "",
			mockBehavior: func(r *mock_usecases.MockUser, email string) {
				expectedUser := &models.UserAccount{
					Email:       "test@gmail.com",
					CompanyName: "Some company",
					UserType:    "employer",
					Image:       "basic_applicant_avatar.webp",
				}
				r.EXPECT().AuthCheck(email).Return(expectedUser, nil)
			},
			sessionRepBehavior: func(sessionRep *mock_session.MockRepository, token string) {
				sessionRep.EXPECT().GetSession(token).Return("", fmt.Errorf("getting session error:"))
			},
			expectedStatusCode:   401,
			expectedResponseBody: "{\"descriptors\":\"\",\"error\":\"Клиент не авторизован\"}",
		},
	}
	for _, test := range testTable {
		testCase := test
		t.Run(testCase.name, func(t *testing.T) {
			t.Parallel()
			c := gomock.NewController(t)
			defer c.Finish()

			mockUseCase := mock_usecases.NewMockUser(c)
			sessionRep := mock_session.NewMockRepository(c)
			sessionMiddlware := middleware.NewSessionMiddleware(sessionRep)
			handler := UserHandler{
				userUseCase: mockUseCase,
				cfg: &configs.Config{
					DefaultExpiringSession: 100,
					Cookie: configs.CookieConfig{
						Secure:   false,
						HTTPOnly: true,
					},
				},
			}
			if testCase.emailFromToken != "" {
				testCase.mockBehavior(mockUseCase, testCase.emailFromToken)
			}
			testCase.sessionRepBehavior(sessionRep, testCase.inputToken)
			r := gin.New()
			r.GET("/auth", sessionMiddlware.Session, handler.AuthCheck, errorHandler2.Middleware())
			w := httptest.NewRecorder()

			cookie := &http.Cookie{
				Name:  "session",
				Value: testCase.inputToken,
			}
			req := httptest.NewRequest("GET", "/auth",
				bytes.NewBufferString(""))

			req.AddCookie(cookie)
			r.ServeHTTP(w, req)

			assert.Equal(t, testCase.expectedStatusCode, w.Code)
			assert.Equal(t, testCase.expectedResponseBody, w.Body.String())
		})
	}
}

func TestUserHandler_GetUser(t *testing.T) {
	type mockBehavior func(r *mock_usecases.MockUser, id uint)
	type sessionRepBehavior func(r *mock_session.MockRepository, token string)

	testTable := []struct {
		name                 string
		inputId              uint
		inputToken           string
		requestParam         string
		emailFromToken       string
		mockBehavior         mockBehavior
		sessionRepBehavior   sessionRepBehavior
		expectedStatusCode   int
		expectedResponseBody string
	}{
		{
			name:           "valid",
			inputId:        42,
			requestParam:   "42",
			emailFromToken: "test@gmail.com",
			mockBehavior: func(r *mock_usecases.MockUser, id uint) {
				expectedUser := &models.UserAccount{
					ID:          42,
					Email:       "test@gmail.com",
					CompanyName: "Some company",
					UserType:    "employer",
					Image:       "basic_applicant_avatar.webp",
				}
				r.EXPECT().GetUser(id).Return(expectedUser, nil)
			},
			sessionRepBehavior: func(sessionRep *mock_session.MockRepository, token string) {
				sessionRep.EXPECT().GetSession(token).Return("test@gmail.com", nil)
			},
			expectedStatusCode:   200,
			expectedResponseBody: "{\"id\":42,\"user_type\":\"employer\",\"email\":\"test@gmail.com\",\"password\":\"\",\"contact_number\":\"\",\"status\":\"\",\"description\":\"\",\"image\":\"basic_applicant_avatar.webp\",\"average_color\":\"\",\"date_of_birth\":\"0001-01-01T00:00:00Z\",\"created_time\":\"0001-01-01T00:00:00Z\",\"company_name\":\"Some company\",\"company_size\":0,\"public_fields\":\"\",\"is_confirmed\":false,\"two_factor_sign_in\":false,\"mailing_approval\":false,\"resumes\":null,\"vacancies\":null,\"favourite_vacancies\":null,\"vacancy_activities\":null}",
		},
		{
			name:           "not found",
			inputId:        42,
			requestParam:   "42",
			emailFromToken: "test@gmail.com",
			mockBehavior: func(r *mock_usecases.MockUser, id uint) {
				expectedUser := &models.UserAccount{
					ID:          42,
					Email:       "test@gmail.com",
					CompanyName: "Some company",
					UserType:    "employer",
					Image:       "basic_applicant_avatar.webp",
				}
				r.EXPECT().GetUser(id).Return(expectedUser, errorHandler2.ErrUserNotExists)
			},
			sessionRepBehavior: func(sessionRep *mock_session.MockRepository, token string) {
				sessionRep.EXPECT().GetSession(token).Return("test@gmail.com", nil)
			},
			expectedStatusCode:   401,
			expectedResponseBody: "{\"descriptors\":\"email\",\"error\":\"Пользователя с таким email не существует\"}",
		},
		{
			name:           "not found",
			inputId:        42,
			requestParam:   "42",
			emailFromToken: "",
			mockBehavior: func(r *mock_usecases.MockUser, id uint) {
				expectedUser := &models.UserAccount{
					ID:          42,
					Email:       "test@gmail.com",
					CompanyName: "Some company",
					UserType:    "employer",
					Image:       "basic_applicant_avatar.webp",
				}
				r.EXPECT().GetUser(id).Return(expectedUser, errorHandler2.ErrUserNotExists)
			},
			sessionRepBehavior: func(sessionRep *mock_session.MockRepository, token string) {
				sessionRep.EXPECT().GetSession(token).Return("", fmt.Errorf("getting session error:"))
			},
			expectedStatusCode:   401,
			expectedResponseBody: "{\"descriptors\":\"\",\"error\":\"Клиент не авторизован\"}",
		},
	}
	for _, test := range testTable {
		testCase := test
		t.Run(testCase.name, func(t *testing.T) {
			t.Parallel()
			c := gomock.NewController(t)
			defer c.Finish()

			mockUseCase := mock_usecases.NewMockUser(c)
			sessionRep := mock_session.NewMockRepository(c)
			sessionMiddlware := middleware.NewSessionMiddleware(sessionRep)

			if testCase.emailFromToken != "" {
				testCase.mockBehavior(mockUseCase, testCase.inputId)
			}

			testCase.sessionRepBehavior(sessionRep, testCase.inputToken)

			handler := UserHandler{
				userUseCase: mockUseCase,
				cfg: &configs.Config{
					DefaultExpiringSession: 100,
					Cookie: configs.CookieConfig{
						Secure:   false,
						HTTPOnly: true,
					},
				},
			}

			r := gin.New()
			r.GET("/:id", sessionMiddlware.Session, handler.GetUser, errorHandler2.Middleware())

			cookie := &http.Cookie{
				Name:  "session",
				Value: testCase.inputToken,
			}

			w := httptest.NewRecorder()
			req := httptest.NewRequest("GET", "/"+testCase.requestParam,
				bytes.NewBufferString(""))

			req.AddCookie(cookie)
			r.ServeHTTP(w, req)

			assert.Equal(t, testCase.expectedStatusCode, w.Code)
			assert.Equal(t, testCase.expectedResponseBody, w.Body.String())
		})
	}

}

func TestUserHandler_GetUserSafety(t *testing.T) {
	type mockBehavior func(r *mock_usecases.MockUser, id uint)

	testTable := []struct {
		name                 string
		inputId              uint
		requestParam         string
		mockBehavior         mockBehavior
		expectedStatusCode   int
		expectedResponseBody string
	}{
		{
			name:         "valid",
			inputId:      42,
			requestParam: "42",
			mockBehavior: func(r *mock_usecases.MockUser, id uint) {
				expectedUser := &models.UserAccount{
					ID:          42,
					Email:       "test@gmail.com",
					CompanyName: "Some company",
					UserType:    "employer",
					Image:       "basic_applicant_avatar.webp",
				}
				r.EXPECT().GetUserSafety(id).Return(expectedUser, nil)
			},
			expectedStatusCode:   200,
			expectedResponseBody: "{\"id\":42,\"user_type\":\"employer\",\"email\":\"test@gmail.com\",\"password\":\"\",\"contact_number\":\"\",\"status\":\"\",\"description\":\"\",\"image\":\"basic_applicant_avatar.webp\",\"average_color\":\"\",\"date_of_birth\":\"0001-01-01T00:00:00Z\",\"created_time\":\"0001-01-01T00:00:00Z\",\"company_name\":\"Some company\",\"company_size\":0,\"public_fields\":\"\",\"is_confirmed\":false,\"two_factor_sign_in\":false,\"mailing_approval\":false,\"resumes\":null,\"vacancies\":null,\"favourite_vacancies\":null,\"vacancy_activities\":null}",
		},
		{
			name:         "not found",
			inputId:      42,
			requestParam: "42",
			mockBehavior: func(r *mock_usecases.MockUser, id uint) {
				expectedUser := &models.UserAccount{
					ID:          42,
					Email:       "test@gmail.com",
					CompanyName: "Some company",
					UserType:    "employer",
					Image:       "basic_applicant_avatar.webp",
				}
				r.EXPECT().GetUserSafety(id).Return(expectedUser, errorHandler2.ErrUserNotExists)
			},
			expectedStatusCode:   401,
			expectedResponseBody: "{\"descriptors\":\"email\",\"error\":\"Пользователя с таким email не существует\"}",
		},
	}
	for _, test := range testTable {
		testCase := test
		t.Run(testCase.name, func(t *testing.T) {
			t.Parallel()
			c := gomock.NewController(t)
			defer c.Finish()

			mockUseCase := mock_usecases.NewMockUser(c)
			testCase.mockBehavior(mockUseCase, testCase.inputId)

			handler := UserHandler{
				userUseCase: mockUseCase,
				cfg: &configs.Config{
					DefaultExpiringSession: 100,
					Cookie: configs.CookieConfig{
						Secure:   false,
						HTTPOnly: true,
					},
				},
			}

			r := gin.New()
			r.GET("/safety/:id", handler.GetUserSafety, errorHandler2.Middleware())

			w := httptest.NewRecorder()
			req := httptest.NewRequest("GET", "/safety/"+testCase.requestParam,
				bytes.NewBufferString(""))

			r.ServeHTTP(w, req)

			assert.Equal(t, testCase.expectedStatusCode, w.Code)
			assert.Equal(t, testCase.expectedResponseBody, w.Body.String())
		})
	}

}

func TestUserHandler_GetPreview(t *testing.T) {
	type mockBehavior func(r *mock_usecases.MockUser, id uint)

	testTable := []struct {
		name                 string
		inputId              uint
		requestParam         string
		mockBehavior         mockBehavior
		expectedStatusCode   int
		expectedResponseBody string
	}{
		{
			name:         "valid",
			inputId:      42,
			requestParam: "42",
			mockBehavior: func(r *mock_usecases.MockUser, id uint) {
				expectedUser := &models.UserAccount{
					ID:          42,
					Email:       "test@gmail.com",
					CompanyName: "Some company",
					UserType:    "employer",
					Image:       "basic_applicant_avatar.webp",
				}
				r.EXPECT().GetUserSafety(id).Return(expectedUser, nil)
			},
			expectedStatusCode:   200,
			expectedResponseBody: "{\"id\":42,\"user_type\":\"employer\",\"email\":\"\",\"password\":\"\",\"contact_number\":\"\",\"status\":\"\",\"description\":\"\",\"image\":\"basic_applicant_avatar.webp\",\"average_color\":\"\",\"date_of_birth\":\"0001-01-01T00:00:00Z\",\"created_time\":\"0001-01-01T00:00:00Z\",\"company_name\":\"Some company\",\"company_size\":0,\"public_fields\":\"\",\"is_confirmed\":false,\"two_factor_sign_in\":false,\"mailing_approval\":false,\"resumes\":null,\"vacancies\":null,\"favourite_vacancies\":null,\"vacancy_activities\":null}",
		},
		{
			name:         "not found",
			inputId:      42,
			requestParam: "42",
			mockBehavior: func(r *mock_usecases.MockUser, id uint) {
				expectedUser := &models.UserAccount{
					ID:          42,
					Email:       "test@gmail.com",
					CompanyName: "Some company",
					UserType:    "employer",
					Image:       "basic_applicant_avatar.webp",
				}
				r.EXPECT().GetUserSafety(id).Return(expectedUser, errorHandler2.ErrUserNotExists)
			},
			expectedStatusCode:   401,
			expectedResponseBody: "{\"descriptors\":\"email\",\"error\":\"Пользователя с таким email не существует\"}",
		},
	}
	for _, test := range testTable {
		testCase := test
		t.Run(testCase.name, func(t *testing.T) {
			t.Parallel()
			c := gomock.NewController(t)
			defer c.Finish()

			mockUseCase := mock_usecases.NewMockUser(c)
			testCase.mockBehavior(mockUseCase, testCase.inputId)

			handler := UserHandler{
				userUseCase: mockUseCase,
				cfg: &configs.Config{
					DefaultExpiringSession: 100,
					Cookie: configs.CookieConfig{
						Secure:   false,
						HTTPOnly: true,
					},
				},
			}

			r := gin.New()
			r.GET("/preview/:id", handler.GetPreview, errorHandler2.Middleware())

			w := httptest.NewRecorder()
			req := httptest.NewRequest("GET", "/preview/"+testCase.requestParam,
				bytes.NewBufferString(""))

			r.ServeHTTP(w, req)

			assert.Equal(t, testCase.expectedStatusCode, w.Code)
			assert.Equal(t, testCase.expectedResponseBody, w.Body.String())
		})
	}

}

func TestUserHandler_UpdateUser(t *testing.T) {
	type mockBehavior func(r *mock_usecases.MockUser, user *models.UserAccount)
	type sessionRepBehavior func(r *mock_session.MockRepository, token string)
	testTable := []struct {
		name                 string
		mockBehavior         mockBehavior
		sessionRepBehavior   sessionRepBehavior
		inputUser            models.UserAccount
		inputBody            string
		inputToken           string
		emailFromToken       string
		expectedStatusCode   int
		expectedResponseBody string
	}{
		{
			name:           "ok",
			inputToken:     "valid_token",
			emailFromToken: "test@gmail.com",
			mockBehavior: func(r *mock_usecases.MockUser, user *models.UserAccount) {
				r.EXPECT().UpdateUser(user).Return(nil)
			},
			inputUser: models.UserAccount{
				UserType:         "applicant",
				ApplicantName:    "Ivan",
				ApplicantSurname: "Kozirev",
				Status:           "Hello World",
				ContactNumber:    "+7 (968) 333-33-33",
				Email:            "test@gmail.com",
			},
			inputBody: `{
				"user_type": "applicant",
				"applicant_name": "Ivan",
				"applicant_surname": "Kozirev",
				"status": "Hello World",
				"contact_number": "+7 (968) 333-33-33",
				"email": "test@gmail.com"
			}`,
			sessionRepBehavior: func(sessionRep *mock_session.MockRepository, token string) {
				sessionRep.EXPECT().GetSession(token).Return("test@gmail.com", nil)
			},
			expectedStatusCode:   200,
			expectedResponseBody: "{\"id\":0,\"user_type\":\"applicant\",\"email\":\"test@gmail.com\",\"password\":\"\",\"contact_number\":\"\",\"status\":\"\",\"description\":\"\",\"image\":\"\",\"average_color\":\"\",\"date_of_birth\":\"0001-01-01T00:00:00Z\",\"created_time\":\"0001-01-01T00:00:00Z\",\"applicant_name\":\"Ivan\",\"applicant_surname\":\"Kozirev\",\"company_size\":0,\"public_fields\":\"\",\"is_confirmed\":false,\"two_factor_sign_in\":false,\"mailing_approval\":false,\"resumes\":null,\"vacancies\":null,\"favourite_vacancies\":null,\"vacancy_activities\":null}",
		},
		{
			name:           "error",
			inputToken:     "valid_token",
			emailFromToken: "test@gmail.com",
			mockBehavior: func(r *mock_usecases.MockUser, user *models.UserAccount) {
				r.EXPECT().UpdateUser(user).Return(errorHandler2.InvalidUserType)
			},
			inputUser: models.UserAccount{
				UserType:         "aplicant",
				ApplicantName:    "Ivan",
				ApplicantSurname: "Kozirev",
				Status:           "Hello World",
				ContactNumber:    "+7 (968) 333-33-33",
				Email:            "test@gmail.com",
			},
			inputBody: `{
				"user_type": "aplicant",
				"applicant_name": "Ivan",
				"applicant_surname": "Kozirev",
				"status": "Hello World",
				"contact_number": "+7 (968) 333-33-33",
				"email": "test@gmail.com"
			}`,
			sessionRepBehavior: func(sessionRep *mock_session.MockRepository, token string) {
				sessionRep.EXPECT().GetSession(token).Return("test@gmail.com", nil)
			},
			expectedStatusCode:   400,
			expectedResponseBody: "{\"descriptors\":\"\",\"error\":\"Некорректный входной тип пользователя\"}",
		},
		{
			name:           "unauthorized",
			inputToken:     "valid_token",
			emailFromToken: "",
			mockBehavior: func(r *mock_usecases.MockUser, user *models.UserAccount) {
				r.EXPECT().UpdateUser(user).Return(errorHandler2.InvalidUserType)
			},
			inputUser: models.UserAccount{
				UserType:         "aplicant",
				ApplicantName:    "Ivan",
				ApplicantSurname: "Kozirev",
				Status:           "Hello World",
				ContactNumber:    "+7 (968) 333-33-33",
				Email:            "test@gmail.com",
			},
			inputBody: `{
				"user_type": "aplicant",
				"applicant_name": "Ivan",
				"applicant_surname": "Kozirev",
				"status": "Hello World",
				"contact_number": "+7 (968) 333-33-33",
				"email": "test@gmail.com"
			}`,
			sessionRepBehavior: func(sessionRep *mock_session.MockRepository, token string) {
				sessionRep.EXPECT().GetSession(token).Return("", fmt.Errorf("getting session error:"))
			},
			expectedStatusCode:   401,
			expectedResponseBody: "{\"descriptors\":\"\",\"error\":\"Клиент не авторизован\"}",
		},
	}
	for _, test := range testTable {
		testCase := test
		t.Run(testCase.name, func(t *testing.T) {
			t.Parallel()
			c := gomock.NewController(t)
			defer c.Finish()

			mockUseCase := mock_usecases.NewMockUser(c)
			sessionRep := mock_session.NewMockRepository(c)
			sessionMiddlware := middleware.NewSessionMiddleware(sessionRep)
			handler := UserHandler{
				userUseCase: mockUseCase,
				cfg: &configs.Config{
					DefaultExpiringSession: 100,
					Cookie: configs.CookieConfig{
						Secure:   false,
						HTTPOnly: true,
					},
				},
			}
			if testCase.emailFromToken != "" {
				testCase.mockBehavior(mockUseCase, &testCase.inputUser)
			}
			testCase.sessionRepBehavior(sessionRep, testCase.inputToken)
			r := gin.New()
			r.POST("/user", sessionMiddlware.Session, handler.UpdateUser, errorHandler2.Middleware())
			w := httptest.NewRecorder()

			cookie := &http.Cookie{
				Name:  "session",
				Value: testCase.inputToken,
			}
			req := httptest.NewRequest("POST", "/user",
				bytes.NewBufferString(testCase.inputBody))

			req.AddCookie(cookie)
			r.ServeHTTP(w, req)

			assert.Equal(t, testCase.expectedStatusCode, w.Code)
			assert.Equal(t, testCase.expectedResponseBody, w.Body.String())
		})
	}
}

func TestUserHandler_UploadUserImage(t *testing.T) {

}

func TestUserHandler_DeleteUserImage(t *testing.T) {
	type deleteImageMockBehavior func(r *mock_usecases.MockUser, user *models.UserAccount)
	type getUserMockBehavior func(r *mock_usecases.MockUser, email string)
	type sessionRepBehavior func(r *mock_session.MockRepository, token string)

	testTable := []struct {
		name                    string
		deleteImageMockBehavior deleteImageMockBehavior
		getUserMockBehavior     getUserMockBehavior
		sessionRepBehavior      sessionRepBehavior
		inputUser               *models.UserAccount
		inputToken              string
		emailFromToken          string
		expectedStatusCode      int
		expectedResponseBody    string
	}{
		{
			name:           "ok",
			inputToken:     "valid_token",
			emailFromToken: "test@gmail.com",
			getUserMockBehavior: func(r *mock_usecases.MockUser, email string) {
				expectedUser := &models.UserAccount{
					ID:          100,
					Email:       "test@gmail.com",
					CompanyName: "Some company",
					UserType:    "employer",
					Image:       "basic_applicant_avatar.webp",
				}
				r.EXPECT().GetUserByEmail(email).Return(expectedUser, nil)
			},
			inputUser: &models.UserAccount{
				ID:          100,
				Email:       "test@gmail.com",
				CompanyName: "Some company",
				UserType:    "employer",
				Image:       "basic_applicant_avatar.webp",
			},
			sessionRepBehavior: func(sessionRep *mock_session.MockRepository, token string) {
				sessionRep.EXPECT().GetSession(token).Return("test@gmail.com", nil)
			},
			deleteImageMockBehavior: func(r *mock_usecases.MockUser, user *models.UserAccount) {
				r.EXPECT().DeleteUserImage(user).Return(nil)
			},
			expectedStatusCode:   200,
			expectedResponseBody: "",
		},
		{
			name:           "cant delete",
			inputToken:     "valid_token",
			emailFromToken: "test@gmail.com",
			getUserMockBehavior: func(r *mock_usecases.MockUser, email string) {
				expectedUser := &models.UserAccount{
					ID:          100,
					Email:       "test@gmail.com",
					CompanyName: "Some company",
					UserType:    "employer",
					Image:       "basic_applicant_avatar.webp",
				}
				r.EXPECT().GetUserByEmail(email).Return(expectedUser, nil)
			},
			inputUser: &models.UserAccount{
				ID:          100,
				Email:       "test@gmail.com",
				CompanyName: "Some company",
				UserType:    "employer",
				Image:       "basic_applicant_avatar.webp",
			},
			sessionRepBehavior: func(sessionRep *mock_session.MockRepository, token string) {
				sessionRep.EXPECT().GetSession(token).Return("test@gmail.com", nil)
			},
			deleteImageMockBehavior: func(r *mock_usecases.MockUser, user *models.UserAccount) {
				r.EXPECT().DeleteUserImage(user).Return(errorHandler2.ErrBadRequest)
			},
			expectedStatusCode:   400,
			expectedResponseBody: "{\"descriptors\":\"\",\"error\":\"Некорректный запрос\"}",
		},
		{
			name:           "invalid session",
			inputToken:     "valid_token",
			emailFromToken: "",
			getUserMockBehavior: func(r *mock_usecases.MockUser, email string) {
				expectedUser := &models.UserAccount{
					ID:          100,
					Email:       "test@gmail.com",
					CompanyName: "Some company",
					UserType:    "employer",
					Image:       "basic_applicant_avatar.webp",
				}
				r.EXPECT().GetUserByEmail(email).Return(expectedUser, nil)
			},
			inputUser: &models.UserAccount{
				ID:          100,
				Email:       "test@gmail.com",
				CompanyName: "Some company",
				UserType:    "employer",
				Image:       "basic_applicant_avatar.webp",
			},
			sessionRepBehavior: func(sessionRep *mock_session.MockRepository, token string) {
				sessionRep.EXPECT().GetSession(token).Return("", fmt.Errorf("getting session error:"))
			},
			deleteImageMockBehavior: func(r *mock_usecases.MockUser, user *models.UserAccount) {
				r.EXPECT().DeleteUserImage(user).Return(errorHandler2.ErrBadRequest)
			},
			expectedStatusCode:   401,
			expectedResponseBody: "{\"descriptors\":\"\",\"error\":\"Клиент не авторизован\"}",
		},
	}
	for _, test := range testTable {
		testCase := test
		t.Run(testCase.name, func(t *testing.T) {
			t.Parallel()
			c := gomock.NewController(t)
			defer c.Finish()

			mockUseCase := mock_usecases.NewMockUser(c)
			sessionRep := mock_session.NewMockRepository(c)
			sessionMiddlware := middleware.NewSessionMiddleware(sessionRep)
			handler := UserHandler{
				userUseCase: mockUseCase,
				cfg: &configs.Config{
					DefaultExpiringSession: 100,
					Cookie: configs.CookieConfig{
						Secure:   false,
						HTTPOnly: true,
					},
				},
			}
			if testCase.emailFromToken != "" {
				testCase.getUserMockBehavior(mockUseCase, testCase.emailFromToken)
				if testCase.inputUser != nil {
					testCase.deleteImageMockBehavior(mockUseCase, testCase.inputUser)
				}
			}
			testCase.sessionRepBehavior(sessionRep, testCase.inputToken)
			r := gin.New()
			r.DELETE("/image", sessionMiddlware.Session, handler.DeleteUserImage, errorHandler2.Middleware())
			w := httptest.NewRecorder()

			cookie := &http.Cookie{
				Name:  "session",
				Value: testCase.inputToken,
			}
			req := httptest.NewRequest("DELETE", "/image",
				bytes.NewBufferString(""))

			req.AddCookie(cookie)
			r.ServeHTTP(w, req)

			assert.Equal(t, testCase.expectedStatusCode, w.Code)
			assert.Equal(t, testCase.expectedResponseBody, w.Body.String())
		})
	}
}

func TestUserHandler_GetAllApplicants(t *testing.T) {
	type mockBehavior func(r *mock_usecases.MockUser, filters models.UserFilter)
	//type sessionRepBehavior func(r *mock_session.MockRepository, token string)

	testTable := []struct {
		name                 string
		inputId              uint
		inputToken           string
		requestParam         string
		emailFromToken       string
		filters              models.UserFilter
		mockBehavior         mockBehavior
		expectedStatusCode   int
		expectedResponseBody string
	}{
		{
			name:           "valid",
			inputId:        42,
			requestParam:   "42",
			emailFromToken: "test@gmail.com",
			mockBehavior: func(r *mock_usecases.MockUser, filters models.UserFilter) {
				expectedResume := []*models.UserAccount{
					{
						ID:            42,
						ApplicantName: "some name",
					},
				}
				r.EXPECT().GetAllApplicants(filters).Return(expectedResume, nil)
			},
			expectedStatusCode:   200,
			expectedResponseBody: "{\"data\":[{\"id\":42,\"user_type\":\"\",\"email\":\"\",\"password\":\"\",\"contact_number\":\"\",\"status\":\"\",\"description\":\"\",\"image\":\"\",\"average_color\":\"\",\"date_of_birth\":\"0001-01-01T00:00:00Z\",\"created_time\":\"0001-01-01T00:00:00Z\",\"applicant_name\":\"some name\",\"company_size\":0,\"public_fields\":\"\",\"is_confirmed\":false,\"two_factor_sign_in\":false,\"mailing_approval\":false,\"resumes\":null,\"vacancies\":null,\"favourite_vacancies\":null,\"vacancy_activities\":null}]}",
		},
		{
			name:           "not found",
			inputId:        42,
			requestParam:   "42",
			emailFromToken: "test@gmail.com",
			mockBehavior: func(r *mock_usecases.MockUser, filters models.UserFilter) {
				expectedResume := []*models.UserAccount{
					{
						ID:            42,
						ApplicantName: "some name",
					},
				}
				r.EXPECT().GetAllApplicants(filters).Return(expectedResume, errorHandler2.ErrResumeNotFound)
			},
			expectedStatusCode:   404,
			expectedResponseBody: "{\"descriptors\":\"\",\"error\":\"Резюме не найдено\"}",
		},
	}
	for _, test := range testTable {
		testCase := test
		t.Run(testCase.name, func(t *testing.T) {
			t.Parallel()
			c := gomock.NewController(t)
			defer c.Finish()

			mockUseCase := mock_usecases.NewMockUser(c)

			if testCase.emailFromToken != "" {
				testCase.mockBehavior(mockUseCase, testCase.filters)
			}

			handler := UserHandler{
				userUseCase: mockUseCase,
			}

			r := gin.New()
			r.GET("/", handler.GetAllApplicants, errorHandler2.Middleware())

			cookie := &http.Cookie{
				Name:  "session",
				Value: testCase.inputToken,
			}

			w := httptest.NewRecorder()
			req := httptest.NewRequest("GET", "/",
				bytes.NewBufferString(""))

			req.AddCookie(cookie)
			r.ServeHTTP(w, req)

			assert.Equal(t, testCase.expectedStatusCode, w.Code)
			assert.Equal(t, testCase.expectedResponseBody, w.Body.String())
		})
	}
}

func TestUserHandler_GetAllEmployers(t *testing.T) {
	type mockBehavior func(r *mock_usecases.MockUser, filters models.UserFilter)
	//type sessionRepBehavior func(r *mock_session.MockRepository, token string)

	testTable := []struct {
		name                 string
		inputId              uint
		inputToken           string
		requestParam         string
		emailFromToken       string
		filters              models.UserFilter
		mockBehavior         mockBehavior
		expectedStatusCode   int
		expectedResponseBody string
	}{
		{
			name:           "valid",
			inputId:        42,
			requestParam:   "42",
			emailFromToken: "test@gmail.com",
			mockBehavior: func(r *mock_usecases.MockUser, filters models.UserFilter) {
				expectedResume := []*models.UserAccount{
					{
						ID:          42,
						CompanyName: "some name",
					},
				}
				r.EXPECT().GetAllEmployers(filters).Return(expectedResume, nil)
			},
			expectedStatusCode:   200,
			expectedResponseBody: "{\"data\":[{\"id\":42,\"user_type\":\"\",\"email\":\"\",\"password\":\"\",\"contact_number\":\"\",\"status\":\"\",\"description\":\"\",\"image\":\"\",\"average_color\":\"\",\"date_of_birth\":\"0001-01-01T00:00:00Z\",\"created_time\":\"0001-01-01T00:00:00Z\",\"company_name\":\"some name\",\"company_size\":0,\"public_fields\":\"\",\"is_confirmed\":false,\"two_factor_sign_in\":false,\"mailing_approval\":false,\"resumes\":null,\"vacancies\":null,\"favourite_vacancies\":null,\"vacancy_activities\":null}]}",
		},
		{
			name:           "not found",
			inputId:        42,
			requestParam:   "42",
			emailFromToken: "test@gmail.com",
			mockBehavior: func(r *mock_usecases.MockUser, filters models.UserFilter) {
				expectedResume := []*models.UserAccount{
					{
						ID:          42,
						CompanyName: "some name",
					},
				}
				r.EXPECT().GetAllEmployers(filters).Return(expectedResume, errorHandler2.ErrResumeNotFound)
			},
			expectedStatusCode:   404,
			expectedResponseBody: "{\"descriptors\":\"\",\"error\":\"Резюме не найдено\"}",
		},
	}
	for _, test := range testTable {
		testCase := test
		t.Run(testCase.name, func(t *testing.T) {
			t.Parallel()
			c := gomock.NewController(t)
			defer c.Finish()

			mockUseCase := mock_usecases.NewMockUser(c)

			if testCase.emailFromToken != "" {
				testCase.mockBehavior(mockUseCase, testCase.filters)
			}

			handler := UserHandler{
				userUseCase: mockUseCase,
			}

			r := gin.New()
			r.GET("/", handler.GetAllEmployers, errorHandler2.Middleware())

			cookie := &http.Cookie{
				Name:  "session",
				Value: testCase.inputToken,
			}

			w := httptest.NewRecorder()
			req := httptest.NewRequest("GET", "/",
				bytes.NewBufferString(""))

			req.AddCookie(cookie)
			r.ServeHTTP(w, req)

			assert.Equal(t, testCase.expectedStatusCode, w.Code)
			assert.Equal(t, testCase.expectedResponseBody, w.Body.String())
		})
	}
}
