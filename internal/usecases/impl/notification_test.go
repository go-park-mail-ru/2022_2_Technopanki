package impl

import (
	"HeadHunter/internal/entity/models"
	mock_repository "HeadHunter/internal/repository/mocks"
	"HeadHunter/pkg/errorHandler"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNotificationService_GetNotificationsByEmail(t *testing.T) {
	type mockBehavior func(r *mock_repository.MockNotificationRepository, u *mock_repository.MockUserRepository, email string)

	testTable := []struct {
		name         string
		email        string
		mockBehavior mockBehavior
		expected     []*models.NotificationPreview
		expectedErr  error
	}{
		{
			name:  "ok apply",
			email: "some@gmail.com",
			mockBehavior: func(r *mock_repository.MockNotificationRepository, u *mock_repository.MockUserRepository, email string) {
				expected := &models.UserAccount{UserType: "employer"}
				u.EXPECT().GetUserByEmail(email).Return(expected, nil)
				r.EXPECT().GetApplyNotificationsByUser(expected.ID).Return([]*models.NotificationPreview{}, nil)

			},
			expected:    []*models.NotificationPreview{},
			expectedErr: nil,
		},
		{
			name:  "ok pdf",
			email: "some@gmail.com",
			mockBehavior: func(r *mock_repository.MockNotificationRepository, u *mock_repository.MockUserRepository, email string) {
				expected := &models.UserAccount{UserType: "applicant"}
				u.EXPECT().GetUserByEmail(email).Return(expected, nil)
				r.EXPECT().GetDownloadPDFNotificationsByUser(expected.ID).Return([]*models.NotificationPreview{}, nil)

			},
			expected:    []*models.NotificationPreview{},
			expectedErr: nil,
		},
		{
			name:  "cannot get applies",
			email: "some@gmail.com",
			mockBehavior: func(r *mock_repository.MockNotificationRepository, u *mock_repository.MockUserRepository, email string) {
				expected := &models.UserAccount{UserType: "employer"}
				u.EXPECT().GetUserByEmail(email).Return(expected, nil)
				r.EXPECT().GetApplyNotificationsByUser(expected.ID).Return([]*models.NotificationPreview{}, errorHandler.ErrBadRequest)

			},
			expected:    []*models.NotificationPreview{},
			expectedErr: errorHandler.ErrBadRequest,
		},
		{
			name:  "not found",
			email: "some@gmail.com",
			mockBehavior: func(r *mock_repository.MockNotificationRepository, u *mock_repository.MockUserRepository, email string) {
				expected := &models.UserAccount{UserType: "employer"}
				u.EXPECT().GetUserByEmail(email).Return(expected, nil)
				r.EXPECT().GetApplyNotificationsByUser(expected.ID).Return([]*models.NotificationPreview{}, errorHandler.ErrNotificationNotFound)

			},
			expected:    []*models.NotificationPreview{},
			expectedErr: nil,
		},
	}

	for _, test := range testTable {
		testCase := test
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()
			c := gomock.NewController(t)
			defer c.Finish()

			mockNotificationRepo := mock_repository.NewMockNotificationRepository(c)
			mockUserRepo := mock_repository.NewMockUserRepository(c)
			testCase.mockBehavior(mockNotificationRepo, mockUserRepo, testCase.email)
			notificationService := NotificationService{notificationRepo: mockNotificationRepo, userRepo: mockUserRepo}
			result, err := notificationService.GetNotificationsByEmail(testCase.email)

			assert.Equal(t, testCase.expected, result)
			assert.Equal(t, testCase.expectedErr, err)
		})
	}
}

func TestNotificationService_CreateNotification(t *testing.T) {
	type mockBehavior func(r *mock_repository.MockNotificationRepository, notic *models.Notification, expected *models.NotificationPreview)

	testTable := []struct {
		name         string
		inputNotic   *models.Notification
		mockBehavior mockBehavior
		expected     *models.NotificationPreview
		expectedErr  error
	}{
		{
			name:       "ok apply",
			inputNotic: &models.Notification{Type: models.AllowedNotificationTypes[0]},
			mockBehavior: func(r *mock_repository.MockNotificationRepository, notic *models.Notification, expected *models.NotificationPreview) {
				r.EXPECT().CreateNotification(notic).Return(nil)
				r.EXPECT().GetNotificationPreviewApply(notic.ID).Return(expected, nil)
			},
			expected:    &models.NotificationPreview{},
			expectedErr: nil,
		},
		{
			name:       "ok download pdf",
			inputNotic: &models.Notification{Type: models.AllowedNotificationTypes[1]},
			mockBehavior: func(r *mock_repository.MockNotificationRepository, notic *models.Notification, expected *models.NotificationPreview) {
				r.EXPECT().CreateNotification(notic).Return(nil)
				r.EXPECT().GetNotificationPreviewDownloadPDF(notic.ID).Return(expected, nil)
			},
			expected:    &models.NotificationPreview{},
			expectedErr: nil,
		},
		{
			name:       "error type of notic",
			inputNotic: &models.Notification{Type: "23232"},
			mockBehavior: func(r *mock_repository.MockNotificationRepository, notic *models.Notification, expected *models.NotificationPreview) {
			},
			expected:    nil,
			expectedErr: errorHandler.ErrBadRequest,
		},
		{
			name:       "cannot create",
			inputNotic: &models.Notification{Type: models.AllowedNotificationTypes[0]},
			mockBehavior: func(r *mock_repository.MockNotificationRepository, notic *models.Notification, expected *models.NotificationPreview) {
				r.EXPECT().CreateNotification(notic).Return(errorHandler.ErrBadRequest)
			},
			expected:    nil,
			expectedErr: errorHandler.ErrBadRequest,
		},
	}

	for _, test := range testTable {
		testCase := test
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()
			c := gomock.NewController(t)
			defer c.Finish()

			mockNotificationRepo := mock_repository.NewMockNotificationRepository(c)
			testCase.mockBehavior(mockNotificationRepo, testCase.inputNotic, testCase.expected)
			notificationService := NotificationService{notificationRepo: mockNotificationRepo}
			result, err := notificationService.CreateNotification(testCase.inputNotic)

			assert.Equal(t, testCase.expected, result)
			assert.Equal(t, testCase.expectedErr, err)
		})
	}
}
