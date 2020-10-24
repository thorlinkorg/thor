package main

import (
	"context"
	"crypto/ecdsa"
	"crypto/elliptic"
	crand "crypto/rand"
	"crypto/sha256"
	"encoding/asn1"
	"errors"
	"flag"
	"fmt"
	"github.com/thorlinkorg/thor/common"
	"github.com/thorlinkorg/thor/common/hexutil"
	"github.com/thorlinkorg/thor/consensus/ethash"
	"github.com/thorlinkorg/thor/core/types"
	"github.com/thorlinkorg/thor/crypto/p256"
	"github.com/thorlinkorg/thor/log"
	"github.com/thorlinkorg/thor/rpc"
	"math"
	"math/big"
	"math/rand"
	"net/http"
	"os"
	"runtime"
	"time"
)
import _ "net/http/pprof"

const (
	fetchWorkInterval = 50 * time.Second
)

var (
	cli    *rpc.Client
	prvKey *ecdsa.PrivateKey
	pubKey []byte
	host   = flag.String("host", "localhost", "rpc host to connect")
	port   = flag.Uint("port", 27455, "rpc port to connect")
)

func init() {
	log.Root().SetHandler(log.LvlFilterHandler(log.Lvl(3), log.StreamHandler(os.Stderr, log.TerminalFormat(true))))
	flag.Parse()
	url := fmt.Sprintf("http://%s:%d", *host, *port)
	var err error
	cli, err = rpc.Dial(url)
	if err != nil {
		log.Error("ERROR: failed to get JRPC client: ", "err", err)
		return
	}

	prvKey, err = ecdsa.GenerateKey(elliptic.P256(), crand.Reader)
	//if err != nil {
	//	panic(err)
	//}
	//log.Info("priv", "prvKey.D", prvKey.D.Text(16))
	//tempKey:=hexutil.Encode(crypto.FromECDSA(prvKey))
	tempKey := "a087bd0d309b8d2b0a573ab34447608e02eb4f7a8e357ea55e807905c067e669" // p256 key\
	prvKey, err = p256.ToECDSA(tempKey)
	//derKey := []byte(tempKey)
	//prvKey, err := x509.ParseECPrivateKey(hexutil.MustDecode(tempKey))
	if err != nil {
		log.Error("ERROR: failed to get decrypt key: ", "err", err)
		return
	}
	//fmt.Println("priv", prvKey.D.Text(16))
	//pubKey = hexutil.MustDecode("0x04C857760EBE36F94D75F7A3103EFB8875993A834DBD6286D5615A90375EDE270FB67EACA30BB836D5FB94FF14256E6F0F73003F0903290FFC374CA00C42DDD0C5")
	pubKey = p256.MarshalPubKey(&prvKey.PublicKey)
	//if !bytes.Equal(pubKey, expPub) {
	//	log.Crit("unmatched pubkey", "pubkey", hexutil.Encode(pubKey))
	//}
	//pub := hexutil.Encode(pubKey)
	sh2 := sha256.Sum256(pubKey)
	sh2Pub := hexutil.Encode(sh2[:])
	log.Info("MINE", "pubKey", hexutil.Encode(pubKey), "sha2", sh2Pub)

}

type result struct {
	nonce     uint64
	mixDigest common.Hash
	sealHash  common.Hash
	signature []byte

	errc chan error
}

type work struct {
	//hash := s.ethash.SealHash(block.Header())
	//s.currentWork[0] = hash.Hex()
	//s.currentWork[1] = common.BytesToHash(SeedHash(block.NumberU64())).Hex()
	//s.currentWork[2] = common.BytesToHash(new(big.Int).Div(two256, block.Difficulty()).Bytes()).Hex()
	//s.currentWork[3] = hexutil.EncodeBig(block.Number())
	headerHash, seedHash, target common.Hash
	height                       *big.Int
}

type worker struct {
	ethsh    *ethash.Ethash
	fetchCh  chan work
	submitCh chan result
	abortCh  chan struct{}
	sign     func([]byte) *[]byte
}

