package main

import (
	"log"

	"github.com/caarlos0/env"

	"github.com/cbi-sh/hashes/internal/hashes/server"
	"github.com/cbi-sh/hashes/internal/hashes/store"
)

type Config struct {
	NDCS   []int  `env:"HASHES_NDCS" envSeparator:"," envDefault:"50"`
	NDCCap int    `env:"HASHES_NDC_CAPACITY" envDefault:"10000000"`
	Salt   string `env:"HASHES_SALT" envDefault:"mySalt"`
}

func main() {
	cfg := Config{}
	if err := env.Parse(&cfg); err != nil {
		log.Fatal(err)
	}
	st := store.New(cfg.NDCS, cfg.NDCCap, cfg.Salt).Generate()
	server.New(st).Start()
}
