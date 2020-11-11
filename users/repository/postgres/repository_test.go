package postgres

import (
	// "database/sql"
	"database/sql/driver"
	"github.com/DATA-DOG/go-sqlmock"
	"github.com/bxcodec/faker"
	"github.com/jmoiron/sqlx"
	"github.com/stretchr/testify/require"
	"park_2020/2020_2_tmp_name/models"
	"testing"
)

type AnyPassword struct{}
func (a AnyPassword) Match(v driver.Value) bool {
	_, ok := v.(string)
	return ok
}

func TestPostgresUserRepository_InsertUser(t *testing.T) {

	type insertUserTestCase struct {
		inputUser models.User
		err error
	}

	db, mock, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
	if err != nil {
		t.Fatalf("error '%s' when opening a stub database connection", err)
	}
	defer db.Close()
	sqlxDB := sqlx.NewDb(db, "sqlmock")

	columns := []string{
		"id",
		"name",
		"telephone",
		"password",
		"date_birth",
		"sex",
		"education",
		"job",
		"about_me",
	}

	query := `INSERT INTO users(name, telephone, password, date_birth, sex, job, education, about_me) 
			  VALUES ($1, $2, $3, $4, $5, $6, $7, $8);`

	var inputUser models.User
	err = faker.FakeData(&inputUser)
	require.NoError(t, err)

	testCases := []insertUserTestCase{
		{
			inputUser: inputUser,
			err: nil,
		},
	}

	for _, testCase := range testCases {
		args := []driver.Value{
			testCase.inputUser.Name,
			testCase.inputUser.Telephone,
			AnyPassword{},
			testCase.inputUser.DateBirth,
			testCase.inputUser.Sex,
			testCase.inputUser.Job,
			testCase.inputUser.Education,
			testCase.inputUser.AboutMe,
		}

		rows := []driver.Value{
			testCase.inputUser.ID,
			testCase.inputUser.Name,
			testCase.inputUser.Telephone,
			testCase.inputUser.Password,
			testCase.inputUser.DateBirth,
			testCase.inputUser.Sex,
			testCase.inputUser.Education,
			testCase.inputUser.Job,
			testCase.inputUser.AboutMe,
		}

		sqlmock.NewRows(columns).AddRow(rows...)
		mock.ExpectExec(query).WithArgs(args...).WillReturnResult(sqlmock.NewResult(int64(testCase.inputUser.ID), 1))

		repo := NewPostgresUserRepository(sqlxDB.DB)

		err = repo.InsertUser(testCase.inputUser)
		require.Equal(t, testCase.err, err)

		err = mock.ExpectationsWereMet()
		require.NoError(t, err, "unfulfilled expectations: %s", err)
	}
}