func (w worker) work() {
	// Another abort channel to stop mining
	abortWorkCh := make(chan struct{})
	resultsCh := make(chan result)
	// Current work
	var currWork *work

	for {
		select {
		case <-w.abortCh:
			close(abortWorkCh)
			return
		case work := <-w.fetchCh:
			// If new work has equal or higher height, abort previous work
			if currWork != nil && work.height.Cmp(currWork.height) >= 0 {
				abortWorkCh <- struct{}{}
			}
			// Start finding the nonce
			target := new(big.Int).SetBytes(work.target.Bytes())
			go w.mine(work.height, target, work.headerHash, resultsCh, abortWorkCh)
			currWork = &work
			log.Info("started new work", "height", work.height, "headerhash", work.headerHash.Hex())

		case res := <-resultsCh:
			nonceBytes := types.EncodeNonce(res.nonce)
			log.Info("MINE", "headerHash", currWork.headerHash, "hex", hexutil.Encode(currWork.headerHash[:]))
			log.Info("MINE", "nonce", res.nonce, "nonceBytes", nonceBytes, "hex", hexutil.Encode(nonceBytes[:]))
			combine := make([]byte, 0)
			combine = append(combine, currWork.headerHash[:]...)
			combine = append(combine, nonceBytes[:]...)
			log.Info("MINE", "contentToHash", combine, "hex", hexutil.Encode(combine))
			hashToSign := sha256.Sum256(combine)
			log.Info("MINE", "hashToSign", hashToSign, "hex", hexutil.Encode(hashToSign[:]))
			sig := w.sign(hashToSign[:])
			res.signature = *sig
			log.Info("MINE", "sig", sig, "hex", hexutil.Encode(*sig))
			w.submitCh <- res
			currWork = nil
		}
	}
}
func (w *worker) mine(height, target *big.Int, sealhash common.Hash, found chan result, abort chan struct{}) {
	// Extract some data from the header
	var (
		number  = height.Uint64()
		dataset = w.ethsh.Dataset(number, false)
	)
	rseed, err := crand.Int(crand.Reader, big.NewInt(math.MaxInt64))
	if err != nil {
		log.Error("rand failed")
	}
	rander := rand.New(rand.NewSource(rseed.Int64()))
	seed := uint64(rander.Int31())
	// Start generating random nonces until we abort or find a good one
	var (
		attempts = int64(0)
		nonce    = seed
	)
	log.Info("MINE", "Started ethash search for new nonce", fmt.Sprintf(" seed=%x, sealHash=%x ", seed, sealhash))
search:
	for {
		select {
		case <-abort:
			// Mining terminated, update stats and abort
			log.Info("Ethash nonce search aborted", "attempts", nonce-seed)
			break search

		default:
			// We don't have to update hash rate on every nonce, so update after after 2^X nonces
			attempts++
			if (attempts % (1 << 18)) == 0 {
				//attempts = 0
				log.Info("mining", "attemps", attempts)
			}
			// Compute the PoW value of this nonce
			digest, value := ethash.HashimotoFull(dataset.Dataset, sealhash.Bytes(), nonce)
			if new(big.Int).SetBytes(value).Cmp(target) <= 0 {
				select {
				case found <- result{nonce, common.BytesToHash(digest), sealhash, nil, nil}:
					log.Info("Ethash nonce found and reported", "attempts", nonce-seed, "nonce", nonce)
				case <-abort:
					log.Info("Ethash nonce found but discarded", "attempts", nonce-seed, "nonce", nonce)
				}
				break search
			}
			nonce++
		}
	}
	// Datasets are unmapped in a finalizer. Ensure that the dataset stays live
	// during sealing so it's not unmapped while being read.
	runtime.KeepAlive(dataset)
}
func (w *worker) fetch() {
	ticker := time.NewTicker(fetchWorkInterval)
	defer ticker.Stop()
	// One-time ticker
	coldStarter := make(chan struct{}, 1)
	coldStarter <- struct{}{}
	// Last fetched work
	var lastFetchedWork *work

	handleWork := func() {
		wrk, err := fetchRPC()
		if err != nil {
			log.Info("WARN: Failed to fetch wrk: ", "err", err)
			return
		}
		if lastFetchedWork != nil {
			if lastFetchedWork.height.Cmp(wrk.height) > 0 {
				log.Info("MINE", "skip wrk with lower height, height", wrk.height)
				return
			}
			if lastFetchedWork.headerHash == wrk.headerHash {
				log.Info("MINE", "skip same wrk, height", wrk.height)
				return
			}
		}
		lastFetchedWork = wrk
		w.fetchCh <- *wrk
	}

	for {
		select {
		case <-w.abortCh:
			return
		case <-coldStarter:
			handleWork()
		case <-ticker.C:
			handleWork()
		}
	}
}

