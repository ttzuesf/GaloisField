package Polyn

import (
	"log"
	"number/field/base"
	"testing"
)

func TestPoly(t *testing.T) {
	var cof = []int{1, 1, 3, 2, 10, 5, 0}
	a := 2
	p := 11
	g := []int{1} // generator polynomial
	for i := 1; i <= 4; i++ {
		b := base.Pow(a, i, p)
		a := []int{1, -b}
		g = Multipoly(g, a, p)
	}
	log.Printf("g(x)=%v\n", g)
	r := Modpoly(cof, g, p)
	log.Printf("r=%v\n", r)
	rs := make([]int, 0)
	for i := 1; i <= 4; i++ {
		b := base.Pow(a, i, p)
		log.Printf("a^%v=%v\n", i, b)
		rs = append(rs, F(r, b, p))
	}
	log.Printf("rs=%v\n", rs)
}

func TestMultipoly(t *testing.T) {
	//test f* constant
	f := []int{1, 2, 3, 4}
	c := []int{-7}
	log.Printf("f=%p,c=%p\n",&f,&c);
	d := Multipoly(f, c, 11)
	log.Printf("d=%+v\n", d);
	log.Printf("f=%p,c=%p\n",&f,&c);
}

func TestFliye(t *testing.T) {
	cof := []int{1, 2, 5}
	l := 10
	S := make([]int, 0)
	for i := 1; i < l; i++ {
		s := 0
		for j := 0; j < len(cof); j++ {
			s = (s + base.Pow(cof[j], i, 11)) % 11
		}
		S = append(S, s)
	}
	log.Println(S)
}

func TestDivpoly(t *testing.T) {
	f := []int{1, -7, 1}
	g := []int{1, -1}
	res := Divpoly(f, g, 11)
	log.Println(res, f, g)
	log.Printf("%+v", Multipoly(res, g, 11))
}

func TestModpoly(t *testing.T) {
	f := []int{1, -2, 1, 7}
	g := []int{1, -1, 2}
	res := Modpoly(f, g, 11)
	p := Divpoly(f, g, 11)
	log.Println(res, f)
	log.Printf("p(x)*g(x)+r(x)=%+v\n", Addpoly(Multipoly(p, g, 11), res, 11))
}

func TestSubpoly(t *testing.T) {
	// test deg(a)>deg(b)
	a := []int{1, 0, 1}
	b := []int{3, 4}
	log.Printf("a(x)-b(x)=%+v\n", Subpoly(a, b, 11))
	// test deg(a)<deg(b);
	a = []int{1, 0}
	b = []int{3, 4, 1}
	log.Printf("a(x)-b(x)=%+v\n", Subpoly(a, b, 11))
}

func TestAddpoly(t *testing.T) {
	// test deg(a)>deg(b)
	a := []int{1, 0, 1}
	b := []int{3, 4}
	log.Printf("a=%v,b=%v\n",a,b);
	c := Subpoly(a, b, 11)
	log.Printf("a=%v,b=%v\n",a,b);
	log.Printf("a(x)-b(x)=%+v\n", Addpoly(c, b, 11))
	log.Printf("a=%v,b=%v\n",a,b);
}
