package bigcurve

import (
	"math/big"
	"roo/coin/pkg/finitefield"
	"testing"
)

func TestValidPointsAreOnCurve(t *testing.T) {

	ops := finitefield.NewBigFiniteField(big.NewInt(223))

	curve := BigCurve{
		a:   0,
		b:   7,
		ops: ops,
	}

	_, err := curve.NewPoint(big.NewInt(192), big.NewInt(105))
	if err != nil {
		t.Error("Point should be on curve!")
	}

	_, err = curve.NewPoint(big.NewInt(17), big.NewInt(56))
	if err != nil {
		t.Error("Point should be on curve!")
	}

	_, err = curve.NewPoint(big.NewInt(1), big.NewInt(193))
	if err != nil {
		t.Error("Point should be on curve!")
	}
}

func TestInvalidPointsAreNotOnCurve(t *testing.T) {

	ops := finitefield.NewBigFiniteField(big.NewInt(223))

	curve := BigCurve{
		a:   0,
		b:   7,
		ops: ops,
	}

	_, err := curve.NewPoint(big.NewInt(200), big.NewInt(119))
	if err == nil {
		t.Error("Point should not be on curve!")
	}

	_, err = curve.NewPoint(big.NewInt(42), big.NewInt(99))
	if err == nil {
		t.Error("Point should not be on curve!")
	}
}

func TestSummationOverCurve(t *testing.T) {

	ops := finitefield.NewBigFiniteField(big.NewInt(223))

	curve := BigCurve{
		a:   0,
		b:   7,
		ops: ops,
	}

	// (170,142) + (60,139)
	point1, _ := curve.NewPoint(big.NewInt(192), big.NewInt(105))
	point2, _ := curve.NewPoint(big.NewInt(17), big.NewInt(56))
	point3, err := point1.Sum(point2)

	if err != nil {
		t.Error("valid point summation failed!")
	}

	if !point3.IsEqual(big.NewInt(170), big.NewInt(142)) {
		t.Error("summation does not match the expected result!")
	}

	// Identity point summation
	point1 = curve.NewIdentityPoint()

	if !point1.IsIdentity() {
		t.Error("does not identify identity point as identity point!")
	}

	point2, _ = curve.NewPoint(big.NewInt(192), big.NewInt(105))
	point3, err = point1.Sum(point2)

	if err != nil {
		t.Error("valid point summation failed!")
	}

	if !point3.IsEqual(big.NewInt(192), big.NewInt(105)) {
		t.Error("summation does not match the expected result!")
	}

	// (192,105) + (192,105)
	point1, _ = curve.NewPoint(big.NewInt(192), big.NewInt(105))
	point2, _ = curve.NewPoint(big.NewInt(192), big.NewInt(105))
	point3, err = point1.Sum(point2)

	if err != nil {
		t.Error("valid point summation failed!")
	}

	if !point3.IsEqual(big.NewInt(49), big.NewInt(71)) {
		t.Error("summation does not match the expected result!")
	}

	// (47,71) + (17,56)
	point1, _ = curve.NewPoint(big.NewInt(47), big.NewInt(71))
	point2, _ = curve.NewPoint(big.NewInt(117), big.NewInt(141))
	point3, err = point1.Sum(point2)

	if err != nil {
		t.Error("valid point summation failed!")
	}

	if !point3.IsEqual(big.NewInt(60), big.NewInt(139)) {
		t.Error("summation does not match the expected result!")
	}

	// (143,98) + (76,66)
	point1, _ = curve.NewPoint(big.NewInt(143), big.NewInt(98))
	point2, _ = curve.NewPoint(big.NewInt(76), big.NewInt(66))
	point3, err = point1.Sum(point2)

	if err != nil {
		t.Error("valid point summation failed!")
	}

	if !point3.IsEqual(big.NewInt(47), big.NewInt(71)) {
		t.Error("summation does not match the expected result!")
	}

}

func TestSelfSummation(t *testing.T) {

	ops := finitefield.NewBigFiniteField(big.NewInt(223))

	curve := BigCurve{
		a:   0,
		b:   7,
		ops: ops,
	}

	two := big.NewInt(2)

	point1, _ := curve.NewPoint(big.NewInt(192), big.NewInt(105))
	selfSummationResult, _ := point1.Mul(two)
	if !selfSummationResult.IsEqual(big.NewInt(49), big.NewInt(71)) {
		t.Error("self summation failed!")
	}

	point1, _ = curve.NewPoint(big.NewInt(143), big.NewInt(98))
	selfSummationResult, _ = point1.Mul(two)
	if !selfSummationResult.IsEqual(big.NewInt(64), big.NewInt(168)) {
		t.Error("self summation failed!")
	}

	point1, _ = curve.NewPoint(big.NewInt(47), big.NewInt(71))
	selfSummationResult, _ = point1.Mul(two)
	if !selfSummationResult.IsEqual(big.NewInt(36), big.NewInt(111)) {
		t.Error("self summation failed!")
	}

	point1, _ = curve.NewPoint(big.NewInt(47), big.NewInt(71))
	selfSummationResult, _ = point1.Mul(big.NewInt(4))
	if !selfSummationResult.IsEqual(big.NewInt(194), big.NewInt(51)) {
		t.Error("self summation failed!")
	}

	point1, _ = curve.NewPoint(big.NewInt(47), big.NewInt(71))
	selfSummationResult, _ = point1.Mul(big.NewInt(8))
	if !selfSummationResult.IsEqual(big.NewInt(116), big.NewInt(55)) {
		t.Error("self summation failed!")
	}

	point1, _ = curve.NewPoint(big.NewInt(47), big.NewInt(71))
	selfSummationResult, _ = point1.Mul(big.NewInt(21))
	if !selfSummationResult.IsIdentity() {
		t.Error("self summation failed!")
	}

}
