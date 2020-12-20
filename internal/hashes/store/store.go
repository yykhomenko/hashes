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

const (
	capacity  = 1000000
	numberMin = capacity - capacity
	numberMax = capacity - 1
)

var ndcs = map[string]int{
	// "50": 50,
	// "63": 63,
	// "66": 66,
	//
	"67": 67,
	// "68": 68,
	// "73": 73,
	//
	// "91": 91,
	// "92": 92,
	// "93": 93,
	//
	// "94": 94,
	// "95": 95,
	// "96": 96,
	//
	// "97": 97,
	// "98": 98,
	// "99": 99,
}

var allCapacity = len(ndcs) * capacity

type Store struct {
	salt string
	sync.RWMutex
	msisdns map[[16]byte]uint32
}

func New(salt string) *Store {
	return &Store{
		salt:    salt,
		msisdns: make(map[[16]byte]uint32, allCapacity),
	}
}

func (s *Store) Generate() *Store {
	log.Printf("generate %d hashes...", allCapacity)
	defer timeTrack(time.Now(), "generate")

	for _, ndc := range ndcs {
		s.generate(ndc)
	}

	return s
}

func (s *Store) generate(ndc int) {
	min := ndc*capacity + numberMin
	max := ndc*capacity + numberMax

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
	if m, ok := s.msisdns[fromHex(hash)]; ok {
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
	elapsed := time.Since(start)
	log.Printf("%s took %s", name, elapsed)
}

func fromHex(hash string) [16]byte {
	key, _ := hex.DecodeString(hash)
	var arr [md5.Size]byte

	for i := 0; i < len(arr); i++ {
		arr[i] = key[i]
	}

	return arr
}
