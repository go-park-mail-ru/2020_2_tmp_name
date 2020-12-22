package models

import (
	"log"
	"strconv"
	"time"

	"golang.org/x/crypto/bcrypt"
)

func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}

func CheckPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

func Age(Day, Month, Year string) (int, error) {
	d1, err := strconv.Atoi(Day)
	if err != nil {
		log.Println(err)
		return 0, err
	}

	var M1 int

	switch Month {

	case "Январь":
		M1 = 1
	case "Февраль":
		M1 = 2
	case "Март":
		M1 = 3
	case "Апрель":
		M1 = 4
	case "Май":
		M1 = 5
	case "Июнь":
		M1 = 6
	case "Июль":
		M1 = 7
	case "Август":
		M1 = 8
	case "Сентябрь":
		M1 = 9
	case "Октябрь":
		M1 = 10
	case "Ноябрь":
		M1 = 11
	case "Декабрь":
		M1 = 12
	}

	y1, err := strconv.Atoi(Year)
	if err != nil {
		log.Println(err)
		return 0, err
	}

	y2, M2, d2 := time.Now().Date()

	year := int(y2 - y1)
	month := int(int(M2) - M1)
	day := int(d2 - d1)

	if day < 0 {
		t := time.Date(y1, time.Month(M1), 32, 0, 0, 0, 0, time.UTC)
		day += 32 - t.Day()
		month--
	}
	if month < 0 {
		month += 12
		year--
	}

	return year, nil
}

func TargetToID(target string) int {
	var id int
	switch target {
	case "love":
		id = 1
	case "friends":
		id = 2
	case "communication":
		id = 3
	}
	return id
}

func IDToTarget(id int) string {
	var target string
	switch id {
	case 1:
		target = "love"
	case 2:
		target = "friends"
	case 3:
		target = "communication"
	}
	return target
}
