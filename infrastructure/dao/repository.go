package dao

import (
	"github.com/achjailani/kickoff-go-rest/domain/repository"
	"github.com/achjailani/kickoff-go-rest/infrastructure/persistence"

	"gorm.io/gorm"
)

// Repositories is a struct collects repositories
type Repositories struct {
	DB                    *gorm.DB
	UserRepository        repository.UserRepository
	RoleRepository        repository.RoleRepository
	UserProfileRepository repository.UserProfileRepository
}

// NewRepo is constructor of Repositories
func NewRepo(db *gorm.DB) *Repositories {
	return &Repositories{
		DB:                    db,
		UserRepository:        persistence.NewUserRepository(db),
		RoleRepository:        persistence.NewRoleRepository(db),
		UserProfileRepository: persistence.NewUserProfileRepository(db),
	}
}
