package hashes

import (
	"crypto/md5"
	"strconv"
	"testing"

	"github.com/stretchr/testify/assert"
)

const (
	number = 500000001
	hash   = "89c664c54118ca2b4e803a6fc58f670d"
)

var testConfig = &Config{
	NDCS:   []int{50},
	NDCCap: 1000000,
	Salt:   "mySalt",
}

func TestStore_Hash(t *testing.T) {
	s := NewStore(testConfig)
	expected := hash
	actual := s.Hash(strconv.Itoa(number))
	assert.Equal(t, expected, actual)
}

func TestStore_Msisdn(t *testing.T) {
	s := NewStore(testConfig)
	s.AddHash(number)
	expected := strconv.Itoa(number)
	actual, _ := s.Msisdn(hash)
	assert.Equal(t, expected, actual)
}

func BenchmarkMD5(b *testing.B) {
	n := []byte(strconv.Itoa(number))
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		md5.Sum(n)
	}
}

func BenchmarkStore_Hash(b *testing.B) {
	s := NewStore(testConfig)
	n := strconv.Itoa(number)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		s.Hash(n)
	}
}

func BenchmarkStore_Msisdn(b *testing.B) {
	s := NewStore(testConfig)
	s.AddHash(number)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		s.Msisdn(hash)
	}
}
