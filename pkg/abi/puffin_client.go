// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package abi

import (
	"errors"
	"math/big"
	"strings"

	ethereum "github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/event"
)

// Reference imports to suppress errors if they are not otherwise used.
var (
	_ = errors.New
	_ = big.NewInt
	_ = strings.NewReader
	_ = ethereum.NotFound
	_ = bind.Bind
	_ = common.Big1
	_ = types.BloomLookup
	_ = event.NewSubscription
)

// PuffinClientMetaData contains all meta data concerning the PuffinClient contract.
var PuffinClientMetaData = &bind.MetaData{
	ABI: "[{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"epoch\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"timestamp\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"users\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_usdcPerUser\",\"type\":\"uint256\"}],\"name\":\"NewEpoch\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"epoch\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"timestamp\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"users\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_usdcPerUser\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"totalPaid\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"remainingBalance\",\"type\":\"uint256\"}],\"name\":\"NewPayment\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"calculatePayout\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_payToAddress\",\"type\":\"address\"}],\"name\":\"changePayToAddress\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_usdc\",\"type\":\"address\"}],\"name\":\"changeUSDC\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_usdPerUser\",\"type\":\"uint256\"}],\"name\":\"changeUSDPerUser\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"delUser\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"epoch\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"epochUsers\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"isCurrent\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_user\",\"type\":\"address\"}],\"name\":\"newUser\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"nextEpoch\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"payToAddress\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"payout\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"usdc\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"usdcPerUser\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
}

// PuffinClientABI is the input ABI used to generate the binding from.
// Deprecated: Use PuffinClientMetaData.ABI instead.
var PuffinClientABI = PuffinClientMetaData.ABI

// PuffinClient is an auto generated Go binding around an Ethereum contract.
type PuffinClient struct {
	PuffinClientCaller     // Read-only binding to the contract
	PuffinClientTransactor // Write-only binding to the contract
	PuffinClientFilterer   // Log filterer for contract events
}

// PuffinClientCaller is an auto generated read-only Go binding around an Ethereum contract.
type PuffinClientCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PuffinClientTransactor is an auto generated write-only Go binding around an Ethereum contract.
type PuffinClientTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PuffinClientFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type PuffinClientFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PuffinClientSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type PuffinClientSession struct {
	Contract     *PuffinClient     // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// PuffinClientCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type PuffinClientCallerSession struct {
	Contract *PuffinClientCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts       // Call options to use throughout this session
}

// PuffinClientTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type PuffinClientTransactorSession struct {
	Contract     *PuffinClientTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts       // Transaction auth options to use throughout this session
}

// PuffinClientRaw is an auto generated low-level Go binding around an Ethereum contract.
type PuffinClientRaw struct {
	Contract *PuffinClient // Generic contract binding to access the raw methods on
}

// PuffinClientCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type PuffinClientCallerRaw struct {
	Contract *PuffinClientCaller // Generic read-only contract binding to access the raw methods on
}

// PuffinClientTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type PuffinClientTransactorRaw struct {
	Contract *PuffinClientTransactor // Generic write-only contract binding to access the raw methods on
}

