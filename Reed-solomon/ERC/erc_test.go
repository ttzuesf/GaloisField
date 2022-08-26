package ERC

import (
	rs "github.com/klauspost/reedsolomon"
	"log"
	"testing"
	"crypto/rand"
	"runtime"
)

func TestErasurecode(t *testing.T){
	datashards:=2;
	parityshards:=2;
	enc,err:=rs.New(datashards, parityshards)
	if err!=nil{
		log.Println("wrong parameters");
	}
	buf:=[]byte("Hello world!");
	shards, err := enc.Split(buf)
	err=enc.Encode(shards);
	//_,err = enc.Verify(data)
	if err==nil{
		log.Println(len(shards));
	}
	log.Println(shards);
	shards[0]=make([]byte,0);
	shards[3]=make([]byte,0);
	log.Println(shards);
	//reconstruct origin data;
	err=enc.Reconstruct(shards);
	if err!=nil{
		log.Println("reconstruct wrong")
	}
	//var dst io.Writer
	//err = enc.Join(dst, shards, len(shards[0])*2)
	//buf1:=make([]byte,len(shards[0])*2)
	//dst.Write(buf1)
	log.Println(shards)

}

func BenchmarkERsureEncode(b *testing.B){
	f:=6;
	total:=3*f+1;
	requid:=f+1;
	parityshards:=total-requid;
	enc,err:=rs.New(requid, parityshards)
	if err!=nil{
		log.Println("wrong parameters");
	}
	buf:=make([]byte,1024*400*requid);
	rand.Read(buf);
	shards, err := enc.Split(buf)
	for i:=0;i<b.N;i++{
		err=enc.Encode(shards);
	}
}

func BenchmarkERsureDecode(b *testing.B){
	f:=1;
	total:=3*f+1;
	requids:=f+1;
	parityshards:=total-requids;
	enc,err:=rs.New(requids, parityshards)
	if err!=nil{
		log.Println("wrong parameters");
	}
	buf:=make([]byte,1024*400*requids);
	rand.Read(buf);
	shards, err := enc.Split(buf)
	err=enc.Encode(shards);
	for i:=0;i<b.N;i++{
		err=enc.Reconstruct(shards);
		if err!=nil{
			log.Println("reconstruct wrong")
		}
		for i:=0;i<requids;i++{
			shards[i]=make([]byte,0);
		}
	}
	//var dst io.Writer
	//err = enc.Join(dst, shards, len(shards[0])*2)
	//buf1:=make([]byte,len(shards[0])*2)
	//dst.Write(buf1)
	//log.Println(shards)
	runtime.GC();
}