package config

type ServerConfig struct {
	GRPC *GRPCConfig
	Rest *RestConfig
}

func LoadServerConfig() *ServerConfig {
	return &ServerConfig{
		GRPC: LoadGRPCConfig(),
		Rest: LoadRestConfig(),
	}
}
