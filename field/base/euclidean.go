package base

// solve gcd(a,b)
func Euclidean(a,b int) int{
	if a<b{
		t:=a;
		a=b;
		a=t;
	}
	for b>0{
		r:=a%b;
		a=b;
		b=r;
	}
	return a;
}

// solve ax+by=gcd(a,b)
func ExtendEuclidean(a,b int) (int,int){
	if a<b{
		t:=a;
		a=b;
		a=t;
	}
	s0:=1;
	t0:=0;
	s1:=0;
	t1:=1;
	for b>0{
		q:=a/b;
		s2:=s0-q*s1;
		t2:=t0-q*t1;
		s0=s1;
		s1=s2;
		t0=t1;
		t1=t2;
		r:=a%b;
		a=b;
		b=r;
	}
	return s0,t0;
}