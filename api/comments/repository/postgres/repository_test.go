package postgres

import (
	"database/sql"
	"fmt"

	"database/sql/driver"
	"park_2020/2020_2_tmp_name/models"
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

type anyTime struct{}

func (a anyTime) Match(v driver.Value) bool {
	_, ok := v.(string)
	return ok
}
func TestPostgresCommentRepository_InsertComment(t *testing.T) {
	type insertCommentTestCase struct {
		comment models.Comment
		err     error
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
		"time_delivery",
		"text",
	}

	query := `INSERT INTO comments(user_id1, user_id2, time_delivery, text) VALUES ($1, $2, $3, $4);`

	var comment models.Comment

	err = faker.FakeData(&comment)
	require.NoError(t, err)

	testCases := []insertCommentTestCase{
		{
			comment: comment,
			err:     nil,
		},
	}

	for _, testCase := range testCases {
		args := []driver.Value{
			comment.Uid1,
			comment.Uid2,
			anyTime{},
			comment.CommentText,
		}

		rows := []driver.Value{
			comment.ID,
			comment.Uid1,
			comment.Uid2,
			comment.TimeDelivery,
			comment.CommentText,
		}

		sqlmock.NewRows(columns).AddRow(rows...)
		mock.ExpectExec(query).WithArgs(args...).WillReturnResult(sqlmock.NewResult(1, 1))

		repo := NewPostgresCommentRepository(sqlxDB.DB)

		err = repo.InsertComment(testCase.comment, testCase.comment.Uid1)
		require.Equal(t, testCase.err, err)

		err = mock.ExpectationsWereMet()
		require.NoError(t, err, "unfulfilled expectations: %s", err)
	}
}

func TestPostgresCommentRepository_SelectImages(t *testing.T) {
	type insertPhotoTestCase struct {
		uid  int
		path []string
		err  error
	}

	db, mock, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
	if err != nil {
		t.Fatalf("error '%s' when opening a stub database connection", err)
	}
	defer db.Close()
	sqlxDB := sqlx.NewDb(db, "sqlmock")

	columns := []string{
		"path",
	}

	query := `SELECT path FROM photo WHERE user_id=$1;`

	var uid int
	err = faker.FakeData(&uid)
	require.NoError(t, err)

	var path []string
	err = faker.FakeData(&path)
	require.NoError(t, err)

	testCases := []insertPhotoTestCase{
		{
			uid:  uid,
			path: path,
			err:  sql.ErrNoRows,
		},
		{
			uid:  uid,
			path: path,
			err:  nil,
		},
	}

	for _, testCase := range testCases {
		if testCase.err != nil {
			mock.ExpectQuery(query).WithArgs(testCase.uid).WillReturnError(testCase.err)
		} else {
			rows := sqlmock.NewRows(columns)
			for _, image := range testCase.path {
				rows.AddRow(image)
			}
			mock.ExpectQuery(query).WithArgs(testCase.uid).WillReturnRows(rows)
		}

		repo := NewPostgresCommentRepository(sqlxDB.DB)

		images, err := repo.SelectImages(testCase.uid)
		t.Log(images)
		require.Equal(t, testCase.err, err)
		if err == nil {
			require.Equal(t, testCase.path, images)
		}

		err = mock.ExpectationsWereMet()
		require.NoError(t, err, "unfulfilled expectations: %s", err)
	}
}

func TestPostgresCommentRepository_SelectUser(t *testing.T) {
	type insertUserTestCase struct {
		telephone  string
		outputUser models.User
		err        error
	}

	db, dbMock, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
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

	query := `SELECT id, name, telephone, password, date_birth, sex, job, education, about_me FROM users
			  WHERE  telephone=$1;`

	var telephone string
	err = faker.FakeData(&telephone)
	require.NoError(t, err)

	var outputUser models.User
	err = faker.FakeData(&outputUser)
	require.NoError(t, err)

	outputUser.Day = ""
	outputUser.Month = ""
	outputUser.Year = ""
	outputUser.DateBirth = 19
	outputUser.Telephone = telephone

	testCases := []insertUserTestCase{
		{
			telephone:  "telephone",
			outputUser: outputUser,
			err:        sql.ErrNoRows,
		},
		{
			telephone:  telephone,
			outputUser: outputUser,
			err:        nil,
		},
	}

	for i, testCase := range testCases {
		msg := fmt.Sprintf("case %d aaaaaaaaaaaa", i)
		data := []driver.Value{
			testCase.outputUser.ID,
			testCase.outputUser.Name,
			testCase.outputUser.Telephone,
			testCase.outputUser.Password,
			testCase.outputUser.DateBirth,
			testCase.outputUser.Sex,
			testCase.outputUser.Education,
			testCase.outputUser.Job,
			testCase.outputUser.AboutMe,
		}

		if testCase.err == nil {
			rows := sqlmock.NewRows(columns).AddRow(data...)
			dbMock.ExpectQuery(query).WithArgs(outputUser.Telephone).WillReturnRows(rows)

			subQuery := `SELECT path FROM photo WHERE user_id=$1;`
			subColumns := []string{
				"path",
			}
			subRows := sqlmock.NewRows(subColumns)
			for _, img := range testCase.outputUser.LinkImages {
				subRows.AddRow(img)
			}
			dbMock.ExpectQuery(subQuery).WithArgs(testCase.outputUser.ID).WillReturnRows(subRows)
		} else {
			dbMock.ExpectQuery(query).WithArgs(outputUser.Telephone).WillReturnError(testCase.err)
		}

		repo := NewPostgresCommentRepository(sqlxDB.DB)

		user, err := repo.SelectUser(testCase.outputUser.Telephone)
		require.Equal(t, testCase.err, err)
		if err == nil {
			require.Equal(t, testCase.outputUser, user)
		}

		err = dbMock.ExpectationsWereMet()
		require.NoError(t, err, "unfulfilled expectations: %s\n%s", err, msg)
	}
}

func TestPostgresCommentRepository_SelectUserFeed(t *testing.T) {
	type insertUserTestCase struct {
		telephone  string
		outputUser models.UserFeed
		err        error
	}

	db, dbMock, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
	if err != nil {
		t.Fatalf("error '%s' when opening a stub database connection", err)
	}
	defer db.Close()
	sqlxDB := sqlx.NewDb(db, "sqlmock")

	columns := []string{
		"id",
		"name",
		"date_birth",
		"education",
		"job",
		"about_me",
	}

	query := `SELECT id, name, date_birth, education, job, about_me FROM users
			  WHERE  telephone=$1;`

	var telephone string
	err = faker.FakeData(&telephone)
	require.NoError(t, err)

	var outputUser models.UserFeed
	err = faker.FakeData(&outputUser)
	require.NoError(t, err)

	testCases := []insertUserTestCase{
		{
			telephone:  "telephone",
			outputUser: outputUser,
			err:        sql.ErrNoRows,
		},
		{
			telephone:  telephone,
			outputUser: outputUser,
			err:        nil,
		},
	}

	for i, testCase := range testCases {
		msg := fmt.Sprintf("case %d aaaaaaaaaaaa", i)
		data := []driver.Value{
			testCase.outputUser.ID,
			testCase.outputUser.Name,
			testCase.outputUser.DateBirth,
			testCase.outputUser.Education,
			testCase.outputUser.Job,
			testCase.outputUser.AboutMe,
		}

		if testCase.err == nil {
			rows := sqlmock.NewRows(columns).AddRow(data...)
			dbMock.ExpectQuery(query).WithArgs(telephone).WillReturnRows(rows)

			subQuery := `SELECT path FROM photo WHERE user_id=$1;`
			subColumns := []string{
				"path",
			}
			subRows := sqlmock.NewRows(subColumns)
			for _, img := range testCase.outputUser.LinkImages {
				subRows.AddRow(img)
			}
			dbMock.ExpectQuery(subQuery).WithArgs(testCase.outputUser.ID).WillReturnRows(subRows)
		} else {
			dbMock.ExpectQuery(query).WithArgs(telephone).WillReturnError(testCase.err)
		}

		repo := NewPostgresCommentRepository(sqlxDB.DB)

		user, err := repo.SelectUserFeed(telephone)
		require.Equal(t, testCase.err, err)
		if err == nil {
			require.Equal(t, testCase.outputUser, user)
		}

		err = dbMock.ExpectationsWereMet()
		require.NoError(t, err, "unfulfilled expectations: %s\n%s", err, msg)
	}
}
