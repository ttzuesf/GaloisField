package math

import (
	"math"
)

func caculate(lambda,k float64) float64{
	fenzi:=lambda*math.Log2(lambda)+0.53*lambda+k;
	fenmu:=k-lambda;
	return fenzi/fenmu
}
