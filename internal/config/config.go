package config

import (
	"errors"
	"sync"

	"github.com/kelseyhightower/envconfig"
)

type Config struct {
	AppName  string `envconfig:"APP_NAME" default:"product service"`
	PORT     string `envconfig:"PORT" default:"8080"`
	DBUser   string `envconfig:"DB_USER" default:"root"`
	DBPass   string `envconfig:"DB_PASS" default:"password"`
	DBHost   string `envconfig:"DB_HOST" default:"localhost"`
	DBPort   string `envconfig:"DB_PORT" default:"3306"`
	DBName   string `envconfig:"DB_NAME" default:"productdb"`
	LogLevel string `envconfig:"LOG_LEVEL" default:"debug"`
}

var (
	instance *Config
	once     sync.Once
)

func GetConfig() (*Config, error) {
	var (
		err error
		cfg Config
	)
	once.Do(func() {
		err = envconfig.Process("", &cfg)
		if err == nil {
			instance = &cfg
		}
	})

	if err != nil {
		return nil, err
	}

	if instance == nil {
		return nil, errors.New("Config is nil")
	}

	return instance, nil
}
