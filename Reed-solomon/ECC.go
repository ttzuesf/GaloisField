package Reed_solomon

import (
	"log"
	"number/field/Polyn"
	"number/field/base"
	"number/field/matrix"
)

// syndrome
func syndrome(cof1 []int, x []int, p int) []int {
	res := make([]int, len(x))
	if len(x) == 0 {
		return nil
	}
	for k, v := range x {
		res[k] = Polyn.F(cof1, v, p)
	}
	return res
}

func lambda(err []int, p int) []int {
	res := make([]int, 0)
	cofs := make([][]int, 0)
	l := len(err)
	for v := l / 2; v >= 0; v-- {
		y := make([]int, v)
		for i := 0; i < v; i++ {
			d := make([]int, v)
			copy(d, err[i:i+v+1])
			y[i] = (p - err[i+v]) % p
			cofs = append(cofs, d)
		}
		//log.Println(cofs,y,v)
		p, err := matrix.Solvequals(cofs, y, p)
		if err != nil {
			cofs = make([][]int, 0)
			continue
		}
		res = p
		break
	}
	res = append(res, 1)
	return res
}

//FFINVFp

func FFTINv(c []int, a, p int) []int {
	n := len(c)
	if n != p-1 {
		return nil
	}
	inv := base.Inverse(n, p)
	cof := make([]int, n)
	for i := 0; i < n; i++ {
		cof[n-1-i] = c[i]
	}
	res := make([]int, n)
	for j := 0; j < n; j++ {
		x := base.Pow(a, n-j, p)
		res[j] = (Polyn.F(cof, x, p) * inv) % p
	}
	return res
}

//FFTFp

func FFT(m []int, a, p int) []int {
	n := len(m)
	if n != p-1 {
		return nil
	}
	cof := make([]int, n)
	for i := 0; i < n; i++ {
		cof[n-1-i] = m[i]
	}
	res := make([]int, n)
	for i := 0; i < n; i++ {
		x := base.Pow(a, i, p)
		res[i] = Polyn.F(cof, x, p) % p
	}
	return res
}

//c(x)=f(x)g(x)

func Encode(f []int, a int, p int) []int {
	k := len(f)
	n := p - 1
	for i := 0; i < k; i++ {
		f[i] = f[i] % p
	}
	log.Println("f1", f)
	e := make([]int, 0)
	for i := 0; i < p; i++ {
		c := base.Pow(a, i, p)
		e = append(e, c)
	}
	s := f
	for i := 1; i <= n-k; i++ {
		cof := []int{1, -e[i]}
		s = Polyn.Multipoly(cof, s, p)
		//log.Println(i,s,len(s))
	}
	for i := 0; i < len(s); i++ {
		if s[i] < 0 {
			s[i] = s[i] + p
		}
	}
	return s
}

// f(x)=c(x) / p(x)

func Decode(c []int, a, k int, p int) []int {
	n := p - 1
	e := make([]int, 0)
	for i := 0; i < p; i++ {
		c := base.Pow(a, i, p)
		e = append(e, c)
	}
	s := []int{1}
	for i := 1; i <= n-k; i++ {
		cof := []int{1, -e[i]}
		s = Polyn.Multipoly(cof, s, p)
		log.Println(i, s, len(s))
	}
	res := Polyn.Divpoly(c, s, p)
	return res
}

// Berlekmap-Massey algorithm

func BerkleMassey(err []int, p int) []int {
	L := 0
	N := len(err)
	cb := []int{1} //c(x)
	px := []int{1} //p(x)
	l := 1
	dm := 1 // previous discrepancy
	for k := 1; k <= N; k++ {
		s := 0
		for i := 1; i <= L; i++ {
			s = (s + cb[i]*err[k-i-1]) % p
		}
		d := (err[k-1] + s) % p
		log.Println(d)
		if d == 0 {
			l = l + 1
		} else {
			if 2*L >= k {
				cb = subf(cb, px, d, dm, l, p)
				log.Println("(2L>k)the length of cb,L", cb, px, k, L)
				l = l + 1
			} else {
				t := make([]int, len(cb))
				copy(t, cb)                    // old c(x);
				cb = subf(cb, px, d, dm, l, p) // new c(x);
				log.Println("(2L<k)the length of cb,L", cb, px, k, L)
				L = k - L
				px = t
				dm = d
				l = 1
				//log.Println("2v<k",cb,cl,d,dm,l)
			}
		}
	}
	return cb
}

// c(x)=c(x)-ddm^{-1}x^{l}p(x)
func subf(cb, px []int, d, dm, l, p int) []int {
	a := (d * base.Inverse(dm, p)) % p //ddm^{-1}
	t := make([]int, 0)
	t = append(t, make([]int, l+1)...)
	log.Println("t=", t)
	t[0] = 1
	t = Polyn.Multipoly(t, px, p)
	t = Polyn.Multipoly([]int{a}, t, p) //t(x)=ddm^{-1}x^{l}p(x);
	log.Printf("t=%v\n", t)
	res := Polyn.Subpoly(cb, t, p) // c(x)-t(x)
	return res
}

// xi=S(x)*G(x) mode x^{2t}
func SG(syndr []int, gamma []int, t, p int) []int {
	res := Polyn.Multipoly(syndr, gamma, 11)
	for i := 1; i < len(res); i++ {
		if res[i] < 0 {
			res[i] = 11 + res[i]
		}
	}
	if len(res) >= 2*t {
		l := len(res)
		res = res[l-2*t:]
	}
	return res
}

// reverse [a1,a2,a3...,an] to [an,an-1,...,a1]
func reverse(a []int) []int {
	res := make([]int, len(a))
	for i := 0; i < len(a); i++ {
		res[i] = a[len(a)-i-1]
	}
	return res
}
