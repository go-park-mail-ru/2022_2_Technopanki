package impl

import (
	"HeadHunter/pkg/errorHandler"
	"fmt"
	"regexp"
	"testing"

	"HeadHunter/internal/entity/models"

	"github.com/stretchr/testify/assert"
	"gopkg.in/DATA-DOG/go-sqlmock.v1"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func CreateResumeMock() (*ResumePostgres, sqlmock.Sqlmock, error) {
	db, mock, mockErr := sqlmock.New()
	if mockErr != nil {
		return nil, nil, mockErr
	}

	gormDB, openErr := gorm.Open(postgres.New(postgres.Config{Conn: db}))

	if openErr != nil {
		closeErr := db.Close()
		if closeErr != nil {
			return nil, nil, closeErr
		}
		return nil, nil, closeErr
	}

	repoBoard := NewResumePostgres(gormDB)
	return repoBoard, mock, nil
}

func TestResumePostgres_GetResume(t *testing.T) {
	t.Parallel()
	ResumeDB, mock, mockErr := CreateResumeMock()
	if mockErr != nil {
		t.Errorf("error with creating mock: %s", mockErr)
	}

	testTable := []struct {
		name           string
		expectedErr    error
		expectedResume *models.Resume
	}{
		{
			name:        "ok",
			expectedErr: nil,
			expectedResume: &models.Resume{
				ID:            1,
				Title:         "title",
				Description:   "desc",
				UserAccountId: 1,
				ExperienceDetail: models.ExperienceDetail{
					ResumeId: 1,
				},
				EducationDetail: models.EducationDetail{
					ResumeId: 1,
				},
			},
		},
	}

	for _, tc := range testTable {
		testCase := tc
		t.Run(testCase.name, func(t *testing.T) {
			t.Parallel()
			resumeRows := sqlmock.NewRows([]string{"id", "title", "description", "user_account_id"})
			resumeRows = resumeRows.AddRow(testCase.expectedResume.ID, testCase.expectedResume.Title,
				testCase.expectedResume.Description, testCase.expectedResume.UserAccountId)
			mock.
				ExpectQuery(regexp.QuoteMeta(`SELECT * FROM "resumes" left join experience_details on resumes.id = experience_details.resume_id left join education_details on resumes.id = education_details.resume_id WHERE resumes.id = $1 `)).
				WithArgs(1).WillReturnRows(resumeRows)

			experienceRows := sqlmock.NewRows([]string{"resume_id"})
			experienceRows = experienceRows.AddRow(testCase.expectedResume.ExperienceDetail.ResumeId)
			mock.
				ExpectQuery(regexp.QuoteMeta(`SELECT * FROM "resumes" left join experience_details on resumes.id = experience_details.resume_id left join education_details on resumes.id = education_details.resume_id WHERE resumes.id = $1 AND "resumes"."id" = $2`)).
				WithArgs(1, 1).WillReturnRows(experienceRows)

			educationRows := sqlmock.NewRows([]string{"resume_id"})
			educationRows = educationRows.AddRow(testCase.expectedResume.ExperienceDetail.ResumeId)
			mock.
				ExpectQuery(regexp.QuoteMeta(`SELECT * FROM "resumes" left join experience_details on resumes.id = experience_details.resume_id left join education_details on resumes.id = education_details.resume_id WHERE resumes.id = $1 AND "resumes"."id" = $2`)).
				WithArgs(1, 1).WillReturnRows(educationRows)

			actualResume, getErr := ResumeDB.GetResume(1)

			assert.Equal(t, testCase.expectedErr, getErr)
			assert.Equal(t, testCase.expectedResume, actualResume)
			if err := mock.ExpectationsWereMet(); err != nil {
				t.Errorf("there were unfulfilled expectations: %s", err)
			}
		})
	}
}

func TestResumePostgres_GetPreviewResumeByApplicant(t *testing.T) {
	t.Parallel()
	ResumeDB, mock, mockErr := CreateResumeMock()
	if mockErr != nil {
		t.Errorf("error with creating mock: %s", mockErr)
	}

	testTable := []struct {
		name        string
		id          uint
		expectedErr error
		expected    []*models.ResumePreview
	}{
		{
			name:        "ok",
			id:          1,
			expectedErr: nil,
			expected: []*models.ResumePreview{{
				ApplicantName:    "Zakhar",
				ApplicantSurname: "Urvancev",
				Id:               1,
				Title:            "Some title",
			},
			},
		},
		{
			name:        "not found",
			id:          5,
			expectedErr: errorHandler.ErrResumeNotFound,
			expected:    []*models.ResumePreview{},
		},
	}

	for _, tc := range testTable {
		testCase := tc
		t.Run(testCase.name, func(t *testing.T) {

			resumeRows := sqlmock.NewRows([]string{"id", "title", "applicant_name", "applicant_surname"})
			if len(testCase.expected) > 0 {
				resumeRows = resumeRows.AddRow(testCase.expected[0].Id, testCase.expected[0].Title,
					testCase.expected[0].ApplicantName, testCase.expected[0].ApplicantSurname)
			}
			mock.
				ExpectQuery(regexp.QuoteMeta(`SELECT user_accounts.applicant_name, user_accounts.applicant_surname, user_accounts.image,resumes.id, resumes.title, resumes.created_time FROM "resumes" left join user_accounts on resumes.user_account_id = user_accounts.id WHERE user_account_id = $1`)).
				WithArgs(1).WillReturnRows(resumeRows)

			actualResume, getErr := ResumeDB.GetPreviewResumeByApplicant(1)

			assert.Equal(t, testCase.expectedErr, getErr)
			assert.Equal(t, len(testCase.expected), len(actualResume))
			for i := range actualResume {
				assert.Equal(t, *testCase.expected[i], *actualResume[i])
			}

			if err := mock.ExpectationsWereMet(); err != nil {
				t.Errorf("there were unfulfilled expectations: %s", err)
			}
		})
	}
}

func TestResumePostgres_GetResumeByApplicant(t *testing.T) {
	t.Parallel()
	ResumeDB, mock, mockErr := CreateResumeMock()
	if mockErr != nil {
		t.Errorf("error with creating mock: %s", mockErr)
	}

	testTable := []struct {
		name        string
		id          uint
		expectedErr error
		expected    []*models.Resume
	}{
		{
			name:        "ok",
			id:          1,
			expectedErr: nil,
			expected: []*models.Resume{{
				ID:            1,
				Title:         "title",
				Description:   "desc",
				UserAccountId: 1,
				ExperienceDetail: models.ExperienceDetail{
					ResumeId: 1,
				},
				EducationDetail: models.EducationDetail{
					ResumeId: 1,
				},
			}},
		},
		{
			name:        "not found",
			id:          5,
			expectedErr: errorHandler.ErrResumeNotFound,
			expected:    []*models.Resume{},
		},
	}

	for _, tc := range testTable {
		testCase := tc
		t.Run(testCase.name, func(t *testing.T) {

			resumeRows := sqlmock.NewRows([]string{"id", "title", "description", "user_account_id"})
			if len(testCase.expected) > 0 {
				resumeRows = resumeRows.AddRow(testCase.expected[0].ID, testCase.expected[0].Title,
					testCase.expected[0].Description, testCase.expected[0].UserAccountId)
			}
			mock.
				ExpectQuery(regexp.QuoteMeta(`SELECT * FROM "resumes" left join experience_details on resumes.id = experience_details.resume_id left join education_details on resumes.id = education_details.resume_id WHERE user_account_id = $1`)).
				WithArgs(1).WillReturnRows(resumeRows)

			experienceRows := sqlmock.NewRows([]string{"resume_id"})
			if len(testCase.expected) > 0 {
				experienceRows = experienceRows.AddRow(testCase.expected[0].ExperienceDetail.ResumeId)
			}
			mock.
				ExpectQuery(regexp.QuoteMeta(`SELECT * FROM "resumes" left join experience_details on resumes.id = experience_details.resume_id left join education_details on resumes.id = education_details.resume_id WHERE user_account_id = $1`)).
				WithArgs(1).WillReturnRows(experienceRows)

			educationRows := sqlmock.NewRows([]string{"resume_id"})
			if len(testCase.expected) > 0 {
				educationRows = educationRows.AddRow(testCase.expected[0].ExperienceDetail.ResumeId)
			}
			mock.
				ExpectQuery(regexp.QuoteMeta(`SELECT * FROM "resumes" left join experience_details on resumes.id = experience_details.resume_id left join education_details on resumes.id = education_details.resume_id WHERE user_account_id = $1`)).
				WithArgs(1).WillReturnRows(educationRows)

			actualResume, getErr := ResumeDB.GetResumeByApplicant(1)

			assert.Equal(t, testCase.expectedErr, getErr)
			assert.Equal(t, len(testCase.expected), len(actualResume))
			for i := range actualResume {
				assert.Equal(t, *testCase.expected[i], *actualResume[i])
			}

			if err := mock.ExpectationsWereMet(); err != nil {
				t.Errorf("there were unfulfilled expectations: %s", err)
			}
		})
	}
}

func TestResumePostgres_DeleteResume(t *testing.T) {
	t.Parallel()
	ResumeDB, mock, mockErr := CreateResumeMock()
	if mockErr != nil {
		t.Errorf("error with creating mock: %s", mockErr)
	}

	testTable := []struct {
		name        string
		id          uint
		expectedErr error
		expected    []*models.Resume
	}{
		{
			name:        "ok",
			id:          1,
			expectedErr: nil,
			expected: []*models.Resume{{
				ID:            1,
				Title:         "title",
				Description:   "desc",
				UserAccountId: 1,
				ExperienceDetail: models.ExperienceDetail{
					ResumeId: 1,
				},
				EducationDetail: models.EducationDetail{
					ResumeId: 1,
				},
			}},
		},
	}

	for _, tc := range testTable {
		testCase := tc
		t.Run(testCase.name, func(t *testing.T) {

			resumeRows := sqlmock.NewRows([]string{"id", "title", "description", "user_account_id"})
			if len(testCase.expected) > 0 {
				resumeRows = resumeRows.AddRow(testCase.expected[0].ID, testCase.expected[0].Title,
					testCase.expected[0].Description, testCase.expected[0].UserAccountId)
			}
			mock.ExpectBegin()
			if testCase.expectedErr == nil {
				mock.
					ExpectExec(regexp.QuoteMeta(`DELETE FROM "resumes" WHERE "resumes"."id" = $1`)).
					WithArgs(1).WillReturnResult(sqlmock.NewResult(1, 1))
			} else {
				mock.
					ExpectExec(regexp.QuoteMeta(`DELETE FROM "resumes" WHERE "resumes"."id" = $1`)).
					WithArgs(1).WillReturnError(fmt.Errorf("%w", errorHandler.ErrResumeNotFound))
			}
			mock.ExpectCommit()

			deleteErr := ResumeDB.DeleteResume(1)
			assert.Equal(t, testCase.expectedErr, deleteErr)
			if err := mock.ExpectationsWereMet(); err != nil {
				t.Errorf("there were unfulfilled expectations: %s", err)
			}
		})
	}
}
