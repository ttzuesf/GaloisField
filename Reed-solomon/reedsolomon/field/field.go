package field

var NO_GFLUT bool=false;
var LINEAR_EXP_LUT bool=true;

type Field_symbol int;
const GFERROR Field_symbol =-1;
//
type Field struct{
	power_ uint;
	prim_poly_deg_ uint; //degree of primitive polynomial
	field_size_ uint; // field size by 2^(m);
	prim_poly_hash_ uint;
	prim_poly_ []uint;
	alpha_to_ []Field_symbol  ;  // aka exponential or anti-log
	index_of_ []Field_symbol;    // aka log
	mul_inverse_ []Field_symbol; // multiplicative inverse table
	mul_table_ [][]Field_symbol ; // multiplicative table
	div_table_ [][]Field_symbol;
	exp_table_ [][]Field_symbol;
	linear_exp_table_ [][]Field_symbol;
	buffer_ []byte;

};

// Given a in GF(2^{m}), to get a exponential of a.
func(f *Field)Index(value Field_symbol) Field_symbol{
	return f.index_of_[value];
}

func(f *Field) Alpha(value Field_symbol) Field_symbol {
	return f.alpha_to_[value];
}

func(f *Field) Size() uint {
	return f.field_size_;
}

func(f *Field)Powr() uint{
	return f.power_;
}
// Output field_size
func(f *Field)Mask() uint{
	return f.field_size_
}

func(f *Field)Add(a,b Field_symbol) Field_symbol{
	return a^b;
}

func(f *Field)Sub(a,b Field_symbol) Field_symbol{
	return a^b;
}

// convert value that are not in range a galois field to the galois field;

func(f *Field) normalize(x Field_symbol) Field_symbol {
	for (x < 0){
		x += Field_symbol(f.field_size_);
	}

	for (x >=Field_symbol(f.field_size_)){
		x -= Field_symbol(f.field_size_);
		x  = (x >> f.power_) + (x & Field_symbol(f.field_size_));
	}
	return x;
}


func(f *Field) Mul(a,b Field_symbol) Field_symbol{
	if NO_GFLUT{
		return f.mul_table_[a][b];
	}
	if a==0 || b==0{
		return 0;
	}else{
		return f.alpha_to_[f.normalize(f.index_of_[a]+f.index_of_[b])]; // a^(log(a)+log(b))
	}
}

func(f *Field)Div(a,b Field_symbol) Field_symbol{
	if NO_GFLUT{
		return f.div_table_[a][b];
	}
	if a==0 || b==0{
		return 0;
	}else{
		return f.alpha_to_[f.normalize(f.index_of_[a]-f.index_of_[b]+Field_symbol(f.field_size_))]; // a^(log(a)-log(b))
	}
}


// a^n
func(f *Field) Exp(a Field_symbol, n int) Field_symbol{
	if NO_GFLUT {
		if n >= 0{
			return f.exp_table_[a][n&int(f.field_size_)];
		}else{
			for (n < 0) {
				n += int(f.field_size_);
			};
			if n>0{
				return f.exp_table_[a][n]
			};
			return 1;
		}
	} else{
		if a != 0 {
			if n < 0 {
				for n < 0{
					n += int(f.field_size_);
				}
				if n>0{
					return f.alpha_to_[f.normalize(f.index_of_[a] * Field_symbol(n))];
				};
			} else if n>0{
				return f.alpha_to_[f.normalize(f.index_of_[a] * Field_symbol(n))];
			}
			return 1;
		} else{
			return 0;
		}
	}
}
// if define LINEAR_EXP_LUT, the executing following function.
func(f *Field) Linear_exp(a Field_symbol) []Field_symbol{
	if NO_GFLUT{
		upper_bound:= Field_symbol(2*f.field_size_);
		if a>=0 && a<upper_bound{
			return f.linear_exp_table_[a];
		}else{
			return nil;
		}
	}else{
		return nil;
	}
}
// inverse
func(f *Field)Inverse(value Field_symbol) Field_symbol{
	if NO_GFLUT{
		return f.mul_inverse_[value];
	}else{
		return f.alpha_to_[f.normalize(Field_symbol(f.field_size_)-f.index_of_[value])];
	}
}

