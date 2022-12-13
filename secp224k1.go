package koblitzcurve

import (
	"crypto/elliptic"
	"math/big"
)

var secp224k1 *curve

func init() {
	secp224k1 = &curve{elliptic.CurveParams{Name: "secp224k1"}}
	secp224k1.P, _ = new(big.Int).SetString("FFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFEFFFFE56D", 16)
	secp224k1.N, _ = new(big.Int).SetString("010000000000000000000000000001DCE8D2EC6184CAF0A971769FB1F7", 16)
	secp224k1.B, _ = new(big.Int).SetString("00000000000000000000000000000000000000000000000000000005", 16)
	secp224k1.Gx, _ = new(big.Int).SetString("A1455B334DF099DF30FC28A169A467E9E47075A90F7E650EB6B7A45C", 16)
	secp224k1.Gy, _ = new(big.Int).SetString("7E089FED7FBA344282CAFBD6F7E319F7C0B0BD59E2CA4BDB556D61A5", 16)
	secp224k1.BitSize = 224
}

// S224 returns an elliptic.Curve which implements secp224k1.
func S224() elliptic.Curve {
	return secp224k1
}
