package server

import (
	"log"
	"park_2020/2020_2_tmp_name/models"
)

func (s *Service) CheckUser(telephone string) bool {
	var count int
	s.DB.QueryRow(`SELECT COUNT(telephone) FROM users WHERE telephone=$1;`, telephone).Scan(&count)
	return count > 0
}

func (s *Service) InsertUser(user models.User, id int) error {
	password, err := HashPassword(user.Password)
	if err != nil {
		log.Println(err)
		return err
	}
	_, err = s.DB.Exec(`INSERT INTO users(id, name, telephone, password, date_birth, sex, job, education, about_me)
						VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9);`, id,
		user.Name,
		user.Telephone,
		password,
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

func (s *Service) SelectUser(telephone string) (models.User, error) {
	var u models.User
	row := s.DB.QueryRow(`SELECT name, telephone, password, date_birth, sex, job, education, about_me FROM users
						WHERE  telephone=$1;`, telephone)
	err := row.Scan(&u.Name, &u.Telephone, &u.Password, &u.DateBirth, &u.Sex, &u.Education, &u.Job, &u.AboutMe)
	if err != nil {
		log.Println(err)
		var eu models.User
		return eu, err
	}
	return u, nil
}

func (s *Service) UpdateUser(user models.User) error {
	if user.Name != "" {
		_, err := s.DB.Exec(`UPDATE users SET name=$1`, user.Name)
		if err != nil {
			log.Println(err)
			return err
		}
	}
	if user.Telephone != "" {
		_, err := s.DB.Exec(`UPDATE users SET telephone=$1`, user.Telephone)
		if err != nil {
			log.Println(err)
			return err
		}
	}
	if user.Job != "" {
		_, err := s.DB.Exec(`UPDATE users SET job=$1;`, user.Job)
		if err != nil {
			log.Println(err)
			return err
		}
	}
	if user.Education != "" {
		_, err := s.DB.Exec(`UPDATE users SET education=$1;`, user.Education)
		if err != nil {
			log.Println(err)
			return err
		}
	}
	if user.AboutMe != "" {
		_, err := s.DB.Exec(`UPDATE users SET about_me=$1;`, user.AboutMe)
		if err != nil {
			log.Println(err)
			return err
		}
	}
	return nil
}

func (s *Service) CheckSession(telephone string) bool {
	var count int
	s.DB.QueryRow(`SELECT COUNT(value) FROM sessions WHERE value=$1;`, telephone).Scan(&count)
	return count > 0
}

func (s *Service) InsertSession(sid, telephone string) error {
	_, err := s.DB.Exec(`INSERT INTO sessions(key, value) VALUES ($1, $2);`, sid, telephone)
	if err != nil {
		log.Println(err)
		return err
	}
	return nil
}

func (s *Service) DeleteSession(sid string) error {
	_, err := s.DB.Exec(`DELETE FROM sessions WHERE key=$1;`, sid)
	if err != nil {
		log.Println(err)
		return err
	}
	return nil
}

func (s *Service) CheckUserBySession(sid string) string {
	var count string
	s.DB.QueryRow(`SELECT value FROM sessions WHERE key=$1;`, sid).Scan(&count)
	return count
}
