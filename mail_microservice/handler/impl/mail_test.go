package impl

import (
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
