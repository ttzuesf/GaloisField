package hash

import (
	"log"
	"math"
)

// Division Method

const n int = 10

const A float64 = 0.357840

func Dhash(a int) int {
	l := 1 << n
	log.Println(l)
	return a % l
}

// Mid Square Method

func Mhash(a int) int {
	r := 2
	l := 1 << n
	return a * a
}

// Digit Folding Method

func DFhash(a []int) int {
	s := 0
	for i := 0; i < len(a); i++ {
		s = a[i] + s
	}
	return s
}

// Multiplication Method

func Mthash(a int) int {
	M := float64(100) // hash table size
	return int(math.Floor(M * (A*a - math.Floor(A*a))))
}
