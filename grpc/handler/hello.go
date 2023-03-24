package handler

// import (
// 	"context"
// 	"fmt"

// 	"github.com/AdiKhoironHasan/bookservices-api-gateway/proto/hello"
// )

// // Ping is a function
// func (c *Handler) Ping(_ context.Context, _ *hello.PingRequest) (*hello.PingReply, error) {
// 	now := []string{
// 		"now",
// 	}

// 	err := c.repo.DB.Raw("SELECT now()").Scan(&now).Error
// 	if err != nil {
// 		return nil, err
// 	}

// 	return &hello.PingReply{
// 		Redis: "Ok",
// 		Db:    fmt.Sprintf("Ok : %v", now),
// 	}, nil
// }
