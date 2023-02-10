package persistence

import (
	"context"
	"gorm.io/gorm"
	"github.com/achjailani/kickoff-go-rest/domain/entity"
	"github.com/achjailani/kickoff-go-rest/domain/repository"
)

// UserProfileRepository is a struct to store db connection.
type UserProfileRepository struct {
	db *gorm.DB
}

// NewUserProfileRepository is a function to build a struct user profile repository.
func NewUserProfileRepository(db *gorm.DB) *UserProfileRepository {
	return &UserProfileRepository{db: db}
}

// implement repository.UserProfileRepository
var _ repository.UserProfileRepository = &UserProfileRepository{}

// Save is function to store new user profile.
func (u *UserProfileRepository) Save(ctx context.Context, userProfile *entity.UserProfile) error {
	return u.db.WithContext(ctx).Create(userProfile).Error
}

// Update is function to update pre-existing user profile By user ID.
func (u *UserProfileRepository) Update(ctx context.Context, userProfile *entity.UserProfile, userId string) error {
	return u.db.WithContext(ctx).Where("user_id = ?", userId).Updates(userProfile).Error
}

// Delete is function to update existing user profile By user ID.
func (u *UserProfileRepository) Delete(ctx context.Context, userId string) error {
	return u.db.WithContext(ctx).Where("user_id = ?", userId).Delete(&entity.UserProfile{}).Error
}
