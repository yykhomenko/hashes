package main

import (
	"github.com/cbi-sh/hashes/internal/hashes/server"
	"github.com/cbi-sh/hashes/internal/hashes/store"
)

func main() {
	st := store.New().Generate()
	server.New(st).Start()
}
