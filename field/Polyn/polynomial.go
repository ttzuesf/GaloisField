package Polyn

import (
	"number/field/matrix"
)

func F1(c []float64, x float64) float64 {
	if len(c) == 0 {
		return 0
	}
	s := c[0]
	for i := 0; i < len(c)-1; i++ {
		s = x*s + c[i+1]
	}
	return s
}

func F2(c []float64, x float64) float64 {
	if len(c) == 0 {
		return 0
	}
	r := c[0]
	for i := 0; i < len(c)-1; i++ {
		r = c[i+1] + x*r
	}
	return r
}

func NewTon(a, b []float64) []float64 {
	x := make([]float64, len(a))
	copy(x, a)
	y := make([]float64, len(b))
	copy(y, b)
	if len(x) != len(y) {
		return nil
	}

	n := len(x)
	f := make([]float64, 0)
	f = append(f, y[0])
	for i := 0; i < n-1; i++ {
		for j := 0; j < n-1-i; j++ {
			y[j] = (y[j] - y[j+1]) / (x[j] - x[j+i+1])
		}
		f1 := []float64{y[0]}
		for j := 0; j <= i; j++ {
			cof := []float64{1}
			cof = append(cof, -x[j]) //x-x_{j}
			f1 = mutipolyf(f1, cof)
		}
		l1 := len(f)
		l2 := len(f1)
		zeros := make([]float64, l2-l1)
		f = append(zeros, f...)
		f = matrix.Add(f, f1)
	}
	return f
}

// f*p
func mutipolyf(cof1, cof2 []float64) []float64 {
	res := make([]float64, len(cof1)+len(cof2)-1)
	for i := 0; i < len(cof1); i++ {
		for j := 0; j < len(cof2); j++ {
			res[i+j] = res[i+j] + cof1[i]*cof2[j]
		}
	}
	return res
}
