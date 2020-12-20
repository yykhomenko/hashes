package main

import (
	"log"
	"os"
	"strings"

	"github.com/cbi-sh/hashes/internal/hashes/server"
	"github.com/cbi-sh/hashes/internal/hashes/store"
)

func main() {
	salt := getEnv("HASHES_SALT", "changeMeSalt")
	st := store.New(salt).Generate()
	server.New(st).Start()
}

func getEnv(key, fallback string) string {
	value, exists := os.LookupEnv(key)
	if !exists {
		log.Printf("env %q is not set, used: %q", key, fallback)
		return fallback
	}
	return strings.TrimSpace(value)
}
