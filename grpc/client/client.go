package client

import (
	"github.com/AdiKhoironHasan/bookservice-protobank/proto/book"
	protoBook "github.com/AdiKhoironHasan/bookservice-protobank/proto/book"
	"github.com/AdiKhoironHasan/bookservices-api-gateway/proto/hello"
	"google.golang.org/grpc"
)

// GRPCClient is a struct
type GRPCClient struct {
	booksHello hello.HelloClient
	book       protoBook.BookServiceClient
}

// NewGRPCClient is constructor
func NewGRPCClient(connBooks grpc.ClientConnInterface) *GRPCClient {
	return &GRPCClient{
		booksHello: hello.NewHelloClient(connBooks),
		book:       book.NewBookServiceClient(connBooks),
	}
}
