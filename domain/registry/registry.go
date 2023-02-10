package registry

import (
	"fmt"
	"log"
	"os"

	"gorm.io/gorm"
)

type Entity struct {
	Entity interface{}
}

type Table struct {
	Name interface{}
}

type Registry struct {
	Entities []Entity
	Tables   []Table
}

type Interface interface {
	AutoMigrate(db *gorm.DB) error
	ResetDatabase(db *gorm.DB) error
}

var _ Interface = &Registry{}

func (r *Registry) AutoMigrate(db *gorm.DB) error {
	var err error

	for _, model := range r.Entities {
		err = db.AutoMigrate(model.Entity)
		if err != nil {
			log.Fatal(err)
		}
	}

	return err
}

func (r *Registry) ResetDatabase(db *gorm.DB) error {
	var err error

	if os.Getenv("APP_ENV") == "production" {
		return nil
	}

	for _, table := range r.Tables {
		errDrop := db.Exec(fmt.Sprintf("DROP TABLE IF EXISTS %s", table.Name))
		if errDrop != nil {
			log.Fatal(errDrop)
		}
	}

	for _, model := range r.Entities {
		err = db.AutoMigrate(model.Entity)
		if err != nil {
			log.Fatal(err)
		}
	}

	return err
}
