package client

import (
	"flag"
	"fmt"

	"github.com/AdiKhoironHasan/bookservices-api-gateway/config"
	"github.com/AdiKhoironHasan/bookservices-api-gateway/grpc/interceptor"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

var (
	ServerHost string
	ServerPort int
	DSN        string
	Address    *string
)

// NewGRPCConn is a constructor
// func NewGRPCConn(cfg *config.Config) (*grpc.ClientConn, error) {
// 	serverHost = "localhost"
// 	serverPort = cfg.GRPCPort
// 	DSN = fmt.Sprintf("%s:%d", serverHost, serverPort)
// 	addr = flag.String("addr", DSN, "The address to connect books service")

// 	flag.Parse()

// 	conn, err := grpc.Dial(*addr,
// 		grpc.WithTransportCredentials(insecure.NewCredentials()),
// 		grpc.WithChainUnaryInterceptor(interceptor.UnaryAuthClientInterceptor()),
// 		grpc.WithStreamInterceptor(interceptor.StreamAuthClientInterceptor()),
// 	)
// 	if err != nil {
// 		return nil, err
// 	}

// 	return conn, nil
// }

func NewGRPCConn_Book(cfg *config.Config) (*grpc.ClientConn, error) {
	ServerHost = cfg.Dependency.BookServices.AppHostGRPC
	ServerPort = cfg.Dependency.BookServices.AppPortGRPC
	DSN = fmt.Sprintf("%s:%d", ServerHost, ServerPort)
	Address = flag.String("bookAddress", DSN, "The address to connect books service")

	flag.Parse()

	conn, err := grpc.Dial(*Address,
		grpc.WithTransportCredentials(insecure.NewCredentials()),
		grpc.WithChainUnaryInterceptor(
			interceptor.UnaryAuthClientInterceptor(),
			interceptor.UnaryConnClientInterceptor(),
		),
		grpc.WithStreamInterceptor(interceptor.StreamAuthClientInterceptor()),
	)
	if err != nil {
		return nil, err
	}

	return conn, nil
}

func NewGRPCConn_User(cfg *config.Config) (*grpc.ClientConn, error) {
	ServerHost = cfg.Dependency.UserServices.AppHostGRPC
	ServerPort = cfg.Dependency.UserServices.AppPortGRPC
	DSN = fmt.Sprintf("%s:%d", ServerHost, ServerPort)
	Address = flag.String("userAddress", DSN, "The address to connect user service")

	flag.Parse()

	conn, err := grpc.Dial(*Address,
		grpc.WithTransportCredentials(insecure.NewCredentials()),
		grpc.WithChainUnaryInterceptor(
			interceptor.UnaryAuthClientInterceptor(),
			interceptor.UnaryConnClientInterceptor(),
		),
		grpc.WithStreamInterceptor(interceptor.StreamAuthClientInterceptor()),
	)
	if err != nil {
		return nil, err
	}

	return conn, nil
}
