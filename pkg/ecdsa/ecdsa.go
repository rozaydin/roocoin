package ecdsa

import (
	"crypto/sha256"
	"math/big"
	"roo/coin/pkg/bigcurve"
)

type Ecdsa struct {
	curve      bigcurve.Sec256k1Curve
	k          *big.Int
	privateKey *big.Int
	R          bigcurve.BigCurvePoint
}

func NewEcdsa(secret string, random *big.Int) (*Ecdsa, error) {

	curve := bigcurve.NewSec256K1Curve()
	k := generateRandomK()
	privateKey := new(big.Int).SetBytes([]byte(secret))
	R, err := curve.SeedPoint.Mul(k)

	if err != nil {
		return nil, err
	}

	ecdsa := Ecdsa{
		curve:      curve,
		k:          k,
		privateKey: privateKey,
		R:          R,
	}

	return &ecdsa, nil

}

func (ecdsa *Ecdsa) sign(message string) {

	

}

func (ecdsa *Ecdsa) verify(msg string, msgHash []byte) {

}

func generateRandomK() *big.Int {
	return big.NewInt(19)
}

func doubleSha256(message string) []byte {

	hash := sha256.New()
	hash.Write([]byte(message))
	secretHash := hash.Sum(nil)

	hash.Reset()
	hash.Write(secretHash)
	return hash.Sum(nil)
}
