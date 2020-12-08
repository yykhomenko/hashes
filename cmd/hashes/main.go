package main

import (
	"github.com/cbi-sh/hashes/internal/mhs/server"
	"github.com/cbi-sh/hashes/internal/mhs/store"
)

func main() {
	st := store.New().Generate()
	server.New(st).Start()
}
