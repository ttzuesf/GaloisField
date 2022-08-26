package Polyn

import (
	"log"
	"testing"
)

func TestF1(t *testing.T) {
	cof:=[]float64{1,2,3}
	res:= F1(cof,3)
	log.Println(res)
	res1:= F2(cof,3)
	log.Println(res1)
}

func TestMultpoly(t *testing.T) {
	cof1:=[]float64{1,2,3}
	cof2:=[]float64{2,3}
	res:=mutipolyf(cof1,cof2)
	log.Println("res:",res)
}

func TestNewTon(t *testing.T){
	x:=[]float64{1,3,2,7,5};
	y:=[]float64{1,2,4,4,3};
	f:=NewTon(x,y)
	for i:=0;i<len(x);i++{
		log.Println(y[i],F2(f,x[i]),"true");
	}
}