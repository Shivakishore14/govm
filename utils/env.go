package utils

import (
	"github.com/pkg/errors"
	"os"
)

var ErrorEnvNotSet = errors.New("env variable not found")

func GetEnv(key string) (string, error) {
	value := os.Getenv(key)
	if value == "" {
		return "", ErrorEnvNotSet
	}
	return value, nil
}

func GetEnvWithDefault(key string, defaultValue string) string {
	value := os.Getenv(key)
	if value == "" {
		value = defaultValue
	}
	return value
}
func AddToPath(data string) {
	pathValue := os.Getenv("PATH")
	os.Setenv("PATH", data+":"+pathValue)
}
