package matrix

import (
	"log"
	"testing"
)

func TestMultiply(t *testing.T) {
	a := make([][]int, 0)
	a = append(a, []int{1, 2}, []int{2, 2})
	b := make([][]int, 0)
	b = append(b, []int{10, 1}, []int{1, 5})
	log.Println(a, b)
	res := Multiply(a, b, 11)
	log.Println(res)
}

func TestDet(t *testing.T) {
	M := make([][]int, 0)
	M = append(M, []int{-2, 3, 1}, []int{0, -7, 2}, []int{0, 6, 6})
	p := 11
	log.Println(Det(M, p))
}

// test LUDecompose
func TestLUDecompose(t *testing.T) {
	M := make([][]int, 0)
	M = append(M, []int{-2, 3, 1}, []int{0, -7, 2}, []int{0, 6, 6})
	p := 11
	L, _ := LUDecompose(M, p)
	log.Println(L, M)
	M1 := Multiply(L, M, p)
	log.Println(M1)
}

// teset Inversion of Matrix
func TestInvmatrix(t *testing.T) {
	m := make([][]int, 0)
	m = append(m, []int{1, 2}, []int{2, 2})
	c := Invmatrix(m, 11)
	log.Println(m)
	log.Println(c)
}

func TestSolvequals(t *testing.T) {
	m := make([][]int, 0)
	m = append(m, []int{3, 5, 4}, []int{5, 4, 9}, []int{4, 9, 7})
	y := []int{2, 4, 0}
	x, _ := Solvequals(m, y, 11)
	log.Println("x:", x)
}
