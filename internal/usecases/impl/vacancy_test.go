package impl

import (
	"HeadHunter/internal/entity/models"
	mock_repository "HeadHunter/internal/repository/mocks"
	"HeadHunter/pkg/errorHandler"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestVacancyService_GetById(t *testing.T) {
	type mockBehavior func(r *mock_repository.MockVacancyRepository, id uint)
	testTable := []struct {
		name            string
		id              uint
		mockBehavior    mockBehavior
		expectedVacancy *models.Vacancy
		expectedErr     error
	}{
		{
			name: "ok",
			mockBehavior: func(r *mock_repository.MockVacancyRepository, id uint) {
				expected := &models.Vacancy{
					Title: "Job",
				}
				r.EXPECT().GetById(id).Return(expected, nil)
			},
			expectedVacancy: &models.Vacancy{
				Title: "Job",
			},
			expectedErr: nil,
		},
		{
			name: "cannot get vacancy",
			mockBehavior: func(r *mock_repository.MockVacancyRepository, id uint) {
				r.EXPECT().GetById(id).Return(nil, errorHandler.ErrBadRequest)
			},
			expectedVacancy: nil,
			expectedErr:     errorHandler.ErrBadRequest,
		},
	}
	for _, test := range testTable {
		testCase := test
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()
			c := gomock.NewController(t)
			defer c.Finish()

			vacancyRepository := mock_repository.NewMockVacancyRepository(c)
			userRep := mock_repository.NewMockUserRepository(c)
			testCase.mockBehavior(vacancyRepository, testCase.id)
			VacancyService := VacancyService{vacancyRep: vacancyRepository, userRep: userRep}
			vacancy, err := VacancyService.GetById(testCase.id)
			if testCase.expectedErr == nil {
				assert.Equal(t, *testCase.expectedVacancy, *vacancy)
			}
			assert.Equal(t, testCase.expectedErr, err)
		})
	}
}

func TestVacancyService_GetByUserId(t *testing.T) {
	type mockBehavior func(r *mock_repository.MockVacancyRepository, userId uint)
	testTable := []struct {
		name              string
		userId            uint
		mockBehavior      mockBehavior
		expectedVacancies []*models.Vacancy
		expectedErr       error
	}{
		{
			name: "ok",
			mockBehavior: func(r *mock_repository.MockVacancyRepository, userId uint) {
				expected := []*models.Vacancy{
					{
						Title: "Job",
					},
				}
				r.EXPECT().GetByUserId(userId).Return(expected, nil)
			},
			expectedVacancies: []*models.Vacancy{
				{
					Title: "Job",
				},
			},
			expectedErr: nil,
		},
		{
			name: "cannot get vacancy",
			mockBehavior: func(r *mock_repository.MockVacancyRepository, userId uint) {
				r.EXPECT().GetByUserId(userId).Return(nil, errorHandler.ErrBadRequest)
			},
			expectedVacancies: nil,
			expectedErr:       errorHandler.ErrBadRequest,
		},
	}
	for _, test := range testTable {
		testCase := test
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()
			c := gomock.NewController(t)
			defer c.Finish()

			vacancyRepository := mock_repository.NewMockVacancyRepository(c)
			userRep := mock_repository.NewMockUserRepository(c)
			testCase.mockBehavior(vacancyRepository, testCase.userId)
			vacancyService := VacancyService{vacancyRep: vacancyRepository, userRep: userRep}
			vacancy, err := vacancyService.GetByUserId(testCase.userId)
			if testCase.expectedErr == nil {
				assert.Equal(t, testCase.expectedVacancies, vacancy)
			}
			assert.Equal(t, testCase.expectedErr, err)
		})

	}
}

