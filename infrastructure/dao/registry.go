package dao

import (
	"github.com/achjailani/kickoff-go-rest/domain/registry"
	"github.com/achjailani/kickoff-go-rest/domain/registry/core"
)

func NewRegistry() *registry.Registry {
	var entityRegistry []registry.Entity
	var tableRegistry []registry.Table

	entityRegistry = append(entityRegistry, core.CollectEntities()...)
	tableRegistry = append(tableRegistry, core.CollectTables()...)

	return &registry.Registry{
		Entities: entityRegistry,
		Tables:   tableRegistry,
	}
}
