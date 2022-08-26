package bloomfilter

import (
	"crypto/rand"
	"crypto/sha256"
	"crypto/sha512"
	"encoding/binary"
	"errors"
	"github.com/bits-and-blooms/bitset"
	"math"
)

type Bloomfilter struct {
	Lambda  uint           // secure parameter
	N       int            // the maximum number
	l       int            // witness length
	Witness *bitset.BitSet //the output of vector
}

func (b *Bloomfilter) AddElement(element []byte) {
	h := baseHashes(element)
	for i := uint(0); i < b.Lambda; i++ {
		b.Witness.Set(b.location(h, i))
	}
}

func (b *Bloomfilter) VerifyElement(element []byte) bool {
	h := baseHashes(element)
	for i := uint(0); i < b.Lambda; i++ {
		if !b.Witness.Test(b.location(h, i)) {
			return false
		}
	}
	return true
}

func (b *Bloomfilter) Encode() []byte {
	return nil
}

// baseHashes returns the four hash values of data that are used to create k
// hashes
/*
func baseHashes(data []byte) [4]uint64 {
	a1 := []byte{1} // to grab another bit of data
	hasher := murmur3.New128()
	hasher.Write(data) // #nosec
	v1, v2 := hasher.Sum128()
	hasher.Write(a1) // #nosec
	v3, v4 := hasher.Sum128()
	return [4]uint64{
		v1, v2, v3, v4,
	}
}*/

func baseHashes(data []byte) [8]uint64 {
	h := sha512.New()
	h.Write(data) // #nosec
	buf := h.Sum(nil)
	h = sha256.New()
	h.Write(buf)
	v0 := binary.BigEndian.Uint64(buf[0:8])
	v1 := binary.BigEndian.Uint64(buf[8:16])
	v2 := binary.BigEndian.Uint64(buf[16:24])
	v3 := binary.BigEndian.Uint64(buf[24:32])
	h.Write([]byte{1})
	buf = h.Sum(nil)
	v4 := binary.BigEndian.Uint64(buf[0:8])
	v5 := binary.BigEndian.Uint64(buf[8:16])
	v6 := binary.BigEndian.Uint64(buf[16:24])
	v7 := binary.BigEndian.Uint64(buf[24:32])
	return [8]uint64{v0, v1, v2, v3, v4, v5, v6, v7}
}

// location returns the ith hashed location using the four base hash values
func (b *Bloomfilter) location(h [8]uint64, i uint) uint {
	ii := uint64(i)
	res := h[ii%4] + ii*h[4+(((ii+(ii%4))%8)/4)]
	return uint(res % uint64(b.l))
}

func Newbloomfilter(n, lambda int) (*Bloomfilter, error) {
	l := math.Ceil(float64(n*lambda) / math.Log(2)) // (ln2)l/n=lambda
	bits := math.Ceil(math.Log2(l))
	pr1, err := rand.Prime(rand.Reader, int(bits))
	if err != nil {
		return nil, errors.New("Wrong paramter")
	}
	pr := pr1.Int64()
	j := 0
	for pr < int64(l) {
		if j > 10000 {
			bits++
			j = 0
		}
		pr1, err = rand.Prime(rand.Reader, int(bits))
		if err != nil {
			return nil, errors.New("Wrong paramter")
		}
		j++
		pr = pr1.Int64()
	}
	return &Bloomfilter{
		Lambda:  uint(lambda),
		N:       n,
		Witness: bitset.New(uint(l)),
		l:       int(pr),
	}, nil
}

func GeneratePrime(n, lambda int) (int, error) {
	l := math.Ceil(float64(n*lambda) / math.Log(2)) // (ln2)l/n=lambda
	bits := math.Ceil(math.Log2(l))
	pr1, err := rand.Prime(rand.Reader, int(bits))
	if err != nil {
		return 0, errors.New("Wrong paramter")
	}
	pr := pr1.Int64()
	//j := 0
	for pr < int64(l) {
		/*if j > 10000 {
			bits++
			j = 0
		}*/
		pr1, err = rand.Prime(rand.Reader, int(bits))
		if err != nil {
			return 0, errors.New("Wrong paramter")
		}
		//j++
		pr = pr1.Int64()
	}
	return int(pr), nil
}
