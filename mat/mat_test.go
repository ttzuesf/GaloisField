package mat

import (
	"crypto/rand"
	"crypto/sha256"
	"fmt"
	"github.com/number/hash/merkletree"
	"testing"
	//mat1 "gonum.org/v1/gonum/mat"
)

type ShaContent struct {
	x []byte
}

// CalculateHash hashes the values of a TestSHA256Content
func (t ShaContent) CalculateHash() ([]byte, error) {
	h := sha256.New()
	if _, err := h.Write(t.x); err != nil {
		return nil, err
	}

	return h.Sum(nil), nil
}

// Equals tests for equality of two Contents
func (t ShaContent) Equals(other merkletree.Content) (bool, error) {
	return true, nil
}

func TestMerkletree(t *testing.T) {
	data := make([]merkletree.Content, 12)
	for i := 0; i < 12; i++ {
		data[i] = ShaContent{nil}
	}
	MT, _ := merkletree.NewTree(data)
	for i, v := range MT.Leafs {
		fmt.Println(i, v)
	}
}

func BenchmarkMerkletree(b *testing.B) {
	f := 50
	n := 3*f + 1
	data := make([]merkletree.Content, n)
	buf := make([]byte, 32)
	for i := 0; i < n; i++ {
		rand.Read(buf)
		data[i] = ShaContent{buf}
	}
	for i := 0; i < b.N; i++ {
		merkletree.NewTreeWithHashStrategy(data, sha256.New)
	}
}
