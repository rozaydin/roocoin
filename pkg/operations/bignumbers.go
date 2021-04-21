package operations

import (
	"math/big"
)

type BigNumbers interface {
	Sum(num1, num2 *big.Int) *big.Int
	Sub(num1, num2 *big.Int) *big.Int
	Mul(num1, num2 *big.Int) *big.Int
	Div(num1, num2 *big.Int) *big.Int
	Exp(num1, power *big.Int) *big.Int
	Sqr(num1 *big.Int) *big.Int
	Cube(num1 *big.Int) *big.Int
	IsInf(num1 *big.Int) bool
	MulInverse(num1 *big.Int) *big.Int
}
