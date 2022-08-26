package base

import (
	"log"
	"math/big"
)

// This package describes base arithmetic operators over GF(2^{k})

// k>=2
func Prm(k int) *big.Int{
	a:=primtable[k-2]
	log.Println("a=",a)
	c:=big.NewInt(0);
	c.Add(c,big.NewInt(1));
	//log.Println(c)
	base:=big.NewInt(int64(1));
	c.Add(c,base.Lsh(base,uint(k)));
	//log.Println(c)
	for _,v:=range(a){
		//log.Println(v);
		j:=1<<v;
		c.Add(c,big.NewInt(int64(j)));
	}
	return c;
}


func XOR(a,b *big.Int) *big.Int{
	return a.Xor(a,b)
}

/*
Multiply two polynomials in the GF(2)[x]
*/

func Mul(a,b *big.Int) *big.Int{
	p:=big.NewInt(0); /* accumulator for the product of the multiplication */
	One:=big.NewInt(int64(1));
	a1:=new(big.Int);
	aa:=big.NewInt(int64(0));
	bb:=big.NewInt(int64(0));
	aa.Set(a);
	bb.Set(b);
	for (aa.Cmp(big.NewInt(0))!= 0 && bb.Cmp(big.NewInt(0))!= 0) {
		//log.Println("Enter Loop, b=",b);
		//log.Println(a1.And(b,One).Cmp(One));
		//log.Println("One=",One);
		if (a1.And(b,One).Cmp(One)==0){
			/* if the Polyn for b has a constant term, add the corresponding a to p */
			p.Xor(p,aa); /* addition in GF(2^m) is an XOR of the Polyn coefficients */
		}
		aa.Lsh(aa,uint(1));/* equivalent to a*x */
		bb.Rsh(bb,uint(1));
	}
	return p;
}

/* Multiply two numbers in the GF(2^k) finite field defined
 * by the modulo Polyn relation x^k
 * (the other way being to do carryless multiplication followed by a modular reduction)
 */
func Mulp(a,b *big.Int,prim *big.Int,k int) *big.Int {
	p:=big.NewInt(0); /* accumulator for the product of the multiplication */
	One:=big.NewInt(int64(1));
	One1:=big.NewInt(int64(1)).Lsh(One,uint(k));
	a1:=new(big.Int);
	/*
	aa:=big.NewInt(int64(0));
	bb:=big.NewInt(int64(0));
	aa.Set(a);
	bb.Set(b);*/
	for (a.Cmp(big.NewInt(0))!= 0 && b.Cmp(big.NewInt(0))!= 0) {
		//log.Println("Enter Loop, b=",b);
		//log.Println(a1.And(b,One).Cmp(One));
		//log.Println("One=",One);
		if (a1.And(b,One).Cmp(One)==0){
			/* if the Polyn for b has a constant term, add the corresponding a to p */
			p.Xor(p,a); /* addition in GF(2^m) is an XOR of the Polyn coefficients */
			//log.Printf("p=%x\n",p);
		}
		if (a.Cmp(One1)==0) {
			/* GF modulo: if a has a nonzero term x^k-1, then must be reduced when it becomes x^k */
			//log.Println("Zero=",Zero.And(a,One1));
			a.Lsh(a,uint(1));
			a.Xor(a,prim); /* subtract (XOR) the primitive Polyn*/
			//log.Println(prim);
			//log.Println("a=",a);
		}else{
			a.Lsh(a,uint(1));/* equivalent to a*x */
			//log.Println("a=",a);
		}
		b.Rsh(b,uint(1));
	}
	return p;
}

/*
The follow function is Div function, a/b mod prim
 */
func Div(a,b *big.Int) *big.Int{
	one:=big.NewInt(int64(1));
	b1:=big.NewInt(int64(0));
	re:=big.NewInt(int64(0)); // Storage final result!
	aa:=big.NewInt(int64(0));
	bb:=big.NewInt(int64(0));
	aa.Set(a);
	bb.Set(b);
	for (aa.BitLen()>=bb.BitLen()){
		l:=a.BitLen()-b.BitLen();
		b1.Lsh(bb,uint(l));
		aa.Xor(aa,b1);
		re.Xor(re,b1.Lsh(one,uint(l)))
	}
	return re;
}

