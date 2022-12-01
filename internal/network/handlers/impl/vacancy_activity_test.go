package impl

import (
	mock_session "HeadHunter/common/session/mocks"
	"HeadHunter/internal/entity/models"
	"HeadHunter/internal/network/middleware"
	mock_usecases "HeadHunter/internal/usecases/mocks"
	"HeadHunter/pkg/errorHandler"
	"bytes"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestVacancyActivityHandler_ApplyForVacancy(t *testing.T) {
	type mockBehavior func(r *mock_usecases.MockVacancyActivity, vacancyActivity *models.VacancyActivity, vacancyId uint)
	type sessionRepBehavior func(r *mock_session.MockRepository, token string)
	testTable := []struct {
		name                 string
		inputToken           string
		requestParam         string
		emailFromToken       string
		createdApply         *models.VacancyActivity
		inputBody            string
		vacancyId            uint
		mockBehavior         mockBehavior
		sessionRepBehavior   sessionRepBehavior
		expectedStatusCode   int
		expectedResponseBody string
	}{
		{
			name:           "valid",
			emailFromToken: "test@gmail.com",
			vacancyId:      1,
			requestParam:   "1",
			createdApply: &models.VacancyActivity{
				ResumeTitle: "some title",
			},
			inputBody: `{
				"title": "some title"
}`,
			mockBehavior: func(r *mock_usecases.MockVacancyActivity, vacancyActivity *models.VacancyActivity, vacancyId uint) {
				createdApply := &models.VacancyActivity{ResumeTitle: "some title"}
				r.EXPECT().ApplyForVacancy("test@gmail.com", vacancyId, createdApply).Return(nil)
			},
			sessionRepBehavior: func(sessionRep *mock_session.MockRepository, token string) {
				sessionRep.EXPECT().GetSession(token).Return("test@gmail.com", nil)
			},
			expectedStatusCode:   200,
			expectedResponseBody: "",
		},
		{
			name:           "user not found",
			emailFromToken: "",
			vacancyId:      1,
			requestParam:   "1",
			createdApply: &models.VacancyActivity{
				ResumeTitle: "some title",
			},
			inputBody: `{
    			"title": "some title"
}`,
			mockBehavior: func(r *mock_usecases.MockVacancyActivity, vacancy *models.VacancyActivity, vacancyId uint) {
				createdApply := &models.VacancyActivity{ResumeTitle: "some title"}
				r.EXPECT().ApplyForVacancy("test@gmail.com", vacancyId, createdApply).Return(errorHandler.ErrBadRequest)
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

			mockUseCase := mock_usecases.NewMockVacancyActivity(c)
			sessionRep := mock_session.NewMockRepository(c)
			sessionMiddlware := middleware.NewSessionMiddleware(sessionRep)

			if testCase.emailFromToken != "" {
				testCase.mockBehavior(mockUseCase, testCase.createdApply, testCase.vacancyId)
			}

			testCase.sessionRepBehavior(sessionRep, testCase.inputToken)

			handler := VacancyActivityHandler{
				vacancyActivityUseCase: mockUseCase,
			}

			r := gin.New()
			r.POST("/apply/:id", sessionMiddlware.Session, handler.ApplyForVacancy, errorHandler.Middleware())

			cookie := &http.Cookie{
				Name:  "session",
				Value: testCase.inputToken,
			}

			w := httptest.NewRecorder()
			req := httptest.NewRequest("POST", "/apply/"+testCase.requestParam,
				bytes.NewBufferString(test.inputBody))

			req.AddCookie(cookie)
			r.ServeHTTP(w, req)

			assert.Equal(t, testCase.expectedStatusCode, w.Code)
			assert.Equal(t, testCase.expectedResponseBody, w.Body.String())
		})
	}
}

func TestVacancyActivityHandler_DeleteUserApply(t *testing.T) {
	type mockBehavior func(r *mock_usecases.MockVacancyActivity, id uint)
	type sessionRepBehavior func(r *mock_session.MockRepository, token string)

	testTable := []struct {
		name                 string
		inputToken           string
		inputId              uint
		inputParam           string
		emailFromToken       string
		mockBehavior         mockBehavior
		sessionRepBehavior   sessionRepBehavior
		expectedStatusCode   int
		expectedResponseBody string
	}{
		{
			name:           "valid",
			inputId:        42,
			inputParam:     "42",
			emailFromToken: "test@gmail.com",
			mockBehavior: func(r *mock_usecases.MockVacancyActivity, id uint) {
				r.EXPECT().DeleteUserApply("test@gmail.com", id).Return(nil)
			},
			sessionRepBehavior: func(sessionRep *mock_session.MockRepository, token string) {
				sessionRep.EXPECT().GetSession(token).Return("test@gmail.com", nil)
			},
			expectedStatusCode:   200,
			expectedResponseBody: "",
		},
		{
			name:           "user not found",
			emailFromToken: "",
			inputId:        42,
			inputParam:     "42",
			mockBehavior: func(r *mock_usecases.MockVacancyActivity, id uint) {
				//createdApply := &models.VacancyActivity{ResumeTitle: "some title"}
				r.EXPECT().DeleteUserApply("test@gmail.com", id).Return(errorHandler.ErrBadRequest)
			},
			sessionRepBehavior: func(sessionRep *mock_session.MockRepository, token string) {
				sessionRep.EXPECT().GetSession(token).Return("", fmt.Errorf("getting session error:"))
			},
			expectedStatusCode:   401,
			expectedResponseBody: "{\"descriptors\":\"\",\"error\":\"Клиент не авторизован\"}",
		},
		{
			name:           "vacancy not found",
			emailFromToken: "test@gmail.com",
			inputId:        42,
			inputParam:     "42",
			mockBehavior: func(r *mock_usecases.MockVacancyActivity, id uint) {
				r.EXPECT().DeleteUserApply("test@gmail.com", id).Return(errorHandler.ErrCannotDeleteVacancyApply)
			},
			sessionRepBehavior: func(sessionRep *mock_session.MockRepository, token string) {
				sessionRep.EXPECT().GetSession(token).Return("test@gmail.com", nil)
			},
			expectedStatusCode:   503,
			expectedResponseBody: "{\"descriptors\":\"\",\"error\":\"Невозможно удалить отклик на вакансию\"}",
		},
	}
	for _, test := range testTable {
		testCase := test
		t.Run(testCase.name, func(t *testing.T) {
			t.Parallel()
			c := gomock.NewController(t)
			defer c.Finish()

			mockUseCase := mock_usecases.NewMockVacancyActivity(c)
			sessionRep := mock_session.NewMockRepository(c)
			sessionMiddlware := middleware.NewSessionMiddleware(sessionRep)

			if testCase.emailFromToken != "" {
				testCase.mockBehavior(mockUseCase, testCase.inputId)
			}

			testCase.sessionRepBehavior(sessionRep, testCase.inputToken)

			handler := VacancyActivityHandler{
				vacancyActivityUseCase: mockUseCase,
			}

			r := gin.New()
			r.POST("/:id", sessionMiddlware.Session, handler.DeleteUserApply, errorHandler.Middleware())

			cookie := &http.Cookie{
				Name:  "session",
				Value: testCase.inputToken,
			}

			w := httptest.NewRecorder()
			req := httptest.NewRequest("POST", "/"+testCase.inputParam,
				bytes.NewBufferString(""))

			req.AddCookie(cookie)
			r.ServeHTTP(w, req)

			assert.Equal(t, testCase.expectedStatusCode, w.Code)
			assert.Equal(t, testCase.expectedResponseBody, w.Body.String())
		})
	}
}

func TestVacancyActivityHandler_GetAllUserApplies(t *testing.T) {
	type mockBehavior func(r *mock_usecases.MockVacancyActivity, id uint)
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
			mockBehavior: func(r *mock_usecases.MockVacancyActivity, id uint) {
				expectedApply := []*models.VacancyActivity{
					{
						UserAccountId: 42,
						ResumeTitle:   "some vacancy",
					},
				}
				r.EXPECT().GetAllUserApplies(id).Return(expectedApply, nil)
			},
			sessionRepBehavior: func(sessionRep *mock_session.MockRepository, token string) {
				sessionRep.EXPECT().GetSession(token).Return("test@gmail.com", nil)
			},
			expectedStatusCode:   200,
			expectedResponseBody: "{\"data\":[{\"user_account_id\":42,\"id\":0,\"vacancy_id\":0,\"applicant_name\":\"\",\"applicant_surname\":\"\",\"title\":\"some vacancy\",\"image\":\"\",\"created_date\":\"0001-01-01T00:00:00Z\"}]}",
		},
		{
			name:           "user not found",
			inputId:        42,
			requestParam:   "42",
			emailFromToken: "test@gmail.com",
			mockBehavior: func(r *mock_usecases.MockVacancyActivity, id uint) {
				expectedVacancy := []*models.VacancyActivity{
					{
						UserAccountId: 42,
						ResumeTitle:   "some vacancy",
					},
				}
				r.EXPECT().GetAllUserApplies(id).Return(expectedVacancy, errorHandler.ErrBadRequest)
			},
			sessionRepBehavior: func(sessionRep *mock_session.MockRepository, token string) {
				sessionRep.EXPECT().GetSession(token).Return("", fmt.Errorf("getting session error:"))
			},
			expectedStatusCode:   400,
			expectedResponseBody: "{\"descriptors\":\"\",\"error\":\"Некорректный запрос\"}",
		},
	}
	for _, test := range testTable {
		testCase := test
		t.Run(testCase.name, func(t *testing.T) {
			t.Parallel()
			c := gomock.NewController(t)
			defer c.Finish()

			mockUseCase := mock_usecases.NewMockVacancyActivity(c)
			sessionRep := mock_session.NewMockRepository(c)
			sessionMiddlware := middleware.NewSessionMiddleware(sessionRep)

			if testCase.emailFromToken != "" {
				testCase.mockBehavior(mockUseCase, testCase.inputId)
			}

			testCase.sessionRepBehavior(sessionRep, testCase.inputToken)

			handler := VacancyActivityHandler{
				vacancyActivityUseCase: mockUseCase,
			}

			r := gin.New()
			r.GET("/:id", sessionMiddlware.Session, handler.GetAllUserApplies, errorHandler.Middleware())

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

func TestVacancyActivityHandler_GetAllVacancyApplies(t *testing.T) {
	type mockBehavior func(r *mock_usecases.MockVacancyActivity, id uint)
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
			mockBehavior: func(r *mock_usecases.MockVacancyActivity, id uint) {
				expectedApply := []*models.VacancyActivity{
					{
						UserAccountId: 42,
						ResumeTitle:   "some vacancy",
					},
				}
				r.EXPECT().GetAllVacancyApplies(id).Return(expectedApply, nil)
			},
			sessionRepBehavior: func(sessionRep *mock_session.MockRepository, token string) {
				sessionRep.EXPECT().GetSession(token).Return("test@gmail.com", nil)
			},
			expectedStatusCode:   200,
			expectedResponseBody: "{\"data\":[{\"user_account_id\":42,\"id\":0,\"vacancy_id\":0,\"applicant_name\":\"\",\"applicant_surname\":\"\",\"title\":\"some vacancy\",\"image\":\"\",\"created_date\":\"0001-01-01T00:00:00Z\"}]}",
		},
		{
			name:           "user not found",
			inputId:        42,
			requestParam:   "42",
			emailFromToken: "test@gmail.com",
			mockBehavior: func(r *mock_usecases.MockVacancyActivity, id uint) {
				expectedVacancy := []*models.VacancyActivity{
					{
						UserAccountId: 42,
						ResumeTitle:   "some vacancy",
					},
				}
				r.EXPECT().GetAllVacancyApplies(id).Return(expectedVacancy, errorHandler.ErrBadRequest)
			},
			sessionRepBehavior: func(sessionRep *mock_session.MockRepository, token string) {
				sessionRep.EXPECT().GetSession(token).Return("", fmt.Errorf("getting session error:"))
			},
			expectedStatusCode:   400,
			expectedResponseBody: "{\"descriptors\":\"\",\"error\":\"Некорректный запрос\"}",
		},
	}
	for _, test := range testTable {
		testCase := test
		t.Run(testCase.name, func(t *testing.T) {
			t.Parallel()
			c := gomock.NewController(t)
			defer c.Finish()

			mockUseCase := mock_usecases.NewMockVacancyActivity(c)
			sessionRep := mock_session.NewMockRepository(c)
			sessionMiddlware := middleware.NewSessionMiddleware(sessionRep)

			if testCase.emailFromToken != "" {
				testCase.mockBehavior(mockUseCase, testCase.inputId)
			}

			testCase.sessionRepBehavior(sessionRep, testCase.inputToken)

			handler := VacancyActivityHandler{
				vacancyActivityUseCase: mockUseCase,
			}

			r := gin.New()
			r.GET("/:id", sessionMiddlware.Session, handler.GetAllVacancyApplies, errorHandler.Middleware())

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
