package handlers

import (
	"HeadHunter/internal/entity/models"
	"HeadHunter/internal/errorHandler"
	mock_usecases "HeadHunter/internal/usecases/mocks"
	"bytes"
	"github.com/gin-gonic/gin"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"net/http/httptest"
	"testing"
)

func TestUserHandler_SignUp(t *testing.T) {
	type mockBehavior func(r *mock_usecases.MockUser, user models.UserAccount)

	testTable := []struct {
		name                 string
		inputUser            models.UserAccount
		expectedUser         models.UserAccount
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
			expectedUser: models.UserAccount{
				ID:               1,
				Email:            "test@gmail.com",
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
			mockBehavior: func(r *mock_usecases.MockUser, user models.UserAccount) {
				expectedUser := models.UserAccount{
					ID:               1,
					Email:            "test@gmail.com",
					ApplicantSurname: "Urvancev",
					ApplicantName:    "Zakhar",
					UserType:         "applicant",
				}
				r.EXPECT().SignUp(user).Return("", errorHandler.ErrBadRequest)
			},
			expectedStatusCode: 200,
			expectedResponseBody: `{
    "id": 1,
    "user_type": "applicant",
    "email": "test@gmail.com",
    "password": "",
    "contact_number": "",
    "status": "",
    "description": "",
    "image": "basic_applicant_avatar.webp",
    "date_of_birth": "0001-01-01T00:00:00Z",
    "applicant_name": "Zakhar",
    "applicant_surname": "Urvancev",
    "company_size": 0,
    "resumes": null,
    "vacancies": null,
    "vacancy_activities": null
}`,
		},
		{
			name: "valid employer",
			inputUser: models.UserAccount{
				Email:            "test2@gmail.com",
				Password:         "123456a!",
				ApplicantSurname: "Urvancev",
				ApplicantName:    "Zakhar",
				UserType:         "employer",
			},
			inputBody: `{
    			"email": "test2@gmail.com",
    			"password": "123456a!",
                "company_name": "Some company",
    			"user_type": "employer"
}`,
			mockBehavior: func(r *mock_usecases.MockUser, user models.UserAccount) {
				r.EXPECT().SignUp(user).Return(1, nil)
			},
			expectedStatusCode: 200,
			expectedResponseBody: `{
    "id": 5,
    "user_type": "employer",
    "email": "test2@gmail.com",
    "password": "",
    "contact_number": "",
    "status": "",
    "description": "",
    "image": "basic_employer_avatar.webp",
    "date_of_birth": "0001-01-01T00:00:00Z",
    "company_name": "Some company",
    "company_size": 0,
    "resumes": null,
    "vacancies": null,
    "vacancy_activities": null
}`,
		},
	}
	for _, test := range testTable {
		t.Run(test.name, func(t *testing.T) {
			// Init Dependencies
			c := gomock.NewController(t)
			defer c.Finish()

			mockUseCase := mock_usecases.NewMockUser(c)
			test.mockBehavior(mockUseCase, test.inputUser)

			handler := UserHandler{userUseCase: mockUseCase}

			// Init Endpoint
			r := gin.New()
			r.POST("/sign-up", handler.SignUp)

			// Create Request
			w := httptest.NewRecorder()
			req := httptest.NewRequest("POST", "/sign-up",
				bytes.NewBufferString(test.inputBody))

			// Make Request
			r.ServeHTTP(w, req)

			// Assert
			assert.Equal(t, w.Code, test.expectedStatusCode)
			assert.Equal(t, w.Body.String(), test.expectedResponseBody)
		})
	}
}

func TestUserHandler_SignIn(t *testing.T) {

}

func TestUserHandler_Logout(t *testing.T) {

}

func TestUserHandler_AuthCheck(t *testing.T) {

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
