package common

import (
	"errors"
	"github.com/thorlinkorg/thor/common/hexutil"
	"os"
	"strings"
)

var (
	Not0TPrefixErr = errors.New("address should start with '0T'")
	check0T        = true
)

type TAddress Address

func init() {
	nocheck, set := os.LookupEnv("nocheck0T")
	if set && nocheck == "true" {
		check0T = false
	}
}

func (a *TAddress) UnmarshalText(input []byte) error {
	if check0T {
		if !strings.HasPrefix(string(input), "0T") {
			return Not0TPrefixErr
		}
	}
	tinput := make([]byte, 0)
	prefix := []byte("0x")
	tinput = append(tinput, prefix...)
	tinput = append(tinput, input[2:]...)
	return hexutil.UnmarshalFixedText("Address", tinput, a[:])
}

func (a TAddress) ToEthAddress() Address {
	return Address(a)
}
