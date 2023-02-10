package route

import (
	"github.com/achjailani/kickoff-go-rest/config"
	"github.com/achjailani/kickoff-go-rest/infrastructure/dao"
	"github.com/achjailani/kickoff-go-rest/interface/handler"
	"github.com/achjailani/kickoff-go-rest/interface/middleware"
	"github.com/achjailani/kickoff-go-rest/pkg/security/jwt"
	"github.com/redis/go-redis/v9"

	"github.com/gin-gonic/gin"
)

// Router is a struct contains dependencies needed
type Router struct {
	config          *config.Config
	redis           *redis.Client
	jwt             *jwt.JWT
	databaseService *dao.Repositories
}

// RouterOption return Router with RouterOption to fill up the dependencies
type RouterOption func(*Router)

// NewRouter is a constructor will initialize Router.
func NewRouter(options ...RouterOption) *Router {
	router := &Router{}

	for _, opt := range options {
		opt(router)
	}

	return router
}

// Init is a function
func (r *Router) Init() *gin.Engine {
	gin.SetMode(gin.ReleaseMode)

	e := gin.Default()

	guard := middleware.NewGuard(r.config, r.redis, r.jwt, r.databaseService)

	useCase := handler.NewHandler(r.config, r.redis, r.databaseService, r.jwt)

	e.GET("/api/ping", useCase.Ping)
	e.GET("/api/dev-error", useCase.DevErrorHandler)
	e.GET("/api/dev-success", useCase.DevSuccessHandler)
	e.GET("/api/dev-async", useCase.DevAsyncHandler)
	e.GET("/api/dev-sync", useCase.DevSyncHandler)

	e.POST("/api/login", useCase.Login)

	authenticated := e.Group("/api/v1", guard.Authenticate())

	owner := authenticated.Group("/", guard.OnlyAgent())

	// Route role
	owner.POST("/roles", useCase.CreateRoleHandler)
	owner.GET("/roles", useCase.GetAllRoleHandler)
	owner.GET("/roles/:id", useCase.GetRoleHandler)

	// Route user
	owner.POST("/users", useCase.UserCreateHandler)
	owner.GET("/users/:id", useCase.UserGetByIDHandler)
	owner.GET("/users", useCase.UserListHandler)

	return e
}
