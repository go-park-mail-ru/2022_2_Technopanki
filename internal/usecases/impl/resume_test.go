package impl

import (
	"HeadHunter/configs"
	"HeadHunter/internal/entity/models"
	mock_repository "HeadHunter/internal/repository/mocks"
	"HeadHunter/pkg/errorHandler"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestResumeService_GetResume(t *testing.T) {
	type mockBehavior func(r *mock_repository.MockResumeRepository, ur *mock_repository.MockUserRepository, id uint, email string)
	testTable := []struct {
		name           string
		id             uint
		email          string
		mockBehavior   mockBehavior
		expectedResume *models.Resume
		expectedErr    error
	}{
		{
			name: "ok",
			mockBehavior: func(r *mock_repository.MockResumeRepository, ur *mock_repository.MockUserRepository, id uint, email string) {
				expected := &models.Resume{
					Title: "Job",
				}
				ur.EXPECT().GetUserByEmail(email).Return(&models.UserAccount{UserType: "applicant"}, nil)
				r.EXPECT().GetResume(id).Return(expected, nil)
			},
			expectedResume: &models.Resume{
				Title: "Job",
			},
			expectedErr: nil,
		},
		{
			name: "user not found",
			mockBehavior: func(r *mock_repository.MockResumeRepository, ur *mock_repository.MockUserRepository, id uint, email string) {
				ur.EXPECT().GetUserByEmail(email).Return(nil, errorHandler.ErrUserNotExists)
			},
			expectedResume: nil,
			expectedErr:    errorHandler.ErrUserNotExists,
		},
		{
			name: "user is not author",
			id:   3,
			mockBehavior: func(r *mock_repository.MockResumeRepository, ur *mock_repository.MockUserRepository, id uint, email string) {
				expected := &models.Resume{
					ID:    1,
					Title: "Job",
				}
				user := &models.UserAccount{
					UserType: "applicant",
					ID:       1,
				}
				ur.EXPECT().GetUserByEmail(email).Return(user, nil)
				r.EXPECT().GetResume(id).Return(expected, nil)
				r.EXPECT().GetEmployerIdByVacancyActivity(id).Return(uint(0), errorHandler.ErrUserNotExists)
			},
			expectedResume: &models.Resume{
				ID:    1,
				Title: "Job",
			},
			expectedErr: errorHandler.ErrUnauthorized,
		},
		{
			name: "cannot get resume",
			mockBehavior: func(r *mock_repository.MockResumeRepository, ur *mock_repository.MockUserRepository, id uint, email string) {
				ur.EXPECT().GetUserByEmail(email).Return(&models.UserAccount{UserType: "applicant"}, nil)
				r.EXPECT().GetResume(id).Return(nil, errorHandler.ErrBadRequest)
			},
			expectedResume: nil,
			expectedErr:    errorHandler.ErrBadRequest,
		},
	}
	for _, test := range testTable {
		testCase := test
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()
			c := gomock.NewController(t)
			defer c.Finish()

			resumeUserRepository := mock_repository.NewMockResumeRepository(c)
			userRep := mock_repository.NewMockUserRepository(c)
			testCase.mockBehavior(resumeUserRepository, userRep, testCase.id, testCase.email)
			resumeService := ResumeService{resumeRep: resumeUserRepository, userRep: userRep}
			user, err := resumeService.GetResume(testCase.id, testCase.email)
			if testCase.expectedErr == nil {
				assert.Equal(t, *testCase.expectedResume, *user)
			}
			assert.Equal(t, testCase.expectedErr, err)
		})
	}
}

