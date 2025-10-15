package hashes

import (
	"log"

	"github.com/caarlos0/env/v11"
)

type Config struct {
	Addr         string `env:"HASHES_ADDR" envDefault:":8082"`
	CC           string `env:"HASHES_CC" envDefault:"380"`
	NDCS         []int  `env:"HASHES_NDCS" envSeparator:"," envDefault:"67"`
	NDCCap       int    `env:"HASHES_NDC_CAPACITY" envDefault:"10000000"`
	Salt         string `env:"HASHES_SALT" envDefault:"mySalt"`
	MsisdnLenMin int    `env:"HASHES_MSISDN_LENGTH_MIN" envDefault:"12"`
	MsisdnLenMax int    `env:"HASHES_MSISDN_LENGTH_MAX" envDefault:"21"`
}

func NewConfig() *Config {
	c := Config{}
	if err := env.Parse(&c); err != nil {
		log.Fatal(err)
	}
	return &c
}
