package base

import "log"

// c=a^x%p
func Pow(a, x, p int) int {
	c := 1
	for x > 0 {
		if x%2 != 0 {
			c = (c * a)%p
		}
		x = x >> 1
		a = (a * a) % p
	}
	return c
}

// c=a^{-1}%p
func Inverse(a, p int) int {
	x, _ := ExtendEuclidean(a, p)
	if x < 0 {
		x = p + x
	}
	return x
}

// primitive elemenet of Z_{p}

func PrimElement(a, p int) bool {
	s := 1
	for i := 1; i < p-1; i++ {
		s = s * a % p
		if s == 1 {
			return false
		}
	}
	log.Println(s)
	return true
}

// Gaosi elimination over fp
func Gaosifp(M [][]int, p int) [][]int {
	m := len(M)
	n := len(M[0])
	if m > n {
		return nil
	}
	for j := 0; j < m; j++ {
		for s := 0; M[j][j] == 0 && j < m-1; s++ {
			if j+s < m && M[j][j+s] != 0 {
				b := make([]int, n)
				copy(b, M[j])
				copy(M[j], M[j+s])
				copy(M[j+s], b)
				break
			}
		}
		if M[j][j] == 0 {
			continue
		}
		a := Inverse(M[j][j], p)
		for k := j; k < n; k++ {
			M[j][k] = (a * M[j][k]) % p
		}
		//log.Println(M[j]);
		for i := 0; i < m; i++ {
			if i == j {
				continue
			}
			b := M[i][j] //element with the same column of a
			//log.Println("b:",b)
			for k := 0; k < n; k++ {
				M[i][k] = (M[i][k] - b*M[j][k]) % p
				if M[i][k] < 0 {
					M[i][k] = p + M[i][k]
				}
			}
		}
		//log.Println(M)
	}
	return M
}

//quadratic residue

func QS(a, p int) bool {
	b := (p - 1) / 2
	if 1 == Pow(a, b, p) {
		return true
	}
	return false
}

func SolveQS(a, v, p int) int {
	if !QS(v, p) == true {
		return -1
	}
	for i := 0; i <= (p-1)/2; i++ {
		if Pow(a, 2*i, p) == v {
			return Pow(a, i, p)
		}
	}
	return -1
}
