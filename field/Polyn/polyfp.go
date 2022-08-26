package Polyn

import (
	"number/field/base"
)

// cof=[an,an-1]
func F(cof []int, x int, p int) int {
	if len(cof) == 0 {
		return 0
	}
	r := cof[0]
	for i := 0; i < len(cof)-1; i++ {
		r = (cof[i+1] + x*r) % p
	}
	return r
}

// a(x)+b(x) mod p

func Addpoly(a, b []int, p int) []int {
	l := len(a)
	n := len(b)
	var c []int
	c = a // small degree polynomial
	//log.Println(len(c));
	res := append(make([]int, 0), b...)
	if l > n { //deg(f)<deg(g)
		c = b
		res = append(make([]int, 0), a...)
	}
	d := len(res) - len(c)
	for i := 0; i < len(c); i++ {
		res[i+d] = (res[i+d] + c[i]) % p
		if res[i+d] < 0 {
			res[i+d] = p + res[i+d]
		}
	}
	for i := 0; i < len(res); i++ {
		if res[i] != 0 {
			return res[i:]
		}
	}
	return []int{0}
}

// a(x)-b(x) mod p
// a(x)=a_{l-1}x^{l-1}+a_{l-1}x^{l-2}+...+a_{0}
// b(x)=b_{l-1}x^{l-1}+b_{l-1}x^{l-2}+...+b_{0}
func Subpoly(a, b []int, p int) []int {
	l := len(a)
	n := len(b)
	var c []int
	c = a // small degree polynomial
	//log.Println(len(c));
	res := append(make([]int, 0), b...)
	if l > n { //deg(f)<deg(g)
		c = b
		res = append(make([]int, 0), a...)
	}
	d := len(res) - len(c)
	for i := 0; i < len(c); i++ {
		res[i+d] = (res[i+d] - c[i]) % p
		if res[i+d] < 0 {
			res[i+d] = p + res[i+d]
		}
	}
	return res
}

// a(x)*b(x) mod p

func Multipoly(a, b []int, p int) []int {
	l := len(a)
	n := len(b)
	if l == 0 || n == 0 {
		return nil
	}
	res := make([]int, l+n-1)
	for i := l - 1; i > -1; i-- {
		for j := n - 1; j > -1; j-- {
			res[i+j] = (res[i+j] + a[i]*b[j]) % p
			if res[i+j] < 0 {
				res[i+j] = p + res[i+j]
			}
		}
	}
	for i := 0; i < len(res); i++ {
		if res[i] != 0 {
			res = res[i:]
			break
		}
	}
	return res
}

// f(x)/g(x) mod p
// f(x)=f_{n}x^{n}+f_{n-1}x^{n-1}+f_{1}x+f_{0}

func Divpoly(f, g []int, p int) []int {
	l := len(f)
	n := len(g)
	f1 := make([]int, l)
	for i := 0; i < l; i++ {
		f1[i] = f[i]
	}
	if l < n || n == 0 {
		return nil
	}
	res := make([]int, 0)
	for i := 0; i < l-n+1; i++ {
		a := f1[i] * (base.Inverse(g[0], p))
		if a < 0 {
			a = p + a
		}
		res = append(res, a)
		for j := 0; j < n; j++ {
			f1[j+i] = (f1[j+i] - a*g[j]) % p
			if f1[j+i] < 0 {
				f1[j+i] = p + f1[j+i]
			}
		}
	}
	return res
}

// r(x)= f(x) mod g(x)
func Modpoly(f, g []int, p int) []int {
	l := len(f)
	n := len(g)
	if l < n {
		return append(make([]int, 0), f...)
	}
	f1 := append(make([]int, 0), f...)
	//log.Println(f1);
	for i := 0; i <= l-n; i++ {
		a := f1[i] * (base.Inverse(g[0], p))
		if a < 0 {
			a = p + a
		}
		for j := 0; j < n; j++ {
			f1[j+i] = (f1[j+i] - a*g[j]) % p
			if f1[j+i] < 0 {
				f1[j+i] = p + f1[j+i]
			}
		}
	}
	k := l - 1
	for i := 0; i < l; i++ {
		if f1[i] != 0 {
			k = i
			break
		}
	}
	//log.Println(f1[k:]);
	return f1[k:]
}
