package bigcurve

import (
	"testing"
)

func TestOrderOfSeedPoint(t *testing.T) {

	curve := NewSec256K1Curve()

	identityPoint, err := curve.SeedPoint.Mul(curve.GroupOrder)

	if err != nil {
		t.Error("Failed to sum up to order N!")
	}

	if !identityPoint.IsIdentity() {
		t.Error("Summing order times did not end up as identity!")
	}

}

// func TestFindOrderOfSeedPoint(t *testing.T) {

// 	// times out :), test sometime again with no timeout to see if it calculates
// 	curve := NewSec256K1Curve()

// 	order, err := curve.SeedPoint.CalculateOrder()
// 	if err != nil {
// 		t.Error(fmt.Sprintf("Failed to calculate order! err: %s", err))
// 	}

// 	if curve.GroupOrder.Cmp(order) != 0 {
// 		t.Error("Calculated order does not match expected order!")
// 	}

// }
