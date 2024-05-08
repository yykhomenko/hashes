package main

import (
	"github.com/yykhomenko/hashes/pkg/hashes"
	"log"
)

func main() {
	log.SetFlags(log.LstdFlags | log.Lmicroseconds)
	config := hashes.NewConfig()

	store := hashes.NewStore(config)
	store.Generate()

	server := hashes.NewServer(config, store)
	server.Start()
}
