package connection

import (
	"fmt"
	"github.com/redis/go-redis/v9"
	"github.com/achjailani/kickoff-go-rest/config"
)

// NewRedisConnection is constructor
func NewRedisConnection(conf *config.Config) *redis.Client {
	var option *redis.Options

	if conf.TestMode {
		option = &redis.Options{
			Addr:     fmt.Sprintf("%s:%s", conf.RedisTestConfig.RedisHost, conf.RedisTestConfig.RedisPort),
			Password: conf.RedisTestConfig.RedisPassword,
			DB:       conf.RedisTestConfig.RedisDB,
		}
	} else {
		option = &redis.Options{
			Addr:     fmt.Sprintf("%s:%s", conf.RedisConfig.RedisHost, conf.RedisConfig.RedisPort),
			Password: conf.RedisConfig.RedisPassword,
			DB:       conf.RedisConfig.RedisDB,
		}
	}

	rdb := redis.NewClient(option)

	return rdb
}
