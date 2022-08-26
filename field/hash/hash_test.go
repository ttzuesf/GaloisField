package hash

import(
	"testing"
	"log"
)

func TestDHash(t *testing.T){
	a:=5;
	log.Printf("r=%+v\n",Dhash(a));
}