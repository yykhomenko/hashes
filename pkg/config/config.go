package config

import (
	"github.com/caarlos0/env/v11"
	"log"
)

type Config struct {
	NDCS   []int  `env:"HASHES_NDCS" envSeparator:"," envDefault:"50"`
	NDCCap int    `env:"HASHES_NDC_CAPACITY" envDefault:"10000000"`
	Salt   string `env:"HASHES_SALT" envDefault:"mySalt"`
}

func New() *Config {
	c := Config{}
	if err := env.Parse(&c); err != nil {
		log.Fatal(err)
	}
	return &c
}
