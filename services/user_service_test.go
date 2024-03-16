package services

import (
	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"testing"
)

func TestVerifyUserEmail(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("An error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	// init
	mock.ExpectQuery("SELECT VERSION()").WillReturnRows(sqlmock.NewRows([]string{"version"}).AddRow("5.7.22"))
	// query user
	rows := sqlmock.NewRows([]string{"id", "email", "verification_code", "verified"}).
		AddRow(1, "aa@aa.com", "1234", false)
	mock.ExpectQuery("SELECT \\* FROM `users` WHERE email = \\? AND verification_code = \\? ORDER BY `users`.`id` LIMIT \\?").
		WithArgs("aa@aa.com", "1234", 1).
		WillReturnRows(rows)

	// update user
	mock.ExpectBegin()
	mock.ExpectExec("UPDATE `users` SET `email`=\\?,`password`=\\?,`verified`=\\?,`verification_code`=\\?,`verification_code_expiry`=\\?,`verification_at`=\\? WHERE `id` = \\?").
		WithArgs("aa@aa.com", "", true, "1234", sqlmock.AnyArg(), sqlmock.AnyArg(), 1).
		WillReturnResult(sqlmock.NewResult(1, 1))
	mock.ExpectCommit()

	gormDB, err := gorm.Open(mysql.New(mysql.Config{
		Conn: db,
	}), &gorm.Config{})
	if err != nil {
		t.Fatalf("An error '%s' was not expected when opening a gorm database connection", err)
	}
	userService := UserService{DB: gormDB}
	resultUser, err := userService.VerifyEmail("aa@aa.com", "1234")

	assert.NoError(t, err)
	assert.NotNil(t, resultUser)
	assert.True(t, resultUser.Verified)
	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("There were unfulfilled expectations: %s", err)
	}
}
