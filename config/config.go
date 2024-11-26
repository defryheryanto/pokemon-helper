package config

import (
	"github.com/spf13/viper"
)

var c config

type config struct {
	HostPort           string `mapstructure:"HOST_PORT"`
	MaxQueueWorker     int    `mapstructure:"MAX_QUEUE_WORKER"`
	RedisNetwork       string `mapstructure:"REDIS_NETWORK"`
	RedisAddress       string `mapstructure:"REDIS_ADDRESS"`
	RedisDB            int    `mapstructure:"REDIS_DB"`
	RedisUsername      string `mapstructure:"REDIS_USERNAME"`
	RedisPassword      string `mapstructure:"REDIS_PASSWORD"`
	TracingEnabled     bool   `mapstructure:"TRACING_ENABLED"`
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
	if cfg.RedisNetwork == "" {
		panic("ENV REDIS_NETWORK is empty")
	}
	if cfg.RedisAddress == "" {
		panic("ENV REDIS_ADDRESS is empty")
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
	viper.BindEnv("REDIS_NETWORK")
	viper.BindEnv("REDIS_ADDRESS")
	viper.BindEnv("REDIS_DB")
	viper.BindEnv("REDIS_USERNAME")
	viper.BindEnv("REDIS_PASSWORD")
	viper.BindEnv("TRACING_ENABLED")
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

func RedisNetwork() string {
	return c.RedisNetwork
}

func RedisAddress() string {
	return c.RedisAddress
}

func RedisDB() int {
	return c.RedisDB
}

func RedisUsername() string {
	return c.RedisUsername
}

func RedisPassword() string {
	return c.RedisPassword
}

func TracingEnabled() bool {
	return c.TracingEnabled
}

func JaegerCollectorURL() string {
	return c.JaegerCollectorURL
}

func Environment() string {
	return c.Environment
}
