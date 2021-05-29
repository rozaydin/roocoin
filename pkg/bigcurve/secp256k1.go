package bigcurve

import (
	"fmt"
	"math/big"
	"roo/coin/pkg/finitefield"
	"roo/coin/pkg/util"
)

type Sec256k1Curve struct {
	Curve      *BigCurve
	Order      *big.Int
	GroupOrder *big.Int
	SeedPoint  *BigCurvePoint
}

func NewSec256K1Curve() Sec256k1Curve {

	order, _ := new(big.Int).SetString("fffffffffffffffffffffffffffffffffffffffffffffffffffffffefffffc2f", 16)
	// max valid private key, excluded
	n, _ := new(big.Int).SetString("fffffffffffffffffffffffffffffffebaaedce6af48a03bbfd25e8cd0364141", 16)
	gx, _ := new(big.Int).SetString("79be667ef9dcbbac55a06295ce870b07029bfcdb2dce28d959f2815b16f81798", 16)
	gy, _ := new(big.Int).SetString("483ada7726a3c4655da4fbfc0e1108a8fd17b448a68554199c47d08ffb10d4b8", 16)

	bigCurve := BigCurve{
		a:   0,
		b:   7,
		ops: finitefield.NewBigFiniteField(order),
	}

	g, _ := bigCurve.NewPoint(gx, gy)

	return Sec256k1Curve{
		Curve:      &bigCurve,
		Order:      order,
		GroupOrder: n,
		SeedPoint:  &g,
	}
}

type PrivateKey struct {
	E     *big.Int
	P     *BigCurvePoint
	Curve *Sec256k1Curve
}

func NewPrivateKey(seed string) (*PrivateKey, error) {
	sec256k1Curve := NewSec256K1Curve()
	e := new(big.Int).SetBytes(util.CalculateSHA3_256Hash(seed))
	if e.Cmp(sec256k1Curve.Order) != -1 {
		return nil, fmt.Errorf("generated PrivateKey: %x is larger then SEC256k1 order: %x", e, sec256k1Curve.Order)
	}

	publicKey, err := sec256k1Curve.SeedPoint.Mul(e)
	if err != nil {
		return nil, err
	}

	privateKey := new(PrivateKey)
	privateKey.Curve = &sec256k1Curve
	privateKey.E = e
	privateKey.P = &publicKey

	return privateKey, nil
}
