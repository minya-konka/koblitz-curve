package secpk1

import (
	"crypto/elliptic"
	"math/big"
)

var secp256k1 *curve

func initSECP256K1() {
	secp256k1 = &curve{elliptic.CurveParams{Name: "secp256k1"}}
	secp256k1.P, _ = new(big.Int).SetString("fffffffffffffffffffffffffffffffffffffffffffffffffffffffefffffc2f", 16)
	secp256k1.N, _ = new(big.Int).SetString("fffffffffffffffffffffffffffffffebaaedce6af48a03bbfd25e8cd0364141", 16)
	secp256k1.B, _ = new(big.Int).SetString("0000000000000000000000000000000000000000000000000000000000000007", 16)
	secp256k1.Gx, _ = new(big.Int).SetString("79be667ef9dcbbac55a06295ce870b07029bfcdb2dce28d959f2815b16f81798", 16)
	secp256k1.Gy, _ = new(big.Int).SetString("483ada7726a3c4655da4fbfc0e1108a8fd17b448a68554199c47d08ffb10d4b8", 16)
	secp256k1.BitSize = 256
}

// S256 returns an elliptic.Curve which implements secp256k1.
func S256() elliptic.Curve {
	initonce.Do(initSECP256K1)
	return secp256k1
}
