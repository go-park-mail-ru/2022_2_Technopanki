package impl

import (
	"HeadHunter/internal/entity/models"
	"github.com/stretchr/testify/assert"
	"gopkg.in/DATA-DOG/go-sqlmock.v1"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"regexp"
	"testing"
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
				ID: 1,
			},
		},
	}

	for _, tc := range testTable {
		testCase := tc
		t.Run(testCase.name, func(t *testing.T) {
			t.Parallel()
			//"is_current_job", "start_date", "end_date", "job_title", "company_name", "job_location_city",
			//			"experience_description"
			rows := sqlmock.NewRows([]string{"id", "title", "description", "user_account_id", "created_date"})
			rows = rows.AddRow(testCase.expectedResume.ID, testCase.expectedResume.Title,
				testCase.expectedResume.Description, testCase.expectedResume.UserAccountId,
				testCase.expectedResume.CreatedTime)
			mock.
				ExpectQuery(regexp.QuoteMeta(`SELECT id FROM "resumes" WHERE "resume"."id" = $1`)).
				WithArgs(1).WillReturnRows(rows)

			actualResume, getErr := ResumeDB.GetResume(1)
			if getErr != nil {
				t.Errorf("unexpected err: %s", getErr)
				return
			}

			assert.Equal(t, testCase.expectedResume, actualResume)
			if err := mock.ExpectationsWereMet(); err != nil {
				t.Errorf("there were unfulfilled expectations: %s", err)
			}
		})
	}
}
