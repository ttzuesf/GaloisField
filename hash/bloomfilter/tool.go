package bloomfilter

import (
	"encoding/binary"
	"math"
)

func floor(a float64) int {
	return int(math.Floor(a));
}

func Int64ToBytes(i int64) []byte {
	var buf = make([]byte, 8)
	binary.BigEndian.PutUint64(buf, uint64(i))
	return buf
}

func BytesToUint64(buf []byte) []uint64 {
	l:=len(buf);
	if l%8!=0{
		a:=make([]byte,8-l%8);
		a=append(a,buf...);
		buf=a;
	}
	//log.Println(len(buf))
	var res []uint64;
	//res:=BytesToUint64(b)
	for i:=0;8*i<l;i++{
		res=append(res,binary.BigEndian.Uint64(buf[8*i:8*(i+1)]))
	}
	return res;
}