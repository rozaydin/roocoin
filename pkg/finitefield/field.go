package finitefield

import (
	"math"
)

type FiniteField struct {
	order int64
}

// NewFiniteField creates a finite field
// of order
func NewFiniteField(order int64) FiniteField {
	return FiniteField{order}
}

// Sum ...
func (ff FiniteField) Sum(elem1 float64, elem2 float64) float64 {
	result := (elem1 + elem2)
	return ff.mod(result)
}

// Subtract ...
func (ff FiniteField) Subtract(elem1 float64, elem2 float64) float64 {
	return ff.Sum(elem1, -1*elem2)
}

// Multiply ...
func (ff FiniteField) Multiply(elem1 float64, elem2 float64) float64 {
	result := elem1 * elem2
	return ff.mod(result)
}

// Divide ...
func (ff FiniteField) Divide(elem1 float64, elem2 float64) float64 {
	elem2Inverse := ff.Power(elem2, ff.order-2)
	return ff.Multiply(elem1, elem2Inverse)
}

// Power ...
// TODO refactor this
func (ff FiniteField) Power(elem1 float64, power int64) float64 {
	var value float64 = 1
	reducedExponent := mod(float64(power), float64(ff.order-1))
	for i := int64(0); i < int64(reducedExponent); i++ {
		value = value * elem1
	}

	return ff.mod(value)
}

func (ff *FiniteField) mod(value float64) float64 {
	return mod(value, float64(ff.order))
}

func mod(value float64, order float64) float64 {
	mod := math.Mod(value, order)
	if mod < 0 {
		// negative number modulus
		mod = mod + order
	}
	return mod
}
