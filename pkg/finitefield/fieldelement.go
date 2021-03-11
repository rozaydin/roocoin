package finitefield

import (
	"fmt"
)

// FieldElement ...
type FieldElement struct {
	value int64
	order int64
}

// NewFieldElement ...
func NewFieldElement(value int64, order int64) (FieldElement, error) {

	var element = FieldElement{}

	if value >= order {
		err := fmt.Errorf(fmt.Sprintf("value: %d is >= order: %d", value, order))
		return element, err
	}

	if value < 0 {
		err := fmt.Errorf(fmt.Sprintf("value: %d is < 0 ", value))
		return element, err
	}

	element.value = value
	element.order = order

	return element, nil
}

// CalculateFromMultiplication ...
func CalculateFromMultiplication(multipliers []int64, order int64) [][]int64 {

	multipliersLen := len(multipliers)
	values := make([][]int64, multipliersLen)

	for index := range values {
		multiplier := multipliers[index]
		for i := int64(0); i < order; i++ {
			values[index] = append(values[index], finiteFieldMultiplication(i, multiplier, order))
		}
	}

	return values
}

// CalculateFromExponent ...
func CalculateFromExponent(exponents []int64, order int64) [][]int64 {

	multipliersLen := len(exponents)
	values := make([][]int64, multipliersLen)

	for index := range values {
		exponent := exponents[index]
		for i := int64(0); i < order; i++ {
			values[index] = append(values[index], finiteFieldExponent(i, exponent, order))
		}
	}

	return values

}

// Sum sums 2 FieldElements and returns new FieldElement without
// changing the original values
func (fe FieldElement) Sum(other FieldElement) (FieldElement, error) {

	var fieldElement = FieldElement{}
	if fe.order != other.order {
		err := fmt.Errorf("order of elements does not match")
		return fieldElement, err
	}

	fieldElement.order = fe.order
	fieldElement.value = finiteFieldSum(fe.value, other.value, fe.order)
	return fieldElement, nil
}

// SumWith Adds a number to the sum and updates the value
func (fe *FieldElement) SumWith(number int64) {
	fe.value = finiteFieldSum(fe.value, number, fe.order)
}

// AdditiveInverse calculates the additive inserve
func (fe *FieldElement) AdditiveInverse() int64 {
	return additiveInverse(fe.value, fe.order)
}

// Subtract subtracts two FieldElements from each other
func (fe FieldElement) Subtract(other FieldElement) (FieldElement, error) {
	negatedOther := other
	negatedOther.value = negatedOther.value * -1
	return fe.Sum(negatedOther)
}

// SubtractWith subtracts a number from FieldElement
func (fe *FieldElement) SubtractWith(number int64) {
	fe.SumWith(-1 * number)
}

// Multiply multiplies two elements and returns a result
func (fe *FieldElement) Multiply(other FieldElement) (FieldElement, error) {

	var fieldElement = FieldElement{}
	if fe.order != other.order {
		err := fmt.Errorf("order of elements does not match")
		return fieldElement, err
	}

	fieldElement.order = fe.order
	fieldElement.value = finiteFieldMultiplication(fe.value, other.value, fe.order)
	return fieldElement, nil
}

// MultiplyWith multiples the field element with some value
func (fe *FieldElement) MultiplyWith(number int64) {
	fe.value = finiteFieldMultiplication(fe.value, number, fe.order)
}

// Divide divides 2 elements and returns result
func (fe *FieldElement) Divide(other FieldElement) (FieldElement, error) {

	var fieldElement = FieldElement{}
	if fe.order != other.order {
		err := fmt.Errorf("order of elements does not match")
		return fieldElement, err
	}

	fieldElement.order = fe.order
	fieldElement.value = finiteFieldDivision(fe.value, other.value, fe.order)
	return fieldElement, nil
}

// DivideBy divides field element with provided number
func (fe *FieldElement) DivideBy(number int64) {
	fe.value = finiteFieldDivision(fe.value, number, fe.order)
}

// Power takes exponent of FiniteField element
func (fe *FieldElement) Power(power int64) {
	fe.value = finiteFieldExponent(fe.value, power, fe.order)
}

func (fe FieldElement) String() string {
	return fmt.Sprintf("value: %d, order: %d", fe.value, fe.order)
}

func modulus(value int64, order int64) int64 {
	mod := value % order
	if mod < 0 {
		// negative number modulus
		mod = mod + order
	}
	return mod
}

func finiteFieldSum(num1 int64, num2 int64, order int64) int64 {
	sum := (num1 + num2)
	return modulus(sum, order)
}

func finiteFieldMultiplication(num1 int64, num2 int64, order int64) int64 {
	multiplication := num1 * num2
	return modulus(multiplication, order)
}

func finiteFieldExponent(num1 int64, exponent int64, order int64) int64 {
	var value int64 = 1
	reducedExponent := modulus(exponent, order-1)
	for i := int64(0); i < reducedExponent; i++ {
		value = value * num1
	}

	return modulus(value, order)
}

func finiteFieldDivision(num1 int64, num2 int64, order int64) int64 {
	// p-2 is the exponent
	num2Inverse := finiteFieldExponent(num2, (order - 2), order)
	return finiteFieldMultiplication(num1, num2Inverse, order)
}

func additiveInverse(num1 int64, order int64) int64 {
	return modulus((-1 * num1), order)
}
