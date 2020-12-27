package usecase

import (
	"image"
	"image/jpeg"
	"os"
	domain "park_2020/2020_2_tmp_name/api/users"
	"park_2020/2020_2_tmp_name/models"
	"time"

	"github.com/disintegration/imaging"
	"github.com/nfnt/resize"
	"github.com/rwcarlsen/goexif/exif"
	"github.com/sirupsen/logrus"
)

type userUsecase struct {
	userRepo domain.UserRepository
}

func NewUserUsecase(u domain.UserRepository) domain.UserUsecase {
	return &userUsecase{
		userRepo: u,
	}
}

func (u *userUsecase) Signup(user models.User) error {
	var check bool
	if check = u.userRepo.CheckUser(user.Telephone); check {
		return models.ErrUnauthorized
	}

	err := u.userRepo.InsertUser(user)
	if err != nil {
		return models.ErrInternalServerError
	}

	return nil
}

func (u *userUsecase) Settings(uid int, userData models.User) error {
	err := u.userRepo.UpdateUser(userData, uid)
	if err != nil {
		return models.ErrInternalServerError
	}
	return nil
}

func (u *userUsecase) IsPremium(uid int) bool {
	return u.userRepo.CheckPremium(uid)
}

func (u *userUsecase) Feed(user models.User) ([]models.UserFeed, error) {
	data, err := u.userRepo.SelectUsers(user)
	if err != nil {
		return data, models.ErrNotFound
	}

	for idxUserFeed, userFeed := range data {
		data[idxUserFeed].IsSuperlike = u.userRepo.CheckSuperLikeMe(user.ID, userFeed.ID)
	}

	return data, nil
}

func (u *userUsecase) ChangeAvatar(user models.User, image models.Image) error {
	err := u.userRepo.ChangeAvatarPath(user.ID, image.LinkImage)
	if err != nil {
		return models.ErrNotFound
	}
	return nil
}

func (u *userUsecase) UserID(uid int) (models.UserFeed, error) {
	user, err := u.userRepo.SelectUserFeedByID(uid)
	if err != nil {
		return user, models.ErrNotFound
	}
	return user, nil
}

func (u *userUsecase) Telephone(telephone string) bool {
	return u.userRepo.CheckUser(telephone)
}

func (u *userUsecase) GetPremium(uid int) error {
	return u.userRepo.InsertPremium(uid, time.Now(), time.Now())
}

func (p *userUsecase) ResizePhoto(path string) error {
	imgIn, err := os.Open(path)
	if err != nil {
		return models.ErrInternalServerError
	}
	defer imgIn.Close()

	fi, err := imgIn.Stat()
	if err != nil {
		return models.ErrInternalServerError
	}
	size := fi.Size()
	if size < 800000 {
		return nil
	}

	err = p.rotateImg(path)
	if err != nil {
		return models.ErrInternalServerError
	}

	imgForDecode, err := os.Open(path)
	if err != nil {
		return models.ErrInternalServerError
	}
	defer imgForDecode.Close()

	imgJpg, err := jpeg.Decode(imgForDecode)
	if err != nil {
		return models.ErrInternalServerError
	}

	width, height := imgJpg.Bounds().Dx(), imgJpg.Bounds().Dy()
	imgJpg = resize.Resize(uint(width)/2, uint(height)/2, imgJpg, resize.Bicubic)

	imgOut, err := os.Create(path)
	if err != nil {
		return models.ErrInternalServerError
	}
	err = jpeg.Encode(imgOut, imgJpg, nil)
	if err != nil {
		return models.ErrInternalServerError
	}
	defer imgOut.Close()
	return nil
}

func (p *userUsecase) reverseOrientation(img image.Image, o string) *image.NRGBA {
	switch o {
	case "1":
		return imaging.Clone(img)
	case "2":
		return imaging.FlipV(img)
	case "3":
		return imaging.Rotate180(img)
	case "4":
		return imaging.Rotate180(imaging.FlipV(img))
	case "5":
		return imaging.Rotate270(imaging.FlipV(img))
	case "6":
		return imaging.Rotate270(img)
	case "7":
		return imaging.Rotate90(imaging.FlipV(img))
	case "8":
		return imaging.Rotate90(img)
	}
	logrus.Errorf("unknown orientation %s, expect 1-8", o)
	return imaging.Clone(img)
}

func (p *userUsecase) rotateImg(path string) error {
	imgInDecode, err := os.Open(path)
	if err != nil {
		return models.ErrInternalServerError
	}
	defer imgInDecode.Close()
	imgDecoded, err := os.Open(path)
	if err != nil {
		return models.ErrInternalServerError
	}
	defer imgDecoded.Close()

	var img image.Image
	img, err = jpeg.Decode(imgInDecode)
	if err != nil {
		return models.ErrInternalServerError
	}
	imgExif, err := exif.Decode(imgDecoded)
	if err != nil {
		return models.ErrInternalServerError
	}
	if imgExif == nil {
		return nil
	}
	orient, err := imgExif.Get(exif.Orientation)
	if err != nil {
		return models.ErrInternalServerError
	}
	if orient != nil {
		img = p.reverseOrientation(img, orient.String())
	} else {
		img = p.reverseOrientation(img, "1")
	}
	err = imaging.Save(img, path)
	if err != nil {
		return models.ErrInternalServerError
	}
	return nil
}