func(f*Field)Prim_poly_term(index Field_symbol) uint{
	return f.prim_poly_[index];  //primitive polynomial;
}


// Generate a field based on a primitive polynomial;
func(f *Field)Generate_field(prim_poly []uint){
	var mask Field_symbol=1;
	f.alpha_to_[f.power_]=0;
	for i:=0;i<int(f.power_);i++{
		f.alpha_to_[i]=mask;
		f.index_of_[f.alpha_to_[i]]=Field_symbol(i);
		if prim_poly[i]!=0{
			f.alpha_to_[f.power_]^=mask;
		}
		mask<<=1;
	}
	f.index_of_[f.alpha_to_[f.power_]] = Field_symbol(f.power_);
	mask >>= 1;

	for i := f.power_ + 1; i<f.field_size_; i++{
	if f.alpha_to_[i - 1] >= mask{
		f.alpha_to_[i] = f.alpha_to_[f.power_] ^ ((f.alpha_to_[i - 1] ^ mask) << 1);
	}else{
		f.alpha_to_[i] = f.alpha_to_[i - 1] << 1;
	}
		f.index_of_[f.alpha_to_[i]] = Field_symbol(i);
	}
	f.index_of_[0] = GFERROR;
	f.alpha_to_[f.field_size_] = 1;
	if NO_GFLUT {
		f.mul_table_=f.creat_array(int(f.field_size_+1),int(f.field_size_+1));
		f.div_table_=f.creat_array(int(f.field_size_+1),int(f.field_size_+1));
		f.exp_table_=f.creat_array(int(f.field_size_+1),int(f.field_size_+1));
		for i := 0; i < int(f.field_size_+1); i++ {
			for j := 0; j < int(f.field_size_+1); j++ {
				f.mul_table_[i][j] = f.gen_mul(i, j);
				f.div_table_[i][j] = f.gen_div(i, j);
				f.exp_table_[i][j] = f.gen_exp(i, j);
			}
		}
		// whether define LINEAR_EXP_LUT
		if LINEAR_EXP_LUT {
			f.linear_exp_table_=f.creat_array(int(f.field_size_+1),int(2*f.field_size_));
			f.mul_inverse_=make([]Field_symbol,2*f.field_size_+2);
			for i := 0; i < int(f.field_size_+1); i++ {
				for j := 0; j < int(2*f.field_size_); j++ {
					f.linear_exp_table_[i][j] = f.gen_exp(i, j);
				}
			}
			for i := 0; i < int(f.field_size_+1); i++ {
				f.mul_inverse_[i] = f.gen_inverse(i);
				f.mul_inverse_[i+int(f.field_size_+1)] = f.mul_inverse_[i];
			}
		}
	}
}
// creat a with dimensions array with m rows, in columns
func(f *Field) creat_array(m,n int) [][]Field_symbol{
	res:=make([][]Field_symbol,0);
	for i:=0;i<m;i++{
		res=append(res,make([]Field_symbol,n));
	}
	return res;
}

func(f *Field) gen_mul(a,b int) Field_symbol{
	if a == 0 || b == 0{
		return 0;
	}
	return f.alpha_to_[f.normalize(f.index_of_[a] + f.index_of_[b])];
}

func(f *Field) gen_div(a,b int) Field_symbol{
	if a == 0 || b == 0{
		return 0;
	}
	return f.alpha_to_[f.normalize(f.index_of_[a] - f.index_of_[b] + Field_symbol(f.field_size_))];
}

func(f *Field) gen_exp(a,n int) Field_symbol {
	if a != 0{
		if n==0{
			return 1
		}
	return f.alpha_to_[f.normalize(f.index_of_[a]*Field_symbol(n))];
	}
	return 0;
}

func(f *Field) gen_inverse(i int) Field_symbol{
	return 0;
}

func NewField(pwr uint,primpoly_deg uint, primitive_poly []uint) *Field{
	field:=new(Field);
	field.power_=pwr;
	field.field_size_=1<<field.power_-1; //2^{power_}-1
	field.prim_poly_deg_=primpoly_deg;
	field.alpha_to_=make([]Field_symbol,field.field_size_+1);
	field.index_of_=make([]Field_symbol,field.field_size_+1);
	buffer_size:=3;
	field.buffer_=make([]byte,buffer_size);
	return field;
}