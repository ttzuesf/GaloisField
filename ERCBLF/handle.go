package ERCBLF

import (
	"bytes"
	"encoding/binary"
)

func Padding(M[]byte,k int) []byte{
	l:=(len(M)/k)*k;
	if len(M)%k!=0{
		l=l+k;
	}
	padding := l - len(M);
	padtext := bytes.Repeat([]byte{byte(padding)}, padding);
	return append(M, padtext...);
}

func Unpadding(M[]byte) []byte{
	length := len(M)
	unpadding := int(M[length-1])
	return M[:(length - unpadding)]
}

func Icvtb(c []uint64) []byte{
	buf:=make([]byte,0)
	for i:=0;i<len(c);i++{
		b := make([]byte, 8)
		binary.LittleEndian.PutUint64(b,c[i])
		buf=append(buf,b...)
	}
	return buf;
}

func BcvtI(source []byte) []uint64{
	buf:=make([]uint64,0)
	for i:=0;i<len(source)/8;i++{
		buf=append(buf,binary.LittleEndian.Uint64(source[i*8:(i+1)*8]))
	}
	return buf;
}