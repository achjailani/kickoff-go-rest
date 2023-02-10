package route

import (
	"github.com/achjailani/kickoff-go-rest/config"
	"github.com/achjailani/kickoff-go-rest/infrastructure/dao"
	"github.com/achjailani/kickoff-go-rest/pkg/security/jwt"
	"github.com/redis/go-redis/v9"
)

// WithConfig is function
func WithConfig(config *config.Config) RouterOption {
	return func(r *Router) {
		r.config = config
	}
}

// WithRedis is function
func WithRedis(redis *redis.Client) RouterOption {
	return func(r *Router) {
		r.redis = redis
	}
}

// WithDatabaseService is a function
func WithDatabaseService(databaseService *dao.Repositories) RouterOption {
	return func(r *Router) {
		r.repo = databaseService
	}
}

// WithJWT is a function to define JWT
func WithJWT(jwt *jwt.JWT) RouterOption {
	return func(r *Router) {
		r.jwt = jwt
	}
}
