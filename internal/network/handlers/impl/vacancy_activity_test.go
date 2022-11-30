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
			vacancyId:      42,
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
			vacancyId:      42,
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
			r.POST("/", sessionMiddlware.Session, handler.ApplyForVacancy, errorHandler.Middleware())

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
