package repository

import (
	"context"
	"github.com/achjailani/kickoff-go-rest/domain/entity"
)

// UserRepository is an interface of user behaviours
type UserRepository interface {
	Save(ctx context.Context, user *entity.User) error
	Update(ctx context.Context, user *entity.User, id string) error
	FindByID(ctx context.Context, id string) (*entity.User, error)
	FindByUsername(ctx context.Context, username string) (*entity.User, error)
	FindAll(ctx context.Context) ([]*entity.User, error)
}
