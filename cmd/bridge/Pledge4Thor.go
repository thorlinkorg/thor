// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package bridge

import (
	"math/big"
	"strings"

	ethereum "github.com/thorlinkorg/thor"
	"github.com/thorlinkorg/thor/accounts/abi"
	"github.com/thorlinkorg/thor/accounts/abi/bind"
	"github.com/thorlinkorg/thor/common"
	"github.com/thorlinkorg/thor/core/types"
	"github.com/thorlinkorg/thor/event"
)

// Reference imports to suppress errors if they are not otherwise used.
var (
	_ = big.NewInt
	_ = strings.NewReader
	_ = ethereum.NotFound
	_ = bind.Bind
	_ = common.Big1
	_ = types.BloomLookup
	_ = event.NewSubscription
)

// Pledge4ThorABI is the input ABI used to generate the binding from.
const Pledge4ThorABI = "[{\"inputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"rig\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"BindRig\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"rig\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"height\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"DeductDeposit\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"Deposit\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"InternalTransfer\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"Paused\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"rig\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"UnBindRig\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"Unpaused\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"Withdraw\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"_fuel\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"_nextFuel\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"_nextFuelHeight\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"rig\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"height\",\"type\":\"uint256\"}],\"name\":\"bindRig\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"blackhole\",\"outputs\":[{\"internalType\":\"addresspayable\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"blockRigHistory\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"decimals\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"rig\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"height\",\"type\":\"uint256\"}],\"name\":\"deductDeposit\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"}],\"name\":\"getUserRig\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"a\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"halfSpan\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"hour\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"initialBindPower\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"isBridge\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"isOwner\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"rig\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"height\",\"type\":\"uint256\"}],\"name\":\"isRigEligible\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"odinContractAddress\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"pause\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"paused\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"name\":\"rigOwners\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"bridge\",\"type\":\"address\"}],\"name\":\"setBridge\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"fuel\",\"type\":\"uint256\"}],\"name\":\"setFuel\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"nextFuel\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"nextFuelHeight\",\"type\":\"uint256\"}],\"name\":\"setNextFuel\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totalDeposit\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"rig\",\"type\":\"bytes32\"}],\"name\":\"unBindRig\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"unpause\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"userDeposits\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"userLastDeductTime\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"addresspayable\",\"name\":\"payee\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"withdraw\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"stateMutability\":\"payable\",\"type\":\"receive\"}]"

// Pledge4Thor is an auto generated Go binding around an Ethereum contract.
type Pledge4Thor struct {
	Pledge4ThorCaller     // Read-only binding to the contract
	Pledge4ThorTransactor // Write-only binding to the contract
	Pledge4ThorFilterer   // Log filterer for contract events
}

// Pledge4ThorCaller is an auto generated read-only Go binding around an Ethereum contract.
type Pledge4ThorCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// Pledge4ThorTransactor is an auto generated write-only Go binding around an Ethereum contract.
type Pledge4ThorTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// Pledge4ThorFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type Pledge4ThorFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// Pledge4ThorSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type Pledge4ThorSession struct {
	Contract     *Pledge4Thor      // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// Pledge4ThorCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type Pledge4ThorCallerSession struct {
	Contract *Pledge4ThorCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts      // Call options to use throughout this session
}

// Pledge4ThorTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type Pledge4ThorTransactorSession struct {
	Contract     *Pledge4ThorTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts      // Transaction auth options to use throughout this session
}

// Pledge4ThorRaw is an auto generated low-level Go binding around an Ethereum contract.
type Pledge4ThorRaw struct {
	Contract *Pledge4Thor // Generic contract binding to access the raw methods on
}

// Pledge4ThorCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type Pledge4ThorCallerRaw struct {
	Contract *Pledge4ThorCaller // Generic read-only contract binding to access the raw methods on
}

// Pledge4ThorTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type Pledge4ThorTransactorRaw struct {
	Contract *Pledge4ThorTransactor // Generic write-only contract binding to access the raw methods on
}

