package config_test

import (
	"reflect"
	"testing"

	"github.com/achjailani/kickoff-go-rest/config"

	"github.com/stretchr/testify/assert"
)

func TestDBConfig(t *testing.T) {
	metaDBFields := []string{
		"DBDriver",
		"DBHost",
		"DBPort",
		"DBUser",
		"DBPassword",
		"DBName",
		"DBTimeZone",
		"DBLog",
	}

	t.Run("if valid DBConfig", func(t *testing.T) {
		metaValue := reflect.ValueOf(new(config.DBConfig)).Elem()

		for _, field := range metaDBFields {
			assert.False(t, (metaValue.FieldByName(field) == (reflect.Value{})))
		}
	})

	t.Run("if valid DBTestConfig", func(t *testing.T) {
		metaValue := reflect.ValueOf(new(config.DBTestConfig)).Elem()

		for _, field := range metaDBFields {
			assert.False(t, (metaValue.FieldByName(field) == (reflect.Value{})))
		}
	})
}
