package async

import (
	"errors"
	"log"
	"number/field/base"
	"number/field/matrix"
)

// v1,v2,v3,v4,....u1,u2,ud+1
func S(y []int, p int) ([]int, error) {
	l := len(y) // 2d+1
	if l%2 == 0 {
		return nil, errors.New("Parameter is wrong!")
	}
	d := (l - 1) / 2
	cofs := make([][]int, 0)
	for i := 1; i <= l; i++ {
		cof := make([]int, l)
		if i <= d {
			for j := 0; j < l; j++ {
				cof[j] = base.Pow(i, j, p)
			}
		} else if i == d+1 {
			cof[0] = 1
		} else {
			for j := 0; j <= d; j++ {
				a := i - d - 1
				cof[j] = base.Pow(a, j, p)
			}
		}
		cofs = append(cofs, cof)
	}
	log.Println(cofs)
	res, err := matrix.Solvequals(cofs, y, p)
	if err != nil {
		return nil, err
	}
	return res, nil
}
