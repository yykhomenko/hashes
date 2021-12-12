package main

import (
	"github.com/yykhomenko/hashes/internal/config"
	"github.com/yykhomenko/hashes/internal/server"
	"github.com/yykhomenko/hashes/internal/store"
)

func main() {
	config := config.New()
	store := store.New(config)
	store.Generate()
	server := server.New(config, store)
	server.Start()
}
