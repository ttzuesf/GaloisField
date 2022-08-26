package bloomfilter

import (
	"github.com/xuri/excelize/v2"
	"log"
	"math"
	"math/big"
	"number/hash/file"
	"strconv"
	"testing"
)

func TestFProTrial(t *testing.T) {
	lamb:=6;
	n:=60;
	l:=int(math.Ceil(float64(lamb*n)/math.Log(2)));
	//l:=lamb*n+1;
	var opt excelize.Options;
	file:=file.NewExcel("record",opt)
	for j:=0;j<50;j++{
		sum:=0;
		vect:=FProTrial(lamb,n,l)
		for i:=0;i<10000;i++{
			if Fverify(lamb,l,vect)==true{
				sum=sum+1;
			}
		}
		a:=float64(sum)/float64(10000);
		file.Record("Sheet3","G"+strconv.Itoa(j+1),a)
	}
	file.Close()
}

func TestFverify(t *testing.T) {
	lamb:=2;
	n:=2;
	l:=int(math.Ceil(float64(lamb*n)/math.Log(2)));
	vect:=FProTrial(lamb,n,l)
	log.Println(vect)
	log.Println(Fverify(lamb,l,vect));
}

func TestRand(t *testing.T){
	lamb:=2;
	n:=4;
	l:=int(math.Ceil(float64(lamb*n)/math.Log(2)));
	//l:=lamb*n+1;
	l1:=new(big.Int);
	l1.SetInt64(int64(l))
	//aver:=float64(0);
	var opt excelize.Options;
	file:=file.NewExcel("record",opt)
/*	if err:=file.Newsheet("Sheet3");err!=nil {
		panic(err)
	}
 */
	for j:=0;j<50;j++{
		sum:=0;
		index:=make([]int,lamb)
		for i:=0;i<len(index);i++{
			index[i]=int(urand(l1));
		}
		for i:=0;i<10000;i++{
			if ProTrial(index,lamb,n,l)==true{
				sum=sum+1;
			}
		}
		a:=float64(sum)/float64(10000);
		file.Record("Sheet3","C"+strconv.Itoa(j+1),a)
	}
	file.Close();

}



func TestPotrial(t *testing.T){
	lamb:=2;
	n:=2;
	//l:=int(math.Ceil(float64(lamb*n)/math.Log(2)));
	l:=lamb*n;
	l1:=new(big.Int);
	l1.SetInt64(int64(l))
	index:=make([]int,lamb)
	for i:=0;i<len(index);i++{
		index[i]=int(urand(l1));
	}
	log.Println("index",index)
	log.Println(ProTrial(index,lamb,n,l))

}

func TestCalPro(t *testing.T) {
	p:=CalProbability(6,4)
	log.Println("probability",p)
}