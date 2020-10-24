package p256

import (
	"crypto/ecdsa"
	"crypto/elliptic"
	"encoding/hex"
	"errors"
	"fmt"
	"github.com/btcsuite/btcd/btcec"
	"math/big"
)

const DerSignatureLength = 72 + 1

var errInvalidPubkey = errors.New("invalid secp256k1 public key")

func UnmarshalSig(b []byte) (*big.Int, *big.Int, error) {
	sig, err := btcec.ParseSignature(b, elliptic.P256())
	if err != nil {
		return nil, nil, err
	}
	return sig.R, sig.S, nil
}

func SamePubKey(key1, key2 *ecdsa.PublicKey) bool {
	x := key1.X.Cmp(key2.X)
	y := key2.Y.Cmp(key2.Y)
	return x == 0 && y == 0
}

func Ecrecovery(data, rawSign []byte) ([]*ecdsa.PublicKey, error) {
	r, s, err := UnmarshalSig(rawSign)
	if err != nil {
		fmt.Println(err)
	}
	expy := new(big.Int).Sub(elliptic.P256().Params().N, big.NewInt(2))
	rinv := new(big.Int).Exp(r, expy, elliptic.P256().Params().N)
	z := new(big.Int).SetBytes(data)

	xxx := new(big.Int).Mul(r, r)
	xxx.Mul(xxx, r)

	ax := new(big.Int).Mul(big.NewInt(3), r)

	yy := new(big.Int).Sub(xxx, ax)
	yy.Add(yy, elliptic.P256().Params().B)

	//y_squard := new(big.Int).Mod(tmp4,elliptic.P256().Params().P)

	y1 := new(big.Int).ModSqrt(yy, elliptic.P256().Params().P)
	if y1 == nil {
		return nil, fmt.Errorf("can not revcovery public key")
	}

	y2 := new(big.Int).Neg(y1)
	y2.Mod(y2, elliptic.P256().Params().P)
	p1, p2 := elliptic.P256().ScalarMult(r, y1, s.Bytes())
	p3, p4 := elliptic.P256().ScalarBaseMult(z.Bytes())

	p5 := new(big.Int).Neg(p4)
	p5.Mod(p5, elliptic.P256().Params().P)

	q1, q2 := elliptic.P256().Add(p1, p2, p3, p5)
	q3, q4 := elliptic.P256().ScalarMult(q1, q2, rinv.Bytes())

	n1, n2 := elliptic.P256().ScalarMult(r, y2, s.Bytes())
	n3, n4 := elliptic.P256().ScalarBaseMult(z.Bytes())

	n5 := new(big.Int).Neg(n4)
	n5.Mod(n5, elliptic.P256().Params().P)

	q5, q6 := elliptic.P256().Add(n1, n2, n3, n5)
	q7, q8 := elliptic.P256().ScalarMult(q5, q6, rinv.Bytes())

	key1 := ecdsa.PublicKey{Curve: elliptic.P256(), X: q3, Y: q4}
	key2 := ecdsa.PublicKey{Curve: elliptic.P256(), X: q7, Y: q8}
	return []*ecdsa.PublicKey{&key1, &key2}, nil
}

func UnmarshalPubKey(pub []byte) (*ecdsa.PublicKey, error) {
	x, y := elliptic.Unmarshal(elliptic.P256(), pub)
	if x == nil {
		return nil, errInvalidPubkey
	}
	return &ecdsa.PublicKey{Curve: elliptic.P256(), X: x, Y: y}, nil
}

func MarshalPubKey(k *ecdsa.PublicKey) []byte {
	return elliptic.Marshal(elliptic.P256(), k.X, k.Y)
}

func ToECDSA(hexString string) (*ecdsa.PrivateKey, error) {
	d, err := hex.DecodeString(hexString)
	if byteErr, ok := err.(hex.InvalidByteError); ok {
		return nil, fmt.Errorf("invalid hex character %q in private key", byte(byteErr))
	} else if err != nil {
		return nil, errors.New("invalid hex data for private key")
	}
	pk := new(ecdsa.PrivateKey)
	pk.D = new(big.Int).SetBytes(d)
	//d, ok := new(big.Int).SetString(hexString, 16)
	//if !ok {
	//	return nil, errors.New("invalid key")
	//}\
	pk.PublicKey.Curve = elliptic.P256()
	pk.PublicKey.X, pk.PublicKey.Y = pk.PublicKey.Curve.ScalarBaseMult(pk.D.Bytes())

	// The priv.D must not be zero or negative.
	if pk.D.Sign() <= 0 {
		return nil, fmt.Errorf("invalid private key, zero or negative")
	}

	//pk.PublicKey.X, pk.PublicKey.Y = pk.PublicKey.Curve.ScalarBaseMult(d)
	if pk.PublicKey.X == nil {
		return nil, errors.New("invalid private key")
	}
	return pk, nil
}
