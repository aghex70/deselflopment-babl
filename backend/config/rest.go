package config

import "github.com/spf13/viper"

type RestConfig struct {
	Host string `mapstructure:"SERVER_HOST"`
	Port int    `mapstructure:"SERVER_PORT"`
}

func LoadRestConfig() *RestConfig {
	cfg := &RestConfig{}
	err := viper.Unmarshal(cfg)
	if err != nil {
		return nil
	}
	return cfg
}
