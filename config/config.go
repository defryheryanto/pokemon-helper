package config

import (
	"github.com/spf13/viper"
)

var c config

type config struct {
	HostPort           string `mapstructure:"HOST_PORT"`
	MaxQueueWorker     int    `mapstructure:"MAX_QUEUE_WORKER"`
	RedisHost          string `mapstructure:"REDIS_HOST"`
	RedisPort          string `mapstructure:"REDIS_PORT"`
	RedisUsername      string `mapstructure:"REDIS_USERNAME"`
	RedisPassword      string `mapstructure:"REDIS_PASSWORD"`
	JaegerCollectorURL string `mapstructure:"JAEGER_COLLECTOR_URL"`
	Environment        string `mapstructure:"ENVIRONMENT"`
}

func (cfg *config) ValidateKeys() {
	if cfg.HostPort == "" {
		panic("ENV HOST_PORT is empty")
	}
	if cfg.MaxQueueWorker == 0 {
		panic("ENV MAX_QUEUE_WORKER is 0")
	}
	if cfg.RedisHost == "" {
		panic("ENV REDIS_HOST is empty")
	}
	if cfg.RedisPort == "" {
		panic("ENV REDIS_PORT is empty")
	}
	if cfg.RedisUsername == "" {
		panic("ENV REDIS_USERNAME is empty")
	}
	if cfg.RedisPassword == "" {
		panic("ENV REDIS_PASSWORD is empty")
	}
	if cfg.JaegerCollectorURL == "" {
		panic("ENV JAEGER_COLLECTOR_URL is empty")
	}
	if cfg.Environment == "" {
		panic("ENV ENVIRONMENT is empty")
	}
}

func Load() {
	viper.BindEnv("HOST_PORT")
	viper.BindEnv("MAX_QUEUE_WORKER")
	viper.BindEnv("REDIS_HOST")
	viper.BindEnv("REDIS_PORT")
	viper.BindEnv("REDIS_USERNAME")
	viper.BindEnv("REDIS_PASSWORD")
	viper.BindEnv("JAEGER_COLLECTOR_URL")
	viper.BindEnv("ENVIRONMENT")

	err := viper.Unmarshal(&c)
	if err != nil {
		panic(err)
	}

	c.ValidateKeys()
}

func HostPort() string {
	return c.HostPort
}

func MaxQueueWorker() int {
	return c.MaxQueueWorker
}

func RedisHost() string {
	return c.RedisHost
}

func RedisPort() string {
	return c.RedisPort
}

func RedisUsername() string {
	return c.RedisUsername
}

func RedisPassword() string {
	return c.RedisPassword
}

func JaegerCollectorURL() string {
	return c.JaegerCollectorURL
}

func Environment() string {
	return c.Environment
}
