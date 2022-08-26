package matrix

import (
	"errors"
	"log"
	"number/field/base"
)

// A*B
func Multiply(a, b [][]int, p int) [][]int {
	m1 := len(a)
	n1 := len(a[0])
	m2 := len(b)
	n2 := len(b[0])
	if n1 != m2 {
		log.Println("error")
		return nil
	}
	res := make([][]int, 0)
	for i := 0; i < m1; i++ {
		c := make([]int, n2)
		for j := 0; j < n2; j++ { // the jth row of a
			s := 0
			for k := 0; k < n1; k++ {
				s = (s + a[i][k]*b[k][j]) % p
			}
			c[j] = s
		}
		res = append(res, c)
	}
	return res
}

// k is an integer, the

// det(a)
func Det(a [][]int, p int) int {
	m := len(a)
	n := len(a[0])
	if m != n {
		return 0
	}
	U := make([][]int, 0) //create a L matrix
	for i := 0; i < n; i++ {
		c := make([]int, n)
		copy(c, a[i])
		U = append(U, c)
	}
	for i := 0; i < n-1; i++ {
	Lab:
		for j := i + 1; j < n; j++ {
			if U[i][i] == 0 {
				c := make([]int, n) // create a cache space;
				for j = i + 1; j < n; j++ {
					if U[j][i] != 0 {
						copy(c, U[i])
						copy(U[i], U[j])
						copy(U[j], c)
						break
					}
					if j == n-1 && U[j][i] == 0 {
						break Lab
					}
				}
			}
			a := (U[j][i] * base.Inverse(U[i][i], p)) % p
			for k := i; k < n; k++ { // column traverse
				b := (U[j][k] - a*U[i][k]) % p
				if b < 0 {
					b = b + p
				}
				U[j][k] = b
			}
		}
	}
	res := 1
	for i := 0; i < n; i++ {
		res = (res * U[i][i]) % p
	}
	if res < 0 {
		res = res + p
	}
	return res
}

// LUDecompose: U=A
func LUDecompose(A [][]int, p int) ([][]int, error) {
	l := len(A)
	if len(A[0]) != l {
		return nil, errors.New("rows doesn't equal columns")
	}
	L := make([][]int, 0) //create a L matrix
	for i := 0; i < l; i++ {
		c := make([]int, l)
		L = append(L, c)
		L[i][i] = 1
	}
	for i := 0; i < l-1; i++ { // calculate the L and U matrix
	Lab:
		for j := i + 1; j < l; j++ { // row traverse
			if A[i][i] == 0 { // partial pivot
				for j := i + 1; j < l; j++ {
					if A[j][i] != 0 {
						c := make([]int, l)
						copy(c, A[i])
						copy(A[i], A[j])
						copy(A[j], c)
						break
					}
					if j == l-1 {
						break Lab
					}
				}
			}
			a := (A[j][i] * base.Inverse(A[i][i], p)) % p
			//log.Println(a)
			L[j][i] = a
			//log.Printf("L[%v,%v]=%v\n", j, i, a)
			for k := i; k < l; k++ { // column traverse
				b := (A[j][k] - a*A[i][k]) % p
				if b < 0 {
					b = b + p
				}
				A[j][k] = b
			}
		}
	}
	return L, nil
}

// matrix inv

func Invmatrix(a [][]int, p int) [][]int {
	m := len(a)
	n := len(a[0])
	if m != n || Det(a, p) == 0 {
		return nil
	}
	for i := 0; i < n; i++ {
		b := make([]int, n)
		b[i] = 1
		a[i] = append(a[i], b...)
	}
	M := base.Gaosifp(a, p)
	res := make([][]int, n)
	for i := 0; i < m; i++ {
		b := make([]int, n)
		copy(b, M[i][n:])
		res[i] = b
	}
	return res
}

// solve equations

func Solvequals(cofs [][]int, y []int, p int) ([]int, error) {
	if len(y) != len(cofs) {
		return nil, errors.New("parameters are wrong!")
	}
	if Det(cofs, p) == 0 {
		return nil, errors.New("No deterministic solution!")
	}
	n := len(cofs)
	for i := 0; i < n; i++ {
		cofs[i] = append(cofs[i], y[i])
	}
	mid := base.Gaosifp(cofs, p)
	//log.Println(len(mid),len(mid[0]),n)
	res := make([]int, n)
	for i := 0; i < n; i++ {
		res[i] = mid[i][n]
	}
	return res, nil
}
