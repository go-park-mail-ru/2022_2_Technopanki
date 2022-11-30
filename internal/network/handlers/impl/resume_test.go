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

func TestResumeHandler_GetResume(t *testing.T) {
	type mockBehavior func(r *mock_usecases.MockResume, id uint)
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
			mockBehavior: func(r *mock_usecases.MockResume, id uint) {
				expectedResume := &models.Resume{
					ID:    42,
					Title: "some resume",
				}
				r.EXPECT().GetResume(id).Return(expectedResume, nil)
			},
			sessionRepBehavior: func(sessionRep *mock_session.MockRepository, token string) {
				sessionRep.EXPECT().GetSession(token).Return("test@gmail.com", nil)
			},
			expectedStatusCode:   200,
			expectedResponseBody: "{\"id\":42,\"user_account_id\":0,\"title\":\"some resume\",\"description\":\"\",\"created_date\":\"0001-01-01T00:00:00Z\",\"education_detail\":{\"resume_id\":0,\"certificate_degree_name\":\"\",\"major\":\"\",\"university_name\":\"\",\"starting_date\":\"0001-01-01T00:00:00Z\",\"completion_date\":\"0001-01-01T00:00:00Z\"},\"experience_detail\":{\"resume_id\":0,\"is_current_job\":\"\",\"start_date\":\"0001-01-01T00:00:00Z\",\"end_date\":\"0001-01-01T00:00:00Z\",\"job_title\":\"\",\"company_name\":\"\",\"job_location_city\":\"\",\"description\":\"\"},\"applicant_skills\":null}",
		},
		{
			name:           "not found",
			inputId:        42,
			requestParam:   "42",
			emailFromToken: "test@gmail.com",
			mockBehavior: func(r *mock_usecases.MockResume, id uint) {
				expectedResume := &models.Resume{
					ID:    42,
					Title: "some resume",
				}
				r.EXPECT().GetResume(id).Return(expectedResume, errorHandler.ErrResumeNotFound)
			},
			sessionRepBehavior: func(sessionRep *mock_session.MockRepository, token string) {
				sessionRep.EXPECT().GetSession(token).Return("test@gmail.com", nil)
			},
			expectedStatusCode:   404,
			expectedResponseBody: "{\"descriptors\":\"\",\"error\":\"Резюме не найдено\"}",
		},

		{
			name:           "resume not found",
			emailFromToken: "test@gmail.com",
			inputId:        42,
			requestParam:   "42",
			mockBehavior: func(r *mock_usecases.MockResume, id uint) {
				r.EXPECT().GetResume(id).Return(&models.Resume{}, errorHandler.ErrResumeNotFound)
			},
			sessionRepBehavior: func(sessionRep *mock_session.MockRepository, token string) {
				sessionRep.EXPECT().GetSession(token).Return("test@gmail.com", nil)
			},
			expectedStatusCode:   404,
			expectedResponseBody: "{\"descriptors\":\"\",\"error\":\"Резюме не найдено\"}",
		},
		{
			name:           "user is not author",
			emailFromToken: "test@gmail.com",
			inputId:        42,
			requestParam:   "42",
			mockBehavior: func(r *mock_usecases.MockResume, id uint) {
				r.EXPECT().GetResume(id).Return(&models.Resume{}, errorHandler.ErrUnauthorized)
			},
			sessionRepBehavior: func(sessionRep *mock_session.MockRepository, token string) {
				sessionRep.EXPECT().GetSession(token).Return("test@gmail.com", nil)
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

			mockUseCase := mock_usecases.NewMockResume(c)
			sessionRep := mock_session.NewMockRepository(c)
			sessionMiddlware := middleware.NewSessionMiddleware(sessionRep)

			if testCase.emailFromToken != "" {
				testCase.mockBehavior(mockUseCase, testCase.inputId)
			}

			testCase.sessionRepBehavior(sessionRep, testCase.inputToken)

			handler := ResumeHandler{
				resumeUseCase: mockUseCase,
			}

			r := gin.New()
			r.GET("/:id", sessionMiddlware.Session, handler.GetResume, errorHandler.Middleware())

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

func TestResumeHandler_GetResumeByApplicant(t *testing.T) {
	type mockBehavior func(r *mock_usecases.MockResume, id uint)
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
			mockBehavior: func(r *mock_usecases.MockResume, id uint) {
				expectedResume := []*models.Resume{
					{
						ID:    42,
						Title: "some resume",
					},
				}
				r.EXPECT().GetResumeByApplicant(id).Return(expectedResume, nil)
			},
			sessionRepBehavior: func(sessionRep *mock_session.MockRepository, token string) {
				sessionRep.EXPECT().GetSession(token).Return("test@gmail.com", nil)
			},
			expectedStatusCode:   200,
			expectedResponseBody: "[{\"id\":42,\"user_account_id\":0,\"title\":\"some resume\",\"description\":\"\",\"created_date\":\"0001-01-01T00:00:00Z\",\"education_detail\":{\"resume_id\":0,\"certificate_degree_name\":\"\",\"major\":\"\",\"university_name\":\"\",\"starting_date\":\"0001-01-01T00:00:00Z\",\"completion_date\":\"0001-01-01T00:00:00Z\"},\"experience_detail\":{\"resume_id\":0,\"is_current_job\":\"\",\"start_date\":\"0001-01-01T00:00:00Z\",\"end_date\":\"0001-01-01T00:00:00Z\",\"job_title\":\"\",\"company_name\":\"\",\"job_location_city\":\"\",\"description\":\"\"},\"applicant_skills\":null}]",
		},
		{
			name:           "resume not found",
			emailFromToken: "test@gmail.com",
			inputId:        42,
			requestParam:   "42",
			mockBehavior: func(r *mock_usecases.MockResume, id uint) {
				r.EXPECT().GetResumeByApplicant(id).Return([]*models.Resume{}, errorHandler.ErrResumeNotFound)
			},
			sessionRepBehavior: func(sessionRep *mock_session.MockRepository, token string) {
				sessionRep.EXPECT().GetSession(token).Return("test@gmail.com", nil)
			},
			expectedStatusCode:   404,
			expectedResponseBody: "{\"descriptors\":\"\",\"error\":\"Резюме не найдено\"}",
		},
		{
			name:           "user is not author",
			emailFromToken: "test@gmail.com",
			inputId:        42,
			requestParam:   "42",
			mockBehavior: func(r *mock_usecases.MockResume, id uint) {
				r.EXPECT().GetResumeByApplicant(id).Return([]*models.Resume{}, errorHandler.ErrUnauthorized)
			},
			sessionRepBehavior: func(sessionRep *mock_session.MockRepository, token string) {
				sessionRep.EXPECT().GetSession(token).Return("test@gmail.com", nil)
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

			mockUseCase := mock_usecases.NewMockResume(c)
			sessionRep := mock_session.NewMockRepository(c)
			sessionMiddlware := middleware.NewSessionMiddleware(sessionRep)

			if testCase.emailFromToken != "" {
				testCase.mockBehavior(mockUseCase, testCase.inputId)
			}

			testCase.sessionRepBehavior(sessionRep, testCase.inputToken)

			handler := ResumeHandler{
				resumeUseCase: mockUseCase,
			}

			r := gin.New()
			r.GET("/:user_id", sessionMiddlware.Session, handler.GetResumeByApplicant, errorHandler.Middleware())

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

func TestResumeHandler_GetPreviewResumeByApplicant(t *testing.T) {
	type mockBehavior func(r *mock_usecases.MockResume, id uint)
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
			mockBehavior: func(r *mock_usecases.MockResume, id uint) {
				expectedResume := []*models.ResumePreview{
					{
						Id:    42,
						Title: "some resume",
					},
				}
				r.EXPECT().GetPreviewResumeByApplicant(id).Return(expectedResume, nil)
			},
			sessionRepBehavior: func(sessionRep *mock_session.MockRepository, token string) {
				sessionRep.EXPECT().GetSession(token).Return("test@gmail.com", nil)
			},
			expectedStatusCode:   200,
			expectedResponseBody: "[{\"image\":\"\",\"applicant_name\":\"\",\"applicant_surname\":\"\",\"id\":42,\"title\":\"some resume\",\"created_date\":\"0001-01-01T00:00:00Z\"}]",
		},
		{
			name:           "resume not found",
			emailFromToken: "test@gmail.com",
			inputId:        42,
			requestParam:   "42",
			mockBehavior: func(r *mock_usecases.MockResume, id uint) {
				r.EXPECT().GetPreviewResumeByApplicant(id).Return([]*models.ResumePreview{}, errorHandler.ErrResumeNotFound)
			},
			sessionRepBehavior: func(sessionRep *mock_session.MockRepository, token string) {
				sessionRep.EXPECT().GetSession(token).Return("test@gmail.com", nil)
			},
			expectedStatusCode:   404,
			expectedResponseBody: "{\"descriptors\":\"\",\"error\":\"Резюме не найдено\"}",
		},
		{
			name:           "user is not author",
			emailFromToken: "test@gmail.com",
			inputId:        42,
			requestParam:   "42",
			mockBehavior: func(r *mock_usecases.MockResume, id uint) {
				r.EXPECT().GetPreviewResumeByApplicant(id).Return([]*models.ResumePreview{}, errorHandler.ErrUnauthorized)
			},
			sessionRepBehavior: func(sessionRep *mock_session.MockRepository, token string) {
				sessionRep.EXPECT().GetSession(token).Return("test@gmail.com", nil)
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

			mockUseCase := mock_usecases.NewMockResume(c)
			sessionRep := mock_session.NewMockRepository(c)
			sessionMiddlware := middleware.NewSessionMiddleware(sessionRep)

			if testCase.emailFromToken != "" {
				testCase.mockBehavior(mockUseCase, testCase.inputId)
			}

			testCase.sessionRepBehavior(sessionRep, testCase.inputToken)

			handler := ResumeHandler{
				resumeUseCase: mockUseCase,
			}

			r := gin.New()
			r.GET("/:user_id", sessionMiddlware.Session, handler.GetPreviewResumeByApplicant, errorHandler.Middleware())

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

func TestResumeHandler_CreateResume(t *testing.T) {
	type mockBehavior func(r *mock_usecases.MockResume, resume *models.Resume)
	type sessionRepBehavior func(r *mock_session.MockRepository, token string)

	testTable := []struct {
		name                 string
		inputToken           string
		emailFromToken       string
		createdResume        *models.Resume
		inputBody            string
		mockBehavior         mockBehavior
		sessionRepBehavior   sessionRepBehavior
		expectedStatusCode   int
		expectedResponseBody string
	}{
		{
			name:           "valid",
			emailFromToken: "test@gmail.com",
			createdResume: &models.Resume{
				Title: "some title",
			},
			inputBody: `{
    			"title": "some title"
}`,
			mockBehavior: func(r *mock_usecases.MockResume, resume *models.Resume) {
				createdResume := &models.Resume{Title: "some title"}
				r.EXPECT().CreateResume(createdResume, "test@gmail.com").Return(nil)
			},
			sessionRepBehavior: func(sessionRep *mock_session.MockRepository, token string) {
				sessionRep.EXPECT().GetSession(token).Return("test@gmail.com", nil)
			},
			expectedStatusCode:   200,
			expectedResponseBody: "{\"id\":0}",
		},
		{
			name:           "user not found",
			emailFromToken: "",
			createdResume: &models.Resume{
				Title: "some title",
			},
			inputBody: `{
    			"title": "some title"
}`,
			mockBehavior: func(r *mock_usecases.MockResume, resume *models.Resume) {
				createdResume := &models.Resume{Title: "some title"}
				r.EXPECT().CreateResume(createdResume, "test@gmail.com").Return(errorHandler.ErrBadRequest)
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

			mockUseCase := mock_usecases.NewMockResume(c)
			sessionRep := mock_session.NewMockRepository(c)
			sessionMiddlware := middleware.NewSessionMiddleware(sessionRep)

			if testCase.emailFromToken != "" {
				testCase.mockBehavior(mockUseCase, testCase.createdResume)
			}

			testCase.sessionRepBehavior(sessionRep, testCase.inputToken)

			handler := ResumeHandler{
				resumeUseCase: mockUseCase,
			}

			r := gin.New()
			r.POST("/", sessionMiddlware.Session, handler.CreateResume, errorHandler.Middleware())

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

func TestResumeHandler_UpdateResume(t *testing.T) {
	type mockBehavior func(r *mock_usecases.MockResume, id uint, resume *models.Resume)
	type sessionRepBehavior func(r *mock_session.MockRepository, token string)

	testTable := []struct {
		name                 string
		inputToken           string
		inputId              uint
		inputParam           string
		emailFromToken       string
		createdResume        *models.Resume
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
			createdResume: &models.Resume{
				ID:    42,
				Title: "some title",
			},
			inputBody: `{
    			"title": "some title"
}`,
			mockBehavior: func(r *mock_usecases.MockResume, id uint, resume *models.Resume) {
				updated := &models.Resume{Title: "some title"}
				r.EXPECT().UpdateResume(id, updated, "test@gmail.com").Return(nil)
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
			createdResume: &models.Resume{
				ID:    42,
				Title: "some title",
			},
			inputBody: `{
    			"title": "some title"
}`,
			mockBehavior: func(r *mock_usecases.MockResume, id uint, resume *models.Resume) {
				updated := &models.Resume{Title: "some title"}
				r.EXPECT().UpdateResume(id, updated, "test@gmail.com").Return(errorHandler.ErrBadRequest)
			},
			sessionRepBehavior: func(sessionRep *mock_session.MockRepository, token string) {
				sessionRep.EXPECT().GetSession(token).Return("", fmt.Errorf("getting session error:"))
			},
			expectedStatusCode:   401,
			expectedResponseBody: "{\"descriptors\":\"\",\"error\":\"Клиент не авторизован\"}",
		},
		{
			name:           "resume not found",
			emailFromToken: "test@gmail.com",
			inputId:        42,
			inputParam:     "42",
			createdResume: &models.Resume{
				ID:    42,
				Title: "some title",
			},
			inputBody: `{
    			"title": "some title"
}`,
			mockBehavior: func(r *mock_usecases.MockResume, id uint, resume *models.Resume) {
				updated := &models.Resume{Title: "some title"}
				r.EXPECT().UpdateResume(id, updated, "test@gmail.com").Return(errorHandler.ErrResumeNotFound)
			},
			sessionRepBehavior: func(sessionRep *mock_session.MockRepository, token string) {
				sessionRep.EXPECT().GetSession(token).Return("test@gmail.com", nil)
			},
			expectedStatusCode:   404,
			expectedResponseBody: "{\"descriptors\":\"\",\"error\":\"Резюме не найдено\"}",
		},
		{
			name:           "user is not an author",
			emailFromToken: "test@gmail.com",
			inputId:        42,
			inputParam:     "42",
			createdResume: &models.Resume{
				ID:    42,
				Title: "some title",
			},
			inputBody: `{
    			"title": "some title"
}`,
			mockBehavior: func(r *mock_usecases.MockResume, id uint, resume *models.Resume) {
				updated := &models.Resume{Title: "some title"}
				r.EXPECT().UpdateResume(id, updated, "test@gmail.com").Return(errorHandler.ErrUnauthorized)
			},
			sessionRepBehavior: func(sessionRep *mock_session.MockRepository, token string) {
				sessionRep.EXPECT().GetSession(token).Return("test@gmail.com", nil)
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

			mockUseCase := mock_usecases.NewMockResume(c)
			sessionRep := mock_session.NewMockRepository(c)
			sessionMiddlware := middleware.NewSessionMiddleware(sessionRep)

			if testCase.emailFromToken != "" {
				testCase.mockBehavior(mockUseCase, testCase.inputId, testCase.createdResume)
			}

			testCase.sessionRepBehavior(sessionRep, testCase.inputToken)

			handler := ResumeHandler{
				resumeUseCase: mockUseCase,
			}

			r := gin.New()
			r.GET("/:id", sessionMiddlware.Session, handler.UpdateResume, errorHandler.Middleware())

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

func TestResumeHandler_DeleteResume(t *testing.T) {
	type mockBehavior func(r *mock_usecases.MockResume, id uint)
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
			mockBehavior: func(r *mock_usecases.MockResume, id uint) {
				r.EXPECT().DeleteResume(id, "test@gmail.com").Return(nil)
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
			mockBehavior: func(r *mock_usecases.MockResume, id uint) {
				r.EXPECT().DeleteResume(id, "test@gmail.com").Return(errorHandler.ErrBadRequest)
			},
			sessionRepBehavior: func(sessionRep *mock_session.MockRepository, token string) {
				sessionRep.EXPECT().GetSession(token).Return("", fmt.Errorf("getting session error:"))
			},
			expectedStatusCode:   401,
			expectedResponseBody: "{\"descriptors\":\"\",\"error\":\"Клиент не авторизован\"}",
		},
		{
			name:           "resume not found",
			emailFromToken: "test@gmail.com",
			inputId:        42,
			inputParam:     "42",
			mockBehavior: func(r *mock_usecases.MockResume, id uint) {
				r.EXPECT().DeleteResume(id, "test@gmail.com").Return(errorHandler.ErrResumeNotFound)
			},
			sessionRepBehavior: func(sessionRep *mock_session.MockRepository, token string) {
				sessionRep.EXPECT().GetSession(token).Return("test@gmail.com", nil)
			},
			expectedStatusCode:   404,
			expectedResponseBody: "{\"descriptors\":\"\",\"error\":\"Резюме не найдено\"}",
		},
		{
			name:           "user is not author",
			emailFromToken: "test@gmail.com",
			inputId:        42,
			inputParam:     "42",
			mockBehavior: func(r *mock_usecases.MockResume, id uint) {
				r.EXPECT().DeleteResume(id, "test@gmail.com").Return(errorHandler.ErrUnauthorized)
			},
			sessionRepBehavior: func(sessionRep *mock_session.MockRepository, token string) {
				sessionRep.EXPECT().GetSession(token).Return("test@gmail.com", nil)
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

			mockUseCase := mock_usecases.NewMockResume(c)
			sessionRep := mock_session.NewMockRepository(c)
			sessionMiddlware := middleware.NewSessionMiddleware(sessionRep)

			if testCase.emailFromToken != "" {
				testCase.mockBehavior(mockUseCase, testCase.inputId)
			}

			testCase.sessionRepBehavior(sessionRep, testCase.inputToken)

			handler := ResumeHandler{
				resumeUseCase: mockUseCase,
			}

			r := gin.New()
			r.POST("/:id", sessionMiddlware.Session, handler.DeleteResume, errorHandler.Middleware())

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
