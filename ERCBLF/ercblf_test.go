package ERCBLF

import (
	"crypto/rand"
	"log"
	"math"
	"testing"
)

func TestNewErcBlf(t *testing.T) {
	var n, f int = 4, 1
	var lambda int = 160
	blf, err := NewErcBlf(n, f, lambda)
	if err != nil {
		return
	}
	log.Println(&blf.Erc, &blf.BLF, &blf.Fec)
}

func TestErcBlfEncode(t *testing.T) {
	var n, f int = 10, 3
	var lambda int = 128
	erbf, err := NewErcBlf(n, f, lambda)
	if err != nil {
		log.Fatal("Creat New ErcBlf wrong!")
	}
	M := []byte("Hello Worlds!000")
	sh, err := erbf.Encode(M)
	if err != nil {
		log.Fatal("Encoding Wrong!")
	}
	for i := 0; i < len(sh); i++ {
		log.Println("message", i, "equal", sh[i].Msg)
	}
	for i := 0; i < len(sh); i++ {
		log.Println("message", i, "equal", len(sh[i].Vec))
	}
}

func TestPadding(t *testing.T) {
	buf := []byte("hello, world! __")

	buf1 := Padding(buf, 5)
	log.Println(string(buf1), len(buf), len(buf1))
	buf2 := Unpadding(buf1)
	log.Println(string(buf2), len(buf1), len(buf2))
}

func TestDecode(t *testing.T) {
	var n, f int = 10, 3
	var lambda int = 160
	erbf, err := NewErcBlf(n, f, lambda)
	if err != nil {
		log.Fatal("Creat New ErcBlf wrong!")
	}
	M := []byte("Hello Worlds!00014102")
	sh, err := erbf.Encode(M)
	/*for i:=0;i<len(sh);i++{
		log.Printf("%p\n",&sh[i]);
	}*/
	for i := 0; i < f+1; i++ {
		log.Println("correct:", sh[i].Vec)
	}
	for i := 0; i < n; i++ {
		if i < f-1 {
			sh[i].Vec = nil
			sh[i].Msg = make([]byte, 0)
		}
		if i == f {
			rand.Read(sh[i].Vec)
			rand.Read(sh[i].Msg)
		}
	}
	for i := 0; i < f+1; i++ {
		log.Println("wrong", sh[i].Vec)
	}
	buf, err := erbf.Decode(sh)
	log.Println(err, string(buf), len(M))
}

func BenchmarkErcBlf_Encode(b *testing.B) {
	var n, f int = 4, 1
	var lambda int = 160
	erbf, err := NewErcBlf(n, f, lambda)
	if err != nil {
		return
	}
	buf := make([]byte, 1024*10*(f+1))
	rand.Read(buf)
	for i := 1; i < b.N; i++ {
		_, err := erbf.Encode(buf)
		if err != nil {
			log.Println(err)
			return
		}
	}
}

func BenchmarkErcBlfDecode(b *testing.B) {
	f := 20
	n := 3*f + 1
	requid := f + 1
	var lambda int = 64
	erbf, err := NewErcBlf(n, f, lambda)
	if err != nil {
		log.Fatal("Creat New ErcBlf wrong!")
	}
	symbolbytes := math.Ceil(1024 * 1024 * 1.0 / float64(requid))
	M := make([]byte, int(symbolbytes)*requid)
	//M:=make([]byte,1024*512*(f+1)*45);
	rand.Read(M)
	sh, err := erbf.Encode(M)
	for i := 0; i < n; i++ {
		if i < f-1 {
			sh[i].Vec = nil
			sh[i].Msg = make([]byte, 0)
		}
		if i == f+2 {
			rand.Read(sh[i].Vec)
			rand.Read(sh[i].Msg)
		}
	}
	for i := 0; i < b.N; i++ {
		erbf.Decode(sh)
	}
}

func TestIntCovBytes(t *testing.T) {
	source := []uint64{1, 2, 4, 11}
	dst := Icvtb(source)
	log.Println(dst)
	dst1 := BcvtI(dst)
	log.Println(dst1)

}
