package ERCMT

import (
	"crypto/sha512"
	"github.com/cbergoon/merkletree"
)

type ShaContent struct {
	x []byte
}

// CalculateHash hashes the values of a TestSHA256Content
func (t ShaContent) CalculateHash() ([]byte, error) {
	h := sha512.New()
	if _, err := h.Write(t.x); err != nil {
		return nil, err
	}

	return h.Sum(nil), nil
}

// Equals tests for equality of two Contents
func (t ShaContent) Equals(other merkletree.Content) (bool, error) {

	return string(t.x) == string(other.(ShaContent).x), nil
}

type Path struct {
	Hval [][]byte
	Indx []int64
}

func (p *Path) Rebuild() {

}
