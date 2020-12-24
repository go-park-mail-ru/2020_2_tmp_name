package postgres

import (
	"database/sql"
	domain "park_2020/2020_2_tmp_name/api/chats"
	"park_2020/2020_2_tmp_name/models"
	"time"
)

type postgresChatRepository struct {
	Conn *sql.DB
}

func NewPostgresChatRepository(Conn *sql.DB) domain.ChatRepository {
	return &postgresChatRepository{Conn}
}

func (p *postgresChatRepository) SelectUserFeed(telephone string) (models.UserFeed, error) {
	var u models.UserFeed
	var tid int
	row := p.Conn.QueryRow(`SELECT id, name, date_birth, education, job, about_me, filter_id FROM users
						WHERE  telephone=$1;`, telephone)
	err := row.Scan(&u.ID, &u.Name, &u.DateBirth, &u.Education, &u.Job, &u.AboutMe, &tid)
	if err != nil {
		return u, err
	}

	u.LinkImages, err = p.SelectImages(u.ID)
	u.Target = models.IDToTarget(tid)
	return u, err
}

func (p *postgresChatRepository) SelectImages(uid int) ([]string, error) {
	var images []string
	rows, err := p.Conn.Query(`SELECT path FROM photo WHERE  user_id=$1;`, uid)
	if err != nil {
		return images, err
	}
	defer rows.Close()

	for rows.Next() {
		var image string
		err := rows.Scan(&image)
		if err != nil {
			return images, err
		}
		images = append(images, image)
	}
	return images, nil
}

func (p *postgresChatRepository) CheckChat(chat models.Chat) bool {
	var id1, id2 int
	row := p.Conn.QueryRow(`SELECT user_id1, user_id2 FROM chat 
							WHERE user_id1 = $1 AND user_id2 = $2 
							OR user_id1 = $2 AND user_id2 = $1`, chat.Uid1, chat.Uid2)
	err := row.Scan(&id1, &id2)
	return err == nil
}

func (p *postgresChatRepository) InsertChat(chat models.Chat) error {
	_, err := p.Conn.Exec(`INSERT INTO chat(user_id1, user_id2, filter_id) VALUES ($1, $2, $3);`,
		chat.Uid1, chat.Uid2, models.TargetToID(chat.Target))
	return err
}

func (p *postgresChatRepository) InsertMessage(text string, chatID, uid int) error {
	_, err := p.Conn.Exec(`INSERT INTO message(text, time_delivery, chat_id, user_id) VALUES ($1, $2, $3, $4);`, text, time.Now().Format("15:04"), chatID, uid)
	return err
}

func (p *postgresChatRepository) SelectMessage(uid, chid int) (models.Msg, error) {
	var message models.Msg
	row := p.Conn.QueryRow(`SELECT text, time_delivery, user_id FROM message WHERE chat_id=$1 order by id desc limit 1;`, chid)
	err := row.Scan(&message.Message, &message.TimeDelivery, &message.UserID)
	if err != nil {
		return message, err
	}

	message.ChatID = chid
	message.UserID = uid
	return message, nil
}

func (p *postgresChatRepository) SelectMessages(chid int) ([]models.Msg, error) {
	var messages []models.Msg
	rows, err := p.Conn.Query(`SELECT m.text, m.time_delivery, m.user_id FROM (SELECT * FROM message WHERE chat_id=$1 ORDER BY id DESC limit 10)
								AS m ORDER BY m.id ASC;`, chid)
	if err != nil {
		return messages, err
	}
	defer rows.Close()

	for rows.Next() {
		var message models.Msg
		err := rows.Scan(&message.Message, &message.TimeDelivery, &message.UserID)
		if err != nil {
			return messages, err
		}
		message.ChatID = chid
		messages = append(messages, message)
	}

	return messages, nil
}

func (p *postgresChatRepository) SelectChatsByID(uid int) ([]models.ChatData, error) {
	var chats []models.ChatData
	rows, err := p.Conn.Query(`SELECT id, user_id1, filter_id FROM chat WHERE user_id2=$1;`, uid)
	if err != nil {
		return chats, err
	}
	defer rows.Close()

	for rows.Next() {
		var chat models.ChatData
		var uid1, fid int
		err := rows.Scan(&chat.ID, &uid1, &fid)
		if err != nil {
			return chats, err
		}
		chat.Target = models.IDToTarget(fid)
		chat.Partner, err = p.SelectUserFeedByID(uid1)
		if err != nil {
			return chats, err
		}
		msg, err := p.SelectMessage(uid1, chat.ID)
		if err != nil {
			return chats, err
		}
		chat.Messages = append(chat.Messages, msg)
		chats = append(chats, chat)
	}

	rows, err = p.Conn.Query(`SELECT id, user_id2, filter_id FROM chat WHERE user_id1=$1;`, uid)
	if err != nil {
		return chats, err
	}
	defer rows.Close()
	for rows.Next() {
		var chat models.ChatData
		var uid2, fid int
		err := rows.Scan(&chat.ID, &uid2, &fid)
		if err != nil {
			return chats, err
		}
		chat.Target = models.IDToTarget(fid)
		chat.Partner, err = p.SelectUserFeedByID(uid2)
		if err != nil {
			return chats, err
		}
		msg, err := p.SelectMessage(uid2, chat.ID)
		if err != nil {
			return chats, err
		}
		chat.Messages = append(chat.Messages, msg)
		chats = append(chats, chat)
	}

	return chats, nil
}

