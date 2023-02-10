package config

import (
	"github.com/achjailani/kickoff-go-rest/util"
)

type DBConfig struct {
	DBDriver   string
	DBHost     string
	DBPort     string
	DBUser     string
	DBPassword string
	DBName     string
	DBTimeZone string
	DBLog      bool
}

type DBTestConfig struct {
	DBDriver   string
	DBHost     string
	DBPort     string
	DBUser     string
	DBPassword string
	DBName     string
	DBTimeZone string
	DBLog      bool
}

type RedisConfig struct {
	RedisHost     string
	RedisPort     string
	RedisPassword string
	RedisDB       int
}

type KeyConfig struct {
	AppPrivateKey string
	AppPublicKey  string
}

type RedisTestConfig struct {
	RedisHost     string
	RedisPort     string
	RedisPassword string
	RedisDB       int
}

type Config struct {
	KeyConfig
	DBConfig
	DBTestConfig
	RedisConfig
	RedisTestConfig
	DebugMode bool
	TestMode  bool
}

func New() *Config {
	return &Config{
		KeyConfig: KeyConfig{
			AppPrivateKey: util.GetEnv("APP_PRIVATE_KEY", "default-private-key"),
			AppPublicKey:  util.GetEnv("APP_PUBLIC_KEY", "default-public-key"),
		},
		DBConfig: DBConfig{
			DBDriver:   util.GetEnv("DB_DRIVER", "postgres"),
			DBHost:     util.GetEnv("DB_HOST", "127.0.0.1"),
			DBPort:     util.GetEnv("DB_PORT", "5432"),
			DBUser:     util.GetEnv("DB_USER", "postgres"),
			DBPassword: util.GetEnv("DB_PASSWORD", ""),
			DBName:     util.GetEnv("DB_NAME", "postgres"),
			DBTimeZone: util.GetEnv("APP_TIMEZONE", "Asia/Jakarta"),
			DBLog:      util.GetEnvAsBool("DB_LOG", false),
		},
		DBTestConfig: DBTestConfig{
			DBDriver:   util.GetEnv("DB_TEST_DRIVER", "postgres"),
			DBHost:     util.GetEnv("DB_TEST_HOST", "127.0.0.1"),
			DBPort:     util.GetEnv("DB_TEST_PORT", "5432"),
			DBUser:     util.GetEnv("DB_TEST_USER", "postgres"),
			DBPassword: util.GetEnv("DB_TEST_PASSWORD", ""),
			DBName:     util.GetEnv("DB_TEST_NAME", "postgres_test"),
			DBTimeZone: util.GetEnv("APP_TIMEZONE", "Asia/Jakarta"),
			DBLog:      util.GetEnvAsBool("DB_TEST_LOG", false),
		},
		RedisConfig: RedisConfig{
			RedisHost:     util.GetEnv("REDIS_HOST", "localhost"),
			RedisPort:     util.GetEnv("REDIS_PORT", "6379"),
			RedisPassword: util.GetEnv("REDIS_PASSWORD", ""),
			RedisDB:       util.GetEnvAsInt("REDIS_DB", 0),
		},
		RedisTestConfig: RedisTestConfig{
			RedisHost:     util.GetEnv("TEST_REDIS_HOST", "localhost"),
			RedisPort:     util.GetEnv("TEST_REDIS_HOST", "6379"),
			RedisPassword: util.GetEnv("TEST_REDIS_HOST", ""),
			RedisDB:       util.GetEnvAsInt("TEST_REDIS_HOST", 0),
		},
		DebugMode: util.GetEnv("APP_ENV", "local") != "production",
		TestMode:  util.GetEnvAsBool("TEST_MODE", false),
	}
}
