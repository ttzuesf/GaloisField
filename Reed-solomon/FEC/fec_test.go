package FEC

import (
	"crypto/rand"
	"crypto/sha256"
	"github.com/vivint/infectious"
	"log"
	"math"
	"runtime"
	"testing"
)

func TestNFC(t *testing.T) {
	var fec *infectious.FEC
	var total int = 10
	var requid int = 5
	fec, err := infectious.NewFEC(requid, total)
	if err != nil {
		log.Println("Initialize wrong")
		return
	}
	// Prepare to receive the shares of encoded data.
	shares := make([]infectious.Share, total)
	output := func(s infectious.Share) {
		// the memory in s gets reused, so we need to make a deep copy
		shares[s.Number] = s.DeepCopy()
	}

	// the data to encode must be padded to a multiple of required, hence the
	// underscores.
	buf := []byte("hello, world! __")
	log.Printf("length buf %d\n", len(buf))
	err = fec.Encode(buf, output)
	if err != nil {
		log.Println("Encoding wrong")
		return
	}

	// we now have total shares.
	for _, share := range shares {
		log.Printf("%d: %#v\n", share.Number, string(share.Data))
	}

	// Let's reconstitute with two pieces missing and one piece corrupted.
	shares = shares[:len(shares)-2] // drop the first piece
	shares[2].Data[1] = '!'         // mutate some data
	shares[0].Data[2] = '*'
	results, err := fec.Decode(nil, shares)
	if err != nil {
		log.Println("Decoding wrong")
		return
	}

	// we have the original data!
	log.Printf("got: %#v\n", string(results))
	runtime.GC()
}

func TestECCDecode(t *testing.T) {
	var fec *infectious.FEC
	var total int = 10
	var requid int = 4
	fec, err := infectious.NewFEC(requid, total)
	if err != nil {
		log.Println("Initialize wrong")
		return
	}
	// Prepare to receive the shares of encoded data.
	shares := make([]infectious.Share, total)
	output := func(s infectious.Share) {
		// the memory in s gets reused, so we need to make a deep copy
		shares[s.Number] = s.DeepCopy()
	}

	// the data to encode must be padded to a multiple of required, hence the
	// underscores.
	err = fec.Encode([]byte("hello, world! __"), output)
	if err != nil {
		log.Println("Encoding wrong")
		return
	}

	// we now have total shares.
	for _, share := range shares {
		log.Printf("%d: %#v\n", share.Number, string(share.Data))
	}
	log.Println("=====================================")
	// Let's reconstitute with two pieces missing and one piece corrupted.
	shares = shares[:len(shares)-2] // drop the first piece
	shares[2].Data[1] = '!'         // mutate some data
	shares[0].Data[2] = '*'
	// we now have total shares.
	for _, share := range shares {
		log.Printf("%d: %#v\n", share.Number, string(share.Data))
	}
	results, err := fec.Decode(nil, shares)
	if err != nil {
		log.Println("Decoding wrong")
		return
	}

	// we have the original data!
	log.Printf("got: %#v\n", string(results))
	runtime.GC()
}

// testing padding

func BenchmarkECCEncode(b *testing.B) {
	var fec *infectious.FEC
	var f int = 1
	var total int = 3*f + 1
	var requid int = f + 1
	fec, err := infectious.NewFEC(requid, total)
	if err != nil {
		log.Println("Initialize wrong")
		return
	}
	// Prepare to receive the shares of encoded data.
	shares := make([]infectious.Share, total)
	output := func(s infectious.Share) {
		// the memory in s gets reused, so we need to make a deep copy
		shares[s.Number] = s.DeepCopy()
	}

	// the data to encode must be padded to a multiple of required, hence the
	// underscores.
	buf := make([]byte, 1024*512*requid)
	rand.Read(buf)
	err = fec.Encode([]byte(buf), output)
	if err != nil {
		log.Println("Encoding wrong")
		return
	}
	for i := 0; i < b.N; i++ {
		fec.Encode(buf, output)
	}
	runtime.GC()
}

func BenchmarkECCDecode(b *testing.B) {
	var fec *infectious.FEC
	var f int = 1
	var total int = 3*f + 1
	var requid int = f + 1
	fec, err := infectious.NewFEC(requid, total)
	if err != nil {
		log.Println("Initialize wrong")
		return
	}
	// Prepare to receive the shares of encoded data.
	shares := make([]infectious.Share, total)
	output := func(s infectious.Share) {
		// the memory in s gets reused, so we need to make a deep copy
		shares[s.Number] = s.DeepCopy()
	}

	// the data to encode must be padded to a multiple of required, hence the
	// underscores.
	symbolbytes := math.Ceil(1024 * 1024 * 15 / float64(requid))
	buf := make([]byte, int(symbolbytes)*requid)
	rand.Read(buf)
	hash := sha256.New()
	hash.Write(buf)
	h := hash.Sum(nil)
	err = fec.Encode([]byte(buf), output)
	if err != nil {
		log.Println("Encoding wrong")
		return
	}
	//log.Println(len(shares[2].Data));
	shares = shares[requid-2:] //drops first f-1 pieces
	rand.Read(shares[3].Data)  // set up the f+2 the location is wrong!
	//rand.Read(shares[3].Data);
	for i := 0; i < b.N; i++ {
		dbuf, _ := fec.Decode(nil, shares)
		hash1 := sha256.New()
		hash1.Write(dbuf)
		if string(h) != string(hash1.Sum(nil)) || err != nil {
			log.Println("out!")
			break
		}
	}
	runtime.GC()
}

func BenchmarkERCDecode(b *testing.B) {
	var fec *infectious.FEC
	var total int = 19
	var requid int = 7
	fec, err := infectious.NewFEC(requid, total)
	if err != nil {
		log.Println("Initialize wrong")
		return
	}
	// Prepare to receive the shares of encoded data.
	shares := make([]infectious.Share, total)
	output := func(s infectious.Share) {
		// the memory in s gets reused, so we need to make a deep copy
		shares[s.Number] = s.DeepCopy()
	}

	// the data to encode must be padded to a multiple of required, hence the
	// underscores.
	buf := make([]byte, 1024*400*requid)
	err = fec.Encode([]byte(buf), output)
	if err != nil {
		log.Println("Encoding wrong")
		return
	}
	//log.Println(len(shares[2].Data));
	rand.Read(buf)
	shares = shares[requid:] // drop the first f pieces
	for i := 0; i < b.N; i++ {
		_, err := fec.Decode(nil, shares)
		if err != nil {
			//log.Println("out!");
			break
		}
	}
	runtime.GC()
}

func TestRand(t *testing.T) {
	b := []byte("Alice!")
	log.Println(b)
	rand.Read(b)
	log.Println(b)
}
