package usecase

import (
	"errors"
	"github.com/golang/mock/gomock"
	"park_2020/2020_2_tmp_name/domain"
	"park_2020/2020_2_tmp_name/domain/mock"
	"park_2020/2020_2_tmp_name/models"

	"github.com/stretchr/testify/require"
	"testing"
	"time"
)

func TestSignUpSuccess(t *testing.T) {
	user := models.User{
		ID:         0,
		Name:       "Misha",
		Telephone:  "909-277-47-21",
		Password:   "password",
		DateBirth:  time.Time{},
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

	us := userUsecase {
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
		DateBirth:  time.Time{},
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

	us := userUsecase {
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
		DateBirth:  time.Time{},
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
		DateBirth:  time.Time{},
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
		err := us.Like(cookie,like)
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
		err := us.Dislike(cookie,dislike)
		require.Equal(t, domain.ErrInternalServerError, err)
	}

}

func TestUserUsecase_CommentSuccess(t *testing.T) {
	cookie := "Something-like-uuid"
	telephone := "909-277-47-21"

	userFeed := models.UserFeed {
		ID:         0,
		Name:       "Misha",
		DateBirth:  0,
		LinkImages: nil,
		Job:        "Fullstack",
		Education:  "BMSTU",
		AboutMe:    "",
	}

	comment := models.Comment {
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

	userFeed := models.UserFeed {
		ID:         0,
		Name:       "Misha",
		DateBirth:  0,
		LinkImages: nil,
		Job:        "Fullstack",
		Education:  "BMSTU",
		AboutMe:    "",
	}

	comment := models.Comment {
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
		err := us.Comment(cookie,comment)
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

	userFeed := models.UserFeed {
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

	userFeed := models.UserFeed {
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
		err := us.Message(cookie,message)
		require.Equal(t, domain.ErrInternalServerError, err)
	}
}