package math

import (
	"fmt"
	"math"
	"testing"
)

func TestP(t *testing.T) {
	a := caculate(256, 512)
	fmt.Println(a)
	fmt.Println(math.Pow(2, a))
}

// Caculate complexity of Bloom filter
func TestQ(t *testing.T) {
	fmt.Println(math.Log2(math.Log(2)))
}

// Caculate complexity of Merkle Trees

func TestCom(t *testing.T) {
	fmt.Println(math.Log2(math.Log(2)))
}
