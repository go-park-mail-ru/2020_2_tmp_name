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
	require.Empty(t, uu)
}

func TestUserUsecase_LoginFail(t *testing.T) {
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

func TestUserUsecase_Logout(t *testing.T) {
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

func TestUserUsecase_SignUpSuccess(t *testing.T) {
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

	uid := 1

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mock := mock.NewMockUserRepository(ctrl)
	mock.EXPECT().UpdateUser(user, uid).Return(nil)

	us := userUsecase{
		userRepo: mock,
	}

	err := us.Settings(uid, user)

	require.NoError(t, err)
	require.Equal(t, nil, err)
}

func TestUserUsecase_SettingsFail(t *testing.T) {
	user := models.User{}
	uid := 1

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mock := mock.NewMockUserRepository(ctrl)
	mock.EXPECT().UpdateUser(user, uid).Return(models.ErrInternalServerError)

	us := userUsecase{
		userRepo: mock,
	}

	err := us.Settings(uid, user)
	require.Equal(t, models.ErrInternalServerError, err)

}

func TestUserUsecase_MeSuccess(t *testing.T) {
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

func TestUserUsecase_MeFail(t *testing.T) {
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

func TestUserUsecase_Feed(t *testing.T) {
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

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mock := mock.NewMockUserRepository(ctrl)
	mock.EXPECT().SelectUsers(user).Times(1).Return(users, nil)

	us := userUsecase{
		userRepo: mock,
	}

	feed, err := us.Feed(user)

	require.NoError(t, err)
	require.Equal(t, feed, users)
}

func TestUserUsecase_FeedFail(t *testing.T) {
	user := models.User{}

	var users []models.UserFeed
	user1 := models.UserFeed{}

	user2 := models.UserFeed{}

	users = append(users, user1, user2)

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mock := mock.NewMockUserRepository(ctrl)
	mock.EXPECT().SelectUsers(user).Return(users, models.ErrNotFound)

	us := userUsecase{
		userRepo: mock,
	}

	_, err := us.Feed(user)

	require.NotEqual(t, err, nil)
}

func TestUserUsecase_FeedSelectFail(t *testing.T) {
	user := models.User{}
	var users []models.UserFeed
	user1 := models.UserFeed{}
	user2 := models.UserFeed{}
	users = append(users, user1, user2)

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mock := mock.NewMockUserRepository(ctrl)
	mock.EXPECT().SelectUsers(user).Times(1).Return(users, models.ErrInternalServerError)

	us := userUsecase{
		userRepo: mock,
	}

	_, err := us.Feed(user)

	require.NotEqual(t, err, nil)
}

func TestUserUsecase_UserSuccess(t *testing.T) {
	user := models.User{
		ID:         0,
		Name:       "Misha",
		DateBirth:  0,
		LinkImages: nil,
		Job:        "Fullstack",
		Education:  "BMSTU",
		AboutMe:    "",
	}

	telephone := "(944) 546 98 24"
	sid := "something-like-this"

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mock := mock.NewMockUserRepository(ctrl)
	mock.EXPECT().CheckUserBySession(sid).Return(telephone)
	mock.EXPECT().SelectUser(telephone).Return(user, nil)

	us := userUsecase{
		userRepo: mock,
	}

	result, err := us.User(sid)

	require.NoError(t, err)
	require.Equal(t, result, user)
}

func TestUserUsecase_UserFail(t *testing.T) {
	user := models.User{
		ID:         0,
		Name:       "Misha",
		DateBirth:  0,
		LinkImages: nil,
		Job:        "Fullstack",
		Education:  "BMSTU",
		AboutMe:    "",
	}

	telephone := "(944) 546 98 24"
	sid := "something-like-this"

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mock := mock.NewMockUserRepository(ctrl)
	mock.EXPECT().CheckUserBySession(sid).Return(telephone)
	mock.EXPECT().SelectUser(telephone).Return(user, models.ErrNotFound)

	us := userUsecase{
		userRepo: mock,
	}

	_, err := us.User(sid)

	require.Equal(t, err, models.ErrNotFound)
}

func TestUserUsecase_UserIDSuccess(t *testing.T) {
	user := models.UserFeed{
		ID:         1,
		Name:       "Misha",
		DateBirth:  0,
		LinkImages: nil,
		Job:        "Fullstack",
		Education:  "BMSTU",
		AboutMe:    "",
	}

	uid := 1

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mock := mock.NewMockUserRepository(ctrl)
	mock.EXPECT().SelectUserFeedByID(uid).Return(user, nil)

	us := userUsecase{
		userRepo: mock,
	}

	result, err := us.UserID(uid)

	require.NoError(t, err)
	require.Equal(t, result, user)
}

func TestUserUsecase_UserIDFail(t *testing.T) {
	user := models.UserFeed{
		ID:         1,
		Name:       "Misha",
		DateBirth:  0,
		LinkImages: nil,
		Job:        "Fullstack",
		Education:  "BMSTU",
		AboutMe:    "",
	}

	uid := 1

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mock := mock.NewMockUserRepository(ctrl)
	mock.EXPECT().SelectUserFeedByID(uid).Return(user, models.ErrNotFound)

	us := userUsecase{
		userRepo: mock,
	}

	_, err := us.UserID(uid)

	require.Equal(t, err, models.ErrNotFound)
}
