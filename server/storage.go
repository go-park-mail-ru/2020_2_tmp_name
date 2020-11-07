package server

import (
	"log"
	"park_2020/2020_2_tmp_name/models"
	"time"
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

	u.LinkImages, err = s.SelectImages(u.ID)
	if err != nil {
		log.Println(err)
		return u, err
	}

	return u, nil
}

func (s *Service) SelectUserFeed(telephone string) (models.UserFeed, error) {
	var u models.UserFeed
	var date time.Time
	row := s.DB.QueryRow(`SELECT id, name, date_birth, job, education, about_me FROM users
						WHERE  telephone=$1;`, telephone)
	err := row.Scan(&u.ID, &u.Name, &date, &u.Education, &u.Job, &u.AboutMe)
	if err != nil {
		log.Println(err)
		return u, err
	}

	u.DateBirth = diff(date, time.Now())
	u.LinkImages, err = s.SelectImages(u.ID)
	if err != nil {
		log.Println(err)
		return u, err
	}

	return u, nil
}

func (s *Service) SelectUserFeedByID(uid int) (models.UserFeed, error) {
	var u models.UserFeed
	var date time.Time
	row := s.DB.QueryRow(`SELECT name, date_birth, job, education, about_me FROM users
						WHERE  id=$1;`, uid)
	err := row.Scan(&u.Name, &date, &u.Job, &u.Education, &u.AboutMe)
	if err != nil {
		log.Println(err)
		return u, err
	}
	u.ID = uid

	u.DateBirth = diff(date, time.Now())
	u.LinkImages, err = s.SelectImages(u.ID)
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

	u.LinkImages, err = s.SelectImages(u.ID)
	if err != nil {
		log.Println(err)
		return u, err
	}

	return u, nil
}

func (s *Service) SelectUsers(user models.User) ([]models.UserFeed, error) {
	var users []models.UserFeed
	rows, err := s.DB.Query(`SELECT id, name, date_birth, job, education, about_me FROM users WHERE sex != $1`, user.Sex)
	if err != nil {
		log.Println(err)
		return users, err
	}
	defer rows.Close()

	for rows.Next() {
		var u models.UserFeed
		var date time.Time
		err := rows.Scan(&u.ID, &u.Name, &date, &u.Education, &u.Job, &u.AboutMe)
		if err != nil {
			log.Println(err)
			continue
		}

		u.DateBirth = diff(date, time.Now())
		u.LinkImages, err = s.SelectImages(u.ID)
		if err != nil {
			log.Println(err)
			return users, err
		}

		users = append(users, u)
	}

	return users, nil
}

func (s *Service) SelectImages(uid int) ([]string, error) {
	var images []string
	rows, err := s.DB.Query(`SELECT path FROM photo WHERE  user_id=$1;`, uid)
	if err != nil {
		log.Println(err)
		return images, err
	}
	defer rows.Close()
	for rows.Next() {
		var image string
		err := rows.Scan(&image)
		if err != nil {
			log.Println(err)
			continue
		}
		images = append(images, image)
	}
	return images, nil
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

func (s *Service) InsertLike(uid1, uid2 int) error {
	_, err := s.DB.Exec(`INSERT INTO likes(user_id1, user_id2) VALUES ($1, $2);`, uid1, uid2)
	if err != nil {
		log.Println(err)
		return err
	}
	return nil
}

func (s *Service) InsertDislike(uid1, uid2 int) error {
	_, err := s.DB.Exec(`INSERT INTO dislikes(user_id1, user_id2) VALUES ($1, $2);`, uid1, uid2)
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

func (s *Service) InsertMessage(text string, chatID, uid int) error {
	_, err := s.DB.Exec(`INSERT INTO message(text, time_delivery, chat_id, user_id) VALUES ($1, $2, $3, $4);`, text, time.Now(), chatID, uid)
	if err != nil {
		log.Println(err)
		return err
	}
	return nil
}

func (s *Service) InsertPhoto(path string, uid int) error {
	_, err := s.DB.Exec(`INSERT INTO photo(path, user_id) VALUES ($1, $2);`, path, uid)
	if err != nil {
		log.Println(err)
		return err
	}
	return nil
}

func (s *Service) SelectMessage(uid, chid int) (models.Msg, error) {
	var message models.Msg
	row := s.DB.QueryRow(`SELECT text, time_delivery, user_id FROM message WHERE user_id=$1 AND chat_id=$2 order by time_delivery desc limit 1;`, uid, chid)
	err := row.Scan(&message.Message, &message.TimeDelivery, &message.UserID)
	if err != nil {
		log.Println(err)
		return message, nil
	}

	return message, nil
}

func (s *Service) SelectMessages(chid int) ([]models.Msg, error) {
	var messages []models.Msg
	rows, err := s.DB.Query(`SELECT text, time_delivery, user_id FROM message WHERE chat_id=$1 order by time_delivery desc limit 10;`, chid)
	if err != nil {
		log.Println(err)
		return messages, err
	}
	defer rows.Close()

	for rows.Next() {
		var message models.Msg
		err := rows.Scan(&message.Message, &message.TimeDelivery, &message.UserID)
		if err != nil {
			log.Println(err)
			continue
		}
		messages = append(messages, message)
	}

	return messages, nil
}

func (s *Service) SelectChatsByID(uid int) ([]models.ChatData, error) {
	var chats []models.ChatData
	rows, err := s.DB.Query(`SELECT id, user_id1 FROM chat WHERE user_id2=$1;`, uid)
	if err != nil {
		log.Println(err)
		return chats, err
	}
	defer rows.Close()

	for rows.Next() {
		var chat models.ChatData
		var uid1 int
		err := rows.Scan(&chat.ID, &uid1)
		if err != nil {
			log.Println(err)
			continue
		}
		chat.Partner, err = s.SelectUserFeedByID(uid1)
		if err != nil {
			log.Println(err)
			return chats, err
		}
		msg, err := s.SelectMessage(uid1, chat.ID)
		if err != nil {
			log.Println(err)
			return chats, err
		}
		chat.Messages = append(chat.Messages, msg)
		chats = append(chats, chat)
	}

	rows, err = s.DB.Query(`SELECT id, user_id2 FROM chat WHERE user_id1=$1;`, uid)
	if err != nil {
		log.Println(err)
		return chats, err
	}
	defer rows.Close()

	for rows.Next() {
		var chat models.ChatData
		var uid2 int
		err := rows.Scan(&chat.ID, &uid2)
		if err != nil {
			log.Println("err")
			continue
		}

		chat.Partner, err = s.SelectUserFeedByID(uid2)
		if err != nil {
			log.Println(err)
			return chats, err
		}
		msg, err := s.SelectMessage(uid2, chat.ID)
		if err != nil {
			log.Println(err)
			return chats, err
		}
		chat.Messages = append(chat.Messages, msg)
		chats = append(chats, chat)
	}

	return chats, nil
}

func (s *Service) SelectChatByID(uid, chid int) (models.ChatData, error) {
	var chat models.ChatData
	chat.ID = chid
	var err error

	chat.Partner, err = s.SelectUserFeedByID(uid)
	if err != nil {
		log.Println(err)
		return chat, err
	}

	chat.Messages, err = s.SelectMessages(chid)
	if err != nil {
		log.Println(err)
		return chat, err
	}

	return chat, nil
}
