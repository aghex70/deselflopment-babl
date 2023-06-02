package config

import "github.com/spf13/viper"

type GRPCConfig struct {
	GRPCHost string `mapstructure:"GRPC_HOST"`
	GRPCPort int    `mapstructure:"GRPC_PORT"`
}

func LoadGRPCConfig() *GRPCConfig {
	cfg := &GRPCConfig{}
	err := viper.Unmarshal(cfg)
	if err != nil {
		return nil
	}
	return cfg
}
