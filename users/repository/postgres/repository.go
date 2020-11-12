package postgres

import (
	"database/sql"
	"park_2020/2020_2_tmp_name/domain"
	"park_2020/2020_2_tmp_name/models"
	"time"

	"fmt"
)

type postgresUserRepository struct {
	Conn *sql.DB
}

func NewPostgresUserRepository(Conn *sql.DB) domain.UserRepository {
	return &postgresUserRepository{Conn}
}

func (p *postgresUserRepository) CheckUser(telephone string) bool {
	var count int
	p.Conn.QueryRow(`SELECT COUNT(telephone) FROM users WHERE telephone=$1;`, telephone).Scan(&count)
	return count > 0
}

func (p *postgresUserRepository) InsertUser(user models.User) error {
	password, err := models.HashPassword(user.Password)
	if err != nil {
		return err
	}
	_, err = p.Conn.Exec(`INSERT INTO users(name, telephone, password, date_birth, sex, job, education, about_me)
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
	return err
}

func (p *postgresUserRepository) SelectUser(telephone string) (models.User, error) {
	var u models.User
	row := p.Conn.QueryRow(`SELECT id, name, telephone, password, date_birth, sex, job, education, about_me FROM users
						WHERE  telephone=$1;`, telephone)
	err := row.Scan(&u.ID, &u.Name, &u.Telephone, &u.Password, &u.DateBirth, &u.Sex, &u.Education, &u.Job, &u.AboutMe)
	if err != nil {
		return u, err
	}

	u.LinkImages, err = p.SelectImages(u.ID)
	return u, err
}

func (p *postgresUserRepository) SelectUserMe(telephone string) (models.UserMe, error) {
	var u models.UserMe
	var date time.Time
	row := p.Conn.QueryRow(`SELECT id, name, telephone, date_birth, job, education, about_me FROM users
						WHERE  telephone=$1;`, telephone)
	err := row.Scan(&u.ID, &u.Name, &u.Telephone, &date, &u.Education, &u.Job, &u.AboutMe)
	if err != nil {
		return u, err
	}

	u.DateBirth = models.Diff(date, time.Now())
	u.LinkImages, err = p.SelectImages(u.ID)
	return u, err
}

func (p *postgresUserRepository) SelectUserFeed(telephone string) (models.UserFeed, error) {
	var u models.UserFeed
	var date time.Time
	row := p.Conn.QueryRow(`SELECT id, name, date_birth, job, education, about_me FROM users
						WHERE  telephone=$1;`, telephone)
	err := row.Scan(&u.ID, &u.Name, &date, &u.Education, &u.Job, &u.AboutMe)
	if err != nil {
		return u, err
	}

	u.DateBirth = models.Diff(date, time.Now())
	u.LinkImages, err = p.SelectImages(u.ID)
	return u, err
}

func (p *postgresUserRepository) SelectUserFeedByID(uid int) (models.UserFeed, error) {
	var u models.UserFeed
	var date time.Time
	row := p.Conn.QueryRow(`SELECT name, date_birth, job, education, about_me FROM users
						WHERE  id=$1;`, uid)
	err := row.Scan(&u.Name, &date, &u.Job, &u.Education, &u.AboutMe)
	if err != nil {
		return u, err
	}
	u.ID = uid

	u.DateBirth = models.Diff(date, time.Now())
	u.LinkImages, err = p.SelectImages(u.ID)
	return u, err
}

func (p *postgresUserRepository) SelectUserByID(uid int) (models.User, error) {
	var u models.User
	row := p.Conn.QueryRow(`SELECT id, name, telephone, password, date_birth, sex, job, education, about_me FROM users
						WHERE  id=$1;`, uid)
	err := row.Scan(&u.ID, &u.Name, &u.Telephone, &u.Password, &u.DateBirth, &u.Sex, &u.Education, &u.Job, &u.AboutMe)
	if err != nil {
		return u, err
	}

	u.LinkImages, err = p.SelectImages(u.ID)
	return u, err
}

func (p *postgresUserRepository) Match(uid1, uid2 int) bool {
	var id1, id2 int
	row := p.Conn.QueryRow(`Select user_id1, user_id2 FROM likes 
							WHERE user_id1 = $1 AND user_id2 = $2;`, uid2, uid1)
	err := row.Scan(&id1, &id2)
	return err == nil
}

func (p *postgresUserRepository) SelectUsers(user models.User) ([]models.UserFeed, error) {
	var users []models.UserFeed
	rows, err := p.Conn.Query(`SELECT id, name, date_birth, job, education, about_me FROM users WHERE sex != $1`, user.Sex)
	if err != nil {
		return users, err
	}
	defer rows.Close()

	for rows.Next() {
		var u models.UserFeed
		var date time.Time
		err := rows.Scan(&u.ID, &u.Name, &date, &u.Education, &u.Job, &u.AboutMe)
		if err != nil {
			continue
		}

		u.DateBirth = models.Diff(date, time.Now())
		u.LinkImages, err = p.SelectImages(u.ID)
		if err != nil {
			return users, err
		}
		users = append(users, u)
	}

	return users, nil
}

func (p *postgresUserRepository) SelectImages(uid int) ([]string, error) {
	var images []string
	rows, err := p.Conn.Query(`SELECT path FROM photo WHERE  user_id=$1;`, uid)
	if err != nil {
		return images, err
	}
	defer rows.Close()

	for rows.Next() {
		var image string
		// var id, uid int
		err := rows.Scan(&image)
		if err != nil {
			continue
		}
		images = append(images, image)
	}
	return images, nil
}

func (p *postgresUserRepository) UpdateUser(user models.User, uid int) error {
	if user.Name != "" {
		_, err := p.Conn.Exec(`UPDATE users SET name=$1 WHERE id = $2;`, user.Name, uid)
		if err != nil {
			return err
		}
	}
	if user.Telephone != "" {
		_, err := p.Conn.Exec(`UPDATE users SET telephone=$1 WHERE id = $2;`, user.Telephone, uid)
		if err != nil {
			return err
		}
	}
	if user.Password != "" {
		password, err := models.HashPassword(user.Password)
		if err != nil {
			return err
		}
		_, err = p.Conn.Exec(`UPDATE users SET password=$1 WHERE id = $2;`, password, uid)
		if err != nil {
			return err
		}
	}
	if user.Job != "" {
		_, err := p.Conn.Exec(`UPDATE users SET job=$1 WHERE id = $2;`, user.Job, uid)
		if err != nil {
			return err
		}
	}
	if user.Education != "" {
		_, err := p.Conn.Exec(`UPDATE users SET education=$1 WHERE id = $2;`, user.Education, uid)
		if err != nil {
			return err
		}
	}
	if user.AboutMe != "" {
		_, err := p.Conn.Exec(`UPDATE users SET about_me=$1 WHERE id = $2;`, user.AboutMe, uid)
		if err != nil {
			return err
		}
	}
	return nil
}

func (p *postgresUserRepository) InsertSession(sid, telephone string) error {
	_, err := p.Conn.Exec(`INSERT INTO sessions(key, value) VALUES ($1, $2);`, sid, telephone)
	return err
}

func (p *postgresUserRepository) DeleteSession(sid string) error {
	_, err := p.Conn.Exec(`DELETE FROM sessions WHERE key=$1;`, sid)
	return err
}

func (p *postgresUserRepository) CheckUserBySession(sid string) string {
	var count string
	p.Conn.QueryRow(`SELECT value FROM sessions WHERE key=$1;`, sid).Scan(&count)
	return count
}

func (p *postgresUserRepository) InsertLike(uid1, uid2 int) error {
	_, err := p.Conn.Exec(`INSERT INTO likes(user_id1, user_id2) VALUES ($1, $2);`, uid1, uid2)
	return err
}

func (p *postgresUserRepository) InsertDislike(uid1, uid2 int) error {
	_, err := p.Conn.Exec(`INSERT INTO dislikes(user_id1, user_id2) VALUES ($1, $2);`, uid1, uid2)
	return err
}

func (p *postgresUserRepository) InsertComment(comment models.Comment, uid int) error {
	_, err := p.Conn.Exec(`INSERT INTO comments(user_id1, user_id2, time_delivery, text) VALUES ($1, $2, $3, $4);`,
		uid, comment.Uid2, time.Now().Format("15:04"), comment.CommentText)
	return err
}

func (p *postgresUserRepository) CheckChat(chat models.Chat) bool {
	var id1, id2 int
	row := p.Conn.QueryRow(`SELECT user_id1, user_id2 FROM chat 
							WHERE user_id1 = $1 AND user_id2 = $2 
							OR user_id1 = $2 AND user_id2 = $1`, chat.Uid1, chat.Uid2)
	err := row.Scan(&id1, &id2)
	return err == nil
}

func (p *postgresUserRepository) InsertChat(chat models.Chat) error {
	_, err := p.Conn.Exec(`INSERT INTO chat(user_id1, user_id2) VALUES ($1, $2);`, chat.Uid1, chat.Uid2)
	return err
}

func (p *postgresUserRepository) InsertMessage(text string, chatID, uid int) error {
	_, err := p.Conn.Exec(`INSERT INTO message(text, time_delivery, chat_id, user_id) VALUES ($1, $2, $3, $4);`, text, time.Now().Format("15:04"), chatID, uid)
	return err
}

func (p *postgresUserRepository) InsertPhoto(path string, uid int) error {
	_, err := p.Conn.Exec(`INSERT INTO photo(path, user_id) VALUES ($1, $2);`, path, uid)
	return err
}

func (p *postgresUserRepository) SelectMessage(uid, chid int) (models.Msg, error) {
	var message models.Msg
    row := p.Conn.QueryRow(`SELECT text, time_delivery, user_id FROM message WHERE user_id=$1 AND chat_id=$2 order by id desc limit 1;`, uid, chid)
    row.Scan(&message.Message, &message.TimeDelivery, &message.UserID)
    return message, nil
}

func (p *postgresUserRepository) SelectMessages(chid int) ([]models.Msg, error) {
	var messages []models.Msg
	rows, err := p.Conn.Query(`SELECT text, time_delivery, user_id FROM message WHERE chat_id=$1 order by id asc limit 10;`, chid)
	if err != nil {
		return messages, err
	}
	defer rows.Close()

	for rows.Next() {
		var message models.Msg
		// var id int
		err := rows.Scan(&message.Message, &message.TimeDelivery, &message.UserID)
		if err != nil {
			continue
		}
		messages = append(messages, message)
	}

	return messages, nil
}

func (p *postgresUserRepository) SelectChatsByID(uid int) ([]models.ChatData, error) {
	var chats []models.ChatData
	rows, err := p.Conn.Query(`SELECT id, user_id1 FROM chat WHERE user_id2=$1;`, uid)
	if err != nil {
		return chats, err
	}
	defer rows.Close()

	for rows.Next() {
		var chat models.ChatData
		var uid1 int
		err := rows.Scan(&chat.ID, &uid1)
		if err != nil {
			continue
		}
		chat.Partner, err = p.SelectUserFeedByID(uid1)
		if err != nil {
			fmt.Println("Select user feed")
			return chats, err
		}
		msg, err := p.SelectMessage(uid1, chat.ID)
		if err != nil {
			fmt.Println("Select message")
			return chats, err
		}
		chat.Messages = append(chat.Messages, msg)
		chats = append(chats, chat)
	}

	rows, err = p.Conn.Query(`SELECT id, user_id2 FROM chat WHERE user_id1=$1;`, uid)
	if err != nil {
		return chats, err
	}
	defer rows.Close()
	for rows.Next() {
		var chat models.ChatData
		var uid2 int
		err := rows.Scan(&chat.ID, &uid2)
		if err != nil {
			continue
		}

		chat.Partner, err = p.SelectUserFeedByID(uid2)
		if err != nil {
			fmt.Println("Select user feed2")
			return chats, err
		}
		msg, err := p.SelectMessage(uid2, chat.ID)
		if err != nil {
			fmt.Println("Select message2")
			return chats, err
		}
		chat.Messages = append(chat.Messages, msg)
		chats = append(chats, chat)
	}

	return chats, nil
}

func (p *postgresUserRepository) SelectChatByID(uid, chid int) (models.ChatData, error) {
	var chat models.ChatData
	chat.ID = chid
	var err error

	chat.Partner, err = p.SelectUserByChat(uid, chid)
	if err != nil {
		return chat, err
	}

	chat.Messages, err = p.SelectMessages(chid)
	return chat, err
}

func (p *postgresUserRepository) SelectUserByChat(uid, chid int) (models.UserFeed, error) {
	var user models.UserFeed
	var id1, id2, id int

	row := p.Conn.QueryRow(`SELECT user_id1, user_id2 FROM chat WHERE id=$1;`, chid)

	err := row.Scan(&id1, &id2)
	if err != nil {
		return user, err
	}

	if id1 != uid {
		id = id1
	} else {
		id = id2
	}

	user, err = p.SelectUserFeedByID(id)
	return user, err
}

func (p *postgresUserRepository) SelectComments(userId int) (models.CommentsById, error) {
	var result models.CommentsById
	var comments []models.CommentId
	comments = make([]models.CommentId, 0, 1)

	rows, err := p.Conn.Query(`SELECT user_id1, text, time_delivery FROM comments
						WHERE  user_id2=$1;`, userId)
	if err != nil {
		return result, err
	}
	defer rows.Close()

	for rows.Next() {
		var comment models.CommentId
		err := rows.Scan(&comment.UserId, &comment.CommentText, &comment.TimeDelivery)
		if err != nil {
			return result, err
		}
		comments = append(comments, comment)
	}

	// TODO: вынести в отдельную функцию бизнес-логики
	for _, comment := range comments {
		user, err := p.SelectUserFeedByID(comment.UserId)
		if err != nil {
			return result, err
		}
		var res models.CommentById
		res.User = user
		res.CommentText = comment.CommentText
		res.TimeDelivery = comment.TimeDelivery
		result.Comments = append(result.Comments, res)
	}
	return result, nil
}
