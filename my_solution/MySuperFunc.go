package titov_solution

import (
	//"fmt"
	//"math"	
	"math/big"
)

type d struct{ p1,p0 big.Int }

var bigTwo = big.NewInt(2)
var bigZero = big.NewInt(0)

func MySuperFuncImpl(x1 float64, x2 float64, n uint8) float64 {
	var p = cf(0,1,n)
	return pw(x1,&p.p1) * bpw(x2,&p.p0)
}

func cf(s0 int64, s1 int64,goal uint8) d {
	var buf int64
	var i uint8 = 1
	for ; i <= goal; i++ {
		buf = s1
		s1 = s1 + s0
		s0 = buf
	}
	return d{p1: *big.NewInt(s1), p0: *big.NewInt(s0)}
}

func pw(x float64, y *big.Int) float64 {
	var tmp = y.Int64()
	var result float64 = 1
	for ; tmp > 0; tmp = tmp >> 1 {
		//предположительно можно избавиться от проверки условия через побитовые операции над x
		if tmp&1 == 1 { result = result * x }
		x = x * x
	}
	return result
}

func bpw(x float64, y *big.Int) float64 {
	var result float64 = 1
	for ; y.Cmp(bigZero) > 0; y.Div(y, bigTwo) {
		if y.Bit(0) == 1 { result = result * x }
		x = x * x
	}
	return result
}
