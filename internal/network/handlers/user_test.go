package handlers

import (
	"HeadHunter/configs"
	"HeadHunter/internal/entity/models"
	"HeadHunter/internal/errorHandler"
	"HeadHunter/internal/network/middleware"
	mock_usecases "HeadHunter/internal/usecases/mocks"
	"bytes"
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
			expectedResponseBody: "{\"id\":0,\"user_type\":\"applicant\",\"email\":\"test@gmail.com\",\"password\":\"\",\"contact_number\":\"\",\"status\":\"\",\"description\":\"\",\"image\":\"\",\"date_of_birth\":\"0001-01-01T00:00:00Z\",\"applicant_name\":\"Zakhar\",\"applicant_surname\":\"Urvancev\",\"company_size\":0,\"resumes\":null,\"vacancies\":null,\"vacancy_activities\":null}",
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
			expectedResponseBody: "{\"id\":0,\"user_type\":\"employer\",\"email\":\"test@gmail.com\",\"password\":\"\",\"contact_number\":\"\",\"status\":\"\",\"description\":\"\",\"image\":\"\",\"date_of_birth\":\"0001-01-01T00:00:00Z\",\"company_name\":\"Some company\",\"company_size\":0,\"resumes\":null,\"vacancies\":null,\"vacancy_activities\":null}",
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
				r.EXPECT().SignUp(user).Return("", errorHandler.ErrUserExists)
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
			test.mockBehavior(mockUseCase, &test.inputUser)

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
			r.POST("/sign-up", handler.SignUp)

			w := httptest.NewRecorder()
			req := httptest.NewRequest("POST", "/sign-up",
				bytes.NewBufferString(test.inputBody))

			r.ServeHTTP(w, req)

			assert.Equal(t, test.expectedStatusCode, w.Code)
			assert.Equal(t, test.expectedResponseBody, w.Body.String())
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
			expectedResponseBody: "{\"id\":0,\"user_type\":\"applicant\",\"email\":\"test@gmail.com\",\"password\":\"\",\"contact_number\":\"\",\"status\":\"\",\"description\":\"\",\"image\":\"\",\"date_of_birth\":\"0001-01-01T00:00:00Z\",\"company_size\":0,\"resumes\":null,\"vacancies\":null,\"vacancy_activities\":null}",
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
				r.EXPECT().SignIn(user).Return(gomock.Any().String(), errorHandler.InvalidUserType)
			},
			expectedStatusCode:   400,
			expectedResponseBody: "{\"descriptors\":[],\"error\":\"invalid input user type\"}",
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
				r.EXPECT().SignIn(user).Return("", errorHandler.ErrUserNotExists)
			},
			expectedStatusCode:   401,
			expectedResponseBody: "{\"descriptors\":[\"email\"],\"error\":\"Пользователя с таким email не существует\"}",
		},
	}
	for _, test := range testTable {
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()
			c := gomock.NewController(t)
			defer c.Finish()

			mockUseCase := mock_usecases.NewMockUser(c)
			test.mockBehavior(mockUseCase, &test.inputUser)

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
			r.POST("/sign-in", handler.SignIn, middleware.ErrorHandler())

			w := httptest.NewRecorder()
			req := httptest.NewRequest("POST", "/sign-in",
				bytes.NewBufferString(test.inputBody))

			r.ServeHTTP(w, req)

			assert.Equal(t, test.expectedStatusCode, w.Code)
			assert.Equal(t, test.expectedResponseBody, w.Body.String())
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
				r.EXPECT().Logout(token).Return(errorHandler.ErrBadRequest)
			},
			expectedStatusCode:   400,
			expectedResponseBody: "{\"descriptors\":[],\"error\":\"bad request\"}",
		},
		{
			name:       "with cookie",
			inputToken: "valid_token",
			mockBehavior: func(r *mock_usecases.MockUser, token string) {
				r.EXPECT().Logout(token).Return(errorHandler.ErrBadRequest)
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

	testTable := []struct {
		name                 string
		mockBehavior         mockBehavior
		inputToken           string
		expectedUser         models.UserAccount
		expectedStatusCode   int
		expectedResponseBody string
	}{
		{
			name:       "no cookie",
			inputToken: "invalid_token",
			mockBehavior: func(r *mock_usecases.MockUser, email string) {
				expectedUser := &models.UserAccount{
					Email:       "test@gmail.com",
					CompanyName: "Some company",
					UserType:    "employer",
					Image:       "basic_applicant_avatar.webp",
				}
				r.EXPECT().AuthCheck(email).Return(expectedUser, nil)
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
			r.GET("/auth", handler.AuthCheck)
			w := httptest.NewRecorder()

			cookie := &http.Cookie{
				Name:  "session",
				Value: "valid_token",
			}
			req := httptest.NewRequest("GET", "/auth",
				bytes.NewBufferString(""))

			req.AddCookie(cookie)
			r.ServeHTTP(w, req)

			assert.Equal(t, test.expectedStatusCode, w.Code)
			assert.Equal(t, test.expectedResponseBody, w.Body.String())
		})
	}
}

func TestUserHandler_UploadUserImage(t *testing.T) {

}

func TestUserHandler_GetUser(t *testing.T) {

}

func TestUserHandler_GetUserSafety(t *testing.T) {

}

func TestUserHandler_UpdateUser(t *testing.T) {

}

func TestUserHandler_DeleteUserImage(t *testing.T) {

}

func TestUserHandler_GetPreview(t *testing.T) {

}
