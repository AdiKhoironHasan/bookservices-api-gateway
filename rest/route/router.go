package route

import (
	"github.com/AdiKhoironHasan/bookservices-api-gateway/config"
	"github.com/AdiKhoironHasan/bookservices-api-gateway/domain/service"
	"github.com/AdiKhoironHasan/bookservices-api-gateway/grpc/client"
	"github.com/AdiKhoironHasan/bookservices-api-gateway/rest/handler"
	"github.com/AdiKhoironHasan/bookservices-api-gateway/rest/middleware"
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

	bookHandler := handler.NewBookHandler(h)

	// Books Services
	e.GET("/api/v1/ping/books", bookHandler.Ping)
	e.GET("/api/v1/books", bookHandler.List)
	e.POST("/api/v1/books", bookHandler.Store)
	e.GET("/api/v1/books/:id", bookHandler.Detail)
	e.PUT("/api/v1/books/:id", bookHandler.Update)
	e.DELETE("/api/v1/books/:id", bookHandler.Delete)

	return e
}
