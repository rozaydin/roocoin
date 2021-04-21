package bigcurve

import (
	"fmt"
	"math/big"
)

type BigCurvePoint struct {
	identity bool
	x, y     *big.Int
	curve    *BigCurve
}

var (
	one = big.NewInt(1)
)

func (p BigCurvePoint) GetX() *big.Int {
	return p.x
}

func (p BigCurvePoint) GetY() *big.Int {
	return p.y
}

func (p BigCurvePoint) String() string {
	return fmt.Sprintf("(%x, %x), curve: %s", p.x, p.y, p.curve.String())
}

func (p BigCurvePoint) Equal(other BigCurvePoint) bool {
	return p.curve.Equal(other.curve) && p.x.Cmp(other.x) == 0 && p.y.Cmp(other.y) == 0
}

func (p BigCurvePoint) IsEqual(x, y *big.Int) bool {
	return p.x.Cmp(x) == 0 && p.y.Cmp(y) == 0
}

func (p BigCurvePoint) IsIdentity() bool {
	return p.identity
}

func (p BigCurvePoint) Clone() BigCurvePoint {

	clone := BigCurvePoint{
		identity: p.identity,
		x:        new(big.Int).Set(p.x),
		y:        new(big.Int).Set(p.y),
		curve:    p.curve, // It is OK to use same curve
	}

	return clone

}

// Calculates how many iterations it take to reach identity point
func (p *BigCurvePoint) CalculateOrder() (*big.Int, error) {

	order := big.NewInt(0)
	temp := p.Clone()
	var err error

	for !temp.IsIdentity() {
		temp, err = temp.Sum(*p)
		order.Add(order, big.NewInt(1))
		if err != nil {
			return big.NewInt(-1), err
		}
	}

	return order, nil
}

// Sum sums 2 points on the curve and returns the result
func (p BigCurvePoint) Sum(other BigCurvePoint) (BigCurvePoint, error) {

	var sum BigCurvePoint

	if !p.curve.Equal(other.curve) {
		return sum, fmt.Errorf("curve functions are different, can not sum points")
	}

	if p.IsIdentity() {
		return other, nil
	}

	if other.IsIdentity() {
		return p, nil
	}

	if isOnSameVerticalLine(p, other) {
		return p.curve.NewIdentityPoint(), nil
	}

	if p.Equal(other) { // sum of same points
		if p.y.Cmp(big.NewInt(0)) == 0 { // p1 == p2 and y = 0 sum at infinitiy
			return p.curve.NewIdentityPoint(), nil
		} else {
			return findSumOfSamePoints(p), nil
		}
	}

	// sum of different points
	return findSumOfDifferentPoints(p, other), nil
}

// Sums point with itself `times` time, scalar multiplication
func (p BigCurvePoint) Mul(times *big.Int) (BigCurvePoint, error) {

	if times.Cmp(one) == -1 {
		return p, fmt.Errorf("Multiplication with zero or negative numbers, are not defined!")
	}

	if times.Cmp(one) == 0 {
		return p, nil
	}

	// TODO this is my bad code, we are subtracting one
	// from the times var and sum with the same number
	// like times is 21, we subtract 1 to make it 20
	// and add 20 times self over the number that makes
	// 21 times summation
	timesCopy := new(big.Int).Set(times)
	timesCopy = timesCopy.Sub(timesCopy, one)

	zero := big.NewInt(0)
	// multiplier
	summer := p.Clone()
	var err error

	for timesCopy.Cmp(zero) != 0 {

		if timesCopy.Bit(0) == 1 {
			p, err = p.Sum(summer)

			if err != nil {
				return p, err
			}
		}

		summer, err = summer.Sum(summer)
		if err != nil {
			return p, err
		}

		timesCopy = timesCopy.Rsh(timesCopy, 1)
	}

	return p, nil
}

func isOnSameVerticalLine(point1 BigCurvePoint, point2 BigCurvePoint) bool {
	return (point1.x.Cmp(point2.x) == 0) && (point1.y.Cmp(point2.y) != 0)
}

func findSumOfDifferentPoints(p1 BigCurvePoint, p2 BigCurvePoint) BigCurvePoint {

	x1 := p1.x
	y1 := p1.y
	x2 := p2.x
	y2 := p2.y

	ops := p1.curve.ops

	// y2 - y1 / x2 - x1
	x2minusx1 := ops.Sub(x2, x1)
	y2minusy1 := ops.Sub(y2, y1)

	slope := ops.Div(y2minusy1, x2minusx1)

	// slope ^2 - x1 - x2
	x3 := ops.Sub(ops.Sub(ops.Sqr(slope), x1), x2)
	// slope * (x1-x3) - y1
	y3 := ops.Sub(ops.Mul(slope, ops.Sub(x1, x3)), y1)

	return BigCurvePoint{
		x:     x3,
		y:     y3,
		curve: p1.curve,
	}
}

func findSumOfSamePoints(p1 BigCurvePoint) BigCurvePoint {

	a := big.NewInt(p1.curve.a)
	x1 := p1.x
	y1 := p1.y

	ops := p1.curve.ops

	// (3*x1^2 + a) / 2*y1
	firstTerm := ops.Sum(ops.Mul(big.NewInt(3), ops.Sqr(x1)), a)
	secondTerm := ops.Mul(big.NewInt(2), y1)
	slope := ops.Div(firstTerm, secondTerm)

	// slope^2 - 2*x1
	x3 := ops.Sub(ops.Sqr(slope), ops.Mul(big.NewInt(2), x1))
	// slope * (x1 - x3) - y1
	y3 := ops.Sub(ops.Mul(slope, ops.Sub(x1, x3)), y1)

	return BigCurvePoint{
		x:     x3,
		y:     y3,
		curve: p1.curve,
	}
}
