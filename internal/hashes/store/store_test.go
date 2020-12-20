package store

import (
	"crypto/md5"
	"strconv"
	"testing"

	"github.com/stretchr/testify/assert"
)

const (
	number = 670001122
	hash   = "ecd6250df6a523b6e4457f09cd2696af"
	salt   = "mySalt"
)

func TestHash(t *testing.T) {
	s := New(salt)

	expected := hash
	actual := s.Hash(strconv.Itoa(number))

	assert.Equal(t, expected, actual)
}

func TestMsisdn(t *testing.T) {
	s := New(salt)
	s.AddHash(number)

	expected := strconv.Itoa(number)
	actual, _ := s.Msisdn(hash)

	assert.Equal(t, expected, actual)
}

func BenchmarkHash(b *testing.B) {
	s := New(salt)
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
