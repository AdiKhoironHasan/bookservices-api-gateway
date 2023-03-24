package route

import (
	"github.com/AdiKhoironHasan/bookservices/config"
	"github.com/AdiKhoironHasan/bookservices/domain/service"
	"github.com/AdiKhoironHasan/bookservices/grpc/client"
	"github.com/AdiKhoironHasan/bookservices/rest/handler"
	"github.com/AdiKhoironHasan/bookservices/rest/middleware"
	"github.com/gin-gonic/gin"
)

// WithConfig is function
func WithConfig(config *config.Config) RouterOption {
	return func(r *Router) {
		r.config = config
	}
}

// WithRepository is function
func WithRepository(repo *service.Repositories) RouterOption {
	return func(r *Router) {
		r.repo = repo
	}
}

// WithGRPCClient is function
func WithGRPCClient(gClient *client.GRPCClient) RouterOption {
	return func(r *Router) {
		r.client = gClient
	}
}

// Init is a function
func (r *Router) Init() *gin.Engine {
	gin.SetMode(gin.ReleaseMode)

	e := gin.Default()
	e.Use(middleware.Logger())

	h := handler.NewHandler(r.repo, r.client)

	helloHandler := handler.NewHelloHandler(h)
	bookHandler := handler.NewBookHandler(h)

	// Books Services
	e.GET("/api/v1/ping/books", helloHandler.Ping)
	e.GET("/api/v1/books", bookHandler.List)
	e.POST("/api/v1/books", bookHandler.Store)

	return e
}
