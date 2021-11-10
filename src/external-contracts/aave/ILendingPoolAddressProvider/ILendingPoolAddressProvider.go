// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package ILendingPoolAddressProvider

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

// AaveMetaData contains all meta data concerning the Aave contract.
var AaveMetaData = &bind.MetaData{
	ABI: "[{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"id\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newAddress\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"hasProxy\",\"type\":\"bool\"}],\"name\":\"AddressSet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newAddress\",\"type\":\"address\"}],\"name\":\"ConfigurationAdminUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newAddress\",\"type\":\"address\"}],\"name\":\"EmergencyAdminUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newAddress\",\"type\":\"address\"}],\"name\":\"LendingPoolCollateralManagerUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newAddress\",\"type\":\"address\"}],\"name\":\"LendingPoolConfiguratorUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newAddress\",\"type\":\"address\"}],\"name\":\"LendingPoolUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newAddress\",\"type\":\"address\"}],\"name\":\"LendingRateOracleUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"string\",\"name\":\"newMarketId\",\"type\":\"string\"}],\"name\":\"MarketIdSet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newAddress\",\"type\":\"address\"}],\"name\":\"PriceOracleUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"id\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newAddress\",\"type\":\"address\"}],\"name\":\"ProxyCreated\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"id\",\"type\":\"bytes32\"}],\"name\":\"getAddress\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getEmergencyAdmin\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getLendingPool\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getLendingPoolCollateralManager\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getLendingPoolConfigurator\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getLendingRateOracle\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getMarketId\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getPoolAdmin\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getPriceOracle\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"id\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"newAddress\",\"type\":\"address\"}],\"name\":\"setAddress\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"id\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"impl\",\"type\":\"address\"}],\"name\":\"setAddressAsProxy\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"admin\",\"type\":\"address\"}],\"name\":\"setEmergencyAdmin\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"manager\",\"type\":\"address\"}],\"name\":\"setLendingPoolCollateralManager\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"configurator\",\"type\":\"address\"}],\"name\":\"setLendingPoolConfiguratorImpl\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"pool\",\"type\":\"address\"}],\"name\":\"setLendingPoolImpl\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"lendingRateOracle\",\"type\":\"address\"}],\"name\":\"setLendingRateOracle\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"marketId\",\"type\":\"string\"}],\"name\":\"setMarketId\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"admin\",\"type\":\"address\"}],\"name\":\"setPoolAdmin\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"priceOracle\",\"type\":\"address\"}],\"name\":\"setPriceOracle\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// AaveABI is the input ABI used to generate the binding from.
// Deprecated: Use AaveMetaData.ABI instead.
var AaveABI = AaveMetaData.ABI

// Aave is an auto generated Go binding around an Ethereum contract.
type Aave struct {
	AaveCaller     // Read-only binding to the contract
	AaveTransactor // Write-only binding to the contract
	AaveFilterer   // Log filterer for contract events
}

// AaveCaller is an auto generated read-only Go binding around an Ethereum contract.
type AaveCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AaveTransactor is an auto generated write-only Go binding around an Ethereum contract.
type AaveTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AaveFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type AaveFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AaveSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type AaveSession struct {
	Contract     *Aave             // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// AaveCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type AaveCallerSession struct {
	Contract *AaveCaller   // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// AaveTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type AaveTransactorSession struct {
	Contract     *AaveTransactor   // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// AaveRaw is an auto generated low-level Go binding around an Ethereum contract.
type AaveRaw struct {
	Contract *Aave // Generic contract binding to access the raw methods on
}

// AaveCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type AaveCallerRaw struct {
	Contract *AaveCaller // Generic read-only contract binding to access the raw methods on
}

// AaveTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type AaveTransactorRaw struct {
	Contract *AaveTransactor // Generic write-only contract binding to access the raw methods on
}

