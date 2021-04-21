package ecdsa

import (
	"math/big"
	"roo/coin/pkg/bigcurve"
	"testing"
)

func TestSigning(t *testing.T) {

	curve := bigcurve.NewSec256K1Curve()

	k := big.NewInt(1234567890)
	N := new(big.Int).Set(curve.GroupOrder)
	G := curve.SeedPoint.Clone()
	R, err := G.Mul(k)

	if err != nil {
		t.Error("Failed to calculate R")
	}

	secretHash, msgHash := getTestSecretAndMessage("my secret", "my message")

	e := new(big.Int).SetBytes(secretHash)
	z := new(big.Int).SetBytes(msgHash)
	r := new(big.Int).Set(R.GetX())

	publicKey, err := G.Mul(e)

	if err != nil {
		t.Error("Failed to generate public key!")
	}

	t.Logf("Public Key: %s", publicKey)
	t.Logf("Message hash: %x", z)
	t.Logf("r: %x", r)

	// Calculate S
	NMinusTwo := new(big.Int).Sub(N, big.NewInt(2))
	k_inv := k.Exp(k, NMinusTwo, N)

	r_mul_e := new(big.Int).Mul(r, e)
	r_mul_e_add_z := new(big.Int).Add(r_mul_e, z)
	r_mul_e_add_z_mul_kinv := new(big.Int).Mul(r_mul_e_add_z, k_inv)
	s := new(big.Int).Mod(r_mul_e_add_z_mul_kinv, N)

	t.Logf("calculated s: %x", s)
}

func TestVerification(t *testing.T) {

	curve := bigcurve.NewSec256K1Curve()

	msgHash := doubleSha256("my message")

	z := new(big.Int).SetBytes(msgHash)
	r, _ := new(big.Int).SetString("2b698a0f0a4041b77e63488ad48c23e8e8838dd1fb7520408b121697b782ef22", 16)
	s, _ := new(big.Int).SetString("bb14e602ef9e3f872e25fad328466b34e6734b7a0fcd58b1eb635447ffae8cb9", 16)

	G := curve.SeedPoint
	N := new(big.Int).Set(curve.GroupOrder)
	NMinusTwo := new(big.Int).Sub(N, big.NewInt(2))

	px, _ := new(big.Int).SetString("28d003eab2e428d11983f3e97c3fa0addf3b42740df0d211795ffb3be2f6c52", 16)
	py, _ := new(big.Int).SetString("ae987b9ec6ea159c78cb2a937ed89096fb218d9e7594f02b547526d8cd309e2", 16)

	publicKey, err := curve.Curve.NewPoint(px, py)

	if err != nil {
		t.Error("Failed to generate public key!")
	}

	// R = uG + vP
	// u = z/s, v = r/s

	s_inv := s.Exp(s, NMinusTwo, N)
	u := new(big.Int).Mul(z, s_inv)
	uModN := u.Mod(u, N)

	v := new(big.Int).Mul(r, s_inv)
	vModN := v.Mod(v, N)

	uPoint, err := G.Mul(uModN)

	if err != nil {
		t.Error("Failed to calculate uPoint")
	}

	vPoint, err := publicKey.Mul(vModN)

	if err != nil {
		t.Error("Failed to calculate vPoint")
	}

	u_Plus_v_Point, _ := uPoint.Sum(vPoint)
	calculated_r := u_Plus_v_Point.GetX()

	t.Logf("Calculated r: %x", calculated_r)

	if calculated_r.Cmp(r) != 0 {
		t.Error("Failed to validate signature!")
	}

}

// Calculates SHA256 hash and returns result
func getTestSecretAndMessage(secret, message string) ([]byte, []byte) {

	secretDoubleHash := doubleSha256(secret)
	messageDoubleHash := doubleSha256(message)

	return secretDoubleHash, messageDoubleHash
}
