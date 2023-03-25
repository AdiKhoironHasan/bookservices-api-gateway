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

	auth := apiV1.Group("").Use(middleware.Authorize())

	// Books Services
	auth.GET("/ping/books", bookHandler.Ping)
	auth.GET("/books", bookHandler.List)
	auth.POST("/books", bookHandler.Store)
	auth.GET("/books/:id", bookHandler.Detail)
	auth.PUT("/books/:id", bookHandler.Update)
	auth.DELETE("/books/:id", bookHandler.Delete)

	// Users Services
	auth.GET("/ping/users", userHandler.Ping)
	auth.GET("/users", userHandler.List)
	auth.POST("/users", userHandler.Store)
	auth.GET("/users/:id", userHandler.Detail)
	auth.PUT("/users/:id", userHandler.Update)
	auth.DELETE("/users/:id", userHandler.Delete)

	// Auth Service
	apiV1.POST("/login", userHandler.Login)
	apiV1.POST("/register", userHandler.Store)
	// auth.POST("/verify", userHandler.Verify)
	// apiV1.POST("/logout", userHandler.Logout)

	return e
}
