package env

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/defry256/pokemon-helper/config"
	"github.com/defry256/pokemon-helper/internal/logger"
	"github.com/joho/godotenv"
)

var (
	HOST_URL  = "HOST_URL"
	HOST_PORT = "HOST_PORT"
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

func getStringOrDefault(key string, def string) string {
	value := os.Getenv(key)
	if value == "" {
		return def
	}
	return value
}
