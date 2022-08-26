package main

import (
	"crypto/rand"
	"crypto/sha512"
	"log"
	"math"
	"number/hash/bloomfilter"
	"time"
)

func main(){
	n:=15;
	lambda:=6;
	hash:=sha512.New();
	blf,_:=bloomfilter.Newbloomfilter(n,lambda,hash);
	buf:=make([]byte,8);
	var cache [][]byte
	for i:=0;i<n;i++{
		rand.Read(buf)
		cache=append(cache,buf);
	}
	for i:=0;i<len(cache);i++{
		//log.Println(cache[i])
		blf.AddElement(cache[i],hash)
	}
	//log.Println("Witness",blf.Witness)
	ch:=make([]byte,8)
	/*
	for i:=0;i<len(cache);i++{
		//log.Println(cache[i])
		if blf.VerifyElement(cache[i],hash){
			log.Println(i,true)
		}
	}*/
	var steps =int(0)
	start:=time.Now()
	var t time.Duration
	for i:=0;;i++{
		rand.Read(ch);
		if blf.VerifyElement(ch,hash){
			t=time.Now().Sub(start)
			steps=i;
			break
		}
	}
	log.Printf("average steps2^(%f),time=%v\n",math.Log2(float64(steps)),t.Seconds())
}