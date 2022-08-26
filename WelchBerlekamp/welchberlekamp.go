package WelchBerlekamp

import (
	"log"
	"number/field/Polyn"
	"number/field/base"
)

type RS struct {
	N int
	K int
	p int
}

func (rs *RS) Encode(f, g []int) []int {
	xn := []int{1}
	shift := make([]int, rs.N-rs.K)
	xn = append(xn, shift...)
	f1 := Polyn.Multipoly(f, xn, rs.p)
	rx := Polyn.Modpoly(f1, g, rs.p)
	log.Println(rx)
	res := Polyn.Subpoly(f1, rx, rs.p)
	return res
}

func (rs *RS) Decode(f, g []int) []int {
	return nil
}

func reminder(r, g []int, p int) []int {
	return nil
}

func Generator(alpha, dist, p int) []int {
	res := []int{1}
	d := dist
	for i := 1; i <= d; i++ {
		root := []int{1}
		root = append(root, -base.Pow(alpha, i, p))
		res = Polyn.Multipoly(res, root, p)
	}
	return res
}

func welchberlekamp(x, y []int, p int) ([]int, []int) {
	if len(x) != len(y) {
		return nil, nil
	}
	l := len(x)
	N := []int{0}
	V := []int{0}
	M := []int{1}
	W := []int{1}
	for i := 0; i < l; i++ {
		b := (Polyn.F(N, x[i], p) - y[i]*Polyn.F(W, x[i], p)) % p
		//log.Printf("b=%v\n", b)
		//log.Println("i=", i)
		if b == 0 {
			copy(N,N);
			copy(W,W);
			M = Polyn.Multipoly(M, []int{1, -x[i]}, p)
			V = Polyn.Multipoly(V, []int{1, -x[i]}, p)
		} else {
			a := (Polyn.F(M, x[i], p) - y[i]*Polyn.F(V, x[i], p)) % p;
			N1 :=make([]int,len(N)); // origin N
			copy(N1,N);
			V1 :=make([]int,len(V)); // orighin V
			copy(V1,V);
			M1:=make([]int,len(M)); // origin M
			copy(M1,M);
			W1:=make([]int,len(W)); // origin W
			copy(W1,W);
			M = Polyn.Multipoly(N, []int{1, -x[i]}, p);
			V = Polyn.Multipoly(W, []int{1, -x[i]}, p);
			N = Polyn.Addpoly(Polyn.Multipoly(M1, []int{b}, p), Polyn.Multipoly(N1, []int{-a}, p), p);
			W = Polyn.Addpoly(Polyn.Multipoly(V1, []int{b}, p), Polyn.Multipoly(W1, []int{-a}, p), p);
			if rank(N, W) > rank(M, V) {
				C :=make([]int,len(N));
				D :=make([]int,len(V));
				copy(C,N);
				copy(D,V);
				copy(N,M);
				copy(W,V);
				copy(M,C);
				copy(V,D);
			}
		}

	}
	return N, W
}

func rank(N, W []int) int {
	l := 2 * (len(W) - 1)
	if 2*(len(N)-1)+1 > l {
		l = 2*(len(N)-1) + 1
	}
	//log.Println("rank", l)
	return l
}
