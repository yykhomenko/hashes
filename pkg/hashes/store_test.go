package hashes

import (
	"crypto/md5"
	"strconv"
	"testing"

	"github.com/stretchr/testify/assert"
)

const (
	testNumber = 500000001
	testHash   = "89c664c54118ca2b4e803a6fc58f670d"
)

var testConfig = &Config{
	NDCS:   []int{50},
	NDCCap: 1000000,
	Salt:   "mySalt",
}

// --- Unit Tests ---

func TestStore_Hash(t *testing.T) {
	s := NewStore(testConfig)
	got := s.Hash(strconv.Itoa(testNumber))
	assert.Equal(t, testHash, got, "Hash() should produce deterministic md5 result")
}

func TestStore_Msisdn(t *testing.T) {
	s := NewStore(testConfig)
	s.AddHash(testNumber)

	msisdn, ok := s.Msisdn(testHash)
	assert.True(t, ok, "Msisdn() should find added hash")
	assert.Equal(t, strconv.Itoa(testNumber), msisdn)
}

func TestStore_Msisdn_NotFound(t *testing.T) {
	s := NewStore(testConfig)
	_, ok := s.Msisdn("ffffffffffffffffffffffffffffffff")
	assert.False(t, ok, "Msisdn() should return false for non-existing hash")
}

func TestStore_AddHash_Idempotent(t *testing.T) {
	s := NewStore(testConfig)
	s.AddHash(testNumber)
	s.AddHash(testNumber)

	msisdn, ok := s.Msisdn(testHash)
	assert.True(t, ok, "hash must still exist after duplicate AddHash")
	assert.Equal(t, strconv.Itoa(testNumber), msisdn)
}

// --- Benchmarks ---

func BenchmarkMD5_Plain(b *testing.B) {
	data := []byte(strconv.Itoa(testNumber))
	b.ReportAllocs()
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		md5.Sum(data)
	}
}

func BenchmarkStore_Hash(b *testing.B) {
	s := NewStore(testConfig)
	input := strconv.Itoa(testNumber)
	b.ReportAllocs()
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		s.Hash(input)
	}
}

func BenchmarkStore_Msisdn(b *testing.B) {
	s := NewStore(testConfig)
	s.AddHash(testNumber)
	b.ReportAllocs()
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		s.Msisdn(testHash)
	}
}

func BenchmarkStore_AddHash(b *testing.B) {
	s := NewStore(testConfig)
	b.ReportAllocs()
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		s.AddHash(testNumber + i)
	}
}

func BenchmarkStore_generateNDC(b *testing.B) {
	cfg := &Config{NDCS: []int{1}, NDCCap: 10_000, Salt: "benchSalt"}
	s := NewStore(cfg)
	b.ReportAllocs()
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		s.generateNDC(1)
	}
}
