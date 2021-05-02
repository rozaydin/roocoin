package bigcurve

import (
	"fmt"
	"strings"
	"testing"
)

func TestSecUncompressedFormat(t *testing.T) {

	point := NewSec256K1Curve().SeedPoint
	secFormat := ExportSecUncompressed(point)

	if !strings.HasPrefix(secFormat, "04") {
		t.Error("Uncompressed SEC format is invalid! does not start with 04")
	}

	expectedSecFormat := "0479BE667EF9DCBBAC55A06295CE870B07029BFCDB2DCE28D959F2815B16F81798483ADA7726A3C4655DA4FBFC0E1108A8FD17B448A68554199C47D08FFB10D4B8"
	if !strings.EqualFold(secFormat, expectedSecFormat) {
		t.Error(fmt.Sprintf("SEC Format: %s does not match expected SEC format: %s", secFormat, expectedSecFormat))
	}
}

func TestSecCompressedFormat(t *testing.T) {
	point := NewSec256K1Curve().SeedPoint
	secFormat := ExportSecCompressed(point)

	if !strings.HasPrefix(secFormat, "02") {
		t.Error("Compressed SEC format is invalid! does not start with 02 (y coordinate is even)")
	}

	expectedSecFormat := "0279be667ef9dcbbac55a06295ce870b07029bfcdb2dce28d959f2815b16f81798"
	if !strings.EqualFold(secFormat, expectedSecFormat) {
		t.Error(fmt.Sprintf("SEC Compressed Format: %s does not match expected SEC Compressed format: %s", secFormat, expectedSecFormat))
	}
}
