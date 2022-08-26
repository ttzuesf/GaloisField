package base

import (
	"log"
	"testing"
)

func TestInverse(t *testing.T) {
	log.Println(Inverse(8, 11))
}

func TestMatrix(t *testing.T) {
	var M [][]int
	M = append(M, []int{4, 6, 0, 1})
	M = append(M, []int{0, 4, 6, 0})
	M = append(M, []int{0, 0, 4, -6})
	log.Println(M[1][2], M[0][1])
	M1 := Gaosifp(M, 11)
	log.Println(M1)
}

func TestPrimElement(t *testing.T) {
	var res []int
	p := 11
	for i := 2; i < p-1; i++ {
		if PrimElement(i, p) {
			res = append(res, i)
		}
	}
	log.Println(res)
}

func TestQS(t *testing.T) {
	p := 11
	log.Println(Pow(5, 5, p))
	log.Println(SolveQS(2, 5, p))
}

func TestPow(t *testing.T) {
	p := 11
	a := 4
	r := Pow(a, 4, p)
	log.Printf("result:=%v\n", r)
}
