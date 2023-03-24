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
	serverHost string
	serverPort int
	DSN        string
	addr       *string
)

// NewGRPCConn is a constructor
func NewGRPCConn(cfg *config.Config) (*grpc.ClientConn, error) {
	serverHost = "localhost"
	serverPort = cfg.GRPCPort
	DSN = fmt.Sprintf("%s:%d", serverHost, serverPort)
	addr = flag.String("addr", DSN, "The address to connect books service")

	flag.Parse()

	conn, err := grpc.Dial(*addr,
		grpc.WithTransportCredentials(insecure.NewCredentials()),
		grpc.WithChainUnaryInterceptor(interceptor.UnaryAuthClientInterceptor()),
		grpc.WithStreamInterceptor(interceptor.StreamAuthClientInterceptor()),
	)
	if err != nil {
		return nil, err
	}

	return conn, nil
}

func NewGRPCConn_Books(cfg *config.Config) (*grpc.ClientConn, error) {
	serverHost = cfg.Dependency.BookServices.AppHostGRPC
	serverPort = cfg.Dependency.BookServices.AppPortGRPC
	DSN = fmt.Sprintf("%s:%d", serverHost, serverPort)
	addr = flag.String("addr", DSN, "The address to connect books service")

	flag.Parse()

	conn, err := grpc.Dial(*addr,
		grpc.WithTransportCredentials(insecure.NewCredentials()),
		grpc.WithChainUnaryInterceptor(interceptor.UnaryAuthClientInterceptor()),
		grpc.WithStreamInterceptor(interceptor.StreamAuthClientInterceptor()),
	)
	if err != nil {
		return nil, err
	}

	return conn, nil
}