// NewPledge4Thor creates a new instance of Pledge4Thor, bound to a specific deployed contract.
func NewPledge4Thor(address common.Address, backend bind.ContractBackend) (*Pledge4Thor, error) {
	contract, err := bindPledge4Thor(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Pledge4Thor{Pledge4ThorCaller: Pledge4ThorCaller{contract: contract}, Pledge4ThorTransactor: Pledge4ThorTransactor{contract: contract}, Pledge4ThorFilterer: Pledge4ThorFilterer{contract: contract}}, nil
}

// NewPledge4ThorCaller creates a new read-only instance of Pledge4Thor, bound to a specific deployed contract.
func NewPledge4ThorCaller(address common.Address, caller bind.ContractCaller) (*Pledge4ThorCaller, error) {
	contract, err := bindPledge4Thor(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &Pledge4ThorCaller{contract: contract}, nil
}

// NewPledge4ThorTransactor creates a new write-only instance of Pledge4Thor, bound to a specific deployed contract.
func NewPledge4ThorTransactor(address common.Address, transactor bind.ContractTransactor) (*Pledge4ThorTransactor, error) {
	contract, err := bindPledge4Thor(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &Pledge4ThorTransactor{contract: contract}, nil
}

// NewPledge4ThorFilterer creates a new log filterer instance of Pledge4Thor, bound to a specific deployed contract.
func NewPledge4ThorFilterer(address common.Address, filterer bind.ContractFilterer) (*Pledge4ThorFilterer, error) {
	contract, err := bindPledge4Thor(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &Pledge4ThorFilterer{contract: contract}, nil
}

// bindPledge4Thor binds a generic wrapper to an already deployed contract.
func bindPledge4Thor(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(Pledge4ThorABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Pledge4Thor *Pledge4ThorRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Pledge4Thor.Contract.Pledge4ThorCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Pledge4Thor *Pledge4ThorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Pledge4Thor.Contract.Pledge4ThorTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Pledge4Thor *Pledge4ThorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Pledge4Thor.Contract.Pledge4ThorTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Pledge4Thor *Pledge4ThorCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Pledge4Thor.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Pledge4Thor *Pledge4ThorTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Pledge4Thor.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Pledge4Thor *Pledge4ThorTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Pledge4Thor.Contract.contract.Transact(opts, method, params...)
}

// Fuel is a free data retrieval call binding the contract method 0xdd07ee50.
//
// Solidity: function _fuel() view returns(uint256)
func (_Pledge4Thor *Pledge4ThorCaller) Fuel(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Pledge4Thor.contract.Call(opts, out, "_fuel")
	return *ret0, err
}

// Fuel is a free data retrieval call binding the contract method 0xdd07ee50.
//
// Solidity: function _fuel() view returns(uint256)
func (_Pledge4Thor *Pledge4ThorSession) Fuel() (*big.Int, error) {
	return _Pledge4Thor.Contract.Fuel(&_Pledge4Thor.CallOpts)
}

// Fuel is a free data retrieval call binding the contract method 0xdd07ee50.
//
// Solidity: function _fuel() view returns(uint256)
func (_Pledge4Thor *Pledge4ThorCallerSession) Fuel() (*big.Int, error) {
	return _Pledge4Thor.Contract.Fuel(&_Pledge4Thor.CallOpts)
}

// NextFuel is a free data retrieval call binding the contract method 0x6e06518a.
//
// Solidity: function _nextFuel() view returns(uint256)
func (_Pledge4Thor *Pledge4ThorCaller) NextFuel(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Pledge4Thor.contract.Call(opts, out, "_nextFuel")
	return *ret0, err
}

// NextFuel is a free data retrieval call binding the contract method 0x6e06518a.
//
// Solidity: function _nextFuel() view returns(uint256)
func (_Pledge4Thor *Pledge4ThorSession) NextFuel() (*big.Int, error) {
	return _Pledge4Thor.Contract.NextFuel(&_Pledge4Thor.CallOpts)
}

// NextFuel is a free data retrieval call binding the contract method 0x6e06518a.
//
// Solidity: function _nextFuel() view returns(uint256)
func (_Pledge4Thor *Pledge4ThorCallerSession) NextFuel() (*big.Int, error) {
	return _Pledge4Thor.Contract.NextFuel(&_Pledge4Thor.CallOpts)
}

// NextFuelHeight is a free data retrieval call binding the contract method 0x1d0d0a1c.
//
// Solidity: function _nextFuelHeight() view returns(uint256)
func (_Pledge4Thor *Pledge4ThorCaller) NextFuelHeight(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Pledge4Thor.contract.Call(opts, out, "_nextFuelHeight")
	return *ret0, err
}

// NextFuelHeight is a free data retrieval call binding the contract method 0x1d0d0a1c.
//
// Solidity: function _nextFuelHeight() view returns(uint256)
func (_Pledge4Thor *Pledge4ThorSession) NextFuelHeight() (*big.Int, error) {
	return _Pledge4Thor.Contract.NextFuelHeight(&_Pledge4Thor.CallOpts)
}

// NextFuelHeight is a free data retrieval call binding the contract method 0x1d0d0a1c.
//
// Solidity: function _nextFuelHeight() view returns(uint256)
func (_Pledge4Thor *Pledge4ThorCallerSession) NextFuelHeight() (*big.Int, error) {
	return _Pledge4Thor.Contract.NextFuelHeight(&_Pledge4Thor.CallOpts)
}

// Blackhole is a free data retrieval call binding the contract method 0xf34447f7.
//
// Solidity: function blackhole() view returns(address)
func (_Pledge4Thor *Pledge4ThorCaller) Blackhole(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _Pledge4Thor.contract.Call(opts, out, "blackhole")
	return *ret0, err
}

// Blackhole is a free data retrieval call binding the contract method 0xf34447f7.
//
// Solidity: function blackhole() view returns(address)
func (_Pledge4Thor *Pledge4ThorSession) Blackhole() (common.Address, error) {
	return _Pledge4Thor.Contract.Blackhole(&_Pledge4Thor.CallOpts)
}

// Blackhole is a free data retrieval call binding the contract method 0xf34447f7.
//
// Solidity: function blackhole() view returns(address)
func (_Pledge4Thor *Pledge4ThorCallerSession) Blackhole() (common.Address, error) {
	return _Pledge4Thor.Contract.Blackhole(&_Pledge4Thor.CallOpts)
}

// BlockRigHistory is a free data retrieval call binding the contract method 0x3ee8ce21.
//
// Solidity: function blockRigHistory(uint256 ) view returns(bytes32)
func (_Pledge4Thor *Pledge4ThorCaller) BlockRigHistory(opts *bind.CallOpts, arg0 *big.Int) ([32]byte, error) {
	var (
		ret0 = new([32]byte)
	)
	out := ret0
	err := _Pledge4Thor.contract.Call(opts, out, "blockRigHistory", arg0)
	return *ret0, err
}

// BlockRigHistory is a free data retrieval call binding the contract method 0x3ee8ce21.
//
// Solidity: function blockRigHistory(uint256 ) view returns(bytes32)
func (_Pledge4Thor *Pledge4ThorSession) BlockRigHistory(arg0 *big.Int) ([32]byte, error) {
	return _Pledge4Thor.Contract.BlockRigHistory(&_Pledge4Thor.CallOpts, arg0)
}

// BlockRigHistory is a free data retrieval call binding the contract method 0x3ee8ce21.
//
// Solidity: function blockRigHistory(uint256 ) view returns(bytes32)
func (_Pledge4Thor *Pledge4ThorCallerSession) BlockRigHistory(arg0 *big.Int) ([32]byte, error) {
	return _Pledge4Thor.Contract.BlockRigHistory(&_Pledge4Thor.CallOpts, arg0)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint256)
func (_Pledge4Thor *Pledge4ThorCaller) Decimals(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Pledge4Thor.contract.Call(opts, out, "decimals")
	return *ret0, err
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint256)
func (_Pledge4Thor *Pledge4ThorSession) Decimals() (*big.Int, error) {
	return _Pledge4Thor.Contract.Decimals(&_Pledge4Thor.CallOpts)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint256)
func (_Pledge4Thor *Pledge4ThorCallerSession) Decimals() (*big.Int, error) {
	return _Pledge4Thor.Contract.Decimals(&_Pledge4Thor.CallOpts)
}

// GetUserRig is a free data retrieval call binding the contract method 0x23d1ae70.
//
// Solidity: function getUserRig(address user, uint256 index) view returns(bytes32 a)
func (_Pledge4Thor *Pledge4ThorCaller) GetUserRig(opts *bind.CallOpts, user common.Address, index *big.Int) ([32]byte, error) {
	var (
		ret0 = new([32]byte)
	)
	out := ret0
	err := _Pledge4Thor.contract.Call(opts, out, "getUserRig", user, index)
	return *ret0, err
}

// GetUserRig is a free data retrieval call binding the contract method 0x23d1ae70.
//
// Solidity: function getUserRig(address user, uint256 index) view returns(bytes32 a)
func (_Pledge4Thor *Pledge4ThorSession) GetUserRig(user common.Address, index *big.Int) ([32]byte, error) {
	return _Pledge4Thor.Contract.GetUserRig(&_Pledge4Thor.CallOpts, user, index)
}

// GetUserRig is a free data retrieval call binding the contract method 0x23d1ae70.
//
// Solidity: function getUserRig(address user, uint256 index) view returns(bytes32 a)
func (_Pledge4Thor *Pledge4ThorCallerSession) GetUserRig(user common.Address, index *big.Int) ([32]byte, error) {
	return _Pledge4Thor.Contract.GetUserRig(&_Pledge4Thor.CallOpts, user, index)
}

// HalfSpan is a free data retrieval call binding the contract method 0x0a75565a.
//
// Solidity: function halfSpan() view returns(uint256)
func (_Pledge4Thor *Pledge4ThorCaller) HalfSpan(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Pledge4Thor.contract.Call(opts, out, "halfSpan")
	return *ret0, err
}

// HalfSpan is a free data retrieval call binding the contract method 0x0a75565a.
//
// Solidity: function halfSpan() view returns(uint256)
func (_Pledge4Thor *Pledge4ThorSession) HalfSpan() (*big.Int, error) {
	return _Pledge4Thor.Contract.HalfSpan(&_Pledge4Thor.CallOpts)
}

// HalfSpan is a free data retrieval call binding the contract method 0x0a75565a.
//
// Solidity: function halfSpan() view returns(uint256)
func (_Pledge4Thor *Pledge4ThorCallerSession) HalfSpan() (*big.Int, error) {
	return _Pledge4Thor.Contract.HalfSpan(&_Pledge4Thor.CallOpts)
}

// Hour is a free data retrieval call binding the contract method 0x23e5f1c5.
//
// Solidity: function hour() view returns(uint256)
func (_Pledge4Thor *Pledge4ThorCaller) Hour(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Pledge4Thor.contract.Call(opts, out, "hour")
	return *ret0, err
}

// Hour is a free data retrieval call binding the contract method 0x23e5f1c5.
//
// Solidity: function hour() view returns(uint256)
func (_Pledge4Thor *Pledge4ThorSession) Hour() (*big.Int, error) {
	return _Pledge4Thor.Contract.Hour(&_Pledge4Thor.CallOpts)
}

// Hour is a free data retrieval call binding the contract method 0x23e5f1c5.
//
// Solidity: function hour() view returns(uint256)
func (_Pledge4Thor *Pledge4ThorCallerSession) Hour() (*big.Int, error) {
	return _Pledge4Thor.Contract.Hour(&_Pledge4Thor.CallOpts)
}

// InitialBindPower is a free data retrieval call binding the contract method 0xcf2333a7.
//
// Solidity: function initialBindPower() view returns(uint256)
func (_Pledge4Thor *Pledge4ThorCaller) InitialBindPower(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Pledge4Thor.contract.Call(opts, out, "initialBindPower")
	return *ret0, err
}

// InitialBindPower is a free data retrieval call binding the contract method 0xcf2333a7.
//
// Solidity: function initialBindPower() view returns(uint256)
func (_Pledge4Thor *Pledge4ThorSession) InitialBindPower() (*big.Int, error) {
	return _Pledge4Thor.Contract.InitialBindPower(&_Pledge4Thor.CallOpts)
}

// InitialBindPower is a free data retrieval call binding the contract method 0xcf2333a7.
//
// Solidity: function initialBindPower() view returns(uint256)
func (_Pledge4Thor *Pledge4ThorCallerSession) InitialBindPower() (*big.Int, error) {
	return _Pledge4Thor.Contract.InitialBindPower(&_Pledge4Thor.CallOpts)
}

// IsBridge is a free data retrieval call binding the contract method 0x4f046c8f.
//
// Solidity: function isBridge() view returns(bool)
func (_Pledge4Thor *Pledge4ThorCaller) IsBridge(opts *bind.CallOpts) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _Pledge4Thor.contract.Call(opts, out, "isBridge")
	return *ret0, err
}

// IsBridge is a free data retrieval call binding the contract method 0x4f046c8f.
//
// Solidity: function isBridge() view returns(bool)
func (_Pledge4Thor *Pledge4ThorSession) IsBridge() (bool, error) {
	return _Pledge4Thor.Contract.IsBridge(&_Pledge4Thor.CallOpts)
}

// IsBridge is a free data retrieval call binding the contract method 0x4f046c8f.
//
// Solidity: function isBridge() view returns(bool)
func (_Pledge4Thor *Pledge4ThorCallerSession) IsBridge() (bool, error) {
	return _Pledge4Thor.Contract.IsBridge(&_Pledge4Thor.CallOpts)
}

// IsOwner is a free data retrieval call binding the contract method 0x8f32d59b.
//
// Solidity: function isOwner() view returns(bool)
func (_Pledge4Thor *Pledge4ThorCaller) IsOwner(opts *bind.CallOpts) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _Pledge4Thor.contract.Call(opts, out, "isOwner")
	return *ret0, err
}

// IsOwner is a free data retrieval call binding the contract method 0x8f32d59b.
//
// Solidity: function isOwner() view returns(bool)
func (_Pledge4Thor *Pledge4ThorSession) IsOwner() (bool, error) {
	return _Pledge4Thor.Contract.IsOwner(&_Pledge4Thor.CallOpts)
}

// IsOwner is a free data retrieval call binding the contract method 0x8f32d59b.
//
// Solidity: function isOwner() view returns(bool)
func (_Pledge4Thor *Pledge4ThorCallerSession) IsOwner() (bool, error) {
	return _Pledge4Thor.Contract.IsOwner(&_Pledge4Thor.CallOpts)
}

// IsRigEligible is a free data retrieval call binding the contract method 0x9bb988fa.
//
// Solidity: function isRigEligible(bytes32 rig, uint256 height) view returns(bool)
func (_Pledge4Thor *Pledge4ThorCaller) IsRigEligible(opts *bind.CallOpts, rig [32]byte, height *big.Int) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _Pledge4Thor.contract.Call(opts, out, "isRigEligible", rig, height)
	return *ret0, err
}

// IsRigEligible is a free data retrieval call binding the contract method 0x9bb988fa.
//
// Solidity: function isRigEligible(bytes32 rig, uint256 height) view returns(bool)
func (_Pledge4Thor *Pledge4ThorSession) IsRigEligible(rig [32]byte, height *big.Int) (bool, error) {
	return _Pledge4Thor.Contract.IsRigEligible(&_Pledge4Thor.CallOpts, rig, height)
}

// IsRigEligible is a free data retrieval call binding the contract method 0x9bb988fa.
//
// Solidity: function isRigEligible(bytes32 rig, uint256 height) view returns(bool)
func (_Pledge4Thor *Pledge4ThorCallerSession) IsRigEligible(rig [32]byte, height *big.Int) (bool, error) {
	return _Pledge4Thor.Contract.IsRigEligible(&_Pledge4Thor.CallOpts, rig, height)
}

// OdinContractAddress is a free data retrieval call binding the contract method 0x91dad278.
//
// Solidity: function odinContractAddress() view returns(address)
func (_Pledge4Thor *Pledge4ThorCaller) OdinContractAddress(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _Pledge4Thor.contract.Call(opts, out, "odinContractAddress")
	return *ret0, err
}

// OdinContractAddress is a free data retrieval call binding the contract method 0x91dad278.
//
// Solidity: function odinContractAddress() view returns(address)
func (_Pledge4Thor *Pledge4ThorSession) OdinContractAddress() (common.Address, error) {
	return _Pledge4Thor.Contract.OdinContractAddress(&_Pledge4Thor.CallOpts)
}

// OdinContractAddress is a free data retrieval call binding the contract method 0x91dad278.
//
// Solidity: function odinContractAddress() view returns(address)
func (_Pledge4Thor *Pledge4ThorCallerSession) OdinContractAddress() (common.Address, error) {
	return _Pledge4Thor.Contract.OdinContractAddress(&_Pledge4Thor.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Pledge4Thor *Pledge4ThorCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _Pledge4Thor.contract.Call(opts, out, "owner")
	return *ret0, err
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Pledge4Thor *Pledge4ThorSession) Owner() (common.Address, error) {
	return _Pledge4Thor.Contract.Owner(&_Pledge4Thor.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Pledge4Thor *Pledge4ThorCallerSession) Owner() (common.Address, error) {
	return _Pledge4Thor.Contract.Owner(&_Pledge4Thor.CallOpts)
}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_Pledge4Thor *Pledge4ThorCaller) Paused(opts *bind.CallOpts) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _Pledge4Thor.contract.Call(opts, out, "paused")
	return *ret0, err
}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_Pledge4Thor *Pledge4ThorSession) Paused() (bool, error) {
	return _Pledge4Thor.Contract.Paused(&_Pledge4Thor.CallOpts)
}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_Pledge4Thor *Pledge4ThorCallerSession) Paused() (bool, error) {
	return _Pledge4Thor.Contract.Paused(&_Pledge4Thor.CallOpts)
}

