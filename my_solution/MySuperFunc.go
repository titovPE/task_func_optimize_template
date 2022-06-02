package titov_solution

import (
	//"fmt"
	//"math"	
	//"math/big"
)

type d struct{ p1,p0 int64 }

func MySuperFuncImpl(x1 float64, x2 float64, n uint8) float64 {
	var p = cf(0,1,n)
	return pw(x1,p.p1) * pw(x2,p.p0)
}

func cf(s0 int64, s1 int64,goal uint8) d {
	var buf int64
	var i uint8 = 1
	for ; i <= goal; i++ {
		buf = s1
		s1 = s1 + s0
		s0 = buf
	}
	return d{p1: s1, p0: s0}
}

func pw(x float64, y int64) float64 {
	var result float64 = 1
	for ; y > 0; y = y >> 1 {
		if y & 1 == 1 { result = result * x }
		//var factor = float64(1&^(1&^(y & 1)))
		//result = result*(1*factor+x*factor)
		x = x * x
	}
	return result
}