func TestResumeService_GetResumeByApplicant(t *testing.T) {
	cfg := &configs.Config{
		Validation: configs.ValidationConfig{
			MaxEmailLength:             30,
			MinSurnameLength:           2,
			MaxSurnameLength:           30,
			MinNameLength:              2,
			MaxNameLength:              30,
			MaxPasswordLength:          20,
			MinPasswordLength:          8,
			MinEmailLength:             8,
			MaxResumeTitleLength:       30,
			MinResumeTitleLength:       3,
			MinResumeDescriptionLength: 3,
		},
	}
	type mockBehavior func(r *mock_repository.MockResumeRepository, ur *mock_repository.MockUserRepository, userId uint, email string)
	testTable := []struct {
		name            string
		id              uint
		email           string
		mockBehavior    mockBehavior
		expectedResumes []*models.Resume
		expectedErr     error
	}{
		{
			name: "ok",
			mockBehavior: func(r *mock_repository.MockResumeRepository, ur *mock_repository.MockUserRepository, userId uint, email string) {
				expected := []*models.Resume{
					{
						Title: "Job",
					},
				}
				ur.EXPECT().GetUserByEmail(email).Return(&models.UserAccount{UserType: "applicant"}, nil)
				r.EXPECT().GetResumeByApplicant(userId).Return(expected, nil)
			},
			expectedResumes: []*models.Resume{
				{
					Title: "Job",
				},
			},
			expectedErr: nil,
		},
		{
			name: "user not found",
			mockBehavior: func(r *mock_repository.MockResumeRepository, ur *mock_repository.MockUserRepository, userId uint, email string) {
				ur.EXPECT().GetUserByEmail(email).Return(nil, errorHandler.ErrUserNotExists)
			},
			expectedResumes: []*models.Resume{},
			expectedErr:     errorHandler.ErrUserNotExists,
		},
		{
			name: "user is not author",
			id:   3,
			mockBehavior: func(r *mock_repository.MockResumeRepository, ur *mock_repository.MockUserRepository, userId uint, email string) {
				ur.EXPECT().GetUserByEmail(email).Return(&models.UserAccount{UserType: "applicant", ID: 1}, nil)
			},
			expectedResumes: []*models.Resume{},
			expectedErr:     errorHandler.ErrUnauthorized,
		},
		{
			name: "cannot get preview",
			mockBehavior: func(r *mock_repository.MockResumeRepository, ur *mock_repository.MockUserRepository, userId uint, email string) {
				ur.EXPECT().GetUserByEmail(email).Return(&models.UserAccount{UserType: "applicant"}, nil)
				r.EXPECT().GetResumeByApplicant(userId).Return([]*models.Resume{}, errorHandler.ErrBadRequest)
			},
			expectedResumes: []*models.Resume{},
			expectedErr:     errorHandler.ErrBadRequest,
		},
	}
	for _, test := range testTable {
		testCase := test
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()
			c := gomock.NewController(t)
			defer c.Finish()

			resumeUserRepository := mock_repository.NewMockResumeRepository(c)
			userRep := mock_repository.NewMockUserRepository(c)
			testCase.mockBehavior(resumeUserRepository, userRep, testCase.id, testCase.email)
			resumeService := ResumeService{resumeRep: resumeUserRepository, cfg: cfg, userRep: userRep}
			user, err := resumeService.GetResumeByApplicant(testCase.id, testCase.email)
			if testCase.expectedErr == nil {
				assert.Equal(t, testCase.expectedResumes, user)
			}
			assert.Equal(t, testCase.expectedErr, err)
		})
	}
}

