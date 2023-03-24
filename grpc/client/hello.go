package client

import (
	"context"

	"github.com/AdiKhoironHasan/bookservices-api-gateway/proto/hello"
)

// Ping is a method
func (r GRPCClient) BooksPing(ctx context.Context) (*hello.PingReply, error) {
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	data, err := r.booksHello.Ping(ctx, &hello.PingRequest{})
	if err != nil {
		return nil, err
	}

	return data, nil
}
