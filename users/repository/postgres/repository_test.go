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

type anyPassword struct{}
func (a anyPassword) Match(v driver.Value) bool {
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
			anyPassword{},
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

func TestPostgresUserRepository_InsertLike(t *testing.T) {
	type insertLikeTestCase struct {
		uid1 int
		uid2 int
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
			err: nil,
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
			err: nil,
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
func TestPostgresUserRepository_InsertComment(t *testing.T) {
	type insertCommentTestCase struct {
		comment models.Comment
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
			err: nil,
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

		repo := NewPostgresUserRepository(sqlxDB.DB)

		err = repo.InsertComment(testCase.comment, testCase.comment.Uid1)
		require.Equal(t, testCase.err, err)

		err = mock.ExpectationsWereMet()
		require.NoError(t, err, "unfulfilled expectations: %s", err)
	}
}

func TestPostgresUserRepository_InsertSession(t *testing.T) {
	type insertSessionTestCase struct {
		 key string
		 value string
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
		"key",
		"value",
	}

	query := `INSERT INTO sessions(key, value) VALUES ($1, $2);`

	var key, value string
	err = faker.FakeData(&key)
	require.NoError(t, err)
	err = faker.FakeData(&value)
	require.NoError(t, err)

	testCases := []insertSessionTestCase{
		{
			key: key,
			value: value,
			err: nil,
		},
	}

	for _, testCase := range testCases {
		args := []driver.Value{
			testCase.key,
			testCase.value,
		}

		rows := []driver.Value{
			1,
			testCase.key,
			testCase.value,
		}

		sqlmock.NewRows(columns).AddRow(rows...)
		mock.ExpectExec(query).WithArgs(args...).WillReturnResult(sqlmock.NewResult(1, 1))

		repo := NewPostgresUserRepository(sqlxDB.DB)

		err = repo.InsertSession(testCase.key, testCase.value)
		require.Equal(t, testCase.err, err)

		err = mock.ExpectationsWereMet()
		require.NoError(t, err, "unfulfilled expectations: %s", err)
	}
}

func TestPostgresUserRepository_InsertChat(t *testing.T) {
	type insertChatTestCase struct {
		chat models.Chat
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
		"user_id1",
		"user_id2",
	}

	query := `INSERT INTO chat(user_id1, user_id2) VALUES ($1, $2);`

	var chat models.Chat
	err = faker.FakeData(&chat)
	require.NoError(t, err)

	testCases := []insertChatTestCase{
		{
			chat: chat,
			err: nil,
		},
	}

	for _, testCase := range testCases {
		args := []driver.Value{
			testCase.chat.Uid1,
			testCase.chat.Uid2,
		}

		rows := []driver.Value{
			testCase.chat.ID,
			testCase.chat.Uid1,
			testCase.chat.Uid2,
		}

		sqlmock.NewRows(columns).AddRow(rows...)
		mock.ExpectExec(query).WithArgs(args...).WillReturnResult(sqlmock.NewResult(1, 1))

		repo := NewPostgresUserRepository(sqlxDB.DB)

		err = repo.InsertChat(testCase.chat)
		require.Equal(t, testCase.err, err)

		err = mock.ExpectationsWereMet()
		require.NoError(t, err, "unfulfilled expectations: %s", err)
	}
}

func TestPostgresUserRepository_InsertMessage(t *testing.T) {
	type insertMessageTestCase struct {
		message models.Message
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
		"text",
		"time_delivery",
		"chat_id",
		"user_id",
	}

	query := `INSERT INTO message(text, time_delivery, chat_id, user_id) VALUES ($1, $2, $3, $4);`

	var message models.Message
	err = faker.FakeData(&message)
	require.NoError(t, err)

	testCases := []insertMessageTestCase{
		{
			message: message,
			err: nil,
		},
	}

	for _, testCase := range testCases {
		args := []driver.Value{
			testCase.message.Text,
			anyTime{},
			testCase.message.ChatID,
			testCase.message.UserID,
		}

		rows := []driver.Value{
			testCase.message.ID,
			testCase.message.Text,
			testCase.message.TimeDelivery,
			testCase.message.ChatID,
			testCase.message.UserID,
		}

		sqlmock.NewRows(columns).AddRow(rows...)
		mock.ExpectExec(query).WithArgs(args...).WillReturnResult(sqlmock.NewResult(1, 1))

		repo := NewPostgresUserRepository(sqlxDB.DB)

		err = repo.InsertMessage(testCase.message.Text, testCase.message.ChatID, testCase.message.UserID)
		require.Equal(t, testCase.err, err)

		err = mock.ExpectationsWereMet()
		require.NoError(t, err, "unfulfilled expectations: %s", err)
	}
}

func TestPostgresUserRepository_InsertPhoto(t *testing.T) {
	type insertPhotoTestCase struct {
		uid int
		path string
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
		"path",
		"user_id",
	}

	query := `INSERT INTO photo(path, user_id) VALUES ($1, $2);`

	var uid int
	err = faker.FakeData(&uid)
	require.NoError(t, err)

	var path string
	err = faker.FakeData(&path)
	require.NoError(t, err)

	testCases := []insertPhotoTestCase{
		{
			uid: uid,
			path: path,
			err: nil,
		},
	}

	for _, testCase := range testCases {
		args := []driver.Value{
			testCase.path,
			testCase.uid,
		}

		rows := []driver.Value{
			1,
			testCase.path,
			testCase.uid,
		}

		sqlmock.NewRows(columns).AddRow(rows...)
		mock.ExpectExec(query).WithArgs(args...).WillReturnResult(sqlmock.NewResult(1, 1))

		repo := NewPostgresUserRepository(sqlxDB.DB)

		err = repo.InsertPhoto(testCase.path, testCase.uid)
		require.Equal(t, testCase.err, err)

		err = mock.ExpectationsWereMet()
		require.NoError(t, err, "unfulfilled expectations: %s", err)
	}
}

func TestPostgresUserRepository_UpdateUser(t *testing.T) {

}