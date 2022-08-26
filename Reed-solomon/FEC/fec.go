package FEC

import (
	"crypto/sha256"
	"github.com/vivint/infectious"
	"log"
)

type Fec struct {
	*infectious.FEC
}

type Msg struct {
	h      []byte
	Symbol infectious.Share
	Number int
}

func (f *Fec) Encode(M []byte, n, k int) []Msg {
	hash := sha256.New()
	hash.Write(M)
	h := hash.Sum(nil)
	log.Printf("%v\n", h)
	fec, err := infectious.NewFEC(k, n)
	if err != nil {
		log.Println("Initialize wrong")
		return nil
	}
	// Prepare to receive the shares of encoded data.
	shares := make([]infectious.Share, n)
	output := func(s infectious.Share) {
		// the memory in s gets reused, so we need to make a deep copy
		shares[s.Number] = s.DeepCopy()
	}
	err = fec.Encode(M, output)
	if err != nil {
		log.Println("Encoding wrong")
		return nil
	}
	for i := 0; i < len(shares); i++ {
		//var msg Msg;

	}
	return nil
}