func TestVacancyService_Create(t *testing.T) {
	type mockBehavior func(r *mock_repository.MockVacancyRepository, ur *mock_repository.MockUserRepository, vacancy *models.Vacancy, userId uint, email string)
	testTable := []struct {
		name         string
		email        string
		userId       uint
		mockBehavior mockBehavior
		input        *models.Vacancy
		expectedId   uint
		expectedErr  error
	}{
		{
			name: "ok",
			mockBehavior: func(r *mock_repository.MockVacancyRepository, ur *mock_repository.MockUserRepository, vacancy *models.Vacancy, userId uint, email string) {
				ur.EXPECT().GetUserByEmail(email).Return(&models.UserAccount{UserType: "employer"}, nil)
				r.EXPECT().Create(vacancy).Return(uint(0), nil)
			},
			input: &models.Vacancy{
				Title:       "Job",
				Description: "Some information for this job",
			},
			expectedId:  0,
			expectedErr: nil,
		},
		{
			name: "user not exists",
			mockBehavior: func(r *mock_repository.MockVacancyRepository, ur *mock_repository.MockUserRepository, vacancy *models.Vacancy, userId uint, email string) {
				ur.EXPECT().GetUserByEmail(email).Return(nil, errorHandler.ErrUserNotExists)
			},
			input: &models.Vacancy{
				Title:       "Job",
				Description: "some description for job",
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

			vacancyRepository := mock_repository.NewMockVacancyRepository(c)
			userRep := mock_repository.NewMockUserRepository(c)
			testCase.mockBehavior(vacancyRepository, userRep, testCase.input, testCase.userId, testCase.email)
			vacancyService := VacancyService{vacancyRep: vacancyRepository, userRep: userRep}
			vacancyId, err := vacancyService.Create(testCase.email, testCase.input)
			if testCase.expectedErr == nil {
				assert.Equal(t, testCase.expectedId, vacancyId)
			}
			assert.Equal(t, testCase.expectedErr, err)
		})
	}
}

func TestVacancyService_Delete(t *testing.T) {
	type mockBehavior func(r *mock_repository.MockVacancyRepository, ur *mock_repository.MockUserRepository, email string, vacancyId uint, userId uint)
	testTable := []struct {
		name         string
		userId       uint
		vacancyId    uint
		email        string
		mockBehavior mockBehavior
		expectedErr  error
	}{
		{
			name: "ok",
			mockBehavior: func(r *mock_repository.MockVacancyRepository, ur *mock_repository.MockUserRepository, email string, vacancyId uint, userId uint) {
				ur.EXPECT().GetUserByEmail(email).Return(&models.UserAccount{UserType: "employer"}, nil)
				r.EXPECT().Delete(userId, vacancyId).Return(nil)
			},
			expectedErr: nil,
		}, {
			name: "cannot get author",
			mockBehavior: func(r *mock_repository.MockVacancyRepository, ur *mock_repository.MockUserRepository, email string, vacancyId uint, userId uint) {
				ur.EXPECT().GetUserByEmail(email).Return(nil, errorHandler.ErrUserNotExists)
			},
			expectedErr: errorHandler.ErrUserNotExists,
		}, {
			name: "cannot delete vacancy",
			mockBehavior: func(r *mock_repository.MockVacancyRepository, ur *mock_repository.MockUserRepository, email string, vacancyId uint, userId uint) {
				ur.EXPECT().GetUserByEmail(email).Return(&models.UserAccount{UserType: "employer"}, nil)
				r.EXPECT().Delete(userId, vacancyId).Return(errorHandler.ErrBadRequest)
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

			vacancyRepository := mock_repository.NewMockVacancyRepository(c)
			userRep := mock_repository.NewMockUserRepository(c)
			testCase.mockBehavior(vacancyRepository, userRep, testCase.email, testCase.vacancyId, testCase.userId)
			vacancyService := VacancyService{vacancyRep: vacancyRepository, userRep: userRep}
			err := vacancyService.Delete(testCase.email, testCase.vacancyId)
			assert.Equal(t, testCase.expectedErr, err)
		})
	}
}

