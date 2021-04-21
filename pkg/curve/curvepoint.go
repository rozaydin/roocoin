package curve

import (
	"fmt"
	"math"
)

type CurvePoint struct {
	x, y  float64
	order int64
	curve *Curve
}

func (p *CurvePoint) String() string {
	return fmt.Sprintf("(%f, %f), curve: %s", p.x, p.y, p.curve.String())
}

func (p *CurvePoint) IsEqual(x, y float64) bool {
	return p.x == x && p.y == y
}

func (p *CurvePoint) IsIdentity() bool {
	return math.IsInf(p.x, 0) && math.IsInf(p.y, 0)
}

// Calculates how many iterations it take to reach identity point
func (p *CurvePoint) CalculateOrder() (int64, error) {

	order := int64(0)
	temp := *p
	var err error

	if p.order == -1 {
		for !temp.IsIdentity() {
			temp, err = temp.Sum(*p)
			order++
			if err != nil {
				return int64(-1), err
			}
		}
	}

	p.order = order

	return order, nil
}

// Sum sums 2 points on the curve and returns the result
func (p CurvePoint) Sum(other CurvePoint) (CurvePoint, error) {

	var sum CurvePoint

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

	if p == other { // sum of same points
		if p.y == 0 { // p1 == p2 and y = 0 sum at infinitiy
			return p.curve.NewIdentityPoint(), nil
		} else {
			return findSumOfSamePoints(p), nil
		}
	}

	// sum of different points
	return findSumOfDifferentPoints(p, other), nil
}

// Sums point with itself `times` time, scalar multiplication
func (p CurvePoint) SelfSum(times int64) (CurvePoint, error) {

	// multiplier
	summer := p
	var err error

	for times != 0 {

		if times&0x0000000000000001 == 1 {
			p, err = p.Sum(summer)

			if err != nil {
				return p, err
			}
		}

		summer, err = summer.Sum(summer)
		if err != nil {
			return p, err
		}

		times = times >> 1
	}

	return p, nil
}
