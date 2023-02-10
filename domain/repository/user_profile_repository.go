package repository

import (
	"context"
	"github.com/achjailani/kickoff-go-rest/domain/entity"
)

// UserProfileRepository is an interface of user profile behaviours
type UserProfileRepository interface {
	Save(ctx context.Context, userProfile *entity.UserProfile) error
	Update(ctx context.Context, userProfile *entity.UserProfile, userId string) error
	Delete(ctx context.Context, userId string) error
}
