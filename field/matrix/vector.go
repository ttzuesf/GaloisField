package matrix

// f=vec1+vec2
func Add(vec1,vec2[]float64) []float64{
	l:=len(vec1)
	res:=make([]float64,0)
	if len(vec2)>l{
		res=append(res,vec2...)
	}else{
		l=len(vec2)
		res=append(res,vec1...)
	}
	for i:=0;i<l;i++{
		res[i]=vec1[i]+vec2[i]
	}
	return res
}

// res=a*vec1 a is a number

func Numtimes(a interface{}, vec interface{}) interface{}{
	switch vec.(type) {
	case []float64:
		if _,ok:=a.(float64);ok!=true{
			return nil;
		}
	case []int64:
		if _,ok:=a.(int64);ok!=true{
			return nil;
		}
	default:
		return nil;
	}
	val:=vec.([]float64)
	b:=a.(float64)
	for i:=0;i<len(val);i++{
		val[i]=b*val[i];
	}
	return val;
}

// res=vec1*vec2'
func Vectimes(vec1 []float64, vec2 []float64) []float64{
	if len(vec1)!=len(vec2){
		return nil;
	}
	return nil;
}