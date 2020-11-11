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
	//user := models.User{
	//	ID:         0,
	//	Name:       "Misha",
	//	Telephone:  "909-277-47-21",
	//	Password:   "password",
	//	DateBirth:  time.Time{},
	//	Sex:        "male",
	//	LinkImages: nil,
	//	Job:        "",
	//	Education:  "BMSTU",
	//	AboutMe:    "",
	//}
	//
	//userFeed := models.UserFeed{
	//	ID:         0,
	//	Name:       "Misha",
	//	DateBirth:  0,
	//	LinkImages: nil,
	//	Job:        "Fullstack",
	//	Education:  "BMSTU",
	//	AboutMe:    "",
	//}
	//
	//cookie := "Something-like-uuid"
}