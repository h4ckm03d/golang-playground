package main

import (
	"crypto/md5"
	"crypto/sha1"
	"crypto/sha256"
	"crypto/sha512"
	"hash"
	"math/rand"
	"testing"

	"golang.org/x/crypto/blake2b"
	"golang.org/x/crypto/sha3"
)

func benchmarkHash(b *testing.B, h hash.Hash) {
	data := make([]byte, 1024)
	rand.Read(data)

	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		h.Write(data)
		h.Sum(nil)
	}
}

func BenchmarkHashMD5(b *testing.B) {
	benchmarkHash(b, md5.New())
}

func BenchmarkHashSHA1(b *testing.B) {
	benchmarkHash(b, sha1.New())
}

func BenchmarkHashSHA256(b *testing.B) {
	benchmarkHash(b, sha256.New())
}

func BenchmarkHashSHA512(b *testing.B) {
	benchmarkHash(b, sha512.New())
}

func BenchmarkHashSHA3256(b *testing.B) {
	benchmarkHash(b, sha3.New256())
}

func BenchmarkHashSHA3512(b *testing.B) {
	benchmarkHash(b, sha3.New512())
}

func BenchmarkHashBLAKE2b256(b *testing.B) {
	h, _ := blake2b.New256(nil)
	benchmarkHash(b, h)
}

func BenchmarkHashBLAKE2b512(b *testing.B) {
	h, _ := blake2b.New512(nil)
	benchmarkHash(b, h)
}