// RigOwners is a free data retrieval call binding the contract method 0xf9cf1940.
//
// Solidity: function rigOwners(bytes32 ) view returns(address)
func (_Pledge4Thor *Pledge4ThorCaller) RigOwners(opts *bind.CallOpts, arg0 [32]byte) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _Pledge4Thor.contract.Call(opts, out, "rigOwners", arg0)
	return *ret0, err
}

// RigOwners is a free data retrieval call binding the contract method 0xf9cf1940.
//
// Solidity: function rigOwners(bytes32 ) view returns(address)
func (_Pledge4Thor *Pledge4ThorSession) RigOwners(arg0 [32]byte) (common.Address, error) {
	return _Pledge4Thor.Contract.RigOwners(&_Pledge4Thor.CallOpts, arg0)
}

// RigOwners is a free data retrieval call binding the contract method 0xf9cf1940.
//
// Solidity: function rigOwners(bytes32 ) view returns(address)
func (_Pledge4Thor *Pledge4ThorCallerSession) RigOwners(arg0 [32]byte) (common.Address, error) {
	return _Pledge4Thor.Contract.RigOwners(&_Pledge4Thor.CallOpts, arg0)
}

// TotalDeposit is a free data retrieval call binding the contract method 0xf6153ccd.
//
// Solidity: function totalDeposit() view returns(uint256)
func (_Pledge4Thor *Pledge4ThorCaller) TotalDeposit(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Pledge4Thor.contract.Call(opts, out, "totalDeposit")
	return *ret0, err
}

