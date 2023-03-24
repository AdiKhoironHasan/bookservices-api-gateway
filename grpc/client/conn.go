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
	bookServerHost string
	bookServerPort string
	bookDSN        string
	bookAddress    *string

	userServerHost string
	userServerPort string
	userDSN        string
	userAddress    *string
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
	var (
		bookServerHost = cfg.Dependency.BookServices.AppHostGRPC
		bookServerPort = cfg.Dependency.BookServices.AppPortGRPC
		bookDSN        = fmt.Sprintf("%s:%d", bookServerHost, bookServerPort)
		bookAddress    = flag.String("addr", bookDSN, "The address to connect books service")
	)

	flag.Parse()

	conn, err := grpc.Dial(*bookAddress,
		grpc.WithTransportCredentials(insecure.NewCredentials()),
		grpc.WithChainUnaryInterceptor(interceptor.UnaryAuthClientInterceptor()),
		grpc.WithStreamInterceptor(interceptor.StreamAuthClientInterceptor()),
	)
	if err != nil {
		return nil, err
	}

	return conn, nil
}

func NewGRPCConn_User(cfg *config.Config) (*grpc.ClientConn, error) {
	var (
		userServerHost = cfg.Dependency.UserServices.AppHostGRPC
		userServerPort = cfg.Dependency.UserServices.AppPortGRPC
		userDSN        = fmt.Sprintf("%s:%d", userServerHost, userServerPort)
		userAddress    = flag.String("addr", userDSN, "The address to connect user service")
	)

	flag.Parse()

	conn, err := grpc.Dial(*userAddress,
		grpc.WithTransportCredentials(insecure.NewCredentials()),
		grpc.WithChainUnaryInterceptor(interceptor.UnaryAuthClientInterceptor()),
		grpc.WithStreamInterceptor(interceptor.StreamAuthClientInterceptor()),
	)
	if err != nil {
		return nil, err
	}

	return conn, nil
}
