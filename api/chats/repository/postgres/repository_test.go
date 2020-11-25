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
func TestPostgresChatRepository_InsertChat(t *testing.T) {
	type insertChatTestCase struct {
		chat models.Chat
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

	query := `INSERT INTO chat(user_id1, user_id2) VALUES ($1, $2);`

	var chat models.Chat
	err = faker.FakeData(&chat)
	require.NoError(t, err)

	testCases := []insertChatTestCase{
		{
			chat: chat,
			err:  nil,
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

		repo := NewPostgresChatRepository(sqlxDB.DB)

		err = repo.InsertChat(testCase.chat)
		require.Equal(t, testCase.err, err)

		err = mock.ExpectationsWereMet()
		require.NoError(t, err, "unfulfilled expectations: %s", err)
	}
}

func TestPostgresChatRepository_InsertMessage(t *testing.T) {
	type insertMessageTestCase struct {
		message models.Message
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
			err:     nil,
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

		repo := NewPostgresChatRepository(sqlxDB.DB)

		err = repo.InsertMessage(testCase.message.Text, testCase.message.ChatID, testCase.message.UserID)
		require.Equal(t, testCase.err, err)

		err = mock.ExpectationsWereMet()
		require.NoError(t, err, "unfulfilled expectations: %s", err)
	}
}

func TestPostgresChatRepository_SelectMessages(t *testing.T) {
	type insertMessageTestCase struct {
		messages []models.Msg
		chatId   int
		err      error
	}

	db, mock, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
	if err != nil {
		t.Fatalf("error '%s' when opening a stub database connection", err)
	}
	defer db.Close()
	sqlxDB := sqlx.NewDb(db, "sqlmock")

	columns := []string{
		"text",
		"time_delivery",
		"user_id",
	}

	query := `SELECT m.text, m.time_delivery, m.user_id FROM (SELECT * FROM message WHERE chat_id=$1 ORDER BY id DESC limit 10) AS m ORDER BY m.id ASC;`

	var messages []models.Msg
	err = faker.FakeData(&messages)
	require.NoError(t, err)

	testCases := []insertMessageTestCase{
		{
			messages: messages,
			chatId:   1,
			err:      sql.ErrNoRows,
		},
		{
			messages: messages,
			chatId:   1,
			err:      nil,
		},
	}

	for _, testCase := range testCases {
		for i := 0; i < len(testCase.messages); i++ {
			testCase.messages[i].ChatID = testCase.chatId
		}
		if testCase.err != nil {
			mock.ExpectQuery(query).WithArgs(testCase.chatId).WillReturnError(testCase.err)
		} else {
			rows := sqlmock.NewRows(columns)
			for _, message := range testCase.messages {
				rows.AddRow(message.Message, message.TimeDelivery, message.UserID)
			}
			mock.ExpectQuery(query).WithArgs(testCase.chatId).WillReturnRows(rows)
		}

		repo := NewPostgresChatRepository(sqlxDB.DB)

		msgs, err := repo.SelectMessages(testCase.chatId)
		require.Equal(t, testCase.err, err)
		if err == nil {
			require.Equal(t, testCase.messages, msgs)
		}

		err = mock.ExpectationsWereMet()
		require.NoError(t, err, "unfulfilled expectations: %s", err)
	}
}

func TestPostgresChatRepository_SelectMessage(t *testing.T) {
	type insertMessageTestCase struct {
		messages models.Msg
		err      error
	}

	db, mock, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
	if err != nil {
		t.Fatalf("error '%s' when opening a stub database connection", err)
	}
	defer db.Close()
	sqlxDB := sqlx.NewDb(db, "sqlmock")

	columns := []string{
		"text",
		"time_delivery",
		"user_id",
	}

	query := `SELECT text, time_delivery, user_id FROM message WHERE chat_id=$1 order by id desc limit 1;`

	var messages models.Msg
	err = faker.FakeData(&messages)
	require.NoError(t, err)

	testCases := []insertMessageTestCase{
		{
			messages: messages,
			err:      nil,
		},
	}

	for _, testCase := range testCases {
		if testCase.err != nil {
			mock.ExpectQuery(query).WithArgs(testCase.messages.ChatID).WillReturnError(testCase.err)
		} else {
			data := []driver.Value{
				testCase.messages.Message,
				testCase.messages.TimeDelivery,
				testCase.messages.UserID,
			}
			rows := sqlmock.NewRows(columns).AddRow(data...)
			mock.ExpectQuery(query).WithArgs(testCase.messages.ChatID).WillReturnRows(rows)
		}

		repo := NewPostgresChatRepository(sqlxDB.DB)

		msg, err := repo.SelectMessage(testCase.messages.UserID, testCase.messages.ChatID)
		require.Equal(t, testCase.err, err)
		if err == nil {
			require.Equal(t, testCase.messages, msg)
		}

		err = mock.ExpectationsWereMet()
		require.NoError(t, err, "unfulfilled expectations: %s", err)
	}
}

func TestPostgresChatRepository_SelectImages(t *testing.T) {
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

		repo := NewPostgresChatRepository(sqlxDB.DB)

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

func TestPostgresChatRepository_SelectUser(t *testing.T) {
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

		repo := NewPostgresChatRepository(sqlxDB.DB)

		user, err := repo.SelectUser(testCase.outputUser.Telephone)
		require.Equal(t, testCase.err, err)
		if err == nil {
			require.Equal(t, testCase.outputUser, user)
		}

		err = dbMock.ExpectationsWereMet()
		require.NoError(t, err, "unfulfilled expectations: %s\n%s", err, msg)
	}
}

func TestPostgresChatRepository_SelectUserFeed(t *testing.T) {
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

		repo := NewPostgresChatRepository(sqlxDB.DB)

		user, err := repo.SelectUserFeed(telephone)
		require.Equal(t, testCase.err, err)
		if err == nil {
			require.Equal(t, testCase.outputUser, user)
		}

		err = dbMock.ExpectationsWereMet()
		require.NoError(t, err, "unfulfilled expectations: %s\n%s", err, msg)
	}
}

func TestPostgresChatRepository_SelectUserFeedByID(t *testing.T) {
	type insertUserTestCase struct {
		id         int
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
		"name",
		"date_birth",
		"job",
		"education",
		"about_me",
	}

	query := `SELECT name, date_birth, job, education, about_me FROM users
			  WHERE id=$1;`

	var telephone string
	err = faker.FakeData(&telephone)
	require.NoError(t, err)

	var outputUser models.UserFeed
	err = faker.FakeData(&outputUser)
	require.NoError(t, err)

	testCases := []insertUserTestCase{
		{
			id:         1,
			outputUser: outputUser,
			err:        sql.ErrNoRows,
		},
		{
			id:         1,
			outputUser: outputUser,
			err:        nil,
		},
	}

	for i, testCase := range testCases {
		msg := fmt.Sprintf("case %d aaaaaaaaaaaa", i)
		data := []driver.Value{
			testCase.outputUser.Name,
			testCase.outputUser.DateBirth,
			testCase.outputUser.Job,
			testCase.outputUser.Education,
			testCase.outputUser.AboutMe,
		}

		if testCase.err == nil {
			rows := sqlmock.NewRows(columns).AddRow(data...)
			dbMock.ExpectQuery(query).WithArgs(outputUser.ID).WillReturnRows(rows)

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
			dbMock.ExpectQuery(query).WithArgs(outputUser.ID).WillReturnError(testCase.err)
		}

		repo := NewPostgresChatRepository(sqlxDB.DB)

		user, err := repo.SelectUserFeedByID(testCase.outputUser.ID)
		require.Equal(t, testCase.err, err)
		if err == nil {
			require.Equal(t, testCase.outputUser, user)
		}

		err = dbMock.ExpectationsWereMet()
		require.NoError(t, err, "unfulfilled expectations: %s\n%s", err, msg)
	}
}
