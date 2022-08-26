package Reed_solomon

import (
	"log"
	"number/field/Polyn"
	"number/field/base"
	"number/field/matrix"
	"testing"
)

func TestSyndrome(t *testing.T) {
	var cof = []int{1, 2, 1, 1}
	a := 2
	//power of a
	e := make([]int, 10)
	for i := 0; i < 10; i++ {
		c := base.Pow(a, i, 11)
		e[i] = c
		//log.Println(c,y)
	}
	//
	cof1 := make([]int, 10)
	log.Println("code", cof1)
	for i := 0; i < 10; i++ {
		y := Polyn.F(cof, e[i], 11)
		log.Println(e[i], y)
		cof1[9-i] = y
	}
	log.Println("code", cof1)
	d1 := syndrome(cof1, e[1:7], 11)
	log.Println("correct", d1)
	rec := []int{1, 5, 2, 5, 2, 2, 1, 9, 0, 0}
	d2 := syndrome(rec, e[1:7], 11)
	log.Println("wrong", d2)
}

func TestSG(t *testing.T) {
	// evaluate S(x)Gama(x) mod x^{6}
	cof1 := []int{3, 4, 1, 1, 8, 0}
	cof2 := []int{2, -3, 1}
	res := Polyn.Multipoly(cof1, cof2, 11)
	log.Println(res)
	for i := 1; i < len(res); i++ {
		if res[i] < 0 {
			res[i] = 11 + res[i]
		}
	}
	if len(res) >= 6 {
		l := len(res)
		log.Println(res[l-6:])
	}
	log.Println("Xi:", SG(cof1, cof2, 3, 11))
}

func TestEncode(t *testing.T) {
	var cof = []int{3, 2, 1, 1}
	log.Println(Encode(cof, 2, 11))
}

func TestDecode(t *testing.T) {
	var cof = []int{1, 4, 2, 5, 10, 2, 1, 9, 2, 7}
	res := Decode(cof, 2, 4, 11)
	for i := 0; i < len(res); i++ {
		res[i] = (10 * res[i]) % 11
	}
	log.Println("res:", res)
}

func TestLambda(t *testing.T) {
	m := make([][]int, 0)
	m = append(m, []int{0, 8, 10}, []int{8, 10, 3}, []int{10, 3, 3})
	y := []int{8, 8, 4}
	x, err := matrix.Solvequals(m, y, 11)
	if err != nil {
		log.Println(err)
		return
	}
	x = append(x, 1)
	log.Println("f:", x)
	lam := Polyn.Multipoly([]int{-3, 1}, []int{1, 1}, 11)
	log.Println(Polyn.F(x, 4, 11), Polyn.F(x, 10, 11))
	log.Println("Lambda", lam)
}

func TestMatrix(t *testing.T) {
	p := 11
	M := make([][]int, 0)
	M = append(M, []int{2, 5, 6}, []int{5, 5, 10})
	M = base.Gaosifp(M, p)
	log.Println(M)
}

func TestFFInv(t *testing.T) {
	var cof = []int{1, 1, 2, 3}
	a := 2
	p := 11
	//power of a
	c := make([]int, 10)
	for i := 0; i < 10; i++ {
		x := base.Pow(a, i, 11)
		c[i] = Polyn.F(cof, x, p)
		//log.Println(c,y)
	}
	log.Println(c)
	res := FFTINv(c, a, p)
	log.Println(res)
}

func TestFFT(t *testing.T) {
	var cof = []int{1, 1, 2, 3}
	a := 2
	p := 11
	//power of a
	m := make([]int, p-1)
	l := len(cof)
	for i := 0; i < l; i++ {
		m[i] = cof[l-1-i]
	}
	res := FFT(m, a, p)
	log.Println(res)
}

func TestCLambda(t *testing.T) {
	p := 11
	errs := []int{3, 7, 6, 7, 10, 8}
	log.Println(lambda(errs, p))
}

func TestSubf(t *testing.T) {
	cb := []int{1, 1}
	cl := []int{0, 4, 5, 6, 7, 1}
	dm := 3
	d := 4
	l := 5
	p := 11
	res := subf(cb, cl, d, dm, l, p)
	log.Println(res, len(res)-len(cl))
}

