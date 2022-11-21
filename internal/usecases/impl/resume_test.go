package impl

import (
	"HeadHunter/internal/entity/models"
	"HeadHunter/internal/errorHandler"
	mock_repository "HeadHunter/internal/repository/mocks"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestResumeService_GetResume(t *testing.T) {
	type mockBehavior func(r *mock_repository.MockResumeRepository, id uint)
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
			mockBehavior: func(r *mock_repository.MockResumeRepository, id uint) {
				expected := &models.Resume{
					Title: "Job",
				}
				r.EXPECT().GetResume(id).Return(expected, nil)
			},
			expectedResume: &models.Resume{
				Title: "Job",
			},
			expectedErr: nil,
		}, {
			name: "error",
			mockBehavior: func(r *mock_repository.MockResumeRepository, id uint) {
				expected := &models.Resume{
					Title: "Job",
				}
				r.EXPECT().GetResume(id).Return(expected, errorHandler.ErrResumeNotFound)
			},
			expectedResume: &models.Resume{
				Title: "Job",
			},
			expectedErr: errorHandler.ErrResumeNotFound,
		},
	}
	for _, test := range testTable {
		testCase := test
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()
			c := gomock.NewController(t)
			defer c.Finish()

			resumeUserRepository := mock_repository.NewMockResumeRepository(c)
			testCase.mockBehavior(resumeUserRepository, testCase.id)
			resumeService := ResumeService{resumeRep: resumeUserRepository}
			user, err := resumeService.GetResume(testCase.id, testCase.email)
			if testCase.expectedErr == nil {
				assert.Equal(t, *testCase.expectedResume, *user)
			}
			assert.Equal(t, testCase.expectedErr, err)
		})
	}
}

func TestResumeService_GetResumeByApplicant(t *testing.T) {
	type mockBehavior func(r *mock_repository.MockResumeRepository, id uint)
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
			mockBehavior: func(r *mock_repository.MockResumeRepository, id uint) {
				expected := []*models.Resume{
					{
						Title: "Job",
					},
				}
				r.EXPECT().GetResumeByApplicant(id).Return(expected, nil)
			},
			expectedResumes: []*models.Resume{
				{
					Title: "Job",
				},
			},
			expectedErr: nil,
		}, {
			name: "error",
			mockBehavior: func(r *mock_repository.MockResumeRepository, id uint) {
				expected := []*models.Resume{
					{
						Title: "Job",
					},
				}
				r.EXPECT().GetResumeByApplicant(id).Return(expected, errorHandler.ErrResumeNotFound)
			},
			expectedResumes: []*models.Resume{
				{
					Title: "Job",
				},
			},
			expectedErr: errorHandler.ErrResumeNotFound,
		},
	}
	for _, test := range testTable {
		testCase := test
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()
			c := gomock.NewController(t)
			defer c.Finish()

			resumeUserRepository := mock_repository.NewMockResumeRepository(c)
			testCase.mockBehavior(resumeUserRepository, testCase.id)
			resumeService := ResumeService{resumeRep: resumeUserRepository}
			user, err := resumeService.GetResumeByApplicant(testCase.id, testCase.email)
			if testCase.expectedErr == nil {
				assert.Equal(t, testCase.expectedResumes, user)
			}
			assert.Equal(t, testCase.expectedErr, err)
		})
	}
}

func TestResumeService_GetPreviewResumeByApplicant(t *testing.T) {
	type mockBehavior func(r *mock_repository.MockResumeRepository, id uint)
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
			mockBehavior: func(r *mock_repository.MockResumeRepository, id uint) {
				expected := []*models.Resume{
					{
						Title: "Job",
					},
				}
				r.EXPECT().GetPreviewResumeByApplicant(id).Return(expected, nil)
			},
			expectedResumes: []*models.Resume{
				{
					Title: "Job",
				},
			},
			expectedErr: nil,
		}, {
			name: "error",
			mockBehavior: func(r *mock_repository.MockResumeRepository, id uint) {
				expected := []*models.Resume{
					{
						Title: "Job",
					},
				}
				r.EXPECT().GetPreviewResumeByApplicant(id).Return(expected, errorHandler.ErrResumeNotFound)
			},
			expectedResumes: []*models.Resume{
				{
					Title: "Job",
				},
			},
			expectedErr: errorHandler.ErrResumeNotFound,
		},
	}
	for _, test := range testTable {
		testCase := test
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()
			c := gomock.NewController(t)
			defer c.Finish()

			resumeUserRepository := mock_repository.NewMockResumeRepository(c)
			testCase.mockBehavior(resumeUserRepository, testCase.id)
			resumeService := ResumeService{resumeRep: resumeUserRepository}
			user, err := resumeService.GetPreviewResumeByApplicant(testCase.id, testCase.email)
			if testCase.expectedErr == nil {
				assert.Equal(t, testCase.expectedResumes, user)
			}
			assert.Equal(t, testCase.expectedErr, err)
		})
	}
}

