package middleware

import (
	"fmt"
	"github.com/achjailani/kickoff-go-rest/config"
	"github.com/achjailani/kickoff-go-rest/domain/entity"
	"github.com/achjailani/kickoff-go-rest/infrastructure/dao"
	"github.com/achjailani/kickoff-go-rest/pkg/constant"
	"github.com/achjailani/kickoff-go-rest/pkg/exception"
	"github.com/achjailani/kickoff-go-rest/pkg/response"
	"github.com/achjailani/kickoff-go-rest/pkg/security/jwt"
	"github.com/gin-gonic/gin"
	"github.com/redis/go-redis/v9"
	"net/http"
	"strings"
)

const (
	// authTypeBearer is constant
	authTypeBearer = "Bearer"
)

// Guard is struct
type Guard struct {
	config *config.Config
	redis  *redis.Client
	jwt    *jwt.JWT
	repo   *dao.Repositories
}

// NewGuard is constructor
func NewGuard(config *config.Config, redis *redis.Client, jwt *jwt.JWT, repo *dao.Repositories) *Guard {
	return &Guard{
		config: config,
		redis:  redis,
		jwt:    jwt,
		repo:   repo,
	}
}

// Authenticate is middleware function for authorization
func (g *Guard) Authenticate() gin.HandlerFunc {
	return func(c *gin.Context) {
		headerAuth := c.Request.Header.Get("Authorization")
		if headerAuth == "" {
			c.Status(http.StatusUnauthorized)
			response.NewHTTPError(c, fmt.Sprintf("missing authorization token")).JSON()
			c.Abort()
			return
		}

		token := strings.Split(headerAuth, " ")
		headerAuthType := token[0]

		if headerAuthType != authTypeBearer || len(token) != 2 {
			c.Status(http.StatusUnauthorized)
			response.NewHTTPError(c, fmt.Sprintf("invalid authorization type")).JSON()
			c.Abort()
			return
		}

		tokenAuth := token[1]

		payload, errV := g.jwt.Verify(tokenAuth)
		if errV != nil {
			var er error
			switch errV {
			case jwt.ErrInvalidToken:
				er = exception.ErrInvalidToken
			case jwt.ErrExpiredToken:
				er = exception.ErrExpiredToken
			default:
				er = errV
			}

			c.Status(http.StatusUnauthorized)
			response.NewHTTPError(c, er.Error()).JSON()
			c.Abort()
			return
		}

		userID := payload.Identifier
		user, errF := g.repo.UserRepository.FindByID(c, userID)
		if errF != nil {
			c.Status(http.StatusUnauthorized)
			response.NewHTTPError(c, exception.ErrorUnauthorized.Error()).JSON()
			c.Abort()
			return
		}

		c.Set("user", user)
		c.Next()
	}
}

// OnlyAgent is function
func (g *Guard) OnlyAgent() gin.HandlerFunc {
	return func(c *gin.Context) {
		userReq, ok := c.Get("user")
		if !ok {
			c.Status(http.StatusInternalServerError)
			response.NewHTTPError(c, exception.ErrorInternalServerError.Error()).JSON()
			c.Abort()
			return
		}

		user := userReq.(*entity.User)
		role := user.Role.Code

		if role != constant.RoleDefaultAgent {
			c.Status(http.StatusForbidden)
			response.NewHTTPError(c, exception.ErrorForbidden.Error()).JSON()
			c.Abort()
			return
		}

		c.Next()
	}
}
