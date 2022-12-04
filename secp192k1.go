package koblitzcurve

import (
	"crypto/elliptic"
	"math/big"
)

var secp192k1 *curve

func initSECP192K1() {
	secp192k1 = &curve{elliptic.CurveParams{Name: "secp192k1"}}
	secp192k1.P, _ = new(big.Int).SetString("FFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFEFFFFEE37", 16)
	secp192k1.N, _ = new(big.Int).SetString("FFFFFFFFFFFFFFFFFFFFFFFE26F2FC170F69466A74DEFD8D", 16)
	secp192k1.B, _ = new(big.Int).SetString("000000000000000000000000000000000000000000000003", 16)
	secp192k1.Gx, _ = new(big.Int).SetString("DB4FF10EC057E9AE26B07D0280B7F4341DA5D1B1EAE06C7D", 16)
	secp192k1.Gy, _ = new(big.Int).SetString("9B2F2F6D9C5628A7844163D015BE86344082AA88D95E2F9D", 16)
	secp192k1.BitSize = 192
}

// S192 returns an elliptic.Curve which implements secp192k1.
func S192() elliptic.Curve {
	initonce.Do(initSECP192K1)
	return secp192k1
}
