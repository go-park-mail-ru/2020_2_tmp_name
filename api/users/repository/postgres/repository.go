package postgres

import (
	"database/sql"
	"fmt"
	domain "park_2020/2020_2_tmp_name/api/users"
	"park_2020/2020_2_tmp_name/models"
	"time"
)

type postgresUserRepository struct {
	Conn *sql.DB
}

func NewPostgresUserRepository(Conn *sql.DB) domain.UserRepository {
	return &postgresUserRepository{Conn}
}

func (p *postgresUserRepository) CheckUser(telephone string) bool {
	var count int
	err := p.Conn.QueryRow(`SELECT COUNT(telephone) FROM users WHERE telephone=$1;`, telephone).Scan(&count)
	if err != nil {
		return false
	}
	return count > 0
}

func (p *postgresUserRepository) InsertUser(user models.User) error {
	password, err := models.HashPassword(user.Password)
	if err != nil {
		return err
	}

	age, err := models.Age(user.Day, user.Month, user.Year)
	if err != nil {
		return err
	}

	_, err = p.Conn.Exec(`INSERT INTO users(name, telephone, password, date_birth, sex, job, education, about_me, filter_id)
						VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9);`,
		user.Name,
		user.Telephone,
		password,
		age,
		user.Sex,
		user.Job,
		user.Education,
		user.AboutMe,
		models.TargetToID("love"),
	)
	if err != nil {
		return err
	}

	var uid int
	row := p.Conn.QueryRow(`SELECT id FROM users WHERE  telephone=$1;`, user.Telephone)
	err = row.Scan(&uid)
	if err != nil {
		return err
	}

	for _, path := range user.LinkImages {
		_, err = p.Conn.Exec(`INSERT INTO photo(path, user_id) VALUES ($1, $2);`, path, uid)
		if err != nil {
			return err
		}
	}

	return nil
}

func (p *postgresUserRepository) SelectUserByID(uid int) (models.User, error) {
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

func (p *postgresUserRepository) SelectUserFeedByID(uid int) (models.UserFeed, error) {
	var u models.UserFeed
	var tid int
	row := p.Conn.QueryRow(`SELECT name, date_birth, job, education, about_me, filter_id FROM users
						WHERE  id=$1;`, uid)
	err := row.Scan(&u.Name, &u.DateBirth, &u.Job, &u.Education, &u.AboutMe, &tid)
	if err != nil {
		return u, err
	}
	u.ID = uid
	u.Target = models.IDToTarget(tid)
	u.LinkImages, err = p.SelectImages(u.ID)
	return u, err
}

func (p *postgresUserRepository) CheckPremium(uid int) bool {
	var count int
	err := p.Conn.QueryRow(`SELECT COUNT(user_id) FROM premium_accounts WHERE user_id=$1;`, uid).Scan(&count)
	if err != nil {
		return false
	}
	return count > 0
}

func (p *postgresUserRepository) SelectUsers(user models.User) ([]models.UserFeed, error) {
	var users []models.UserFeed
	var rows *sql.Rows
	var err error
	fmt.Println(user.Sex, user.ID, models.TargetToID(user.Target))
	if user.Target == "love" {
		rows, err = p.Conn.Query(`SELECT u.id, u.name, u.date_birth, u.education, u.job, u.about_me, u.filter_id FROM users AS u
								WHERE u.sex != $1 AND u.filter_id=$3 AND u.id != $2
								EXCEPT (
								SELECT u.id, u.name, u.date_birth, u.education, u.job, u.about_me, u.filter_id FROM users AS u
								JOIN likes AS l ON u.id=l.user_id2 WHERE u.sex != $1 AND l.user_id1=$2 AND u.filter_id=$3
								UNION
								SELECT u.id, u.name, u.date_birth, u.education, u.job, u.about_me, u.filter_id FROM users AS u
								JOIN dislikes AS d ON u.id=d.user_id2 WHERE u.sex != $1 AND d.user_id1=$2 AND u.filter_id=$3
								);`, user.Sex, user.ID, models.TargetToID(user.Target))
	} else {
		rows, err = p.Conn.Query(`SELECT u.id, u.name, u.date_birth, u.education, u.job, u.about_me, u.filter_id FROM users AS u
								WHERE u.filter_id=$2 AND u.id != $1
								EXCEPT (
								SELECT u.id, u.name, u.date_birth, u.education, u.job, u.about_me, u.filter_id FROM users AS u
								JOIN likes AS l ON u.id=l.user_id2 WHERE l.user_id1=$1 AND u.filter_id=$2
								UNION
								SELECT u.id, u.name, u.date_birth, u.education, u.job, u.about_me, u.filter_id FROM users AS u
								JOIN dislikes AS d ON u.id=d.user_id2 WHERE d.user_id1=$1 AND u.filter_id=$2
								);`, user.ID, models.TargetToID(user.Target))
	}
	if err != nil {
		return users, err
	}
	defer rows.Close()

	for rows.Next() {
		var u models.UserFeed
		var tid int
		err := rows.Scan(&u.ID, &u.Name, &u.DateBirth, &u.Education, &u.Job, &u.AboutMe, &tid)
		if err != nil {
			return users, err
		}

		u.LinkImages, err = p.SelectImages(u.ID)
		u.Target = models.IDToTarget(tid)
		if err != nil {
			return users, err
		}
		users = append(users, u)

	}

	return users, nil
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
	if user.Target != "" {
		_, err := p.Conn.Exec(`UPDATE users SET filter_id=$1 WHERE id = $2;`, models.TargetToID(user.Target), uid)
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

func (p *postgresUserRepository) SelectImages(uid int) ([]string, error) {
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

func (p *postgresUserRepository) ChangeAvatarPath(uid int, newpath string) error {
	row := p.Conn.QueryRow(`SELECT path FROM photo WHERE user_id=$1 ORDER BY id LIMIT 1;`, uid)

	var path string
	err := row.Scan(&path)
	if err != nil {
		return err
	}

	tempPath := newpath
	_, err = p.Conn.Exec(`UPDATE path SET path=$1 WHERE path=$2`, newpath, path)
	if err != nil {
		return err
	}

	_, err = p.Conn.Exec(`UPDATE path SET path=$1 WHERE path=$2`, path, tempPath)
	if err != nil {
		return err
	}

	return nil
}

func (p *postgresUserRepository) InsertPremium(uid int, dateFrom time.Time, dateTo time.Time) error {
	_, err := p.Conn.Exec(`INSERT INTO premium_accounts(user_id, date_to, date_from) 
								 VALUES ($1, $2, $3);`, uid, dateTo, dateFrom)

	return err
}

func (p *postgresUserRepository) CheckSuperLikeMe(me, userId int) bool {
	var count int
	err := p.Conn.QueryRow(`SELECT COUNT(*) FROM superlikes WHERE user_id2 = $1;`, me).Scan(&count)
	if err != nil {
		return false
	}

	return count > 0
}
