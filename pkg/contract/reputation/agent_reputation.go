// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package reputation

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
	_ = abi.ConvertType
)

// ReputationMetaData contains all meta data concerning the Reputation contract.
var ReputationMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"target\",\"type\":\"address\"}],\"name\":\"AddressEmptyCode\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"implementation\",\"type\":\"address\"}],\"name\":\"ERC1967InvalidImplementation\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ERC1967NonPayable\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"FailedCall\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidInitialization\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NotInitializing\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"OwnableInvalidOwner\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"OwnableUnauthorizedAccount\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"UUPSUnauthorizedCallContext\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"slot\",\"type\":\"bytes32\"}],\"name\":\"UUPSUnsupportedProxiableUUID\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"agentId\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"clientAddress\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint64\",\"name\":\"feedbackIndex\",\"type\":\"uint64\"}],\"name\":\"FeedbackRevoked\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"version\",\"type\":\"uint64\"}],\"name\":\"Initialized\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"agentId\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"clientAddress\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"feedbackIndex\",\"type\":\"uint64\"},{\"indexed\":false,\"internalType\":\"int128\",\"name\":\"value\",\"type\":\"int128\"},{\"indexed\":false,\"internalType\":\"uint8\",\"name\":\"valueDecimals\",\"type\":\"uint8\"},{\"indexed\":true,\"internalType\":\"string\",\"name\":\"indexedTag1\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"tag1\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"tag2\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"endpoint\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"feedbackURI\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"feedbackHash\",\"type\":\"bytes32\"}],\"name\":\"NewFeedback\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"agentId\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"clientAddress\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"feedbackIndex\",\"type\":\"uint64\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"responder\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"responseURI\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"responseHash\",\"type\":\"bytes32\"}],\"name\":\"ResponseAppended\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"implementation\",\"type\":\"address\"}],\"name\":\"Upgraded\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"UPGRADE_INTERFACE_VERSION\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"agentId\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"clientAddress\",\"type\":\"address\"},{\"internalType\":\"uint64\",\"name\":\"feedbackIndex\",\"type\":\"uint64\"},{\"internalType\":\"string\",\"name\":\"responseURI\",\"type\":\"string\"},{\"internalType\":\"bytes32\",\"name\":\"responseHash\",\"type\":\"bytes32\"}],\"name\":\"appendResponse\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"agentId\",\"type\":\"uint256\"}],\"name\":\"getClients\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"\",\"type\":\"address[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getIdentityRegistry\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"agentId\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"clientAddress\",\"type\":\"address\"}],\"name\":\"getLastIndex\",\"outputs\":[{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"agentId\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"clientAddress\",\"type\":\"address\"},{\"internalType\":\"uint64\",\"name\":\"feedbackIndex\",\"type\":\"uint64\"},{\"internalType\":\"address[]\",\"name\":\"responders\",\"type\":\"address[]\"}],\"name\":\"getResponseCount\",\"outputs\":[{\"internalType\":\"uint64\",\"name\":\"count\",\"type\":\"uint64\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"agentId\",\"type\":\"uint256\"},{\"internalType\":\"address[]\",\"name\":\"clientAddresses\",\"type\":\"address[]\"},{\"internalType\":\"string\",\"name\":\"tag1\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"tag2\",\"type\":\"string\"}],\"name\":\"getSummary\",\"outputs\":[{\"internalType\":\"uint64\",\"name\":\"count\",\"type\":\"uint64\"},{\"internalType\":\"int128\",\"name\":\"summaryValue\",\"type\":\"int128\"},{\"internalType\":\"uint8\",\"name\":\"summaryValueDecimals\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getVersion\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"agentId\",\"type\":\"uint256\"},{\"internalType\":\"int128\",\"name\":\"value\",\"type\":\"int128\"},{\"internalType\":\"uint8\",\"name\":\"valueDecimals\",\"type\":\"uint8\"},{\"internalType\":\"string\",\"name\":\"tag1\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"tag2\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"endpoint\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"feedbackURI\",\"type\":\"string\"},{\"internalType\":\"bytes32\",\"name\":\"feedbackHash\",\"type\":\"bytes32\"}],\"name\":\"giveFeedback\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"identityRegistry_\",\"type\":\"address\"}],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"proxiableUUID\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"agentId\",\"type\":\"uint256\"},{\"internalType\":\"address[]\",\"name\":\"clientAddresses\",\"type\":\"address[]\"},{\"internalType\":\"string\",\"name\":\"tag1\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"tag2\",\"type\":\"string\"},{\"internalType\":\"bool\",\"name\":\"includeRevoked\",\"type\":\"bool\"}],\"name\":\"readAllFeedback\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"clients\",\"type\":\"address[]\"},{\"internalType\":\"uint64[]\",\"name\":\"feedbackIndexes\",\"type\":\"uint64[]\"},{\"internalType\":\"int128[]\",\"name\":\"values\",\"type\":\"int128[]\"},{\"internalType\":\"uint8[]\",\"name\":\"valueDecimals\",\"type\":\"uint8[]\"},{\"internalType\":\"string[]\",\"name\":\"tag1s\",\"type\":\"string[]\"},{\"internalType\":\"string[]\",\"name\":\"tag2s\",\"type\":\"string[]\"},{\"internalType\":\"bool[]\",\"name\":\"revokedStatuses\",\"type\":\"bool[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"agentId\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"clientAddress\",\"type\":\"address\"},{\"internalType\":\"uint64\",\"name\":\"feedbackIndex\",\"type\":\"uint64\"}],\"name\":\"readFeedback\",\"outputs\":[{\"internalType\":\"int128\",\"name\":\"value\",\"type\":\"int128\"},{\"internalType\":\"uint8\",\"name\":\"valueDecimals\",\"type\":\"uint8\"},{\"internalType\":\"string\",\"name\":\"tag1\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"tag2\",\"type\":\"string\"},{\"internalType\":\"bool\",\"name\":\"isRevoked\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"agentId\",\"type\":\"uint256\"},{\"internalType\":\"uint64\",\"name\":\"feedbackIndex\",\"type\":\"uint64\"}],\"name\":\"revokeFeedback\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newImplementation\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"upgradeToAndCall\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"}]",
}

// ReputationABI is the input ABI used to generate the binding from.
// Deprecated: Use ReputationMetaData.ABI instead.
var ReputationABI = ReputationMetaData.ABI

// Reputation is an auto generated Go binding around an Ethereum contract.
type Reputation struct {
	ReputationCaller     // Read-only binding to the contract
	ReputationTransactor // Write-only binding to the contract
	ReputationFilterer   // Log filterer for contract events
}

