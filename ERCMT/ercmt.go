package ERCMT

import (
	"crypto/sha512"
	"errors"
	merkle "github.com/cbergoon/merkletree"
	rs "github.com/klauspost/reedsolomon"
)

type Share struct {
	Idx  int    // index of coded Message
	Msg  []byte // coded Message
	Tr   Path   //coded vector
	Root []byte
	Dig  []byte //digest of vector
}

type ErcMT struct {
	N   int // total nodes;
	F   int // maximum fault nodes;
	Erc rs.Encoder
}

func (em *ErcMT) Encode(M []byte) ([]Share, error) {
	//Encode original message!
	shards, err := em.Erc.Split(M)
	//log.Println(M);
	if err != nil {
		return nil, err
	}
	em.Erc.Encode(shards)
	//log.Println(shards);
	// build Merkel Tree
	data := make([]merkle.Content, em.N)
	for k, v := range shards {
		data[k] = ShaContent{v}
		//log.Println(&data[k]);
	}
	tree, _ := merkle.NewTreeWithHashStrategy(data, sha512.New)
	tree.MerkleRoot()
	//a:=tree.Leafs[0];
	res := make([]Share, 0)
	for i := 0; i < len(shards); i++ {
		var sh Share
		sh.Msg = append(sh.Msg, shards[i]...)
		//log.Println(data[i]);
		c, b, _ := tree.GetMerklePath(data[i])
		sh.Tr = Path{Hval: c, Indx: b}
		//log.Printf("%v,%v\n",c,b);
		sh.Root = append(sh.Root, tree.MerkleRoot()...)
		sh.Idx = i
		res = append(res, sh)
	}
	return res, nil
}

func (em *ErcMT) Decode(M []Share, root []byte) ([]byte, error) {
	// checke root whether is right
	// stored sorted codes
	mshs := make([][]byte, em.N)
	// get the validate messages
	for _, v := range M {
		if string(v.Root) != string(root) {
			continue
		}
		h := sha512.New()
		h.Write(v.Msg)
		msg := h.Sum(nil)
		for i := 0; i < len(v.Tr.Hval); i++ {
			h := sha512.New()
			if v.Tr.Indx[i] == 1 {
				h.Write(append(msg, v.Tr.Hval[i]...))
			} else {
				h.Write(append(v.Tr.Hval[i], msg...))
			}
			msg = h.Sum(nil)
		}
		if string(msg) == string(root) {
			mshs[v.Idx] = v.Msg
			//log.Println("true");
		}
		if len(mshs) == em.F+1 {
			break
		}
	}
	//log.Println("mshs:",mshs);
	err := em.Erc.Reconstruct(mshs)
	if err != nil {
		return nil, err
	}
	//log.Println("Hello:Msg");
	res := make([]byte, 0)
	for i := 0; i < em.F+1; i++ {
		res = append(res, mshs[i]...)
	}
	return res, nil
}

func (em *ErcMT) recombineMT() {

}

func NewErcMT(n, f int) (*ErcMT, error) {
	erc := new(ErcMT)
	if n != 3*f+1 {
		return nil, errors.New("Wrong Parameter!")
	}
	erc.F = f
	erc.N = n
	//Error Erasure encoding origin message!
	erc.Erc, _ = rs.New(erc.F+1, erc.N-erc.F-1)

	return erc, nil
}
