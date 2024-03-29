package impl

import (
	"HeadHunter/common/session/mocks"
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

func TestVacancyHandler_GetVacancyById(t *testing.T) {
	type mockBehavior func(r *mock_usecases.MockVacancy, id uint)
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
			mockBehavior: func(r *mock_usecases.MockVacancy, id uint) {
				expectedVacancy := &models.Vacancy{
					ID:    42,
					Title: "some vacancy",
				}
				r.EXPECT().GetById(id).Return(expectedVacancy, nil)
			},
			//sessionRepBehavior: func(sessionRep *mock_session.MockRepository, token string) {
			//	sessionRep.EXPECT().GetSession(token).Return("test@gmail.com", nil)
			//},
			expectedStatusCode:   200,
			expectedResponseBody: "{\"id\":42,\"postedByUserId\":0,\"title\":\"some vacancy\",\"createdDate\":\"0001-01-01T00:00:00Z\",\"vacancyActivities\":null,\"skills\":null}",
		},
		{
			name:           "not found",
			inputId:        42,
			requestParam:   "42",
			emailFromToken: "test@gmail.com",
			mockBehavior: func(r *mock_usecases.MockVacancy, id uint) {
				expectedVacancy := &models.Vacancy{
					ID:    42,
					Title: "some vacancy",
				}
				r.EXPECT().GetById(id).Return(expectedVacancy, errorHandler.ErrVacancyNotFound)
			},
			//sessionRepBehavior: func(sessionRep *mock_session.MockRepository, token string) {
			//	sessionRep.EXPECT().GetSession(token).Return("test@gmail.com", nil)
			//},
			expectedStatusCode:   404,
			expectedResponseBody: "{\"descriptors\":\"\",\"error\":\"Вакансия не найдена\"}",
		},
	}
	for _, test := range testTable {
		testCase := test
		t.Run(testCase.name, func(t *testing.T) {
			t.Parallel()
			c := gomock.NewController(t)
			defer c.Finish()

			mockUseCase := mock_usecases.NewMockVacancy(c)
			//sessionRep := mock_session.NewMockRepository(c)
			//sessionMiddlware := middleware.NewSessionMiddleware(sessionRep)

			if testCase.emailFromToken != "" {
				testCase.mockBehavior(mockUseCase, testCase.inputId)
			}

			//testCase.sessionRepBehavior(sessionRep, testCase.inputToken)

			handler := VacancyHandler{
				vacancyUseCase: mockUseCase,
			}

			r := gin.New()
			r.GET("/:id", handler.GetVacancyById, errorHandler.Middleware())

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

func TestVacancyHandler_CreateVacancy(t *testing.T) {
	type mockBehavior func(r *mock_usecases.MockVacancy, vacancy *models.Vacancy)
	type sessionRepBehavior func(r *mock_session.MockRepository, token string)
	testTable := []struct {
		name                 string
		inputToken           string
		emailFromToken       string
		createdVacancy       *models.Vacancy
		inputBody            string
		mockBehavior         mockBehavior
		sessionRepBehavior   sessionRepBehavior
		expectedStatusCode   int
		expectedResponseBody string
	}{
		{
			name:           "valid",
			emailFromToken: "test@gmail.com",
			createdVacancy: &models.Vacancy{
				Title: "some title",
			},
			inputBody: `{
				"title": "some title"
}`,
			mockBehavior: func(r *mock_usecases.MockVacancy, vacancy *models.Vacancy) {
				createdVacancy := &models.Vacancy{Title: "some title"}
				r.EXPECT().Create("test@gmail.com", createdVacancy).Return(uint(1), nil)
			},
			sessionRepBehavior: func(sessionRep *mock_session.MockRepository, token string) {
				sessionRep.EXPECT().GetSession(token).Return("test@gmail.com", nil)
			},
			expectedStatusCode:   200,
			expectedResponseBody: "{\"id\":1}",
		},
		{
			name:           "user not found",
			emailFromToken: "",
			createdVacancy: &models.Vacancy{
				Title: "some title",
			},
			inputBody: `{
    			"title": "some title"
}`,
			mockBehavior: func(r *mock_usecases.MockVacancy, vacancy *models.Vacancy) {
				createdVacancy := &models.Vacancy{Title: "some title"}
				r.EXPECT().Create("test@gmail.com", createdVacancy).Return(errorHandler.ErrBadRequest)
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

			mockUseCase := mock_usecases.NewMockVacancy(c)
			sessionRep := mock_session.NewMockRepository(c)
			sessionMiddlware := middleware.NewSessionMiddleware(sessionRep)

			if testCase.emailFromToken != "" {
				testCase.mockBehavior(mockUseCase, testCase.createdVacancy)
			}

			testCase.sessionRepBehavior(sessionRep, testCase.inputToken)

			handler := VacancyHandler{
				vacancyUseCase: mockUseCase,
			}

			r := gin.New()
			r.POST("/", sessionMiddlware.Session, handler.CreateVacancy, errorHandler.Middleware())

			cookie := &http.Cookie{
				Name:  "session",
				Value: testCase.inputToken,
			}

			w := httptest.NewRecorder()
			req := httptest.NewRequest("POST", "/",
				bytes.NewBufferString(test.inputBody))

			req.AddCookie(cookie)
			r.ServeHTTP(w, req)

			assert.Equal(t, testCase.expectedStatusCode, w.Code)
			assert.Equal(t, testCase.expectedResponseBody, w.Body.String())
		})
	}
}

func TestVacancyHandler_DeleteVacancy(t *testing.T) {
	type mockBehavior func(r *mock_usecases.MockVacancy, id uint)
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
			mockBehavior: func(r *mock_usecases.MockVacancy, id uint) {
				r.EXPECT().Delete("test@gmail.com", id).Return(nil)
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
			mockBehavior: func(r *mock_usecases.MockVacancy, id uint) {
				createdVacancy := &models.Vacancy{Title: "some title"}
				r.EXPECT().Delete("test@gmail.com", createdVacancy).Return(errorHandler.ErrBadRequest)
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
			mockBehavior: func(r *mock_usecases.MockVacancy, id uint) {
				r.EXPECT().Delete("test@gmail.com", id).Return(errorHandler.ErrVacancyNotFound)
			},
			sessionRepBehavior: func(sessionRep *mock_session.MockRepository, token string) {
				sessionRep.EXPECT().GetSession(token).Return("test@gmail.com", nil)
			},
			expectedStatusCode:   404,
			expectedResponseBody: "{\"descriptors\":\"\",\"error\":\"Вакансия не найдена\"}",
		},
	}
	for _, test := range testTable {
		testCase := test
		t.Run(testCase.name, func(t *testing.T) {
			t.Parallel()
			c := gomock.NewController(t)
			defer c.Finish()

			mockUseCase := mock_usecases.NewMockVacancy(c)
			sessionRep := mock_session.NewMockRepository(c)
			sessionMiddlware := middleware.NewSessionMiddleware(sessionRep)

			if testCase.emailFromToken != "" {
				testCase.mockBehavior(mockUseCase, testCase.inputId)
			}

			testCase.sessionRepBehavior(sessionRep, testCase.inputToken)

			handler := VacancyHandler{
				vacancyUseCase: mockUseCase,
			}

			r := gin.New()
			r.POST("/:id", sessionMiddlware.Session, handler.DeleteVacancy, errorHandler.Middleware())

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

func TestVacancyHandler_UpdateVacancy(t *testing.T) {
	type mockBehavior func(r *mock_usecases.MockVacancy, id uint, vacancy *models.Vacancy)
	type sessionRepBehavior func(r *mock_session.MockRepository, token string)

	testTable := []struct {
		name                 string
		inputToken           string
		inputId              uint
		inputParam           string
		emailFromToken       string
		createdVacancy       *models.Vacancy
		inputBody            string
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
			createdVacancy: &models.Vacancy{
				ID:    42,
				Title: "some title",
			},
			inputBody: `{
    			"title": "some title"
}`,
			mockBehavior: func(r *mock_usecases.MockVacancy, id uint, vacancy *models.Vacancy) {
				updated := &models.Vacancy{Title: "some title"}
				r.EXPECT().Update("test@gmail.com", id, updated).Return(nil)
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
			createdVacancy: &models.Vacancy{
				ID:    42,
				Title: "some title",
			},
			inputBody: `{
    			"title": "some title"
}`,
			mockBehavior: func(r *mock_usecases.MockVacancy, id uint, vacancy *models.Vacancy) {
				updated := &models.Vacancy{Title: "some title"}
				r.EXPECT().Update("test@gmail.com", id, updated).Return(errorHandler.ErrBadRequest)
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
			createdVacancy: &models.Vacancy{
				ID:    42,
				Title: "some title",
			},
			inputBody: `{
    			"title": "some title"
}`,
			mockBehavior: func(r *mock_usecases.MockVacancy, id uint, vacancy *models.Vacancy) {
				updated := &models.Vacancy{Title: "some title"}
				r.EXPECT().Update("test@gmail.com", id, updated).Return(errorHandler.ErrBadRequest)
			},
			sessionRepBehavior: func(sessionRep *mock_session.MockRepository, token string) {
				sessionRep.EXPECT().GetSession(token).Return("test@gmail.com", nil)
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

			mockUseCase := mock_usecases.NewMockVacancy(c)
			sessionRep := mock_session.NewMockRepository(c)
			sessionMiddlware := middleware.NewSessionMiddleware(sessionRep)

			if testCase.emailFromToken != "" {
				testCase.mockBehavior(mockUseCase, testCase.inputId, testCase.createdVacancy)
			}

			testCase.sessionRepBehavior(sessionRep, testCase.inputToken)

			handler := VacancyHandler{
				vacancyUseCase: mockUseCase,
			}

			r := gin.New()
			r.GET("/:id", sessionMiddlware.Session, handler.UpdateVacancy, errorHandler.Middleware())

			cookie := &http.Cookie{
				Name:  "session",
				Value: testCase.inputToken,
			}

			w := httptest.NewRecorder()
			req := httptest.NewRequest("GET", "/"+testCase.inputParam,
				bytes.NewBufferString(test.inputBody))

			req.AddCookie(cookie)
			r.ServeHTTP(w, req)

			assert.Equal(t, testCase.expectedStatusCode, w.Code)
			assert.Equal(t, testCase.expectedResponseBody, w.Body.String())
		})
	}
}

func TestVacancyHandler_GetUserVacancies(t *testing.T) {
	type mockBehavior func(r *mock_usecases.MockVacancy, id uint)
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
			mockBehavior: func(r *mock_usecases.MockVacancy, id uint) {
				expectedVacancy := []*models.Vacancy{
					{
						ID:    42,
						Title: "some vacancy",
					},
				}
				r.EXPECT().GetByUserId(id).Return(expectedVacancy, nil)
			},
			sessionRepBehavior: func(sessionRep *mock_session.MockRepository, token string) {
				sessionRep.EXPECT().GetSession(token).Return("test@gmail.com", nil)
			},
			expectedStatusCode:   200,
			expectedResponseBody: "{\"data\":[{\"id\":42,\"postedByUserId\":0,\"title\":\"some vacancy\",\"createdDate\":\"0001-01-01T00:00:00Z\",\"vacancyActivities\":null,\"skills\":null}]}",
		},
		{
			name:           "user not found",
			inputId:        42,
			requestParam:   "42",
			emailFromToken: "test@gmail.com",
			mockBehavior: func(r *mock_usecases.MockVacancy, id uint) {
				expectedVacancy := []*models.Vacancy{
					{
						ID:    42,
						Title: "some vacancy",
					},
				}
				r.EXPECT().GetByUserId(id).Return(expectedVacancy, errorHandler.ErrBadRequest)
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

			mockUseCase := mock_usecases.NewMockVacancy(c)
			sessionRep := mock_session.NewMockRepository(c)
			sessionMiddlware := middleware.NewSessionMiddleware(sessionRep)

			if testCase.emailFromToken != "" {
				testCase.mockBehavior(mockUseCase, testCase.inputId)
			}

			testCase.sessionRepBehavior(sessionRep, testCase.inputToken)

			handler := VacancyHandler{
				vacancyUseCase: mockUseCase,
			}

			r := gin.New()
			r.GET("/:id", sessionMiddlware.Session, handler.GetUserVacancies, errorHandler.Middleware())

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

func TestVacancyHandler_GetAllVacancies(t *testing.T) {
	type mockBehavior func(r *mock_usecases.MockVacancy, filters models.VacancyFilter)
	//type sessionRepBehavior func(r *mock_session.MockRepository, token string)

	testTable := []struct {
		name                 string
		inputToken           string
		requestParam         string
		emailFromToken       string
		mockBehavior         mockBehavior
		expectedStatusCode   int
		expectedResponseBody string
		filters              models.VacancyFilter
	}{
		{
			name:           "valid",
			requestParam:   "42",
			emailFromToken: "test@gmail.com",
			mockBehavior: func(r *mock_usecases.MockVacancy, filters models.VacancyFilter) {
				expectedVacancy := []*models.Vacancy{
					{
						ID:    42,
						Title: "some vacancy",
					},
				}
				r.EXPECT().GetAll(filters).Return(expectedVacancy, nil)
			},
			expectedStatusCode:   200,
			expectedResponseBody: "{\"data\":[{\"id\":42,\"postedByUserId\":0,\"title\":\"some vacancy\",\"createdDate\":\"0001-01-01T00:00:00Z\",\"vacancyActivities\":null,\"skills\":null}]}",
		},
		{
			name:           "user not found",
			requestParam:   "42",
			emailFromToken: "test@gmail.com",
			mockBehavior: func(r *mock_usecases.MockVacancy, filters models.VacancyFilter) {
				expectedVacancy := []*models.Vacancy{
					{
						ID:    42,
						Title: "some vacancy",
					},
				}
				r.EXPECT().GetAll(filters).Return(expectedVacancy, errorHandler.ErrBadRequest)
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

			mockUseCase := mock_usecases.NewMockVacancy(c)

			if testCase.emailFromToken != "" {
				testCase.mockBehavior(mockUseCase, testCase.filters)
			}

			handler := VacancyHandler{
				vacancyUseCase: mockUseCase,
			}

			r := gin.New()
			r.GET("/", handler.GetAllVacancies, errorHandler.Middleware())

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

func TestVacancyHandler_GetPreviewVacanciesByEmployer(t *testing.T) {
	type mockBehavior func(r *mock_usecases.MockVacancy, id uint)
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
			mockBehavior: func(r *mock_usecases.MockVacancy, id uint) {
				expectedVacancy := []*models.VacancyPreview{
					{
						Id:    42,
						Title: "some vacancy",
					},
				}
				r.EXPECT().GetPreviewVacanciesByEmployer(id).Return(expectedVacancy, nil)
			},
			sessionRepBehavior: func(sessionRep *mock_session.MockRepository, token string) {
				sessionRep.EXPECT().GetSession(token).Return("test@gmail.com", nil)
			},
			expectedStatusCode:   200,
			expectedResponseBody: "[{\"id\":42,\"title\":\"some vacancy\",\"image\":\"\",\"salary\":0,\"location\":\"\",\"format\":\"\",\"hours\":\"\",\"description\":\"\"}]",
		},
		{
			name:           "user not found",
			inputId:        42,
			requestParam:   "42",
			emailFromToken: "test@gmail.com",
			mockBehavior: func(r *mock_usecases.MockVacancy, id uint) {
				expectedVacancy := []*models.VacancyPreview{
					{
						Id:    42,
						Title: "some vacancy",
					},
				}
				r.EXPECT().GetPreviewVacanciesByEmployer(id).Return(expectedVacancy, errorHandler.ErrBadRequest)
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

			mockUseCase := mock_usecases.NewMockVacancy(c)
			sessionRep := mock_session.NewMockRepository(c)
			sessionMiddlware := middleware.NewSessionMiddleware(sessionRep)

			if testCase.emailFromToken != "" {
				testCase.mockBehavior(mockUseCase, testCase.inputId)
			}

			testCase.sessionRepBehavior(sessionRep, testCase.inputToken)

			handler := VacancyHandler{
				vacancyUseCase: mockUseCase,
			}

			r := gin.New()
			r.GET("/:id", sessionMiddlware.Session, handler.GetPreviewVacanciesByEmployer, errorHandler.Middleware())

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

func TestVacancyHandler_AddVacancyToFavorites(t *testing.T) {
	type mockBehavior func(r *mock_usecases.MockVacancy, vacancyId uint)
	type sessionRepBehavior func(r *mock_session.MockRepository, token string)

	testTable := []struct {
		name                 string
		vacancyId            uint
		inputToken           string
		requestParam         string
		inputBody            string
		emailFromToken       string
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
			mockBehavior: func(r *mock_usecases.MockVacancy, vacancyId uint) {
				r.EXPECT().AddVacancyToFavorites("test@gmail.com", vacancyId).Return(nil)
			},
			sessionRepBehavior: func(sessionRep *mock_session.MockRepository, token string) {
				sessionRep.EXPECT().GetSession(token).Return("test@gmail.com", nil)
			},
			expectedStatusCode: 200,
		},
		{
			name:           "user not found",
			emailFromToken: "",
			vacancyId:      1,
			requestParam:   "1",
			mockBehavior: func(r *mock_usecases.MockVacancy, vacancyId uint) {
				r.EXPECT().AddVacancyToFavorites("test@gmail.com", vacancyId).Return(errorHandler.ErrBadRequest)
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

			mockUseCase := mock_usecases.NewMockVacancy(c)
			sessionRep := mock_session.NewMockRepository(c)
			sessionMiddlware := middleware.NewSessionMiddleware(sessionRep)

			if testCase.emailFromToken != "" {
				testCase.mockBehavior(mockUseCase, testCase.vacancyId)
			}

			testCase.sessionRepBehavior(sessionRep, testCase.inputToken)

			handler := VacancyHandler{
				vacancyUseCase: mockUseCase,
			}

			r := gin.New()
			r.POST("/favorites/:id", sessionMiddlware.Session, handler.AddVacancyToFavorites, errorHandler.Middleware())

			cookie := &http.Cookie{
				Name:  "session",
				Value: testCase.inputToken,
			}

			w := httptest.NewRecorder()
			req := httptest.NewRequest("POST", "/favorites/"+testCase.requestParam,
				bytes.NewBufferString(test.inputBody))

			req.AddCookie(cookie)
			r.ServeHTTP(w, req)

			assert.Equal(t, testCase.expectedStatusCode, w.Code)
			assert.Equal(t, testCase.expectedResponseBody, w.Body.String())
		})
	}
}

func TestVacancyHandler_GetUserFavoriteVacancies(t *testing.T) {
	type mockBehavior func(r *mock_usecases.MockVacancy)
	type sessionRepBehavior func(r *mock_session.MockRepository, token string)

	testTable := []struct {
		name                 string
		inputToken           string
		emailFromToken       string
		mockBehavior         mockBehavior
		sessionRepBehavior   sessionRepBehavior
		expectedStatusCode   int
		expectedResponseBody string
	}{
		{
			name:           "valid",
			emailFromToken: "test@gmail.com",
			mockBehavior: func(r *mock_usecases.MockVacancy) {
				expectedApply := []*models.Vacancy{
					{
						Title:       "title",
						Description: "description",
					},
				}
				r.EXPECT().GetUserFavoriteVacancies("test@gmail.com").Return(expectedApply, nil)
			},
			sessionRepBehavior: func(sessionRep *mock_session.MockRepository, token string) {
				sessionRep.EXPECT().GetSession(token).Return("test@gmail.com", nil)
			},
			expectedStatusCode:   200,
			expectedResponseBody: "{\"data\":[{\"id\":0,\"postedByUserId\":0,\"title\":\"title\",\"description\":\"description\",\"createdDate\":\"0001-01-01T00:00:00Z\",\"vacancyActivities\":null,\"skills\":null}]}",
		},
	}
	for _, test := range testTable {
		testCase := test
		t.Run(testCase.name, func(t *testing.T) {
			t.Parallel()
			c := gomock.NewController(t)
			defer c.Finish()

			mockUseCase := mock_usecases.NewMockVacancy(c)
			sessionRep := mock_session.NewMockRepository(c)
			sessionMiddlware := middleware.NewSessionMiddleware(sessionRep)

			if testCase.emailFromToken != "" {
				testCase.mockBehavior(mockUseCase)
			}

			testCase.sessionRepBehavior(sessionRep, testCase.inputToken)

			handler := VacancyHandler{
				vacancyUseCase: mockUseCase,
			}

			r := gin.New()
			r.GET("/favorites", sessionMiddlware.Session, handler.GetUserFavoriteVacancies, errorHandler.Middleware())

			cookie := &http.Cookie{
				Name:  "session",
				Value: testCase.inputToken,
			}

			w := httptest.NewRecorder()
			req := httptest.NewRequest("GET", "/favorites",
				bytes.NewBufferString(""))

			req.AddCookie(cookie)
			r.ServeHTTP(w, req)

			assert.Equal(t, testCase.expectedStatusCode, w.Code)
			assert.Equal(t, testCase.expectedResponseBody, w.Body.String())
		})
	}
}

func TestVacancyHandler_DeleteVacancyFromFavorites(t *testing.T) {
	type mockBehavior func(r *mock_usecases.MockVacancy, id uint)
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
			mockBehavior: func(r *mock_usecases.MockVacancy, id uint) {
				r.EXPECT().DeleteVacancyFromFavorites("test@gmail.com", id).Return(nil)
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
			requestParam:   "42",
			mockBehavior: func(r *mock_usecases.MockVacancy, id uint) {
				//createdApply := &models.VacancyActivity{ResumeTitle: "some title"}
				r.EXPECT().DeleteVacancyFromFavorites("test@gmail.com", id).Return(errorHandler.ErrBadRequest)
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

			mockUseCase := mock_usecases.NewMockVacancy(c)
			sessionRep := mock_session.NewMockRepository(c)
			sessionMiddlware := middleware.NewSessionMiddleware(sessionRep)

			if testCase.emailFromToken != "" {
				testCase.mockBehavior(mockUseCase, testCase.inputId)
			}

			testCase.sessionRepBehavior(sessionRep, testCase.inputToken)

			handler := VacancyHandler{
				vacancyUseCase: mockUseCase,
			}

			r := gin.New()
			r.DELETE("/:id", sessionMiddlware.Session, handler.DeleteVacancyFromFavorites, errorHandler.Middleware())

			cookie := &http.Cookie{
				Name:  "session",
				Value: testCase.inputToken,
			}

			w := httptest.NewRecorder()
			req := httptest.NewRequest("DELETE", "/"+testCase.requestParam,
				bytes.NewBufferString(""))

			req.AddCookie(cookie)
			r.ServeHTTP(w, req)

			assert.Equal(t, testCase.expectedStatusCode, w.Code)
			assert.Equal(t, testCase.expectedResponseBody, w.Body.String())
		})
	}
}
