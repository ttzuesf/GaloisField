package field

import (
	"log"
	"testing"
)

func TestNewFieldElement(t *testing.T){
	prim_poly:=Primitive_polynomial00;
	power:=uint(3);
	field:=NewField(power,3,prim_poly);
	var a =Field_element{field_: field, poly_value: 12};
	var b =Field_element{field_:field,poly_value: 10};
	var c =Field_element{field_:field,poly_value: 0};
	c.Set(15);
	c.Printf(c);
	log.Println(c.Add(&a,&b).poly_value);
}
