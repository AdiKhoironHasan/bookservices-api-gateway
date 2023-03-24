package client

import (
	"google.golang.org/grpc"

	protoBook "github.com/AdiKhoironHasan/bookservice-protobank/proto/book"
	protoUser "github.com/AdiKhoironHasan/bookservice-protobank/proto/user"
)

// GRPCClient is a struct
type GRPCClient struct {
	book protoBook.BookServiceClient
	user protoUser.UserServiceClient
}

// NewGRPCClient is constructor
func NewGRPCClient(
	connBook grpc.ClientConnInterface,
	connUser grpc.ClientConnInterface,
) *GRPCClient {
	return &GRPCClient{
		book: protoBook.NewBookServiceClient(connBook),
		user: protoUser.NewUserServiceClient(connUser),
	}
}
