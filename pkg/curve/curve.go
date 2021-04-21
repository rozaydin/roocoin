package curve

import (
	"fmt"
	"math"
	"roo/coin/pkg/operations"
)

type Curve struct {
	a, b float64
	ops  operations.Operations
}

func (c *Curve) String() string {
	return fmt.Sprintf("y^2=x^3+%fx+%f", c.a, c.b)
}

func NewCurve(a, b float64, ops operations.Operations) Curve {
	return Curve{
		a,
		b,
		ops,
	}
}

func (c *Curve) Equal(other *Curve) bool {

	if c.a == other.a && c.b == other.b {
		return true
	}

	return false
}

// NewPoint creates a new point
func (c *Curve) NewPoint(x, y float64) (CurvePoint, error) {

	var point CurvePoint

	if !c.IsPointOnCurve(x, y) {
		err := fmt.Errorf("Point (%f, %f) is not on y^2 = x^3 + %fx + %f", x, y, c.a, c.b)
		return point, err
	}

	point.x = x
	point.y = y
	point.order = -1
	point.curve = c

	return point, nil
}

// NewIdentityPoint returns a point at positive infinity
func (c *Curve) NewIdentityPoint() CurvePoint {
	return CurvePoint{
		x:     math.Inf(0),
		y:     math.Inf(0),
		curve: c,
	}
}

func (c *Curve) IsPointOnCurve(x, y float64) bool {

	if math.IsInf(x, 0) && math.IsInf(y, 0) {
		return true
	}

	a := c.a
	b := c.b
	ops := c.ops

	// y^2 = x^3 + ax + b

	y2 := ops.Exp(y, 2)

	x3 := ops.Exp(x, 3)
	ax := ops.Multiply(a, x)
	xSum := ops.Sum(ops.Sum(x3, ax), b)

	return y2 == xSum
}

func isOnSameVerticalLine(point1 CurvePoint, point2 CurvePoint) bool {
	return (point1.x == point2.x) && (point1.y != point2.y)
}

func findSumOfDifferentPoints(p1 CurvePoint, p2 CurvePoint) CurvePoint {

	x1 := p1.x
	y1 := p1.y
	x2 := p2.x
	y2 := p2.y

	ops := p1.curve.ops

	// y2 - y1 / x2 - x1

	x2minusx1 := ops.Subtract(x2, x1)
	y2minusy1 := ops.Subtract(y2, y1)

	slope := ops.Divide(y2minusy1, x2minusx1)

	// slope ^2 - x1 - x2
	var x3 float64 = ops.Subtract(ops.Subtract(ops.Exp(slope, 2), x1), x2)
	// slope * (x1-x3) - y1
	var y3 float64 = ops.Subtract(ops.Multiply(slope, ops.Subtract(x1, x3)), y1)

	return CurvePoint{
		x:     x3,
		y:     y3,
		curve: p1.curve,
	}
}

func findSumOfSamePoints(p1 CurvePoint) CurvePoint {

	a := p1.curve.a
	x1 := p1.x
	y1 := p1.y

	ops := p1.curve.ops

	// (3*x1^2 + a) / 2*y1
	slope := ops.Divide(ops.Sum(ops.Multiply(3, ops.Exp(x1, 2)), a), ops.Multiply(2, y1))

	// slope^2 - 2*x1
	x3 := ops.Subtract(ops.Exp(slope, 2), ops.Multiply(2, x1))
	// slope * (x1 - x3) - y1
	y3 := ops.Subtract(ops.Multiply(slope, ops.Subtract(x1, x3)), y1)

	return CurvePoint{
		x:     x3,
		y:     y3,
		curve: p1.curve,
	}
}