func (p *postgresChatRepository) SelectChatByID(uid, chid int) (models.ChatData, error) {
	var chat models.ChatData
	chat.ID = chid
	var err error

	chat.Partner, err = p.SelectUserByChat(uid, chid)
	if err != nil {
		return chat, err
	}

	chat.Messages, err = p.SelectMessages(chid)
	chat.Target = models.IDToTarget(uid)
	return chat, err
}

func (p *postgresChatRepository) SelectUserByChat(uid, chid int) (models.UserFeed, error) {
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

func (p *postgresChatRepository) SelectUserFeedByID(uid int) (models.UserFeed, error) {
	var u models.UserFeed
	var tid int
	row := p.Conn.QueryRow(`SELECT name, date_birth, job, education, about_me, filter_id FROM users
						WHERE  id=$1;`, uid)
	err := row.Scan(&u.Name, &u.DateBirth, &u.Job, &u.Education, &u.AboutMe, &tid)
	if err != nil {
		return u, err
	}
	u.ID = uid

	u.LinkImages, err = p.SelectImages(u.ID)
	u.Target = models.IDToTarget(tid)
	return u, err
}

func (p *postgresChatRepository) SelectUserByID(uid int) (models.User, error) {
	var u models.User
	var tid int
	row := p.Conn.QueryRow(`SELECT id, name, telephone, password, date_birth, sex, job, education, about_me, filter_id FROM users
						WHERE  id=$1;`, uid)
	err := row.Scan(&u.ID, &u.Name, &u.Telephone, &u.Password, &u.DateBirth, &u.Sex, &u.Education, &u.Job, &u.AboutMe, &tid)
	if err != nil {
		return u, err
	}

	u.LinkImages, err = p.SelectImages(u.ID)
	u.Target = models.IDToTarget(tid)
	return u, err
}

func (p *postgresChatRepository) CheckUserBySession(sid string) string {
	var str string
	err := p.Conn.QueryRow(`SELECT value FROM sessions WHERE key=$1;`, sid).Scan(&str)
	if err != nil {
		return ""
	}
	return str
}

func (p *postgresChatRepository) SelectSessions(uid int) ([]string, error) {
	var sessions []string
	user, err := p.SelectUserByID(uid)
	if err != nil {
		return sessions, err
	}

	rows, err := p.Conn.Query(`SELECT key FROM sessions WHERE value=$1;`, user.Telephone)
	if err != nil {
		return sessions, err
	}
	defer rows.Close()

	for rows.Next() {
		var session string
		err := rows.Scan(&session)
		if err != nil {
			return sessions, err
		}
		sessions = append(sessions, session)
	}
	return sessions, nil
}

func (p *postgresChatRepository) Match(uid1, uid2, fid int) bool {
	var id1, id2 int
	row := p.Conn.QueryRow(`Select user_id1, user_id2 FROM likes 
							WHERE user_id1 = $1 AND user_id2 = $2 AND filter_id = $3;`, uid2, uid1, fid)
	err := row.Scan(&id1, &id2)
	return err == nil
}

func (p *postgresChatRepository) InsertLike(uid1, uid2, fid int) error {
	_, err := p.Conn.Exec(`INSERT INTO likes(user_id1, user_id2, filter_id) VALUES ($1, $2, $3);`, uid1, uid2, fid)
	return err
}

func (p *postgresChatRepository) InsertDislike(uid1, uid2, fid int) error {
	_, err := p.Conn.Exec(`INSERT INTO dislikes(user_id1, user_id2, filter_id) VALUES ($1, $2, $3);`, uid1, uid2, fid)
	return err
}

func (p *postgresChatRepository) CheckLike(uid1, uid2 int) bool {
	var count int
	err := p.Conn.QueryRow(`SELECT COUNT(id) FROM likes WHERE user_id1=$1 AND user_id2 = $2;`, uid1, uid2).Scan(&count)
	if err != nil {
		return false
	}
	return count > 0
}

func (p *postgresChatRepository) CheckDislike(uid1, uid2 int) bool {
	var count int
	err := p.Conn.QueryRow(`SELECT COUNT(id) FROM dislikes WHERE user_id1=$1 AND user_id2 = $2;`, uid1, uid2).Scan(&count)
	if err != nil {
		return false
	}
	return count > 0
}

func (p *postgresChatRepository) DeleteLike(uid1, uid2 int) error {
	_, err := p.Conn.Exec(`DELETE FROM likes WHERE user_id1=$1 AND user_id2=$2;`, uid1, uid2)
	return err
}

func (p *postgresChatRepository) DeleteDislike(uid1, uid2 int) error {
	_, err := p.Conn.Exec(`DELETE FROM dislikes WHERE user_id1=$1 AND user_id2=$2;`, uid1, uid2)
	return err
}

func (p *postgresChatRepository) SelectChatID(uid1, uid2 int) (int, error) {
	var chid int
	row := p.Conn.QueryRow(`SELECT id FROM chat WHERE user_id1=$1 AND user_id2=$2;`, uid1, uid2)
	err := row.Scan(&chid)
	return chid, err
}

func (p *postgresChatRepository) InsertSuperlike(uid1, uid2, fid int) error {
	_, err := p.Conn.Exec(`INSERT INTO superlikes(user_id1, user_id2, filter_id) VALUES ($1, $2, $3);`, uid1, uid2, fid)
	return err
}
