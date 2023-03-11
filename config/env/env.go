package env

import (
	"fmt"
	"os"
	"path/filepath"
	"strconv"

	"github.com/defry256/pokemon-helper/config"
	"github.com/defry256/pokemon-helper/internal/logger"
	"github.com/joho/godotenv"
)

var (
	HOST_URL         = "HOST_URL"
	HOST_PORT        = "HOST_PORT"
	MAX_QUEUE_WORKER = "MAX_QUEUE_WORKER"
	REDIS_HOST       = "REDIS_HOST"
	REDIS_PORT       = "REDIS_PORT"
	REDIS_USERNAME   = "REDIS_USERNAME"
	REDIS_PASSWORD   = "REDIS_PASSWORD"
)

type env struct{}

func LoadEnv() {
	envPath, _ := filepath.Abs("../.env")
	err := godotenv.Load(envPath)
	if err != nil {
		logger.Error(fmt.Sprintf("Failed load env - %v", err), err)
	}
	config.SetConfig(&env{})
}

func (e *env) HOST_URL() string {
	return getStringOrDefault(HOST_URL, "localhost")
}

func (e *env) HOST_PORT() string {
	return getStringOrDefault(HOST_PORT, "8000")
}

func (e *env) MAX_QUEUE_WORKER() int {
	return getIntOrDefault(HOST_PORT, 1)
}

func (e *env) REDIS_HOST() string {
	return getStringOrDefault(REDIS_HOST, "")
}

func (e *env) REDIS_PORT() string {
	return getStringOrDefault(REDIS_PORT, "")
}

func (e *env) REDIS_USERNAME() string {
	return getStringOrDefault(REDIS_USERNAME, "")
}

func (e *env) REDIS_PASSWORD() string {
	return getStringOrDefault(REDIS_PASSWORD, "")
}

func getStringOrDefault(key string, def string) string {
	value := os.Getenv(key)
	if value == "" {
		return def
	}
	return value
}

func getIntOrDefault(key string, def int) int {
	value := os.Getenv(key)
	if value == "" {
		return def
	}
	intVal, err := strconv.Atoi(value)
	if err != nil {
		panic(err)
	}
	return intVal
}
