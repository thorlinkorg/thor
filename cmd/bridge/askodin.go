package bridge

import (
	"context"
	"errors"
	"github.com/thorlinkorg/thor/common"
	"github.com/thorlinkorg/thor/consensus/ethash"
	"github.com/thorlinkorg/thor/core/types"
	"github.com/thorlinkorg/thor/log"
	"github.com/thorlinkorg/thor/thorclient"
	"time"
)

var OdinAssetsUnsatisfiedErr = errors.New("Odin assets not satifisfied")

type OdinBridge struct {
	ContractAddress common.Address
	RpcUrl          string
}

func (ob *OdinBridge) VerifyOdinProof(header *types.Header, publicKey []byte) error {
	signer, err := ethash.GetSigner(header, publicKey)
	if err != nil {
		log.Error("Verify signature failed", "err", err)
		return err
	}
	for i := 0; ; i++ {
		if i > 0 {
			time.Sleep(2 * time.Second)
			log.Warn("VerifyOdinProof", "retry", i)
		}
		if err := ob.verifyOdinProof(header, signer); err == OdinAssetsUnsatisfiedErr || err == nil {
			return err
		} else if i == 3 {
			log.Error("VerifyOdinProof failed for 3 times", "err", err)
			return err
		}
	}
}

func (ob *OdinBridge) verifyOdinProof(header *types.Header, signer common.Hash) error {
	n := time.Now()
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer func(start time.Time) {
		cancel()
		//log.Info("verifyOdin", "elapsed", common.PrettyDuration(time.Since(start)))
	}(n)

	if ob.RpcUrl == "" {
		return errors.New("empty Odin RPC URL")
	}
	if ob.ContractAddress == (common.Address{}) {
		return errors.New("empty Odin asset contract")
	}
	number := header.Number
	log.Info("VerifyOdinProof", "signer", signer, "number", number)
	ethcli, err := thorclient.DialContext(ctx, ob.RpcUrl)
	if err != nil {
		log.Error("Dial Odin RPC failed", "err", err, "odinrpc", ob.RpcUrl)
		return err
	}
	contract, err := NewPledge4Thor(ob.ContractAddress, ethcli)
	if err != nil {
		log.Error("Accessing contract failed", "err", err)
		return err
	}
	result, err := contract.IsRigEligible(nil, signer, number)
	if err != nil {
		log.Error("CallContract failed", "err", err)
		return err
	}
	if !result {
		log.Error("IsRigEligible", "number", number, "result", result)
		return OdinAssetsUnsatisfiedErr
	}
	log.Info("IsRigEligible", "number", number, "result", result)
	return nil
}
