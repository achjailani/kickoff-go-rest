package connection

import (
	"github.com/achjailani/kickoff-go-rest/config"
	"log"
	"os"
	"time"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

const (
	driverPostgres = "postgres"
)

func NewDBConnection(conf *config.Config) (*gorm.DB, error) {
	newLogger := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags),
		logger.Config{
			SlowThreshold: time.Second,
			LogLevel:      logger.Info,
			Colorful:      true,
		},
	)

	gormConfig := &gorm.Config{
		DisableForeignKeyConstraintWhenMigrating: true,
	}

	if conf.DBConfig.DBLog {
		gormConfig.Logger = newLogger
	}

	var db *gorm.DB

	switch conf.DBConfig.DBDriver {
	case driverPostgres:
		var dns *PostgresDNS

		if conf.TestMode == true {
			dns = &PostgresDNS{
				Host:     conf.DBTestConfig.DBHost,
				Port:     conf.DBTestConfig.DBPort,
				User:     conf.DBTestConfig.DBUser,
				Password: conf.DBTestConfig.DBPassword,
				DBName:   conf.DBTestConfig.DBName,
				SSLMode:  false,
				Timezone: conf.DBTestConfig.DBTimeZone,
			}
		} else {
			dns = &PostgresDNS{
				Host:     conf.DBConfig.DBHost,
				Port:     conf.DBConfig.DBPort,
				User:     conf.DBConfig.DBUser,
				Password: conf.DBConfig.DBPassword,
				DBName:   conf.DBConfig.DBName,
				SSLMode:  false,
				Timezone: conf.DBConfig.DBTimeZone,
			}
		}

		dbConn, err := gorm.Open(postgres.Open(dns.ToString()), gormConfig)
		if err != nil {
			return nil, err
		}

		return dbConn, nil
	}

	return db, nil
}
