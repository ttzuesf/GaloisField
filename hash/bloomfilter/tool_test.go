package bloomfilter

import (
	"crypto/rand"
	"crypto/sha256"
	"fmt"
	"log"
	"testing"
)

func TestBytesToUint64(t *testing.T) {
	b := make([]byte, 7)
	rand.Read(b)
	res := BytesToUint64(b)
	log.Println(res)
}

func TestFunction(t *testing.T) {
	h := sha256.New()
	h.Write([]byte("hello!"))
	for i := 0; i < 7; i++ {
		h.Write([]byte{1})
		fmt.Println(h.Sum(nil))
	}
}

func BenchmarkSha256(b *testing.B) {
	buf := make([]byte, 1024)
	rand.Read(buf)
	for i := 0; i < b.N; i++ {
		h := sha256.New()
		h.Write(buf[0:])
		h.Write([]byte{1})
		//h.Sum(nil)
	}
}