/*
The follow function is Div function, a%b mod prim
*/
func Rem(a,b *big.Int) *big.Int{
	b1:=big.NewInt(int64(0));
	aa:=big.NewInt(int64(0));
	bb:=big.NewInt(int64(0));
	aa.Set(a);
	bb.Set(b);
	for (aa.BitLen()>=bb.BitLen()){
		l:=aa.BitLen()-bb.BitLen();
		b1.Lsh(bb,uint(l));
		aa.Xor(aa,b1);
	}
	return aa;
}

/*
The follow function is achieving extending Euclid algorithm Over GF(2^{k})!
*/

func ExtendEuclid(a,b *big.Int)(*big.Int,*big.Int){
	if a.Cmp(b)==-1{ // keep a>b;
		t:=new(big.Int);
		*t=*a;
		*a=*b;
		*b=*t;
	}
	s:=big.NewInt(0);
	r:=big.NewInt(0);
	zero:=big.NewInt(0);
	quotient:=big.NewInt(0);
	temp:=big.NewInt(0);
	r.Set(b);
	old_s:=big.NewInt(1);
	old_r:=big.NewInt(0).Set(a);
	bezout_t:=big.NewInt(0);
	// When r not equal to zero
	for(r.Cmp(zero)!= 0) {
		quotient.Set(Div(old_r,r));
		temp.Set(r);
		r.Xor(old_r, Mul(quotient,r));
		old_r.Set(temp);
		temp.Set(s);
		s.Xor(old_s, Mul(quotient,s));
		old_s.Set(temp);
	}
	if (b.Cmp(zero)!=0) {
		bezout_t =Div(old_r.Xor(r,Mul(old_s,a)),b);
	}
	return old_s,bezout_t;
}

// n = 2 to 256;
var primtable=[256][]int{
	{1},{1},{1},{1,2,3},{1,4,5},{2,3,4},{1,2,7},{3,5,6},{2,3,8},{1,8,10},
	{1,2,10},{3,5,8},{1,11,12},{3,4,12},{10,12,15},{4,12,16},{4,11,16},{3,9,10},{2,7,13},{3,4,9},
	{3,7,12},{4,8,15},{2,5,11},{7,12,13},{13,15,23},{17,22,23},{5,8,24},{2,6,16},{9,10,27},{8,23,25},
	{2,7,16},{11,16,26},{8,12,17},{9,17,27},{7,12,33},{2,14,22},{5,6,27},{16,23,35},{23,27,29},{27,31,32},
	{30,31,34},{5,22,27},{18,35,39},{4,28,39},{18,31,40},{11,24,32},{1,9,19},{16,18,24},{17,31,34},{15,24,46},
	{17,18,22},{20,41,50},{29,49,53},{19,38,50},{29,39,41},{1,16,42},{4,37,52},{26,46,54},{27,28,34},{15,19,44},
	{3,26,57},{20,44,54},{9,34,61},{10,18,38},{39,48,55},{3,33,61},{29,47,62},{20,27,63 },{3,57,69},{48,53,59},
	{2,14,23},{11,50,58},{7,43,68},{14,18,33},{14,29,52},{2,36,52},{16,20,47},{24,28,44},{17,27,75},{9,34,43},
	{27,41,68},{16,33,55},{45,51,59},{11,36,50},{7,10,80},{21,53,56},{15,53,86},{34,67,77},{10,58,71},{29,31,50},
	{13,24,32},{67,77,88},{18,29,80},{11,77,83},{15,17,84},{17,44,93},{26,85,87},{11,38,68},{36,60,81},{26,74,83},
	{15,19,27},{60,80,83},{6,49,89},{70,87,96},{19,86,96},{39,54,59},{3,24,59},{25,58,102},{21,55,97},{5,67,77},
	{2,19,68},{25,80,96},{54,72,103},{8,20,30},{24,27,95},{64,73,74},{50,106,117},{36,52,82},{9,46,88},{33,42,43},
	{35,39,54},{23,51,113},{15,31,43},{65,90,103},{10,70,117},{13,45,54},{11,35,77},
}