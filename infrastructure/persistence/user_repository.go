package persistence

import (
	"context"
	"github.com/achjailani/kickoff-go-rest/domain/entity"
	"github.com/achjailani/kickoff-go-rest/domain/repository"

	"gorm.io/gorm"
)

// UserRepository is a struct to store db connection.
type UserRepository struct {
	db *gorm.DB
}

// NewUserRepository is a function to build a struct user repository.
func NewUserRepository(db *gorm.DB) *UserRepository {
	return &UserRepository{db: db}
}

// implement repository.UserRepository
var _ repository.UserRepository = &UserRepository{}

// Save is function to store new User.
func (u *UserRepository) Save(ctx context.Context, user *entity.User) error {
	return u.db.WithContext(ctx).Create(user).Error
}

// Update is function to update pre-existing User By ID.
func (u *UserRepository) Update(ctx context.Context, user *entity.User, id string) error {
	return u.db.WithContext(ctx).Where("id = ?", id).Updates(user).Error
}

// FindByID is function to retrieve User data By ID.
func (u *UserRepository) FindByID(ctx context.Context, id string) (*entity.User, error) {
	var user entity.User

	err := u.db.WithContext(ctx).
		Joins("Role").
		Joins("UserProfile").
		Where("users.id = ?", id).
		Take(&user).Error
	if err != nil {
		return nil, err
	}

	return &user, nil
}

// FindByUsername is function to retrieve User data By Username.
func (u *UserRepository) FindByUsername(ctx context.Context, username string) (*entity.User, error) {
	var user entity.User

	err := u.db.WithContext(ctx).
		Joins("Role").
		Joins("UserProfile").
		Where("username = ?", username).Take(&user).Error
	if err != nil {
		return nil, err
	}

	return &user, nil
}

// FindAll is function to retrieve all data
func (u *UserRepository) FindAll(ctx context.Context) ([]*entity.User, error) {
	var users []*entity.User

	err := u.db.WithContext(ctx).Find(&users).Error
	if err != nil {
		return nil, err
	}

	return users, nil
}
