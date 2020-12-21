package main

import (
	"log"
	"os"
	"strconv"
	"strings"

	"github.com/cbi-sh/hashes/internal/hashes/server"
	"github.com/cbi-sh/hashes/internal/hashes/store"
)

func main() {
	ndcs := parseNdcs(getEnv("HASHES_NDCS", "50"), ",")
	capacity, err := strconv.Atoi(getEnv("HASHES_CAPACITY", "1000000"))
	if err != nil {
		log.Fatalf("HASHES_CAPACITY: %v\n", err)
	}
	salt := getEnv("HASHES_SALT", "changeMeSalt")

	st := store.New(ndcs, capacity, salt).Generate()
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

func parseNdcs(value string, sep string) (ndcs []int) {
	for _, ndc := range strings.Split(value, sep) {
		n, err := strconv.Atoi(ndc)
		if err != nil {
			log.Fatalf("parseNdcs: %v\n", err)
		}
		ndcs = append(ndcs, n)
	}
	return
}
