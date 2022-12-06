package impl

import (
	"HeadHunter/internal/entity/models"
	"HeadHunter/pkg/errorHandler"
	"fmt"
	"github.com/stretchr/testify/assert"
	"gopkg.in/DATA-DOG/go-sqlmock.v1"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"regexp"
	"testing"
	"time"
)

func CreateVacancyMock() (*VacancyPostgres, sqlmock.Sqlmock, error) {
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

	repoVacancy := NewVacancyPostgres(gormDB)
	return repoVacancy, mock, nil
}

func TestVacancyPostgres_GetById(t *testing.T) {
	t.Parallel()
	ResumeDB, mock, mockErr := CreateVacancyMock()
	if mockErr != nil {
		t.Errorf("error with creating mock: %s", mockErr)
	}

	testTable := []struct {
		name            string
		expectedErr     error
		expectedVacancy *models.Vacancy
	}{
		{
			name:        "ok",
			expectedErr: nil,
			expectedVacancy: &models.Vacancy{
				ID:          1,
				Title:       "title",
				Description: "desc",
			},
		},
	}

	for _, tc := range testTable {
		testCase := tc
		t.Run(testCase.name, func(t *testing.T) {
			t.Parallel()
			vacancyRows := sqlmock.NewRows([]string{"id", "title", "description"})
			vacancyRows.AddRow(testCase.expectedVacancy.ID, testCase.expectedVacancy.Title,
				testCase.expectedVacancy.Description)
			mock.
				ExpectQuery(regexp.QuoteMeta(`SELECT * FROM "vacancies" WHERE id = $1`)).
				WithArgs(1).WillReturnRows(vacancyRows)

			actualResume, getErr := ResumeDB.GetById(1)

			assert.Equal(t, testCase.expectedErr, getErr)
			assert.Equal(t, testCase.expectedVacancy, actualResume)
			if err := mock.ExpectationsWereMet(); err != nil {
				t.Errorf("there were unfulfilled expectations: %s", err)
			}
		})
	}
}

