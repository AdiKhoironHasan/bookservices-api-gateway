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
	userHandler := handler.NewUserHandler(h)

	// api version
	apiV1 := e.Group("api/v1")

	// Books Services
	apiV1.GET("/ping/books", bookHandler.Ping)
	apiV1.GET("/books", bookHandler.List)
	apiV1.POST("/books", bookHandler.Store)
	apiV1.GET("/books/:id", bookHandler.Detail)
	apiV1.PUT("/books/:id", bookHandler.Update)
	apiV1.DELETE("/books/:id", bookHandler.Delete)

	// Users Services
	apiV1.GET("/ping/users", userHandler.Ping)
	apiV1.GET("/users", userHandler.List)
	apiV1.POST("/users", userHandler.Store)
	apiV1.GET("/users/:id", userHandler.Detail)
	apiV1.PUT("/users/:id", userHandler.Update)
	apiV1.DELETE("/users/:id", userHandler.Delete)

	return e
}
