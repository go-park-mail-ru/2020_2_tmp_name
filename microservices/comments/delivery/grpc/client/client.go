package client

import (
	"context"
	"github.com/sirupsen/logrus"
	"google.golang.org/genproto/googleapis/datastore/admin/v1"
	"google.golang.org/grpc"
	"google.golang.org/grpc/keepalive"
	"google.golang.org/grpc/reflection"
	"net"
	comments "park_2020/2020_2_tmp_name/microservices/comments"
	proto "park_2020/2020_2_tmp_name/microservices/comments/delivery/grpc/protobuf"
	"time"
)

type CommentClient struct {
	client proto.CommentsGRPCHandlerClient
}

func NewCommentsClientGRPC(conn *grpc.ClientConn) *CommentClient{
	c := proto.NewCommentsGRPCHandlerClient(conn)
	return &CommentClient{
		client: c,
	}
}

func (c *CommentClient) Comment(ctx context.Context, userComment *proto.UserComment) (*proto.Empty, error) {

}

func (c *CommentClient) CommentsById(ctx context.Context, id *proto.Id) (*proto.CommentsData, error) {

}