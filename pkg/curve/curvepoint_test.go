package curve

import (
	"fmt"
	"roo/coin/pkg/finitefield"
	"testing"
)

func TestCalculateOrder(t *testing.T) {

	var curve Curve = Curve{
		a:   0,
		b:   7,
		ops: finitefield.NewFiniteField(223),
	}

	seedPoint, err := curve.NewPoint(47, 71)
	if err != nil {
		t.Error("Provided point is not on curve!")
	}

	order, err := seedPoint.CalculateOrder()
	if err != nil {
		errMsg := fmt.Sprintf("Failed to calculate order! err: %s", err)
		t.Error(errMsg)
	}

	expectedOrder := int64(20)

	if order != expectedOrder {
		errMsg := fmt.Sprintf("calculated order: %d does not match expected order: %d!", order, expectedOrder)
		t.Error(errMsg)
	}

}