func TestResumeService_GetPreviewResumeByApplicant(t *testing.T) {
	cfg := &configs.Config{
		Validation: configs.ValidationConfig{
			MaxEmailLength:             30,
			MinSurnameLength:           2,
			MaxSurnameLength:           30,
			MinNameLength:              2,
			MaxNameLength:              30,
			MaxPasswordLength:          20,
			MinPasswordLength:          8,
			MinEmailLength:             8,
			MaxResumeTitleLength:       30,
			MinResumeTitleLength:       3,
			MinResumeDescriptionLength: 3,
		},
	}
	type mockBehavior func(r *mock_repository.MockResumeRepository, ur *mock_repository.MockUserRepository, userId uint, email string)
	testTable := []struct {
		name            string
		id              uint
		email           string
		mockBehavior    mockBehavior
		expectedResumes []*models.ResumePreview
		expectedErr     error
	}{
		{
			name: "ok",
			mockBehavior: func(r *mock_repository.MockResumeRepository, ur *mock_repository.MockUserRepository, userId uint, email string) {
				expected := []*models.ResumePreview{
					{
						Title: "Job",
					},
				}
				ur.EXPECT().GetUserByEmail(email).Return(&models.UserAccount{UserType: "applicant"}, nil)
				r.EXPECT().GetPreviewResumeByApplicant(userId).Return(expected, nil)
			},
			expectedResumes: []*models.ResumePreview{
				{
					Title: "Job",
				},
			},
			expectedErr: nil,
		}, {
			name: "user not found",
			mockBehavior: func(r *mock_repository.MockResumeRepository, ur *mock_repository.MockUserRepository, userId uint, email string) {
				ur.EXPECT().GetUserByEmail(email).Return(nil, errorHandler.ErrUserNotExists)
			},
			expectedResumes: []*models.ResumePreview{},
			expectedErr:     errorHandler.ErrUserNotExists,
		},
		{
			name: "user is not author",
			id:   3,
			mockBehavior: func(r *mock_repository.MockResumeRepository, ur *mock_repository.MockUserRepository, userId uint, email string) {
				ur.EXPECT().GetUserByEmail(email).Return(&models.UserAccount{UserType: "applicant", ID: 1}, nil)
			},
			expectedResumes: []*models.ResumePreview{},
			expectedErr:     errorHandler.ErrUnauthorized,
		},
		{
			name: "cannot get preview",
			mockBehavior: func(r *mock_repository.MockResumeRepository, ur *mock_repository.MockUserRepository, userId uint, email string) {
				ur.EXPECT().GetUserByEmail(email).Return(&models.UserAccount{UserType: "applicant"}, nil)
				r.EXPECT().GetPreviewResumeByApplicant(userId).Return([]*models.ResumePreview{}, errorHandler.ErrBadRequest)
			},
			expectedResumes: []*models.ResumePreview{},
			expectedErr:     errorHandler.ErrBadRequest,
		},
	}
	for _, test := range testTable {
		testCase := test
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()
			c := gomock.NewController(t)
			defer c.Finish()

			resumeUserRepository := mock_repository.NewMockResumeRepository(c)
			userRep := mock_repository.NewMockUserRepository(c)
			testCase.mockBehavior(resumeUserRepository, userRep, testCase.id, testCase.email)
			resumeService := ResumeService{resumeRep: resumeUserRepository, cfg: cfg, userRep: userRep}
			user, err := resumeService.GetPreviewResumeByApplicant(testCase.id, testCase.email)
			if testCase.expectedErr == nil {
				assert.Equal(t, testCase.expectedResumes, user)
			}
			assert.Equal(t, testCase.expectedErr, err)
		})
	}
}

func TestResumeService_DeleteResume(t *testing.T) {
	type mockBehavior func(r *mock_repository.MockResumeRepository, ur *mock_repository.MockUserRepository, id uint, email string)
	cfg := &configs.Config{
		Validation: configs.ValidationConfig{
			MaxEmailLength:             30,
			MinSurnameLength:           2,
			MaxSurnameLength:           30,
			MinNameLength:              2,
			MaxNameLength:              30,
			MaxPasswordLength:          20,
			MinPasswordLength:          8,
			MinEmailLength:             8,
			MaxResumeTitleLength:       30,
			MinResumeTitleLength:       3,
			MinResumeDescriptionLength: 3,
		},
	}
	testTable := []struct {
		name         string
		id           uint
		email        string
		mockBehavior mockBehavior
		expectedErr  error
	}{
		{
			name: "ok",
			mockBehavior: func(r *mock_repository.MockResumeRepository, ur *mock_repository.MockUserRepository, id uint, email string) {
				ur.EXPECT().GetUserByEmail(email).Return(&models.UserAccount{UserType: "applicant"}, nil)
				r.EXPECT().GetResume(id).Return(&models.Resume{}, nil)
				r.EXPECT().DeleteResume(id).Return(nil)
			},
			expectedErr: nil,
		}, {
			name: "cannot get author",
			mockBehavior: func(r *mock_repository.MockResumeRepository, ur *mock_repository.MockUserRepository, id uint, email string) {
				ur.EXPECT().GetUserByEmail(email).Return(nil, errorHandler.ErrUserNotExists)
			},
			expectedErr: errorHandler.ErrUserNotExists,
		},
		{
			name: "resume not found",
			mockBehavior: func(r *mock_repository.MockResumeRepository, ur *mock_repository.MockUserRepository, id uint, email string) {
				ur.EXPECT().GetUserByEmail(email).Return(&models.UserAccount{UserType: "applicant"}, nil)
				r.EXPECT().GetResume(id).Return(nil, errorHandler.ErrResumeNotFound)
			},
			expectedErr: errorHandler.ErrResumeNotFound,
		},
		{
			name: "user is not author",
			mockBehavior: func(r *mock_repository.MockResumeRepository, ur *mock_repository.MockUserRepository, id uint, email string) {
				ur.EXPECT().GetUserByEmail(email).Return(&models.UserAccount{UserType: "applicant", ID: 1}, nil)
				r.EXPECT().GetResume(id).Return(&models.Resume{UserAccountId: 2}, nil)
			},
			expectedErr: errorHandler.ErrUnauthorized,
		},
		{
			name: "cannot delete resume",
			mockBehavior: func(r *mock_repository.MockResumeRepository, ur *mock_repository.MockUserRepository, id uint, email string) {
				ur.EXPECT().GetUserByEmail(email).Return(&models.UserAccount{UserType: "applicant"}, nil)
				r.EXPECT().GetResume(id).Return(&models.Resume{}, nil)
				r.EXPECT().DeleteResume(id).Return(errorHandler.ErrBadRequest)
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

			resumeUserRepository := mock_repository.NewMockResumeRepository(c)
			userRep := mock_repository.NewMockUserRepository(c)
			testCase.mockBehavior(resumeUserRepository, userRep, testCase.id, testCase.email)
			resumeService := ResumeService{resumeRep: resumeUserRepository, cfg: cfg, userRep: userRep}
			err := resumeService.DeleteResume(testCase.id, testCase.email)
			assert.Equal(t, testCase.expectedErr, err)
		})
	}
}

