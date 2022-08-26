package base

import (
	"log"
	"math/rand"
	"testing"
	"time"
)

func TestEuclidean(t *testing.T) {
	var a,b int=245,10;
	log.Println(Euclidean(a,b))
}

func TestExtendEuclidean(t *testing.T) {
	rand.Seed(time.Now().Unix())
	a:=rand.Intn(78);
	b:=rand.Intn(32);
	log.Println(a,b,Euclidean(a,b))
	x,y:=ExtendEuclidean(a,b)
	if x*a+y*b==Euclidean(a,b){
		log.Println("True")
	}
}