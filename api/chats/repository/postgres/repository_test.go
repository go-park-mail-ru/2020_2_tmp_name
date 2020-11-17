package postgres

import (
	"database/sql"

	"database/sql/driver"
	"math/rand"
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

	query := `SELECT text, time_delivery, user_id FROM message WHERE chat_id=$1 order by id asc limit 10;`

	var messages []models.Msg
	err = faker.FakeData(&messages)
	require.NoError(t, err)

	testCases := []insertMessageTestCase{
		{
			messages: messages,
			chatId:   messages[1+rand.Int()%len(messages)].ChatID,
			err:      sql.ErrNoRows,
		},
		{
			messages: messages,
			chatId:   messages[1+rand.Int()%len(messages)].ChatID,
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

	query := `SELECT text, time_delivery, user_id FROM message WHERE user_id=$1 AND chat_id=$2 order by id desc limit 1;`

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
			mock.ExpectQuery(query).WithArgs(testCase.messages.UserID, testCase.messages.ChatID).WillReturnError(testCase.err)
		} else {
			data := []driver.Value{
				testCase.messages.Message,
				testCase.messages.TimeDelivery,
				testCase.messages.UserID,
			}
			rows := sqlmock.NewRows(columns).AddRow(data...)
			mock.ExpectQuery(query).WithArgs(testCase.messages.UserID, testCase.messages.ChatID).WillReturnRows(rows)
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
