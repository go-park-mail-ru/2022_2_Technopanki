package impl

import (
	"HeadHunter/internal/entity/models"
	mock_repository "HeadHunter/internal/repository/mocks"
	"HeadHunter/pkg/errorHandler"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestVacancyActivityService_ApplyForVacancy(t *testing.T) {
	type mockBehavior func(r *mock_repository.MockVacancyActivityRepository, vr *mock_repository.MockVacancyRepository,
		ur *mock_repository.MockUserRepository, nr *mock_repository.MockNotificationRepository,
		email string, userId uint, vacancyId uint, input *models.VacancyActivity)
	testTable := []struct {
		name         string
		email        string
		userId       uint
		vacancyId    uint
		mockBehavior mockBehavior
		input        *models.VacancyActivity
		expectedErr  error
	}{
		{
			name: "ok",
			mockBehavior: func(r *mock_repository.MockVacancyActivityRepository, vr *mock_repository.MockVacancyRepository,
				ur *mock_repository.MockUserRepository, nr *mock_repository.MockNotificationRepository,
				email string, userId uint, vacancyId uint, input *models.VacancyActivity) {
				expectedVacancy := &models.Vacancy{}
				ur.EXPECT().GetUserByEmail(email).Return(&models.UserAccount{UserType: "applicant"}, nil)
				r.EXPECT().ApplyForVacancy(input).Return(nil)
				vr.EXPECT().GetById(uint(0)).Return(expectedVacancy, nil)
			},
			input: &models.VacancyActivity{
				ResumeId: 1,
			},
			expectedErr: nil,
		},
		{
			name: "user not exists",
			mockBehavior: func(r *mock_repository.MockVacancyActivityRepository, vr *mock_repository.MockVacancyRepository,
				ur *mock_repository.MockUserRepository, nr *mock_repository.MockNotificationRepository,
				email string, userId uint, vacancyId uint, input *models.VacancyActivity) {
				ur.EXPECT().GetUserByEmail(email).Return(nil, errorHandler.ErrUserNotExists)
			},
			input: &models.VacancyActivity{
				ResumeId: 1,
			},
			expectedErr: errorHandler.ErrUserNotExists,
		},
	}
	for _, test := range testTable {
		testCase := test
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()
			c := gomock.NewController(t)
			defer c.Finish()

			vacancyActivityRepository := mock_repository.NewMockVacancyActivityRepository(c)
			userRep := mock_repository.NewMockUserRepository(c)
			vacancyRep := mock_repository.NewMockVacancyRepository(c)
			notificationRep := mock_repository.NewMockNotificationRepository(c)
			testCase.mockBehavior(vacancyActivityRepository, vacancyRep, userRep, notificationRep, testCase.email, testCase.userId, testCase.vacancyId, testCase.input)
			vacancyActivityService := VacancyActivityService{
				vacancyActivityRep: vacancyActivityRepository,
				userRep:            userRep,
				notificationRep:    notificationRep,
				vacancyRep:         vacancyRep,
			}
			_, err := vacancyActivityService.ApplyForVacancy(testCase.email, testCase.vacancyId, testCase.input)
			assert.Equal(t, testCase.expectedErr, err)
		})
	}
}

func TestVacancyActivityService_DeleteUserApply(t *testing.T) {
	type mockBehavior func(r *mock_repository.MockVacancyActivityRepository, ur *mock_repository.MockUserRepository, email string, userId uint, applyId uint)
	testTable := []struct {
		name         string
		userId       uint
		applyId      uint
		email        string
		mockBehavior mockBehavior
		expectedErr  error
	}{
		{
			name: "ok",
			mockBehavior: func(r *mock_repository.MockVacancyActivityRepository, ur *mock_repository.MockUserRepository, email string, userId uint, applyId uint) {
				ur.EXPECT().GetUserByEmail(email).Return(&models.UserAccount{UserType: "applicant"}, nil)
				r.EXPECT().DeleteUserApply(userId, applyId).Return(nil)
			},
			expectedErr: nil,
		}, {
			name: "cannot get author",
			mockBehavior: func(r *mock_repository.MockVacancyActivityRepository, ur *mock_repository.MockUserRepository, email string, userId uint, applyId uint) {
				ur.EXPECT().GetUserByEmail(email).Return(nil, errorHandler.ErrUserNotExists)
			},
			expectedErr: errorHandler.ErrUserNotExists,
		}, {
			name: "cannot delete apply",
			mockBehavior: func(r *mock_repository.MockVacancyActivityRepository, ur *mock_repository.MockUserRepository, email string, userId uint, applyId uint) {
				ur.EXPECT().GetUserByEmail(email).Return(&models.UserAccount{UserType: "applicant"}, nil)
				r.EXPECT().DeleteUserApply(userId, applyId).Return(errorHandler.ErrBadRequest)
			},
			expectedErr: errorHandler.ErrBadRequest,
		},
	}
	for _, test := range testTable {
		testCase := test
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()
			c := gomock.NewController(t)
			defer c.Finish()

			vacancyActivityRepository := mock_repository.NewMockVacancyActivityRepository(c)
			userRep := mock_repository.NewMockUserRepository(c)
			vacancyRep := mock_repository.NewMockVacancyRepository(c)
			notificationRep := mock_repository.NewMockNotificationRepository(c)
			testCase.mockBehavior(vacancyActivityRepository, userRep, testCase.email, testCase.userId, testCase.applyId)
			vacancyActivityService := VacancyActivityService{
				vacancyActivityRep: vacancyActivityRepository,
				userRep:            userRep,
				notificationRep:    notificationRep,
				vacancyRep:         vacancyRep,
			}
			err := vacancyActivityService.DeleteUserApply(testCase.email, testCase.applyId)
			assert.Equal(t, testCase.expectedErr, err)
		})
	}
}

