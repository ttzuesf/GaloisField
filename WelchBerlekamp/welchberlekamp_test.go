package WelchBerlekamp

import (
	"log"
	"number/field/Polyn"
	"number/field/base"
	"testing"
)

func TestRation(t *testing.T) {
	rs := RS{N: 7, K: 3, p: 11}
	f := []int{1, 1, 4}
	g := Generator(2, rs.N-rs.K, 11)
	log.Println("generator polynomial:", g)
	c := rs.Encode(f, g)
	log.Println("coded word", c)
	log.Printf("f=%d\n", Polyn.Modpoly(c, g, 11))
	for i := 0; i < rs.N; i++ {
		a := base.Pow(2, i, rs.p)
		log.Printf("c(alpha^%v)=%+v\n", i, Polyn.F(c, a, rs.p))
	}
}

func TestReminder(t *testing.T) {
	n := 7
	k := 3
	g := Generator(2, n-k, 11)
	c := []int{1, 1, 4, 2, 10, 5, 5}
	e1 := Polyn.Modpoly(c, g, 11)
	log.Printf("correct reminder:%v\n", e1)
	r := []int{1, 1, 3, 2, 10, 5, 0}
	e2 := Polyn.Modpoly(r, g, 11)
	log.Printf("error reminder:%v\n", e2)
}

func TestBW(t *testing.T) {
	p := 11
	f := []int{1, 1, 4}
	var x []int
	var y []int
	for i := 1; i <= 7; i++ {
		a := base.Pow(2, i, p)
		x = append(x, a)
		y = append(y, Polyn.F(f, a, p))
	}
	log.Println(x)
	log.Println(y)
	y[3] = 5
	log.Printf("modified y%d\n", y)
	N, W := welchberlekamp(x, y, p)
	log.Printf("N=%v\n", N)
	log.Printf("W=%v\n", W)
	for i := 0; i < len(x); i++ {
		log.Printf("W(x[%v])=%v\n", i+1, Polyn.F(W, x[i], p))
	}
}

func TestRank(t *testing.T) {
	N := []int{1, 1, 3}
	W := []int{2, 1, 4}
	l := rank(N, W)
	log.Printf("rank[N,W]=%v\n", l)
}