func TestResumeService_DeleteResume(t *testing.T) {
	type mockBehavior func(r *mock_repository.MockResumeRepository, id uint)
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
			mockBehavior: func(r *mock_repository.MockResumeRepository, id uint) {
				r.EXPECT().DeleteResume(id).Return(nil)
			},
			expectedErr: nil,
		}, {
			name: "error",
			mockBehavior: func(r *mock_repository.MockResumeRepository, id uint) {
				r.EXPECT().DeleteResume(id).Return(errorHandler.ErrResumeNotFound)
			},
			expectedErr: errorHandler.ErrResumeNotFound,
		},
	}
	for _, test := range testTable {
		testCase := test
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()
			c := gomock.NewController(t)
			defer c.Finish()

			resumeUserRepository := mock_repository.NewMockResumeRepository(c)
			testCase.mockBehavior(resumeUserRepository, testCase.id)
			resumeService := ResumeService{resumeRep: resumeUserRepository}
			err := resumeService.DeleteResume(testCase.id, testCase.email)
			assert.Equal(t, testCase.expectedErr, err)
		})
	}
}

func TestResumeService_CreateResume(t *testing.T) {
	type mockBehavior func(r *mock_repository.MockResumeRepository, resume *models.Resume, email string)
	testTable := []struct {
		name         string
		email        string
		mockBehavior mockBehavior
		input        *models.Resume
		expectedId   uint
		expectedErr  error
	}{
		{
			name: "ok",
			mockBehavior: func(r *mock_repository.MockResumeRepository, resume *models.Resume, email string) {
				r.EXPECT().CreateResume(resume, email).Return(nil)
			},
			input: &models.Resume{
				Title:       "Job",
				Description: "Some information for this job",
			},
			expectedId:  1,
			expectedErr: nil,
		}, {
			name: "error",
			mockBehavior: func(r *mock_repository.MockResumeRepository, resume *models.Resume, email string) {
				r.EXPECT().CreateResume(resume, email).Return(errorHandler.ErrResumeNotFound)
			},
			input: &models.Resume{
				Title: "Job",
			},
			expectedId:  1,
			expectedErr: errorHandler.ErrResumeNotFound,
		},
	}
	for _, test := range testTable {
		testCase := test
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()
			c := gomock.NewController(t)
			defer c.Finish()

			resumeUserRepository := mock_repository.NewMockResumeRepository(c)
			testCase.mockBehavior(resumeUserRepository, testCase.input, testCase.email)
			resumeService := ResumeService{resumeRep: resumeUserRepository}
			err := resumeService.CreateResume(testCase.input, testCase.email)
			assert.Equal(t, testCase.expectedErr, err)
		})
	}
}

func TestResumeService_UpdateResume(t *testing.T) {
	type mockBehavior func(r *mock_repository.MockResumeRepository, resume *models.Resume, id uint)
	testTable := []struct {
		name         string
		id           uint
		email        string
		mockBehavior mockBehavior
		input        *models.Resume
		expectedId   uint
		expectedErr  error
	}{
		{
			name: "ok",
			mockBehavior: func(r *mock_repository.MockResumeRepository, resume *models.Resume, id uint) {
				r.EXPECT().UpdateResume(id, resume).Return(nil)
			},
			input: &models.Resume{
				Title: "Job",
			},
			expectedId:  1,
			expectedErr: nil,
		}, {
			name: "error",
			mockBehavior: func(r *mock_repository.MockResumeRepository, resume *models.Resume, id uint) {
				r.EXPECT().UpdateResume(id, resume).Return(errorHandler.ErrResumeNotFound)
			},
			input: &models.Resume{
				Title: "Job",
			},
			expectedId:  1,
			expectedErr: errorHandler.ErrResumeNotFound,
		},
	}
	for _, test := range testTable {
		testCase := test
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()
			c := gomock.NewController(t)
			defer c.Finish()

			resumeUserRepository := mock_repository.NewMockResumeRepository(c)
			testCase.mockBehavior(resumeUserRepository, testCase.input, testCase.id)
			resumeService := ResumeService{resumeRep: resumeUserRepository}
			err := resumeService.UpdateResume(testCase.id, testCase.input, testCase.email)
			assert.Equal(t, testCase.expectedErr, err)
		})
	}
}
