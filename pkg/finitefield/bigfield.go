package finitefield

import (
	"math/big"
)

var (
	// TODO for big.Int i failed to find a build in
	// infinity checker so relying on the finitefield
	// wont permit negative values, we will rely on
	// -1 value to represent infinity
	infinity = big.NewInt(-1)
)

type BigFiniteField struct {
	order *big.Int
}

// NewBigFiniteField creates a finite field
func NewBigFiniteField(order *big.Int) BigFiniteField {
	return BigFiniteField{
		order: new(big.Int).Set(order),
	}
}

// Sum ...
func (ff BigFiniteField) Sum(elem1, elem2 *big.Int) *big.Int {

	result := big.NewInt(0)
	result.Add(elem1, elem2)

	return ff.mod(result)
}

// Subtract ...
func (ff BigFiniteField) Sub(elem1, elem2 *big.Int) *big.Int {

	elem2Neg := big.NewInt(0)
	elem2Neg = elem2Neg.Neg(elem2)

	return ff.Sum(elem1, elem2Neg)
}

// Multiply ...
func (ff BigFiniteField) Mul(elem1, elem2 *big.Int) *big.Int {

	result := big.NewInt(0)
	result = result.Mul(elem1, elem2)

	return ff.mod(result)
}

// Divide ...
func (ff BigFiniteField) Div(elem1, elem2 *big.Int) *big.Int {

	orderMinusTwo := big.NewInt(0)
	orderMinusTwo.Sub(ff.order, big.NewInt(2))

	elem2Inverse := ff.Exp(elem2, orderMinusTwo)
	return ff.Mul(elem1, elem2Inverse)
}

// Exp ...
func (ff BigFiniteField) Exp(elem1, power *big.Int) *big.Int {

	orderMinusOne := new(big.Int)
	orderMinusOne = orderMinusOne.Sub(ff.order, big.NewInt(1))

	reducedExponent := BigMod(power, orderMinusOne)
	value := big.NewInt(1)

	return value.Exp(elem1, reducedExponent, ff.order)
}

func (ff BigFiniteField) Sqr(num1 *big.Int) *big.Int {
	return ff.Exp(num1, big.NewInt(2))
}

func (ff BigFiniteField) Cube(num1 *big.Int) *big.Int {
	return ff.Exp(num1, big.NewInt(3))
}

func (ff BigFiniteField) IsInf(num1 *big.Int) bool {
	return num1.Cmp(infinity) == 0
}

func (ff BigFiniteField) MulInverse(num1 *big.Int) *big.Int {
	orderMinusTwo := big.NewInt(0)
	orderMinusTwo.Sub(ff.order, big.NewInt(2))
	return ff.Exp(num1, orderMinusTwo)
}

func (ff *BigFiniteField) mod(value *big.Int) *big.Int {
	return BigMod(value, ff.order)
}

func BigMod(value *big.Int, order *big.Int) *big.Int {
	var mod = big.NewInt(0)
	return mod.Mod(value, order)
}
