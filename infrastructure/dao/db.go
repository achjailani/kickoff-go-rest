package dao

import (
	"github.com/achjailani/kickoff-go-rest/domain/repository"
	"github.com/achjailani/kickoff-go-rest/infrastructure/persistence"

	"gorm.io/gorm"
)

type Repositories struct {
	DB                    *gorm.DB
	UserRepository        repository.UserRepository
	RoleRepository        repository.RoleRepository
	UserProfileRepository repository.UserProfileRepository
}

func NewDBService(db *gorm.DB) *Repositories {
	return &Repositories{
		DB:                    db,
		UserRepository:        persistence.NewUserRepository(db),
		RoleRepository:        persistence.NewRoleRepository(db),
		UserProfileRepository: persistence.NewUserProfileRepository(db),
	}
}
