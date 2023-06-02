package config

import (
	"github.com/spf13/viper"
)

const (
	CONFIG_NAME string = ".env"
	CONFIG_TYPE string = "env"
	CONFIG_PATH string = "./"
)

type Config struct {
	Cache     *CacheConfig
	Database  *DatabaseConfig
	Server    *ServerConfig
}

func NewConfig() (*Config, error) {
	viper.AddConfigPath(CONFIG_PATH)
	viper.SetConfigName(CONFIG_NAME)
	viper.SetConfigType(CONFIG_TYPE)
	viper.AutomaticEnv()

	err := viper.ReadInConfig()
	if err != nil {
		panic(err)
	}

	return &Config{
		Cache:     LoadCacheConfig(),
		Database:  LoadDatabaseConfig(),
		Server:    LoadServerConfig(),
	}, nil
}