var rpcTimeout = 10 * time.Second

func fetchRPC() (*work, error) {
	ctx, cancel := context.WithTimeout(context.Background(), rpcTimeout)
	defer cancel()
	ret := make([]string, 4)
	if err := cli.CallContext(ctx, &ret, "eth_getWork"); err != nil {
		return nil, err
	}
	//   result[0], 32 bytes hex encoded current block header pow-hash
	//   result[1], 32 bytes hex encoded seed hash used for DAG
	//   result[2], 32 bytes hex encoded boundary condition ("target"), 2^256/difficulty
	//   result[3], hex encoded block number

	height, _ := new(big.Int).SetString(ret[3][2:], 16)
	work := &work{
		headerHash: common.HexToHash(ret[0]),
		seedHash:   common.HexToHash(ret[1]),
		target:     common.HexToHash(ret[2]),
		height:     height,
	}
	log.Info("MINE", "getwork raw", ret)
	log.Info("MINE", "getwork height", height)
	return work, nil
}

func submitRPC(res result) error {
	ctx, cancel := context.WithTimeout(context.Background(), rpcTimeout)
	defer cancel()
	var success bool
	nonceEncode := types.EncodeNonce(res.nonce)
	sealHash := res.sealHash.Hex()
	md := res.mixDigest.Hex()
	sig := hexutil.Encode(res.signature)
	log.Info("MINE", "nonce:", nonceEncode, "sealHash:", sealHash, "mixDigest:", md, "signature:", sig, "publickey", hexutil.Encode(pubKey))
	if err := cli.CallContext(ctx, &success, "eth_submitWork", nonceEncode, sealHash, md, sig, hexutil.Encode(pubKey)); err != nil {
		return err
	}
	if !success {
		return errors.New("submit work failed")
	}
	log.Info("SubmitWork", "success", success)
	return nil
}

func main() {
	go func() {
		_ = http.ListenAndServe("0.0.0.0:8080", nil)
	}()
	submitCh := make(chan result)
	wrkr := &worker{
		ethsh: ethash.New(ethash.Config{
			CacheDir:         "ethash",
			CachesInMem:      1,
			CachesOnDisk:     3,
			CachesLockMmap:   false,
			DatasetsInMem:    1,
			DatasetsOnDisk:   2,
			DatasetsLockMmap: false,
		}, nil, false),
		fetchCh:  make(chan work),
		submitCh: submitCh,
		abortCh:  nil,
		sign: func(sealHashNonce []byte) *[]byte {
			r, s, err := ecdsa.Sign(crand.Reader, prvKey, sealHashNonce[:])
			if err != nil {
				panic(err)
			}
			type ecdsaSignature struct {
				R, S *big.Int
			}
			//fmt.Printf("signature: (0x%x, 0x%x)\n", r, s)
			sigB, err := asn1.Marshal(ecdsaSignature{r, s})
			if err != nil {
				log.Error("ERROR: fail to sign", "err", err)
			}
			fmt.Println("signature", hexutil.Encode(sigB))
			return &sigB
		},
	}
	go wrkr.work()
	go wrkr.fetch()

	for {
		select {
		case res := <-submitCh:
			if err := submitRPC(res); err != nil {
				log.Info("MINE", "failed to submit work", err)
			} else {
				log.Info("submitted work")
			}
		}
	}
}
