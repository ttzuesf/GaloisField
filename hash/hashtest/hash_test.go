package hashtest

import (
	"crypto/rand"
	"crypto/sha256"
	"github.com/spaolacci/murmur3"
	"math"
	"testing"
)

func BenchmarkMurmur128(b *testing.B) {
	f := 70
	n := 3*f + 1
	lambda := 160
	buf := make([]byte, 128)
	var cache [][]byte
	for i := 0; i < n; i++ {
		rand.Read(buf)
		cache = append(cache, buf)
	}
	for i := 0; i < b.N; i++ {
		for i := 0; i < lambda/4; i++ {
			hash := murmur3.New128()
			hash.Write(cache[i])
			hash.Sum(nil)
		}
	}
}

func BenchmarkSha256(b *testing.B) {
	f := 10
	n := 3*f + 1
	buf := make([]byte, 256)
	var cache [][]byte
	for i := 0; i < n; i++ {
		rand.Read(buf)
		cache = append(cache, buf)
	}
	for i := 0; i < b.N; i++ {
		for i := 0; i < int(math.Ceil(math.Log2(float64(n)))); i++ {
			hash := sha256.New()
			hash.Write(cache[i])
			hash.Sum(nil)
		}
	}
}
