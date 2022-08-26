package bloomfilter

import (
	"crypto/rand"
	"fmt"
	"log"
	"math"
	"math/big"
	"strconv"
	"testing"
)

func TestNewBloomfilter(t *testing.T) {
	n := 12
	lambda := 256
	blf, _ := Newbloomfilter(n, lambda)
	log.Println(blf)
}

func TestNewblf(t *testing.T) {
	n := 12
	lambda := 60
	blf, _ := Newbloomfilter(n, lambda)
	log.Println(blf.l, blf.N, blf.Witness)
}

func TestAddelment(t *testing.T) {
	n := 4
	lambda := 60
	blf, _ := Newbloomfilter(n, lambda)
	log.Println(blf.l, blf.N, blf.Witness.Len())
	blf.AddElement([]byte("hello"))
	log.Println(blf.Witness)
}

func TestVerifyElement(t *testing.T) {
	n := 4
	lambda := 6
	blf, _ := Newbloomfilter(n, lambda)
	var buf []string
	for i := 0; i < n; i++ {
		buf = append(buf, "alice"+strconv.Itoa(i+19))
	}
	for i := 0; i < len(buf); i++ {
		blf.AddElement([]byte(buf[i]))
	}
	log.Printf("buf[0]=%s\n", string(buf[0]))
	if blf.VerifyElement([]byte(buf[1])) {
		fmt.Println(buf[0], "true")
	}
	if blf.VerifyElement([]byte("hello")) {
		fmt.Println("false")
	}
}

func BenchmarkAddElement(b *testing.B) {
	f := 50
	n := 3*f + 1
	lambda := 128
	blf, _ := Newbloomfilter(n, lambda)
	buf := make([]byte, 32)
	var cache [][]byte
	for i := 0; i < n; i++ {
		rand.Read(buf)
		cache = append(cache, buf)
	}
	for i := 0; i < b.N; i++ {
		for i := 0; i < len(cache); i++ {
			blf.AddElement(cache[i])
		}
	}
}

// test condition probabiltiy
func TestCProbability(t *testing.T) {
	n := 1000
	lambda := 100
	blf, _ := Newbloomfilter(n, lambda)
	//log.Println("Witness",blf.Witness)
	cache := make([][]byte, 0)
	for i := 0; i < n; i++ {
		num := make([]byte, 8)
		rand.Read(num)
		cache = append(cache, num)
	}
	count := 0
	for i := 0; i < 50; i++ {
		j := 0
		for i := 0; i < len(cache); i++ {
			//log.Println(cache[i])
			blf.AddElement(cache[i])
		}
		for ; ; j++ {
			var a *big.Int
			num := new(big.Int)
			num.SetInt64(10000000000000000)
			a, _ = rand.Int(rand.Reader, num)
			/*
				tag:
				for ;;{
					num:=new(big.Int)
					num.SetInt64(10000000000000000)
					num1,_:=rand.Int(rand.Reader,num)
					for k,v:=range cache{
						if num1.Cmp(v)==0{
							break
						}
						if k==len(cache)-1{
							a=num1;
							break tag;
						}
					}
				}*/
			if blf.VerifyElement(a.Bytes()) {
				break
			}
		}
		count += j
	}
	log.Printf("average steps 2^(%f)\n", math.Log2(float64(count)/float64(50)))
}

// test whole probability
func TestProbability(t *testing.T) {
	n := 10
	lambda := 12
	//log.Println("Witness",blf.Witness)
	aver := float64(0)
	for i := 0; i < 50; i++ {
		blf, _ := Newbloomfilter(n, lambda)
		var cache []*big.Int
		for i := 0; i < n; i++ {
			num := new(big.Int)
			num.SetInt64(1000)
			num1, _ := rand.Int(rand.Reader, num)
			cache = append(cache, num1)
		}
		//log.Println(len(cache),cache)
		for i := 0; i < len(cache); i++ {
			//log.Println(cache[i])
			blf.AddElement(cache[i].Bytes())
		}
		sum := 0
		for j := 0; j < 100000; j++ {
			var a *big.Int
		tag:
			for { // determing element different from testing caching
				num := new(big.Int)
				num.SetInt64(1000)
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
			if blf.VerifyElement(a.Bytes()) {
				sum++
			}
		}
		aver += float64(sum) / float64(100000)
		//log.Println(float64(sum) / float64(10000))
	}
	log.Printf("probability (%f)", aver/50)
}

func TestFBHash(t *testing.T) {
	data := [][]byte{[]byte("Hello1"), []byte("Hello2"), []byte("Hello3")}
	f := 1
	n := 3*f + 1
	lambda := 160
	blf, _ := Newbloomfilter(n, lambda)
	for i := 0; i < len(data); i++ {
		blf.AddElement(data[i])
	}
	data[0] = []byte("00000")
	for i := 0; i < len(data); i++ {
		log.Println(blf.VerifyElement(data[i]))
	}
}

// noting deepcopy().
func TestBVerifyElement(t *testing.T) {
	f := 1
	n := 3*f + 1
	lambda := 160
	blf, _ := Newbloomfilter(n, lambda)
	var cache [][]byte
	for i := 0; i < n; i++ {
		buf := make([]byte, 10)
		rand.Read(buf)
		cache = append(cache, buf)
	}
	log.Println(cache)
	for i := 0; i < len(cache); i++ {
		blf.AddElement(cache[i])
	}
	log.Println(cache)
	rand.Read(cache[1])
	log.Println(cache)
	for i := 0; i < len(cache); i++ {
		fmt.Println(blf.VerifyElement(cache[i]))
	}
}

func TestGenearte(t *testing.T) {
	n := 100
	lambda := 256
	pr, _ := GeneratePrime(n, lambda)
	//l := math.Ceil(float64(n*lambda) / math.Log(2)) // (ln2)l/n=lambda
	p := math.Pow(1-math.Pow(math.E, -float64(lambda*n)/float64(pr)), float64(lambda))
	log.Println(math.Log2(p))
}
