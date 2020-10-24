package ethash

import (
	"crypto/ecdsa"
	"crypto/sha256"
	"github.com/thorlinkorg/thor/common"
	"github.com/thorlinkorg/thor/common/hexutil"
	"github.com/thorlinkorg/thor/core/types"
	"github.com/thorlinkorg/thor/crypto/p256"
	"github.com/thorlinkorg/thor/log"
	"github.com/thorlinkorg/thor/rlp"
	"golang.org/x/crypto/sha3"
)

func SealHash(header *types.Header) (hash common.Hash) {
	hasher := sha3.NewLegacyKeccak256()

	rlp.Encode(hasher, []interface{}{
		header.ParentHash,
		header.UncleHash,
		header.Coinbase,
		header.Root,
		header.TxHash,
		header.ReceiptHash,
		header.Bloom,
		header.Difficulty,
		header.Number,
		header.GasLimit,
		header.GasUsed,
		header.Time,
		header.Extra[:len(header.Extra)-p256.DerSignatureLength], // Yes, this will panic if extra is too short
	})
	hasher.Sum(hash[:0])
	return hash
}

// extracts the Ethereum account address from a signed header.
func GetSigner(header *types.Header, publicKey []byte) (common.Hash, error) {
	signature := header.Extra[extraVanity : len(header.Extra)-1]
	log.Debug("GetSigner", "input pubkey", hexutil.Encode(publicKey))
	log.Debug("GetSigner", "signature", hexutil.Encode(signature), "actual sig length", len(signature))
	log.Debug("GetSigner", "SealHash", hexutil.Encode(SealHash(header).Bytes()))
	log.Debug("GetSigner", "header.nonce", hexutil.Encode(header.Nonce[:]))
	combine := make([]byte, 0)
	combine = append(combine, SealHash(header).Bytes()...)
	combine = append(combine, header.Nonce[:]...)
	log.Debug("GetSigner", "contentToHash", hexutil.Encode(combine))
	hashToSign := sha256.Sum256(combine)
	log.Debug("GetSigner", "hashToSign", hexutil.Encode(hashToSign[:]))
	pubkeys, err := p256.Ecrecovery(hashToSign[:], signature)
	if err != nil {
		log.Error("GetSigner", "err", err)
		return common.Hash{}, err
	}
	var decidedPubkey []byte
	if publicKey != nil {
		exp, err := p256.UnmarshalPubKey(publicKey)
		if err != nil {
			return common.Hash{}, err
		}
		r, s, err := p256.UnmarshalSig(signature)
		if !ecdsa.Verify(exp, hashToSign[:], r, s) {
			log.Error("GetSigner", "verify sig with given pubkey failed. err", err)
			return common.Hash{}, err
		}
		recoverId := byte(0xff)
		for i, k := range pubkeys {
			if p256.SamePubKey(k, exp) {
				recoverId = byte(i)
				decidedPubkey = publicKey
			}
		}
		if decidedPubkey == nil {
			log.Error("GetSigner", "recover pubkey failed. err", err)
			return common.Hash{}, err
		}
		log.Debug("GetSigner", "recovered pubkey", hexutil.Encode(decidedPubkey))
		header.Extra[len(header.Extra)-1] = recoverId
		log.Debug("GetSigner", "recoverId", recoverId)
	} else {
		recoverId := header.Extra[len(header.Extra)-1]
		decidedPubkey = p256.MarshalPubKey(pubkeys[int(recoverId)])
		log.Debug("GetSigner", "recovered pubkey", hexutil.Encode(decidedPubkey))
	}
	log.Debug("GetSigner", "extra", hexutil.Encode(header.Extra))
	hashedPubkey := sha256.Sum256(decidedPubkey)
	log.Debug("GetSigner", "hashedPubkey", hexutil.Encode(hashedPubkey[:]))
	return hashedPubkey, nil
}

type OdinAccessor interface {
	VerifyOdinProof(header *types.Header, publicKey []byte) error
}
