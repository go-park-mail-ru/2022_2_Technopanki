package impl

import (
	mock_session "HeadHunter/common/session/mocks"
	"HeadHunter/internal/entity/models"
	"HeadHunter/internal/network/middleware"
	mock_usecases "HeadHunter/internal/usecases/mocks"
	errorHandler2 "HeadHunter/pkg/errorHandler"
	"bytes"
	"github.com/gin-gonic/gin"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestNotificationHandler_GetNotifications(t *testing.T) {
	type mockBehavior func(n *mock_usecases.MockNotification, s *mock_session.MockRepository, token string, email string)
	testTable := []struct {
		name                 string
		inputToken           string
		emailFromToken       string
		mockBehavior         mockBehavior
		expectedStatusCode   int
		expectedResponseBody string
	}{
		{
			name:           "valid",
			emailFromToken: "test@gmail.com",
			mockBehavior: func(n *mock_usecases.MockNotification, s *mock_session.MockRepository, token string, email string) {
				var notifications []*models.NotificationPreview
				n.EXPECT().GetNotificationsByEmail(email).Return(notifications, nil)
				s.EXPECT().GetSession(token).Return(email, nil)
			},
			expectedStatusCode:   200,
			expectedResponseBody: "{\"data\":null}",
		},
		{
			name:           "unauthorized",
			emailFromToken: "",
			mockBehavior: func(n *mock_usecases.MockNotification, s *mock_session.MockRepository, token string, email string) {
				s.EXPECT().GetSession(token).Return(email, errorHandler2.ErrUnauthorized)
			},
			expectedStatusCode:   401,
			expectedResponseBody: "{\"descriptors\":\"\",\"error\":\"Клиент не авторизован\"}",
		},
		{
			name:           "cannot get",
			emailFromToken: "dfsfsf",
			mockBehavior: func(n *mock_usecases.MockNotification, s *mock_session.MockRepository, token string, email string) {
				s.EXPECT().GetSession(token).Return(email, nil)
				n.EXPECT().GetNotificationsByEmail(email).Return(nil, errorHandler2.ErrBadRequest)
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

			mockUseCase := mock_usecases.NewMockNotification(c)
			sessionRep := mock_session.NewMockRepository(c)
			sessionMiddlware := middleware.NewSessionMiddleware(sessionRep)

			testCase.mockBehavior(mockUseCase, sessionRep, testCase.inputToken, testCase.emailFromToken)

			handler := NotificationHandler{
				notificationUseCase: mockUseCase,
			}

			r := gin.New()
			r.GET("/", sessionMiddlware.Session, handler.GetNotifications, errorHandler2.Middleware())

			cookie := &http.Cookie{
				Name:  "session",
				Value: testCase.inputToken,
			}

			w := httptest.NewRecorder()
			req := httptest.NewRequest("GET", "/", bytes.NewBufferString(""))

			req.AddCookie(cookie)
			r.ServeHTTP(w, req)

			assert.Equal(t, testCase.expectedStatusCode, w.Code)
			assert.Equal(t, testCase.expectedResponseBody, w.Body.String())
		})
	}
}
