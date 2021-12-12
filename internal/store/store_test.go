package store

import (
	"crypto/md5"
	"strconv"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/yykhomenko/hashes/internal/config"
)

const (
	number = 500000001
	hash   = "89c664c54118ca2b4e803a6fc58f670d"
)

var testConfig = &config.Config{
	NDCS:   []int{50},
	NDCCap: 10000000,
	Salt:   "mySalt",
}

func TestHash(t *testing.T) {

	s := New(testConfig)

	expected := hash
	actual := s.Hash(strconv.Itoa(number))

	assert.Equal(t, expected, actual)
}

func TestMsisdn(t *testing.T) {
	s := New(testConfig)
	s.AddHash(number)

	expected := strconv.Itoa(number)
	actual, _ := s.Msisdn(hash)

	assert.Equal(t, expected, actual)
}

func BenchmarkHash(b *testing.B) {
	s := New(testConfig)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		s.Hash(strconv.Itoa(number))
	}
}

func BenchmarkMD5(b *testing.B) {
	s := []byte("57149cb6-991c-4ffd-9c9")
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		md5.Sum(s)
	}
}
