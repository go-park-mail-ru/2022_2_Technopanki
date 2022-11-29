package impl

import (
	"fmt"
	"regexp"
	"testing"

	"HeadHunter/internal/entity/models"

	"gopkg.in/DATA-DOG/go-sqlmock.v1"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func CreateUserMock() (*UserPostgres, sqlmock.Sqlmock, error) {
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

	userRepo := NewUserPostgres(gormDB)
	return userRepo, mock, nil
}

func TestUserPostgres_CreateUser(t *testing.T) {
	UserDB, mock, mockErr := CreateUserMock()
	if mockErr != nil {
		t.Errorf("error with creating mock: %s", mockErr)
	}

	user := models.UserAccount{UserType: "applicant", ApplicantName: "Zakhar",
		ApplicantSurname: "Urvancev"}

	mock.ExpectBegin()
	mock.
		ExpectQuery(regexp.QuoteMeta(`INSERT INTO "user_accounts" ("user_type","email","password","contact_number","status","description","image","date_of_birth","applicant_name","applicant_surname","applicant_current_salary","company_name","business_type","company_website_url","location","company_size","public_fields") VALUES ($1,$2,$3,$4,$5,$6,$7,$8,$9,$10,$11,$12,$13,$14,$15,$16,$17)`)).
		WithArgs(
			user.UserType,
			user.Email,
			user.Password,
			user.ContactNumber,
			user.Status,
			user.Description,
			user.Image,
			user.DateOfBirth,
			user.ApplicantName,
			user.ApplicantSurname,
			user.ApplicantCurrentSalary,
			user.CompanyName,
			user.BusinessType,
			user.CompanyWebsiteUrl,
			user.Location,
			user.CompanySize,
			user.PublicFields).
		WillReturnRows(sqlmock.NewRows([]string{"1"}))
	mock.ExpectCommit()

	err := UserDB.CreateUser(&user)

	if err != nil {
		t.Errorf("unexpected err: %s", err)
		return
	}

	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
	}

	mock.ExpectBegin()
	mock.
		ExpectQuery(regexp.QuoteMeta(`INSERT INTO "user_accounts" ("user_type","email","password","contact_number","status","description","image","date_of_birth","applicant_name","applicant_surname","applicant_current_salary","company_name","business_type","company_website_url","location","company_size","public_fields") VALUES ($1,$2,$3,$4,$5,$6,$7,$8,$9,$10,$11,$12,$13,$14,$15,$16,$17)`)).
		WithArgs(
			user.UserType,
			user.Email,
			user.Password,
			user.ContactNumber,
			user.Status,
			user.Description,
			user.Image,
			user.DateOfBirth,
			user.ApplicantName,
			user.ApplicantSurname,
			user.ApplicantCurrentSalary,
			user.CompanyName,
			user.BusinessType,
			user.CompanyWebsiteUrl,
			user.Location,
			user.CompanySize,
			user.PublicFields).
		WillReturnError(fmt.Errorf("bad_result"))
	mock.ExpectRollback()

	err = UserDB.CreateUser(&user)

	if err == nil {
		t.Errorf("expected error, got nil")
		return
	}

	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
	}
}
