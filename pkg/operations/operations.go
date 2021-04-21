package operations

type Operations interface {
	Sum(num1, num2 float64) float64
	Subtract(num1, num2 float64) float64
	Multiply(num1, num2 float64) float64
	Divide(num1, num2 float64) float64
	Exp(num1 float64, power int64) float64
}