func TestResumeService_CreateResume(t *testing.T) {
	type mockBehavior func(r *mock_repository.MockResumeRepository, ur *mock_repository.MockUserRepository, resume *models.Resume, userId uint, email string)
	cfg := &configs.Config{
		Validation: configs.ValidationConfig{
			MaxEmailLength:             30,
			MinSurnameLength:           2,
			MaxSurnameLength:           30,
			MinNameLength:              2,
			MaxNameLength:              30,
			MaxPasswordLength:          20,
			MinPasswordLength:          8,
			MinEmailLength:             8,
			MaxResumeTitleLength:       30,
			MinResumeTitleLength:       3,
			MinResumeDescriptionLength: 3,
		},
	}
	testTable := []struct {
		name         string
		email        string
		userId       uint
		mockBehavior mockBehavior
		input        *models.Resume
		expectedErr  error
	}{
		{
			name: "ok",
			mockBehavior: func(r *mock_repository.MockResumeRepository, ur *mock_repository.MockUserRepository, resume *models.Resume, userId uint, email string) {
				ur.EXPECT().GetUserByEmail(email).Return(&models.UserAccount{UserType: "applicant"}, nil)
				r.EXPECT().CreateResume(resume, userId).Return(nil)
			},
			input: &models.Resume{
				Title:       "Job",
				Description: "Some information for this job",
			},
			expectedErr: nil,
		}, {
			name: "error",
			mockBehavior: func(r *mock_repository.MockResumeRepository, ur *mock_repository.MockUserRepository, resume *models.Resume, userId uint, email string) {
			},
			input: &models.Resume{
				Title: "Job",
			},
			expectedErr: errorHandler.InvalidResumeDescriptionLength,
		},
		{
			name: "user not exists",
			mockBehavior: func(r *mock_repository.MockResumeRepository, ur *mock_repository.MockUserRepository, resume *models.Resume, userId uint, email string) {
				ur.EXPECT().GetUserByEmail(email).Return(nil, errorHandler.ErrUserNotExists)
			},
			input: &models.Resume{
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

			resumeUserRepository := mock_repository.NewMockResumeRepository(c)
			userRep := mock_repository.NewMockUserRepository(c)
			testCase.mockBehavior(resumeUserRepository, userRep, testCase.input, testCase.userId, testCase.email)
			resumeService := ResumeService{resumeRep: resumeUserRepository, cfg: cfg, userRep: userRep}
			err := resumeService.CreateResume(testCase.input, testCase.email)
			assert.Equal(t, testCase.expectedErr, err)
		})
	}
}

