package postgres

import (
	"database/sql/driver"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/bxcodec/faker"
	"github.com/jmoiron/sqlx"
	"github.com/stretchr/testify/require"
)

type anyPassword struct{}

func (a anyPassword) Match(v driver.Value) bool {
	_, ok := v.(string)
	return ok
}

func TestPostgresUserRepository_InsertLike(t *testing.T) {
	type insertLikeTestCase struct {
		uid1 int
		uid2 int
		err  error
	}

	db, mock, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
	if err != nil {
		t.Fatalf("error '%s' when opening a stub database connection", err)
	}
	defer db.Close()
	sqlxDB := sqlx.NewDb(db, "sqlmock")

	columns := []string{
		"id",
		"user_id1",
		"user_id2",
	}

	query := `INSERT INTO likes(user_id1, user_id2) VALUES ($1, $2);`

	var uid1, uid2 int

	err = faker.FakeData(&uid1)
	require.NoError(t, err)

	err = faker.FakeData(&uid2)
	require.NoError(t, err)

	testCases := []insertLikeTestCase{
		{
			uid1: uid1,
			uid2: uid2,
			err:  nil,
		},
	}

	for _, testCase := range testCases {
		args := []driver.Value{
			testCase.uid1,
			testCase.uid2,
		}

		rows := []driver.Value{
			1,
			testCase.uid1,
			testCase.uid2,
		}

		sqlmock.NewRows(columns).AddRow(rows...)
		mock.ExpectExec(query).WithArgs(args...).WillReturnResult(sqlmock.NewResult(1, 1))

		repo := NewPostgresUserRepository(sqlxDB.DB)

		err = repo.InsertLike(testCase.uid1, testCase.uid2)
		require.Equal(t, testCase.err, err)

		err = mock.ExpectationsWereMet()
		require.NoError(t, err, "unfulfilled expectations: %s", err)
	}
}

func TestPostgresUserRepository_InsertDisLike(t *testing.T) {
	type insertDisLikeTestCase struct {
		uid1 int
		uid2 int
		err  error
	}

	db, mock, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
	if err != nil {
		t.Fatalf("error '%s' when opening a stub database connection", err)
	}
	defer db.Close()
	sqlxDB := sqlx.NewDb(db, "sqlmock")

	columns := []string{
		"id",
		"user_id1",
		"user_id2",
	}

	query := `INSERT INTO dislikes(user_id1, user_id2) VALUES ($1, $2);`

	var uid1, uid2 int

	err = faker.FakeData(&uid1)
	require.NoError(t, err)

	err = faker.FakeData(&uid2)
	require.NoError(t, err)

	testCases := []insertDisLikeTestCase{
		{
			uid1: uid1,
			uid2: uid2,
			err:  nil,
		},
	}

	for _, testCase := range testCases {
		args := []driver.Value{
			testCase.uid1,
			testCase.uid2,
		}

		rows := []driver.Value{
			1,
			testCase.uid1,
			testCase.uid2,
		}

		sqlmock.NewRows(columns).AddRow(rows...)
		mock.ExpectExec(query).WithArgs(args...).WillReturnResult(sqlmock.NewResult(1, 1))

		repo := NewPostgresUserRepository(sqlxDB.DB)

		err = repo.InsertDislike(testCase.uid1, testCase.uid2)
		require.Equal(t, testCase.err, err)

		err = mock.ExpectationsWereMet()
		require.NoError(t, err, "unfulfilled expectations: %s", err)
	}
}

type anyTime struct{}

func (a anyTime) Match(v driver.Value) bool {
	_, ok := v.(string)
	return ok
}
