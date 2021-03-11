package finitefield

import (
	"testing"
)

func TestFieldElement(t *testing.T) {

	value, order := int64(1), int64(19)

	elem1, err := NewFieldElement(value, order)
	if err != nil {
		t.Error("should create FieldElement from valid values!")
	}

	if elem1.value != value {
		t.Error("should have set value of finitefield properly!")
	}

	if elem1.order != order {
		t.Error("should have set order of finitefield properly")
	}
}

func TestInvalidFieldElement(t *testing.T) {

	value, order := int64(21), int64(19)
	_, err := NewFieldElement(value, order)
	if err == nil {
		t.Error("should raise an error when value is greater/equal than order!")
	}
}

func TestNegativeElement(t *testing.T) {
	value, order := int64(-8), int64(19)
	_, err := NewFieldElement(value, order)
	if err == nil {
		t.Error("should raise an error when value is negative!")
	}
}

func TestAdditiveInverse(t *testing.T) {
	fe, _ := NewFieldElement(9, 19)
	additiveInverse := fe.AdditiveInverse()

	if additiveInverse != 10 {
		t.Error("should properly calculate additive inverses of numbers!")
	}
}

func TestFieldElementEquality(t *testing.T) {

	elem1, _ := NewFieldElement(7, 19)
	elem2, _ := NewFieldElement(7, 19)

	if elem1 != elem2 {
		t.Error("FiniteField Elements with same value should be equal")
	}
}

func TestFiniteFieldSum(t *testing.T) {

	elem1, _ := NewFieldElement(1, 19)
	elem1.SumWith(-19)

	if elem1.value != 1 {
		t.Error("Should properly calculate the modulus with negative numbers")
	}

	elem2, _ := NewFieldElement(11, 19)
	elem2.SumWith(17)

	if elem2.value != 9 {
		t.Error("Should properly calculate the modulus with positive numbers")
	}
}

func TestFiniteFieldSubtraction(t *testing.T) {

	elem1, _ := NewFieldElement(1, 19)
	elem2, _ := NewFieldElement(2, 19)

	elem3, _ := elem1.Subtract(elem2)
	if elem3.value != 18 {
		t.Error("Failed to subtract 2 elements from each other!")
	}

	elem4, _ := NewFieldElement(18, 19)
	elem5, _ := NewFieldElement(17, 19)

	elem6, _ := elem4.Subtract(elem5)
	if elem6.value != 1 {
		t.Error("Failed to subtract 2 elements from each other!")
	}
}

func TestFiniteFieldSubtractWith(t *testing.T) {

	elem1, _ := NewFieldElement(1, 19)
	elem1.SubtractWith(1)

	if elem1.value != 0 {
		t.Error("Failed to subtract value from elem1!")
	}

	elem1.SubtractWith(-19)
	if elem1.value != 0 {
		t.Error("Failed to subtract value from elem1!")
	}

}

func TestFiniteFieldMultiplication(t *testing.T) {

	elem1, _ := NewFieldElement(1, 19)
	elem2, _ := NewFieldElement(18, 19)

	elem3, _ := elem1.Multiply(elem2)
	if elem3.value != 18 {
		t.Error("Failed to multiply finite elements!")
	}

	elem4, _ := NewFieldElement(1, 19)
	elem5, _ := NewFieldElement(1, 18)

	_, err := elem4.Multiply(elem5)
	if err == nil {
		t.Error("Different order multiplication should fail!")
	}

}

func TestFiniteFiledMultiplyWith(t *testing.T) {

	elem1, _ := NewFieldElement(1, 19)
	elem1.MultiplyWith(19)

	if elem1.value != 0 {
		t.Error("Failed to multiply finite elements!")
	}
}

func TestNegativeMultiplyWith(t *testing.T) {
	elem1, _ := NewFieldElement(1, 19)
	elem1.MultiplyWith(-5)

	if elem1.value != 14 {
		t.Error("Negative multiplication do not work!")
	}
}

func TestPower(t *testing.T) {

	elem1, _ := NewFieldElement(2, 19)
	elem1.Power(2)

	if elem1.value != 4 {
		t.Error("Power function is not working!")
	}

	elem1.value = 2
	elem1.Power(4)

	if elem1.value != 16 {
		t.Error("Power function is not working!")
	}

	elem1.value = 18
	elem1.Power(2)

	if elem1.value != 1 {
		t.Error("Power function is not working!")
	}
}

func TestDivision(t *testing.T) {

	elem1, _ := NewFieldElement(1, 3)
	elem1.DivideBy(2)

	if elem1.value != 2 {
		t.Error("Division between FiniteElements not working!")
	}

}
