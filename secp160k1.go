package secpk1

import (
	"crypto/elliptic"
	"math/big"
)

var secp160k1 *curve

func initSECP160K1() {
	secp160k1 = &curve{elliptic.CurveParams{Name: "secp160k1"}}
	secp160k1.P, _ = new(big.Int).SetString("FFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFEFFFFAC73", 16)
	secp160k1.N, _ = new(big.Int).SetString("0100000000000000000001B8FA16DFAB9ACA16B6B3", 16)
	secp160k1.B, _ = new(big.Int).SetString("0000000000000000000000000000000000000007", 16)
	secp160k1.Gx, _ = new(big.Int).SetString("3B4C382CE37AA192A4019E763036F4F5DD4D7EBB", 16)
	secp160k1.Gy, _ = new(big.Int).SetString("938CF935318FDCED6BC28286531733C3F03C4FEE", 16)
	secp160k1.BitSize = 160
}

// S160 returns an elliptic.Curve which implements secp160k1.
func S160() elliptic.Curve {
	initonce.Do(initSECP160K1)
	return secp160k1
}
