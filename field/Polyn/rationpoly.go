package Polyn

import "number/field/base"

//P(x)/Q(x)

func QuotPQnomial(P []int, Q []int, x int, p int) int {
	a := F(P, x, p)
	b := F(Q, x, p)
	return a * base.Inverse(b, p) % p
}
