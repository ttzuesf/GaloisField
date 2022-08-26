package ERCBLF

import (
	"crypto/sha256"
	"crypto/sha512"
	"errors"
	"log"

	//"github.com/bits-and-blooms/bitset"
	blf "github.com/bits-and-blooms/bloom"
	rs "github.com/klauspost/reedsolomon"
	"math"

	//blf "github.com/number/hash/bloomfilter" //own code
	"github.com/vivint/infectious"
	"hash"
)

type Share struct {
	Idx int    // index of coded Message
	Msg []byte // coded Message
	Vec []byte //coded vector
	Dig []byte //digest of vector
}

type ErcBlf struct {
	N   int // total nodes;
	F   int // maximum fault nodes;
	Erc rs.Encoder
	//BLF   *blf.Bloomfilter
	BLF   *blf.BloomFilter
	Fec   *infectious.FEC
	hf256 hash.Hash
	hf512 hash.Hash
}

//Encode M to M,b,h

func (eb *ErcBlf) Encode(M []byte) ([]Share, error) {

	//Encode original message!
	shards, err := eb.Erc.Split(M)
	//log.Println(M);
	if err != nil {
		return nil, err
	}
	eb.Erc.Encode(shards)
	// Add verification information to a bits vector
	for i := 0; i < len(shards); i++ {
		//eb.BLF.Add(shards[i])
		h := sha512.New()
		h.Write(shards[i])
		buf := h.Sum(nil)
		eb.BLF.Add(buf)
	}
	//cache := eb.BLF.Witness.Bytes()
	//vect := Icvtb(cache)
	vect, _ := eb.BLF.GobEncode()

	vect = Padding(vect, eb.F+1)
	eb.hf256.Write(vect)
	h := eb.hf256.Sum(nil)
	eb.hf256.Reset()
	//Error Correct Encoding the bits vector, vect;
	vshards := make([]infectious.Share, eb.N)
	output := func(s infectious.Share) {
		// the memory in s gets reused, so we need to make a deep copy
		vshards[s.Number] = s.DeepCopy()
	}
	err = eb.Fec.Encode([]byte(vect), output)
	if err != nil {
		return nil, errors.New("Encoding Wrong Bits vector!")
	}
	res := make([]Share, 0)
	for i := 0; i < len(shards); i++ {
		var sh Share
		sh.Msg = append(sh.Msg, shards[i]...)       // Message.
		sh.Vec = append(sh.Vec, vshards[i].Data...) // coded vector of Bloom filter!
		sh.Idx = i
		sh.Dig = append(sh.Dig, h...)
		res = append(res, sh)
	}
	return res, nil
}

func (eb *ErcBlf) Decode(M []Share) ([]byte, error) {
	//recontruct origin vect by ECC;
	vshs := make([]infectious.Share, 0)
	for _, v := range M {
		if v.Vec != nil {
			var sh infectious.Share
			sh.Data = v.Vec
			sh.Number = v.Idx
			vshs = append(vshs, sh)
		}
	}
	//log.Println("Hello:vect",vshs);
	vect, err := eb.Fec.Decode(nil, vshs)
	if err != nil {
		return nil, err
	}
	eb.hf256.Write(vect)
	h := eb.hf256.Sum(nil)
	if string(M[0].Dig) != string(h) {
		log.Printf("Wrong vect")
		return nil, errors.New("Wrong vect")
	}
	eb.hf256.Reset()
	eb.BLF.GobDecode(vect)
	mshs := make([][]byte, eb.N)
	//buf:=make([]byte,len(M[0].Msg));
	for i := 0; i < len(M); i++ {
		/*
			    	if eb.BLF.VerifyElement(M[i].Msg){
			    		mshs[M[i].Idx]=M[i].Msg;
					};
		*/
		h := sha512.New()
		h.Write(M[i].Msg)
		buf := h.Sum(nil)
		if eb.BLF.Test(buf) {
			mshs[M[i].Idx] = M[i].Msg
		}
		if len(mshs) == eb.F+1 {
			break
		}
	}
	//log.Println("mshs:",mshs);
	err = eb.Erc.Reconstruct(mshs)
	if err != nil {
		return nil, err
	}
	//log.Println("Hello:Msg");
	res := make([]byte, 0)
	for i := 0; i < eb.F+1; i++ {
		res = append(res, mshs[i]...)
	}
	return res, nil
}

func NewErcBlf(n, f int, lambda int) (*ErcBlf, error) {
	erc := new(ErcBlf)
	if n != 3*f+1 {
		return nil, errors.New("Wrong Parameter!")
	}
	erc.F = f
	erc.N = n
	erc.hf256 = sha256.New()
	erc.hf512 = sha512.New()
	//blf, err := blf.Newbloomfilter(erc.N, lambda, erc.hf512)
	l := int(math.Ceil(float64(lambda*n) / math.Log(2)))
	blf := blf.New(uint(l), uint(lambda))
	erc.BLF = blf
	//Error Erasure encoding origin message!
	//err := nil;
	erc.Erc, _ = rs.New(erc.F+1, erc.N-erc.F-1)
	/*if err != nil {
		return nil, err
	}*/
	//Error correct coding origin message!
	erc.Fec, _ = infectious.NewFEC(erc.F+1, erc.N)
	return erc, nil
}
