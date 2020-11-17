package usecase

import (
	"errors"
	domain "park_2020/2020_2_tmp_name/api/users"
	"park_2020/2020_2_tmp_name/api/users/mock"
	"park_2020/2020_2_tmp_name/models"

	"github.com/golang/mock/gomock"

	"testing"

	"github.com/stretchr/testify/require"
)

func TestNewUserUsecase(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	var u domain.UserRepository
	uu := NewUserUsecase(u)
	require.NotEmpty(t, uu)
}

func TestLoginFail(t *testing.T) {
	login := models.LoginData{
		Telephone: "944-739-32-28",
		Password:  "password",
	}

	user := models.User{
		ID:         1,
		Name:       "Andrey",
		Telephone:  "944-739-32-28",
		Password:   "password",
		DateBirth:  20,
		Sex:        "male",
		LinkImages: nil,
		Job:        "",
		Education:  "BMSTU",
		AboutMe:    "",
	}

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	mock := mock.NewMockUserRepository(ctrl)
	mock.EXPECT().CheckUser(login.Telephone).Times(1).Return(true)
	mock.EXPECT().SelectUser(login.Telephone).Times(1).Return(user, nil)

	us := userUsecase{
		userRepo: mock,
	}

	_, err := us.Login(login)

	require.NotEqual(t, err, nil)
}

func TestLogout(t *testing.T) {
	sid := "something-like-this"

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mock := mock.NewMockUserRepository(ctrl)
	mock.EXPECT().DeleteSession(sid).Times(1).Return(nil)

	us := userUsecase{
		userRepo: mock,
	}

	err := us.Logout(sid)
	require.NoError(t, err)
}

func TestSignUpSuccess(t *testing.T) {
	user := models.User{
		ID:         0,
		Name:       "Misha",
		Telephone:  "909-277-47-21",
		Password:   "password",
		DateBirth:  20,
		Sex:        "male",
		LinkImages: nil,
		Job:        "",
		Education:  "BMSTU",
		AboutMe:    "",
	}
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mock := mock.NewMockUserRepository(ctrl)
	mock.EXPECT().CheckUser(user.Telephone).Return(false)
	mock.EXPECT().InsertUser(user).Return(nil)

	us := userUsecase{
		userRepo: mock,
	}

	err := us.Signup(user)

	require.NoError(t, err)
	require.Equal(t, nil, err)
}

func TestUserUsecase_SignupFail(t *testing.T) {
	user := models.User{
		ID:         0,
		Name:       "Misha",
		Telephone:  "909-277-47-21",
		Password:   "password",
		DateBirth:  20,
		Sex:        "male",
		LinkImages: nil,
		Job:        "",
		Education:  "BMSTU",
		AboutMe:    "",
	}
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mock := mock.NewMockUserRepository(ctrl)
	firstCall := mock.EXPECT().CheckUser(user.Telephone).Return(true)
	secondCall := mock.EXPECT().CheckUser(user.Telephone).After(firstCall).Return(false)
	mock.EXPECT().InsertUser(user).After(secondCall).Return(errors.New("Fail to insert"))

	us := userUsecase{
		userRepo: mock,
	}

	var errors []error
	errors = make([]error, 0, 1)
	errors = append(errors, models.ErrUnauthorized)
	errors = append(errors, models.ErrInternalServerError)

	for i := 0; i < 2; i++ {
		err := us.Signup(user)
		require.Equal(t, errors[i], err)
	}

}

func TestUserUsecase_SettingsSuccess(t *testing.T) {
	user := models.User{
		ID:         0,
		Name:       "Misha",
		Telephone:  "909-277-47-21",
		Password:   "password",
		DateBirth:  20,
		Sex:        "male",
		LinkImages: nil,
		Job:        "",
		Education:  "BMSTU",
		AboutMe:    "",
	}

	userFeed := models.UserFeed{
		ID:         0,
		Name:       "Misha",
		DateBirth:  0,
		LinkImages: nil,
		Job:        "Fullstack",
		Education:  "BMSTU",
		AboutMe:    "",
	}

	cookie := "Something-like-uuid"

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mock := mock.NewMockUserRepository(ctrl)
	firstCall := mock.EXPECT().CheckUserBySession(cookie).Return(user.Telephone)
	mock.EXPECT().SelectUserFeed(user.Telephone).After(firstCall).Return(userFeed, nil)
	mock.EXPECT().UpdateUser(user, userFeed.ID).Return(nil)

	us := userUsecase{
		userRepo: mock,
	}

	err := us.Settings(cookie, user)

	require.NoError(t, err)
	require.Equal(t, nil, err)
}

