package async

import (
	"log"
	"testing"
)

func TestS(t *testing.T) {
	ys := []int{1, 7, 10, 3, 2, 0, 4}
	p := 11
	log.Printf("v1,v2,v3,....,u1,u2,...,ud:%d\n", ys)
	res, err := S(ys, p)
	if err != nil {
		return
	}
	log.Printf("res:%v\n", res)
}
