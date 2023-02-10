package util

import (
	"os"
	"strconv"
	"strings"
)

func GetEnv(key string, defaultVal string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}

	if nextValue := os.Getenv(key); nextValue != "" {
		return nextValue
	}

	return defaultVal
}

func GetEnvAsInt(key string, defaultVal int) int {
	varStr := GetEnv(key, "")
	if val, err := strconv.Atoi(varStr); err == nil {
		return val
	}

	return defaultVal
}

func GetEnvAsBool(key string, defaultVal bool) bool {
	varStr := GetEnv(key, "")
	if val, err := strconv.ParseBool(varStr); err == nil {
		return val
	}

	return defaultVal
}

func GetEnvAsSlice(key string, defaultVal []string, sep string) []string {
	valStr := GetEnv(key, "")
	if valStr == "" {
		return defaultVal
	}

	val := strings.Split(valStr, sep)

	return val
}
