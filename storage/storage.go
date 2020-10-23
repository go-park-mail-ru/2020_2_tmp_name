package storage

import (
	"database/sql"
	"fmt"
	"log"
	"park_2020/2020_2_tmp_name/models"

	_ "github.com/lib/pq"
)

func DBConnection(conf *models.Config) (*sql.DB, error) {
	connString := fmt.Sprintf("host=%v user=%v password=%v dbname=%v sslmode=disable",
		conf.SQLDataBase.Server,
		conf.SQLDataBase.UserID,
		conf.SQLDataBase.Password,
		conf.SQLDataBase.Database,
	)

	db, err := sql.Open("postgres", connString)
	if err != nil {
		log.Println(err)
		return nil, err
	}

	db.SetMaxOpenConns(10)

	err = db.Ping()
	if err != nil {
		log.Println(err)
		return nil, err
	}

	return db, nil
}
