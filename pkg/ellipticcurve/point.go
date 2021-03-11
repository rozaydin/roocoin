package ellipticcurve

import (
	"fmt"
	"math"
)

type Point struct {
	a, b, x, y float64
}

// NewPoint creates a new point and returns
// if point is valid
func NewPoint(a, b, x, y float64) (Point, error) {

	var point Point

	if !isPointOnCurve(a, b, x, y) {
		err := fmt.Errorf("Point (%f, %f) is not on y^2 = x^3 + %fx + %f", x, y, a, b)
		return point, err
	}

	point.x = x
	point.y = y
	point.a = a
	point.b = b

	return point, nil
}

// NewIdentityPoint returns a point at positive infinity
func NewIdentityPoint(a, b float64) Point {
	return Point{
		x: math.Inf(0),
		y: math.Inf(0),
		a: a,
		b: b,
	}
}

func isPointOnCurve(a, b, x, y float64) bool {

	if x == math.Inf(-1) && y == math.Inf(-1) {
		return true
	}

	if x == math.Inf(+1) && y == math.Inf(+1) {
		return true
	}

	// y^2 = x^3 + ax + b
	y2 := y * y
	xSum := (x * x * x) + a*x + b

	return y2 == xSum
}

func (p Point) isIdentity() bool {

	pointX := p.x
	pointY := p.y

	return math.IsInf(float64(pointX), 0) && math.IsInf(float64(pointY), 0)
}

// Sum sums 2 points on the curve and returns the result
func (p Point) Sum(other Point) (Point, error) {

	var sum Point

	if p.a != other.a || p.b != other.b {
		return sum, fmt.Errorf("elliptic curve functions are different, can not sum points")
	}

	if p.isIdentity() {
		return other, nil
	}

	if other.isIdentity() {
		return p, nil
	}

	if isOnSameVerticalLine(p, other) {
		return NewIdentityPoint(p.a, p.b), nil
	}

	if p == other { // sum of same points
		if p.y == 0 { // p1 == p2 and y = 0 sum at infinitiy
			return NewIdentityPoint(p.a, p.b), nil
		} else {
			return findSumOfSamePoints(p), nil
		}
	}

	// sum of different points
	return findSumOfDifferentPoints(p, other), nil
}

func isOnSameVerticalLine(point1 Point, point2 Point) bool {
	return (point1.x == point2.x) && (point1.y != point2.y)
}

func findSumOfDifferentPoints(p1 Point, p2 Point) Point {

	a := p1.a
	b := p1.b

	x1 := p1.x
	y1 := p1.y
	x2 := p2.x
	y2 := p2.y

	slope := (y2 - y1) / (x2 - x1)

	var x3 float64 = (slope * slope) - x1 - x2
	var y3 float64 = slope*(x1-x3) - y1

	return Point{
		x: x3,
		y: y3,
		a: a,
		b: b,
	}
}

func findSumOfSamePoints(p1 Point) Point {

	a := p1.a
	b := p1.b

	x1 := p1.x
	y1 := p1.y

	slope := (3*x1*x1 + a) / 2 * y1
	x3 := (slope * slope) - (2 * x1)
	y3 := slope*(x1-x3) - y1

	return Point{a: a, b: b, x: x3, y: y3}

}
