package models

import (
	"errors"
	"net/http"

	"github.com/sirupsen/logrus"
)

var (
	ErrBadRequest          = errors.New("Bad request")
	ErrNotFound            = errors.New("Your requested Item is not found")
	ErrConflict            = errors.New("Your Item already exist")
	ErrUnauthorized        = errors.New("User not authorised or not found")
	ErrInternalServerError = errors.New("Internal Server Error")
)

func GetStatusCode(err error) int {
	if err == nil {
		return http.StatusOK
	}

	logrus.Error(err)
	switch err {
	case ErrBadRequest: // 400
		return http.StatusBadRequest
	case ErrNotFound:
		return http.StatusNotFound // 404
	case ErrConflict:
		return http.StatusConflict // 409
	case ErrUnauthorized:
		return http.StatusUnauthorized // 401
	default:
		return http.StatusInternalServerError // 500
	}
}
