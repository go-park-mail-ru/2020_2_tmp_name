package postgres

import (
	"database/sql"
	"fmt"
	"time"

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

func TestPostgresUserRepository_InsertUser(t *testing.T) {
	type insertUserTestCase struct {
		inputUser models.User
		err       error
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

	inputUser.Day = "21"
	inputUser.Month = "Май"
	inputUser.Year = "2001"
	inputUser.DateBirth = 19

	testCases := []insertUserTestCase{
		{
			inputUser: inputUser,
			err:       sql.ErrNoRows,
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
		mock.ExpectExec(query).WithArgs(args...).WillReturnError(testCase.err)

		repo := NewPostgresUserRepository(sqlxDB.DB)

		err = repo.InsertUser(testCase.inputUser)
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

func TestPostgresUserRepository_InsertSession(t *testing.T) {
	type insertSessionTestCase struct {
		key   string
		value string
		err   error
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
			key:   key,
			value: value,
			err:   nil,
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

func TestPostgresUserRepository_UpdateUser(t *testing.T) {
	type updateUserTestCase struct {
		user models.User
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
		"name",
		"telephone",
		"password",
		"date_birth",
		"sex",
		"education",
		"job",
		"about_me",
	}

	var user models.User
	err = faker.FakeData(&user)
	require.NoError(t, err)

	testCases := []updateUserTestCase{
		{
			user: user,
			err:  nil,
		},
	}

	type query struct {
		q    string
		data interface{}
	}

	queries := []query{
		{
			q:    `UPDATE users SET name=$1 WHERE id = $2;`,
			data: user.Name,
		},
		{
			q:    `UPDATE users SET telephone=$1 WHERE id = $2;`,
			data: user.Telephone,
		},
		{
			q:    `UPDATE users SET password=$1 WHERE id = $2;`,
			data: anyPassword{},
		},
		{
			q:    `UPDATE users SET job=$1 WHERE id = $2;`,
			data: user.Job,
		},
		{
			q:    `UPDATE users SET education=$1 WHERE id = $2;`,
			data: user.Education,
		},
		{
			q:    `UPDATE users SET about_me=$1 WHERE id = $2;`,
			data: user.AboutMe,
		},
	}

	for _, testCase := range testCases {
		rows := []driver.Value{
			testCase.user.ID,
			testCase.user.Name,
			testCase.user.Telephone,
			testCase.user.Password,
			testCase.user.DateBirth,
			testCase.user.Sex,
			testCase.user.Education,
			testCase.user.Job,
			testCase.user.AboutMe,
		}

		sqlmock.NewRows(columns).AddRow(rows...)
		for _, query := range queries {
			mock.ExpectExec(query.q).WithArgs(query.data, testCase.user.ID).WillReturnResult(sqlmock.NewResult(int64(testCase.user.ID), 1))
		}

		repo := NewPostgresUserRepository(sqlxDB.DB)

		err = repo.UpdateUser(testCase.user, testCase.user.ID)
		require.Equal(t, testCase.err, err)

		err = mock.ExpectationsWereMet()
		require.NoError(t, err, "unfulfilled expectations: %s", err)
	}
}

func TestPostgresUserRepository_SelectUserByID(t *testing.T) {
	type insertUserTestCase struct {
		id         int
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

	query := `SELECT id, name, telephone, password, date_birth, sex, job, education, about_me FROM users WHERE  id=$1;`

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

		repo := NewPostgresUserRepository(sqlxDB.DB)

		user, err := repo.SelectUserByID(testCase.outputUser.ID)
		require.Equal(t, testCase.err, err)
		if err == nil {
			require.Equal(t, testCase.outputUser, user)
		}

		err = dbMock.ExpectationsWereMet()
		require.NoError(t, err, "unfulfilled expectations: %s\n%s", err, msg)
	}
}

func TestPostgresUserRepository_SelectUserFeedByID(t *testing.T) {
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

	query := `SELECT name, date_birth, job, education, about_me FROM users WHERE id=$1;`

	var outputUser models.UserFeed
	err = faker.FakeData(&outputUser)
	require.NoError(t, err)
	outputUser.IsSuperlike = false

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

		repo := NewPostgresUserRepository(sqlxDB.DB)

		user, err := repo.SelectUserFeedByID(testCase.outputUser.ID)
		require.Equal(t, testCase.err, err)
		if err == nil {
			require.Equal(t, testCase.outputUser, user)
		}

		err = dbMock.ExpectationsWereMet()
		require.NoError(t, err, "unfulfilled expectations: %s\n%s", err, msg)
	}
}

func TestPostgresUserRepository_SelectUsers(t *testing.T) {
	type selectUsersTestCase struct {
		inputUser   models.User
		outputUsers []models.UserFeed
		err         error
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
		"job",
		"education",
		"about_me",
	}

	query := `SELECT id, name, date_birth, education, job, about_me FROM users WHERE sex != $1`

	var inputUser models.User
	err = faker.FakeData(&inputUser)
	require.NoError(t, err)

	var outputUsers []models.UserFeed
	var outputUser models.UserFeed
	err = faker.FakeData(&outputUser)
	require.NoError(t, err)
	outputUser.IsSuperlike = false

	outputUsers = append(outputUsers, outputUser)

	testCases := []selectUsersTestCase{
		{
			inputUser:   inputUser,
			outputUsers: outputUsers,
			err:         sql.ErrNoRows,
		},
	}

	for i, testCase := range testCases {
		msg := fmt.Sprintf("case %d aaaaaaaaaaaa", i)
		data := []driver.Value{
			1,
			testCase.outputUsers[0].Name,
			testCase.outputUsers[0].DateBirth,
			testCase.outputUsers[0].Education,
			testCase.outputUsers[0].Job,
			testCase.outputUsers[0].AboutMe,
		}

		if testCase.err != nil {

			dbMock.ExpectQuery(query).WithArgs(inputUser.Sex).WillReturnError(testCase.err)

		} else {
			rows := sqlmock.NewRows(columns).AddRow(data...)
			for _, user := range testCase.outputUsers {
				rows.AddRow(user.ID, user.Name, user.DateBirth, user.Job, user.Education, user.AboutMe)
			}
			dbMock.ExpectQuery(query).WithArgs(inputUser.Sex).WillReturnRows(rows)
		}

		repo := NewPostgresUserRepository(sqlxDB.DB)
		_, err := repo.SelectUsers(inputUser)
		require.Equal(t, testCase.err, err)

		err = dbMock.ExpectationsWereMet()
		require.NoError(t, err, "unfulfilled expectations: %s\n%s", err, msg)
	}
}

func TestPostgresUserRepository_SelectImages(t *testing.T) {
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

		repo := NewPostgresUserRepository(sqlxDB.DB)

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

func TestPostgresUserRepository_DeleteSession(t *testing.T) {
	type deleteSessionTestCase struct {
		key string
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
	}

	query := `DELETE FROM sessions WHERE key=$1;`

	var key string
	err = faker.FakeData(&key)
	require.NoError(t, err)

	testCases := []deleteSessionTestCase{
		{
			key: key,
			err: nil,
		},
	}

	for _, testCase := range testCases {
		args := []driver.Value{
			testCase.key,
		}

		rows := []driver.Value{
			1,
			testCase.key,
		}

		sqlmock.NewRows(columns).AddRow(rows...)
		mock.ExpectExec(query).WithArgs(args...).WillReturnResult(sqlmock.NewResult(1, 1))

		repo := NewPostgresUserRepository(sqlxDB.DB)

		err = repo.DeleteSession(testCase.key)
		require.Equal(t, testCase.err, err)

		err = mock.ExpectationsWereMet()
		require.NoError(t, err, "unfulfilled expectations: %s", err)
	}
}

func TestPostgresUserRepository_InsertPremium(t *testing.T) {
	type insertPremiumTestCase struct {
		uid      int
		dateFrom time.Time
		dateTo   time.Time
		err      error
	}

	db, mock, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
	if err != nil {
		t.Fatalf("error '%s' when opening a stub database connection", err)
	}
	defer db.Close()
	sqlxDB := sqlx.NewDb(db, "sqlmock")

	columns := []string{
		"id",
		"user_id",
		"date_from",
		"date_to",
	}

	query := `INSERT INTO premium_accounts(user_id, date_to, date_from) VALUES ($1, $2, $3);`

	var uid int
	err = faker.FakeData(&uid)
	require.NoError(t, err)
	dateFrom := time.Now()
	dateTo := dateFrom

	testCases := []insertPremiumTestCase{
		{
			uid:      uid,
			dateFrom: dateFrom,
			dateTo:   dateTo,
			err:      nil,
		},
	}

	for _, testCase := range testCases {
		args := []driver.Value{
			testCase.uid,
			testCase.dateFrom,
			testCase.dateTo,
		}

		rows := []driver.Value{
			1,
			testCase.uid,
			testCase.dateFrom,
			testCase.dateTo,
		}

		sqlmock.NewRows(columns).AddRow(rows...)
		mock.ExpectExec(query).WithArgs(args...).WillReturnResult(sqlmock.NewResult(1, 1))

		repo := NewPostgresUserRepository(sqlxDB.DB)

		err = repo.InsertPremium(testCase.uid, testCase.dateFrom, testCase.dateTo)
		require.Equal(t, testCase.err, err)

		err = mock.ExpectationsWereMet()
		require.NoError(t, err, "unfulfilled expectations: %s", err)
	}
}