// ReputationCaller is an auto generated read-only Go binding around an Ethereum contract.
type ReputationCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ReputationTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ReputationTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ReputationFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ReputationFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ReputationSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ReputationSession struct {
	Contract     *Reputation       // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ReputationCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ReputationCallerSession struct {
	Contract *ReputationCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts     // Call options to use throughout this session
}

// ReputationTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ReputationTransactorSession struct {
	Contract     *ReputationTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts     // Transaction auth options to use throughout this session
}

// ReputationRaw is an auto generated low-level Go binding around an Ethereum contract.
type ReputationRaw struct {
	Contract *Reputation // Generic contract binding to access the raw methods on
}

// ReputationCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ReputationCallerRaw struct {
	Contract *ReputationCaller // Generic read-only contract binding to access the raw methods on
}

// ReputationTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ReputationTransactorRaw struct {
	Contract *ReputationTransactor // Generic write-only contract binding to access the raw methods on
}

// NewReputation creates a new instance of Reputation, bound to a specific deployed contract.
func NewReputation(address common.Address, backend bind.ContractBackend) (*Reputation, error) {
	contract, err := bindReputation(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Reputation{ReputationCaller: ReputationCaller{contract: contract}, ReputationTransactor: ReputationTransactor{contract: contract}, ReputationFilterer: ReputationFilterer{contract: contract}}, nil
}

// NewReputationCaller creates a new read-only instance of Reputation, bound to a specific deployed contract.
func NewReputationCaller(address common.Address, caller bind.ContractCaller) (*ReputationCaller, error) {
	contract, err := bindReputation(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ReputationCaller{contract: contract}, nil
}

// NewReputationTransactor creates a new write-only instance of Reputation, bound to a specific deployed contract.
func NewReputationTransactor(address common.Address, transactor bind.ContractTransactor) (*ReputationTransactor, error) {
	contract, err := bindReputation(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ReputationTransactor{contract: contract}, nil
}

// NewReputationFilterer creates a new log filterer instance of Reputation, bound to a specific deployed contract.
func NewReputationFilterer(address common.Address, filterer bind.ContractFilterer) (*ReputationFilterer, error) {
	contract, err := bindReputation(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ReputationFilterer{contract: contract}, nil
}

// bindReputation binds a generic wrapper to an already deployed contract.
func bindReputation(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := ReputationMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Reputation *ReputationRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Reputation.Contract.ReputationCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Reputation *ReputationRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Reputation.Contract.ReputationTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Reputation *ReputationRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Reputation.Contract.ReputationTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Reputation *ReputationCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Reputation.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Reputation *ReputationTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Reputation.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Reputation *ReputationTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Reputation.Contract.contract.Transact(opts, method, params...)
}

// UPGRADEINTERFACEVERSION is a free data retrieval call binding the contract method 0xad3cb1cc.
//
// Solidity: function UPGRADE_INTERFACE_VERSION() view returns(string)
func (_Reputation *ReputationCaller) UPGRADEINTERFACEVERSION(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _Reputation.contract.Call(opts, &out, "UPGRADE_INTERFACE_VERSION")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// UPGRADEINTERFACEVERSION is a free data retrieval call binding the contract method 0xad3cb1cc.
//
// Solidity: function UPGRADE_INTERFACE_VERSION() view returns(string)
func (_Reputation *ReputationSession) UPGRADEINTERFACEVERSION() (string, error) {
	return _Reputation.Contract.UPGRADEINTERFACEVERSION(&_Reputation.CallOpts)
}

// UPGRADEINTERFACEVERSION is a free data retrieval call binding the contract method 0xad3cb1cc.
//
// Solidity: function UPGRADE_INTERFACE_VERSION() view returns(string)
func (_Reputation *ReputationCallerSession) UPGRADEINTERFACEVERSION() (string, error) {
	return _Reputation.Contract.UPGRADEINTERFACEVERSION(&_Reputation.CallOpts)
}

// GetClients is a free data retrieval call binding the contract method 0x42dd519c.
//
// Solidity: function getClients(uint256 agentId) view returns(address[])
func (_Reputation *ReputationCaller) GetClients(opts *bind.CallOpts, agentId *big.Int) ([]common.Address, error) {
	var out []interface{}
	err := _Reputation.contract.Call(opts, &out, "getClients", agentId)

	if err != nil {
		return *new([]common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)

	return out0, err

}

// GetClients is a free data retrieval call binding the contract method 0x42dd519c.
//
// Solidity: function getClients(uint256 agentId) view returns(address[])
func (_Reputation *ReputationSession) GetClients(agentId *big.Int) ([]common.Address, error) {
	return _Reputation.Contract.GetClients(&_Reputation.CallOpts, agentId)
}

// GetClients is a free data retrieval call binding the contract method 0x42dd519c.
//
// Solidity: function getClients(uint256 agentId) view returns(address[])
func (_Reputation *ReputationCallerSession) GetClients(agentId *big.Int) ([]common.Address, error) {
	return _Reputation.Contract.GetClients(&_Reputation.CallOpts, agentId)
}

// GetIdentityRegistry is a free data retrieval call binding the contract method 0xbc4d861b.
//
// Solidity: function getIdentityRegistry() view returns(address)
func (_Reputation *ReputationCaller) GetIdentityRegistry(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Reputation.contract.Call(opts, &out, "getIdentityRegistry")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetIdentityRegistry is a free data retrieval call binding the contract method 0xbc4d861b.
//
// Solidity: function getIdentityRegistry() view returns(address)
func (_Reputation *ReputationSession) GetIdentityRegistry() (common.Address, error) {
	return _Reputation.Contract.GetIdentityRegistry(&_Reputation.CallOpts)
}

// GetIdentityRegistry is a free data retrieval call binding the contract method 0xbc4d861b.
//
// Solidity: function getIdentityRegistry() view returns(address)
func (_Reputation *ReputationCallerSession) GetIdentityRegistry() (common.Address, error) {
	return _Reputation.Contract.GetIdentityRegistry(&_Reputation.CallOpts)
}

// GetLastIndex is a free data retrieval call binding the contract method 0xf2d81759.
//
// Solidity: function getLastIndex(uint256 agentId, address clientAddress) view returns(uint64)
func (_Reputation *ReputationCaller) GetLastIndex(opts *bind.CallOpts, agentId *big.Int, clientAddress common.Address) (uint64, error) {
	var out []interface{}
	err := _Reputation.contract.Call(opts, &out, "getLastIndex", agentId, clientAddress)

	if err != nil {
		return *new(uint64), err
	}

	out0 := *abi.ConvertType(out[0], new(uint64)).(*uint64)

	return out0, err

}

// GetLastIndex is a free data retrieval call binding the contract method 0xf2d81759.
//
// Solidity: function getLastIndex(uint256 agentId, address clientAddress) view returns(uint64)
func (_Reputation *ReputationSession) GetLastIndex(agentId *big.Int, clientAddress common.Address) (uint64, error) {
	return _Reputation.Contract.GetLastIndex(&_Reputation.CallOpts, agentId, clientAddress)
}

// GetLastIndex is a free data retrieval call binding the contract method 0xf2d81759.
//
// Solidity: function getLastIndex(uint256 agentId, address clientAddress) view returns(uint64)
func (_Reputation *ReputationCallerSession) GetLastIndex(agentId *big.Int, clientAddress common.Address) (uint64, error) {
	return _Reputation.Contract.GetLastIndex(&_Reputation.CallOpts, agentId, clientAddress)
}

// GetResponseCount is a free data retrieval call binding the contract method 0x6e04cacd.
//
// Solidity: function getResponseCount(uint256 agentId, address clientAddress, uint64 feedbackIndex, address[] responders) view returns(uint64 count)
func (_Reputation *ReputationCaller) GetResponseCount(opts *bind.CallOpts, agentId *big.Int, clientAddress common.Address, feedbackIndex uint64, responders []common.Address) (uint64, error) {
	var out []interface{}
	err := _Reputation.contract.Call(opts, &out, "getResponseCount", agentId, clientAddress, feedbackIndex, responders)

	if err != nil {
		return *new(uint64), err
	}

	out0 := *abi.ConvertType(out[0], new(uint64)).(*uint64)

	return out0, err

}

// GetResponseCount is a free data retrieval call binding the contract method 0x6e04cacd.
//
// Solidity: function getResponseCount(uint256 agentId, address clientAddress, uint64 feedbackIndex, address[] responders) view returns(uint64 count)
func (_Reputation *ReputationSession) GetResponseCount(agentId *big.Int, clientAddress common.Address, feedbackIndex uint64, responders []common.Address) (uint64, error) {
	return _Reputation.Contract.GetResponseCount(&_Reputation.CallOpts, agentId, clientAddress, feedbackIndex, responders)
}

// GetResponseCount is a free data retrieval call binding the contract method 0x6e04cacd.
//
// Solidity: function getResponseCount(uint256 agentId, address clientAddress, uint64 feedbackIndex, address[] responders) view returns(uint64 count)
func (_Reputation *ReputationCallerSession) GetResponseCount(agentId *big.Int, clientAddress common.Address, feedbackIndex uint64, responders []common.Address) (uint64, error) {
	return _Reputation.Contract.GetResponseCount(&_Reputation.CallOpts, agentId, clientAddress, feedbackIndex, responders)
}

// GetSummary is a free data retrieval call binding the contract method 0x81bbba58.
//
// Solidity: function getSummary(uint256 agentId, address[] clientAddresses, string tag1, string tag2) view returns(uint64 count, int128 summaryValue, uint8 summaryValueDecimals)
func (_Reputation *ReputationCaller) GetSummary(opts *bind.CallOpts, agentId *big.Int, clientAddresses []common.Address, tag1 string, tag2 string) (struct {
	Count                uint64
	SummaryValue         *big.Int
	SummaryValueDecimals uint8
}, error) {
	var out []interface{}
	err := _Reputation.contract.Call(opts, &out, "getSummary", agentId, clientAddresses, tag1, tag2)

	outstruct := new(struct {
		Count                uint64
		SummaryValue         *big.Int
		SummaryValueDecimals uint8
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Count = *abi.ConvertType(out[0], new(uint64)).(*uint64)
	outstruct.SummaryValue = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)
	outstruct.SummaryValueDecimals = *abi.ConvertType(out[2], new(uint8)).(*uint8)

	return *outstruct, err

}

// GetSummary is a free data retrieval call binding the contract method 0x81bbba58.
//
// Solidity: function getSummary(uint256 agentId, address[] clientAddresses, string tag1, string tag2) view returns(uint64 count, int128 summaryValue, uint8 summaryValueDecimals)
func (_Reputation *ReputationSession) GetSummary(agentId *big.Int, clientAddresses []common.Address, tag1 string, tag2 string) (struct {
	Count                uint64
	SummaryValue         *big.Int
	SummaryValueDecimals uint8
}, error) {
	return _Reputation.Contract.GetSummary(&_Reputation.CallOpts, agentId, clientAddresses, tag1, tag2)
}

// GetSummary is a free data retrieval call binding the contract method 0x81bbba58.
//
// Solidity: function getSummary(uint256 agentId, address[] clientAddresses, string tag1, string tag2) view returns(uint64 count, int128 summaryValue, uint8 summaryValueDecimals)
func (_Reputation *ReputationCallerSession) GetSummary(agentId *big.Int, clientAddresses []common.Address, tag1 string, tag2 string) (struct {
	Count                uint64
	SummaryValue         *big.Int
	SummaryValueDecimals uint8
}, error) {
	return _Reputation.Contract.GetSummary(&_Reputation.CallOpts, agentId, clientAddresses, tag1, tag2)
}

// GetVersion is a free data retrieval call binding the contract method 0x0d8e6e2c.
//
// Solidity: function getVersion() pure returns(string)
func (_Reputation *ReputationCaller) GetVersion(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _Reputation.contract.Call(opts, &out, "getVersion")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// GetVersion is a free data retrieval call binding the contract method 0x0d8e6e2c.
//
// Solidity: function getVersion() pure returns(string)
func (_Reputation *ReputationSession) GetVersion() (string, error) {
	return _Reputation.Contract.GetVersion(&_Reputation.CallOpts)
}

// GetVersion is a free data retrieval call binding the contract method 0x0d8e6e2c.
//
// Solidity: function getVersion() pure returns(string)
func (_Reputation *ReputationCallerSession) GetVersion() (string, error) {
	return _Reputation.Contract.GetVersion(&_Reputation.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Reputation *ReputationCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Reputation.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Reputation *ReputationSession) Owner() (common.Address, error) {
	return _Reputation.Contract.Owner(&_Reputation.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Reputation *ReputationCallerSession) Owner() (common.Address, error) {
	return _Reputation.Contract.Owner(&_Reputation.CallOpts)
}

// ProxiableUUID is a free data retrieval call binding the contract method 0x52d1902d.
//
// Solidity: function proxiableUUID() view returns(bytes32)
func (_Reputation *ReputationCaller) ProxiableUUID(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _Reputation.contract.Call(opts, &out, "proxiableUUID")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// ProxiableUUID is a free data retrieval call binding the contract method 0x52d1902d.
//
// Solidity: function proxiableUUID() view returns(bytes32)
func (_Reputation *ReputationSession) ProxiableUUID() ([32]byte, error) {
	return _Reputation.Contract.ProxiableUUID(&_Reputation.CallOpts)
}

// ProxiableUUID is a free data retrieval call binding the contract method 0x52d1902d.
//
// Solidity: function proxiableUUID() view returns(bytes32)
func (_Reputation *ReputationCallerSession) ProxiableUUID() ([32]byte, error) {
	return _Reputation.Contract.ProxiableUUID(&_Reputation.CallOpts)
}

// ReadAllFeedback is a free data retrieval call binding the contract method 0xd9d84224.
//
// Solidity: function readAllFeedback(uint256 agentId, address[] clientAddresses, string tag1, string tag2, bool includeRevoked) view returns(address[] clients, uint64[] feedbackIndexes, int128[] values, uint8[] valueDecimals, string[] tag1s, string[] tag2s, bool[] revokedStatuses)
func (_Reputation *ReputationCaller) ReadAllFeedback(opts *bind.CallOpts, agentId *big.Int, clientAddresses []common.Address, tag1 string, tag2 string, includeRevoked bool) (struct {
	Clients         []common.Address
	FeedbackIndexes []uint64
	Values          []*big.Int
	ValueDecimals   []uint8
	Tag1s           []string
	Tag2s           []string
	RevokedStatuses []bool
}, error) {
	var out []interface{}
	err := _Reputation.contract.Call(opts, &out, "readAllFeedback", agentId, clientAddresses, tag1, tag2, includeRevoked)

	outstruct := new(struct {
		Clients         []common.Address
		FeedbackIndexes []uint64
		Values          []*big.Int
		ValueDecimals   []uint8
		Tag1s           []string
		Tag2s           []string
		RevokedStatuses []bool
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Clients = *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)
	outstruct.FeedbackIndexes = *abi.ConvertType(out[1], new([]uint64)).(*[]uint64)
	outstruct.Values = *abi.ConvertType(out[2], new([]*big.Int)).(*[]*big.Int)
	outstruct.ValueDecimals = *abi.ConvertType(out[3], new([]uint8)).(*[]uint8)
	outstruct.Tag1s = *abi.ConvertType(out[4], new([]string)).(*[]string)
	outstruct.Tag2s = *abi.ConvertType(out[5], new([]string)).(*[]string)
	outstruct.RevokedStatuses = *abi.ConvertType(out[6], new([]bool)).(*[]bool)

	return *outstruct, err

}

// ReadAllFeedback is a free data retrieval call binding the contract method 0xd9d84224.
//
// Solidity: function readAllFeedback(uint256 agentId, address[] clientAddresses, string tag1, string tag2, bool includeRevoked) view returns(address[] clients, uint64[] feedbackIndexes, int128[] values, uint8[] valueDecimals, string[] tag1s, string[] tag2s, bool[] revokedStatuses)
func (_Reputation *ReputationSession) ReadAllFeedback(agentId *big.Int, clientAddresses []common.Address, tag1 string, tag2 string, includeRevoked bool) (struct {
	Clients         []common.Address
	FeedbackIndexes []uint64
	Values          []*big.Int
	ValueDecimals   []uint8
	Tag1s           []string
	Tag2s           []string
	RevokedStatuses []bool
}, error) {
	return _Reputation.Contract.ReadAllFeedback(&_Reputation.CallOpts, agentId, clientAddresses, tag1, tag2, includeRevoked)
}

// ReadAllFeedback is a free data retrieval call binding the contract method 0xd9d84224.
//
// Solidity: function readAllFeedback(uint256 agentId, address[] clientAddresses, string tag1, string tag2, bool includeRevoked) view returns(address[] clients, uint64[] feedbackIndexes, int128[] values, uint8[] valueDecimals, string[] tag1s, string[] tag2s, bool[] revokedStatuses)
func (_Reputation *ReputationCallerSession) ReadAllFeedback(agentId *big.Int, clientAddresses []common.Address, tag1 string, tag2 string, includeRevoked bool) (struct {
	Clients         []common.Address
	FeedbackIndexes []uint64
	Values          []*big.Int
	ValueDecimals   []uint8
	Tag1s           []string
	Tag2s           []string
	RevokedStatuses []bool
}, error) {
	return _Reputation.Contract.ReadAllFeedback(&_Reputation.CallOpts, agentId, clientAddresses, tag1, tag2, includeRevoked)
}

// ReadFeedback is a free data retrieval call binding the contract method 0x232b0810.
//
// Solidity: function readFeedback(uint256 agentId, address clientAddress, uint64 feedbackIndex) view returns(int128 value, uint8 valueDecimals, string tag1, string tag2, bool isRevoked)
func (_Reputation *ReputationCaller) ReadFeedback(opts *bind.CallOpts, agentId *big.Int, clientAddress common.Address, feedbackIndex uint64) (struct {
	Value         *big.Int
	ValueDecimals uint8
	Tag1          string
	Tag2          string
	IsRevoked     bool
}, error) {
	var out []interface{}
	err := _Reputation.contract.Call(opts, &out, "readFeedback", agentId, clientAddress, feedbackIndex)

	outstruct := new(struct {
		Value         *big.Int
		ValueDecimals uint8
		Tag1          string
		Tag2          string
		IsRevoked     bool
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Value = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.ValueDecimals = *abi.ConvertType(out[1], new(uint8)).(*uint8)
	outstruct.Tag1 = *abi.ConvertType(out[2], new(string)).(*string)
	outstruct.Tag2 = *abi.ConvertType(out[3], new(string)).(*string)
	outstruct.IsRevoked = *abi.ConvertType(out[4], new(bool)).(*bool)

	return *outstruct, err

}

// ReadFeedback is a free data retrieval call binding the contract method 0x232b0810.
//
// Solidity: function readFeedback(uint256 agentId, address clientAddress, uint64 feedbackIndex) view returns(int128 value, uint8 valueDecimals, string tag1, string tag2, bool isRevoked)
func (_Reputation *ReputationSession) ReadFeedback(agentId *big.Int, clientAddress common.Address, feedbackIndex uint64) (struct {
	Value         *big.Int
	ValueDecimals uint8
	Tag1          string
	Tag2          string
	IsRevoked     bool
}, error) {
	return _Reputation.Contract.ReadFeedback(&_Reputation.CallOpts, agentId, clientAddress, feedbackIndex)
}

// ReadFeedback is a free data retrieval call binding the contract method 0x232b0810.
//
// Solidity: function readFeedback(uint256 agentId, address clientAddress, uint64 feedbackIndex) view returns(int128 value, uint8 valueDecimals, string tag1, string tag2, bool isRevoked)
func (_Reputation *ReputationCallerSession) ReadFeedback(agentId *big.Int, clientAddress common.Address, feedbackIndex uint64) (struct {
	Value         *big.Int
	ValueDecimals uint8
	Tag1          string
	Tag2          string
	IsRevoked     bool
}, error) {
	return _Reputation.Contract.ReadFeedback(&_Reputation.CallOpts, agentId, clientAddress, feedbackIndex)
}

// AppendResponse is a paid mutator transaction binding the contract method 0xc2349ab2.
//
// Solidity: function appendResponse(uint256 agentId, address clientAddress, uint64 feedbackIndex, string responseURI, bytes32 responseHash) returns()
func (_Reputation *ReputationTransactor) AppendResponse(opts *bind.TransactOpts, agentId *big.Int, clientAddress common.Address, feedbackIndex uint64, responseURI string, responseHash [32]byte) (*types.Transaction, error) {
	return _Reputation.contract.Transact(opts, "appendResponse", agentId, clientAddress, feedbackIndex, responseURI, responseHash)
}

// AppendResponse is a paid mutator transaction binding the contract method 0xc2349ab2.
//
// Solidity: function appendResponse(uint256 agentId, address clientAddress, uint64 feedbackIndex, string responseURI, bytes32 responseHash) returns()
func (_Reputation *ReputationSession) AppendResponse(agentId *big.Int, clientAddress common.Address, feedbackIndex uint64, responseURI string, responseHash [32]byte) (*types.Transaction, error) {
	return _Reputation.Contract.AppendResponse(&_Reputation.TransactOpts, agentId, clientAddress, feedbackIndex, responseURI, responseHash)
}

// AppendResponse is a paid mutator transaction binding the contract method 0xc2349ab2.
//
// Solidity: function appendResponse(uint256 agentId, address clientAddress, uint64 feedbackIndex, string responseURI, bytes32 responseHash) returns()
func (_Reputation *ReputationTransactorSession) AppendResponse(agentId *big.Int, clientAddress common.Address, feedbackIndex uint64, responseURI string, responseHash [32]byte) (*types.Transaction, error) {
	return _Reputation.Contract.AppendResponse(&_Reputation.TransactOpts, agentId, clientAddress, feedbackIndex, responseURI, responseHash)
}

// GiveFeedback is a paid mutator transaction binding the contract method 0x3c036a7e.
//
// Solidity: function giveFeedback(uint256 agentId, int128 value, uint8 valueDecimals, string tag1, string tag2, string endpoint, string feedbackURI, bytes32 feedbackHash) returns()
func (_Reputation *ReputationTransactor) GiveFeedback(opts *bind.TransactOpts, agentId *big.Int, value *big.Int, valueDecimals uint8, tag1 string, tag2 string, endpoint string, feedbackURI string, feedbackHash [32]byte) (*types.Transaction, error) {
	return _Reputation.contract.Transact(opts, "giveFeedback", agentId, value, valueDecimals, tag1, tag2, endpoint, feedbackURI, feedbackHash)
}

// GiveFeedback is a paid mutator transaction binding the contract method 0x3c036a7e.
//
// Solidity: function giveFeedback(uint256 agentId, int128 value, uint8 valueDecimals, string tag1, string tag2, string endpoint, string feedbackURI, bytes32 feedbackHash) returns()
func (_Reputation *ReputationSession) GiveFeedback(agentId *big.Int, value *big.Int, valueDecimals uint8, tag1 string, tag2 string, endpoint string, feedbackURI string, feedbackHash [32]byte) (*types.Transaction, error) {
	return _Reputation.Contract.GiveFeedback(&_Reputation.TransactOpts, agentId, value, valueDecimals, tag1, tag2, endpoint, feedbackURI, feedbackHash)
}

// GiveFeedback is a paid mutator transaction binding the contract method 0x3c036a7e.
//
// Solidity: function giveFeedback(uint256 agentId, int128 value, uint8 valueDecimals, string tag1, string tag2, string endpoint, string feedbackURI, bytes32 feedbackHash) returns()
func (_Reputation *ReputationTransactorSession) GiveFeedback(agentId *big.Int, value *big.Int, valueDecimals uint8, tag1 string, tag2 string, endpoint string, feedbackURI string, feedbackHash [32]byte) (*types.Transaction, error) {
	return _Reputation.Contract.GiveFeedback(&_Reputation.TransactOpts, agentId, value, valueDecimals, tag1, tag2, endpoint, feedbackURI, feedbackHash)
}

// Initialize is a paid mutator transaction binding the contract method 0xc4d66de8.
//
// Solidity: function initialize(address identityRegistry_) returns()
func (_Reputation *ReputationTransactor) Initialize(opts *bind.TransactOpts, identityRegistry_ common.Address) (*types.Transaction, error) {
	return _Reputation.contract.Transact(opts, "initialize", identityRegistry_)
}

// Initialize is a paid mutator transaction binding the contract method 0xc4d66de8.
//
// Solidity: function initialize(address identityRegistry_) returns()
func (_Reputation *ReputationSession) Initialize(identityRegistry_ common.Address) (*types.Transaction, error) {
	return _Reputation.Contract.Initialize(&_Reputation.TransactOpts, identityRegistry_)
}

// Initialize is a paid mutator transaction binding the contract method 0xc4d66de8.
//
// Solidity: function initialize(address identityRegistry_) returns()
func (_Reputation *ReputationTransactorSession) Initialize(identityRegistry_ common.Address) (*types.Transaction, error) {
	return _Reputation.Contract.Initialize(&_Reputation.TransactOpts, identityRegistry_)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Reputation *ReputationTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Reputation.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Reputation *ReputationSession) RenounceOwnership() (*types.Transaction, error) {
	return _Reputation.Contract.RenounceOwnership(&_Reputation.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Reputation *ReputationTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _Reputation.Contract.RenounceOwnership(&_Reputation.TransactOpts)
}

// RevokeFeedback is a paid mutator transaction binding the contract method 0x4ab3ca99.
//
// Solidity: function revokeFeedback(uint256 agentId, uint64 feedbackIndex) returns()
func (_Reputation *ReputationTransactor) RevokeFeedback(opts *bind.TransactOpts, agentId *big.Int, feedbackIndex uint64) (*types.Transaction, error) {
	return _Reputation.contract.Transact(opts, "revokeFeedback", agentId, feedbackIndex)
}

// RevokeFeedback is a paid mutator transaction binding the contract method 0x4ab3ca99.
//
// Solidity: function revokeFeedback(uint256 agentId, uint64 feedbackIndex) returns()
func (_Reputation *ReputationSession) RevokeFeedback(agentId *big.Int, feedbackIndex uint64) (*types.Transaction, error) {
	return _Reputation.Contract.RevokeFeedback(&_Reputation.TransactOpts, agentId, feedbackIndex)
}

// RevokeFeedback is a paid mutator transaction binding the contract method 0x4ab3ca99.
//
// Solidity: function revokeFeedback(uint256 agentId, uint64 feedbackIndex) returns()
func (_Reputation *ReputationTransactorSession) RevokeFeedback(agentId *big.Int, feedbackIndex uint64) (*types.Transaction, error) {
	return _Reputation.Contract.RevokeFeedback(&_Reputation.TransactOpts, agentId, feedbackIndex)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Reputation *ReputationTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _Reputation.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Reputation *ReputationSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Reputation.Contract.TransferOwnership(&_Reputation.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Reputation *ReputationTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Reputation.Contract.TransferOwnership(&_Reputation.TransactOpts, newOwner)
}

// UpgradeToAndCall is a paid mutator transaction binding the contract method 0x4f1ef286.
//
// Solidity: function upgradeToAndCall(address newImplementation, bytes data) payable returns()
func (_Reputation *ReputationTransactor) UpgradeToAndCall(opts *bind.TransactOpts, newImplementation common.Address, data []byte) (*types.Transaction, error) {
	return _Reputation.contract.Transact(opts, "upgradeToAndCall", newImplementation, data)
}

// UpgradeToAndCall is a paid mutator transaction binding the contract method 0x4f1ef286.
//
// Solidity: function upgradeToAndCall(address newImplementation, bytes data) payable returns()
func (_Reputation *ReputationSession) UpgradeToAndCall(newImplementation common.Address, data []byte) (*types.Transaction, error) {
	return _Reputation.Contract.UpgradeToAndCall(&_Reputation.TransactOpts, newImplementation, data)
}

// UpgradeToAndCall is a paid mutator transaction binding the contract method 0x4f1ef286.
//
// Solidity: function upgradeToAndCall(address newImplementation, bytes data) payable returns()
func (_Reputation *ReputationTransactorSession) UpgradeToAndCall(newImplementation common.Address, data []byte) (*types.Transaction, error) {
	return _Reputation.Contract.UpgradeToAndCall(&_Reputation.TransactOpts, newImplementation, data)
}

// ReputationFeedbackRevokedIterator is returned from FilterFeedbackRevoked and is used to iterate over the raw logs and unpacked data for FeedbackRevoked events raised by the Reputation contract.
type ReputationFeedbackRevokedIterator struct {
	Event *ReputationFeedbackRevoked // Event containing the contract specifics and raw log

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
func (it *ReputationFeedbackRevokedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ReputationFeedbackRevoked)
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
		it.Event = new(ReputationFeedbackRevoked)
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
func (it *ReputationFeedbackRevokedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ReputationFeedbackRevokedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ReputationFeedbackRevoked represents a FeedbackRevoked event raised by the Reputation contract.
type ReputationFeedbackRevoked struct {
	AgentId       *big.Int
	ClientAddress common.Address
	FeedbackIndex uint64
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterFeedbackRevoked is a free log retrieval operation binding the contract event 0x25156fd3288212246d8b008d5921fde376c71ed14ac2e072a506eb06fde6d09d.
//
// Solidity: event FeedbackRevoked(uint256 indexed agentId, address indexed clientAddress, uint64 indexed feedbackIndex)
func (_Reputation *ReputationFilterer) FilterFeedbackRevoked(opts *bind.FilterOpts, agentId []*big.Int, clientAddress []common.Address, feedbackIndex []uint64) (*ReputationFeedbackRevokedIterator, error) {

	var agentIdRule []interface{}
	for _, agentIdItem := range agentId {
		agentIdRule = append(agentIdRule, agentIdItem)
	}
	var clientAddressRule []interface{}
	for _, clientAddressItem := range clientAddress {
		clientAddressRule = append(clientAddressRule, clientAddressItem)
	}
	var feedbackIndexRule []interface{}
	for _, feedbackIndexItem := range feedbackIndex {
		feedbackIndexRule = append(feedbackIndexRule, feedbackIndexItem)
	}

	logs, sub, err := _Reputation.contract.FilterLogs(opts, "FeedbackRevoked", agentIdRule, clientAddressRule, feedbackIndexRule)
	if err != nil {
		return nil, err
	}
	return &ReputationFeedbackRevokedIterator{contract: _Reputation.contract, event: "FeedbackRevoked", logs: logs, sub: sub}, nil
}

// WatchFeedbackRevoked is a free log subscription operation binding the contract event 0x25156fd3288212246d8b008d5921fde376c71ed14ac2e072a506eb06fde6d09d.
//
// Solidity: event FeedbackRevoked(uint256 indexed agentId, address indexed clientAddress, uint64 indexed feedbackIndex)
func (_Reputation *ReputationFilterer) WatchFeedbackRevoked(opts *bind.WatchOpts, sink chan<- *ReputationFeedbackRevoked, agentId []*big.Int, clientAddress []common.Address, feedbackIndex []uint64) (event.Subscription, error) {

	var agentIdRule []interface{}
	for _, agentIdItem := range agentId {
		agentIdRule = append(agentIdRule, agentIdItem)
	}
	var clientAddressRule []interface{}
	for _, clientAddressItem := range clientAddress {
		clientAddressRule = append(clientAddressRule, clientAddressItem)
	}
	var feedbackIndexRule []interface{}
	for _, feedbackIndexItem := range feedbackIndex {
		feedbackIndexRule = append(feedbackIndexRule, feedbackIndexItem)
	}

	logs, sub, err := _Reputation.contract.WatchLogs(opts, "FeedbackRevoked", agentIdRule, clientAddressRule, feedbackIndexRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ReputationFeedbackRevoked)
				if err := _Reputation.contract.UnpackLog(event, "FeedbackRevoked", log); err != nil {
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

// ParseFeedbackRevoked is a log parse operation binding the contract event 0x25156fd3288212246d8b008d5921fde376c71ed14ac2e072a506eb06fde6d09d.
//
// Solidity: event FeedbackRevoked(uint256 indexed agentId, address indexed clientAddress, uint64 indexed feedbackIndex)
func (_Reputation *ReputationFilterer) ParseFeedbackRevoked(log types.Log) (*ReputationFeedbackRevoked, error) {
	event := new(ReputationFeedbackRevoked)
	if err := _Reputation.contract.UnpackLog(event, "FeedbackRevoked", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ReputationInitializedIterator is returned from FilterInitialized and is used to iterate over the raw logs and unpacked data for Initialized events raised by the Reputation contract.
type ReputationInitializedIterator struct {
	Event *ReputationInitialized // Event containing the contract specifics and raw log

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
func (it *ReputationInitializedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ReputationInitialized)
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
		it.Event = new(ReputationInitialized)
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
func (it *ReputationInitializedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ReputationInitializedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ReputationInitialized represents a Initialized event raised by the Reputation contract.
type ReputationInitialized struct {
	Version uint64
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterInitialized is a free log retrieval operation binding the contract event 0xc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d2.
//
// Solidity: event Initialized(uint64 version)
func (_Reputation *ReputationFilterer) FilterInitialized(opts *bind.FilterOpts) (*ReputationInitializedIterator, error) {

	logs, sub, err := _Reputation.contract.FilterLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return &ReputationInitializedIterator{contract: _Reputation.contract, event: "Initialized", logs: logs, sub: sub}, nil
}

// WatchInitialized is a free log subscription operation binding the contract event 0xc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d2.
//
// Solidity: event Initialized(uint64 version)
func (_Reputation *ReputationFilterer) WatchInitialized(opts *bind.WatchOpts, sink chan<- *ReputationInitialized) (event.Subscription, error) {

	logs, sub, err := _Reputation.contract.WatchLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ReputationInitialized)
				if err := _Reputation.contract.UnpackLog(event, "Initialized", log); err != nil {
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

// ParseInitialized is a log parse operation binding the contract event 0xc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d2.
//
// Solidity: event Initialized(uint64 version)
func (_Reputation *ReputationFilterer) ParseInitialized(log types.Log) (*ReputationInitialized, error) {
	event := new(ReputationInitialized)
	if err := _Reputation.contract.UnpackLog(event, "Initialized", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ReputationNewFeedbackIterator is returned from FilterNewFeedback and is used to iterate over the raw logs and unpacked data for NewFeedback events raised by the Reputation contract.
type ReputationNewFeedbackIterator struct {
	Event *ReputationNewFeedback // Event containing the contract specifics and raw log

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
func (it *ReputationNewFeedbackIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ReputationNewFeedback)
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
		it.Event = new(ReputationNewFeedback)
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
func (it *ReputationNewFeedbackIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ReputationNewFeedbackIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ReputationNewFeedback represents a NewFeedback event raised by the Reputation contract.
type ReputationNewFeedback struct {
	AgentId       *big.Int
	ClientAddress common.Address
	FeedbackIndex uint64
	Value         *big.Int
	ValueDecimals uint8
	IndexedTag1   common.Hash
	Tag1          string
	Tag2          string
	Endpoint      string
	FeedbackURI   string
	FeedbackHash  [32]byte
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterNewFeedback is a free log retrieval operation binding the contract event 0x6a4a61743519c9d648a14e6493f47dbe3ff1aa29e7785c96c8326a205e58febc.
//
// Solidity: event NewFeedback(uint256 indexed agentId, address indexed clientAddress, uint64 feedbackIndex, int128 value, uint8 valueDecimals, string indexed indexedTag1, string tag1, string tag2, string endpoint, string feedbackURI, bytes32 feedbackHash)
func (_Reputation *ReputationFilterer) FilterNewFeedback(opts *bind.FilterOpts, agentId []*big.Int, clientAddress []common.Address, indexedTag1 []string) (*ReputationNewFeedbackIterator, error) {

	var agentIdRule []interface{}
	for _, agentIdItem := range agentId {
		agentIdRule = append(agentIdRule, agentIdItem)
	}
	var clientAddressRule []interface{}
	for _, clientAddressItem := range clientAddress {
		clientAddressRule = append(clientAddressRule, clientAddressItem)
	}

	var indexedTag1Rule []interface{}
	for _, indexedTag1Item := range indexedTag1 {
		indexedTag1Rule = append(indexedTag1Rule, indexedTag1Item)
	}

	logs, sub, err := _Reputation.contract.FilterLogs(opts, "NewFeedback", agentIdRule, clientAddressRule, indexedTag1Rule)
	if err != nil {
		return nil, err
	}
	return &ReputationNewFeedbackIterator{contract: _Reputation.contract, event: "NewFeedback", logs: logs, sub: sub}, nil
}

// WatchNewFeedback is a free log subscription operation binding the contract event 0x6a4a61743519c9d648a14e6493f47dbe3ff1aa29e7785c96c8326a205e58febc.
//
// Solidity: event NewFeedback(uint256 indexed agentId, address indexed clientAddress, uint64 feedbackIndex, int128 value, uint8 valueDecimals, string indexed indexedTag1, string tag1, string tag2, string endpoint, string feedbackURI, bytes32 feedbackHash)
func (_Reputation *ReputationFilterer) WatchNewFeedback(opts *bind.WatchOpts, sink chan<- *ReputationNewFeedback, agentId []*big.Int, clientAddress []common.Address, indexedTag1 []string) (event.Subscription, error) {

	var agentIdRule []interface{}
	for _, agentIdItem := range agentId {
		agentIdRule = append(agentIdRule, agentIdItem)
	}
	var clientAddressRule []interface{}
	for _, clientAddressItem := range clientAddress {
		clientAddressRule = append(clientAddressRule, clientAddressItem)
	}

	var indexedTag1Rule []interface{}
	for _, indexedTag1Item := range indexedTag1 {
		indexedTag1Rule = append(indexedTag1Rule, indexedTag1Item)
	}

	logs, sub, err := _Reputation.contract.WatchLogs(opts, "NewFeedback", agentIdRule, clientAddressRule, indexedTag1Rule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ReputationNewFeedback)
				if err := _Reputation.contract.UnpackLog(event, "NewFeedback", log); err != nil {
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

// ParseNewFeedback is a log parse operation binding the contract event 0x6a4a61743519c9d648a14e6493f47dbe3ff1aa29e7785c96c8326a205e58febc.
//
// Solidity: event NewFeedback(uint256 indexed agentId, address indexed clientAddress, uint64 feedbackIndex, int128 value, uint8 valueDecimals, string indexed indexedTag1, string tag1, string tag2, string endpoint, string feedbackURI, bytes32 feedbackHash)
func (_Reputation *ReputationFilterer) ParseNewFeedback(log types.Log) (*ReputationNewFeedback, error) {
	event := new(ReputationNewFeedback)
	if err := _Reputation.contract.UnpackLog(event, "NewFeedback", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ReputationOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the Reputation contract.
type ReputationOwnershipTransferredIterator struct {
	Event *ReputationOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *ReputationOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ReputationOwnershipTransferred)
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
		it.Event = new(ReputationOwnershipTransferred)
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
func (it *ReputationOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ReputationOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ReputationOwnershipTransferred represents a OwnershipTransferred event raised by the Reputation contract.
type ReputationOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Reputation *ReputationFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*ReputationOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Reputation.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &ReputationOwnershipTransferredIterator{contract: _Reputation.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Reputation *ReputationFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *ReputationOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Reputation.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ReputationOwnershipTransferred)
				if err := _Reputation.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
func (_Reputation *ReputationFilterer) ParseOwnershipTransferred(log types.Log) (*ReputationOwnershipTransferred, error) {
	event := new(ReputationOwnershipTransferred)
	if err := _Reputation.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ReputationResponseAppendedIterator is returned from FilterResponseAppended and is used to iterate over the raw logs and unpacked data for ResponseAppended events raised by the Reputation contract.
type ReputationResponseAppendedIterator struct {
	Event *ReputationResponseAppended // Event containing the contract specifics and raw log

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
func (it *ReputationResponseAppendedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ReputationResponseAppended)
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
		it.Event = new(ReputationResponseAppended)
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
func (it *ReputationResponseAppendedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ReputationResponseAppendedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ReputationResponseAppended represents a ResponseAppended event raised by the Reputation contract.
type ReputationResponseAppended struct {
	AgentId       *big.Int
	ClientAddress common.Address
	FeedbackIndex uint64
	Responder     common.Address
	ResponseURI   string
	ResponseHash  [32]byte
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterResponseAppended is a free log retrieval operation binding the contract event 0xb1c6be0b5b8aef6539e2fac0fd131a2faa7b49edf8e505b5eb0ad487d56051d4.
//
// Solidity: event ResponseAppended(uint256 indexed agentId, address indexed clientAddress, uint64 feedbackIndex, address indexed responder, string responseURI, bytes32 responseHash)
func (_Reputation *ReputationFilterer) FilterResponseAppended(opts *bind.FilterOpts, agentId []*big.Int, clientAddress []common.Address, responder []common.Address) (*ReputationResponseAppendedIterator, error) {

	var agentIdRule []interface{}
	for _, agentIdItem := range agentId {
		agentIdRule = append(agentIdRule, agentIdItem)
	}
	var clientAddressRule []interface{}
	for _, clientAddressItem := range clientAddress {
		clientAddressRule = append(clientAddressRule, clientAddressItem)
	}

	var responderRule []interface{}
	for _, responderItem := range responder {
		responderRule = append(responderRule, responderItem)
	}

	logs, sub, err := _Reputation.contract.FilterLogs(opts, "ResponseAppended", agentIdRule, clientAddressRule, responderRule)
	if err != nil {
		return nil, err
	}
	return &ReputationResponseAppendedIterator{contract: _Reputation.contract, event: "ResponseAppended", logs: logs, sub: sub}, nil
}

// WatchResponseAppended is a free log subscription operation binding the contract event 0xb1c6be0b5b8aef6539e2fac0fd131a2faa7b49edf8e505b5eb0ad487d56051d4.
//
// Solidity: event ResponseAppended(uint256 indexed agentId, address indexed clientAddress, uint64 feedbackIndex, address indexed responder, string responseURI, bytes32 responseHash)
func (_Reputation *ReputationFilterer) WatchResponseAppended(opts *bind.WatchOpts, sink chan<- *ReputationResponseAppended, agentId []*big.Int, clientAddress []common.Address, responder []common.Address) (event.Subscription, error) {

	var agentIdRule []interface{}
	for _, agentIdItem := range agentId {
		agentIdRule = append(agentIdRule, agentIdItem)
	}
	var clientAddressRule []interface{}
	for _, clientAddressItem := range clientAddress {
		clientAddressRule = append(clientAddressRule, clientAddressItem)
	}

	var responderRule []interface{}
	for _, responderItem := range responder {
		responderRule = append(responderRule, responderItem)
	}

	logs, sub, err := _Reputation.contract.WatchLogs(opts, "ResponseAppended", agentIdRule, clientAddressRule, responderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ReputationResponseAppended)
				if err := _Reputation.contract.UnpackLog(event, "ResponseAppended", log); err != nil {
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

// ParseResponseAppended is a log parse operation binding the contract event 0xb1c6be0b5b8aef6539e2fac0fd131a2faa7b49edf8e505b5eb0ad487d56051d4.
//
// Solidity: event ResponseAppended(uint256 indexed agentId, address indexed clientAddress, uint64 feedbackIndex, address indexed responder, string responseURI, bytes32 responseHash)
func (_Reputation *ReputationFilterer) ParseResponseAppended(log types.Log) (*ReputationResponseAppended, error) {
	event := new(ReputationResponseAppended)
	if err := _Reputation.contract.UnpackLog(event, "ResponseAppended", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ReputationUpgradedIterator is returned from FilterUpgraded and is used to iterate over the raw logs and unpacked data for Upgraded events raised by the Reputation contract.
type ReputationUpgradedIterator struct {
	Event *ReputationUpgraded // Event containing the contract specifics and raw log

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
func (it *ReputationUpgradedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ReputationUpgraded)
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
		it.Event = new(ReputationUpgraded)
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
func (it *ReputationUpgradedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ReputationUpgradedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ReputationUpgraded represents a Upgraded event raised by the Reputation contract.
type ReputationUpgraded struct {
	Implementation common.Address
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterUpgraded is a free log retrieval operation binding the contract event 0xbc7cd75a20ee27fd9adebab32041f755214dbc6bffa90cc0225b39da2e5c2d3b.
//
// Solidity: event Upgraded(address indexed implementation)
func (_Reputation *ReputationFilterer) FilterUpgraded(opts *bind.FilterOpts, implementation []common.Address) (*ReputationUpgradedIterator, error) {

	var implementationRule []interface{}
	for _, implementationItem := range implementation {
		implementationRule = append(implementationRule, implementationItem)
	}

	logs, sub, err := _Reputation.contract.FilterLogs(opts, "Upgraded", implementationRule)
	if err != nil {
		return nil, err
	}
	return &ReputationUpgradedIterator{contract: _Reputation.contract, event: "Upgraded", logs: logs, sub: sub}, nil
}

// WatchUpgraded is a free log subscription operation binding the contract event 0xbc7cd75a20ee27fd9adebab32041f755214dbc6bffa90cc0225b39da2e5c2d3b.
//
// Solidity: event Upgraded(address indexed implementation)
func (_Reputation *ReputationFilterer) WatchUpgraded(opts *bind.WatchOpts, sink chan<- *ReputationUpgraded, implementation []common.Address) (event.Subscription, error) {

	var implementationRule []interface{}
	for _, implementationItem := range implementation {
		implementationRule = append(implementationRule, implementationItem)
	}

	logs, sub, err := _Reputation.contract.WatchLogs(opts, "Upgraded", implementationRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ReputationUpgraded)
				if err := _Reputation.contract.UnpackLog(event, "Upgraded", log); err != nil {
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

// ParseUpgraded is a log parse operation binding the contract event 0xbc7cd75a20ee27fd9adebab32041f755214dbc6bffa90cc0225b39da2e5c2d3b.
//
// Solidity: event Upgraded(address indexed implementation)
func (_Reputation *ReputationFilterer) ParseUpgraded(log types.Log) (*ReputationUpgraded, error) {
	event := new(ReputationUpgraded)
	if err := _Reputation.contract.UnpackLog(event, "Upgraded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