func TestVacancyActivityService_GetAllVacancyApplies(t *testing.T) {
	type mockBehavior func(r *mock_repository.MockVacancyActivityRepository, ur *mock_repository.MockUserRepository, vacancyId uint)
	testTable := []struct {
		name            string
		vacancyId       uint
		mockBehavior    mockBehavior
		expectedApplies []*models.VacancyActivityPreview
		expectedErr     error
	}{
		{
			name: "ok",
			mockBehavior: func(r *mock_repository.MockVacancyActivityRepository, ur *mock_repository.MockUserRepository, vacancyId uint) {
				expected := []*models.VacancyActivityPreview{
					{
						ResumeId:         1,
						ApplicantName:    "User",
						ApplicantSurname: "User",
					},
				}
				r.EXPECT().GetAllVacancyApplies(vacancyId).Return(expected, nil)
			},
			expectedApplies: []*models.VacancyActivityPreview{
				{
					ResumeId:         1,
					ApplicantName:    "User",
					ApplicantSurname: "User",
				},
			},
			expectedErr: nil,
		},
		{
			name: "cannot get applies",
			mockBehavior: func(r *mock_repository.MockVacancyActivityRepository, ur *mock_repository.MockUserRepository, vacancyId uint) {
				r.EXPECT().GetAllVacancyApplies(vacancyId).Return(nil, errorHandler.ErrBadRequest)
			},
			expectedApplies: nil,
			expectedErr:     errorHandler.ErrBadRequest,
		},
	}
	for _, test := range testTable {
		testCase := test
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()
			c := gomock.NewController(t)
			defer c.Finish()

			vacancyActivityRepository := mock_repository.NewMockVacancyActivityRepository(c)
			userRep := mock_repository.NewMockUserRepository(c)
			testCase.mockBehavior(vacancyActivityRepository, userRep, testCase.vacancyId)
			vacancyActivityService := VacancyActivityService{vacancyActivityRep: vacancyActivityRepository, userRep: userRep}
			applies, err := vacancyActivityService.GetAllVacancyApplies(testCase.vacancyId)
			if testCase.expectedErr == nil {
				assert.Equal(t, testCase.expectedApplies, applies)
			}
			assert.Equal(t, testCase.expectedErr, err)
		})
	}
}

func TestVacancyActivityService_GetAllUserApplies(t *testing.T) {
	type mockBehavior func(r *mock_repository.MockVacancyActivityRepository, ur *mock_repository.MockUserRepository, userId uint)
	testTable := []struct {
		name            string
		userId          uint
		mockBehavior    mockBehavior
		expectedApplies []*models.VacancyActivityPreview
		expectedErr     error
	}{
		{
			name: "ok",
			mockBehavior: func(r *mock_repository.MockVacancyActivityRepository, ur *mock_repository.MockUserRepository, userId uint) {
				expected := []*models.VacancyActivityPreview{
					{
						ResumeId:         1,
						ApplicantName:    "User",
						ApplicantSurname: "User",
					},
				}
				r.EXPECT().GetAllUserApplies(userId).Return(expected, nil)
			},
			expectedApplies: []*models.VacancyActivityPreview{
				{
					ResumeId:         1,
					ApplicantName:    "User",
					ApplicantSurname: "User",
				},
			},
			expectedErr: nil,
		},
		{
			name: "cannot get applies",
			mockBehavior: func(r *mock_repository.MockVacancyActivityRepository, ur *mock_repository.MockUserRepository, userId uint) {
				r.EXPECT().GetAllUserApplies(userId).Return(nil, errorHandler.ErrBadRequest)
			},
			expectedApplies: nil,
			expectedErr:     errorHandler.ErrBadRequest,
		},
	}
	for _, test := range testTable {
		testCase := test
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()
			c := gomock.NewController(t)
			defer c.Finish()

			vacancyActivityRepository := mock_repository.NewMockVacancyActivityRepository(c)
			userRep := mock_repository.NewMockUserRepository(c)
			testCase.mockBehavior(vacancyActivityRepository, userRep, testCase.userId)
			vacancyActivityService := VacancyActivityService{vacancyActivityRep: vacancyActivityRepository, userRep: userRep}
			applies, err := vacancyActivityService.GetAllUserApplies(testCase.userId)
			if testCase.expectedErr == nil {
				assert.Equal(t, testCase.expectedApplies, applies)
			}
			assert.Equal(t, testCase.expectedErr, err)
		})
	}
}
