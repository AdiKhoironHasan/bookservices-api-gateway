package handler

import (
	"github.com/AdiKhoironHasan/bookservices/config"
	"github.com/AdiKhoironHasan/bookservices/domain/service"
	"github.com/AdiKhoironHasan/bookservices/proto/book"
	"github.com/AdiKhoironHasan/bookservices/proto/hello"
	// "github.com/AdiKhoironHasan/bookservices/proto/foo"
)

// Interface is an interface
type Interface interface {
	// interface of grpc handler
	hello.HelloServer
	book.BookServiceServer
}

// Handler is struct
type Handler struct {
	config *config.Config
	repo   *service.Repositories

	hello.UnimplementedHelloServer
	book.UnimplementedBookServiceServer
}

// NewHandler is a constructor
func NewHandler(conf *config.Config, repo *service.Repositories) *Handler {
	return &Handler{
		config: conf,
		repo:   repo,
	}
}

var _ Interface = &Handler{}
