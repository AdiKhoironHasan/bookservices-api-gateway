package client

import (
	"github.com/AdiKhoironHasan/bookservices/proto/book"
	"github.com/AdiKhoironHasan/bookservices/proto/hello"
	"google.golang.org/grpc"
)

// GRPCClient is a struct
type GRPCClient struct {
	booksHello hello.HelloClient
	book  book.BookServiceClient
}

// NewGRPCClient is constructor
func NewGRPCClient(connBooks grpc.ClientConnInterface) *GRPCClient {
	return &GRPCClient{
		booksHello: hello.NewHelloClient(connBooks),
		book:  book.NewBookServiceClient(connBooks),
	}
}
