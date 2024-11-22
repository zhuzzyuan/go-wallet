package config

import "fmt"

type APIConfig struct {
	Host string `mapstructure:"host"`
	Port string `mapstructure:"port"`
}

func GetAPIEndpoint() string {
	return fmt.Sprintf("%s:%s", cfg.API.Host, cfg.API.Port)
}