// NewAave creates a new instance of Aave, bound to a specific deployed contract.
func NewAave(address common.Address, backend bind.ContractBackend) (*Aave, error) {
	contract, err := bindAave(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Aave{AaveCaller: AaveCaller{contract: contract}, AaveTransactor: AaveTransactor{contract: contract}, AaveFilterer: AaveFilterer{contract: contract}}, nil
}

// NewAaveCaller creates a new read-only instance of Aave, bound to a specific deployed contract.
func NewAaveCaller(address common.Address, caller bind.ContractCaller) (*AaveCaller, error) {
	contract, err := bindAave(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &AaveCaller{contract: contract}, nil
}

// NewAaveTransactor creates a new write-only instance of Aave, bound to a specific deployed contract.
func NewAaveTransactor(address common.Address, transactor bind.ContractTransactor) (*AaveTransactor, error) {
	contract, err := bindAave(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &AaveTransactor{contract: contract}, nil
}

// NewAaveFilterer creates a new log filterer instance of Aave, bound to a specific deployed contract.
func NewAaveFilterer(address common.Address, filterer bind.ContractFilterer) (*AaveFilterer, error) {
	contract, err := bindAave(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &AaveFilterer{contract: contract}, nil
}

// bindAave binds a generic wrapper to an already deployed contract.
func bindAave(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(AaveABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Aave *AaveRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Aave.Contract.AaveCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Aave *AaveRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Aave.Contract.AaveTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Aave *AaveRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Aave.Contract.AaveTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Aave *AaveCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Aave.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Aave *AaveTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Aave.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Aave *AaveTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Aave.Contract.contract.Transact(opts, method, params...)
}

// GetAddress is a free data retrieval call binding the contract method 0x21f8a721.
//
// Solidity: function getAddress(bytes32 id) view returns(address)
func (_Aave *AaveCaller) GetAddress(opts *bind.CallOpts, id [32]byte) (common.Address, error) {
	var out []interface{}
	err := _Aave.contract.Call(opts, &out, "getAddress", id)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetAddress is a free data retrieval call binding the contract method 0x21f8a721.
//
// Solidity: function getAddress(bytes32 id) view returns(address)
func (_Aave *AaveSession) GetAddress(id [32]byte) (common.Address, error) {
	return _Aave.Contract.GetAddress(&_Aave.CallOpts, id)
}

// GetAddress is a free data retrieval call binding the contract method 0x21f8a721.
//
// Solidity: function getAddress(bytes32 id) view returns(address)
func (_Aave *AaveCallerSession) GetAddress(id [32]byte) (common.Address, error) {
	return _Aave.Contract.GetAddress(&_Aave.CallOpts, id)
}

// GetEmergencyAdmin is a free data retrieval call binding the contract method 0xddcaa9ea.
//
// Solidity: function getEmergencyAdmin() view returns(address)
func (_Aave *AaveCaller) GetEmergencyAdmin(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Aave.contract.Call(opts, &out, "getEmergencyAdmin")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetEmergencyAdmin is a free data retrieval call binding the contract method 0xddcaa9ea.
//
// Solidity: function getEmergencyAdmin() view returns(address)
func (_Aave *AaveSession) GetEmergencyAdmin() (common.Address, error) {
	return _Aave.Contract.GetEmergencyAdmin(&_Aave.CallOpts)
}

// GetEmergencyAdmin is a free data retrieval call binding the contract method 0xddcaa9ea.
//
// Solidity: function getEmergencyAdmin() view returns(address)
func (_Aave *AaveCallerSession) GetEmergencyAdmin() (common.Address, error) {
	return _Aave.Contract.GetEmergencyAdmin(&_Aave.CallOpts)
}

// GetLendingPool is a free data retrieval call binding the contract method 0x0261bf8b.
//
// Solidity: function getLendingPool() view returns(address)
func (_Aave *AaveCaller) GetLendingPool(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Aave.contract.Call(opts, &out, "getLendingPool")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetLendingPool is a free data retrieval call binding the contract method 0x0261bf8b.
//
// Solidity: function getLendingPool() view returns(address)
func (_Aave *AaveSession) GetLendingPool() (common.Address, error) {
	return _Aave.Contract.GetLendingPool(&_Aave.CallOpts)
}

// GetLendingPool is a free data retrieval call binding the contract method 0x0261bf8b.
//
// Solidity: function getLendingPool() view returns(address)
func (_Aave *AaveCallerSession) GetLendingPool() (common.Address, error) {
	return _Aave.Contract.GetLendingPool(&_Aave.CallOpts)
}

// GetLendingPoolCollateralManager is a free data retrieval call binding the contract method 0x712d9171.
//
// Solidity: function getLendingPoolCollateralManager() view returns(address)
func (_Aave *AaveCaller) GetLendingPoolCollateralManager(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Aave.contract.Call(opts, &out, "getLendingPoolCollateralManager")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetLendingPoolCollateralManager is a free data retrieval call binding the contract method 0x712d9171.
//
// Solidity: function getLendingPoolCollateralManager() view returns(address)
func (_Aave *AaveSession) GetLendingPoolCollateralManager() (common.Address, error) {
	return _Aave.Contract.GetLendingPoolCollateralManager(&_Aave.CallOpts)
}

// GetLendingPoolCollateralManager is a free data retrieval call binding the contract method 0x712d9171.
//
// Solidity: function getLendingPoolCollateralManager() view returns(address)
func (_Aave *AaveCallerSession) GetLendingPoolCollateralManager() (common.Address, error) {
	return _Aave.Contract.GetLendingPoolCollateralManager(&_Aave.CallOpts)
}

// GetLendingPoolConfigurator is a free data retrieval call binding the contract method 0x85c858b1.
//
// Solidity: function getLendingPoolConfigurator() view returns(address)
func (_Aave *AaveCaller) GetLendingPoolConfigurator(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Aave.contract.Call(opts, &out, "getLendingPoolConfigurator")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetLendingPoolConfigurator is a free data retrieval call binding the contract method 0x85c858b1.
//
// Solidity: function getLendingPoolConfigurator() view returns(address)
func (_Aave *AaveSession) GetLendingPoolConfigurator() (common.Address, error) {
	return _Aave.Contract.GetLendingPoolConfigurator(&_Aave.CallOpts)
}

// GetLendingPoolConfigurator is a free data retrieval call binding the contract method 0x85c858b1.
//
// Solidity: function getLendingPoolConfigurator() view returns(address)
func (_Aave *AaveCallerSession) GetLendingPoolConfigurator() (common.Address, error) {
	return _Aave.Contract.GetLendingPoolConfigurator(&_Aave.CallOpts)
}

// GetLendingRateOracle is a free data retrieval call binding the contract method 0x3618abba.
//
// Solidity: function getLendingRateOracle() view returns(address)
func (_Aave *AaveCaller) GetLendingRateOracle(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Aave.contract.Call(opts, &out, "getLendingRateOracle")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetLendingRateOracle is a free data retrieval call binding the contract method 0x3618abba.
//
// Solidity: function getLendingRateOracle() view returns(address)
func (_Aave *AaveSession) GetLendingRateOracle() (common.Address, error) {
	return _Aave.Contract.GetLendingRateOracle(&_Aave.CallOpts)
}

// GetLendingRateOracle is a free data retrieval call binding the contract method 0x3618abba.
//
// Solidity: function getLendingRateOracle() view returns(address)
func (_Aave *AaveCallerSession) GetLendingRateOracle() (common.Address, error) {
	return _Aave.Contract.GetLendingRateOracle(&_Aave.CallOpts)
}

// GetMarketId is a free data retrieval call binding the contract method 0x568ef470.
//
// Solidity: function getMarketId() view returns(string)
func (_Aave *AaveCaller) GetMarketId(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _Aave.contract.Call(opts, &out, "getMarketId")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// GetMarketId is a free data retrieval call binding the contract method 0x568ef470.
//
// Solidity: function getMarketId() view returns(string)
func (_Aave *AaveSession) GetMarketId() (string, error) {
	return _Aave.Contract.GetMarketId(&_Aave.CallOpts)
}

// GetMarketId is a free data retrieval call binding the contract method 0x568ef470.
//
// Solidity: function getMarketId() view returns(string)
func (_Aave *AaveCallerSession) GetMarketId() (string, error) {
	return _Aave.Contract.GetMarketId(&_Aave.CallOpts)
}

// GetPoolAdmin is a free data retrieval call binding the contract method 0xaecda378.
//
// Solidity: function getPoolAdmin() view returns(address)
func (_Aave *AaveCaller) GetPoolAdmin(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Aave.contract.Call(opts, &out, "getPoolAdmin")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetPoolAdmin is a free data retrieval call binding the contract method 0xaecda378.
//
// Solidity: function getPoolAdmin() view returns(address)
func (_Aave *AaveSession) GetPoolAdmin() (common.Address, error) {
	return _Aave.Contract.GetPoolAdmin(&_Aave.CallOpts)
}

// GetPoolAdmin is a free data retrieval call binding the contract method 0xaecda378.
//
// Solidity: function getPoolAdmin() view returns(address)
func (_Aave *AaveCallerSession) GetPoolAdmin() (common.Address, error) {
	return _Aave.Contract.GetPoolAdmin(&_Aave.CallOpts)
}

// GetPriceOracle is a free data retrieval call binding the contract method 0xfca513a8.
//
// Solidity: function getPriceOracle() view returns(address)
func (_Aave *AaveCaller) GetPriceOracle(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Aave.contract.Call(opts, &out, "getPriceOracle")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetPriceOracle is a free data retrieval call binding the contract method 0xfca513a8.
//
// Solidity: function getPriceOracle() view returns(address)
func (_Aave *AaveSession) GetPriceOracle() (common.Address, error) {
	return _Aave.Contract.GetPriceOracle(&_Aave.CallOpts)
}

// GetPriceOracle is a free data retrieval call binding the contract method 0xfca513a8.
//
// Solidity: function getPriceOracle() view returns(address)
func (_Aave *AaveCallerSession) GetPriceOracle() (common.Address, error) {
	return _Aave.Contract.GetPriceOracle(&_Aave.CallOpts)
}

// SetAddress is a paid mutator transaction binding the contract method 0xca446dd9.
//
// Solidity: function setAddress(bytes32 id, address newAddress) returns()
func (_Aave *AaveTransactor) SetAddress(opts *bind.TransactOpts, id [32]byte, newAddress common.Address) (*types.Transaction, error) {
	return _Aave.contract.Transact(opts, "setAddress", id, newAddress)
}

// SetAddress is a paid mutator transaction binding the contract method 0xca446dd9.
//
// Solidity: function setAddress(bytes32 id, address newAddress) returns()
func (_Aave *AaveSession) SetAddress(id [32]byte, newAddress common.Address) (*types.Transaction, error) {
	return _Aave.Contract.SetAddress(&_Aave.TransactOpts, id, newAddress)
}

// SetAddress is a paid mutator transaction binding the contract method 0xca446dd9.
//
// Solidity: function setAddress(bytes32 id, address newAddress) returns()
func (_Aave *AaveTransactorSession) SetAddress(id [32]byte, newAddress common.Address) (*types.Transaction, error) {
	return _Aave.Contract.SetAddress(&_Aave.TransactOpts, id, newAddress)
}

// SetAddressAsProxy is a paid mutator transaction binding the contract method 0x5dcc528c.
//
// Solidity: function setAddressAsProxy(bytes32 id, address impl) returns()
func (_Aave *AaveTransactor) SetAddressAsProxy(opts *bind.TransactOpts, id [32]byte, impl common.Address) (*types.Transaction, error) {
	return _Aave.contract.Transact(opts, "setAddressAsProxy", id, impl)
}

// SetAddressAsProxy is a paid mutator transaction binding the contract method 0x5dcc528c.
//
// Solidity: function setAddressAsProxy(bytes32 id, address impl) returns()
func (_Aave *AaveSession) SetAddressAsProxy(id [32]byte, impl common.Address) (*types.Transaction, error) {
	return _Aave.Contract.SetAddressAsProxy(&_Aave.TransactOpts, id, impl)
}

// SetAddressAsProxy is a paid mutator transaction binding the contract method 0x5dcc528c.
//
// Solidity: function setAddressAsProxy(bytes32 id, address impl) returns()
func (_Aave *AaveTransactorSession) SetAddressAsProxy(id [32]byte, impl common.Address) (*types.Transaction, error) {
	return _Aave.Contract.SetAddressAsProxy(&_Aave.TransactOpts, id, impl)
}

// SetEmergencyAdmin is a paid mutator transaction binding the contract method 0x35da3394.
//
// Solidity: function setEmergencyAdmin(address admin) returns()
func (_Aave *AaveTransactor) SetEmergencyAdmin(opts *bind.TransactOpts, admin common.Address) (*types.Transaction, error) {
	return _Aave.contract.Transact(opts, "setEmergencyAdmin", admin)
}

// SetEmergencyAdmin is a paid mutator transaction binding the contract method 0x35da3394.
//
// Solidity: function setEmergencyAdmin(address admin) returns()
func (_Aave *AaveSession) SetEmergencyAdmin(admin common.Address) (*types.Transaction, error) {
	return _Aave.Contract.SetEmergencyAdmin(&_Aave.TransactOpts, admin)
}

// SetEmergencyAdmin is a paid mutator transaction binding the contract method 0x35da3394.
//
// Solidity: function setEmergencyAdmin(address admin) returns()
func (_Aave *AaveTransactorSession) SetEmergencyAdmin(admin common.Address) (*types.Transaction, error) {
	return _Aave.Contract.SetEmergencyAdmin(&_Aave.TransactOpts, admin)
}

// SetLendingPoolCollateralManager is a paid mutator transaction binding the contract method 0x398e5553.
//
// Solidity: function setLendingPoolCollateralManager(address manager) returns()
func (_Aave *AaveTransactor) SetLendingPoolCollateralManager(opts *bind.TransactOpts, manager common.Address) (*types.Transaction, error) {
	return _Aave.contract.Transact(opts, "setLendingPoolCollateralManager", manager)
}

// SetLendingPoolCollateralManager is a paid mutator transaction binding the contract method 0x398e5553.
//
// Solidity: function setLendingPoolCollateralManager(address manager) returns()
func (_Aave *AaveSession) SetLendingPoolCollateralManager(manager common.Address) (*types.Transaction, error) {
	return _Aave.Contract.SetLendingPoolCollateralManager(&_Aave.TransactOpts, manager)
}

// SetLendingPoolCollateralManager is a paid mutator transaction binding the contract method 0x398e5553.
//
// Solidity: function setLendingPoolCollateralManager(address manager) returns()
func (_Aave *AaveTransactorSession) SetLendingPoolCollateralManager(manager common.Address) (*types.Transaction, error) {
	return _Aave.Contract.SetLendingPoolCollateralManager(&_Aave.TransactOpts, manager)
}

// SetLendingPoolConfiguratorImpl is a paid mutator transaction binding the contract method 0xc12542df.
//
// Solidity: function setLendingPoolConfiguratorImpl(address configurator) returns()
func (_Aave *AaveTransactor) SetLendingPoolConfiguratorImpl(opts *bind.TransactOpts, configurator common.Address) (*types.Transaction, error) {
	return _Aave.contract.Transact(opts, "setLendingPoolConfiguratorImpl", configurator)
}

// SetLendingPoolConfiguratorImpl is a paid mutator transaction binding the contract method 0xc12542df.
//
// Solidity: function setLendingPoolConfiguratorImpl(address configurator) returns()
func (_Aave *AaveSession) SetLendingPoolConfiguratorImpl(configurator common.Address) (*types.Transaction, error) {
	return _Aave.Contract.SetLendingPoolConfiguratorImpl(&_Aave.TransactOpts, configurator)
}

// SetLendingPoolConfiguratorImpl is a paid mutator transaction binding the contract method 0xc12542df.
//
// Solidity: function setLendingPoolConfiguratorImpl(address configurator) returns()
func (_Aave *AaveTransactorSession) SetLendingPoolConfiguratorImpl(configurator common.Address) (*types.Transaction, error) {
	return _Aave.Contract.SetLendingPoolConfiguratorImpl(&_Aave.TransactOpts, configurator)
}

// SetLendingPoolImpl is a paid mutator transaction binding the contract method 0x5aef021f.
//
// Solidity: function setLendingPoolImpl(address pool) returns()
func (_Aave *AaveTransactor) SetLendingPoolImpl(opts *bind.TransactOpts, pool common.Address) (*types.Transaction, error) {
	return _Aave.contract.Transact(opts, "setLendingPoolImpl", pool)
}

// SetLendingPoolImpl is a paid mutator transaction binding the contract method 0x5aef021f.
//
// Solidity: function setLendingPoolImpl(address pool) returns()
func (_Aave *AaveSession) SetLendingPoolImpl(pool common.Address) (*types.Transaction, error) {
	return _Aave.Contract.SetLendingPoolImpl(&_Aave.TransactOpts, pool)
}

// SetLendingPoolImpl is a paid mutator transaction binding the contract method 0x5aef021f.
//
// Solidity: function setLendingPoolImpl(address pool) returns()
func (_Aave *AaveTransactorSession) SetLendingPoolImpl(pool common.Address) (*types.Transaction, error) {
	return _Aave.Contract.SetLendingPoolImpl(&_Aave.TransactOpts, pool)
}

// SetLendingRateOracle is a paid mutator transaction binding the contract method 0x820d1274.
//
// Solidity: function setLendingRateOracle(address lendingRateOracle) returns()
func (_Aave *AaveTransactor) SetLendingRateOracle(opts *bind.TransactOpts, lendingRateOracle common.Address) (*types.Transaction, error) {
	return _Aave.contract.Transact(opts, "setLendingRateOracle", lendingRateOracle)
}

// SetLendingRateOracle is a paid mutator transaction binding the contract method 0x820d1274.
//
// Solidity: function setLendingRateOracle(address lendingRateOracle) returns()
func (_Aave *AaveSession) SetLendingRateOracle(lendingRateOracle common.Address) (*types.Transaction, error) {
	return _Aave.Contract.SetLendingRateOracle(&_Aave.TransactOpts, lendingRateOracle)
}

// SetLendingRateOracle is a paid mutator transaction binding the contract method 0x820d1274.
//
// Solidity: function setLendingRateOracle(address lendingRateOracle) returns()
func (_Aave *AaveTransactorSession) SetLendingRateOracle(lendingRateOracle common.Address) (*types.Transaction, error) {
	return _Aave.Contract.SetLendingRateOracle(&_Aave.TransactOpts, lendingRateOracle)
}

// SetMarketId is a paid mutator transaction binding the contract method 0xf67b1847.
//
// Solidity: function setMarketId(string marketId) returns()
func (_Aave *AaveTransactor) SetMarketId(opts *bind.TransactOpts, marketId string) (*types.Transaction, error) {
	return _Aave.contract.Transact(opts, "setMarketId", marketId)
}

// SetMarketId is a paid mutator transaction binding the contract method 0xf67b1847.
//
// Solidity: function setMarketId(string marketId) returns()
func (_Aave *AaveSession) SetMarketId(marketId string) (*types.Transaction, error) {
	return _Aave.Contract.SetMarketId(&_Aave.TransactOpts, marketId)
}

// SetMarketId is a paid mutator transaction binding the contract method 0xf67b1847.
//
// Solidity: function setMarketId(string marketId) returns()
func (_Aave *AaveTransactorSession) SetMarketId(marketId string) (*types.Transaction, error) {
	return _Aave.Contract.SetMarketId(&_Aave.TransactOpts, marketId)
}

// SetPoolAdmin is a paid mutator transaction binding the contract method 0x283d62ad.
//
// Solidity: function setPoolAdmin(address admin) returns()
func (_Aave *AaveTransactor) SetPoolAdmin(opts *bind.TransactOpts, admin common.Address) (*types.Transaction, error) {
	return _Aave.contract.Transact(opts, "setPoolAdmin", admin)
}

// SetPoolAdmin is a paid mutator transaction binding the contract method 0x283d62ad.
//
// Solidity: function setPoolAdmin(address admin) returns()
func (_Aave *AaveSession) SetPoolAdmin(admin common.Address) (*types.Transaction, error) {
	return _Aave.Contract.SetPoolAdmin(&_Aave.TransactOpts, admin)
}

// SetPoolAdmin is a paid mutator transaction binding the contract method 0x283d62ad.
//
// Solidity: function setPoolAdmin(address admin) returns()
func (_Aave *AaveTransactorSession) SetPoolAdmin(admin common.Address) (*types.Transaction, error) {
	return _Aave.Contract.SetPoolAdmin(&_Aave.TransactOpts, admin)
}

// SetPriceOracle is a paid mutator transaction binding the contract method 0x530e784f.
//
// Solidity: function setPriceOracle(address priceOracle) returns()
func (_Aave *AaveTransactor) SetPriceOracle(opts *bind.TransactOpts, priceOracle common.Address) (*types.Transaction, error) {
	return _Aave.contract.Transact(opts, "setPriceOracle", priceOracle)
}

// SetPriceOracle is a paid mutator transaction binding the contract method 0x530e784f.
//
// Solidity: function setPriceOracle(address priceOracle) returns()
func (_Aave *AaveSession) SetPriceOracle(priceOracle common.Address) (*types.Transaction, error) {
	return _Aave.Contract.SetPriceOracle(&_Aave.TransactOpts, priceOracle)
}

// SetPriceOracle is a paid mutator transaction binding the contract method 0x530e784f.
//
// Solidity: function setPriceOracle(address priceOracle) returns()
func (_Aave *AaveTransactorSession) SetPriceOracle(priceOracle common.Address) (*types.Transaction, error) {
	return _Aave.Contract.SetPriceOracle(&_Aave.TransactOpts, priceOracle)
}

// AaveAddressSetIterator is returned from FilterAddressSet and is used to iterate over the raw logs and unpacked data for AddressSet events raised by the Aave contract.
type AaveAddressSetIterator struct {
	Event *AaveAddressSet // Event containing the contract specifics and raw log

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
func (it *AaveAddressSetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AaveAddressSet)
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
		it.Event = new(AaveAddressSet)
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
func (it *AaveAddressSetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AaveAddressSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AaveAddressSet represents a AddressSet event raised by the Aave contract.
type AaveAddressSet struct {
	Id         [32]byte
	NewAddress common.Address
	HasProxy   bool
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterAddressSet is a free log retrieval operation binding the contract event 0xf2689d5d5cd0c639e137642cae5d40afced201a1a0327e7ac9358461dc9fff31.
//
// Solidity: event AddressSet(bytes32 id, address indexed newAddress, bool hasProxy)
func (_Aave *AaveFilterer) FilterAddressSet(opts *bind.FilterOpts, newAddress []common.Address) (*AaveAddressSetIterator, error) {

	var newAddressRule []interface{}
	for _, newAddressItem := range newAddress {
		newAddressRule = append(newAddressRule, newAddressItem)
	}

	logs, sub, err := _Aave.contract.FilterLogs(opts, "AddressSet", newAddressRule)
	if err != nil {
		return nil, err
	}
	return &AaveAddressSetIterator{contract: _Aave.contract, event: "AddressSet", logs: logs, sub: sub}, nil
}

// WatchAddressSet is a free log subscription operation binding the contract event 0xf2689d5d5cd0c639e137642cae5d40afced201a1a0327e7ac9358461dc9fff31.
//
// Solidity: event AddressSet(bytes32 id, address indexed newAddress, bool hasProxy)
func (_Aave *AaveFilterer) WatchAddressSet(opts *bind.WatchOpts, sink chan<- *AaveAddressSet, newAddress []common.Address) (event.Subscription, error) {

	var newAddressRule []interface{}
	for _, newAddressItem := range newAddress {
		newAddressRule = append(newAddressRule, newAddressItem)
	}

	logs, sub, err := _Aave.contract.WatchLogs(opts, "AddressSet", newAddressRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AaveAddressSet)
				if err := _Aave.contract.UnpackLog(event, "AddressSet", log); err != nil {
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

// ParseAddressSet is a log parse operation binding the contract event 0xf2689d5d5cd0c639e137642cae5d40afced201a1a0327e7ac9358461dc9fff31.
//
// Solidity: event AddressSet(bytes32 id, address indexed newAddress, bool hasProxy)
func (_Aave *AaveFilterer) ParseAddressSet(log types.Log) (*AaveAddressSet, error) {
	event := new(AaveAddressSet)
	if err := _Aave.contract.UnpackLog(event, "AddressSet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AaveConfigurationAdminUpdatedIterator is returned from FilterConfigurationAdminUpdated and is used to iterate over the raw logs and unpacked data for ConfigurationAdminUpdated events raised by the Aave contract.
type AaveConfigurationAdminUpdatedIterator struct {
	Event *AaveConfigurationAdminUpdated // Event containing the contract specifics and raw log

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
func (it *AaveConfigurationAdminUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AaveConfigurationAdminUpdated)
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
		it.Event = new(AaveConfigurationAdminUpdated)
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
func (it *AaveConfigurationAdminUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AaveConfigurationAdminUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AaveConfigurationAdminUpdated represents a ConfigurationAdminUpdated event raised by the Aave contract.
type AaveConfigurationAdminUpdated struct {
	NewAddress common.Address
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterConfigurationAdminUpdated is a free log retrieval operation binding the contract event 0xc20a317155a9e7d84e06b716b4b355d47742ab9f8c5d630e7f556553f582430d.
//
// Solidity: event ConfigurationAdminUpdated(address indexed newAddress)
func (_Aave *AaveFilterer) FilterConfigurationAdminUpdated(opts *bind.FilterOpts, newAddress []common.Address) (*AaveConfigurationAdminUpdatedIterator, error) {

	var newAddressRule []interface{}
	for _, newAddressItem := range newAddress {
		newAddressRule = append(newAddressRule, newAddressItem)
	}

	logs, sub, err := _Aave.contract.FilterLogs(opts, "ConfigurationAdminUpdated", newAddressRule)
	if err != nil {
		return nil, err
	}
	return &AaveConfigurationAdminUpdatedIterator{contract: _Aave.contract, event: "ConfigurationAdminUpdated", logs: logs, sub: sub}, nil
}

// WatchConfigurationAdminUpdated is a free log subscription operation binding the contract event 0xc20a317155a9e7d84e06b716b4b355d47742ab9f8c5d630e7f556553f582430d.
//
// Solidity: event ConfigurationAdminUpdated(address indexed newAddress)
func (_Aave *AaveFilterer) WatchConfigurationAdminUpdated(opts *bind.WatchOpts, sink chan<- *AaveConfigurationAdminUpdated, newAddress []common.Address) (event.Subscription, error) {

	var newAddressRule []interface{}
	for _, newAddressItem := range newAddress {
		newAddressRule = append(newAddressRule, newAddressItem)
	}

	logs, sub, err := _Aave.contract.WatchLogs(opts, "ConfigurationAdminUpdated", newAddressRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AaveConfigurationAdminUpdated)
				if err := _Aave.contract.UnpackLog(event, "ConfigurationAdminUpdated", log); err != nil {
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

// ParseConfigurationAdminUpdated is a log parse operation binding the contract event 0xc20a317155a9e7d84e06b716b4b355d47742ab9f8c5d630e7f556553f582430d.
//
// Solidity: event ConfigurationAdminUpdated(address indexed newAddress)
func (_Aave *AaveFilterer) ParseConfigurationAdminUpdated(log types.Log) (*AaveConfigurationAdminUpdated, error) {
	event := new(AaveConfigurationAdminUpdated)
	if err := _Aave.contract.UnpackLog(event, "ConfigurationAdminUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AaveEmergencyAdminUpdatedIterator is returned from FilterEmergencyAdminUpdated and is used to iterate over the raw logs and unpacked data for EmergencyAdminUpdated events raised by the Aave contract.
type AaveEmergencyAdminUpdatedIterator struct {
	Event *AaveEmergencyAdminUpdated // Event containing the contract specifics and raw log

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
func (it *AaveEmergencyAdminUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AaveEmergencyAdminUpdated)
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
		it.Event = new(AaveEmergencyAdminUpdated)
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
func (it *AaveEmergencyAdminUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AaveEmergencyAdminUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AaveEmergencyAdminUpdated represents a EmergencyAdminUpdated event raised by the Aave contract.
type AaveEmergencyAdminUpdated struct {
	NewAddress common.Address
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterEmergencyAdminUpdated is a free log retrieval operation binding the contract event 0xe19673fc861bfeb894cf2d6b7662505497ef31c0f489b742db24ee3310826916.
//
// Solidity: event EmergencyAdminUpdated(address indexed newAddress)
func (_Aave *AaveFilterer) FilterEmergencyAdminUpdated(opts *bind.FilterOpts, newAddress []common.Address) (*AaveEmergencyAdminUpdatedIterator, error) {

	var newAddressRule []interface{}
	for _, newAddressItem := range newAddress {
		newAddressRule = append(newAddressRule, newAddressItem)
	}

	logs, sub, err := _Aave.contract.FilterLogs(opts, "EmergencyAdminUpdated", newAddressRule)
	if err != nil {
		return nil, err
	}
	return &AaveEmergencyAdminUpdatedIterator{contract: _Aave.contract, event: "EmergencyAdminUpdated", logs: logs, sub: sub}, nil
}

// WatchEmergencyAdminUpdated is a free log subscription operation binding the contract event 0xe19673fc861bfeb894cf2d6b7662505497ef31c0f489b742db24ee3310826916.
//
// Solidity: event EmergencyAdminUpdated(address indexed newAddress)
func (_Aave *AaveFilterer) WatchEmergencyAdminUpdated(opts *bind.WatchOpts, sink chan<- *AaveEmergencyAdminUpdated, newAddress []common.Address) (event.Subscription, error) {

	var newAddressRule []interface{}
	for _, newAddressItem := range newAddress {
		newAddressRule = append(newAddressRule, newAddressItem)
	}

	logs, sub, err := _Aave.contract.WatchLogs(opts, "EmergencyAdminUpdated", newAddressRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AaveEmergencyAdminUpdated)
				if err := _Aave.contract.UnpackLog(event, "EmergencyAdminUpdated", log); err != nil {
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

// ParseEmergencyAdminUpdated is a log parse operation binding the contract event 0xe19673fc861bfeb894cf2d6b7662505497ef31c0f489b742db24ee3310826916.
//
// Solidity: event EmergencyAdminUpdated(address indexed newAddress)
func (_Aave *AaveFilterer) ParseEmergencyAdminUpdated(log types.Log) (*AaveEmergencyAdminUpdated, error) {
	event := new(AaveEmergencyAdminUpdated)
	if err := _Aave.contract.UnpackLog(event, "EmergencyAdminUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AaveLendingPoolCollateralManagerUpdatedIterator is returned from FilterLendingPoolCollateralManagerUpdated and is used to iterate over the raw logs and unpacked data for LendingPoolCollateralManagerUpdated events raised by the Aave contract.
type AaveLendingPoolCollateralManagerUpdatedIterator struct {
	Event *AaveLendingPoolCollateralManagerUpdated // Event containing the contract specifics and raw log

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
func (it *AaveLendingPoolCollateralManagerUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AaveLendingPoolCollateralManagerUpdated)
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
		it.Event = new(AaveLendingPoolCollateralManagerUpdated)
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
func (it *AaveLendingPoolCollateralManagerUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AaveLendingPoolCollateralManagerUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AaveLendingPoolCollateralManagerUpdated represents a LendingPoolCollateralManagerUpdated event raised by the Aave contract.
type AaveLendingPoolCollateralManagerUpdated struct {
	NewAddress common.Address
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterLendingPoolCollateralManagerUpdated is a free log retrieval operation binding the contract event 0x991888326f0eab3df6084aadb82bee6781b5c9aa75379e8bc50ae86934541638.
//
// Solidity: event LendingPoolCollateralManagerUpdated(address indexed newAddress)
func (_Aave *AaveFilterer) FilterLendingPoolCollateralManagerUpdated(opts *bind.FilterOpts, newAddress []common.Address) (*AaveLendingPoolCollateralManagerUpdatedIterator, error) {

	var newAddressRule []interface{}
	for _, newAddressItem := range newAddress {
		newAddressRule = append(newAddressRule, newAddressItem)
	}

	logs, sub, err := _Aave.contract.FilterLogs(opts, "LendingPoolCollateralManagerUpdated", newAddressRule)
	if err != nil {
		return nil, err
	}
	return &AaveLendingPoolCollateralManagerUpdatedIterator{contract: _Aave.contract, event: "LendingPoolCollateralManagerUpdated", logs: logs, sub: sub}, nil
}

// WatchLendingPoolCollateralManagerUpdated is a free log subscription operation binding the contract event 0x991888326f0eab3df6084aadb82bee6781b5c9aa75379e8bc50ae86934541638.
//
// Solidity: event LendingPoolCollateralManagerUpdated(address indexed newAddress)
func (_Aave *AaveFilterer) WatchLendingPoolCollateralManagerUpdated(opts *bind.WatchOpts, sink chan<- *AaveLendingPoolCollateralManagerUpdated, newAddress []common.Address) (event.Subscription, error) {

	var newAddressRule []interface{}
	for _, newAddressItem := range newAddress {
		newAddressRule = append(newAddressRule, newAddressItem)
	}

	logs, sub, err := _Aave.contract.WatchLogs(opts, "LendingPoolCollateralManagerUpdated", newAddressRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AaveLendingPoolCollateralManagerUpdated)
				if err := _Aave.contract.UnpackLog(event, "LendingPoolCollateralManagerUpdated", log); err != nil {
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

// ParseLendingPoolCollateralManagerUpdated is a log parse operation binding the contract event 0x991888326f0eab3df6084aadb82bee6781b5c9aa75379e8bc50ae86934541638.
//
// Solidity: event LendingPoolCollateralManagerUpdated(address indexed newAddress)
func (_Aave *AaveFilterer) ParseLendingPoolCollateralManagerUpdated(log types.Log) (*AaveLendingPoolCollateralManagerUpdated, error) {
	event := new(AaveLendingPoolCollateralManagerUpdated)
	if err := _Aave.contract.UnpackLog(event, "LendingPoolCollateralManagerUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AaveLendingPoolConfiguratorUpdatedIterator is returned from FilterLendingPoolConfiguratorUpdated and is used to iterate over the raw logs and unpacked data for LendingPoolConfiguratorUpdated events raised by the Aave contract.
type AaveLendingPoolConfiguratorUpdatedIterator struct {
	Event *AaveLendingPoolConfiguratorUpdated // Event containing the contract specifics and raw log

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
func (it *AaveLendingPoolConfiguratorUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AaveLendingPoolConfiguratorUpdated)
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
		it.Event = new(AaveLendingPoolConfiguratorUpdated)
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
func (it *AaveLendingPoolConfiguratorUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AaveLendingPoolConfiguratorUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AaveLendingPoolConfiguratorUpdated represents a LendingPoolConfiguratorUpdated event raised by the Aave contract.
type AaveLendingPoolConfiguratorUpdated struct {
	NewAddress common.Address
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterLendingPoolConfiguratorUpdated is a free log retrieval operation binding the contract event 0xdfabe479bad36782fb1e77fbfddd4e382671713527e4786cfc93a022ae763729.
//
// Solidity: event LendingPoolConfiguratorUpdated(address indexed newAddress)
func (_Aave *AaveFilterer) FilterLendingPoolConfiguratorUpdated(opts *bind.FilterOpts, newAddress []common.Address) (*AaveLendingPoolConfiguratorUpdatedIterator, error) {

	var newAddressRule []interface{}
	for _, newAddressItem := range newAddress {
		newAddressRule = append(newAddressRule, newAddressItem)
	}

	logs, sub, err := _Aave.contract.FilterLogs(opts, "LendingPoolConfiguratorUpdated", newAddressRule)
	if err != nil {
		return nil, err
	}
	return &AaveLendingPoolConfiguratorUpdatedIterator{contract: _Aave.contract, event: "LendingPoolConfiguratorUpdated", logs: logs, sub: sub}, nil
}

// WatchLendingPoolConfiguratorUpdated is a free log subscription operation binding the contract event 0xdfabe479bad36782fb1e77fbfddd4e382671713527e4786cfc93a022ae763729.
//
// Solidity: event LendingPoolConfiguratorUpdated(address indexed newAddress)
func (_Aave *AaveFilterer) WatchLendingPoolConfiguratorUpdated(opts *bind.WatchOpts, sink chan<- *AaveLendingPoolConfiguratorUpdated, newAddress []common.Address) (event.Subscription, error) {

	var newAddressRule []interface{}
	for _, newAddressItem := range newAddress {
		newAddressRule = append(newAddressRule, newAddressItem)
	}

	logs, sub, err := _Aave.contract.WatchLogs(opts, "LendingPoolConfiguratorUpdated", newAddressRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AaveLendingPoolConfiguratorUpdated)
				if err := _Aave.contract.UnpackLog(event, "LendingPoolConfiguratorUpdated", log); err != nil {
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

// ParseLendingPoolConfiguratorUpdated is a log parse operation binding the contract event 0xdfabe479bad36782fb1e77fbfddd4e382671713527e4786cfc93a022ae763729.
//
// Solidity: event LendingPoolConfiguratorUpdated(address indexed newAddress)
func (_Aave *AaveFilterer) ParseLendingPoolConfiguratorUpdated(log types.Log) (*AaveLendingPoolConfiguratorUpdated, error) {
	event := new(AaveLendingPoolConfiguratorUpdated)
	if err := _Aave.contract.UnpackLog(event, "LendingPoolConfiguratorUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AaveLendingPoolUpdatedIterator is returned from FilterLendingPoolUpdated and is used to iterate over the raw logs and unpacked data for LendingPoolUpdated events raised by the Aave contract.
type AaveLendingPoolUpdatedIterator struct {
	Event *AaveLendingPoolUpdated // Event containing the contract specifics and raw log

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
func (it *AaveLendingPoolUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AaveLendingPoolUpdated)
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
		it.Event = new(AaveLendingPoolUpdated)
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
func (it *AaveLendingPoolUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AaveLendingPoolUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AaveLendingPoolUpdated represents a LendingPoolUpdated event raised by the Aave contract.
type AaveLendingPoolUpdated struct {
	NewAddress common.Address
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterLendingPoolUpdated is a free log retrieval operation binding the contract event 0xc4e6c6cdf28d0edbd8bcf071d724d33cc2e7a30be7d06443925656e9cb492aa4.
//
// Solidity: event LendingPoolUpdated(address indexed newAddress)
func (_Aave *AaveFilterer) FilterLendingPoolUpdated(opts *bind.FilterOpts, newAddress []common.Address) (*AaveLendingPoolUpdatedIterator, error) {

	var newAddressRule []interface{}
	for _, newAddressItem := range newAddress {
		newAddressRule = append(newAddressRule, newAddressItem)
	}

	logs, sub, err := _Aave.contract.FilterLogs(opts, "LendingPoolUpdated", newAddressRule)
	if err != nil {
		return nil, err
	}
	return &AaveLendingPoolUpdatedIterator{contract: _Aave.contract, event: "LendingPoolUpdated", logs: logs, sub: sub}, nil
}

// WatchLendingPoolUpdated is a free log subscription operation binding the contract event 0xc4e6c6cdf28d0edbd8bcf071d724d33cc2e7a30be7d06443925656e9cb492aa4.
//
// Solidity: event LendingPoolUpdated(address indexed newAddress)
func (_Aave *AaveFilterer) WatchLendingPoolUpdated(opts *bind.WatchOpts, sink chan<- *AaveLendingPoolUpdated, newAddress []common.Address) (event.Subscription, error) {

	var newAddressRule []interface{}
	for _, newAddressItem := range newAddress {
		newAddressRule = append(newAddressRule, newAddressItem)
	}

	logs, sub, err := _Aave.contract.WatchLogs(opts, "LendingPoolUpdated", newAddressRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AaveLendingPoolUpdated)
				if err := _Aave.contract.UnpackLog(event, "LendingPoolUpdated", log); err != nil {
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

// ParseLendingPoolUpdated is a log parse operation binding the contract event 0xc4e6c6cdf28d0edbd8bcf071d724d33cc2e7a30be7d06443925656e9cb492aa4.
//
// Solidity: event LendingPoolUpdated(address indexed newAddress)
func (_Aave *AaveFilterer) ParseLendingPoolUpdated(log types.Log) (*AaveLendingPoolUpdated, error) {
	event := new(AaveLendingPoolUpdated)
	if err := _Aave.contract.UnpackLog(event, "LendingPoolUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AaveLendingRateOracleUpdatedIterator is returned from FilterLendingRateOracleUpdated and is used to iterate over the raw logs and unpacked data for LendingRateOracleUpdated events raised by the Aave contract.
type AaveLendingRateOracleUpdatedIterator struct {
	Event *AaveLendingRateOracleUpdated // Event containing the contract specifics and raw log

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
func (it *AaveLendingRateOracleUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AaveLendingRateOracleUpdated)
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
		it.Event = new(AaveLendingRateOracleUpdated)
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
func (it *AaveLendingRateOracleUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AaveLendingRateOracleUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AaveLendingRateOracleUpdated represents a LendingRateOracleUpdated event raised by the Aave contract.
type AaveLendingRateOracleUpdated struct {
	NewAddress common.Address
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterLendingRateOracleUpdated is a free log retrieval operation binding the contract event 0x5c29179aba6942020a8a2d38f65de02fb6b7f784e7f049ed3a3cab97621859b5.
//
// Solidity: event LendingRateOracleUpdated(address indexed newAddress)
func (_Aave *AaveFilterer) FilterLendingRateOracleUpdated(opts *bind.FilterOpts, newAddress []common.Address) (*AaveLendingRateOracleUpdatedIterator, error) {

	var newAddressRule []interface{}
	for _, newAddressItem := range newAddress {
		newAddressRule = append(newAddressRule, newAddressItem)
	}

	logs, sub, err := _Aave.contract.FilterLogs(opts, "LendingRateOracleUpdated", newAddressRule)
	if err != nil {
		return nil, err
	}
	return &AaveLendingRateOracleUpdatedIterator{contract: _Aave.contract, event: "LendingRateOracleUpdated", logs: logs, sub: sub}, nil
}

// WatchLendingRateOracleUpdated is a free log subscription operation binding the contract event 0x5c29179aba6942020a8a2d38f65de02fb6b7f784e7f049ed3a3cab97621859b5.
//
// Solidity: event LendingRateOracleUpdated(address indexed newAddress)
func (_Aave *AaveFilterer) WatchLendingRateOracleUpdated(opts *bind.WatchOpts, sink chan<- *AaveLendingRateOracleUpdated, newAddress []common.Address) (event.Subscription, error) {

	var newAddressRule []interface{}
	for _, newAddressItem := range newAddress {
		newAddressRule = append(newAddressRule, newAddressItem)
	}

	logs, sub, err := _Aave.contract.WatchLogs(opts, "LendingRateOracleUpdated", newAddressRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AaveLendingRateOracleUpdated)
				if err := _Aave.contract.UnpackLog(event, "LendingRateOracleUpdated", log); err != nil {
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

// ParseLendingRateOracleUpdated is a log parse operation binding the contract event 0x5c29179aba6942020a8a2d38f65de02fb6b7f784e7f049ed3a3cab97621859b5.
//
// Solidity: event LendingRateOracleUpdated(address indexed newAddress)
func (_Aave *AaveFilterer) ParseLendingRateOracleUpdated(log types.Log) (*AaveLendingRateOracleUpdated, error) {
	event := new(AaveLendingRateOracleUpdated)
	if err := _Aave.contract.UnpackLog(event, "LendingRateOracleUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AaveMarketIdSetIterator is returned from FilterMarketIdSet and is used to iterate over the raw logs and unpacked data for MarketIdSet events raised by the Aave contract.
type AaveMarketIdSetIterator struct {
	Event *AaveMarketIdSet // Event containing the contract specifics and raw log

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
func (it *AaveMarketIdSetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AaveMarketIdSet)
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
		it.Event = new(AaveMarketIdSet)
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
func (it *AaveMarketIdSetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AaveMarketIdSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AaveMarketIdSet represents a MarketIdSet event raised by the Aave contract.
type AaveMarketIdSet struct {
	NewMarketId string
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterMarketIdSet is a free log retrieval operation binding the contract event 0x5e667c32fd847cf8bce48ab3400175cbf107bdc82b2dea62e3364909dfaee799.
//
// Solidity: event MarketIdSet(string newMarketId)
func (_Aave *AaveFilterer) FilterMarketIdSet(opts *bind.FilterOpts) (*AaveMarketIdSetIterator, error) {

	logs, sub, err := _Aave.contract.FilterLogs(opts, "MarketIdSet")
	if err != nil {
		return nil, err
	}
	return &AaveMarketIdSetIterator{contract: _Aave.contract, event: "MarketIdSet", logs: logs, sub: sub}, nil
}

// WatchMarketIdSet is a free log subscription operation binding the contract event 0x5e667c32fd847cf8bce48ab3400175cbf107bdc82b2dea62e3364909dfaee799.
//
// Solidity: event MarketIdSet(string newMarketId)
func (_Aave *AaveFilterer) WatchMarketIdSet(opts *bind.WatchOpts, sink chan<- *AaveMarketIdSet) (event.Subscription, error) {

	logs, sub, err := _Aave.contract.WatchLogs(opts, "MarketIdSet")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AaveMarketIdSet)
				if err := _Aave.contract.UnpackLog(event, "MarketIdSet", log); err != nil {
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

// ParseMarketIdSet is a log parse operation binding the contract event 0x5e667c32fd847cf8bce48ab3400175cbf107bdc82b2dea62e3364909dfaee799.
//
// Solidity: event MarketIdSet(string newMarketId)
func (_Aave *AaveFilterer) ParseMarketIdSet(log types.Log) (*AaveMarketIdSet, error) {
	event := new(AaveMarketIdSet)
	if err := _Aave.contract.UnpackLog(event, "MarketIdSet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AavePriceOracleUpdatedIterator is returned from FilterPriceOracleUpdated and is used to iterate over the raw logs and unpacked data for PriceOracleUpdated events raised by the Aave contract.
type AavePriceOracleUpdatedIterator struct {
	Event *AavePriceOracleUpdated // Event containing the contract specifics and raw log

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
func (it *AavePriceOracleUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AavePriceOracleUpdated)
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
		it.Event = new(AavePriceOracleUpdated)
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
func (it *AavePriceOracleUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AavePriceOracleUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AavePriceOracleUpdated represents a PriceOracleUpdated event raised by the Aave contract.
type AavePriceOracleUpdated struct {
	NewAddress common.Address
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterPriceOracleUpdated is a free log retrieval operation binding the contract event 0xefe8ab924ca486283a79dc604baa67add51afb82af1db8ac386ebbba643cdffd.
//
// Solidity: event PriceOracleUpdated(address indexed newAddress)
func (_Aave *AaveFilterer) FilterPriceOracleUpdated(opts *bind.FilterOpts, newAddress []common.Address) (*AavePriceOracleUpdatedIterator, error) {

	var newAddressRule []interface{}
	for _, newAddressItem := range newAddress {
		newAddressRule = append(newAddressRule, newAddressItem)
	}

	logs, sub, err := _Aave.contract.FilterLogs(opts, "PriceOracleUpdated", newAddressRule)
	if err != nil {
		return nil, err
	}
	return &AavePriceOracleUpdatedIterator{contract: _Aave.contract, event: "PriceOracleUpdated", logs: logs, sub: sub}, nil
}

// WatchPriceOracleUpdated is a free log subscription operation binding the contract event 0xefe8ab924ca486283a79dc604baa67add51afb82af1db8ac386ebbba643cdffd.
//
// Solidity: event PriceOracleUpdated(address indexed newAddress)
func (_Aave *AaveFilterer) WatchPriceOracleUpdated(opts *bind.WatchOpts, sink chan<- *AavePriceOracleUpdated, newAddress []common.Address) (event.Subscription, error) {

	var newAddressRule []interface{}
	for _, newAddressItem := range newAddress {
		newAddressRule = append(newAddressRule, newAddressItem)
	}

	logs, sub, err := _Aave.contract.WatchLogs(opts, "PriceOracleUpdated", newAddressRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AavePriceOracleUpdated)
				if err := _Aave.contract.UnpackLog(event, "PriceOracleUpdated", log); err != nil {
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

// ParsePriceOracleUpdated is a log parse operation binding the contract event 0xefe8ab924ca486283a79dc604baa67add51afb82af1db8ac386ebbba643cdffd.
//
// Solidity: event PriceOracleUpdated(address indexed newAddress)
func (_Aave *AaveFilterer) ParsePriceOracleUpdated(log types.Log) (*AavePriceOracleUpdated, error) {
	event := new(AavePriceOracleUpdated)
	if err := _Aave.contract.UnpackLog(event, "PriceOracleUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AaveProxyCreatedIterator is returned from FilterProxyCreated and is used to iterate over the raw logs and unpacked data for ProxyCreated events raised by the Aave contract.
type AaveProxyCreatedIterator struct {
	Event *AaveProxyCreated // Event containing the contract specifics and raw log

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
func (it *AaveProxyCreatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AaveProxyCreated)
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
		it.Event = new(AaveProxyCreated)
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
func (it *AaveProxyCreatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AaveProxyCreatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AaveProxyCreated represents a ProxyCreated event raised by the Aave contract.
type AaveProxyCreated struct {
	Id         [32]byte
	NewAddress common.Address
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterProxyCreated is a free log retrieval operation binding the contract event 0x1eb35cb4b5bbb23d152f3b4016a5a46c37a07ae930ed0956aba951e231142438.
//
// Solidity: event ProxyCreated(bytes32 id, address indexed newAddress)
func (_Aave *AaveFilterer) FilterProxyCreated(opts *bind.FilterOpts, newAddress []common.Address) (*AaveProxyCreatedIterator, error) {

	var newAddressRule []interface{}
	for _, newAddressItem := range newAddress {
		newAddressRule = append(newAddressRule, newAddressItem)
	}

	logs, sub, err := _Aave.contract.FilterLogs(opts, "ProxyCreated", newAddressRule)
	if err != nil {
		return nil, err
	}
	return &AaveProxyCreatedIterator{contract: _Aave.contract, event: "ProxyCreated", logs: logs, sub: sub}, nil
}

// WatchProxyCreated is a free log subscription operation binding the contract event 0x1eb35cb4b5bbb23d152f3b4016a5a46c37a07ae930ed0956aba951e231142438.
//
// Solidity: event ProxyCreated(bytes32 id, address indexed newAddress)
func (_Aave *AaveFilterer) WatchProxyCreated(opts *bind.WatchOpts, sink chan<- *AaveProxyCreated, newAddress []common.Address) (event.Subscription, error) {

	var newAddressRule []interface{}
	for _, newAddressItem := range newAddress {
		newAddressRule = append(newAddressRule, newAddressItem)
	}

	logs, sub, err := _Aave.contract.WatchLogs(opts, "ProxyCreated", newAddressRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AaveProxyCreated)
				if err := _Aave.contract.UnpackLog(event, "ProxyCreated", log); err != nil {
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

// ParseProxyCreated is a log parse operation binding the contract event 0x1eb35cb4b5bbb23d152f3b4016a5a46c37a07ae930ed0956aba951e231142438.
//
// Solidity: event ProxyCreated(bytes32 id, address indexed newAddress)
func (_Aave *AaveFilterer) ParseProxyCreated(log types.Log) (*AaveProxyCreated, error) {
	event := new(AaveProxyCreated)
	if err := _Aave.contract.UnpackLog(event, "ProxyCreated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
