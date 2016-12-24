package app

import (
	"sync"

	"github.com/caarlos0/env"
)

type Config struct {
	Port      int    `env:"PORT"        envDefault:"5000"`
	DBConnStr string `env:"DB_CONN_STR" envDefault:"root:password@tcp(localhost:3306)/wikid"`
}

var config *Config
var configOnce sync.Once

// GetConfig panics if it fails to get the configuration values.
func GetConfig() *Config {
	configOnce.Do(func() {
		config = &Config{}
		if err := env.Parse(config); err != nil {
			panic(err)
		}
	})

	return config
}
