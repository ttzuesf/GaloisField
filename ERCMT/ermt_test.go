package ERCMT

import (
	"crypto/rand"
	"log"
	"math"
	"testing"
)

func TestERCMTEncode(t *testing.T) {
	M := []byte("Hello Worlds!000")
	f := 1
	n := 3*f + 1
	erc, _ := NewErcMT(n, f)
	res, _ := erc.Encode(M)
	for _, v := range res {
		log.Println(v)
	}
}

func TestERCMTDecode(t *testing.T) {
	M := []byte("Hello Worlds!000")
	f := 3
	n := 3*f + 1
	erc, _ := NewErcMT(n, f)
	res, _ := erc.Encode(M)
	res = res[f-1:]
	rand.Read(res[1].Msg)
	rand.Read(res[1].Tr.Hval[0])
	msg, err := erc.Decode(res, res[f].Root)
	if err != nil {
		log.Fatal("wrong!")
	}
	log.Println(string(msg))
}

func BenchmarkErcMTEncode(b *testing.B) {
	f := 5
	n := 3*f + 1
	buf := make([]byte, 1024*512*(f+1))
	ermt, _ := NewErcMT(n, f)
	rand.Read(buf)
	for i := 1; i < b.N; i++ {
		_, err := ermt.Encode(buf)
		if err != nil {
			log.Println(err)
			return
		}
	}
}

func BenchmarkErcMTDecode(b *testing.B) {
	f := 20
	n := 3*f + 1
	requid := f + 1
	symbolbytes := math.Ceil(1024 * 1024 * 1 / float64(requid))
	M := make([]byte, int(symbolbytes)*requid)
	//M:=make([]byte,1024*512*(f+1)*45);
	rand.Read(M)
	erc, _ := NewErcMT(n, f)
	res, _ := erc.Encode(M)
	res = res[f-1:]
	rand.Read(res[0].Msg)
	//log.Fatal("wrong!");
	rand.Read(res[1].Tr.Hval[0])
	for i := 0; i < b.N; i++ {
		_, err := erc.Decode(res, res[f+1].Root)
		if err != nil {
			log.Fatal("wrong!")
		}
	}
}
