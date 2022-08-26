package base

func Cfraction(a []float64) float64 {
	var s float64 = a[len(a)-1]
	for i := len(a) - 1; i >= 0; i-- {
		s = a[i] + 1/s
	}
	return s
}
