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
	row := p.Conn.QueryRow(`SELECT id, name, date_birth, education, job, about_me FROM users
						WHERE  telephone=$1;`, telephone)
	err := row.Scan(&u.ID, &u.Name, &u.DateBirth, &u.Education, &u.Job, &u.AboutMe)
	if err != nil {
		return u, err
	}

	u.LinkImages, err = p.SelectImages(u.ID)
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
			continue
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
	_, err := p.Conn.Exec(`INSERT INTO chat(user_id1, user_id2) VALUES ($1, $2);`, chat.Uid1, chat.Uid2)
	return err
}

func (p *postgresChatRepository) InsertMessage(text string, chatID, uid int) error {
	_, err := p.Conn.Exec(`INSERT INTO message(text, time_delivery, chat_id, user_id) VALUES ($1, $2, $3, $4);`, text, time.Now().Format("15:04"), chatID, uid)
	return err
}

func (p *postgresChatRepository) SelectMessage(uid, chid int) (models.Msg, error) {
	var message models.Msg
	row := p.Conn.QueryRow(`SELECT text, time_delivery, user_id FROM message WHERE user_id=$1 AND chat_id=$2 order by id desc limit 1;`, uid, chid)
	row.Scan(&message.Message, &message.TimeDelivery, &message.UserID)
	message.ChatID = chid
	message.UserID = uid
	return message, nil
}

func (p *postgresChatRepository) SelectMessages(chid int) ([]models.Msg, error) {
	var messages []models.Msg
	rows, err := p.Conn.Query(`SELECT text, time_delivery, user_id FROM message WHERE chat_id=$1 order by id asc limit 10;`, chid)
	if err != nil {
		return messages, err
	}
	defer rows.Close()

	for rows.Next() {
		var message models.Msg
		err := rows.Scan(&message.Message, &message.TimeDelivery, &message.UserID)
		if err != nil {
			continue
		}
		message.ChatID = chid
		messages = append(messages, message)
	}

	return messages, nil
}

func (p *postgresChatRepository) SelectChatsByID(uid int) ([]models.ChatData, error) {
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
			return chats, err
		}
		msg, err := p.SelectMessage(uid1, chat.ID)
		if err != nil {
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
	row := p.Conn.QueryRow(`SELECT name, date_birth, job, education, about_me FROM users
						WHERE  id=$1;`, uid)
	err := row.Scan(&u.Name, &u.DateBirth, &u.Job, &u.Education, &u.AboutMe)
	if err != nil {
		return u, err
	}
	u.ID = uid

	u.LinkImages, err = p.SelectImages(u.ID)
	return u, err
}

func (p *postgresChatRepository) CheckUserBySession(sid string) string {
	var count string
	p.Conn.QueryRow(`SELECT value FROM sessions WHERE key=$1;`, sid).Scan(&count)
	return count
}