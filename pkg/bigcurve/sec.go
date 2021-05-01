package bigcurve

import (
	"fmt"
	"math/big"
)

// Exports Point as uncomressed SEC formatted string
func ExportSecUncompressed(point *BigCurvePoint) string {
	secUncompressed := fmt.Sprintf("04%x%x", to32Byte(point.x), to32Byte(point.y))
	return secUncompressed
}

func ExportSecCompressed(point *BigCurvePoint) string {
	y_copy := new(big.Int).Set(point.y)
	y_mod_two := y_copy.Mod(y_copy, big.NewInt(2))

	if y_mod_two.Cmp(big.NewInt(0)) == 0 {
		// y is even 02
		secCompressed := fmt.Sprintf("02%x", to32Byte(point.x))
		return secCompressed
	} else {
		// y is odd 03
		secCompressed := fmt.Sprintf("03%x", to32Byte(point.x))
		return secCompressed
	}
}

func to32Byte(value *big.Int) []byte {
	valueContainer := make([]byte, 32)
	value.FillBytes(valueContainer)
	return valueContainer
}