// TotalDeposit is a free data retrieval call binding the contract method 0xf6153ccd.
//
// Solidity: function totalDeposit() view returns(uint256)
func (_Pledge4Thor *Pledge4ThorSession) TotalDeposit() (*big.Int, error) {
	return _Pledge4Thor.Contract.TotalDeposit(&_Pledge4Thor.CallOpts)
}

// TotalDeposit is a free data retrieval call binding the contract method 0xf6153ccd.
//
// Solidity: function totalDeposit() view returns(uint256)
func (_Pledge4Thor *Pledge4ThorCallerSession) TotalDeposit() (*big.Int, error) {
	return _Pledge4Thor.Contract.TotalDeposit(&_Pledge4Thor.CallOpts)
}

// UserDeposits is a free data retrieval call binding the contract method 0x0ba36dcd.
//
// Solidity: function userDeposits(address ) view returns(uint256)
func (_Pledge4Thor *Pledge4ThorCaller) UserDeposits(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Pledge4Thor.contract.Call(opts, out, "userDeposits", arg0)
	return *ret0, err
}

// UserDeposits is a free data retrieval call binding the contract method 0x0ba36dcd.
//
// Solidity: function userDeposits(address ) view returns(uint256)
func (_Pledge4Thor *Pledge4ThorSession) UserDeposits(arg0 common.Address) (*big.Int, error) {
	return _Pledge4Thor.Contract.UserDeposits(&_Pledge4Thor.CallOpts, arg0)
}

// UserDeposits is a free data retrieval call binding the contract method 0x0ba36dcd.
//
// Solidity: function userDeposits(address ) view returns(uint256)
func (_Pledge4Thor *Pledge4ThorCallerSession) UserDeposits(arg0 common.Address) (*big.Int, error) {
	return _Pledge4Thor.Contract.UserDeposits(&_Pledge4Thor.CallOpts, arg0)
}

// UserLastDeductTime is a free data retrieval call binding the contract method 0x7ba028a8.
//
// Solidity: function userLastDeductTime(address ) view returns(uint256)
func (_Pledge4Thor *Pledge4ThorCaller) UserLastDeductTime(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Pledge4Thor.contract.Call(opts, out, "userLastDeductTime", arg0)
	return *ret0, err
}

// UserLastDeductTime is a free data retrieval call binding the contract method 0x7ba028a8.
//
// Solidity: function userLastDeductTime(address ) view returns(uint256)
func (_Pledge4Thor *Pledge4ThorSession) UserLastDeductTime(arg0 common.Address) (*big.Int, error) {
	return _Pledge4Thor.Contract.UserLastDeductTime(&_Pledge4Thor.CallOpts, arg0)
}

// UserLastDeductTime is a free data retrieval call binding the contract method 0x7ba028a8.
//
// Solidity: function userLastDeductTime(address ) view returns(uint256)
func (_Pledge4Thor *Pledge4ThorCallerSession) UserLastDeductTime(arg0 common.Address) (*big.Int, error) {
	return _Pledge4Thor.Contract.UserLastDeductTime(&_Pledge4Thor.CallOpts, arg0)
}

// BindRig is a paid mutator transaction binding the contract method 0x21e9f172.
//
// Solidity: function bindRig(bytes32 rig, uint256 height) returns()
func (_Pledge4Thor *Pledge4ThorTransactor) BindRig(opts *bind.TransactOpts, rig [32]byte, height *big.Int) (*types.Transaction, error) {
	return _Pledge4Thor.contract.Transact(opts, "bindRig", rig, height)
}

// BindRig is a paid mutator transaction binding the contract method 0x21e9f172.
//
// Solidity: function bindRig(bytes32 rig, uint256 height) returns()
func (_Pledge4Thor *Pledge4ThorSession) BindRig(rig [32]byte, height *big.Int) (*types.Transaction, error) {
	return _Pledge4Thor.Contract.BindRig(&_Pledge4Thor.TransactOpts, rig, height)
}

// BindRig is a paid mutator transaction binding the contract method 0x21e9f172.
//
// Solidity: function bindRig(bytes32 rig, uint256 height) returns()
func (_Pledge4Thor *Pledge4ThorTransactorSession) BindRig(rig [32]byte, height *big.Int) (*types.Transaction, error) {
	return _Pledge4Thor.Contract.BindRig(&_Pledge4Thor.TransactOpts, rig, height)
}

// DeductDeposit is a paid mutator transaction binding the contract method 0x02d901d4.
//
// Solidity: function deductDeposit(bytes32 rig, uint256 height) returns()
func (_Pledge4Thor *Pledge4ThorTransactor) DeductDeposit(opts *bind.TransactOpts, rig [32]byte, height *big.Int) (*types.Transaction, error) {
	return _Pledge4Thor.contract.Transact(opts, "deductDeposit", rig, height)
}

// DeductDeposit is a paid mutator transaction binding the contract method 0x02d901d4.
//
// Solidity: function deductDeposit(bytes32 rig, uint256 height) returns()
func (_Pledge4Thor *Pledge4ThorSession) DeductDeposit(rig [32]byte, height *big.Int) (*types.Transaction, error) {
	return _Pledge4Thor.Contract.DeductDeposit(&_Pledge4Thor.TransactOpts, rig, height)
}

