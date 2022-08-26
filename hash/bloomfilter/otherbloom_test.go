package bloomfilter

import (
	"crypto/rand"
	"github.com/bits-and-blooms/bloom"
	"log"
	"math"
	"math/big"
	"testing"
)

func TestOtherBlf(t *testing.T) {
	f := 1
	lambda := 128
	n := 3*f + 1
	l := int(math.Ceil(float64(lambda*n) / math.Log(2)))
	blf := bloom.New(uint(l), uint(lambda))
	c, _ := blf.GobEncode()
	log.Println(c)
}

func BenchmarkOtherBlf(b *testing.B) {
	f := 1
	lambda := 128
	n := 3*f + 1
	l := int(math.Ceil(float64(lambda*n) / math.Log(2)))
	blf := bloom.New(uint(l), uint(lambda))
	buf := make([]byte, 1024*400)
	var cache [][]byte
	for i := 0; i < n; i++ {
		rand.Read(buf)
		cache = append(cache, buf)
	}
	for i := 0; i < b.N; i++ {
		for i := 0; i < len(cache); i++ {
			blf.Add(cache[i])
		}
	}
	c, _ := blf.GobEncode()
	log.Println(c)
}

// test whole probability
func TestOProbability(t *testing.T) {
	n := 5
	lambda := 10
	l := int(math.Ceil(float64(lambda*n) / math.Log(2)))
	aver := float64(0)
	for i := 0; i < 50; i++ {
		blf := bloom.New(uint(l), uint(lambda))
		blf.GobEncode()
		var cache []*big.Int
		for i := 0; i < n; i++ {
			num := new(big.Int)
			num.SetInt64(10000000000000000)
			num1, _ := rand.Int(rand.Reader, num)
			cache = append(cache, num1)
		}
		//log.Println(len(cache),cache)
		for i := 0; i < len(cache); i++ {
			//log.Println(cache[i])
			blf.Add(cache[i].Bytes())
		}
		sum := 0
		for j := 0; j < 100000; j++ {
			var a *big.Int
		tag:
			for { // determing element different from testing caching
				num := new(big.Int)
				num.SetInt64(10000000000000000)
				num1, _ := rand.Int(rand.Reader, num)
				for k, v := range cache {
					if num1.Cmp(v) == 0 {
						break
					}
					if k == len(cache)-1 {
						a = num1
						break tag
					}
				}
			}
			if blf.Test(a.Bytes()) {
				sum++
			}
		}
		aver += float64(sum) / float64(100000)
		//log.Println(float64(sum) / float64(10000))
	}
	log.Printf("probability (%f)", aver/50)
}
