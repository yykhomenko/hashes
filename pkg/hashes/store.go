package hashes

import (
	"context"
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"log"
	"runtime"
	"strconv"
	"sync"
	"time"
)

type Store struct {
	config  *Config
	mu      sync.RWMutex
	msisdns map[[16]byte]uint32
}

func NewStore(cfg *Config) *Store {
	capacity := len(cfg.NDCS) * cfg.NDCCap
	if capacity <= 0 {
		capacity = 1_000_000 // fallback
	}
	return &Store{
		config:  cfg,
		msisdns: make(map[[16]byte]uint32, capacity),
	}
}

func (s *Store) Generate(ctx context.Context) error {
	total := len(s.config.NDCS) * s.config.NDCCap
	log.Printf("🔁 Generating %d hashes across %d NDCs...", total, len(s.config.NDCS))

	for _, ndc := range s.config.NDCS {
		select {
		case <-ctx.Done():
			return fmt.Errorf("generation cancelled: %w", ctx.Err())
		default:
			s.generateNDC(ndc)
		}
	}
	log.Printf("✅ Generation completed (%d entries)", len(s.msisdns))
	return nil
}

func (s *Store) generateNDC(ndc int) {
	min := ndc * s.config.NDCCap
	max := min + s.config.NDCCap

	workers := runtime.GOMAXPROCS(0)
	ch := make(chan int, 100*workers)

	go func() {
		defer close(ch)
		for n := min; n < max; n++ {
			ch <- n
		}
	}()

	wg := sync.WaitGroup{}
	wg.Add(workers)

	for i := 0; i < workers; i++ {
		go func() {
			defer wg.Done()
			local := make(map[[16]byte]uint32, 5000)

			for n := range ch {
				sum := md5.Sum([]byte(strconv.Itoa(n) + s.config.Salt))
				local[sum] = uint32(n)

				// невелике батч-злиття для зменшення lock contention
				if len(local) >= 5000 {
					s.mergeLocal(local)
					local = make(map[[16]byte]uint32, 5000)
				}
			}

			if len(local) > 0 {
				s.mergeLocal(local)
			}
		}()
	}

	wg.Wait()
}

func (s *Store) mergeLocal(local map[[16]byte]uint32) {
	s.mu.Lock()
	defer s.mu.Unlock()
	for k, v := range local {
		s.msisdns[k] = v
	}
}

func (s *Store) Hash(msisdn string) string {
	sum := md5.Sum([]byte(msisdn + s.config.Salt))
	return hex.EncodeToString(sum[:])
}

func (s *Store) Msisdn(hash string) (string, bool) {
	h, ok := fromHex(hash)
	if !ok {
		return "", false
	}
	s.mu.RLock()
	defer s.mu.RUnlock()
	if m, found := s.msisdns[h]; found {
		return strconv.Itoa(int(m)), true
	}
	return "", false
}

func (s *Store) AddHash(number int) {
	sum := md5.Sum([]byte(strconv.Itoa(number) + s.config.Salt))
	s.mu.Lock()
	s.msisdns[sum] = uint32(number)
	s.mu.Unlock()
}

func timeTrack(start time.Time, name string) {
	log.Printf("%s took %s", name, time.Since(start).Truncate(time.Millisecond))
}

func fromHex(hash string) ([16]byte, bool) {
	data, err := hex.DecodeString(hash)
	if err != nil || len(data) != md5.Size {
		return [16]byte{}, false
	}
	var arr [16]byte
	copy(arr[:], data)
	return arr, true
}