func TestVacancyService_Update(t *testing.T) {
	type mockBehavior func(r *mock_repository.MockVacancyRepository, ur *mock_repository.MockUserRepository, email string, vacancyId uint, userId uint, updates *models.Vacancy, old *models.Vacancy)
	testTable := []struct {
		name         string
		email        string
		vacancyId    uint
		userId       uint
		mockBehavior mockBehavior
		updates      *models.Vacancy
		old          *models.Vacancy
		expectedErr  error
	}{
		{
			name: "ok",
			updates: &models.Vacancy{
				Title:       "Job",
				Description: "Some description",
			},
			old: &models.Vacancy{},
			mockBehavior: func(r *mock_repository.MockVacancyRepository, ur *mock_repository.MockUserRepository, email string, vacancyId uint, userId uint, updates *models.Vacancy, old *models.Vacancy) {
				ur.EXPECT().GetUserByEmail(email).Return(&models.UserAccount{UserType: "employer"}, nil)
				r.EXPECT().GetById(vacancyId).Return(&models.Vacancy{}, nil)
				r.EXPECT().Update(userId, vacancyId, old, updates).Return(nil)
			},
			expectedErr: nil,
		},
		{
			name: "user not found",
			mockBehavior: func(r *mock_repository.MockVacancyRepository, ur *mock_repository.MockUserRepository, email string, vacancyId uint, userId uint, updates *models.Vacancy, old *models.Vacancy) {
				ur.EXPECT().GetUserByEmail(email).Return(nil, errorHandler.ErrUserNotExists)
			},
			updates: &models.Vacancy{
				Title:       "Job",
				Description: "Some description",
			},
			expectedErr: errorHandler.ErrUserNotExists,
		},
		{
			name: "vacancy not found",
			mockBehavior: func(r *mock_repository.MockVacancyRepository, ur *mock_repository.MockUserRepository, email string, vacancyId uint, userId uint, updates *models.Vacancy, old *models.Vacancy) {
				ur.EXPECT().GetUserByEmail(email).Return(&models.UserAccount{UserType: "employer"}, nil)
				r.EXPECT().GetById(vacancyId).Return(nil, errorHandler.ErrVacancyNotFound)
			},
			updates: &models.Vacancy{
				Title:       "Job",
				Description: "Some description",
			},
			expectedErr: errorHandler.ErrVacancyNotFound,
		},
		{
			name: "cannot update vacancy",
			mockBehavior: func(r *mock_repository.MockVacancyRepository, ur *mock_repository.MockUserRepository, email string, vacancyId uint, userId uint, updates *models.Vacancy, old *models.Vacancy) {
				ur.EXPECT().GetUserByEmail(email).Return(&models.UserAccount{UserType: "employer"}, nil)
				r.EXPECT().GetById(vacancyId).Return(&models.Vacancy{}, nil)
				r.EXPECT().Update(userId, vacancyId, old, updates).Return(errorHandler.ErrCannotUpdateVacancy)
			},
			updates: &models.Vacancy{
				Title:       "Job",
				Description: "Some description",
			},
			old:         &models.Vacancy{},
			expectedErr: errorHandler.ErrCannotUpdateVacancy,
		},
	}
	for _, test := range testTable {
		testCase := test
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()
			c := gomock.NewController(t)
			defer c.Finish()

			vacancyRepository := mock_repository.NewMockVacancyRepository(c)
			userRep := mock_repository.NewMockUserRepository(c)
			testCase.mockBehavior(vacancyRepository, userRep, testCase.email, testCase.vacancyId, testCase.userId, testCase.updates, testCase.old)
			vacancyService := VacancyService{vacancyRep: vacancyRepository, userRep: userRep}
			err := vacancyService.Update(testCase.email, testCase.vacancyId, testCase.updates)
			assert.Equal(t, testCase.expectedErr, err)
		})
	}
}

//func TestVacancyService_GetAll(t *testing.T) {
//	type mockBehavior func(r *mock_repository.MockVacancyRepository)
//	testTable := []struct {
//		name              string
//		mockBehavior      mockBehavior
//		expectedVacancies []*models.Vacancy
//		expectedErr       error
//	}{
//		{
//			name: "ok",
//			mockBehavior: func(r *mock_repository.MockVacancyRepository) {
//				expected := []*models.Vacancy{
//					{
//						Title: "Job",
//					},
//				}
//				r.EXPECT().GetAll().Return(expected, nil)
//			},
//			expectedVacancies: []*models.Vacancy{
//				{
//					Title: "Job",
//				},
//			},
//			expectedErr: nil,
//		},
//		{
//			name: "cannot get vacancies",
//			mockBehavior: func(r *mock_repository.MockVacancyRepository) {
//				r.EXPECT().GetAll().Return([]*models.Vacancy{}, errorHandler.ErrBadRequest)
//			},
//			expectedVacancies: []*models.Vacancy{},
//			expectedErr:       errorHandler.ErrBadRequest,
//		},
//	}
//	for _, test := range testTable {
//		testCase := test
//		t.Run(test.name, func(t *testing.T) {
//			t.Parallel()
//			c := gomock.NewController(t)
//			defer c.Finish()
//
//			vacancyRepository := mock_repository.NewMockVacancyRepository(c)
//			userRep := mock_repository.NewMockUserRepository(c)
//			testCase.mockBehavior(vacancyRepository)
//			vacancyService := VacancyService{vacancyRep: vacancyRepository, userRep: userRep}
//			vacancy, err := vacancyService.GetAll()
//			if testCase.expectedErr == nil {
//				assert.Equal(t, testCase.expectedVacancies, vacancy)
//			}
//			assert.Equal(t, testCase.expectedErr, err)
//		})
//
//	}
//}
