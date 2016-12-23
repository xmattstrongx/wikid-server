package app

import "github.com/caarlos0/env"

type Config struct {
	Port      int    `env:"PORT"        envDefault:"5000"`
	DbConnStr string `env:"DB_CONN_STR" envDefault:"root:password@tcp(localhost:3306)/wikid"`
}

func GetConfig() (*Config, error) {
	config := &Config{}
	err := env.Parse(&config)
	return config, err
}
