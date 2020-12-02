package user

import (
	proto "park_2020/2020_2_tmp_name/api/users/delivery/grpc/protobuf"

	"google.golang.org/grpc"
)

type UserClient struct {
	client proto.UserGRPCHandlerClient
}

func NewUserClient(conn *grpc.ClientConn) *UserClient {
	c := proto.NewUserGRPCHandlerClient(conn)
	return &UserClient{
		client: c,
	}
}
