package repository

import (
	"context"
	"github.com/achjailani/kickoff-go-rest/domain/entity"
)

// RoleRepository is an interface of role behaviours
type RoleRepository interface {
	Save(ctx context.Context, role *entity.Role) error
	Update(ctx context.Context, role *entity.Role, id string) error
	Delete(ctx context.Context, id string) error
	FindByID(ctx context.Context, id string) (*entity.Role, error)
	FindByCode(ctx context.Context, code string) (*entity.Role, error)
	FindAll(ctx context.Context) ([]*entity.Role, error)
}
