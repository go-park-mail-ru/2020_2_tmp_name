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

	_, err = p.Conn.Exec(`INSERT INTO users(name, telephone, password, date_birth, sex, job, education, about_me)
						VALUES ($1, $2, $3, $4, $5, $6, $7, $8);`,
		user.Name,
		user.Telephone,
		password,
		age,
		user.Sex,
		user.Job,
		user.Education,
		user.AboutMe,
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
	row := p.Conn.QueryRow(`SELECT id, name, telephone, password, date_birth, sex, job, education, about_me FROM users
						WHERE  id=$1;`, uid)
	err := row.Scan(&u.ID, &u.Name, &u.Telephone, &u.Password, &u.DateBirth, &u.Sex, &u.Education, &u.Job, &u.AboutMe)
	if err != nil {
		return u, err
	}

	u.LinkImages, err = p.SelectImages(u.ID)
	return u, err
}

func (p *postgresUserRepository) SelectUserFeedByID(uid int) (models.UserFeed, error) {
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

func (p *postgresUserRepository) Match(uid1, uid2 int) bool {
	var id1, id2 int
	row := p.Conn.QueryRow(`Select user_id1, user_id2 FROM likes 
							WHERE user_id1 = $1 AND user_id2 = $2;`, uid2, uid1)
	err := row.Scan(&id1, &id2)
	return err == nil
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
	// rows, err := p.Conn.Query(`SELECT id, name, date_birth, education, job,  about_me FROM users WHERE sex != $1`, user.Sex)
	rows, err := p.Conn.Query(`SELECT u.id, u.name, u.date_birth, u.education, u.job, u.about_me FROM users as u
								where u.sex != $1
								except (
								SELECT u.id, u.name, u.date_birth, u.education, u.job, u.about_me FROM users as u
								join likes as l on u.id=l.user_id2 where u.sex != $1 and l.user_id1=$2
								union
								SELECT u.id, u.name, u.date_birth, u.education, u.job, u.about_me FROM users as u
								join dislikes as d on u.id=d.user_id2 where u.sex != $1 and d.user_id1=$2
								);`, user.Sex, user.ID)
	if err != nil {
		return users, err
	}
	defer rows.Close()

	for rows.Next() {
		var u models.UserFeed
		err := rows.Scan(&u.ID, &u.Name, &u.DateBirth, &u.Education, &u.Job, &u.AboutMe)
		if err != nil {
			continue
		}

		u.LinkImages, err = p.SelectImages(u.ID)
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
			continue
		}
		images = append(images, image)
	}
	return images, nil
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

	fmt.Println("--------------------------------------------------------__*******************")
	fmt.Println(count)

	return count > 0
}
