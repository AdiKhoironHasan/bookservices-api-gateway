package client

import (
	"github.com/AdiKhoironHasan/bookservice-protobank/proto/book"
	protoBook "github.com/AdiKhoironHasan/bookservice-protobank/proto/book"
	"google.golang.org/grpc"
)

// GRPCClient is a struct
type GRPCClient struct {
	book protoBook.BookServiceClient
}

// NewGRPCClient is constructor
func NewGRPCClient(connBooks grpc.ClientConnInterface) *GRPCClient {
	return &GRPCClient{
		book: book.NewBookServiceClient(connBooks),
	}
}
