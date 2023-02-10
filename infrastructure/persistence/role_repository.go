package persistence

import (
	"context"
	"gorm.io/gorm"
	"github.com/achjailani/kickoff-go-rest/domain/entity"
	"github.com/achjailani/kickoff-go-rest/domain/repository"
)

// RoleRepository is a struct to store db connection.
type RoleRepository struct {
	db *gorm.DB
}

// NewRoleRepository is a function to build a struct role repository.
func NewRoleRepository(db *gorm.DB) *RoleRepository {
	return &RoleRepository{db: db}
}

// implement repository.RoleRepository
var _ repository.RoleRepository = &RoleRepository{}

// Save is function to store new Role.
func (o *RoleRepository) Save(ctx context.Context, role *entity.Role) error {
	return o.db.WithContext(ctx).Create(role).Error
}

// Update is function to update pre-existing Role By ID.
func (o *RoleRepository) Update(ctx context.Context, role *entity.Role, id string) error {
	return o.db.WithContext(ctx).Where("id = ?", id).Updates(role).Error
}

// Delete is function to remove existing Role By ID.
func (o *RoleRepository) Delete(ctx context.Context, id string) error {
	return o.db.WithContext(ctx).Delete(&entity.User{}, id).Error
}

// FindByID is function to retrieve Role data By ID.
func (o *RoleRepository) FindByID(ctx context.Context, id string) (*entity.Role, error) {
	var role entity.Role

	err := o.db.WithContext(ctx).Where("id = ?", id).Take(&role).Error
	if err != nil {
		return nil, err
	}

	return &role, nil
}

// FindByCode is function to retrieve Role data By Code.
func (o *RoleRepository) FindByCode(ctx context.Context, code string) (*entity.Role, error) {
	var role entity.Role

	err := o.db.WithContext(ctx).Where("code = ?", code).Take(&role).Error
	if err != nil {
		return nil, err
	}

	return &role, nil
}

func (o *RoleRepository) FindAll(ctx context.Context) ([]*entity.Role, error) {
	var roles []*entity.Role

	err := o.db.WithContext(ctx).Find(&roles).Error
	if err != nil {
		return nil, err
	}

	return roles, nil
}