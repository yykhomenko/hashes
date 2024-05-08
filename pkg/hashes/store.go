package hashes

import (
	"crypto/md5"
	"encoding/hex"
	"log"
	"runtime"
	"strconv"
	"sync"
	"time"
)

type Store struct {
	config *Config
	sync.RWMutex
	msisdns map[[16]byte]uint32
}

func NewStore(config *Config) *Store {
	return &Store{
		config:  config,
		msisdns: make(map[[16]byte]uint32, len(config.NDCS)*config.NDCCap),
	}
}

func (s *Store) Generate() {
	log.Printf("generate %d hashes...", len(s.config.NDCS)*s.config.NDCCap)
	defer timeTrack(time.Now(), "generate")

	for _, ndc := range s.config.NDCS {
		s.generate(ndc)
	}
}

func (s *Store) generate(ndc int) {
	minNum := ndc*s.config.NDCCap + 0
	maxNum := ndc*s.config.NDCCap + s.config.NDCCap - 1

	var workers = runtime.GOMAXPROCS(-1)
	numbers := make(chan int, 10*workers)
	go func() {
		defer close(numbers)
		for number := minNum; number <= maxNum; number++ {
			numbers <- number
		}
	}()

	wg := &sync.WaitGroup{}
	for i := 0; i < workers; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for number := range numbers {
				hash := md5.Sum([]byte(strconv.Itoa(number) + s.config.Salt))
				s.Lock()
				s.msisdns[hash] = uint32(number)
				s.Unlock()
			}
		}()
	}
	wg.Wait()
}

func (s *Store) Hash(msisdn string) string {
	sum := md5.Sum([]byte(msisdn + s.config.Salt))
	return hex.EncodeToString(sum[:])
}

func (s *Store) Msisdn(hash string) (string, bool) {
	s.RLock()
	h := fromHex(hash)
	if m, ok := s.msisdns[h]; ok {
		s.RUnlock()
		return strconv.Itoa(int(m)), true
	}
	s.RUnlock()
	return "", false
}

func (s *Store) AddHash(number int) {
	hash := md5.Sum([]byte(strconv.Itoa(number) + s.config.Salt))
	s.msisdns[hash] = uint32(number)
}

func timeTrack(start time.Time, name string) {
	log.Printf("%s took %s", name, time.Since(start))
}

func fromHex(hash string) [16]byte {
	key, _ := hex.DecodeString(hash)
	var arr [md5.Size]byte

	for i := 0; i < len(arr); i++ {
		arr[i] = key[i]
	}

	return arr
}
