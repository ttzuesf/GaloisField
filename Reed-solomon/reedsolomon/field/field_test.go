package field

import (
	"log"
	"testing"
	"unsafe"
)

func TestNewfield(t *testing.T){
	prim_poly:=Primitive_polynomial00;
	power:=uint(3);
	field:=NewField(power,3,prim_poly);
	// testing Add opertation;
	field.Generate_field(prim_poly);
	log.Println("power",field.power_);
	log.Println("alpher_to_i",field.alpha_to_);
	log.Println("index of alpha",field.index_of_);
	//var a,b Field_symbol=3,4;
	var a,b Field_symbol=3,4;
	//testing adding operation;
	log.Printf("%v+%v=%v\n",a,b,field.Add(a,b));
	//testing sub operation;
	log.Printf("%v-%v=%v\n",a,b,field.Sub(a,b));
	//testing multiplication ;
	log.Printf("%v*%v=%v\n",a,b,field.Mul(a,b));
	//testing division;
	log.Printf("%v/%v=%v\n",a,b,field.Div(a,b));
	//testing exponential operator;
	log.Printf("%v^%v=%v\n",a,2,field.Exp(a,10));
	field.mul_table_=field.creat_array(3,4)
	//testing creat a 2 dimensions array
	log.Println(field.mul_table_);
}

func TestFieldcache(t *testing.T){
	NO_GFLUT=true;
	//LINEAR_EXP_LUT=false;
	prim_poly:=Primitive_polynomial00;
	power:=uint(3);
	field:=NewField(power,3,prim_poly);
	field.Generate_field(prim_poly);
	//print mul_table;
	log.Printf("mul_table:%v\n",field.mul_table_);
	//print div_table;
	log.Printf("mul_table:%v\n",field.div_table_);
	//print exp_table;
	log.Printf("exp_table:%v\n",field.exp_table_)
	// print memory of field
	log.Printf("Memory:%v\n",unsafe.Sizeof(field.exp_table_));
}

func TestShift(t *testing.T){
	a:=1024;
	log.Printf("%d\n",a>>10);
}