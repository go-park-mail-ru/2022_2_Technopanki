package impl

import (
	mock_session "HeadHunter/common/session/mocks"
	"HeadHunter/internal/entity/models"
	mock_sender "HeadHunter/mail_microservice/usecase/sender/mocks"
	"HeadHunter/pkg/errorHandler"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"testing"
)

func TestMailService_SendConfirmCode(t *testing.T) {
	type mockBehavior func(sender *mock_sender.MockSender, session *mock_session.MockRepository, email string)
	testTable := []struct {
		name         string
		email        string
		expectedErr  error
		mockBehavior mockBehavior
	}{
		{
			name:        "ok",
			email:       "example@gmail.com",
			expectedErr: nil,
			mockBehavior: func(sender *mock_sender.MockSender, session *mock_session.MockRepository, email string) {
				code := "some_code"
				session.EXPECT().GetCodeFromEmail(email).Return("", errorHandler.ErrBadRequest)
				session.EXPECT().CreateConfirmationCode(email).Return(code, nil)
				sender.EXPECT().SendConfirmCode(email, code).Return(nil)
			},
		},
		{
			name:        "cannot create code",
			email:       "example@gmail.com",
			expectedErr: errorHandler.ErrBadRequest,
			mockBehavior: func(sender *mock_sender.MockSender, session *mock_session.MockRepository, email string) {
				session.EXPECT().GetCodeFromEmail(email).Return("", errorHandler.ErrBadRequest)
				session.EXPECT().CreateConfirmationCode(email).Return("", errorHandler.ErrBadRequest)
			},
		},
		{
			name:        "cannot send message",
			email:       "example@gmail.com",
			expectedErr: errorHandler.ErrBadRequest,
			mockBehavior: func(sender *mock_sender.MockSender, session *mock_session.MockRepository, email string) {
				code := "some_code"
				session.EXPECT().GetCodeFromEmail(email).Return("", errorHandler.ErrBadRequest)
				session.EXPECT().CreateConfirmationCode(email).Return(code, nil)
				sender.EXPECT().SendConfirmCode(email, code).Return(errorHandler.ErrBadRequest)
			},
		},
		{
			name:        "code already exists",
			email:       "example@gmail.com",
			expectedErr: status.Error(codes.AlreadyExists, errorHandler.ErrCodeAlreadyExists.Error()),
			mockBehavior: func(sender *mock_sender.MockSender, session *mock_session.MockRepository, email string) {
				session.EXPECT().GetCodeFromEmail(email).Return("", nil)
			},
		},
	}
	for _, test := range testTable {
		testCase := test
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()
			c := gomock.NewController(t)
			defer c.Finish()

			sessionRepo := mock_session.NewMockRepository(c)
			sender := mock_sender.NewMockSender(c)

			testCase.mockBehavior(sender, sessionRepo, testCase.email)
			mailService := MailService{sender: sender, sessionRepo: sessionRepo}
			err := mailService.SendConfirmCode(testCase.email)
			assert.Equal(t, testCase.expectedErr, err)
		})
	}
}

