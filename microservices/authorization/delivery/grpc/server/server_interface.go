package server

import (
	auth "park_2020/2020_2_tmp_name/microservices/authorization/delivery/grpc/protobuf"
)

type UserUsecase interface {
	Login(data auth.LoginData) (string, error)
	Logout(session string) error
	CheckSession(cookie string) (auth.User, error)
}
}
