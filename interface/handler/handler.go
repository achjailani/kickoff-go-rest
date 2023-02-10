package handler

import (
	"github.com/achjailani/kickoff-go-rest/config"
	"github.com/achjailani/kickoff-go-rest/infrastructure/dao"
	"github.com/achjailani/kickoff-go-rest/pkg/security/jwt"
	"github.com/redis/go-redis/v9"
)

// Handler is struct
type Handler struct {
	config     *config.Config
	redis      *redis.Client
	jwt        *jwt.JWT
	repository *dao.Repositories
}

// NewHandler is constructor
func NewHandler(config *config.Config, redis *redis.Client, repo *dao.Repositories, jwt *jwt.JWT) *Handler {
	return &Handler{
		config:     config,
		redis:      redis,
		repository: repo,
		jwt:        jwt,
	}
}
