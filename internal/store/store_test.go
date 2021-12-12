package store

import (
	"crypto/md5"
	"strconv"
	"testing"

	"github.com/stretchr/testify/assert"
)

const (
	number   = 500000001
	hash     = "89c664c54118ca2b4e803a6fc58f670d"
	salt     = "mySalt"
	capacity = 1000000
)

var ndcs = []int{67}

func TestHash(t *testing.T) {
	s := New(ndcs, capacity, salt)

	expected := hash
	actual := s.Hash(strconv.Itoa(number))

	assert.Equal(t, expected, actual)
}

func TestMsisdn(t *testing.T) {
	s := New(ndcs, capacity, salt)
	s.AddHash(number)

	expected := strconv.Itoa(number)
	actual, _ := s.Msisdn(hash)

	assert.Equal(t, expected, actual)
}

func BenchmarkHash(b *testing.B) {
	s := New(ndcs, capacity, salt)
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