// DeductDeposit is a paid mutator transaction binding the contract method 0x02d901d4.
//
// Solidity: function deductDeposit(bytes32 rig, uint256 height) returns()
func (_Pledge4Thor *Pledge4ThorTransactorSession) DeductDeposit(rig [32]byte, height *big.Int) (*types.Transaction, error) {
	return _Pledge4Thor.Contract.DeductDeposit(&_Pledge4Thor.TransactOpts, rig, height)
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_Pledge4Thor *Pledge4ThorTransactor) Pause(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Pledge4Thor.contract.Transact(opts, "pause")
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_Pledge4Thor *Pledge4ThorSession) Pause() (*types.Transaction, error) {
	return _Pledge4Thor.Contract.Pause(&_Pledge4Thor.TransactOpts)
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_Pledge4Thor *Pledge4ThorTransactorSession) Pause() (*types.Transaction, error) {
	return _Pledge4Thor.Contract.Pause(&_Pledge4Thor.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Pledge4Thor *Pledge4ThorTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Pledge4Thor.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Pledge4Thor *Pledge4ThorSession) RenounceOwnership() (*types.Transaction, error) {
	return _Pledge4Thor.Contract.RenounceOwnership(&_Pledge4Thor.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Pledge4Thor *Pledge4ThorTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _Pledge4Thor.Contract.RenounceOwnership(&_Pledge4Thor.TransactOpts)
}

// SetBridge is a paid mutator transaction binding the contract method 0x8dd14802.
//
// Solidity: function setBridge(address bridge) returns()
func (_Pledge4Thor *Pledge4ThorTransactor) SetBridge(opts *bind.TransactOpts, bridge common.Address) (*types.Transaction, error) {
	return _Pledge4Thor.contract.Transact(opts, "setBridge", bridge)
}

// SetBridge is a paid mutator transaction binding the contract method 0x8dd14802.
//
// Solidity: function setBridge(address bridge) returns()
func (_Pledge4Thor *Pledge4ThorSession) SetBridge(bridge common.Address) (*types.Transaction, error) {
	return _Pledge4Thor.Contract.SetBridge(&_Pledge4Thor.TransactOpts, bridge)
}

// SetBridge is a paid mutator transaction binding the contract method 0x8dd14802.
//
// Solidity: function setBridge(address bridge) returns()
func (_Pledge4Thor *Pledge4ThorTransactorSession) SetBridge(bridge common.Address) (*types.Transaction, error) {
	return _Pledge4Thor.Contract.SetBridge(&_Pledge4Thor.TransactOpts, bridge)
}

// SetFuel is a paid mutator transaction binding the contract method 0x99b96329.
//
// Solidity: function setFuel(uint256 fuel) returns()
func (_Pledge4Thor *Pledge4ThorTransactor) SetFuel(opts *bind.TransactOpts, fuel *big.Int) (*types.Transaction, error) {
	return _Pledge4Thor.contract.Transact(opts, "setFuel", fuel)
}

// SetFuel is a paid mutator transaction binding the contract method 0x99b96329.
//
// Solidity: function setFuel(uint256 fuel) returns()
func (_Pledge4Thor *Pledge4ThorSession) SetFuel(fuel *big.Int) (*types.Transaction, error) {
	return _Pledge4Thor.Contract.SetFuel(&_Pledge4Thor.TransactOpts, fuel)
}

// SetFuel is a paid mutator transaction binding the contract method 0x99b96329.
//
// Solidity: function setFuel(uint256 fuel) returns()
func (_Pledge4Thor *Pledge4ThorTransactorSession) SetFuel(fuel *big.Int) (*types.Transaction, error) {
	return _Pledge4Thor.Contract.SetFuel(&_Pledge4Thor.TransactOpts, fuel)
}

// SetNextFuel is a paid mutator transaction binding the contract method 0xc4cb6971.
//
// Solidity: function setNextFuel(uint256 nextFuel, uint256 nextFuelHeight) returns()
func (_Pledge4Thor *Pledge4ThorTransactor) SetNextFuel(opts *bind.TransactOpts, nextFuel *big.Int, nextFuelHeight *big.Int) (*types.Transaction, error) {
	return _Pledge4Thor.contract.Transact(opts, "setNextFuel", nextFuel, nextFuelHeight)
}

// SetNextFuel is a paid mutator transaction binding the contract method 0xc4cb6971.
//
// Solidity: function setNextFuel(uint256 nextFuel, uint256 nextFuelHeight) returns()
func (_Pledge4Thor *Pledge4ThorSession) SetNextFuel(nextFuel *big.Int, nextFuelHeight *big.Int) (*types.Transaction, error) {
	return _Pledge4Thor.Contract.SetNextFuel(&_Pledge4Thor.TransactOpts, nextFuel, nextFuelHeight)
}

// SetNextFuel is a paid mutator transaction binding the contract method 0xc4cb6971.
//
// Solidity: function setNextFuel(uint256 nextFuel, uint256 nextFuelHeight) returns()
func (_Pledge4Thor *Pledge4ThorTransactorSession) SetNextFuel(nextFuel *big.Int, nextFuelHeight *big.Int) (*types.Transaction, error) {
	return _Pledge4Thor.Contract.SetNextFuel(&_Pledge4Thor.TransactOpts, nextFuel, nextFuelHeight)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Pledge4Thor *Pledge4ThorTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _Pledge4Thor.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Pledge4Thor *Pledge4ThorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Pledge4Thor.Contract.TransferOwnership(&_Pledge4Thor.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Pledge4Thor *Pledge4ThorTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Pledge4Thor.Contract.TransferOwnership(&_Pledge4Thor.TransactOpts, newOwner)
}

// UnBindRig is a paid mutator transaction binding the contract method 0x72d45cb2.
//
// Solidity: function unBindRig(bytes32 rig) returns()
func (_Pledge4Thor *Pledge4ThorTransactor) UnBindRig(opts *bind.TransactOpts, rig [32]byte) (*types.Transaction, error) {
	return _Pledge4Thor.contract.Transact(opts, "unBindRig", rig)
}

// UnBindRig is a paid mutator transaction binding the contract method 0x72d45cb2.
//
// Solidity: function unBindRig(bytes32 rig) returns()
func (_Pledge4Thor *Pledge4ThorSession) UnBindRig(rig [32]byte) (*types.Transaction, error) {
	return _Pledge4Thor.Contract.UnBindRig(&_Pledge4Thor.TransactOpts, rig)
}

// UnBindRig is a paid mutator transaction binding the contract method 0x72d45cb2.
//
// Solidity: function unBindRig(bytes32 rig) returns()
func (_Pledge4Thor *Pledge4ThorTransactorSession) UnBindRig(rig [32]byte) (*types.Transaction, error) {
	return _Pledge4Thor.Contract.UnBindRig(&_Pledge4Thor.TransactOpts, rig)
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns()
func (_Pledge4Thor *Pledge4ThorTransactor) Unpause(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Pledge4Thor.contract.Transact(opts, "unpause")
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns()
func (_Pledge4Thor *Pledge4ThorSession) Unpause() (*types.Transaction, error) {
	return _Pledge4Thor.Contract.Unpause(&_Pledge4Thor.TransactOpts)
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns()
func (_Pledge4Thor *Pledge4ThorTransactorSession) Unpause() (*types.Transaction, error) {
	return _Pledge4Thor.Contract.Unpause(&_Pledge4Thor.TransactOpts)
}

// Withdraw is a paid mutator transaction binding the contract method 0xf3fef3a3.
//
// Solidity: function withdraw(address payee, uint256 amount) returns()
func (_Pledge4Thor *Pledge4ThorTransactor) Withdraw(opts *bind.TransactOpts, payee common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Pledge4Thor.contract.Transact(opts, "withdraw", payee, amount)
}

// Withdraw is a paid mutator transaction binding the contract method 0xf3fef3a3.
//
// Solidity: function withdraw(address payee, uint256 amount) returns()
func (_Pledge4Thor *Pledge4ThorSession) Withdraw(payee common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Pledge4Thor.Contract.Withdraw(&_Pledge4Thor.TransactOpts, payee, amount)
}

// Withdraw is a paid mutator transaction binding the contract method 0xf3fef3a3.
//
// Solidity: function withdraw(address payee, uint256 amount) returns()
func (_Pledge4Thor *Pledge4ThorTransactorSession) Withdraw(payee common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Pledge4Thor.Contract.Withdraw(&_Pledge4Thor.TransactOpts, payee, amount)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_Pledge4Thor *Pledge4ThorTransactor) Receive(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Pledge4Thor.contract.RawTransact(opts, nil) // calldata is disallowed for receive function
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_Pledge4Thor *Pledge4ThorSession) Receive() (*types.Transaction, error) {
	return _Pledge4Thor.Contract.Receive(&_Pledge4Thor.TransactOpts)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_Pledge4Thor *Pledge4ThorTransactorSession) Receive() (*types.Transaction, error) {
	return _Pledge4Thor.Contract.Receive(&_Pledge4Thor.TransactOpts)
}

// Pledge4ThorBindRigIterator is returned from FilterBindRig and is used to iterate over the raw logs and unpacked data for BindRig events raised by the Pledge4Thor contract.
type Pledge4ThorBindRigIterator struct {
	Event *Pledge4ThorBindRig // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *Pledge4ThorBindRigIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(Pledge4ThorBindRig)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(Pledge4ThorBindRig)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *Pledge4ThorBindRigIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *Pledge4ThorBindRigIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// Pledge4ThorBindRig represents a BindRig event raised by the Pledge4Thor contract.
type Pledge4ThorBindRig struct {
	Rig   [32]byte
	Owner common.Address
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterBindRig is a free log retrieval operation binding the contract event 0x50f0ee7af49e6ace8107b6f934a6a53ca37b9a1625d53bed9ab5c9c76faeb5e0.
//
// Solidity: event BindRig(bytes32 rig, address indexed owner)
func (_Pledge4Thor *Pledge4ThorFilterer) FilterBindRig(opts *bind.FilterOpts, owner []common.Address) (*Pledge4ThorBindRigIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}

	logs, sub, err := _Pledge4Thor.contract.FilterLogs(opts, "BindRig", ownerRule)
	if err != nil {
		return nil, err
	}
	return &Pledge4ThorBindRigIterator{contract: _Pledge4Thor.contract, event: "BindRig", logs: logs, sub: sub}, nil
}

// WatchBindRig is a free log subscription operation binding the contract event 0x50f0ee7af49e6ace8107b6f934a6a53ca37b9a1625d53bed9ab5c9c76faeb5e0.
//
// Solidity: event BindRig(bytes32 rig, address indexed owner)
func (_Pledge4Thor *Pledge4ThorFilterer) WatchBindRig(opts *bind.WatchOpts, sink chan<- *Pledge4ThorBindRig, owner []common.Address) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}

	logs, sub, err := _Pledge4Thor.contract.WatchLogs(opts, "BindRig", ownerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(Pledge4ThorBindRig)
				if err := _Pledge4Thor.contract.UnpackLog(event, "BindRig", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseBindRig is a log parse operation binding the contract event 0x50f0ee7af49e6ace8107b6f934a6a53ca37b9a1625d53bed9ab5c9c76faeb5e0.
//
// Solidity: event BindRig(bytes32 rig, address indexed owner)
func (_Pledge4Thor *Pledge4ThorFilterer) ParseBindRig(log types.Log) (*Pledge4ThorBindRig, error) {
	event := new(Pledge4ThorBindRig)
	if err := _Pledge4Thor.contract.UnpackLog(event, "BindRig", log); err != nil {
		return nil, err
	}
	return event, nil
}

// Pledge4ThorDeductDepositIterator is returned from FilterDeductDeposit and is used to iterate over the raw logs and unpacked data for DeductDeposit events raised by the Pledge4Thor contract.
type Pledge4ThorDeductDepositIterator struct {
	Event *Pledge4ThorDeductDeposit // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *Pledge4ThorDeductDepositIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(Pledge4ThorDeductDeposit)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(Pledge4ThorDeductDeposit)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *Pledge4ThorDeductDepositIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *Pledge4ThorDeductDepositIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// Pledge4ThorDeductDeposit represents a DeductDeposit event raised by the Pledge4Thor contract.
type Pledge4ThorDeductDeposit struct {
	Rig    [32]byte
	Height *big.Int
	Owner  common.Address
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterDeductDeposit is a free log retrieval operation binding the contract event 0x741d859d6fdf97ec7399b68da994d7aa873c4f189cf3a5ef58310885f235fa4f.
//
// Solidity: event DeductDeposit(bytes32 rig, uint256 height, address indexed owner)
func (_Pledge4Thor *Pledge4ThorFilterer) FilterDeductDeposit(opts *bind.FilterOpts, owner []common.Address) (*Pledge4ThorDeductDepositIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}

	logs, sub, err := _Pledge4Thor.contract.FilterLogs(opts, "DeductDeposit", ownerRule)
	if err != nil {
		return nil, err
	}
	return &Pledge4ThorDeductDepositIterator{contract: _Pledge4Thor.contract, event: "DeductDeposit", logs: logs, sub: sub}, nil
}

// WatchDeductDeposit is a free log subscription operation binding the contract event 0x741d859d6fdf97ec7399b68da994d7aa873c4f189cf3a5ef58310885f235fa4f.
//
// Solidity: event DeductDeposit(bytes32 rig, uint256 height, address indexed owner)
func (_Pledge4Thor *Pledge4ThorFilterer) WatchDeductDeposit(opts *bind.WatchOpts, sink chan<- *Pledge4ThorDeductDeposit, owner []common.Address) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}

	logs, sub, err := _Pledge4Thor.contract.WatchLogs(opts, "DeductDeposit", ownerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(Pledge4ThorDeductDeposit)
				if err := _Pledge4Thor.contract.UnpackLog(event, "DeductDeposit", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseDeductDeposit is a log parse operation binding the contract event 0x741d859d6fdf97ec7399b68da994d7aa873c4f189cf3a5ef58310885f235fa4f.
//
// Solidity: event DeductDeposit(bytes32 rig, uint256 height, address indexed owner)
func (_Pledge4Thor *Pledge4ThorFilterer) ParseDeductDeposit(log types.Log) (*Pledge4ThorDeductDeposit, error) {
	event := new(Pledge4ThorDeductDeposit)
	if err := _Pledge4Thor.contract.UnpackLog(event, "DeductDeposit", log); err != nil {
		return nil, err
	}
	return event, nil
}

// Pledge4ThorDepositIterator is returned from FilterDeposit and is used to iterate over the raw logs and unpacked data for Deposit events raised by the Pledge4Thor contract.
type Pledge4ThorDepositIterator struct {
	Event *Pledge4ThorDeposit // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *Pledge4ThorDepositIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(Pledge4ThorDeposit)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(Pledge4ThorDeposit)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *Pledge4ThorDepositIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *Pledge4ThorDepositIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// Pledge4ThorDeposit represents a Deposit event raised by the Pledge4Thor contract.
type Pledge4ThorDeposit struct {
	From   common.Address
	Amount *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterDeposit is a free log retrieval operation binding the contract event 0xe1fffcc4923d04b559f4d29a8bfc6cda04eb5b0d3c460751c2402c5c5cc9109c.
//
// Solidity: event Deposit(address indexed from, uint256 amount)
func (_Pledge4Thor *Pledge4ThorFilterer) FilterDeposit(opts *bind.FilterOpts, from []common.Address) (*Pledge4ThorDepositIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}

	logs, sub, err := _Pledge4Thor.contract.FilterLogs(opts, "Deposit", fromRule)
	if err != nil {
		return nil, err
	}
	return &Pledge4ThorDepositIterator{contract: _Pledge4Thor.contract, event: "Deposit", logs: logs, sub: sub}, nil
}

// WatchDeposit is a free log subscription operation binding the contract event 0xe1fffcc4923d04b559f4d29a8bfc6cda04eb5b0d3c460751c2402c5c5cc9109c.
//
// Solidity: event Deposit(address indexed from, uint256 amount)
func (_Pledge4Thor *Pledge4ThorFilterer) WatchDeposit(opts *bind.WatchOpts, sink chan<- *Pledge4ThorDeposit, from []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}

	logs, sub, err := _Pledge4Thor.contract.WatchLogs(opts, "Deposit", fromRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(Pledge4ThorDeposit)
				if err := _Pledge4Thor.contract.UnpackLog(event, "Deposit", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseDeposit is a log parse operation binding the contract event 0xe1fffcc4923d04b559f4d29a8bfc6cda04eb5b0d3c460751c2402c5c5cc9109c.
//
// Solidity: event Deposit(address indexed from, uint256 amount)
func (_Pledge4Thor *Pledge4ThorFilterer) ParseDeposit(log types.Log) (*Pledge4ThorDeposit, error) {
	event := new(Pledge4ThorDeposit)
	if err := _Pledge4Thor.contract.UnpackLog(event, "Deposit", log); err != nil {
		return nil, err
	}
	return event, nil
}

// Pledge4ThorInternalTransferIterator is returned from FilterInternalTransfer and is used to iterate over the raw logs and unpacked data for InternalTransfer events raised by the Pledge4Thor contract.
type Pledge4ThorInternalTransferIterator struct {
	Event *Pledge4ThorInternalTransfer // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *Pledge4ThorInternalTransferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(Pledge4ThorInternalTransfer)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(Pledge4ThorInternalTransfer)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *Pledge4ThorInternalTransferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *Pledge4ThorInternalTransferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// Pledge4ThorInternalTransfer represents a InternalTransfer event raised by the Pledge4Thor contract.
type Pledge4ThorInternalTransfer struct {
	From   common.Address
	To     common.Address
	Amount *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterInternalTransfer is a free log retrieval operation binding the contract event 0xe2080c8fc8d86c864d8dc081fadaebf2be7191086615e786f954420f13ed122a.
//
// Solidity: event InternalTransfer(address indexed from, address indexed to, uint256 amount)
func (_Pledge4Thor *Pledge4ThorFilterer) FilterInternalTransfer(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*Pledge4ThorInternalTransferIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _Pledge4Thor.contract.FilterLogs(opts, "InternalTransfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &Pledge4ThorInternalTransferIterator{contract: _Pledge4Thor.contract, event: "InternalTransfer", logs: logs, sub: sub}, nil
}

// WatchInternalTransfer is a free log subscription operation binding the contract event 0xe2080c8fc8d86c864d8dc081fadaebf2be7191086615e786f954420f13ed122a.
//
// Solidity: event InternalTransfer(address indexed from, address indexed to, uint256 amount)
func (_Pledge4Thor *Pledge4ThorFilterer) WatchInternalTransfer(opts *bind.WatchOpts, sink chan<- *Pledge4ThorInternalTransfer, from []common.Address, to []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _Pledge4Thor.contract.WatchLogs(opts, "InternalTransfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(Pledge4ThorInternalTransfer)
				if err := _Pledge4Thor.contract.UnpackLog(event, "InternalTransfer", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseInternalTransfer is a log parse operation binding the contract event 0xe2080c8fc8d86c864d8dc081fadaebf2be7191086615e786f954420f13ed122a.
//
// Solidity: event InternalTransfer(address indexed from, address indexed to, uint256 amount)
func (_Pledge4Thor *Pledge4ThorFilterer) ParseInternalTransfer(log types.Log) (*Pledge4ThorInternalTransfer, error) {
	event := new(Pledge4ThorInternalTransfer)
	if err := _Pledge4Thor.contract.UnpackLog(event, "InternalTransfer", log); err != nil {
		return nil, err
	}
	return event, nil
}

// Pledge4ThorOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the Pledge4Thor contract.
type Pledge4ThorOwnershipTransferredIterator struct {
	Event *Pledge4ThorOwnershipTransferred // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *Pledge4ThorOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(Pledge4ThorOwnershipTransferred)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(Pledge4ThorOwnershipTransferred)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *Pledge4ThorOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *Pledge4ThorOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// Pledge4ThorOwnershipTransferred represents a OwnershipTransferred event raised by the Pledge4Thor contract.
type Pledge4ThorOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Pledge4Thor *Pledge4ThorFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*Pledge4ThorOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Pledge4Thor.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &Pledge4ThorOwnershipTransferredIterator{contract: _Pledge4Thor.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Pledge4Thor *Pledge4ThorFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *Pledge4ThorOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Pledge4Thor.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(Pledge4ThorOwnershipTransferred)
				if err := _Pledge4Thor.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseOwnershipTransferred is a log parse operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Pledge4Thor *Pledge4ThorFilterer) ParseOwnershipTransferred(log types.Log) (*Pledge4ThorOwnershipTransferred, error) {
	event := new(Pledge4ThorOwnershipTransferred)
	if err := _Pledge4Thor.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	return event, nil
}

// Pledge4ThorPausedIterator is returned from FilterPaused and is used to iterate over the raw logs and unpacked data for Paused events raised by the Pledge4Thor contract.
type Pledge4ThorPausedIterator struct {
	Event *Pledge4ThorPaused // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *Pledge4ThorPausedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(Pledge4ThorPaused)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(Pledge4ThorPaused)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *Pledge4ThorPausedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *Pledge4ThorPausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// Pledge4ThorPaused represents a Paused event raised by the Pledge4Thor contract.
type Pledge4ThorPaused struct {
	Account common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterPaused is a free log retrieval operation binding the contract event 0x62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a258.
//
// Solidity: event Paused(address account)
func (_Pledge4Thor *Pledge4ThorFilterer) FilterPaused(opts *bind.FilterOpts) (*Pledge4ThorPausedIterator, error) {

	logs, sub, err := _Pledge4Thor.contract.FilterLogs(opts, "Paused")
	if err != nil {
		return nil, err
	}
	return &Pledge4ThorPausedIterator{contract: _Pledge4Thor.contract, event: "Paused", logs: logs, sub: sub}, nil
}

// WatchPaused is a free log subscription operation binding the contract event 0x62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a258.
//
// Solidity: event Paused(address account)
func (_Pledge4Thor *Pledge4ThorFilterer) WatchPaused(opts *bind.WatchOpts, sink chan<- *Pledge4ThorPaused) (event.Subscription, error) {

	logs, sub, err := _Pledge4Thor.contract.WatchLogs(opts, "Paused")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(Pledge4ThorPaused)
				if err := _Pledge4Thor.contract.UnpackLog(event, "Paused", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParsePaused is a log parse operation binding the contract event 0x62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a258.
//
// Solidity: event Paused(address account)
func (_Pledge4Thor *Pledge4ThorFilterer) ParsePaused(log types.Log) (*Pledge4ThorPaused, error) {
	event := new(Pledge4ThorPaused)
	if err := _Pledge4Thor.contract.UnpackLog(event, "Paused", log); err != nil {
		return nil, err
	}
	return event, nil
}

// Pledge4ThorUnBindRigIterator is returned from FilterUnBindRig and is used to iterate over the raw logs and unpacked data for UnBindRig events raised by the Pledge4Thor contract.
type Pledge4ThorUnBindRigIterator struct {
	Event *Pledge4ThorUnBindRig // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *Pledge4ThorUnBindRigIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(Pledge4ThorUnBindRig)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(Pledge4ThorUnBindRig)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *Pledge4ThorUnBindRigIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *Pledge4ThorUnBindRigIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// Pledge4ThorUnBindRig represents a UnBindRig event raised by the Pledge4Thor contract.
type Pledge4ThorUnBindRig struct {
	Rig   [32]byte
	Owner common.Address
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterUnBindRig is a free log retrieval operation binding the contract event 0x7310aed68a6f42bbcee631bcb319c0eead9b2f286ee39d57183a7460a0193f10.
//
// Solidity: event UnBindRig(bytes32 rig, address indexed owner)
func (_Pledge4Thor *Pledge4ThorFilterer) FilterUnBindRig(opts *bind.FilterOpts, owner []common.Address) (*Pledge4ThorUnBindRigIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}

	logs, sub, err := _Pledge4Thor.contract.FilterLogs(opts, "UnBindRig", ownerRule)
	if err != nil {
		return nil, err
	}
	return &Pledge4ThorUnBindRigIterator{contract: _Pledge4Thor.contract, event: "UnBindRig", logs: logs, sub: sub}, nil
}

// WatchUnBindRig is a free log subscription operation binding the contract event 0x7310aed68a6f42bbcee631bcb319c0eead9b2f286ee39d57183a7460a0193f10.
//
// Solidity: event UnBindRig(bytes32 rig, address indexed owner)
func (_Pledge4Thor *Pledge4ThorFilterer) WatchUnBindRig(opts *bind.WatchOpts, sink chan<- *Pledge4ThorUnBindRig, owner []common.Address) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}

	logs, sub, err := _Pledge4Thor.contract.WatchLogs(opts, "UnBindRig", ownerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(Pledge4ThorUnBindRig)
				if err := _Pledge4Thor.contract.UnpackLog(event, "UnBindRig", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseUnBindRig is a log parse operation binding the contract event 0x7310aed68a6f42bbcee631bcb319c0eead9b2f286ee39d57183a7460a0193f10.
//
// Solidity: event UnBindRig(bytes32 rig, address indexed owner)
func (_Pledge4Thor *Pledge4ThorFilterer) ParseUnBindRig(log types.Log) (*Pledge4ThorUnBindRig, error) {
	event := new(Pledge4ThorUnBindRig)
	if err := _Pledge4Thor.contract.UnpackLog(event, "UnBindRig", log); err != nil {
		return nil, err
	}
	return event, nil
}

// Pledge4ThorUnpausedIterator is returned from FilterUnpaused and is used to iterate over the raw logs and unpacked data for Unpaused events raised by the Pledge4Thor contract.
type Pledge4ThorUnpausedIterator struct {
	Event *Pledge4ThorUnpaused // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *Pledge4ThorUnpausedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(Pledge4ThorUnpaused)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(Pledge4ThorUnpaused)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *Pledge4ThorUnpausedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *Pledge4ThorUnpausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// Pledge4ThorUnpaused represents a Unpaused event raised by the Pledge4Thor contract.
type Pledge4ThorUnpaused struct {
	Account common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterUnpaused is a free log retrieval operation binding the contract event 0x5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa.
//
// Solidity: event Unpaused(address account)
func (_Pledge4Thor *Pledge4ThorFilterer) FilterUnpaused(opts *bind.FilterOpts) (*Pledge4ThorUnpausedIterator, error) {

	logs, sub, err := _Pledge4Thor.contract.FilterLogs(opts, "Unpaused")
	if err != nil {
		return nil, err
	}
	return &Pledge4ThorUnpausedIterator{contract: _Pledge4Thor.contract, event: "Unpaused", logs: logs, sub: sub}, nil
}

// WatchUnpaused is a free log subscription operation binding the contract event 0x5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa.
//
// Solidity: event Unpaused(address account)
func (_Pledge4Thor *Pledge4ThorFilterer) WatchUnpaused(opts *bind.WatchOpts, sink chan<- *Pledge4ThorUnpaused) (event.Subscription, error) {

	logs, sub, err := _Pledge4Thor.contract.WatchLogs(opts, "Unpaused")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(Pledge4ThorUnpaused)
				if err := _Pledge4Thor.contract.UnpackLog(event, "Unpaused", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseUnpaused is a log parse operation binding the contract event 0x5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa.
//
// Solidity: event Unpaused(address account)
func (_Pledge4Thor *Pledge4ThorFilterer) ParseUnpaused(log types.Log) (*Pledge4ThorUnpaused, error) {
	event := new(Pledge4ThorUnpaused)
	if err := _Pledge4Thor.contract.UnpackLog(event, "Unpaused", log); err != nil {
		return nil, err
	}
	return event, nil
}

// Pledge4ThorWithdrawIterator is returned from FilterWithdraw and is used to iterate over the raw logs and unpacked data for Withdraw events raised by the Pledge4Thor contract.
type Pledge4ThorWithdrawIterator struct {
	Event *Pledge4ThorWithdraw // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *Pledge4ThorWithdrawIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(Pledge4ThorWithdraw)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(Pledge4ThorWithdraw)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *Pledge4ThorWithdrawIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *Pledge4ThorWithdrawIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// Pledge4ThorWithdraw represents a Withdraw event raised by the Pledge4Thor contract.
type Pledge4ThorWithdraw struct {
	To     common.Address
	Amount *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterWithdraw is a free log retrieval operation binding the contract event 0x884edad9ce6fa2440d8a54cc123490eb96d2768479d49ff9c7366125a9424364.
//
// Solidity: event Withdraw(address indexed to, uint256 amount)
func (_Pledge4Thor *Pledge4ThorFilterer) FilterWithdraw(opts *bind.FilterOpts, to []common.Address) (*Pledge4ThorWithdrawIterator, error) {

	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _Pledge4Thor.contract.FilterLogs(opts, "Withdraw", toRule)
	if err != nil {
		return nil, err
	}
	return &Pledge4ThorWithdrawIterator{contract: _Pledge4Thor.contract, event: "Withdraw", logs: logs, sub: sub}, nil
}

// WatchWithdraw is a free log subscription operation binding the contract event 0x884edad9ce6fa2440d8a54cc123490eb96d2768479d49ff9c7366125a9424364.
//
// Solidity: event Withdraw(address indexed to, uint256 amount)
func (_Pledge4Thor *Pledge4ThorFilterer) WatchWithdraw(opts *bind.WatchOpts, sink chan<- *Pledge4ThorWithdraw, to []common.Address) (event.Subscription, error) {

	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _Pledge4Thor.contract.WatchLogs(opts, "Withdraw", toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(Pledge4ThorWithdraw)
				if err := _Pledge4Thor.contract.UnpackLog(event, "Withdraw", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseWithdraw is a log parse operation binding the contract event 0x884edad9ce6fa2440d8a54cc123490eb96d2768479d49ff9c7366125a9424364.
//
// Solidity: event Withdraw(address indexed to, uint256 amount)
func (_Pledge4Thor *Pledge4ThorFilterer) ParseWithdraw(log types.Log) (*Pledge4ThorWithdraw, error) {
	event := new(Pledge4ThorWithdraw)
	if err := _Pledge4Thor.contract.UnpackLog(event, "Withdraw", log); err != nil {
		return nil, err
	}
	return event, nil
}
