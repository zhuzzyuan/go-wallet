package config

type PostgresConfig struct {
	Url string `mapstructure:"url"`
}

func GetPostgresConfig() PostgresConfig {
	return cfg.Postgres
}

func GetPostgresConnectionURL() string {
	return cfg.Postgres.Url
}
