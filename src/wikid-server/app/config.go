package app

import (
	"log"

	"github.com/caarlos0/env"
)

// Config holds application configuration values.
var Config config

type config struct {
	Port int `env:"PORT" envDefault:"5000"`
}

func init() {
	if err := env.Parse(&Config); err != nil {
		log.Fatalln("Failed to set up configuration values.")
	}
}
