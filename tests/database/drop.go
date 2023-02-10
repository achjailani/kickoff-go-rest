package database

import (
	"context"
	"gorm.io/gorm"
	"github.com/achjailani/kickoff-go-rest/domain/registry/core"
	"github.com/achjailani/kickoff-go-rest/infrastructure/dao"
)

// Drop is struct to hold connection to DB
type Drop struct {
	db *gorm.DB
}

// NewDrop is constructor
func NewDrop(db *gorm.DB) *Drop {
	return &Drop{db: db}
}

// DropPostgresql is function drop all tables of postgresql
func (op *Drop) DropPostgresql(ctx context.Context) error {
	for _, table := range core.CollectTables() {
		ok := op.db.WithContext(ctx).Migrator().HasTable(table.Name)
		if ok {
			err := op.db.WithContext(ctx).Migrator().DropTable(table.Name)
			if err != nil {
				return err
			}
		}
	}

	return nil
}

// Reset is function drop all tables & recreate them
func (op *Drop) Reset(ctx context.Context) error {
	err := op.DropPostgresql(ctx)
	if err != nil {
		return err
	}

	err = dao.NewRegistry().AutoMigrate(op.db)
	if err != nil {
		return err
	}

	return nil
}