func TestUserUsecase_SettingsFail(t *testing.T) {
	user := models.User{
		ID:         0,
		Name:       "Misha",
		Telephone:  "909-277-47-21",
		Password:   "password",
		DateBirth:  20,
		Sex:        "male",
		LinkImages: nil,
		Job:        "",
		Education:  "BMSTU",
		AboutMe:    "",
	}

	userFeed := models.UserFeed{
		ID:         0,
		Name:       "Misha",
		DateBirth:  0,
		LinkImages: nil,
		Job:        "Fullstack",
		Education:  "BMSTU",
		AboutMe:    "",
	}

	cookie := "Something-like-uuid"

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mock := mock.NewMockUserRepository(ctrl)
	mock.EXPECT().CheckUserBySession(cookie).Times(2).Return(user.Telephone)
	firstCall := mock.EXPECT().SelectUserFeed(user.Telephone).Return(userFeed, errors.New("Have not this user"))
	secondCall := mock.EXPECT().SelectUserFeed(user.Telephone).After(firstCall).Return(userFeed, nil)
	mock.EXPECT().UpdateUser(user, userFeed.ID).After(secondCall).Return(errors.New("Could not update"))

	us := userUsecase{
		userRepo: mock,
	}

	for i := 0; i < 2; i++ {
		err := us.Settings(cookie, user)
		require.Equal(t, models.ErrInternalServerError, err)
	}
}

func TestMeSuccess(t *testing.T) {
	sid := "something-like-this"

	user := models.UserFeed{
		ID:         1,
		Name:       "Andrey",
		DateBirth:  20,
		LinkImages: nil,
		Job:        "",
		Education:  "BMSTU",
		AboutMe:    "",
	}

	telephone := "944-739-32-28"

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mock := mock.NewMockUserRepository(ctrl)
	mock.EXPECT().CheckUserBySession(sid).Times(1).Return(telephone)
	mock.EXPECT().SelectUserFeed(telephone).Times(1).Return(user, nil)

	us := userUsecase{
		userRepo: mock,
	}

	me, err := us.Me(sid)

	require.NoError(t, err)
	require.Equal(t, me, user)
}

func TestMeFail(t *testing.T) {
	sid := "something-like-this"

	user := models.UserFeed{}

	telephone := "944-739-32-28"

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mock := mock.NewMockUserRepository(ctrl)
	mock.EXPECT().CheckUserBySession(sid).Times(1).Return(telephone)
	mock.EXPECT().SelectUserFeed(telephone).Times(1).Return(user, models.ErrInternalServerError)

	us := userUsecase{
		userRepo: mock,
	}

	_, err := us.Me(sid)

	require.NotEqual(t, err, nil)
}

func TestFeed(t *testing.T) {
	sid := "something-like-this"

	user := models.User{
		ID:         1,
		Name:       "Andrey",
		Telephone:  "944-739-32-28",
		Password:   "password",
		DateBirth:  20,
		Sex:        "male",
		LinkImages: nil,
		Job:        "",
		Education:  "BMSTU",
		AboutMe:    "",
	}

	var users []models.UserFeed
	user1 := models.UserFeed{
		ID:         3,
		Name:       "Masha",
		DateBirth:  20,
		LinkImages: nil,
		Job:        "",
		Education:  "BMSTU",
		AboutMe:    "",
	}

	user2 := models.UserFeed{
		ID:         4,
		Name:       "Dasha",
		DateBirth:  20,
		LinkImages: nil,
		Job:        "",
		Education:  "BMSTU",
		AboutMe:    "",
	}

	users = append(users, user1, user2)

	telephone := "944-739-32-28"

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mock := mock.NewMockUserRepository(ctrl)
	mock.EXPECT().CheckUserBySession(sid).Times(1).Return(telephone)
	mock.EXPECT().SelectUser(telephone).Times(1).Return(user, nil)
	mock.EXPECT().SelectUsers(user).Times(1).Return(users, nil)

	us := userUsecase{
		userRepo: mock,
	}

	feed, err := us.Feed(sid)

	require.NoError(t, err)
	require.Equal(t, feed, users)
}

func TestFeedFail(t *testing.T) {
	sid := "something-like-this"

	user := models.User{}

	var users []models.UserFeed
	user1 := models.UserFeed{}

	user2 := models.UserFeed{}

	users = append(users, user1, user2)

	telephone := "944-739-32-28"

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mock := mock.NewMockUserRepository(ctrl)
	mock.EXPECT().CheckUserBySession(sid).Times(1).Return(telephone)
	mock.EXPECT().SelectUser(telephone).Times(1).Return(user, models.ErrInternalServerError)

	us := userUsecase{
		userRepo: mock,
	}

	_, err := us.Feed(sid)

	require.NotEqual(t, err, nil)
}

func TestFeedSelectFail(t *testing.T) {
	sid := "something-like-this"

	user := models.User{}

	var users []models.UserFeed
	user1 := models.UserFeed{}

	user2 := models.UserFeed{}

	users = append(users, user1, user2)

	telephone := "944-739-32-28"

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mock := mock.NewMockUserRepository(ctrl)
	mock.EXPECT().CheckUserBySession(sid).Times(1).Return(telephone)
	mock.EXPECT().SelectUser(telephone).Times(1).Return(user, nil)
	mock.EXPECT().SelectUsers(user).Times(1).Return(users, models.ErrInternalServerError)

	us := userUsecase{
		userRepo: mock,
	}

	_, err := us.Feed(sid)

	require.NotEqual(t, err, nil)
}
