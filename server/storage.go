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

func (s *Service) InsertUser(user models.User) error {
	password, err := HashPassword(user.Password)
	if err != nil {
		log.Println(err)
		return err
	}
	_, err = s.DB.Exec(`INSERT INTO users(name, telephone, password, date_birth, sex, job, education, about_me)
						VALUES ($1, $2, $3, $4, $5, $6, $7, $8);`,
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
	row := s.DB.QueryRow(`SELECT id, name, telephone, password, date_birth, sex, job, education, about_me FROM users
						WHERE  telephone=$1;`, telephone)
	err := row.Scan(&u.ID, &u.Name, &u.Telephone, &u.Password, &u.DateBirth, &u.Sex, &u.Education, &u.Job, &u.AboutMe)
	if err != nil {
		log.Println(err)
		return u, err
	}
	return u, nil
}

func (s *Service) SelectUserByID(uid int) (models.User, error) {
	var u models.User
	row := s.DB.QueryRow(`SELECT id, name, telephone, password, date_birth, sex, job, education, about_me FROM users
						WHERE  id=$1;`, uid)
	err := row.Scan(&u.ID, &u.Name, &u.Telephone, &u.Password, &u.DateBirth, &u.Sex, &u.Education, &u.Job, &u.AboutMe)
	if err != nil {
		log.Println(err)
		return u, err
	}
	return u, nil
}

func (s *Service) SelectUsers() ([]models.UserFeed, error) {
	var users []models.UserFeed
	rows, err := s.DB.Query(`SELECT id, name, date_birth, job, education, about_me FROM users`)
	if err != nil {
		log.Println(err)
		return users, err
	}
	defer rows.Close()
	for rows.Next() {
		var u models.UserFeed
		err := rows.Scan(&u.ID, &u.Name, &u.DateBirth, &u.Education, &u.Job, &u.AboutMe)
		if err != nil {
			log.Println(err)
			continue
		}
		users = append(users, u)
	}
	users = users[0:5]
	users[0].LinkImages = append(users[0].LinkImages, "/static/avatars/3.jpg")
	users[1].LinkImages = append(users[1].LinkImages, "/static/avatars/4.jpg")
	users[2].LinkImages = append(users[2].LinkImages, "/static/avatars/9.jpg")
	users[3].LinkImages = append(users[3].LinkImages, "/static/avatars/6.jpg")
	users[4].LinkImages = append(users[4].LinkImages, "/static/avatars/7.jpg")
	return users, nil
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

func (s *Service) InsertLike(like models.Like) error {
	_, err := s.DB.Exec(`INSERT INTO likes(user_id1, user_id2) VALUES ($1, $2);`, like.Uid1, like.Uid2)
	if err != nil {
		log.Println(err)
		return err
	}
	return nil
}

func (s *Service) InsertDislike(dislike models.Dislike) error {
	_, err := s.DB.Exec(`INSERT INTO dislikes(user_id1, user_id2) VALUES ($1, $2);`, dislike.Uid1, dislike.Uid2)
	if err != nil {
		log.Println(err)
		return err
	}
	return nil
}

func (s *Service) InsertComment(comment models.Comment) error {
	_, err := s.DB.Exec(`INSERT INTO comments(user_id1, user_id2) VALUES ($1, $2);`, comment.PhotoID, comment.Text)
	if err != nil {
		log.Println(err)
		return err
	}
	return nil
}

func (s *Service) InsertChat(chat models.Chat) error {
	_, err := s.DB.Exec(`INSERT INTO chat(user_id1, user_id2) VALUES ($1, $2);`, chat.Uid1, chat.Uid2)
	if err != nil {
		log.Println(err)
		return err
	}
	return nil
}

func (s *Service) InsertPhoto(photo models.Photo) error {
	_, err := s.DB.Exec(`INSERT INTO photo(path, user_id) VALUES ($1, $2);`, photo.Path, photo.UID)
	if err != nil {
		log.Println(err)
		return err
	}
	return nil
}
