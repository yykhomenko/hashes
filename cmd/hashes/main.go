package main

import (
	"github.com/yykhomenko/hashes/pkg/config"
	"github.com/yykhomenko/hashes/pkg/server"
	"github.com/yykhomenko/hashes/pkg/store"
)

func main() {
	config := config.New()
	store := store.New(config)
	store.Generate()
	server := server.New(config, store)
	server.Start()
}
