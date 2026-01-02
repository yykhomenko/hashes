package main

import (
	"context"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/yykhomenko/hashes/pkg/hashes"
)

func main() {
	log.SetFlags(log.LstdFlags | log.Lmicroseconds)

	cfg := hashes.NewConfig()

	store := hashes.NewStore(cfg)

	ctx, cancel := signal.NotifyContext(context.Background(), os.Interrupt, syscall.SIGTERM)
	defer cancel()

	start := time.Now()
	if err := store.Generate(ctx); err != nil {
		log.Fatalf("❌ failed to generate hashes: %v", err)
	}
	log.Printf("✅ hashes generated in %s", time.Since(start).Truncate(time.Millisecond))

	server := hashes.NewServer(cfg, store)
	go func() {
		if err := server.Start(); err != nil {
			log.Fatalf("❌ server exited with error: %v", err)
		}
	}()

	<-ctx.Done()
	log.Println("🛑 Shutting down...")
	log.Println("✅ Server stopped gracefully")
}
