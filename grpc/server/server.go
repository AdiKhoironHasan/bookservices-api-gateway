package server

import (
	"github.com/AdiKhoironHasan/bookservices-api-gateway/config"
	"github.com/AdiKhoironHasan/bookservices-api-gateway/domain/service"
	"github.com/AdiKhoironHasan/bookservices-api-gateway/grpc/handler"
	"github.com/AdiKhoironHasan/bookservices-api-gateway/grpc/interceptor"
	"github.com/AdiKhoironHasan/bookservices-api-gateway/proto/hello"

	protoBook "github.com/AdiKhoironHasan/bookservice-protobank/proto/book"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

// Server is struct to hold any dependencies used for server
type Server struct {
	config *config.Config
	repo   *service.Repositories
}

// NewGRPCServer is constructor
func NewGRPCServer(conf *config.Config, repo *service.Repositories) *Server {
	return &Server{
		config: conf,
		repo:   repo,
	}
}

// Run is a method gRPC server
func (s *Server) Run(port int) error {
	//server := grpc.NewServer(grpc.UnaryInterceptor(interceptor.AuthorizationInterceptor))
	server := grpc.NewServer(
		grpc.ChainUnaryInterceptor(
			interceptor.UnaryLoggerServerInterceptor(),
			interceptor.UnaryAuthServerInterceptor(),
		),
		grpc.ChainStreamInterceptor(
			interceptor.StreamLoggerServerInterceptor(),
			interceptor.StreamAuthServerInterceptor(),
		),
	)

	handlers := handler.NewHandler(s.config, s.repo)

	// register from proto
	hello.RegisterHelloServer(server, handlers)
	protoBook.RegisterBookServiceServer(server, handlers)

	// register reflection
	reflection.Register(server)

	return RunGRPCServer(server, port)
}
