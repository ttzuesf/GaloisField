package base

import (
	"crypto/rand"
	"fmt"
	"log"
	"math/big"
	"testing"
)

func BenchmarkBase(t *testing.B){
	a:=big.NewInt(56)
	b:=big.NewInt(32)
	c:=big.NewInt(0)
	t.Run("Testing1",func(t *testing.B){
		c= XOR(a,b)
	})
	log.Println("c=",c)
	a1:=65
	b1:=32
	c1:=0
	t.Run("Testing2",func(t *testing.B){
		c1=a1^b1
	})
	log.Println("c=",c1)
}

func TestPrimitive(t *testing.T){
	c:=Prm(128)
	log.Println(c)
}

func TestMul(t *testing.T){
	prim:=Prm(8);
	log.Printf("prim=%x\n",prim)
	result:=Mulp(big.NewInt(4),big.NewInt(12),prim,8)
	fmt.Printf("result=%x\n",result);
}

func TestMul1(t *testing.T){
	a:=big.NewInt(11);
	b:=big.NewInt(4);
	res:=Mul(a,b);
	log.Printf("a=%d,b=%d,result=%d\n",a,b,res);
}

func BenchmarkMulp(t *testing.B) {
	k:=8;
	max:=big.NewInt(1);
	max.Lsh(max,uint(k))
	a,_:=rand.Int(rand.Reader,max);
	b,_:=rand.Int(rand.Reader,max);
	log.Printf("a=%d,b=%d\n",a,b)
	prim:=Prm(k);
	log.Println("prim=",prim)
	t.Run("Mul on GF(2^{k})", func(t *testing.B) {
		for i:=0;i<t.N;i++{
			Mulp(a,b,prim,k)
		}
	})
	p, _ := rand.Prime(rand.Reader, k+1)
	log.Printf("p=%d\n",p);
	t.Run("Mul on GF(p)", func(t *testing.B) {
		for i:=0;i<t.N;i++{
			a.Mul(a,b).Mod(a,p)
		}
	})
	// add on GF(p)
	t.Run("Add on GF(2^{k})", func(t *testing.B) {
		for i:=0;i<t.N;i++ {
			XOR(a, b)
		}
	})
	t.Run("Add on GF(p)", func(t *testing.B) {
		for i:=0;i<t.N;i++{
			a.Add(a,b).Mod(a,p)
		}
	})
}

func TestDivRem(t *testing.T){
	a:=big.NewInt(11);
	b:=big.NewInt(4);
	res:=Div(a,b)
	fmt.Printf("a=%d,b=%d,bresult=%x\n",a,b,res);
	rem:=Rem(big.NewInt(11),big.NewInt(4))
	fmt.Printf("a=%d,b=%d,rem=%x\n",a,b,rem);
}

func TestExtendEuclid(t *testing.T) {
	a,b:=ExtendEuclid(big.NewInt(11),big.NewInt(4))
	log.Println(a,b)
}

func TestBaseOperation(t *testing.T){
	a := big.NewInt(4);
	c := big.NewInt(12);
	c.And(a, big.NewInt(1))
	log.Println(c, a)
}