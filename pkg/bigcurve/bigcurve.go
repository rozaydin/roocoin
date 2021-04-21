package bigcurve

import (
	"fmt"
	"math/big"
	"roo/coin/pkg/operations"
)

type BigCurve struct {
	a, b int64
	ops  operations.BigNumbers
}

func (c *BigCurve) String() string {
	return fmt.Sprintf("y^2=x^3+%dx+%d", c.a, c.b)
}

func (c *BigCurve) GetOps() operations.BigNumbers {
	return c.ops
}

func NewBigCurve(a, b int64, ops operations.BigNumbers) BigCurve {
	return BigCurve{
		a,
		b,
		ops,
	}
}

func (c *BigCurve) Equal(other *BigCurve) bool {

	if c.a == other.a && c.b == other.b {
		return true
	}

	return false
}

// NewPoint creates a new point
func (c *BigCurve) NewPoint(x, y *big.Int) (BigCurvePoint, error) {

	var point BigCurvePoint

	if !c.IsPointOnCurve(x, y) {
		err := fmt.Errorf("Point (%s, %s) is not on y^2 = x^3 + %dx + %d", x, y, c.a, c.b)
		return point, err
	}

	point.x = x
	point.y = y
	point.curve = c

	return point, nil
}

// NewIdentityPoint returns a point at infinity
func (c *BigCurve) NewIdentityPoint() BigCurvePoint {
	return BigCurvePoint{
		identity: true,
		x:        big.NewInt(-1),
		y:        big.NewInt(-1),
		curve:    c,
	}
}

func (c *BigCurve) IsPointOnCurve(x, y *big.Int) bool {

	// TODO skips cases where point is at infinity

	a := big.NewInt(c.a)
	b := big.NewInt(c.b)
	ops := c.ops

	// y^2 = x^3 + ax + b
	y2 := ops.Sqr(y)

	x3 := ops.Cube(x)
	ax := ops.Mul(a, x)
	xSum := ops.Sum(ops.Sum(x3, ax), b)

	return y2.Cmp(xSum) == 0
}