func TestResumeService_UpdateResume(t *testing.T) {
	type mockBehavior func(r *mock_repository.MockResumeRepository, ur *mock_repository.MockUserRepository, resume *models.Resume, id uint, email string)
	cfg := &configs.Config{
		Validation: configs.ValidationConfig{
			MaxEmailLength:             30,
			MinSurnameLength:           2,
			MaxSurnameLength:           30,
			MinNameLength:              2,
			MaxNameLength:              30,
			MaxPasswordLength:          20,
			MinPasswordLength:          8,
			MinEmailLength:             8,
			MaxResumeTitleLength:       30,
			MinResumeTitleLength:       3,
			MinResumeDescriptionLength: 3,
		},
	}
	testTable := []struct {
		name         string
		id           uint
		email        string
		mockBehavior mockBehavior
		input        *models.Resume
		expectedErr  error
	}{
		{
			name: "ok",
			mockBehavior: func(r *mock_repository.MockResumeRepository, ur *mock_repository.MockUserRepository, resume *models.Resume, id uint, email string) {
				ur.EXPECT().GetUserByEmail(email).Return(&models.UserAccount{UserType: "applicant"}, nil)
				r.EXPECT().GetResume(id).Return(&models.Resume{}, nil)
				r.EXPECT().UpdateResume(id, resume).Return(nil)
			},
			input: &models.Resume{
				Title:       "Job",
				Description: "Some description",
			},
			expectedErr: nil,
		}, {
			name: "validation error",
			mockBehavior: func(r *mock_repository.MockResumeRepository, ur *mock_repository.MockUserRepository, resume *models.Resume, id uint, email string) {
			},
			input: &models.Resume{
				Title: "Job",
			},
			expectedErr: errorHandler.InvalidResumeDescriptionLength,
		},
		{
			name: "user not found",
			mockBehavior: func(r *mock_repository.MockResumeRepository, ur *mock_repository.MockUserRepository, resume *models.Resume, id uint, email string) {
				ur.EXPECT().GetUserByEmail(email).Return(nil, errorHandler.ErrUserNotExists)
			},
			input: &models.Resume{
				Title:       "Job",
				Description: "Some description",
			},
			expectedErr: errorHandler.ErrUserNotExists,
		},
		{
			name: "resume not found",
			mockBehavior: func(r *mock_repository.MockResumeRepository, ur *mock_repository.MockUserRepository, resume *models.Resume, id uint, email string) {
				ur.EXPECT().GetUserByEmail(email).Return(&models.UserAccount{UserType: "applicant"}, nil)
				r.EXPECT().GetResume(id).Return(nil, errorHandler.ErrResumeNotFound)
			},
			input: &models.Resume{
				Title:       "Job",
				Description: "Some description",
			},
			expectedErr: errorHandler.ErrResumeNotFound,
		},
		{
			name: "user is not author",
			mockBehavior: func(r *mock_repository.MockResumeRepository, ur *mock_repository.MockUserRepository, resume *models.Resume, id uint, email string) {
				ur.EXPECT().GetUserByEmail(email).Return(&models.UserAccount{UserType: "applicant", ID: 1}, nil)
				r.EXPECT().GetResume(id).Return(&models.Resume{UserAccountId: 2}, nil)
			},
			input: &models.Resume{
				Title:       "Job",
				Description: "Some description",
			},
			expectedErr: errorHandler.ErrUnauthorized,
		},
		{
			name: "cannot update resume",
			mockBehavior: func(r *mock_repository.MockResumeRepository, ur *mock_repository.MockUserRepository, resume *models.Resume, id uint, email string) {
				ur.EXPECT().GetUserByEmail(email).Return(&models.UserAccount{UserType: "applicant"}, nil)
				r.EXPECT().GetResume(id).Return(&models.Resume{}, nil)
				r.EXPECT().UpdateResume(id, resume).Return(errorHandler.ErrBadRequest)
			},
			input: &models.Resume{
				Title:       "Job",
				Description: "Some description",
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

			resumeUserRepository := mock_repository.NewMockResumeRepository(c)
			userRep := mock_repository.NewMockUserRepository(c)
			testCase.mockBehavior(resumeUserRepository, userRep, testCase.input, testCase.id, testCase.email)
			resumeService := ResumeService{resumeRep: resumeUserRepository, cfg: cfg, userRep: userRep}
			err := resumeService.UpdateResume(testCase.id, testCase.input, testCase.email)
			assert.Equal(t, testCase.expectedErr, err)
		})
	}
}
