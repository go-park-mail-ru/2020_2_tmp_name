package usecase

import (
	"github.com/nfnt/resize"
	"image/jpeg"
	"os"
	domain "park_2020/2020_2_tmp_name/api/photos"
	"park_2020/2020_2_tmp_name/models"
	"path/filepath"
	"strings"
)

type photoUsecase struct {
	photoRepo domain.PhotoRepository
}

func NewPhotoUsecase(p domain.PhotoRepository) domain.PhotoUsecase {
	return &photoUsecase{
		photoRepo: p,
	}
}

func (p *photoUsecase) AddPhoto(photo models.Photo) error {
	user, err := p.photoRepo.SelectUserFeed(photo.Telephone)
	if err != nil {
		return models.ErrNotFound
	}
	user.LinkImages = append(user.LinkImages, photo.Path)

	err = p.photoRepo.InsertPhoto(photo.Path, user.ID)
	if err != nil {
		return models.ErrInternalServerError
	}
	return nil
}

func (p *photoUsecase) RemovePhoto(path string, uid int) error {
	err := p.ClearPhotos(path)
	if err != nil {
		return err
	}

	err = p.photoRepo.DeletePhoto(path, uid)
	if err != nil {
		return models.ErrInternalServerError
	}

	return nil
}

func (p *photoUsecase) ClearPhotos(path string) error {
	localPath := strings.Replace(path, "https://mi-ami.ru/static/avatars/", "", -1)
	photoPath := "/app/static/avatars/"

	err := os.RemoveAll(photoPath + localPath)
	if err != nil {
		return models.ErrNotFound
	}

	return nil
}

func (p *photoUsecase) FindPhotoWithMask(path string) ([]string, error) {
	under := strings.LastIndex(path, "_")
	if under != -1 {
		path = strings.Replace(path, path[under:], "", -1)
	} else {
		dot := strings.LastIndex(path, ".")
		if dot != -1 {
			path = strings.Replace(path, path[dot:], "", -1)
		}
	}

	photos, err := p.photoRepo.SelectPhotoWithMask(path)

	return photos, err
}

func (p *photoUsecase) FindPhotoWithoutMask(path string) (string, error) {
	photoName := strings.Replace(path, "https://mi-ami.ru/static/avatars/", "", -1)
	under := strings.LastIndex(photoName, "_")
	if under != -1 {
		photoName = photoName[:under]
	} else {
		dot := strings.LastIndex(photoName, ".")
		if dot != -1 {
			photoName = photoName[:dot]
		}
	}

	photoPath := "/app/static/avatars/"

	files, _ := filepath.Glob(photoPath + photoName + "*")

	for _, file := range files {
		if !strings.Contains(file, photoPath+photoName+"_") {
			return file, nil
		}
	}
	return "", models.ErrNotFound
}

func (p *photoUsecase) ResizePhoto(path string) error {
	imgIn, err := os.Open(path)
	if err != nil {
		return models.ErrInternalServerError
	}

	fi, err := imgIn.Stat()
	if err != nil {
		return models.ErrInternalServerError
	}
	size := fi.Size()
	if (size < 800000) {
		return nil
	}


	imgJpg, err := jpeg.Decode(imgIn)
	if err != nil {
		return models.ErrInternalServerError
	}
	defer imgIn.Close()

	width, height := imgJpg.Bounds().Dx(), imgJpg.Bounds().Dy()
	////
	//dst := image.NewRGBA(image.Rect(0, 0, width/10, height/10))
	//fmt.Println(dst.Bounds())


	//draw.ApproxBiLinear.Scale(dst, dst.Bounds(), imgJpg,
	//	imgJpg.Bounds(), draw.Over, nil)

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
