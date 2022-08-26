package congruence

import (
	"log"
	"testing"
)

func TestSolvcongr(t *testing.T) {
	a := 4
	b := 8
	m := 12
	log.Printf("congruence solution: %v\n", Solvcongr(a, b, m))
	log.Printf("all congruence solution: %+v\n", Solvcongrall(a, b, m))
}
