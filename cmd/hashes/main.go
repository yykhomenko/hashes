package main

import (
	"log"
	"os"
	"strings"

	"github.com/cbi-sh/hashes/internal/hashes/server"
	"github.com/cbi-sh/hashes/internal/hashes/store"
)

func main() {
	salt := os.Getenv("HASHES_SALT")
	salt = strings.TrimSpace(salt)
	if salt == "" {
		log.Println("env HASHES_SALT is not set, used: mySalt")
	}

	st := store.New(salt).Generate()
	server.New(st).Start()
}