func TestVacancyPostgres_GetPreviewVacanciesByEmployer(t *testing.T) {
	t.Parallel()
	VacancyDB, mock, mockErr := CreateVacancyMock()
	if mockErr != nil {
		t.Errorf("error with creating mock: %s", mockErr)
	}

	testTable := []struct {
		name        string
		id          uint
		expectedErr error
		expected    []*models.VacancyPreview
	}{
		{
			name:        "ok",
			id:          1,
			expectedErr: nil,
			expected: []*models.VacancyPreview{{
				Id:    1,
				Title: "Some title",
			},
			},
		},
		{
			name:        "not found",
			id:          5,
			expectedErr: errorHandler.ErrVacancyNotFound,
			expected:    []*models.VacancyPreview{},
		},
	}

	for _, tc := range testTable {
		testCase := tc
		t.Run(testCase.name, func(t *testing.T) {

			vacancyRows := sqlmock.NewRows([]string{"id", "title"})
			if len(testCase.expected) > 0 {
				vacancyRows.AddRow(testCase.expected[0].Id, testCase.expected[0].Title)
			}
			mock.
				ExpectQuery(regexp.QuoteMeta(`SELECT user_accounts.image,vacancies.id, vacancies.title, vacancies.salary, vacancies.location, vacancies.format, vacancies.hours, vacancies.description FROM "vacancies" left join user_accounts on vacancies.posted_by_user_id = user_accounts.id WHERE posted_by_user_id = $1`)).
				WithArgs(1).WillReturnRows(vacancyRows)

			actualResume, getErr := VacancyDB.GetPreviewVacanciesByEmployer(1)

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

func TestVacancyPostgres_GetByUserId(t *testing.T) {
	t.Parallel()
	VacancyDB, mock, mockErr := CreateVacancyMock()
	if mockErr != nil {
		t.Errorf("error with creating mock: %s", mockErr)
	}

	testTable := []struct {
		name        string
		id          uint
		expectedErr error
		expected    []*models.Vacancy
	}{
		{
			name:        "ok",
			id:          1,
			expectedErr: nil,
			expected: []*models.Vacancy{{
				ID:          1,
				Title:       "title",
				Description: "desc",
			}},
		},
		{
			name:        "not found",
			id:          5,
			expectedErr: errorHandler.ErrVacancyNotFound,
			expected:    []*models.Vacancy{},
		},
	}

	for _, tc := range testTable {
		testCase := tc
		t.Run(testCase.name, func(t *testing.T) {

			vacancyRows := sqlmock.NewRows([]string{"id", "title", "description"})
			if len(testCase.expected) > 0 {
				vacancyRows.AddRow(testCase.expected[0].ID, testCase.expected[0].Title,
					testCase.expected[0].Description)
			}
			mock.
				ExpectQuery(regexp.QuoteMeta(`SELECT * FROM "vacancies" WHERE posted_by_user_id = $1`)).
				WithArgs(1).WillReturnRows(vacancyRows)

			actualVacancies, getErr := VacancyDB.GetByUserId(1)

			assert.Equal(t, testCase.expectedErr, getErr)
			assert.Equal(t, len(testCase.expected), len(actualVacancies))
			for i := range actualVacancies {
				assert.Equal(t, *testCase.expected[i], *actualVacancies[i])
			}

			if err := mock.ExpectationsWereMet(); err != nil {
				t.Errorf("there were unfulfilled expectations: %s", err)
			}
		})
	}
}

func TestResumePostgres_DeleteVacancy(t *testing.T) {
	t.Parallel()
	VacancyDB, mock, mockErr := CreateVacancyMock()
	if mockErr != nil {
		t.Errorf("error with creating mock: %s", mockErr)
	}

	testTable := []struct {
		name        string
		id          uint
		expectedErr error
		expected    []*models.Vacancy
	}{
		{
			name:        "ok",
			id:          1,
			expectedErr: nil,
			expected: []*models.Vacancy{{
				ID:          1,
				Title:       "title",
				Description: "desc",
			}},
		},
	}

	for _, tc := range testTable {
		testCase := tc
		t.Run(testCase.name, func(t *testing.T) {

			vacancyRows := sqlmock.NewRows([]string{"id", "title", "description"})
			if len(testCase.expected) > 0 {
				vacancyRows.AddRow(testCase.expected[0].ID, testCase.expected[0].Title,
					testCase.expected[0].Description)
			}
			mock.ExpectBegin()
			if testCase.expectedErr == nil {
				mock.
					ExpectExec(regexp.QuoteMeta(`DELETE FROM "vacancies" WHERE posted_by_user_id = $1 AND "vacancies"."id" = $2`)).
					WithArgs(1, 1).WillReturnResult(sqlmock.NewResult(1, 1))
			} else {
				mock.
					ExpectExec(regexp.QuoteMeta(`DELETE FROM "vacancies" WHERE posted_by_user_id = $1 AND "vacancies"."id" = $2`)).
					WithArgs(1, 1).WillReturnError(fmt.Errorf("%w", errorHandler.ErrVacancyNotFound))
			}
			mock.ExpectCommit()

			deleteErr := VacancyDB.Delete(1, 1)
			assert.Equal(t, testCase.expectedErr, deleteErr)
			if err := mock.ExpectationsWereMet(); err != nil {
				t.Errorf("there were unfulfilled expectations: %s", err)
			}
		})
	}
}

func TestResumePostgres_CreateVacancy(t *testing.T) {
	t.Parallel()
	VacancyDB, mock, mockErr := CreateVacancyMock()
	if mockErr != nil {
		t.Errorf("error with creating mock: %s", mockErr)
	}

	testTable := []struct {
		name        string
		id          uint
		expectedErr error
		expected    []*models.Vacancy
	}{
		{
			name:        "ok",
			id:          1,
			expectedErr: nil,
			expected: []*models.Vacancy{{
				ID:             1,
				PostedByUserId: 1,
				Title:          "title",
				Description:    "desc",
				Location:       "loc",
				Salary:         1,
			}},
		},
	}

	for _, tc := range testTable {
		testCase := tc
		t.Run(testCase.name, func(t *testing.T) {

			vacancyRows := sqlmock.NewRows([]string{"id", "posted_by_user_id", "title", "description", "tasks", "requirements", "extra", "created_date", "salary", "location", "is_active", "experience", "format", "hours", "image"})
			if len(testCase.expected) > 0 {
				vacancyRows.AddRow(testCase.expected[0].ID, testCase.expected[0].PostedByUserId, testCase.expected[0].Title,
					testCase.expected[0].Description, testCase.expected[0].Tasks, testCase.expected[0].Requirements, testCase.expected[0].Extra,
					testCase.expected[0].CreatedDate, testCase.expected[0].Salary, testCase.expected[0].Location, testCase.expected[0].IsActive, testCase.expected[0].Experience, testCase.expected[0].Format, testCase.expected[0].Hours, testCase.expected[0].Image)
			}

			mock.
				ExpectQuery(regexp.QuoteMeta(`SELECT * FROM "user_accounts" WHERE id = $1`)).
				WithArgs(1).WillReturnRows(vacancyRows)

			timeNow := time.Now()

			mock.ExpectBegin()
			mock.
				ExpectQuery(regexp.QuoteMeta(`INSERT INTO "vacancies" ("posted_by_user_id","title","description","tasks","requirements","extra","created_date","salary","location","is_active","experience","format","hours","image") VALUES ($1,$2,$3,$4,$5,$6,$7,$8,$9,$10,$11,$12,$13,$14) RETURNING "id"`)).
				WithArgs(1,
					"title",
					"desc",
					"tasks",
					"requere",
					"extra",
					timeNow,
					100,
					"loc",
					true,
					"exp",
					"format",
					"hoyrs",
					"",
				).
				WillReturnRows(sqlmock.NewRows([]string{"1"}))
			mock.ExpectCommit()

			vacancy := &models.Vacancy{
				PostedByUserId: 1,
				Title:          "title",
				Description:    "desc",
				CreatedDate:    timeNow,
				Location:       "loc",
				Salary:         100,
				Tasks:          "tasks",
				Requirements:   "requere",
				Extra:          "extra",
				IsActive:       true,
				Experience:     "exp",
				Format:         "format",
				Hours:          "hoyrs",
				Image:          "",
			}

			_, createErr := VacancyDB.Create(vacancy)
			assert.Equal(t, testCase.expectedErr, createErr)

			if err := mock.ExpectationsWereMet(); err != nil {
				t.Errorf("there were unfulfilled expectations: %s", err)
			}
		})
	}
}
