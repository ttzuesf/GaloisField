package field

import (
	"testing"
)
func Sum[T int64|float64](x T){
	fmt.Println(x);
}

func TestGenerics(t *testing.T){
	Sum[int64](16);
}