func TestBerkleMassey(t *testing.T) {
	err := []int{8, 8, 2, 4, 1, 4, 0, 8, 5}
	cof := BerkleMassey(err, 11)
	cof1 := make([]int, len(cof))
	for i := 0; i < len(cof); i++ {
		cof1[i] = cof[len(cof)-1-i]
	}
	log.Println("cof:", cof)
	log.Println("check result:", Polyn.F(cof, 2, 11), Polyn.F(cof, 5, 11), Polyn.F(cof, 1, 11))
}

func TestSugiAlg(t *testing.T) {
	p := 11
	a := 2
	// give a message f(x)
	var f = []int{3, 2, 1, 1}
	c := Encode(f, a, p)
	log.Println("codeword vector", c)
	// evaluate syndrome
	l := len(c) - len(f)
	e := make([]int, 0) // storage from a^{0} to a^{2t}
	for i := 0; i < len(c); i++ {
		c := base.Pow(a, i, p)
		e = append(e, c)
	}
	rec := c
	rec[9] = 0  // x^{0}
	rec[8] = 0  // x^{1}
	rec[1] = 10 //x^{8}
	S := syndrome(c, e[1:l+1], p)
	log.Printf("syndrome S[1] to S[%d]:%v\n", l, S)
	lamb := lambda(S, p)
	log.Printf("coffiece of lambda:%v\n", lamb)
	log.Println(Polyn.F(lamb, base.Inverse(e[10-9-1], p), p), Polyn.F(lamb, base.Inverse(e[10-1-1], p), p))
}

func TestFSugiAlg(t *testing.T) {
	p := 11
	a := 2
	// give a message f(x)
	var f = []int{3, 2, 1, 1}
	c := Encode(f, a, p)
	log.Println("codeword vector", c)
	// evaluate syndrome
	l := len(c) - len(f)
	e := make([]int, 0) // storage from a^{0} to a^{2t}
	for i := 0; i < len(c); i++ {
		d := base.Pow(a, i, p)
		e = append(e, d)
	}
	rec := make([]int, len(c))
	copy(rec, c)
	rec[9] = 0  // x^{0}
	rec[8] = 0  // x^{1}
	rec[1] = 10 //x^{8}
	rec[4] = 5  // x^{5}
	log.Printf("received codeword vector:rec=%v\n", rec)
	S := syndrome(rec, e[1:l+1], p)
	S = reverse(S)
	log.Printf("syndrome S[1] to S[%d]:%v\n", l, S)
	s := S[:4]
	log.Printf("syndrome s[1] to s[%d]:%v\n", 4, s)
	gamma := Polyn.Multipoly([]int{-1, 1}, []int{-2, 1}, p)
	log.Printf("gamma:%d\n", gamma)
	Xi := SG(s, gamma, 2, p)
	xi := reverse(Xi)
	log.Printf("xi:%d\n", xi)
	lamb := lambda(xi, p)
	log.Printf("coefficient of lambda:%v\n", lamb)
	log.Println(Polyn.F(lamb, base.Inverse(e[5], p), p), Polyn.F(lamb, base.Inverse(e[8], p), p))
}

func TestKcz(t *testing.T) {
	p := 11
	a := 2
	// give a message f(x)
	var f = []int{8, 9, -1, -1}
	f1 := make([]int, len(f))
	copy(f1, f)
	c := Encode(f, a, p)
	c = reverse(c)
	log.Println("codeword vector", c)
	d := FFT(c, a, p)
	log.Println(d, f1)
	e := make([]int, 0)
	for i := 0; i < p; i++ {
		c := base.Pow(a, i, p)
		e = append(e, c)
	}
	s := []int{1}
	for i := 1; i <= 6; i++ {
		cof := []int{1, -e[i]}
		s = Polyn.Multipoly(cof, s, p)
		//log.Println(i,s,len(s))
	}
	for i := 0; i < len(s); i++ {
		if s[i] < 0 {
			s[i] = s[i] + p
		}
	}
	log.Println("c", s)
}

func TestGS(t *testing.T) {
	m := make([][]int, 0)
	m = append(m, []int{2, 9, 4})
	m = append(m, []int{9, 4, 1})
	m = append(m, []int{4, 1, 3})
	log.Println(matrix.Det(m, 11))
}
