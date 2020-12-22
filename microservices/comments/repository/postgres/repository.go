package postgres

import (
	"database/sql"
	domain "park_2020/2020_2_tmp_name/microservices/comments"
	"park_2020/2020_2_tmp_name/models"
	"time"
)

type postgresCommentRepository struct {
	Conn *sql.DB
}

func NewPostgresCommentRepository(Conn *sql.DB) domain.CommentRepository {
	return &postgresCommentRepository{Conn}
}

func (p *postgresCommentRepository) SelectUserFeed(telephone string) (models.UserFeed, error) {
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

func (p *postgresCommentRepository) SelectUserFeedByID(uid int) (models.UserFeed, error) {
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

func (p *postgresCommentRepository) SelectImages(uid int) ([]string, error) {
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

func (p *postgresCommentRepository) InsertComment(comment models.Comment, uid int) error {
	_, err := p.Conn.Exec(`INSERT INTO comments(user_id1, user_id2, time_delivery, text) VALUES ($1, $2, $3, $4);`,
		uid, comment.Uid2, time.Now().Format("15:04"), comment.CommentText)
	return err
}

func (p *postgresCommentRepository) SelectComments(userId int) (models.CommentsById, error) {
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
