package server

import (
	"log"
	"park_2020/2020_2_tmp_name/models"
)

func (s *Service) CheckUserSignup(telephone string) (bool, error) {
	var count int
	err := s.DB.QueryRow(`SELECT COUNT(telephone) FROM users WHERE telephone=$1;`, telephone).Scan(&count)
	if err != nil {
		log.Println(err)
		return false, err
	}
	return (count > 0), nil
}

func (s *Service) InsertUserSignup(user models.User, id int) error {
	_, err := s.DB.Exec(`INSERT INTO users(id, name, telephone, password, date_birth, sex, job, education, about_me)
						VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9);`, id,
		user.Name,
		user.Telephone,
		user.Password,
		user.DateBirth,
		user.Sex,
		user.Job,
		user.Education,
		user.AboutMe,
	)
	if err != nil {
		log.Println(err)
		return err
	}
	return nil
}
