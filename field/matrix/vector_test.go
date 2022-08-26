package matrix

import (
	"log"
	"testing"
)

func TestAdd(t *testing.T) {
	a:=[]float64{1,2,7}
	b:=[]float64{3,4,}
	log.Println(Add(a,b))
	log.Println(Add(b,a))
}