// NewPuffinClient creates a new instance of PuffinClient, bound to a specific deployed contract.
func NewPuffinClient(address common.Address, backend bind.ContractBackend) (*PuffinClient, error) {
	contract, err := bindPuffinClient(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &PuffinClient{PuffinClientCaller: PuffinClientCaller{contract: contract}, PuffinClientTransactor: PuffinClientTransactor{contract: contract}, PuffinClientFilterer: PuffinClientFilterer{contract: contract}}, nil
}

// NewPuffinClientCaller creates a new read-only instance of PuffinClient, bound to a specific deployed contract.
func NewPuffinClientCaller(address common.Address, caller bind.ContractCaller) (*PuffinClientCaller, error) {
	contract, err := bindPuffinClient(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &PuffinClientCaller{contract: contract}, nil
}

// NewPuffinClientTransactor creates a new write-only instance of PuffinClient, bound to a specific deployed contract.
func NewPuffinClientTransactor(address common.Address, transactor bind.ContractTransactor) (*PuffinClientTransactor, error) {
	contract, err := bindPuffinClient(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &PuffinClientTransactor{contract: contract}, nil
}

// NewPuffinClientFilterer creates a new log filterer instance of PuffinClient, bound to a specific deployed contract.
func NewPuffinClientFilterer(address common.Address, filterer bind.ContractFilterer) (*PuffinClientFilterer, error) {
	contract, err := bindPuffinClient(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &PuffinClientFilterer{contract: contract}, nil
}

// bindPuffinClient binds a generic wrapper to an already deployed contract.
func bindPuffinClient(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(PuffinClientABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_PuffinClient *PuffinClientRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _PuffinClient.Contract.PuffinClientCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_PuffinClient *PuffinClientRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _PuffinClient.Contract.PuffinClientTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_PuffinClient *PuffinClientRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _PuffinClient.Contract.PuffinClientTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_PuffinClient *PuffinClientCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _PuffinClient.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_PuffinClient *PuffinClientTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _PuffinClient.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_PuffinClient *PuffinClientTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _PuffinClient.Contract.contract.Transact(opts, method, params...)
}

// CalculatePayout is a free data retrieval call binding the contract method 0x83c64ec4.
//
// Solidity: function calculatePayout() view returns(uint256)
func (_PuffinClient *PuffinClientCaller) CalculatePayout(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _PuffinClient.contract.Call(opts, &out, "calculatePayout")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// CalculatePayout is a free data retrieval call binding the contract method 0x83c64ec4.
//
// Solidity: function calculatePayout() view returns(uint256)
func (_PuffinClient *PuffinClientSession) CalculatePayout() (*big.Int, error) {
	return _PuffinClient.Contract.CalculatePayout(&_PuffinClient.CallOpts)
}

// CalculatePayout is a free data retrieval call binding the contract method 0x83c64ec4.
//
// Solidity: function calculatePayout() view returns(uint256)
func (_PuffinClient *PuffinClientCallerSession) CalculatePayout() (*big.Int, error) {
	return _PuffinClient.Contract.CalculatePayout(&_PuffinClient.CallOpts)
}

// Epoch is a free data retrieval call binding the contract method 0x900cf0cf.
//
// Solidity: function epoch() view returns(uint256)
func (_PuffinClient *PuffinClientCaller) Epoch(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _PuffinClient.contract.Call(opts, &out, "epoch")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Epoch is a free data retrieval call binding the contract method 0x900cf0cf.
//
// Solidity: function epoch() view returns(uint256)
func (_PuffinClient *PuffinClientSession) Epoch() (*big.Int, error) {
	return _PuffinClient.Contract.Epoch(&_PuffinClient.CallOpts)
}

// Epoch is a free data retrieval call binding the contract method 0x900cf0cf.
//
// Solidity: function epoch() view returns(uint256)
func (_PuffinClient *PuffinClientCallerSession) Epoch() (*big.Int, error) {
	return _PuffinClient.Contract.Epoch(&_PuffinClient.CallOpts)
}

// EpochUsers is a free data retrieval call binding the contract method 0xb4e03125.
//
// Solidity: function epochUsers(uint256 ) view returns(uint256)
func (_PuffinClient *PuffinClientCaller) EpochUsers(opts *bind.CallOpts, arg0 *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _PuffinClient.contract.Call(opts, &out, "epochUsers", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// EpochUsers is a free data retrieval call binding the contract method 0xb4e03125.
//
// Solidity: function epochUsers(uint256 ) view returns(uint256)
func (_PuffinClient *PuffinClientSession) EpochUsers(arg0 *big.Int) (*big.Int, error) {
	return _PuffinClient.Contract.EpochUsers(&_PuffinClient.CallOpts, arg0)
}

// EpochUsers is a free data retrieval call binding the contract method 0xb4e03125.
//
// Solidity: function epochUsers(uint256 ) view returns(uint256)
func (_PuffinClient *PuffinClientCallerSession) EpochUsers(arg0 *big.Int) (*big.Int, error) {
	return _PuffinClient.Contract.EpochUsers(&_PuffinClient.CallOpts, arg0)
}

// IsCurrent is a free data retrieval call binding the contract method 0x3d100d3f.
//
// Solidity: function isCurrent() view returns(bool)
func (_PuffinClient *PuffinClientCaller) IsCurrent(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _PuffinClient.contract.Call(opts, &out, "isCurrent")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsCurrent is a free data retrieval call binding the contract method 0x3d100d3f.
//
// Solidity: function isCurrent() view returns(bool)
func (_PuffinClient *PuffinClientSession) IsCurrent() (bool, error) {
	return _PuffinClient.Contract.IsCurrent(&_PuffinClient.CallOpts)
}

// IsCurrent is a free data retrieval call binding the contract method 0x3d100d3f.
//
// Solidity: function isCurrent() view returns(bool)
func (_PuffinClient *PuffinClientCallerSession) IsCurrent() (bool, error) {
	return _PuffinClient.Contract.IsCurrent(&_PuffinClient.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_PuffinClient *PuffinClientCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _PuffinClient.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_PuffinClient *PuffinClientSession) Owner() (common.Address, error) {
	return _PuffinClient.Contract.Owner(&_PuffinClient.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_PuffinClient *PuffinClientCallerSession) Owner() (common.Address, error) {
	return _PuffinClient.Contract.Owner(&_PuffinClient.CallOpts)
}

// PayToAddress is a free data retrieval call binding the contract method 0xbe11fdc7.
//
// Solidity: function payToAddress() view returns(address)
func (_PuffinClient *PuffinClientCaller) PayToAddress(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _PuffinClient.contract.Call(opts, &out, "payToAddress")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// PayToAddress is a free data retrieval call binding the contract method 0xbe11fdc7.
//
// Solidity: function payToAddress() view returns(address)
func (_PuffinClient *PuffinClientSession) PayToAddress() (common.Address, error) {
	return _PuffinClient.Contract.PayToAddress(&_PuffinClient.CallOpts)
}

// PayToAddress is a free data retrieval call binding the contract method 0xbe11fdc7.
//
// Solidity: function payToAddress() view returns(address)
func (_PuffinClient *PuffinClientCallerSession) PayToAddress() (common.Address, error) {
	return _PuffinClient.Contract.PayToAddress(&_PuffinClient.CallOpts)
}

// Usdc is a free data retrieval call binding the contract method 0x3e413bee.
//
// Solidity: function usdc() view returns(address)
func (_PuffinClient *PuffinClientCaller) Usdc(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _PuffinClient.contract.Call(opts, &out, "usdc")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Usdc is a free data retrieval call binding the contract method 0x3e413bee.
//
// Solidity: function usdc() view returns(address)
func (_PuffinClient *PuffinClientSession) Usdc() (common.Address, error) {
	return _PuffinClient.Contract.Usdc(&_PuffinClient.CallOpts)
}

// Usdc is a free data retrieval call binding the contract method 0x3e413bee.
//
// Solidity: function usdc() view returns(address)
func (_PuffinClient *PuffinClientCallerSession) Usdc() (common.Address, error) {
	return _PuffinClient.Contract.Usdc(&_PuffinClient.CallOpts)
}

// UsdcPerUser is a free data retrieval call binding the contract method 0xbd8be057.
//
// Solidity: function usdcPerUser() view returns(uint256)
func (_PuffinClient *PuffinClientCaller) UsdcPerUser(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _PuffinClient.contract.Call(opts, &out, "usdcPerUser")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// UsdcPerUser is a free data retrieval call binding the contract method 0xbd8be057.
//
// Solidity: function usdcPerUser() view returns(uint256)
func (_PuffinClient *PuffinClientSession) UsdcPerUser() (*big.Int, error) {
	return _PuffinClient.Contract.UsdcPerUser(&_PuffinClient.CallOpts)
}

// UsdcPerUser is a free data retrieval call binding the contract method 0xbd8be057.
//
// Solidity: function usdcPerUser() view returns(uint256)
func (_PuffinClient *PuffinClientCallerSession) UsdcPerUser() (*big.Int, error) {
	return _PuffinClient.Contract.UsdcPerUser(&_PuffinClient.CallOpts)
}

// ChangePayToAddress is a paid mutator transaction binding the contract method 0x5cb39e00.
//
// Solidity: function changePayToAddress(address _payToAddress) returns()
func (_PuffinClient *PuffinClientTransactor) ChangePayToAddress(opts *bind.TransactOpts, _payToAddress common.Address) (*types.Transaction, error) {
	return _PuffinClient.contract.Transact(opts, "changePayToAddress", _payToAddress)
}

// ChangePayToAddress is a paid mutator transaction binding the contract method 0x5cb39e00.
//
// Solidity: function changePayToAddress(address _payToAddress) returns()
func (_PuffinClient *PuffinClientSession) ChangePayToAddress(_payToAddress common.Address) (*types.Transaction, error) {
	return _PuffinClient.Contract.ChangePayToAddress(&_PuffinClient.TransactOpts, _payToAddress)
}

// ChangePayToAddress is a paid mutator transaction binding the contract method 0x5cb39e00.
//
// Solidity: function changePayToAddress(address _payToAddress) returns()
func (_PuffinClient *PuffinClientTransactorSession) ChangePayToAddress(_payToAddress common.Address) (*types.Transaction, error) {
	return _PuffinClient.Contract.ChangePayToAddress(&_PuffinClient.TransactOpts, _payToAddress)
}

// ChangeUSDC is a paid mutator transaction binding the contract method 0x1ff415d9.
//
// Solidity: function changeUSDC(address _usdc) returns()
func (_PuffinClient *PuffinClientTransactor) ChangeUSDC(opts *bind.TransactOpts, _usdc common.Address) (*types.Transaction, error) {
	return _PuffinClient.contract.Transact(opts, "changeUSDC", _usdc)
}

// ChangeUSDC is a paid mutator transaction binding the contract method 0x1ff415d9.
//
// Solidity: function changeUSDC(address _usdc) returns()
func (_PuffinClient *PuffinClientSession) ChangeUSDC(_usdc common.Address) (*types.Transaction, error) {
	return _PuffinClient.Contract.ChangeUSDC(&_PuffinClient.TransactOpts, _usdc)
}

// ChangeUSDC is a paid mutator transaction binding the contract method 0x1ff415d9.
//
// Solidity: function changeUSDC(address _usdc) returns()
func (_PuffinClient *PuffinClientTransactorSession) ChangeUSDC(_usdc common.Address) (*types.Transaction, error) {
	return _PuffinClient.Contract.ChangeUSDC(&_PuffinClient.TransactOpts, _usdc)
}

// ChangeUSDPerUser is a paid mutator transaction binding the contract method 0x9dde76fe.
//
// Solidity: function changeUSDPerUser(uint256 _usdPerUser) returns()
func (_PuffinClient *PuffinClientTransactor) ChangeUSDPerUser(opts *bind.TransactOpts, _usdPerUser *big.Int) (*types.Transaction, error) {
	return _PuffinClient.contract.Transact(opts, "changeUSDPerUser", _usdPerUser)
}

// ChangeUSDPerUser is a paid mutator transaction binding the contract method 0x9dde76fe.
//
// Solidity: function changeUSDPerUser(uint256 _usdPerUser) returns()
func (_PuffinClient *PuffinClientSession) ChangeUSDPerUser(_usdPerUser *big.Int) (*types.Transaction, error) {
	return _PuffinClient.Contract.ChangeUSDPerUser(&_PuffinClient.TransactOpts, _usdPerUser)
}

// ChangeUSDPerUser is a paid mutator transaction binding the contract method 0x9dde76fe.
//
// Solidity: function changeUSDPerUser(uint256 _usdPerUser) returns()
func (_PuffinClient *PuffinClientTransactorSession) ChangeUSDPerUser(_usdPerUser *big.Int) (*types.Transaction, error) {
	return _PuffinClient.Contract.ChangeUSDPerUser(&_PuffinClient.TransactOpts, _usdPerUser)
}

// DelUser is a paid mutator transaction binding the contract method 0x899478a6.
//
// Solidity: function delUser() returns()
func (_PuffinClient *PuffinClientTransactor) DelUser(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _PuffinClient.contract.Transact(opts, "delUser")
}

// DelUser is a paid mutator transaction binding the contract method 0x899478a6.
//
// Solidity: function delUser() returns()
func (_PuffinClient *PuffinClientSession) DelUser() (*types.Transaction, error) {
	return _PuffinClient.Contract.DelUser(&_PuffinClient.TransactOpts)
}

// DelUser is a paid mutator transaction binding the contract method 0x899478a6.
//
// Solidity: function delUser() returns()
func (_PuffinClient *PuffinClientTransactorSession) DelUser() (*types.Transaction, error) {
	return _PuffinClient.Contract.DelUser(&_PuffinClient.TransactOpts)
}

// NewUser is a paid mutator transaction binding the contract method 0x0cb5f653.
//
// Solidity: function newUser(address _user) returns()
func (_PuffinClient *PuffinClientTransactor) NewUser(opts *bind.TransactOpts, _user common.Address) (*types.Transaction, error) {
	return _PuffinClient.contract.Transact(opts, "newUser", _user)
}

// NewUser is a paid mutator transaction binding the contract method 0x0cb5f653.
//
// Solidity: function newUser(address _user) returns()
func (_PuffinClient *PuffinClientSession) NewUser(_user common.Address) (*types.Transaction, error) {
	return _PuffinClient.Contract.NewUser(&_PuffinClient.TransactOpts, _user)
}

// NewUser is a paid mutator transaction binding the contract method 0x0cb5f653.
//
// Solidity: function newUser(address _user) returns()
func (_PuffinClient *PuffinClientTransactorSession) NewUser(_user common.Address) (*types.Transaction, error) {
	return _PuffinClient.Contract.NewUser(&_PuffinClient.TransactOpts, _user)
}

// NextEpoch is a paid mutator transaction binding the contract method 0xaea0e78b.
//
// Solidity: function nextEpoch() returns()
func (_PuffinClient *PuffinClientTransactor) NextEpoch(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _PuffinClient.contract.Transact(opts, "nextEpoch")
}

// NextEpoch is a paid mutator transaction binding the contract method 0xaea0e78b.
//
// Solidity: function nextEpoch() returns()
func (_PuffinClient *PuffinClientSession) NextEpoch() (*types.Transaction, error) {
	return _PuffinClient.Contract.NextEpoch(&_PuffinClient.TransactOpts)
}

// NextEpoch is a paid mutator transaction binding the contract method 0xaea0e78b.
//
// Solidity: function nextEpoch() returns()
func (_PuffinClient *PuffinClientTransactorSession) NextEpoch() (*types.Transaction, error) {
	return _PuffinClient.Contract.NextEpoch(&_PuffinClient.TransactOpts)
}

// Payout is a paid mutator transaction binding the contract method 0x63bd1d4a.
//
// Solidity: function payout() returns()
func (_PuffinClient *PuffinClientTransactor) Payout(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _PuffinClient.contract.Transact(opts, "payout")
}

// Payout is a paid mutator transaction binding the contract method 0x63bd1d4a.
//
// Solidity: function payout() returns()
func (_PuffinClient *PuffinClientSession) Payout() (*types.Transaction, error) {
	return _PuffinClient.Contract.Payout(&_PuffinClient.TransactOpts)
}

// Payout is a paid mutator transaction binding the contract method 0x63bd1d4a.
//
// Solidity: function payout() returns()
func (_PuffinClient *PuffinClientTransactorSession) Payout() (*types.Transaction, error) {
	return _PuffinClient.Contract.Payout(&_PuffinClient.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_PuffinClient *PuffinClientTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _PuffinClient.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_PuffinClient *PuffinClientSession) RenounceOwnership() (*types.Transaction, error) {
	return _PuffinClient.Contract.RenounceOwnership(&_PuffinClient.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_PuffinClient *PuffinClientTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _PuffinClient.Contract.RenounceOwnership(&_PuffinClient.TransactOpts)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_PuffinClient *PuffinClientTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _PuffinClient.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_PuffinClient *PuffinClientSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _PuffinClient.Contract.TransferOwnership(&_PuffinClient.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_PuffinClient *PuffinClientTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _PuffinClient.Contract.TransferOwnership(&_PuffinClient.TransactOpts, newOwner)
}

// PuffinClientNewEpochIterator is returned from FilterNewEpoch and is used to iterate over the raw logs and unpacked data for NewEpoch events raised by the PuffinClient contract.
type PuffinClientNewEpochIterator struct {
	Event *PuffinClientNewEpoch // Event containing the contract specifics and raw log

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
func (it *PuffinClientNewEpochIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PuffinClientNewEpoch)
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
		it.Event = new(PuffinClientNewEpoch)
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
func (it *PuffinClientNewEpochIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PuffinClientNewEpochIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PuffinClientNewEpoch represents a NewEpoch event raised by the PuffinClient contract.
type PuffinClientNewEpoch struct {
	Epoch       *big.Int
	Timestamp   *big.Int
	Users       *big.Int
	UsdcPerUser *big.Int
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterNewEpoch is a free log retrieval operation binding the contract event 0xcc53a96dddb8d3b8710c4932d9e90652a4eaef0bf4c37a13dc4d806e1b54cb4d.
//
// Solidity: event NewEpoch(uint256 epoch, uint256 timestamp, uint256 users, uint256 _usdcPerUser)
func (_PuffinClient *PuffinClientFilterer) FilterNewEpoch(opts *bind.FilterOpts) (*PuffinClientNewEpochIterator, error) {

	logs, sub, err := _PuffinClient.contract.FilterLogs(opts, "NewEpoch")
	if err != nil {
		return nil, err
	}
	return &PuffinClientNewEpochIterator{contract: _PuffinClient.contract, event: "NewEpoch", logs: logs, sub: sub}, nil
}

// WatchNewEpoch is a free log subscription operation binding the contract event 0xcc53a96dddb8d3b8710c4932d9e90652a4eaef0bf4c37a13dc4d806e1b54cb4d.
//
// Solidity: event NewEpoch(uint256 epoch, uint256 timestamp, uint256 users, uint256 _usdcPerUser)
func (_PuffinClient *PuffinClientFilterer) WatchNewEpoch(opts *bind.WatchOpts, sink chan<- *PuffinClientNewEpoch) (event.Subscription, error) {

	logs, sub, err := _PuffinClient.contract.WatchLogs(opts, "NewEpoch")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PuffinClientNewEpoch)
				if err := _PuffinClient.contract.UnpackLog(event, "NewEpoch", log); err != nil {
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

// ParseNewEpoch is a log parse operation binding the contract event 0xcc53a96dddb8d3b8710c4932d9e90652a4eaef0bf4c37a13dc4d806e1b54cb4d.
//
// Solidity: event NewEpoch(uint256 epoch, uint256 timestamp, uint256 users, uint256 _usdcPerUser)
func (_PuffinClient *PuffinClientFilterer) ParseNewEpoch(log types.Log) (*PuffinClientNewEpoch, error) {
	event := new(PuffinClientNewEpoch)
	if err := _PuffinClient.contract.UnpackLog(event, "NewEpoch", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PuffinClientNewPaymentIterator is returned from FilterNewPayment and is used to iterate over the raw logs and unpacked data for NewPayment events raised by the PuffinClient contract.
type PuffinClientNewPaymentIterator struct {
	Event *PuffinClientNewPayment // Event containing the contract specifics and raw log

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
func (it *PuffinClientNewPaymentIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PuffinClientNewPayment)
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
		it.Event = new(PuffinClientNewPayment)
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
func (it *PuffinClientNewPaymentIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PuffinClientNewPaymentIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PuffinClientNewPayment represents a NewPayment event raised by the PuffinClient contract.
type PuffinClientNewPayment struct {
	Epoch            *big.Int
	Timestamp        *big.Int
	Users            *big.Int
	UsdcPerUser      *big.Int
	TotalPaid        *big.Int
	RemainingBalance *big.Int
	Raw              types.Log // Blockchain specific contextual infos
}

// FilterNewPayment is a free log retrieval operation binding the contract event 0xe6445facd3b49f2a8c7b165887e7c442d1a03e0140c91c836d2baded9b9b7f08.
//
// Solidity: event NewPayment(uint256 epoch, uint256 timestamp, uint256 users, uint256 _usdcPerUser, uint256 totalPaid, uint256 remainingBalance)
func (_PuffinClient *PuffinClientFilterer) FilterNewPayment(opts *bind.FilterOpts) (*PuffinClientNewPaymentIterator, error) {

	logs, sub, err := _PuffinClient.contract.FilterLogs(opts, "NewPayment")
	if err != nil {
		return nil, err
	}
	return &PuffinClientNewPaymentIterator{contract: _PuffinClient.contract, event: "NewPayment", logs: logs, sub: sub}, nil
}

// WatchNewPayment is a free log subscription operation binding the contract event 0xe6445facd3b49f2a8c7b165887e7c442d1a03e0140c91c836d2baded9b9b7f08.
//
// Solidity: event NewPayment(uint256 epoch, uint256 timestamp, uint256 users, uint256 _usdcPerUser, uint256 totalPaid, uint256 remainingBalance)
func (_PuffinClient *PuffinClientFilterer) WatchNewPayment(opts *bind.WatchOpts, sink chan<- *PuffinClientNewPayment) (event.Subscription, error) {

	logs, sub, err := _PuffinClient.contract.WatchLogs(opts, "NewPayment")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PuffinClientNewPayment)
				if err := _PuffinClient.contract.UnpackLog(event, "NewPayment", log); err != nil {
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

// ParseNewPayment is a log parse operation binding the contract event 0xe6445facd3b49f2a8c7b165887e7c442d1a03e0140c91c836d2baded9b9b7f08.
//
// Solidity: event NewPayment(uint256 epoch, uint256 timestamp, uint256 users, uint256 _usdcPerUser, uint256 totalPaid, uint256 remainingBalance)
func (_PuffinClient *PuffinClientFilterer) ParseNewPayment(log types.Log) (*PuffinClientNewPayment, error) {
	event := new(PuffinClientNewPayment)
	if err := _PuffinClient.contract.UnpackLog(event, "NewPayment", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PuffinClientOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the PuffinClient contract.
type PuffinClientOwnershipTransferredIterator struct {
	Event *PuffinClientOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *PuffinClientOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PuffinClientOwnershipTransferred)
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
		it.Event = new(PuffinClientOwnershipTransferred)
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
func (it *PuffinClientOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PuffinClientOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PuffinClientOwnershipTransferred represents a OwnershipTransferred event raised by the PuffinClient contract.
type PuffinClientOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_PuffinClient *PuffinClientFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*PuffinClientOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _PuffinClient.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &PuffinClientOwnershipTransferredIterator{contract: _PuffinClient.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_PuffinClient *PuffinClientFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *PuffinClientOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _PuffinClient.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PuffinClientOwnershipTransferred)
				if err := _PuffinClient.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
func (_PuffinClient *PuffinClientFilterer) ParseOwnershipTransferred(log types.Log) (*PuffinClientOwnershipTransferred, error) {
	event := new(PuffinClientOwnershipTransferred)
	if err := _PuffinClient.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
