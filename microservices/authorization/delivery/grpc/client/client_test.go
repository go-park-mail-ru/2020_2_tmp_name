package client

import (
	"testing"

	"google.golang.org/grpc"
)

func TestNewAuthClient(t *testing.T) {
	conn := &grpc.ClientConn{}
	NewAuthClient(conn)
}
