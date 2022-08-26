package field

import (
	"errors"
	"fmt"
)

type Field_element struct{
	field_ *Field;
	poly_value Field_symbol;
}

func(fe *Field_element)Printf(gfe Field_element){
	fmt.Println(gfe.poly_value);
}
// a+b
func(fe *Field_element)Add(a,b *Field_element) *Field_element{
	if fe==nil{
		panic(errors.New("The pointer Field element is nil"));
	}
	fe.poly_value =a.poly_value^b.poly_value;
	return fe;
}

//a-b;
func(fe *Field_element)Sub(a,b*Field_element) *Field_element{
	if fe==nil{
		panic(errors.New("The pointer Field element is nil"));
	}
	fe.poly_value=a.poly_value^b.poly_value;
	return fe;
}

//a*b;
func(fe *Field_element)Mul(a,b*Field_element) *Field_element{
	if fe==nil{
		panic(errors.New("The pointer Field element is nil"));
	}
	fe.poly_value=fe.field_.Mul(a.poly_value,b.poly_value)
	return fe;
}

//a/b
func(fe *Field_element)Div(a,b*Field_element)*Field_element{
	if fe==nil{
		panic(errors.New("The pointer Field element is nil"));
	}
	fe.poly_value=fe.field_.Div(a.poly_value,b.poly_value)
	return fe;

}
//a^b
func(fe *Field_element)Exp(a *Field_element,n int)*Field_element{
	if fe==nil{
		panic(errors.New("The pointer Field element is nil"));
	}
	fe.poly_value=fe.field_.Exp(a.poly_value,n)
	return fe;

}

//if a>b, return 1, a<b return -1, a==b return 0;
func(fe *Field_element)Cmp(a,b*Field_element) int{
	if fe==nil{
		panic(errors.New("The pointer Field element is nil"));
	}
	if a.poly_value>b.poly_value{
		return 1;
	}else if a.poly_value<b.poly_value{
		return -1;
	}else{
		return 0;
	}
}

func(fe *Field_element)Set(val int){
	if fe==nil{
		panic(errors.New("The pointer Field element is nil"));
	}
	fe.poly_value=fe.field_.normalize(Field_symbol(val));
}