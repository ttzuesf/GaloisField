package bloomfilter

import (
	"crypto/rand"
	"errors"
	"math"
	"math/big"
)

func ProTrial(index []int,lamb,n,l int) bool{
	if l<lamb*n{
		panic(errors.New("Rong Paramter"));
	}
	vect:=make([]bool,l);
	l1:=new(big.Int)
	l1.SetInt64(int64(l));
	m:=lamb*n;
	for i:=0;i<m;i++{
		r:=urand(l1);
		for _,v:=range index{
			if r==int64(v) {
				vect[v]=true;
			}
		}
		//log.Printf("vect[%d]=%t\n",i,vect)
	}
	// check whether occuring false mischecking
	for _,v:=range index{
		if vect[v]!=true{
			return false
		}
	}
	return true;
}
// verify trial
func FProTrial(lamb,n,l int) []bool{
	if l<lamb*n{
		panic(errors.New("Rong Paramter"));
	}
	vect:=make([]bool,l);
	l1:=new(big.Int)
	l1.SetInt64(int64(l));
	m:=lamb*n;
	for i:=0;i<m;i++{
		r:=urand(l1);
		vect[r]=true;
		//log.Printf("vect[%d]=%t\n",i,vect)
	}
	return vect;
}

func Fverify(lamb,l int,vect[]bool) bool{
	l1:=new(big.Int)
	l1.SetInt64(int64(l));
	//log.Println(l1);
	index:=make([]int,0)
	for i:=0;i<lamb;i++{
		r:=urand(l1);
		index=append(index,int(r));
	}
	//log.Println(index);
	for _,v:=range index{
		if vect[v]!=true{
			return false
		}
	}
	//log.Println(vect)
	return true;
}

func urand(l *big.Int) int64{
	r,_:=rand.Int(rand.Reader,l)
	return r.Int64()
}
//
func CalProbability(lamb,n int) float64{
	l:=math.Ceil(float64(lamb*n)/math.Log(2));
	m:=float64(lamb*n)
	p1:=math.Pow((1-1/l),m)
	p:=math.Pow(1-p1,float64(lamb))
	return p;
}