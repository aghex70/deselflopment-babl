package config

import "github.com/spf13/viper"

type RestConfig struct {
	Host string `mapstructure:"HOST"`
	Port int    `mapstructure:"PORT"`
}

func LoadRestConfig() *RestConfig {
	cfg := &RestConfig{}
	err := viper.Unmarshal(cfg)
	if err != nil {
		return nil
	}
	return cfg
}
