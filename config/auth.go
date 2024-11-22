package config

import (
	"time"
)

// JWTConfig model.
type JWTConfig struct {
	AccessTokenExpiry  time.Duration `mapstructure:"access_token_expiry"`
	RefreshTokenExpiry time.Duration `mapstructure:"refresh_token_expiry"`
	JWTSecret          string        `mapstructure:"jwt_secret"`
}

// GetTokenConfig returns jwt configuration.
func GetTokenConfig() JWTConfig {
	return JWTConfig{
		AccessTokenExpiry:  time.Duration(cfg.JWT.AccessTokenExpiry) * time.Second,
		RefreshTokenExpiry: time.Duration(cfg.JWT.RefreshTokenExpiry) * time.Second,
		JWTSecret:          cfg.JWT.JWTSecret,
	}
}
