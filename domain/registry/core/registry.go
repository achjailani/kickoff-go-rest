package core

import (
	"github.com/achjailani/kickoff-go-rest/domain/entity"
	"github.com/achjailani/kickoff-go-rest/domain/registry"
)

// CollectEntities is function collects entities
func CollectEntities() []registry.Entity {
	return []registry.Entity{
		{Entity: entity.Role{}},
		{Entity: entity.User{}},
		{Entity: entity.UserProfile{}},
	}
}

// CollectTables is function collects entity names
func CollectTables() []registry.Table {
	var user entity.User
	var role entity.Role
	var userProfile entity.UserProfile

	return []registry.Table{
		{Name: role.TableName()},
		{Name: user.TableName()},
		{Name: userProfile.TableName()},
	}
}
