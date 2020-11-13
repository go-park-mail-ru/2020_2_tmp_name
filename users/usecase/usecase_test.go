package usecase

import (
	"errors"
	"park_2020/2020_2_tmp_name/domain"
	"park_2020/2020_2_tmp_name/domain/mock"
	"park_2020/2020_2_tmp_name/models"

	"github.com/golang/mock/gomock"

	"testing"
	"time"

	"github.com/stretchr/testify/require"
)

func TestNewUserUsecase(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	var u domain.UserRepository
	uu := NewUserUsecase(u, time.Duration(10*time.Second))
	require.NotEmpty(t, uu)
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
	errors = append(errors, domain.ErrUnauthorized)
	errors = append(errors, domain.ErrInternalServerError)

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
		require.Equal(t, domain.ErrInternalServerError, err)
	}
}
func TestUserUsecase_LikeSuccess(t *testing.T) {
	like := models.Like{
		ID:   0,
		Uid1: 1,
		Uid2: 2,
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

	chat := models.Chat{
		ID:      0,
		Uid1:    userFeed.ID,
		Uid2:    like.Uid2,
		LastMsg: "",
	}
	cookie := "Something-like-uuid"
	telephone := "909-277-47-21"

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mock := mock.NewMockUserRepository(ctrl)
	firstCall := mock.EXPECT().CheckUserBySession(cookie).Return(telephone)
	mock.EXPECT().SelectUserFeed(telephone).After(firstCall).Return(userFeed, nil)
	mock.EXPECT().InsertLike(userFeed.ID, like.Uid2).Return(nil)
	mock.EXPECT().Match(userFeed.ID, like.Uid2).Return(true)
	mock.EXPECT().CheckChat(chat).Return(false)
	mock.EXPECT().InsertChat(chat).Return(nil)

	us := userUsecase{
		userRepo: mock,
	}

	err := us.Like(cookie, like)

	require.NoError(t, err)
	require.Equal(t, nil, err)

}

func TestUserUsecase_LikeFail(t *testing.T) {
	like := models.Like{
		ID:   0,
		Uid1: 1,
		Uid2: 2,
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
	chat := models.Chat{
		ID:      0,
		Uid1:    userFeed.ID,
		Uid2:    like.Uid2,
		LastMsg: "",
	}

	cookie := "Something-like-uuid"
	telephone := "909-277-47-21"

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mock := mock.NewMockUserRepository(ctrl)
	mock.EXPECT().CheckUserBySession(cookie).Times(3).Return(telephone)
	gomock.InOrder(
		mock.EXPECT().SelectUserFeed(telephone).Return(userFeed, errors.New("error select user")),
		mock.EXPECT().SelectUserFeed(telephone).Return(userFeed, nil),
		mock.EXPECT().InsertLike(userFeed.ID, like.Uid2).Return(errors.New("error of insert")),
		mock.EXPECT().SelectUserFeed(telephone).Return(userFeed, nil),
		mock.EXPECT().InsertLike(userFeed.ID, like.Uid2).Return(nil),
		mock.EXPECT().Match(userFeed.ID, like.Uid2).Return(true),
		mock.EXPECT().CheckChat(chat).Return(false),
		mock.EXPECT().InsertChat(chat).Return(errors.New("error of insert")),
	)

	us := userUsecase{
		userRepo: mock,
	}

	for i := 0; i < 3; i++ {
		err := us.Like(cookie, like)
		require.Equal(t, domain.ErrInternalServerError, err)
	}
}

func TestUserUsecase_DislikeSuccess(t *testing.T) {
	cookie := "Something-like-uuid"
	telephone := "909-277-47-21"

	userFeed := models.UserFeed{
		ID:         0,
		Name:       "Misha",
		DateBirth:  0,
		LinkImages: nil,
		Job:        "Fullstack",
		Education:  "BMSTU",
		AboutMe:    "",
	}

	dislike := models.Dislike{
		ID:   0,
		Uid1: userFeed.ID,
		Uid2: 2,
	}

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mock := mock.NewMockUserRepository(ctrl)
	mock.EXPECT().CheckUserBySession(cookie).Times(1).Return(telephone)
	mock.EXPECT().SelectUserFeed(telephone).Return(userFeed, nil)
	mock.EXPECT().InsertDislike(userFeed.ID, dislike.Uid2).Return(nil)

	us := userUsecase{
		userRepo: mock,
	}

	err := us.Dislike(cookie, dislike)

	require.NoError(t, err)
	require.Equal(t, nil, err)
}

func TestUserUsecase_DislikeFail(t *testing.T) {
	cookie := "Something-like-uuid"
	telephone := "909-277-47-21"

	userFeed := models.UserFeed{
		ID:         0,
		Name:       "Misha",
		DateBirth:  0,
		LinkImages: nil,
		Job:        "Fullstack",
		Education:  "BMSTU",
		AboutMe:    "",
	}

	dislike := models.Dislike{
		ID:   0,
		Uid1: userFeed.ID,
		Uid2: 2,
	}

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mock := mock.NewMockUserRepository(ctrl)
	mock.EXPECT().CheckUserBySession(cookie).Times(2).Return(telephone)
	gomock.InOrder(
		mock.EXPECT().SelectUserFeed(telephone).Return(userFeed, errors.New("error")),
		mock.EXPECT().SelectUserFeed(telephone).Return(userFeed, nil),
		mock.EXPECT().InsertDislike(userFeed.ID, dislike.Uid2).Return(errors.New("error")),
	)

	us := userUsecase{
		userRepo: mock,
	}

	for i := 0; i < 2; i++ {
		err := us.Dislike(cookie, dislike)
		require.Equal(t, domain.ErrInternalServerError, err)
	}

}

func TestUserUsecase_CommentSuccess(t *testing.T) {
	cookie := "Something-like-uuid"
	telephone := "909-277-47-21"

	userFeed := models.UserFeed{
		ID:         0,
		Name:       "Misha",
		DateBirth:  0,
		LinkImages: nil,
		Job:        "Fullstack",
		Education:  "BMSTU",
		AboutMe:    "",
	}

	comment := models.Comment{
		ID:           0,
		Uid1:         1,
		Uid2:         2,
		TimeDelivery: "7:23",
		CommentText:  "I love tests very much",
	}

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mock := mock.NewMockUserRepository(ctrl)
	mock.EXPECT().CheckUserBySession(cookie).Times(1).Return(telephone)
	mock.EXPECT().SelectUserFeed(telephone).Return(userFeed, nil)
	mock.EXPECT().InsertComment(comment, userFeed.ID).Return(nil)

	us := userUsecase{
		userRepo: mock,
	}

	err := us.Comment(cookie, comment)

	require.NoError(t, err)
	require.Equal(t, nil, err)
}

func TestUserUsecase_CommentFail(t *testing.T) {
	cookie := "Something-like-uuid"
	telephone := "909-277-47-21"

	userFeed := models.UserFeed{
		ID:         0,
		Name:       "Misha",
		DateBirth:  0,
		LinkImages: nil,
		Job:        "Fullstack",
		Education:  "BMSTU",
		AboutMe:    "",
	}

	comment := models.Comment{
		ID:           0,
		Uid1:         1,
		Uid2:         2,
		TimeDelivery: "7:23",
		CommentText:  "I love tests very much",
	}

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mock := mock.NewMockUserRepository(ctrl)
	mock.EXPECT().CheckUserBySession(cookie).Times(2).Return(telephone)
	gomock.InOrder(
		mock.EXPECT().SelectUserFeed(telephone).Return(userFeed, errors.New("error")),
		mock.EXPECT().SelectUserFeed(telephone).Return(userFeed, nil),
		mock.EXPECT().InsertComment(comment, userFeed.ID).Return(errors.New("error")),
	)

	us := userUsecase{
		userRepo: mock,
	}

	for i := 0; i < 2; i++ {
		err := us.Comment(cookie, comment)
		require.Equal(t, domain.ErrInternalServerError, err)
	}

}

func TestUserUsecase_CommentsByIDSuccess(t *testing.T) {
	comments := models.CommentsById{}
	Data := models.CommentsData{}
	Data.Data = comments
	id := 2

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mock := mock.NewMockUserRepository(ctrl)
	mock.EXPECT().SelectComments(id).Return(comments, nil)

	us := userUsecase{
		userRepo: mock,
	}

	data, err := us.CommentsByID(id)

	require.NoError(t, err)
	require.Equal(t, Data, data)
}

func TestUserUsecase_CommentsByIDFail(t *testing.T) {
	comments := models.CommentsById{}
	Data := models.CommentsData{}
	Data.Data = comments
	id := 2

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mock := mock.NewMockUserRepository(ctrl)
	mock.EXPECT().SelectComments(id).Return(comments, errors.New("error"))

	us := userUsecase{
		userRepo: mock,
	}

	_, err := us.CommentsByID(id)

	require.Equal(t, domain.ErrInternalServerError, err)
}

func TestUserUsecase_MessageSuccess(t *testing.T) {
	cookie := "Something-like-uuid"
	telephone := "909-277-47-21"

	userFeed := models.UserFeed{
		ID:         0,
		Name:       "Misha",
		DateBirth:  0,
		LinkImages: nil,
		Job:        "Fullstack",
		Education:  "BMSTU",
		AboutMe:    "",
	}

	message := models.Message{
		ID:           0,
		Text:         "Save me from tests",
		TimeDelivery: time.Time{},
		ChatID:       1,
		UserID:       2,
	}

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mock := mock.NewMockUserRepository(ctrl)
	mock.EXPECT().CheckUserBySession(cookie).Times(1).Return(telephone)
	mock.EXPECT().SelectUserFeed(telephone).Return(userFeed, nil)
	mock.EXPECT().InsertMessage(message.Text, message.ChatID, userFeed.ID).Return(nil)

	us := userUsecase{
		userRepo: mock,
	}

	err := us.Message(cookie, message)

	require.NoError(t, err)
	require.Equal(t, nil, err)

}

func TestUserUsecase_MessageFail(t *testing.T) {
	cookie := "Something-like-uuid"
	telephone := "909-277-47-21"

	userFeed := models.UserFeed{
		ID:         0,
		Name:       "Misha",
		DateBirth:  0,
		LinkImages: nil,
		Job:        "Fullstack",
		Education:  "BMSTU",
		AboutMe:    "",
	}

	message := models.Message{
		ID:           0,
		Text:         "Save me from tests",
		TimeDelivery: time.Time{},
		ChatID:       1,
		UserID:       2,
	}

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mock := mock.NewMockUserRepository(ctrl)
	mock.EXPECT().CheckUserBySession(cookie).Times(2).Return(telephone)
	gomock.InOrder(
		mock.EXPECT().SelectUserFeed(telephone).Return(userFeed, errors.New("error")),
		mock.EXPECT().SelectUserFeed(telephone).Return(userFeed, nil),
		mock.EXPECT().InsertMessage(message.Text, message.ChatID, userFeed.ID).Return(errors.New("error")),
	)

	us := userUsecase{
		userRepo: mock,
	}

	for i := 0; i < 2; i++ {
		err := us.Message(cookie, message)
		require.Equal(t, domain.ErrInternalServerError, err)
	}
}

func TestUserUsecase_ChatSuccess(t *testing.T) {
	chat := models.Chat{
		ID:      0,
		Uid1:    4,
		Uid2:    3,
		LastMsg: "Save me from tests",
	}

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mock := mock.NewMockUserRepository(ctrl)
	mock.EXPECT().InsertChat(chat).Times(1).Return(nil)

	us := userUsecase{
		userRepo: mock,
	}

	err := us.Chat(chat)

	require.NoError(t, err)
	require.Equal(t, nil, err)
}

func TestUserUsecase_ChatFail(t *testing.T) {
	chat := models.Chat{
		ID:      0,
		Uid1:    4,
		Uid2:    3,
		LastMsg: "Save me from tests",
	}

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mock := mock.NewMockUserRepository(ctrl)
	mock.EXPECT().InsertChat(chat).Times(1).Return(domain.ErrInternalServerError)

	us := userUsecase{
		userRepo: mock,
	}

	err := us.Chat(chat)

	require.NotEqual(t, nil, err)
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

func TestAddPhotoSuccess(t *testing.T) {
	photo := models.Photo{
		Path:      "path",
		Telephone: "944-739-32-28",
	}

	user := models.UserFeed{
		ID:         1,
		Name:       "Andrey",
		DateBirth:  20,
		LinkImages: nil,
		Job:        "",
		Education:  "BMSTU",
		AboutMe:    "",
	}

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mock := mock.NewMockUserRepository(ctrl)
	mock.EXPECT().SelectUserFeed(photo.Telephone).Times(1).Return(user, nil)
	mock.EXPECT().InsertPhoto(photo.Path, user.ID).Times(1).Return(nil)

	us := userUsecase{
		userRepo: mock,
	}

	err := us.AddPhoto(photo)

	require.NoError(t, err)
}

func TestAddPhotoFail(t *testing.T) {
	photo := models.Photo{
		Path:      "path",
		Telephone: "944-739-32-28",
	}

	user := models.UserFeed{
		ID:         1,
		Name:       "Andrey",
		DateBirth:  20,
		LinkImages: nil,
		Job:        "",
		Education:  "BMSTU",
		AboutMe:    "",
	}

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mock := mock.NewMockUserRepository(ctrl)
	mock.EXPECT().SelectUserFeed(photo.Telephone).Times(1).Return(user, domain.ErrInternalServerError)

	us := userUsecase{
		userRepo: mock,
	}

	err := us.AddPhoto(photo)

	require.NotEqual(t, err, nil)
}

func TestAddPhotoFailSelect(t *testing.T) {
	photo := models.Photo{
		Path:      "path",
		Telephone: "944-739-32-28",
	}

	user := models.UserFeed{
		ID:         1,
		Name:       "Andrey",
		DateBirth:  20,
		LinkImages: nil,
		Job:        "",
		Education:  "BMSTU",
		AboutMe:    "",
	}

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mock := mock.NewMockUserRepository(ctrl)
	mock.EXPECT().SelectUserFeed(photo.Telephone).Times(1).Return(user, nil)
	mock.EXPECT().InsertPhoto(photo.Path, user.ID).Times(1).Return(domain.ErrInternalServerError)

	us := userUsecase{
		userRepo: mock,
	}

	err := us.AddPhoto(photo)

	require.NotEqual(t, err, nil)
}

func TestUploadAvatarSuccess(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mock := mock.NewMockUserRepository(ctrl)

	us := userUsecase{
		userRepo: mock,
	}

	uid, err := us.UploadAvatar()

	require.NoError(t, err)
	require.NotEqual(t, uid.String(), "")
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
	mock.EXPECT().SelectUserFeed(telephone).Times(1).Return(user, domain.ErrInternalServerError)

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
	mock.EXPECT().SelectUser(telephone).Times(1).Return(user, domain.ErrInternalServerError)

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
	mock.EXPECT().SelectUsers(user).Times(1).Return(users, domain.ErrInternalServerError)

	us := userUsecase{
		userRepo: mock,
	}

	_, err := us.Feed(sid)

	require.NotEqual(t, err, nil)
}

func TestСhats(t *testing.T) {
	sid := "something-like-this"
	var chatModel models.ChatModel
	user := models.UserFeed{
		ID:         1,
		Name:       "Andrey",
		DateBirth:  20,
		LinkImages: nil,
		Job:        "",
		Education:  "BMSTU",
		AboutMe:    "",
	}

	msg1 := models.Msg{
		UserID:       1,
		ChatID:       3,
		Message:      "Hi",
		TimeDelivery: "18:40",
	}

	msg2 := models.Msg{
		UserID:       6,
		ChatID:       3,
		Message:      "Hi",
		TimeDelivery: "18:41",
	}

	var chats []models.ChatData

	chat1 := models.ChatData{
		ID: 1,
		Partner: models.UserFeed{
			ID:         6,
			Name:       "Natasha",
			DateBirth:  20,
			LinkImages: nil,
			Job:        "",
			Education:  "BMSTU",
			AboutMe:    "",
		},
		Messages: []models.Msg{msg1, msg2},
	}

	chat2 := models.ChatData{
		ID: 2,
		Partner: models.UserFeed{
			ID:         4,
			Name:       "Dasha",
			DateBirth:  20,
			LinkImages: nil,
			Job:        "",
			Education:  "BMSTU",
			AboutMe:    "",
		},
		Messages: []models.Msg{msg1, msg2},
	}

	chats = append(chats, chat1, chat2)

	telephone := "944-739-32-28"

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mock := mock.NewMockUserRepository(ctrl)
	mock.EXPECT().CheckUserBySession(sid).Times(1).Return(telephone)
	mock.EXPECT().SelectUserFeed(telephone).Times(1).Return(user, nil)
	mock.EXPECT().SelectChatsByID(user.ID).Times(1).Return(chats, nil)
	chatModel.Data = chats

	us := userUsecase{
		userRepo: mock,
	}

	result, err := us.Chats(sid)

	require.NoError(t, err)
	require.Equal(t, result, chatModel)
}

func TestСhatID(t *testing.T) {
	sid := "something-like-this"
	var chid = 1
	user := models.UserFeed{
		ID:         1,
		Name:       "Andrey",
		DateBirth:  20,
		LinkImages: nil,
		Job:        "",
		Education:  "BMSTU",
		AboutMe:    "",
	}

	msg1 := models.Msg{
		UserID:       1,
		ChatID:       3,
		Message:      "Hi",
		TimeDelivery: "18:40",
	}

	msg2 := models.Msg{
		UserID:       6,
		ChatID:       3,
		Message:      "Hi",
		TimeDelivery: "18:41",
	}

	chat := models.ChatData{
		ID: 1,
		Partner: models.UserFeed{
			ID:         6,
			Name:       "Natasha",
			DateBirth:  20,
			LinkImages: nil,
			Job:        "",
			Education:  "BMSTU",
			AboutMe:    "",
		},
		Messages: []models.Msg{msg1, msg2},
	}

	telephone := "944-739-32-28"

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mock := mock.NewMockUserRepository(ctrl)
	mock.EXPECT().CheckUserBySession(sid).Times(1).Return(telephone)
	mock.EXPECT().SelectUserFeed(telephone).Times(1).Return(user, nil)
	mock.EXPECT().SelectChatByID(user.ID, chid).Times(1).Return(chat, nil)

	us := userUsecase{
		userRepo: mock,
	}

	result, err := us.ChatID(sid, chid)

	require.NoError(t, err)
	require.Equal(t, result, chat)
}

func TestСhatIDFail(t *testing.T) {
	sid := "something-like-this"
	var chid = 1
	user := models.UserFeed{}

	telephone := "944-739-32-28"

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mock := mock.NewMockUserRepository(ctrl)
	mock.EXPECT().CheckUserBySession(sid).Times(1).Return(telephone)
	mock.EXPECT().SelectUserFeed(telephone).Times(1).Return(user, domain.ErrInternalServerError)

	us := userUsecase{
		userRepo: mock,
	}

	_, err := us.ChatID(sid, chid)

	require.NotEqual(t, err, nil)
}

func TestСhatSelectIDFail(t *testing.T) {
	sid := "something-like-this"
	var chid = 1
	user := models.UserFeed{}

	msg1 := models.Msg{}

	msg2 := models.Msg{}

	chat := models.ChatData{
		ID:       1,
		Partner:  models.UserFeed{},
		Messages: []models.Msg{msg1, msg2},
	}

	telephone := "944-739-32-28"

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mock := mock.NewMockUserRepository(ctrl)
	mock.EXPECT().CheckUserBySession(sid).Times(1).Return(telephone)
	mock.EXPECT().SelectUserFeed(telephone).Times(1).Return(user, nil)
	mock.EXPECT().SelectChatByID(user.ID, chid).Times(1).Return(chat, domain.ErrInternalServerError)

	us := userUsecase{
		userRepo: mock,
	}

	_, err := us.ChatID(sid, chid)

	require.NotEqual(t, err, nil)
}

func TestGochat(t *testing.T) {
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

	result, err := us.Gochat(sid)

	require.NoError(t, err)
	require.Equal(t, result, user)
}

func TestGochatFail(t *testing.T) {
	sid := "something-like-this"
	user := models.UserFeed{}

	telephone := "944-739-32-28"

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mock := mock.NewMockUserRepository(ctrl)
	mock.EXPECT().CheckUserBySession(sid).Times(1).Return(telephone)
	mock.EXPECT().SelectUserFeed(telephone).Times(1).Return(user, domain.ErrInternalServerError)

	us := userUsecase{
		userRepo: mock,
	}

	_, err := us.Gochat(sid)

	require.NotEqual(t, err, nil)
}
