package impl

import (
	"HeadHunter/internal/entity/models"
	"HeadHunter/mail_microservice/handler"
	mock_usecase "HeadHunter/mail_microservice/usecase/mocks"
	"HeadHunter/pkg/errorHandler"
	"context"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestMailService_SendConfirmCode(t *testing.T) {
	type mockBehavior func(service *mock_usecase.MockMail, in *handler.Email)
	ctx := context.Background()
	testTable := []struct {
		name         string
		in           *handler.Email
		expectedErr  error
		mockBehavior mockBehavior
	}{
		{
			name:        "ok",
			in:          &handler.Email{Value: "example@gmail.com"},
			expectedErr: nil,
			mockBehavior: func(service *mock_usecase.MockMail, in *handler.Email) {
				service.EXPECT().SendConfirmCode(in.Value).Return(nil)
			},
		},
		{
			name:        "cannot send message",
			in:          &handler.Email{Value: "example@gmail.com"},
			expectedErr: errorHandler.ErrBadRequest,
			mockBehavior: func(service *mock_usecase.MockMail, in *handler.Email) {
				service.EXPECT().SendConfirmCode(in.Value).Return(errorHandler.ErrBadRequest)
			},
		},
		{
			name:        "empty in",
			in:          nil,
			expectedErr: errorHandler.ErrBadRequest,
			mockBehavior: func(service *mock_usecase.MockMail, in *handler.Email) {
			},
		},
	}
	for _, test := range testTable {
		testCase := test
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()
			c := gomock.NewController(t)
			defer c.Finish()

			usecase := mock_usecase.NewMockMail(c)

			testCase.mockBehavior(usecase, testCase.in)
			mailHandler := MailHandler{mailUseCase: usecase}
			_, err := mailHandler.SendConfirmCode(ctx, testCase.in)
			assert.Equal(t, testCase.expectedErr, err)
		})
	}
}

func TestMailService_SendApplicantMailing(t *testing.T) {
	ctx := context.Background()
	type mockBehavior func(service *mock_usecase.MockMail, in *handler.ApplicantMailingData)
	testTable := []struct {
		name         string
		in           *handler.ApplicantMailingData
		expectedErr  error
		mockBehavior mockBehavior
	}{
		{
			name:        "ok",
			in:          &handler.ApplicantMailingData{},
			expectedErr: nil,
			mockBehavior: func(service *mock_usecase.MockMail, in *handler.ApplicantMailingData) {
				vacancies := make([]*models.Vacancy, len(in.Vac))
				for i, preview := range in.Vac {
					vacancies[i].ID = uint(preview.Id)
					vacancies[i].Title = preview.Title
					vacancies[i].Image = preview.Image
				}
				service.EXPECT().SendApplicantMailing(in.Emails, vacancies).Return(nil)
			},
		},
		{
			name:        "cannot send message",
			in:          &handler.ApplicantMailingData{},
			expectedErr: errorHandler.ErrBadRequest,
			mockBehavior: func(service *mock_usecase.MockMail, in *handler.ApplicantMailingData) {
				vacancies := make([]*models.Vacancy, len(in.Vac))
				for i, preview := range in.Vac {
					vacancies[i].ID = uint(preview.Id)
					vacancies[i].Title = preview.Title
					vacancies[i].Image = preview.Image
				}
				service.EXPECT().SendApplicantMailing(in.Emails, vacancies).Return(errorHandler.ErrBadRequest)
			},
		},
		{
			name:        "empty in",
			in:          nil,
			expectedErr: errorHandler.ErrBadRequest,
			mockBehavior: func(service *mock_usecase.MockMail, in *handler.ApplicantMailingData) {
			},
		},
	}
	for _, test := range testTable {
		testCase := test
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()
			c := gomock.NewController(t)
			defer c.Finish()

			usecase := mock_usecase.NewMockMail(c)

			testCase.mockBehavior(usecase, testCase.in)
			mailHandler := MailHandler{mailUseCase: usecase}
			_, err := mailHandler.SendApplicantMailing(ctx, testCase.in)
			assert.Equal(t, testCase.expectedErr, err)
		})
	}
}

func TestMailService_SendEmployerMailing(t *testing.T) {
	ctx := context.Background()
	type mockBehavior func(service *mock_usecase.MockMail, in *handler.EmployerMailingData)
	testTable := []struct {
		name         string
		in           *handler.EmployerMailingData
		expectedErr  error
		mockBehavior mockBehavior
	}{
		{
			name:        "ok",
			in:          &handler.EmployerMailingData{},
			expectedErr: nil,
			mockBehavior: func(service *mock_usecase.MockMail, in *handler.EmployerMailingData) {
				applicants := make([]*models.UserAccount, len(in.Emp))
				for i, preview := range in.Emp {
					applicants[i].ID = uint(preview.Id)
					applicants[i].ApplicantName = preview.ApplicantName
					applicants[i].ApplicantSurname = preview.ApplicantSurname
					applicants[i].Image = preview.Image
					applicants[i].Status = preview.Status
					applicants[i].Location = preview.Location
				}
				service.EXPECT().SendEmployerMailing(in.Emails, applicants).Return(nil)
			},
		},
		{
			name:        "cannot send message",
			in:          &handler.EmployerMailingData{},
			expectedErr: errorHandler.ErrBadRequest,
			mockBehavior: func(service *mock_usecase.MockMail, in *handler.EmployerMailingData) {
				applicants := make([]*models.UserAccount, len(in.Emp))
				for i, preview := range in.Emp {
					applicants[i].ID = uint(preview.Id)
					applicants[i].ApplicantName = preview.ApplicantName
					applicants[i].ApplicantSurname = preview.ApplicantSurname
					applicants[i].Image = preview.Image
					applicants[i].Status = preview.Status
					applicants[i].Location = preview.Location
				}
				service.EXPECT().SendEmployerMailing(in.Emails, applicants).Return(errorHandler.ErrBadRequest)
			},
		},
		{
			name:        "empty in",
			in:          nil,
			expectedErr: errorHandler.ErrBadRequest,
			mockBehavior: func(service *mock_usecase.MockMail, in *handler.EmployerMailingData) {
			},
		},
	}
	for _, test := range testTable {
		testCase := test
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()
			c := gomock.NewController(t)
			defer c.Finish()

			usecase := mock_usecase.NewMockMail(c)

			testCase.mockBehavior(usecase, testCase.in)
			mailHandler := MailHandler{mailUseCase: usecase}
			_, err := mailHandler.SendEmployerMailing(ctx, testCase.in)
			assert.Equal(t, testCase.expectedErr, err)
		})
	}
}
