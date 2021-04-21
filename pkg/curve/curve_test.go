package curve

import (
	"math"
	"roo/coin/pkg/finitefield"
	"roo/coin/pkg/operations"
	"testing"
)

func TestValidPointsAreCreated(t *testing.T) {

	curve := NewCurve(5, 7, operations.RealNumbers{})

	_, err := curve.NewPoint(-1, -1)
	if err != nil {
		t.Error("Failed to create a valid point on the curve!")
	}

	_, err = curve.NewPoint(18, 77)
	if err != nil {
		t.Error("Failed to create a valid point on the curve!")
	}
}

func TestInvalidPointsAreNotCreated(t *testing.T) {

	curve := NewCurve(5, 7, operations.RealNumbers{})

	_, err := curve.NewPoint(2, 4)
	if err == nil {
		t.Error("Should not permit invalid points to be created!")
	}

	_, err = curve.NewPoint(5, 7)
	if err == nil {
		t.Error("Should not permit invalid points to be created!")
	}

}

func TestPointsAtInfinityAreOnCurve(t *testing.T) {

	curve := NewCurve(5, 7, operations.RealNumbers{})

	var x = math.Inf(-1)
	var y = math.Inf(-1)

	_, err := curve.NewPoint(x, y)

	if err != nil {
		t.Error("Should permit point at negative infinity as valid")
	}

	x *= -1
	y *= -1

	_, err = curve.NewPoint(x, y)

	if err != nil {
		t.Error("Should permit point at infinity as valid")
	}
}

func TestSummingPoints(t *testing.T) {

	curve := NewCurve(0, 1, operations.RealNumbers{})

	// when summed with identity should return same point
	point1, _ := curve.NewPoint(0, 1)
	point2, _ := curve.NewPoint(0, -1)
	point3, _ := curve.NewPoint(-1, 0)

	identity := curve.NewIdentityPoint()
	sum, err := point1.Sum(identity)

	if err != nil || sum != point1 {
		t.Error("Failed to sum identity point with a point")
	}

	// when on same vertical line, should return identity
	// when summed with inverse should return identity
	sum, err = point1.Sum(point2)
	if err != nil || sum != identity {
		t.Error("Failed to sum point with it's inverse")
	}

	sum, err = point3.Sum(point3)
	if err != nil || sum != identity {
		t.Error("Failed to sum point where y=0 and x's are same")
	}

}

func TestValidPointsOnCurve(t *testing.T) {

	ops := finitefield.NewFiniteField(223)

	curve := Curve{
		a:   0,
		b:   7,
		ops: ops,
	}

	_, err := curve.NewPoint(192, 105)
	if err != nil {
		t.Error("Point should be on curve!")
	}

	_, err = curve.NewPoint(17, 56)
	if err != nil {
		t.Error("Point should be on curve!")
	}

	_, err = curve.NewPoint(1, 193)
	if err != nil {
		t.Error("Point should be on curve!")
	}

}

func TestInvalidPointsOnCurve(t *testing.T) {

	ops := finitefield.NewFiniteField(223)

	curve := Curve{
		a:   0,
		b:   7,
		ops: ops,
	}

	_, err := curve.NewPoint(200, 119)
	if err == nil {
		t.Error("Point should not be on curve!")
	}

	_, err = curve.NewPoint(42, 99)
	if err == nil {
		t.Error("Point should not be on curve!")
	}
}

func TestSummationOverFiniteField(t *testing.T) {

	ops := finitefield.NewFiniteField(223)

	curve := Curve{
		a:   0,
		b:   7,
		ops: ops,
	}

	// (170,142) + (60,139)
	point1, _ := curve.NewPoint(192, 105)
	point2, _ := curve.NewPoint(17, 56)
	point3, err := point1.Sum(point2)

	if err != nil {
		t.Error("valid point summation failed!")
	}

	if !point3.IsEqual(170, 142) {
		t.Error("summation does not match the expected result!")
	}

	// (47,71) + (17,56)
	point1, _ = curve.NewPoint(47, 71)
	point2, _ = curve.NewPoint(117, 141)
	point3, err = point1.Sum(point2)

	if err != nil {
		t.Error("valid point summation failed!")
	}

	if !point3.IsEqual(60, 139) {
		t.Error("summation does not match the expected result!")
	}

	// (143,98) + (76,66)
	point1, _ = curve.NewPoint(143, 98)
	point2, _ = curve.NewPoint(76, 66)
	point3, err = point1.Sum(point2)

	if err != nil {
		t.Error("valid point summation failed!")
	}

	if !point3.IsEqual(47, 71) {
		t.Error("summation does not match the expected result!")
	}
}

func TestSamePointAddition(t *testing.T) {

	ops := finitefield.NewFiniteField(223)

	curve := Curve{
		a:   0,
		b:   7,
		ops: ops,
	}

	point1, _ := curve.NewPoint(192, 105)
	selfSummationResult, _ := point1.SelfSum(1)
	if !selfSummationResult.IsEqual(49, 71) {
		t.Error("self summation failed!")
	}

	point1, _ = curve.NewPoint(143, 98)
	selfSummationResult, _ = point1.SelfSum(1)
	if !selfSummationResult.IsEqual(64, 168) {
		t.Error("self summation failed!")
	}

	point1, _ = curve.NewPoint(47, 71)
	selfSummationResult, _ = point1.SelfSum(1)
	if !selfSummationResult.IsEqual(36, 111) {
		t.Error("self summation failed!")
	}

	point1, _ = curve.NewPoint(47, 71)
	selfSummationResult, _ = point1.SelfSum(3)
	if !selfSummationResult.IsEqual(194, 51) {
		t.Error("self summation failed!")
	}

	point1, _ = curve.NewPoint(47, 71)
	selfSummationResult, _ = point1.SelfSum(7)
	if !selfSummationResult.IsEqual(116, 55) {
		t.Error("self summation failed!")
	}

	point1, _ = curve.NewPoint(47, 71)
	selfSummationResult, _ = point1.SelfSum(20)
	if !selfSummationResult.IsIdentity() {
		t.Error("self summation failed!")
	}

}