func TestMailService_SendApplicantMailing(t *testing.T) {
	type mockBehavior func(sender *mock_sender.MockSender, emails []string, vacancies []models.VacancyPreview)
	testTable := []struct {
		name         string
		emails       []string
		vacancies    []*models.VacancyPreview
		expectedErr  error
		mockBehavior mockBehavior
	}{
		{
			name:        "ok",
			emails:      []string{"example@gmail.com"},
			expectedErr: nil,
			mockBehavior: func(sender *mock_sender.MockSender, emails []string, vacancies []models.VacancyPreview) {
				sender.EXPECT().SendApplicantMailing(emails[0], vacancies).Return(nil)
			},
		},
		{
			name:        "ok 2",
			emails:      []string{"example@gmail.com", "example2@gmail.com"},
			expectedErr: nil,
			mockBehavior: func(sender *mock_sender.MockSender, emails []string, vacancies []models.VacancyPreview) {
				sender.EXPECT().SendApplicantMailing(emails[0], vacancies).Return(nil)
				sender.EXPECT().SendApplicantMailing(emails[1], vacancies).Return(nil)
			},
		},
		{
			name:         "ok 3",
			emails:       []string{},
			expectedErr:  nil,
			mockBehavior: func(sender *mock_sender.MockSender, emails []string, vacancies []models.VacancyPreview) {},
		},
		{
			name:        "cannot send message",
			emails:      []string{"example@gmail.com"},
			expectedErr: nil,
			mockBehavior: func(sender *mock_sender.MockSender, emails []string, vacancies []models.VacancyPreview) {
				sender.EXPECT().SendApplicantMailing(emails[0], vacancies).Return(errorHandler.ErrBadRequest)
			},
		},
	}
	for _, test := range testTable {
		testCase := test
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()
			c := gomock.NewController(t)
			defer c.Finish()

			sessionRepo := mock_session.NewMockRepository(c)
			sender := mock_sender.NewMockSender(c)
			objects := make([]models.VacancyPreview, len(testCase.vacancies))
			for i, elem := range testCase.vacancies {
				objects[i] = *elem
			}
			testCase.mockBehavior(sender, testCase.emails, objects)
			mailService := MailService{sender: sender, sessionRepo: sessionRepo}

			err := mailService.SendApplicantMailing(testCase.emails, testCase.vacancies)
			assert.Equal(t, testCase.expectedErr, err)
		})
	}
}

func TestMailService_SendEmployerMailing(t *testing.T) {
	type mockBehavior func(sender *mock_sender.MockSender, emails []string, previews []models.ResumePreview)
	testTable := []struct {
		name         string
		emails       []string
		previews     []*models.ResumePreview
		expectedErr  error
		mockBehavior mockBehavior
	}{
		{
			name:        "ok",
			emails:      []string{"example@gmail.com"},
			expectedErr: nil,
			mockBehavior: func(sender *mock_sender.MockSender, emails []string, previews []models.ResumePreview) {
				sender.EXPECT().SendEmployerMailing(emails[0], previews).Return(nil)
			},
		},
		{
			name:        "ok 2",
			emails:      []string{"example@gmail.com", "example2@gmail.com"},
			expectedErr: nil,
			mockBehavior: func(sender *mock_sender.MockSender, emails []string, previews []models.ResumePreview) {
				sender.EXPECT().SendEmployerMailing(emails[0], previews).Return(nil)
				sender.EXPECT().SendEmployerMailing(emails[1], previews).Return(nil)
			},
		},
		{
			name:         "ok 3",
			emails:       []string{},
			expectedErr:  nil,
			mockBehavior: func(sender *mock_sender.MockSender, emails []string, previews []models.ResumePreview) {},
		},
		{
			name:        "cannot send message",
			emails:      []string{"example@gmail.com"},
			expectedErr: nil,
			mockBehavior: func(sender *mock_sender.MockSender, emails []string, previews []models.ResumePreview) {
				sender.EXPECT().SendEmployerMailing(emails[0], previews).Return(errorHandler.ErrBadRequest)
			},
		},
	}
	for _, test := range testTable {
		testCase := test
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()
			c := gomock.NewController(t)
			defer c.Finish()

			sessionRepo := mock_session.NewMockRepository(c)
			sender := mock_sender.NewMockSender(c)

			objects := make([]models.ResumePreview, len(testCase.previews))
			for i, elem := range testCase.previews {
				objects[i] = *elem
			}
			testCase.mockBehavior(sender, testCase.emails, objects)
			mailService := MailService{sender: sender, sessionRepo: sessionRepo}
			err := mailService.SendEmployerMailing(testCase.emails, testCase.previews)
			assert.Equal(t, testCase.expectedErr, err)
		})
	}
}
