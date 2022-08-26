package Polyn

import (
	"log"
	"testing"
)

func TestQuotPQnomial(t *testing.T) {
	P := []int{3, 1, 1}
	Q := []int{4, 0, 3}
	p := 11
	res := make([]int, 0)
	for x := 1; x < p; x++ {
		r := QuotPQnomial(P, Q, x, p)
		res = append(res, r)
	}
	log.Printf("res=[%v]\n", res)
}
