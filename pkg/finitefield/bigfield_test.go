package finitefield

import (
	"fmt"
	"math"
	"math/big"
	"testing"
)

// This test checks how modulus works with big numbers
// this suite is mostly getting accustomed to math.big package
func TestPositiveModulus(t *testing.T) {

	x := big.NewInt(10)
	y := big.NewInt(19)
	z := big.NewInt(0)

	// Positive Mod
	z = z.Mod(x, y)

	if z.Int64() != 10 {
		t.Error("calculated value does not match expected!")
	}

}

func TestNegativeModulus(t *testing.T) {

	x := big.NewInt(-10)
	y := big.NewInt(19)
	z := big.NewInt(0)

	// Positive Mod
	z = z.Mod(x, y)

	if z.Int64() != 9 {
		t.Error("calculated value does not match expected!")
	}

}

func TestInfinityCheck(t *testing.T) {

	inf := big.NewFloat(math.Inf(0))
	if !inf.IsInf() {
		t.Error("math.inf is not recognized as inf in big.Float")
	}
}

// func TestIntInfinityCheck(t *testing.T) {

// 	inf := big.NewFloat(math.Inf(0))

// 	infAsInt := big.NewInt(0)
// 	infAsInt.SetString(inf.String(), 10)

// 	convertedInf := big.NewFloat(0)
// 	convertedInf.SetInt(infAsInt)

// 	if !convertedInf.IsInf() {
// 		t.Error("converted int is not recognized as inf in big.Float")
// 	}
// }

func TestSum(t *testing.T) {
	sum(-10, 29, 19, 0, t)
	sum(11, 19, 19, 11, t)
	sum(11, 17, 19, 9, t)
}

func TestSub(t *testing.T) {
	sub(0, 18, 19, 1, t)
	sub(19, 1, 19, 18, t)
	sub(191, 1, 19, 0, t)
	sub(-10, -5, 19, 14, t)
}

func TestDiv(t *testing.T) {
	div(1, 2, 3, 2, t)
}

func TestMul(t *testing.T) {
	multiply(1, 18, 19, 18, t)
	multiply(19, 10, 19, 0, t)
	multiply(-1, 19, 19, 0, t)
}

func TestExp(t *testing.T) {
	exp(2, 4, 19, 16, t)
	exp(2, 0, 19, 1, t)
	exp(18, 2, 19, 1, t)
}

func TestSqr(t *testing.T) {
	sqr(2, 19, 4, t)
	sqr(4, 19, 16, t)
	sqr(8, 19, 7, t)
}

func TestCube(t *testing.T) {
	cube(3, 19, 8, t)
	cube(4, 19, 7, t)
	cube(5, 19, 11, t)
}

func TestInfinity(t *testing.T) {

	bigField := NewBigFiniteField(
		big.NewInt(19),
	)

	if !bigField.IsInf(big.NewInt(-1)) {
		t.Error("Infinity check failed, -1 value should be equal to infinity")
	}
}

// Helper Functions, for easy testing

func operation(x, y, order, expected int64, opType string, t *testing.T) {

	xBig := big.NewInt(x)
	yBig := big.NewInt(y)

	bigField := NewBigFiniteField(
		big.NewInt(order),
	)

	result := big.NewInt(0)

	switch opType {

	case "SUM":
		result = bigField.Sum(xBig, yBig)
	case "SUB":
		result = bigField.Sub(xBig, yBig)
	case "DIV":
		result = bigField.Div(xBig, yBig)
	case "MUL":
		result = bigField.Mul(xBig, yBig)
	case "EXP":
		result = bigField.Exp(xBig, yBig)
	default:
		t.Error(fmt.Sprintf("Unknown opType: %s is provided", opType))
	}

	if result.Int64() != expected {
		t.Error(fmt.Sprintf("Calculated %s: %s does not match expected sum: %d", opType, result.Text(10), expected))
	}
}

func sum(x, y, order, expected int64, t *testing.T) {
	operation(x, y, order, expected, "SUM", t)
}

func sub(x, y, order, expected int64, t *testing.T) {
	operation(x, y, order, expected, "SUB", t)
}

func div(x, y, order, expected int64, t *testing.T) {
	operation(x, y, order, expected, "DIV", t)
}

func multiply(x, y, order, expected int64, t *testing.T) {
	operation(x, y, order, expected, "MUL", t)
}

func exp(x, y, order, expected int64, t *testing.T) {
	operation(x, y, order, expected, "EXP", t)
}

func sqr(x, order, expected int64, t *testing.T) {

	xBig := big.NewInt(x)

	bigField := NewBigFiniteField(
		big.NewInt(order),
	)

	result := bigField.Sqr(xBig)

	if result.Int64() != expected {
		t.Error(fmt.Sprintf("Calculated SQR: %s does not match expected sum: %d", result.Text(10), expected))
	}
}

func cube(x, order, expected int64, t *testing.T) {

	xBig := big.NewInt(x)

	bigField := NewBigFiniteField(
		big.NewInt(order),
	)

	result := bigField.Cube(xBig)

	if result.Int64() != expected {
		t.Error(fmt.Sprintf("Calculated Cube: %s does not match expected sum: %d", result.Text(10), expected))
	}

}
