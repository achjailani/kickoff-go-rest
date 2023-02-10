package seeder

import (
	"context"
	"fmt"
	"gorm.io/gorm"
	"github.com/achjailani/kickoff-go-rest/domain/entity"
	"github.com/achjailani/kickoff-go-rest/infrastructure/dao"
	"github.com/achjailani/kickoff-go-rest/pkg/constant"
	"github.com/achjailani/kickoff-go-rest/pkg/security"
)

// UserSeeder is struct user seeder
type UserSeeder struct{}

// implement Seeder interface
var _ Seeder = &UserSeeder{}

// Seed is a method implementation
func (u *UserSeeder) Seed(ctx context.Context, repo *dao.Repositories) error {
	roleAdmin := &entity.Role{
		Name: constant.RoleDefaultAgent,
		Code: constant.RoleDefaultAgent,
	}

	roleUser := &entity.Role{
		Name: constant.RoleDefaultUser,
		Code: constant.RoleDefaultUser,
	}

	resultRoleAdmin, err := repo.RoleRepository.FindByCode(ctx, constant.RoleDefaultAgent)
	if err != nil {
		switch err {
		case gorm.ErrRecordNotFound:
			errAdmin := repo.RoleRepository.Save(ctx, roleAdmin)
			if errAdmin != nil {
				return fmt.Errorf("seed role admin: %w", errAdmin)
			}
		default:
			return fmt.Errorf("find admin: %w", err)
		}
	} else {
		roleAdmin = resultRoleAdmin
	}

	resultRoleUser, err := repo.RoleRepository.FindByCode(ctx, constant.RoleDefaultUser)
	if err != nil {
		switch err {
		case gorm.ErrRecordNotFound:
			errAdmin := repo.RoleRepository.Save(ctx, roleUser)
			if errAdmin != nil {
				return fmt.Errorf("seed role admin: %w", err)
			}
		default:
			return fmt.Errorf("find admin: %w", err)
		}
	} else {
		roleUser = resultRoleUser
	}

	_, err = repo.UserRepository.FindByUsername(ctx, constant.UserDefaultUsername)
	if err != nil {
		switch err {
		case gorm.ErrRecordNotFound:
			err = repo.UserRepository.Save(ctx, &entity.User{
				RoleID:   roleAdmin.ID,
				Name:     constant.UserDefaultName,
				Username: constant.UserDefaultUsername,
				Password: security.HashMake(constant.UserDefaultPassword),
				UserProfile: entity.UserProfile{
					Email: constant.UserDefaultEmail,
				}})

			if err != nil {
				return fmt.Errorf("seed user: %w", err)
			}
		}
	}

	return nil
}
