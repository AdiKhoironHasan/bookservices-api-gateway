package handler

import (
	"github.com/AdiKhoironHasan/bookservices-api-gateway/config"
	"github.com/AdiKhoironHasan/bookservices-api-gateway/domain/service"
	"github.com/AdiKhoironHasan/bookservices-api-gateway/proto/hello"

	protoBook "github.com/AdiKhoironHasan/bookservice-protobank/proto/book"
)

// Interface is an interface
type Interface interface {
	// interface of grpc handler
	hello.HelloServer
	protoBook.BookServiceServer
}

// Handler is struct
type Handler struct {
	config *config.Config
	repo   *service.Repositories

	hello.UnimplementedHelloServer
	protoBook.UnimplementedBookServiceServer
}

// NewHandler is a constructor
func NewHandler(conf *config.Config, repo *service.Repositories) *Handler {
	return &Handler{
		config: conf,
		repo:   repo,
	}
}

var _ Interface = &Handler{}
