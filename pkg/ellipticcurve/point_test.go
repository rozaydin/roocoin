package ellipticcurve

import (
	"math"
	"testing"
)

func TestValidPointsAreCreated(t *testing.T) {

	_, err := NewPoint(5, 7, -1, -1)
	if err != nil {
		t.Error("Failed to create a valid point on the curve!")
	}

	_, err = NewPoint(5, 7, 18, 77)
	if err != nil {
		t.Error("Failed to create a valid point on the curve!")
	}
}

func TestInvalidPointsAreNotCreated(t *testing.T) {

	_, err := NewPoint(5, 7, 2, 4)
	if err == nil {
		t.Error("Should not permit invalid points to be created!")
	}

	_, err = NewPoint(5, 7, 5, 7)
	if err == nil {
		t.Error("Should not permit invalid points to be created!")
	}

}

func TestPointsAtInfinityAreOnCurve(t *testing.T) {

	var x = math.Inf(-1)
	var y = math.Inf(-1)

	_, err := NewPoint(5, 7, x, y)

	if err != nil {
		t.Error("Should permit point at negative infinity as valid")
	}

	x *= -1
	y *= -1

	_, err = NewPoint(5, 7, x, y)

	if err != nil {
		t.Error("Should permit point at infinity as valid")
	}
}

func TestSummingPoints(t *testing.T) {

	// when summed with identity should return same point
	point1, _ := NewPoint(0, 1, 0, 1)
	point2, _ := NewPoint(0, 1, 0, -1)
	point3, _ := NewPoint(0, 1, -1, 0)

	identity := NewIdentityPoint(0, 1)
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
