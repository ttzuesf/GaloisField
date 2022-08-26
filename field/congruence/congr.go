package congruence

import (
	"log"
	"number/field/base"
)

// solve the congruence ax=b mod m

func Solvcongr(a, b, m int) int {
	gcd := base.Euclidean(a, m)
	if b%gcd != 0 {
		log.Fatalf("The congruence:%vx=%v mod %v has no solution!", a, b, m)
	}
	b1 := b / gcd
	s, _ := base.ExtendEuclidean(a, m)
	return s * b1
}

// solve all solution satisfying ax=b mod m
func Solvcongrall(a, b, m int) []int {
	gcd := base.Euclidean(a, m)
	if b%gcd != 0 {
		log.Fatalf("The congruence:%vx=%v mod %v has no solution!", a, b, m)
	}
	b1 := b / gcd
	s, _ := base.ExtendEuclidean(a, m)
	log.Printf("s=%v\n", s)
	x0 := b1 * s % m
	m1 := m / gcd
	re := make([]int, gcd)
	for i := 0; i < gcd; i++ {
		re[i] = x0 + i*m1%m
		if re[i] < 0 {
			re[i] = re[i] + m
		}
	}
	return re
}
