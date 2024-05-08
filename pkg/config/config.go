package config

import (
	"github.com/caarlos0/env/v11"
	"log"
)

type Config struct {
	Addr         string `env:"HASHES_ADDR" envSeparator:"," envDefault:":8080"`
	CC           string `env:"HASHES_CC" envSeparator:"," envDefault:"380"`
	NDCS         []int  `env:"HASHES_NDCS" envSeparator:"," envDefault:"50"`
	NDCCap       int    `env:"HASHES_NDC_CAPACITY" envDefault:"10000000"`
	Salt         string `env:"HASHES_SALT" envDefault:"mySalt"`
	MsisdnLenMin int    `env:"HASHES_MSISDN_LENGTH_MIN" envDefault:"12"`
	MsisdnLenMax int    `env:"HASHES_MSISDN_LENGTH_MIN" envDefault:"21"`
}

func New() *Config {
	c := Config{}
	if err := env.Parse(&c); err != nil {
		log.Fatal(err)
	}
	return &c
}
