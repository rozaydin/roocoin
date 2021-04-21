package operations

import (
	"math"
)

type RealNumbers struct {
}

func (r RealNumbers) Sum(num1, num2 float64) float64 {
	return num1 + num2
}

func (r RealNumbers) Subtract(num1, num2 float64) float64 {
	return num1 - num2
}

func (r RealNumbers) Multiply(num1, num2 float64) float64 {
	return num1 * num2
}

func (r RealNumbers) Divide(num1, num2 float64) float64 {
	return num1 / num2
}

func (r RealNumbers) Exp(num1 float64, power int64) float64 {
	return math.Pow(num1, float64(power))
}
