package p256

import (
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/rand"
	"crypto/sha256"
	"encoding/asn1"
	"fmt"
	"github.com/thorlinkorg/thor/common/hexutil"
	"math/big"
)

func positive() {
	privateKey, err := ecdsa.GenerateKey(elliptic.P256(), rand.Reader)
	if err != nil {
		panic(err)
	}

	pub := privateKey.PublicKey
	pubB := elliptic.Marshal(pub.Curve, pub.X, pub.Y)
	fmt.Println("pub", hexutil.Encode(pubB))
	msg := "helloworld"
	hash := sha256.Sum256([]byte(msg))

	r, s, err := ecdsa.Sign(rand.Reader, privateKey, hash[:])
	if err != nil {
		panic(err)
	}
	type ecdsaSignature struct {
		R, S *big.Int
	}

	//fmt.Printf("signature: (0x%x, 0x%x)\n", r, s)
	sigB, err := asn1.Marshal(ecdsaSignature{r, s})
	if err != nil {
		panic(err)
	}
	fmt.Println("signature", hexutil.Encode(sigB))
	valid := ecdsa.Verify(&privateKey.PublicKey, hash[:], r, s)
	fmt.Println("signature verified:", valid)
}

func reversem() {
	msg := "helloworld"
	msgB := make([]byte, 0)
	msgB = append([]byte(msg), byte(0x0a))
	hash := sha256.Sum256(msgB)
	pubS := "0x04CFC43FC80EACDD27A9D6C962CA24A3FF181A4793FE25531ED086CD7B0C08B2071EAF8874FB0FD6D0F8AD7D0E2756815C4368C342C30099070F19CC792456AF04"
	pubB := hexutil.MustDecode(pubS)
	sigS := "0x3046022100649133B4EF09E25566CF499972319A573117E9FD837FAF80EDA94CFFF4C239200221002DAECA445B74DDBB0686448ABC5636F0B803FBEC553428279B770857DD410629"
	//sigS := "0x3046022120cd93880b5187f2427a29ea6a5971ac8ede4e27422661aae94b7d1c6603871523022120b7628b117f4978389bbca5771a2359f994b702620d311cd41dfff27d2092570f"
	sigB := hexutil.MustDecode(sigS)
	pub, err := UnmarshalPubKey(pubB)
	if err != nil {
		panic(err)
	}
	r, s, err := UnmarshalSig(sigB)
	if err != nil {
		panic(err)
	}
	valid := ecdsa.Verify(pub, hash[:], r, s)

	fmt.Println("signature verified:", valid)
}
func reverse() {
	msg := "helloworld"
	hash := sha256.Sum256([]byte(msg))
	pubS := "0x0464c25e8f741ad20d600c7597e8ae0b3793dbcbdb78f682c00493fa080d57ab4a65aed02a7f76ec017cd56ea245b9d68678555258ffbd9345e755a755294117ff"
	pubB := hexutil.MustDecode(pubS)
	sigS := "0x3046022100ffa915ff4a1aa0722a1d6f818e302c837302617917e06d5fc6f15588b38d1aae022100fddb18b751aa4f708a5412810d868a16cfab0519e161d023580d6a275d4a9487"
	sigB := hexutil.MustDecode(sigS)
	pub, err := UnmarshalPubKey(pubB)
	if err != nil {
		panic(err)
	}
	r, s, err := UnmarshalSig(sigB)
	if err != nil {
		panic(err)
	}
	valid := ecdsa.Verify(pub, hash[:], r, s)

	fmt.Println("signature verified:", valid)
}
func recoverAndCompare(pubS, sigS string, msgB []byte) bool {
	//pubS := "0x0464c25e8f741ad20d600c7597e8ae0b3793dbcbdb78f682c00493fa080d57ab4a65aed02a7f76ec017cd56ea245b9d68678555258ffbd9345e755a755294117ff"
	pubB := hexutil.MustDecode(pubS)
	pub, err := UnmarshalPubKey(pubB)
	if err != nil {
		panic(err)
	}
	printPub("k0", pub)
	//msg := "helloworld"
	//msgB := make([]byte, 0)
	//msgB = append([]byte(msg), byte(0x0a))
	hash := sha256.Sum256(msgB)
	//sigS := "0x3046022100ffa915ff4a1aa0722a1d6f818e302c837302617917e06d5fc6f15588b38d1aae022100fddb18b751aa4f708a5412810d868a16cfab0519e161d023580d6a275d4a9487"
	sigB := hexutil.MustDecode(sigS)
	fmt.Println("sig", sigS, "len", len(sigB))
	k, err := Ecrecovery(hash[:], sigB)
	if err != nil {
		panic(err)
	}
	printPub("k1", k[0])
	printPub("k2", k[1])

	var recoverId int = -1

	if SamePubKey(k[0], pub) {
		recoverId = 0
	} else if SamePubKey(k[1], pub) {
		recoverId = 1
	} else {
		panic("failed")
	}

	fmt.Println("recoverId", recoverId)
	return true
}

func printPub(lab string, pub *ecdsa.PublicKey) {
	pubB := elliptic.Marshal(pub.Curve, pub.X, pub.Y)
	fmt.Println(lab, hexutil.Encode(pubB))
}

func verifyFromMiner() {
	msg := "0x8c76f377bc31a484d1dc417d9a1ba49ad0a0d6d7391cd74f950421441738b7218d0e0000264cd0b9"
	hash := sha256.Sum256(hexutil.MustDecode(msg))
	fmt.Println("hash", hexutil.Encode(hash[:]))
	pubS := "0x04cfc43fc80eacdd27a9d6c962ca24a3ff181a4793fe25531ed086cd7b0c08b2071eaf8874fb0fd6d0f8ad7d0e2756815c4368c342c30099070f19cc792456af04"
	pubB := hexutil.MustDecode(pubS)
	sigS := "0x304502210087d26afa67316792373937dd1f786fa0ff2fc8ba9c5e2e75082b1925f47e829c02203f14081aae81930e1c03ce83e97efeacc89711859189315c6a5b30168c07dead"
	sigB := hexutil.MustDecode(sigS)
	pub, err := UnmarshalPubKey(pubB)
	if err != nil {
		panic(err)
	}
	fmt.Println("X", pub.X, "Y", pub.Y)
	r, s, err := UnmarshalSig(sigB)
	if err != nil {
		panic(err)
	}
	fmt.Println("r", r.Text(16), "s", s.Text(16))
	valid := ecdsa.Verify(pub, hash[:], r, s)
	if !valid {
		panic("verify failed")
	}
	recoverAndCompare(pubS, sigS, hexutil.MustDecode(msg))
}
