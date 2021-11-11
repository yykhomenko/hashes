package store

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
	ndcs   []int
	ndcCap int
	salt   string
	sync.RWMutex
	msisdns map[[16]byte]uint32
}

func New(ndcs []int, ndcCap int, salt string) *Store {
	return &Store{
		ndcs:    ndcs,
		ndcCap:  ndcCap,
		salt:    salt,
		msisdns: make(map[[16]byte]uint32, len(ndcs)*ndcCap),
	}
}

func (s *Store) Generate() *Store {
	log.Printf("generate %d hashes...", len(s.ndcs)*s.ndcCap)
	defer timeTrack(time.Now(), "generate")

	for _, ndc := range s.ndcs {
		s.generate(ndc)
	}

	return s
}

func (s *Store) generate(ndc int) {
	min := ndc*s.ndcCap + 0
	max := ndc*s.ndcCap + s.ndcCap - 1

	var workers = runtime.GOMAXPROCS(-1)
	numbers := make(chan int, 10*workers)
	go func() {
		defer close(numbers)
		for number := min; number <= max; number++ {
			numbers <- number
		}
	}()

	wg := &sync.WaitGroup{}
	for i := 0; i < workers; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for number := range numbers {
				hash := md5.Sum([]byte(strconv.Itoa(number) + s.salt))
				s.Lock()
				s.msisdns[hash] = uint32(number)
				s.Unlock()
			}
		}()
	}
	wg.Wait()
}

func (s *Store) Hash(msisdn string) string {
	sum := md5.Sum([]byte(msisdn + s.salt))
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
	hash := md5.Sum([]byte(strconv.Itoa(number) + s.salt))
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
