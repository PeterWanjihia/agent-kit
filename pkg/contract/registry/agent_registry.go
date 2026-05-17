// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package registry

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

// IdentityRegistryUpgradeableMetadataEntry is an auto generated low-level Go binding around an user-defined struct.
type IdentityRegistryUpgradeableMetadataEntry struct {
	MetadataKey   string
	MetadataValue []byte
}

// RegistryMetaData contains all meta data concerning the Registry contract.
var RegistryMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"target\",\"type\":\"address\"}],\"name\":\"AddressEmptyCode\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ECDSAInvalidSignature\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"length\",\"type\":\"uint256\"}],\"name\":\"ECDSAInvalidSignatureLength\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"s\",\"type\":\"bytes32\"}],\"name\":\"ECDSAInvalidSignatureS\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"implementation\",\"type\":\"address\"}],\"name\":\"ERC1967InvalidImplementation\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ERC1967NonPayable\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"ERC721IncorrectOwner\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"ERC721InsufficientApproval\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"approver\",\"type\":\"address\"}],\"name\":\"ERC721InvalidApprover\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"}],\"name\":\"ERC721InvalidOperator\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"ERC721InvalidOwner\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"receiver\",\"type\":\"address\"}],\"name\":\"ERC721InvalidReceiver\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"ERC721InvalidSender\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"ERC721NonexistentToken\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"FailedCall\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidInitialization\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NotInitializing\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"OwnableInvalidOwner\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"OwnableUnauthorizedAccount\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"UUPSUnauthorizedCallContext\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"slot\",\"type\":\"bytes32\"}],\"name\":\"UUPSUnsupportedProxiableUUID\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"approved\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"Approval\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"approved\",\"type\":\"bool\"}],\"name\":\"ApprovalForAll\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_fromTokenId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_toTokenId\",\"type\":\"uint256\"}],\"name\":\"BatchMetadataUpdate\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[],\"name\":\"EIP712DomainChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"version\",\"type\":\"uint64\"}],\"name\":\"Initialized\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"agentId\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"string\",\"name\":\"indexedMetadataKey\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"metadataKey\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"metadataValue\",\"type\":\"bytes\"}],\"name\":\"MetadataSet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_tokenId\",\"type\":\"uint256\"}],\"name\":\"MetadataUpdate\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"agentId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"agentURI\",\"type\":\"string\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"Registered\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"Transfer\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"agentId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"newURI\",\"type\":\"string\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"updatedBy\",\"type\":\"address\"}],\"name\":\"URIUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"implementation\",\"type\":\"address\"}],\"name\":\"Upgraded\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"UPGRADE_INTERFACE_VERSION\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"approve\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"eip712Domain\",\"outputs\":[{\"internalType\":\"bytes1\",\"name\":\"fields\",\"type\":\"bytes1\"},{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"version\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"chainId\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"verifyingContract\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"salt\",\"type\":\"bytes32\"},{\"internalType\":\"uint256[]\",\"name\":\"extensions\",\"type\":\"uint256[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"agentId\",\"type\":\"uint256\"}],\"name\":\"getAgentWallet\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"getApproved\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"agentId\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"metadataKey\",\"type\":\"string\"}],\"name\":\"getMetadata\",\"outputs\":[{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getVersion\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"}],\"name\":\"isApprovedForAll\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"name\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"ownerOf\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"proxiableUUID\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"register\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"agentId\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"agentURI\",\"type\":\"string\"},{\"components\":[{\"internalType\":\"string\",\"name\":\"metadataKey\",\"type\":\"string\"},{\"internalType\":\"bytes\",\"name\":\"metadataValue\",\"type\":\"bytes\"}],\"internalType\":\"structIdentityRegistryUpgradeable.MetadataEntry[]\",\"name\":\"metadata\",\"type\":\"tuple[]\"}],\"name\":\"register\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"agentId\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"agentURI\",\"type\":\"string\"}],\"name\":\"register\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"agentId\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"safeTransferFrom\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"safeTransferFrom\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"agentId\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"newURI\",\"type\":\"string\"}],\"name\":\"setAgentURI\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"agentId\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"newWallet\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"signature\",\"type\":\"bytes\"}],\"name\":\"setAgentWallet\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"approved\",\"type\":\"bool\"}],\"name\":\"setApprovalForAll\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"agentId\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"metadataKey\",\"type\":\"string\"},{\"internalType\":\"bytes\",\"name\":\"metadataValue\",\"type\":\"bytes\"}],\"name\":\"setMetadata\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes4\",\"name\":\"interfaceId\",\"type\":\"bytes4\"}],\"name\":\"supportsInterface\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"symbol\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"tokenURI\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"transferFrom\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"agentId\",\"type\":\"uint256\"}],\"name\":\"unsetAgentWallet\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newImplementation\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"upgradeToAndCall\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"}]",
}

// RegistryABI is the input ABI used to generate the binding from.
// Deprecated: Use RegistryMetaData.ABI instead.
var RegistryABI = RegistryMetaData.ABI

// Registry is an auto generated Go binding around an Ethereum contract.
type Registry struct {
	RegistryCaller     // Read-only binding to the contract
	RegistryTransactor // Write-only binding to the contract
	RegistryFilterer   // Log filterer for contract events
}

// RegistryCaller is an auto generated read-only Go binding around an Ethereum contract.
type RegistryCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// RegistryTransactor is an auto generated write-only Go binding around an Ethereum contract.
type RegistryTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// RegistryFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type RegistryFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// RegistrySession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type RegistrySession struct {
	Contract     *Registry         // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// RegistryCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type RegistryCallerSession struct {
	Contract *RegistryCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts   // Call options to use throughout this session
}

// RegistryTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type RegistryTransactorSession struct {
	Contract     *RegistryTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts   // Transaction auth options to use throughout this session
}

// RegistryRaw is an auto generated low-level Go binding around an Ethereum contract.
type RegistryRaw struct {
	Contract *Registry // Generic contract binding to access the raw methods on
}

// RegistryCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type RegistryCallerRaw struct {
	Contract *RegistryCaller // Generic read-only contract binding to access the raw methods on
}

// RegistryTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type RegistryTransactorRaw struct {
	Contract *RegistryTransactor // Generic write-only contract binding to access the raw methods on
}

// NewRegistry creates a new instance of Registry, bound to a specific deployed contract.
func NewRegistry(address common.Address, backend bind.ContractBackend) (*Registry, error) {
	contract, err := bindRegistry(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Registry{RegistryCaller: RegistryCaller{contract: contract}, RegistryTransactor: RegistryTransactor{contract: contract}, RegistryFilterer: RegistryFilterer{contract: contract}}, nil
}

// NewRegistryCaller creates a new read-only instance of Registry, bound to a specific deployed contract.
func NewRegistryCaller(address common.Address, caller bind.ContractCaller) (*RegistryCaller, error) {
	contract, err := bindRegistry(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &RegistryCaller{contract: contract}, nil
}

// NewRegistryTransactor creates a new write-only instance of Registry, bound to a specific deployed contract.
func NewRegistryTransactor(address common.Address, transactor bind.ContractTransactor) (*RegistryTransactor, error) {
	contract, err := bindRegistry(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &RegistryTransactor{contract: contract}, nil
}

// NewRegistryFilterer creates a new log filterer instance of Registry, bound to a specific deployed contract.
func NewRegistryFilterer(address common.Address, filterer bind.ContractFilterer) (*RegistryFilterer, error) {
	contract, err := bindRegistry(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &RegistryFilterer{contract: contract}, nil
}

// bindRegistry binds a generic wrapper to an already deployed contract.
func bindRegistry(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := RegistryMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Registry *RegistryRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Registry.Contract.RegistryCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Registry *RegistryRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Registry.Contract.RegistryTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Registry *RegistryRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Registry.Contract.RegistryTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Registry *RegistryCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Registry.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Registry *RegistryTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Registry.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Registry *RegistryTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Registry.Contract.contract.Transact(opts, method, params...)
}

// UPGRADEINTERFACEVERSION is a free data retrieval call binding the contract method 0xad3cb1cc.
//
// Solidity: function UPGRADE_INTERFACE_VERSION() view returns(string)
func (_Registry *RegistryCaller) UPGRADEINTERFACEVERSION(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _Registry.contract.Call(opts, &out, "UPGRADE_INTERFACE_VERSION")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// UPGRADEINTERFACEVERSION is a free data retrieval call binding the contract method 0xad3cb1cc.
//
// Solidity: function UPGRADE_INTERFACE_VERSION() view returns(string)
func (_Registry *RegistrySession) UPGRADEINTERFACEVERSION() (string, error) {
	return _Registry.Contract.UPGRADEINTERFACEVERSION(&_Registry.CallOpts)
}

// UPGRADEINTERFACEVERSION is a free data retrieval call binding the contract method 0xad3cb1cc.
//
// Solidity: function UPGRADE_INTERFACE_VERSION() view returns(string)
func (_Registry *RegistryCallerSession) UPGRADEINTERFACEVERSION() (string, error) {
	return _Registry.Contract.UPGRADEINTERFACEVERSION(&_Registry.CallOpts)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address owner) view returns(uint256)
func (_Registry *RegistryCaller) BalanceOf(opts *bind.CallOpts, owner common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Registry.contract.Call(opts, &out, "balanceOf", owner)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address owner) view returns(uint256)
func (_Registry *RegistrySession) BalanceOf(owner common.Address) (*big.Int, error) {
	return _Registry.Contract.BalanceOf(&_Registry.CallOpts, owner)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address owner) view returns(uint256)
func (_Registry *RegistryCallerSession) BalanceOf(owner common.Address) (*big.Int, error) {
	return _Registry.Contract.BalanceOf(&_Registry.CallOpts, owner)
}

// Eip712Domain is a free data retrieval call binding the contract method 0x84b0196e.
//
// Solidity: function eip712Domain() view returns(bytes1 fields, string name, string version, uint256 chainId, address verifyingContract, bytes32 salt, uint256[] extensions)
func (_Registry *RegistryCaller) Eip712Domain(opts *bind.CallOpts) (struct {
	Fields            [1]byte
	Name              string
	Version           string
	ChainId           *big.Int
	VerifyingContract common.Address
	Salt              [32]byte
	Extensions        []*big.Int
}, error) {
	var out []interface{}
	err := _Registry.contract.Call(opts, &out, "eip712Domain")

	outstruct := new(struct {
		Fields            [1]byte
		Name              string
		Version           string
		ChainId           *big.Int
		VerifyingContract common.Address
		Salt              [32]byte
		Extensions        []*big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Fields = *abi.ConvertType(out[0], new([1]byte)).(*[1]byte)
	outstruct.Name = *abi.ConvertType(out[1], new(string)).(*string)
	outstruct.Version = *abi.ConvertType(out[2], new(string)).(*string)
	outstruct.ChainId = *abi.ConvertType(out[3], new(*big.Int)).(**big.Int)
	outstruct.VerifyingContract = *abi.ConvertType(out[4], new(common.Address)).(*common.Address)
	outstruct.Salt = *abi.ConvertType(out[5], new([32]byte)).(*[32]byte)
	outstruct.Extensions = *abi.ConvertType(out[6], new([]*big.Int)).(*[]*big.Int)

	return *outstruct, err

}

// Eip712Domain is a free data retrieval call binding the contract method 0x84b0196e.
//
// Solidity: function eip712Domain() view returns(bytes1 fields, string name, string version, uint256 chainId, address verifyingContract, bytes32 salt, uint256[] extensions)
func (_Registry *RegistrySession) Eip712Domain() (struct {
	Fields            [1]byte
	Name              string
	Version           string
	ChainId           *big.Int
	VerifyingContract common.Address
	Salt              [32]byte
	Extensions        []*big.Int
}, error) {
	return _Registry.Contract.Eip712Domain(&_Registry.CallOpts)
}

// Eip712Domain is a free data retrieval call binding the contract method 0x84b0196e.
//
// Solidity: function eip712Domain() view returns(bytes1 fields, string name, string version, uint256 chainId, address verifyingContract, bytes32 salt, uint256[] extensions)
func (_Registry *RegistryCallerSession) Eip712Domain() (struct {
	Fields            [1]byte
	Name              string
	Version           string
	ChainId           *big.Int
	VerifyingContract common.Address
	Salt              [32]byte
	Extensions        []*big.Int
}, error) {
	return _Registry.Contract.Eip712Domain(&_Registry.CallOpts)
}

// GetAgentWallet is a free data retrieval call binding the contract method 0x00339509.
//
// Solidity: function getAgentWallet(uint256 agentId) view returns(address)
func (_Registry *RegistryCaller) GetAgentWallet(opts *bind.CallOpts, agentId *big.Int) (common.Address, error) {
	var out []interface{}
	err := _Registry.contract.Call(opts, &out, "getAgentWallet", agentId)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetAgentWallet is a free data retrieval call binding the contract method 0x00339509.
//
// Solidity: function getAgentWallet(uint256 agentId) view returns(address)
func (_Registry *RegistrySession) GetAgentWallet(agentId *big.Int) (common.Address, error) {
	return _Registry.Contract.GetAgentWallet(&_Registry.CallOpts, agentId)
}

// GetAgentWallet is a free data retrieval call binding the contract method 0x00339509.
//
// Solidity: function getAgentWallet(uint256 agentId) view returns(address)
func (_Registry *RegistryCallerSession) GetAgentWallet(agentId *big.Int) (common.Address, error) {
	return _Registry.Contract.GetAgentWallet(&_Registry.CallOpts, agentId)
}

// GetApproved is a free data retrieval call binding the contract method 0x081812fc.
//
// Solidity: function getApproved(uint256 tokenId) view returns(address)
func (_Registry *RegistryCaller) GetApproved(opts *bind.CallOpts, tokenId *big.Int) (common.Address, error) {
	var out []interface{}
	err := _Registry.contract.Call(opts, &out, "getApproved", tokenId)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetApproved is a free data retrieval call binding the contract method 0x081812fc.
//
// Solidity: function getApproved(uint256 tokenId) view returns(address)
func (_Registry *RegistrySession) GetApproved(tokenId *big.Int) (common.Address, error) {
	return _Registry.Contract.GetApproved(&_Registry.CallOpts, tokenId)
}

// GetApproved is a free data retrieval call binding the contract method 0x081812fc.
//
// Solidity: function getApproved(uint256 tokenId) view returns(address)
func (_Registry *RegistryCallerSession) GetApproved(tokenId *big.Int) (common.Address, error) {
	return _Registry.Contract.GetApproved(&_Registry.CallOpts, tokenId)
}

// GetMetadata is a free data retrieval call binding the contract method 0xcb4799f2.
//
// Solidity: function getMetadata(uint256 agentId, string metadataKey) view returns(bytes)
func (_Registry *RegistryCaller) GetMetadata(opts *bind.CallOpts, agentId *big.Int, metadataKey string) ([]byte, error) {
	var out []interface{}
	err := _Registry.contract.Call(opts, &out, "getMetadata", agentId, metadataKey)

	if err != nil {
		return *new([]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([]byte)).(*[]byte)

	return out0, err

}

// GetMetadata is a free data retrieval call binding the contract method 0xcb4799f2.
//
// Solidity: function getMetadata(uint256 agentId, string metadataKey) view returns(bytes)
func (_Registry *RegistrySession) GetMetadata(agentId *big.Int, metadataKey string) ([]byte, error) {
	return _Registry.Contract.GetMetadata(&_Registry.CallOpts, agentId, metadataKey)
}

// GetMetadata is a free data retrieval call binding the contract method 0xcb4799f2.
//
// Solidity: function getMetadata(uint256 agentId, string metadataKey) view returns(bytes)
func (_Registry *RegistryCallerSession) GetMetadata(agentId *big.Int, metadataKey string) ([]byte, error) {
	return _Registry.Contract.GetMetadata(&_Registry.CallOpts, agentId, metadataKey)
}

// GetVersion is a free data retrieval call binding the contract method 0x0d8e6e2c.
//
// Solidity: function getVersion() pure returns(string)
func (_Registry *RegistryCaller) GetVersion(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _Registry.contract.Call(opts, &out, "getVersion")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// GetVersion is a free data retrieval call binding the contract method 0x0d8e6e2c.
//
// Solidity: function getVersion() pure returns(string)
func (_Registry *RegistrySession) GetVersion() (string, error) {
	return _Registry.Contract.GetVersion(&_Registry.CallOpts)
}

// GetVersion is a free data retrieval call binding the contract method 0x0d8e6e2c.
//
// Solidity: function getVersion() pure returns(string)
func (_Registry *RegistryCallerSession) GetVersion() (string, error) {
	return _Registry.Contract.GetVersion(&_Registry.CallOpts)
}

// IsApprovedForAll is a free data retrieval call binding the contract method 0xe985e9c5.
//
// Solidity: function isApprovedForAll(address owner, address operator) view returns(bool)
func (_Registry *RegistryCaller) IsApprovedForAll(opts *bind.CallOpts, owner common.Address, operator common.Address) (bool, error) {
	var out []interface{}
	err := _Registry.contract.Call(opts, &out, "isApprovedForAll", owner, operator)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsApprovedForAll is a free data retrieval call binding the contract method 0xe985e9c5.
//
// Solidity: function isApprovedForAll(address owner, address operator) view returns(bool)
func (_Registry *RegistrySession) IsApprovedForAll(owner common.Address, operator common.Address) (bool, error) {
	return _Registry.Contract.IsApprovedForAll(&_Registry.CallOpts, owner, operator)
}

// IsApprovedForAll is a free data retrieval call binding the contract method 0xe985e9c5.
//
// Solidity: function isApprovedForAll(address owner, address operator) view returns(bool)
func (_Registry *RegistryCallerSession) IsApprovedForAll(owner common.Address, operator common.Address) (bool, error) {
	return _Registry.Contract.IsApprovedForAll(&_Registry.CallOpts, owner, operator)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_Registry *RegistryCaller) Name(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _Registry.contract.Call(opts, &out, "name")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_Registry *RegistrySession) Name() (string, error) {
	return _Registry.Contract.Name(&_Registry.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_Registry *RegistryCallerSession) Name() (string, error) {
	return _Registry.Contract.Name(&_Registry.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Registry *RegistryCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Registry.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Registry *RegistrySession) Owner() (common.Address, error) {
	return _Registry.Contract.Owner(&_Registry.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Registry *RegistryCallerSession) Owner() (common.Address, error) {
	return _Registry.Contract.Owner(&_Registry.CallOpts)
}

// OwnerOf is a free data retrieval call binding the contract method 0x6352211e.
//
// Solidity: function ownerOf(uint256 tokenId) view returns(address)
func (_Registry *RegistryCaller) OwnerOf(opts *bind.CallOpts, tokenId *big.Int) (common.Address, error) {
	var out []interface{}
	err := _Registry.contract.Call(opts, &out, "ownerOf", tokenId)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// OwnerOf is a free data retrieval call binding the contract method 0x6352211e.
//
// Solidity: function ownerOf(uint256 tokenId) view returns(address)
func (_Registry *RegistrySession) OwnerOf(tokenId *big.Int) (common.Address, error) {
	return _Registry.Contract.OwnerOf(&_Registry.CallOpts, tokenId)
}

// OwnerOf is a free data retrieval call binding the contract method 0x6352211e.
//
// Solidity: function ownerOf(uint256 tokenId) view returns(address)
func (_Registry *RegistryCallerSession) OwnerOf(tokenId *big.Int) (common.Address, error) {
	return _Registry.Contract.OwnerOf(&_Registry.CallOpts, tokenId)
}

// ProxiableUUID is a free data retrieval call binding the contract method 0x52d1902d.
//
// Solidity: function proxiableUUID() view returns(bytes32)
func (_Registry *RegistryCaller) ProxiableUUID(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _Registry.contract.Call(opts, &out, "proxiableUUID")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// ProxiableUUID is a free data retrieval call binding the contract method 0x52d1902d.
//
// Solidity: function proxiableUUID() view returns(bytes32)
func (_Registry *RegistrySession) ProxiableUUID() ([32]byte, error) {
	return _Registry.Contract.ProxiableUUID(&_Registry.CallOpts)
}

// ProxiableUUID is a free data retrieval call binding the contract method 0x52d1902d.
//
// Solidity: function proxiableUUID() view returns(bytes32)
func (_Registry *RegistryCallerSession) ProxiableUUID() ([32]byte, error) {
	return _Registry.Contract.ProxiableUUID(&_Registry.CallOpts)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_Registry *RegistryCaller) SupportsInterface(opts *bind.CallOpts, interfaceId [4]byte) (bool, error) {
	var out []interface{}
	err := _Registry.contract.Call(opts, &out, "supportsInterface", interfaceId)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_Registry *RegistrySession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _Registry.Contract.SupportsInterface(&_Registry.CallOpts, interfaceId)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_Registry *RegistryCallerSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _Registry.Contract.SupportsInterface(&_Registry.CallOpts, interfaceId)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_Registry *RegistryCaller) Symbol(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _Registry.contract.Call(opts, &out, "symbol")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_Registry *RegistrySession) Symbol() (string, error) {
	return _Registry.Contract.Symbol(&_Registry.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_Registry *RegistryCallerSession) Symbol() (string, error) {
	return _Registry.Contract.Symbol(&_Registry.CallOpts)
}

// TokenURI is a free data retrieval call binding the contract method 0xc87b56dd.
//
// Solidity: function tokenURI(uint256 tokenId) view returns(string)
func (_Registry *RegistryCaller) TokenURI(opts *bind.CallOpts, tokenId *big.Int) (string, error) {
	var out []interface{}
	err := _Registry.contract.Call(opts, &out, "tokenURI", tokenId)

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// TokenURI is a free data retrieval call binding the contract method 0xc87b56dd.
//
// Solidity: function tokenURI(uint256 tokenId) view returns(string)
func (_Registry *RegistrySession) TokenURI(tokenId *big.Int) (string, error) {
	return _Registry.Contract.TokenURI(&_Registry.CallOpts, tokenId)
}

// TokenURI is a free data retrieval call binding the contract method 0xc87b56dd.
//
// Solidity: function tokenURI(uint256 tokenId) view returns(string)
func (_Registry *RegistryCallerSession) TokenURI(tokenId *big.Int) (string, error) {
	return _Registry.Contract.TokenURI(&_Registry.CallOpts, tokenId)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address to, uint256 tokenId) returns()
func (_Registry *RegistryTransactor) Approve(opts *bind.TransactOpts, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _Registry.contract.Transact(opts, "approve", to, tokenId)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address to, uint256 tokenId) returns()
func (_Registry *RegistrySession) Approve(to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _Registry.Contract.Approve(&_Registry.TransactOpts, to, tokenId)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address to, uint256 tokenId) returns()
func (_Registry *RegistryTransactorSession) Approve(to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _Registry.Contract.Approve(&_Registry.TransactOpts, to, tokenId)
}

// Initialize is a paid mutator transaction binding the contract method 0x8129fc1c.
//
// Solidity: function initialize() returns()
func (_Registry *RegistryTransactor) Initialize(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Registry.contract.Transact(opts, "initialize")
}

// Initialize is a paid mutator transaction binding the contract method 0x8129fc1c.
//
// Solidity: function initialize() returns()
func (_Registry *RegistrySession) Initialize() (*types.Transaction, error) {
	return _Registry.Contract.Initialize(&_Registry.TransactOpts)
}

// Initialize is a paid mutator transaction binding the contract method 0x8129fc1c.
//
// Solidity: function initialize() returns()
func (_Registry *RegistryTransactorSession) Initialize() (*types.Transaction, error) {
	return _Registry.Contract.Initialize(&_Registry.TransactOpts)
}

// Register is a paid mutator transaction binding the contract method 0x1aa3a008.
//
// Solidity: function register() returns(uint256 agentId)
func (_Registry *RegistryTransactor) Register(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Registry.contract.Transact(opts, "register")
}

// Register is a paid mutator transaction binding the contract method 0x1aa3a008.
//
// Solidity: function register() returns(uint256 agentId)
func (_Registry *RegistrySession) Register() (*types.Transaction, error) {
	return _Registry.Contract.Register(&_Registry.TransactOpts)
}

// Register is a paid mutator transaction binding the contract method 0x1aa3a008.
//
// Solidity: function register() returns(uint256 agentId)
func (_Registry *RegistryTransactorSession) Register() (*types.Transaction, error) {
	return _Registry.Contract.Register(&_Registry.TransactOpts)
}

// Register0 is a paid mutator transaction binding the contract method 0x8ea42286.
//
// Solidity: function register(string agentURI, (string,bytes)[] metadata) returns(uint256 agentId)
func (_Registry *RegistryTransactor) Register0(opts *bind.TransactOpts, agentURI string, metadata []IdentityRegistryUpgradeableMetadataEntry) (*types.Transaction, error) {
	return _Registry.contract.Transact(opts, "register0", agentURI, metadata)
}

// Register0 is a paid mutator transaction binding the contract method 0x8ea42286.
//
// Solidity: function register(string agentURI, (string,bytes)[] metadata) returns(uint256 agentId)
func (_Registry *RegistrySession) Register0(agentURI string, metadata []IdentityRegistryUpgradeableMetadataEntry) (*types.Transaction, error) {
	return _Registry.Contract.Register0(&_Registry.TransactOpts, agentURI, metadata)
}

// Register0 is a paid mutator transaction binding the contract method 0x8ea42286.
//
// Solidity: function register(string agentURI, (string,bytes)[] metadata) returns(uint256 agentId)
func (_Registry *RegistryTransactorSession) Register0(agentURI string, metadata []IdentityRegistryUpgradeableMetadataEntry) (*types.Transaction, error) {
	return _Registry.Contract.Register0(&_Registry.TransactOpts, agentURI, metadata)
}

// Register1 is a paid mutator transaction binding the contract method 0xf2c298be.
//
// Solidity: function register(string agentURI) returns(uint256 agentId)
func (_Registry *RegistryTransactor) Register1(opts *bind.TransactOpts, agentURI string) (*types.Transaction, error) {
	return _Registry.contract.Transact(opts, "register1", agentURI)
}

// Register1 is a paid mutator transaction binding the contract method 0xf2c298be.
//
// Solidity: function register(string agentURI) returns(uint256 agentId)
func (_Registry *RegistrySession) Register1(agentURI string) (*types.Transaction, error) {
	return _Registry.Contract.Register1(&_Registry.TransactOpts, agentURI)
}

// Register1 is a paid mutator transaction binding the contract method 0xf2c298be.
//
// Solidity: function register(string agentURI) returns(uint256 agentId)
func (_Registry *RegistryTransactorSession) Register1(agentURI string) (*types.Transaction, error) {
	return _Registry.Contract.Register1(&_Registry.TransactOpts, agentURI)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Registry *RegistryTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Registry.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Registry *RegistrySession) RenounceOwnership() (*types.Transaction, error) {
	return _Registry.Contract.RenounceOwnership(&_Registry.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Registry *RegistryTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _Registry.Contract.RenounceOwnership(&_Registry.TransactOpts)
}

// SafeTransferFrom is a paid mutator transaction binding the contract method 0x42842e0e.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId) returns()
func (_Registry *RegistryTransactor) SafeTransferFrom(opts *bind.TransactOpts, from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _Registry.contract.Transact(opts, "safeTransferFrom", from, to, tokenId)
}

// SafeTransferFrom is a paid mutator transaction binding the contract method 0x42842e0e.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId) returns()
func (_Registry *RegistrySession) SafeTransferFrom(from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _Registry.Contract.SafeTransferFrom(&_Registry.TransactOpts, from, to, tokenId)
}

// SafeTransferFrom is a paid mutator transaction binding the contract method 0x42842e0e.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId) returns()
func (_Registry *RegistryTransactorSession) SafeTransferFrom(from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _Registry.Contract.SafeTransferFrom(&_Registry.TransactOpts, from, to, tokenId)
}

// SafeTransferFrom0 is a paid mutator transaction binding the contract method 0xb88d4fde.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId, bytes data) returns()
func (_Registry *RegistryTransactor) SafeTransferFrom0(opts *bind.TransactOpts, from common.Address, to common.Address, tokenId *big.Int, data []byte) (*types.Transaction, error) {
	return _Registry.contract.Transact(opts, "safeTransferFrom0", from, to, tokenId, data)
}

// SafeTransferFrom0 is a paid mutator transaction binding the contract method 0xb88d4fde.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId, bytes data) returns()
func (_Registry *RegistrySession) SafeTransferFrom0(from common.Address, to common.Address, tokenId *big.Int, data []byte) (*types.Transaction, error) {
	return _Registry.Contract.SafeTransferFrom0(&_Registry.TransactOpts, from, to, tokenId, data)
}

// SafeTransferFrom0 is a paid mutator transaction binding the contract method 0xb88d4fde.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId, bytes data) returns()
func (_Registry *RegistryTransactorSession) SafeTransferFrom0(from common.Address, to common.Address, tokenId *big.Int, data []byte) (*types.Transaction, error) {
	return _Registry.Contract.SafeTransferFrom0(&_Registry.TransactOpts, from, to, tokenId, data)
}

// SetAgentURI is a paid mutator transaction binding the contract method 0x0af28bd3.
//
// Solidity: function setAgentURI(uint256 agentId, string newURI) returns()
func (_Registry *RegistryTransactor) SetAgentURI(opts *bind.TransactOpts, agentId *big.Int, newURI string) (*types.Transaction, error) {
	return _Registry.contract.Transact(opts, "setAgentURI", agentId, newURI)
}

// SetAgentURI is a paid mutator transaction binding the contract method 0x0af28bd3.
//
// Solidity: function setAgentURI(uint256 agentId, string newURI) returns()
func (_Registry *RegistrySession) SetAgentURI(agentId *big.Int, newURI string) (*types.Transaction, error) {
	return _Registry.Contract.SetAgentURI(&_Registry.TransactOpts, agentId, newURI)
}

// SetAgentURI is a paid mutator transaction binding the contract method 0x0af28bd3.
//
// Solidity: function setAgentURI(uint256 agentId, string newURI) returns()
func (_Registry *RegistryTransactorSession) SetAgentURI(agentId *big.Int, newURI string) (*types.Transaction, error) {
	return _Registry.Contract.SetAgentURI(&_Registry.TransactOpts, agentId, newURI)
}

// SetAgentWallet is a paid mutator transaction binding the contract method 0x2d1ef5ae.
//
// Solidity: function setAgentWallet(uint256 agentId, address newWallet, uint256 deadline, bytes signature) returns()
func (_Registry *RegistryTransactor) SetAgentWallet(opts *bind.TransactOpts, agentId *big.Int, newWallet common.Address, deadline *big.Int, signature []byte) (*types.Transaction, error) {
	return _Registry.contract.Transact(opts, "setAgentWallet", agentId, newWallet, deadline, signature)
}

// SetAgentWallet is a paid mutator transaction binding the contract method 0x2d1ef5ae.
//
// Solidity: function setAgentWallet(uint256 agentId, address newWallet, uint256 deadline, bytes signature) returns()
func (_Registry *RegistrySession) SetAgentWallet(agentId *big.Int, newWallet common.Address, deadline *big.Int, signature []byte) (*types.Transaction, error) {
	return _Registry.Contract.SetAgentWallet(&_Registry.TransactOpts, agentId, newWallet, deadline, signature)
}

// SetAgentWallet is a paid mutator transaction binding the contract method 0x2d1ef5ae.
//
// Solidity: function setAgentWallet(uint256 agentId, address newWallet, uint256 deadline, bytes signature) returns()
func (_Registry *RegistryTransactorSession) SetAgentWallet(agentId *big.Int, newWallet common.Address, deadline *big.Int, signature []byte) (*types.Transaction, error) {
	return _Registry.Contract.SetAgentWallet(&_Registry.TransactOpts, agentId, newWallet, deadline, signature)
}

// SetApprovalForAll is a paid mutator transaction binding the contract method 0xa22cb465.
//
// Solidity: function setApprovalForAll(address operator, bool approved) returns()
func (_Registry *RegistryTransactor) SetApprovalForAll(opts *bind.TransactOpts, operator common.Address, approved bool) (*types.Transaction, error) {
	return _Registry.contract.Transact(opts, "setApprovalForAll", operator, approved)
}

// SetApprovalForAll is a paid mutator transaction binding the contract method 0xa22cb465.
//
// Solidity: function setApprovalForAll(address operator, bool approved) returns()
func (_Registry *RegistrySession) SetApprovalForAll(operator common.Address, approved bool) (*types.Transaction, error) {
	return _Registry.Contract.SetApprovalForAll(&_Registry.TransactOpts, operator, approved)
}

// SetApprovalForAll is a paid mutator transaction binding the contract method 0xa22cb465.
//
// Solidity: function setApprovalForAll(address operator, bool approved) returns()
func (_Registry *RegistryTransactorSession) SetApprovalForAll(operator common.Address, approved bool) (*types.Transaction, error) {
	return _Registry.Contract.SetApprovalForAll(&_Registry.TransactOpts, operator, approved)
}

// SetMetadata is a paid mutator transaction binding the contract method 0x466648da.
//
// Solidity: function setMetadata(uint256 agentId, string metadataKey, bytes metadataValue) returns()
func (_Registry *RegistryTransactor) SetMetadata(opts *bind.TransactOpts, agentId *big.Int, metadataKey string, metadataValue []byte) (*types.Transaction, error) {
	return _Registry.contract.Transact(opts, "setMetadata", agentId, metadataKey, metadataValue)
}

// SetMetadata is a paid mutator transaction binding the contract method 0x466648da.
//
// Solidity: function setMetadata(uint256 agentId, string metadataKey, bytes metadataValue) returns()
func (_Registry *RegistrySession) SetMetadata(agentId *big.Int, metadataKey string, metadataValue []byte) (*types.Transaction, error) {
	return _Registry.Contract.SetMetadata(&_Registry.TransactOpts, agentId, metadataKey, metadataValue)
}

// SetMetadata is a paid mutator transaction binding the contract method 0x466648da.
//
// Solidity: function setMetadata(uint256 agentId, string metadataKey, bytes metadataValue) returns()
func (_Registry *RegistryTransactorSession) SetMetadata(agentId *big.Int, metadataKey string, metadataValue []byte) (*types.Transaction, error) {
	return _Registry.Contract.SetMetadata(&_Registry.TransactOpts, agentId, metadataKey, metadataValue)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 tokenId) returns()
func (_Registry *RegistryTransactor) TransferFrom(opts *bind.TransactOpts, from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _Registry.contract.Transact(opts, "transferFrom", from, to, tokenId)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 tokenId) returns()
func (_Registry *RegistrySession) TransferFrom(from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _Registry.Contract.TransferFrom(&_Registry.TransactOpts, from, to, tokenId)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 tokenId) returns()
func (_Registry *RegistryTransactorSession) TransferFrom(from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _Registry.Contract.TransferFrom(&_Registry.TransactOpts, from, to, tokenId)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Registry *RegistryTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _Registry.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Registry *RegistrySession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Registry.Contract.TransferOwnership(&_Registry.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Registry *RegistryTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Registry.Contract.TransferOwnership(&_Registry.TransactOpts, newOwner)
}

// UnsetAgentWallet is a paid mutator transaction binding the contract method 0x3fddcf19.
//
// Solidity: function unsetAgentWallet(uint256 agentId) returns()
func (_Registry *RegistryTransactor) UnsetAgentWallet(opts *bind.TransactOpts, agentId *big.Int) (*types.Transaction, error) {
	return _Registry.contract.Transact(opts, "unsetAgentWallet", agentId)
}

// UnsetAgentWallet is a paid mutator transaction binding the contract method 0x3fddcf19.
//
// Solidity: function unsetAgentWallet(uint256 agentId) returns()
func (_Registry *RegistrySession) UnsetAgentWallet(agentId *big.Int) (*types.Transaction, error) {
	return _Registry.Contract.UnsetAgentWallet(&_Registry.TransactOpts, agentId)
}

// UnsetAgentWallet is a paid mutator transaction binding the contract method 0x3fddcf19.
//
// Solidity: function unsetAgentWallet(uint256 agentId) returns()
func (_Registry *RegistryTransactorSession) UnsetAgentWallet(agentId *big.Int) (*types.Transaction, error) {
	return _Registry.Contract.UnsetAgentWallet(&_Registry.TransactOpts, agentId)
}

// UpgradeToAndCall is a paid mutator transaction binding the contract method 0x4f1ef286.
//
// Solidity: function upgradeToAndCall(address newImplementation, bytes data) payable returns()
func (_Registry *RegistryTransactor) UpgradeToAndCall(opts *bind.TransactOpts, newImplementation common.Address, data []byte) (*types.Transaction, error) {
	return _Registry.contract.Transact(opts, "upgradeToAndCall", newImplementation, data)
}

// UpgradeToAndCall is a paid mutator transaction binding the contract method 0x4f1ef286.
//
// Solidity: function upgradeToAndCall(address newImplementation, bytes data) payable returns()
func (_Registry *RegistrySession) UpgradeToAndCall(newImplementation common.Address, data []byte) (*types.Transaction, error) {
	return _Registry.Contract.UpgradeToAndCall(&_Registry.TransactOpts, newImplementation, data)
}

// UpgradeToAndCall is a paid mutator transaction binding the contract method 0x4f1ef286.
//
// Solidity: function upgradeToAndCall(address newImplementation, bytes data) payable returns()
func (_Registry *RegistryTransactorSession) UpgradeToAndCall(newImplementation common.Address, data []byte) (*types.Transaction, error) {
	return _Registry.Contract.UpgradeToAndCall(&_Registry.TransactOpts, newImplementation, data)
}

// RegistryApprovalIterator is returned from FilterApproval and is used to iterate over the raw logs and unpacked data for Approval events raised by the Registry contract.
type RegistryApprovalIterator struct {
	Event *RegistryApproval // Event containing the contract specifics and raw log

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
func (it *RegistryApprovalIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RegistryApproval)
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
		it.Event = new(RegistryApproval)
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
func (it *RegistryApprovalIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RegistryApprovalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RegistryApproval represents a Approval event raised by the Registry contract.
type RegistryApproval struct {
	Owner    common.Address
	Approved common.Address
	TokenId  *big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterApproval is a free log retrieval operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed approved, uint256 indexed tokenId)
func (_Registry *RegistryFilterer) FilterApproval(opts *bind.FilterOpts, owner []common.Address, approved []common.Address, tokenId []*big.Int) (*RegistryApprovalIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var approvedRule []interface{}
	for _, approvedItem := range approved {
		approvedRule = append(approvedRule, approvedItem)
	}
	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}

	logs, sub, err := _Registry.contract.FilterLogs(opts, "Approval", ownerRule, approvedRule, tokenIdRule)
	if err != nil {
		return nil, err
	}
	return &RegistryApprovalIterator{contract: _Registry.contract, event: "Approval", logs: logs, sub: sub}, nil
}

// WatchApproval is a free log subscription operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed approved, uint256 indexed tokenId)
func (_Registry *RegistryFilterer) WatchApproval(opts *bind.WatchOpts, sink chan<- *RegistryApproval, owner []common.Address, approved []common.Address, tokenId []*big.Int) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var approvedRule []interface{}
	for _, approvedItem := range approved {
		approvedRule = append(approvedRule, approvedItem)
	}
	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}

	logs, sub, err := _Registry.contract.WatchLogs(opts, "Approval", ownerRule, approvedRule, tokenIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RegistryApproval)
				if err := _Registry.contract.UnpackLog(event, "Approval", log); err != nil {
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

// ParseApproval is a log parse operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed approved, uint256 indexed tokenId)
func (_Registry *RegistryFilterer) ParseApproval(log types.Log) (*RegistryApproval, error) {
	event := new(RegistryApproval)
	if err := _Registry.contract.UnpackLog(event, "Approval", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// RegistryApprovalForAllIterator is returned from FilterApprovalForAll and is used to iterate over the raw logs and unpacked data for ApprovalForAll events raised by the Registry contract.
type RegistryApprovalForAllIterator struct {
	Event *RegistryApprovalForAll // Event containing the contract specifics and raw log

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
func (it *RegistryApprovalForAllIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RegistryApprovalForAll)
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
		it.Event = new(RegistryApprovalForAll)
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
func (it *RegistryApprovalForAllIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RegistryApprovalForAllIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RegistryApprovalForAll represents a ApprovalForAll event raised by the Registry contract.
type RegistryApprovalForAll struct {
	Owner    common.Address
	Operator common.Address
	Approved bool
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterApprovalForAll is a free log retrieval operation binding the contract event 0x17307eab39ab6107e8899845ad3d59bd9653f200f220920489ca2b5937696c31.
//
// Solidity: event ApprovalForAll(address indexed owner, address indexed operator, bool approved)
func (_Registry *RegistryFilterer) FilterApprovalForAll(opts *bind.FilterOpts, owner []common.Address, operator []common.Address) (*RegistryApprovalForAllIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}

	logs, sub, err := _Registry.contract.FilterLogs(opts, "ApprovalForAll", ownerRule, operatorRule)
	if err != nil {
		return nil, err
	}
	return &RegistryApprovalForAllIterator{contract: _Registry.contract, event: "ApprovalForAll", logs: logs, sub: sub}, nil
}

// WatchApprovalForAll is a free log subscription operation binding the contract event 0x17307eab39ab6107e8899845ad3d59bd9653f200f220920489ca2b5937696c31.
//
// Solidity: event ApprovalForAll(address indexed owner, address indexed operator, bool approved)
func (_Registry *RegistryFilterer) WatchApprovalForAll(opts *bind.WatchOpts, sink chan<- *RegistryApprovalForAll, owner []common.Address, operator []common.Address) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}

	logs, sub, err := _Registry.contract.WatchLogs(opts, "ApprovalForAll", ownerRule, operatorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RegistryApprovalForAll)
				if err := _Registry.contract.UnpackLog(event, "ApprovalForAll", log); err != nil {
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

// ParseApprovalForAll is a log parse operation binding the contract event 0x17307eab39ab6107e8899845ad3d59bd9653f200f220920489ca2b5937696c31.
//
// Solidity: event ApprovalForAll(address indexed owner, address indexed operator, bool approved)
func (_Registry *RegistryFilterer) ParseApprovalForAll(log types.Log) (*RegistryApprovalForAll, error) {
	event := new(RegistryApprovalForAll)
	if err := _Registry.contract.UnpackLog(event, "ApprovalForAll", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// RegistryBatchMetadataUpdateIterator is returned from FilterBatchMetadataUpdate and is used to iterate over the raw logs and unpacked data for BatchMetadataUpdate events raised by the Registry contract.
type RegistryBatchMetadataUpdateIterator struct {
	Event *RegistryBatchMetadataUpdate // Event containing the contract specifics and raw log

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
func (it *RegistryBatchMetadataUpdateIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RegistryBatchMetadataUpdate)
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
		it.Event = new(RegistryBatchMetadataUpdate)
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
func (it *RegistryBatchMetadataUpdateIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RegistryBatchMetadataUpdateIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RegistryBatchMetadataUpdate represents a BatchMetadataUpdate event raised by the Registry contract.
type RegistryBatchMetadataUpdate struct {
	FromTokenId *big.Int
	ToTokenId   *big.Int
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterBatchMetadataUpdate is a free log retrieval operation binding the contract event 0x6bd5c950a8d8df17f772f5af37cb3655737899cbf903264b9795592da439661c.
//
// Solidity: event BatchMetadataUpdate(uint256 _fromTokenId, uint256 _toTokenId)
func (_Registry *RegistryFilterer) FilterBatchMetadataUpdate(opts *bind.FilterOpts) (*RegistryBatchMetadataUpdateIterator, error) {

	logs, sub, err := _Registry.contract.FilterLogs(opts, "BatchMetadataUpdate")
	if err != nil {
		return nil, err
	}
	return &RegistryBatchMetadataUpdateIterator{contract: _Registry.contract, event: "BatchMetadataUpdate", logs: logs, sub: sub}, nil
}

// WatchBatchMetadataUpdate is a free log subscription operation binding the contract event 0x6bd5c950a8d8df17f772f5af37cb3655737899cbf903264b9795592da439661c.
//
// Solidity: event BatchMetadataUpdate(uint256 _fromTokenId, uint256 _toTokenId)
func (_Registry *RegistryFilterer) WatchBatchMetadataUpdate(opts *bind.WatchOpts, sink chan<- *RegistryBatchMetadataUpdate) (event.Subscription, error) {

	logs, sub, err := _Registry.contract.WatchLogs(opts, "BatchMetadataUpdate")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RegistryBatchMetadataUpdate)
				if err := _Registry.contract.UnpackLog(event, "BatchMetadataUpdate", log); err != nil {
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

// ParseBatchMetadataUpdate is a log parse operation binding the contract event 0x6bd5c950a8d8df17f772f5af37cb3655737899cbf903264b9795592da439661c.
//
// Solidity: event BatchMetadataUpdate(uint256 _fromTokenId, uint256 _toTokenId)
func (_Registry *RegistryFilterer) ParseBatchMetadataUpdate(log types.Log) (*RegistryBatchMetadataUpdate, error) {
	event := new(RegistryBatchMetadataUpdate)
	if err := _Registry.contract.UnpackLog(event, "BatchMetadataUpdate", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// RegistryEIP712DomainChangedIterator is returned from FilterEIP712DomainChanged and is used to iterate over the raw logs and unpacked data for EIP712DomainChanged events raised by the Registry contract.
type RegistryEIP712DomainChangedIterator struct {
	Event *RegistryEIP712DomainChanged // Event containing the contract specifics and raw log

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
func (it *RegistryEIP712DomainChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RegistryEIP712DomainChanged)
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
		it.Event = new(RegistryEIP712DomainChanged)
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
func (it *RegistryEIP712DomainChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RegistryEIP712DomainChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RegistryEIP712DomainChanged represents a EIP712DomainChanged event raised by the Registry contract.
type RegistryEIP712DomainChanged struct {
	Raw types.Log // Blockchain specific contextual infos
}

// FilterEIP712DomainChanged is a free log retrieval operation binding the contract event 0x0a6387c9ea3628b88a633bb4f3b151770f70085117a15f9bf3787cda53f13d31.
//
// Solidity: event EIP712DomainChanged()
func (_Registry *RegistryFilterer) FilterEIP712DomainChanged(opts *bind.FilterOpts) (*RegistryEIP712DomainChangedIterator, error) {

	logs, sub, err := _Registry.contract.FilterLogs(opts, "EIP712DomainChanged")
	if err != nil {
		return nil, err
	}
	return &RegistryEIP712DomainChangedIterator{contract: _Registry.contract, event: "EIP712DomainChanged", logs: logs, sub: sub}, nil
}

// WatchEIP712DomainChanged is a free log subscription operation binding the contract event 0x0a6387c9ea3628b88a633bb4f3b151770f70085117a15f9bf3787cda53f13d31.
//
// Solidity: event EIP712DomainChanged()
func (_Registry *RegistryFilterer) WatchEIP712DomainChanged(opts *bind.WatchOpts, sink chan<- *RegistryEIP712DomainChanged) (event.Subscription, error) {

	logs, sub, err := _Registry.contract.WatchLogs(opts, "EIP712DomainChanged")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RegistryEIP712DomainChanged)
				if err := _Registry.contract.UnpackLog(event, "EIP712DomainChanged", log); err != nil {
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

// ParseEIP712DomainChanged is a log parse operation binding the contract event 0x0a6387c9ea3628b88a633bb4f3b151770f70085117a15f9bf3787cda53f13d31.
//
// Solidity: event EIP712DomainChanged()
func (_Registry *RegistryFilterer) ParseEIP712DomainChanged(log types.Log) (*RegistryEIP712DomainChanged, error) {
	event := new(RegistryEIP712DomainChanged)
	if err := _Registry.contract.UnpackLog(event, "EIP712DomainChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// RegistryInitializedIterator is returned from FilterInitialized and is used to iterate over the raw logs and unpacked data for Initialized events raised by the Registry contract.
type RegistryInitializedIterator struct {
	Event *RegistryInitialized // Event containing the contract specifics and raw log

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
func (it *RegistryInitializedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RegistryInitialized)
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
		it.Event = new(RegistryInitialized)
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
func (it *RegistryInitializedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RegistryInitializedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RegistryInitialized represents a Initialized event raised by the Registry contract.
type RegistryInitialized struct {
	Version uint64
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterInitialized is a free log retrieval operation binding the contract event 0xc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d2.
//
// Solidity: event Initialized(uint64 version)
func (_Registry *RegistryFilterer) FilterInitialized(opts *bind.FilterOpts) (*RegistryInitializedIterator, error) {

	logs, sub, err := _Registry.contract.FilterLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return &RegistryInitializedIterator{contract: _Registry.contract, event: "Initialized", logs: logs, sub: sub}, nil
}

// WatchInitialized is a free log subscription operation binding the contract event 0xc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d2.
//
// Solidity: event Initialized(uint64 version)
func (_Registry *RegistryFilterer) WatchInitialized(opts *bind.WatchOpts, sink chan<- *RegistryInitialized) (event.Subscription, error) {

	logs, sub, err := _Registry.contract.WatchLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RegistryInitialized)
				if err := _Registry.contract.UnpackLog(event, "Initialized", log); err != nil {
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
func (_Registry *RegistryFilterer) ParseInitialized(log types.Log) (*RegistryInitialized, error) {
	event := new(RegistryInitialized)
	if err := _Registry.contract.UnpackLog(event, "Initialized", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// RegistryMetadataSetIterator is returned from FilterMetadataSet and is used to iterate over the raw logs and unpacked data for MetadataSet events raised by the Registry contract.
type RegistryMetadataSetIterator struct {
	Event *RegistryMetadataSet // Event containing the contract specifics and raw log

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
func (it *RegistryMetadataSetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RegistryMetadataSet)
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
		it.Event = new(RegistryMetadataSet)
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
func (it *RegistryMetadataSetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RegistryMetadataSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RegistryMetadataSet represents a MetadataSet event raised by the Registry contract.
type RegistryMetadataSet struct {
	AgentId            *big.Int
	IndexedMetadataKey common.Hash
	MetadataKey        string
	MetadataValue      []byte
	Raw                types.Log // Blockchain specific contextual infos
}

// FilterMetadataSet is a free log retrieval operation binding the contract event 0x2c149ed548c6d2993cd73efe187df6eccabe4538091b33adbd25fafdb8a1468b.
//
// Solidity: event MetadataSet(uint256 indexed agentId, string indexed indexedMetadataKey, string metadataKey, bytes metadataValue)
func (_Registry *RegistryFilterer) FilterMetadataSet(opts *bind.FilterOpts, agentId []*big.Int, indexedMetadataKey []string) (*RegistryMetadataSetIterator, error) {

	var agentIdRule []interface{}
	for _, agentIdItem := range agentId {
		agentIdRule = append(agentIdRule, agentIdItem)
	}
	var indexedMetadataKeyRule []interface{}
	for _, indexedMetadataKeyItem := range indexedMetadataKey {
		indexedMetadataKeyRule = append(indexedMetadataKeyRule, indexedMetadataKeyItem)
	}

	logs, sub, err := _Registry.contract.FilterLogs(opts, "MetadataSet", agentIdRule, indexedMetadataKeyRule)
	if err != nil {
		return nil, err
	}
	return &RegistryMetadataSetIterator{contract: _Registry.contract, event: "MetadataSet", logs: logs, sub: sub}, nil
}

// WatchMetadataSet is a free log subscription operation binding the contract event 0x2c149ed548c6d2993cd73efe187df6eccabe4538091b33adbd25fafdb8a1468b.
//
// Solidity: event MetadataSet(uint256 indexed agentId, string indexed indexedMetadataKey, string metadataKey, bytes metadataValue)
func (_Registry *RegistryFilterer) WatchMetadataSet(opts *bind.WatchOpts, sink chan<- *RegistryMetadataSet, agentId []*big.Int, indexedMetadataKey []string) (event.Subscription, error) {

	var agentIdRule []interface{}
	for _, agentIdItem := range agentId {
		agentIdRule = append(agentIdRule, agentIdItem)
	}
	var indexedMetadataKeyRule []interface{}
	for _, indexedMetadataKeyItem := range indexedMetadataKey {
		indexedMetadataKeyRule = append(indexedMetadataKeyRule, indexedMetadataKeyItem)
	}

	logs, sub, err := _Registry.contract.WatchLogs(opts, "MetadataSet", agentIdRule, indexedMetadataKeyRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RegistryMetadataSet)
				if err := _Registry.contract.UnpackLog(event, "MetadataSet", log); err != nil {
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

// ParseMetadataSet is a log parse operation binding the contract event 0x2c149ed548c6d2993cd73efe187df6eccabe4538091b33adbd25fafdb8a1468b.
//
// Solidity: event MetadataSet(uint256 indexed agentId, string indexed indexedMetadataKey, string metadataKey, bytes metadataValue)
func (_Registry *RegistryFilterer) ParseMetadataSet(log types.Log) (*RegistryMetadataSet, error) {
	event := new(RegistryMetadataSet)
	if err := _Registry.contract.UnpackLog(event, "MetadataSet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// RegistryMetadataUpdateIterator is returned from FilterMetadataUpdate and is used to iterate over the raw logs and unpacked data for MetadataUpdate events raised by the Registry contract.
type RegistryMetadataUpdateIterator struct {
	Event *RegistryMetadataUpdate // Event containing the contract specifics and raw log

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
func (it *RegistryMetadataUpdateIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RegistryMetadataUpdate)
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
		it.Event = new(RegistryMetadataUpdate)
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
func (it *RegistryMetadataUpdateIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RegistryMetadataUpdateIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RegistryMetadataUpdate represents a MetadataUpdate event raised by the Registry contract.
type RegistryMetadataUpdate struct {
	TokenId *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterMetadataUpdate is a free log retrieval operation binding the contract event 0xf8e1a15aba9398e019f0b49df1a4fde98ee17ae345cb5f6b5e2c27f5033e8ce7.
//
// Solidity: event MetadataUpdate(uint256 _tokenId)
func (_Registry *RegistryFilterer) FilterMetadataUpdate(opts *bind.FilterOpts) (*RegistryMetadataUpdateIterator, error) {

	logs, sub, err := _Registry.contract.FilterLogs(opts, "MetadataUpdate")
	if err != nil {
		return nil, err
	}
	return &RegistryMetadataUpdateIterator{contract: _Registry.contract, event: "MetadataUpdate", logs: logs, sub: sub}, nil
}

// WatchMetadataUpdate is a free log subscription operation binding the contract event 0xf8e1a15aba9398e019f0b49df1a4fde98ee17ae345cb5f6b5e2c27f5033e8ce7.
//
// Solidity: event MetadataUpdate(uint256 _tokenId)
func (_Registry *RegistryFilterer) WatchMetadataUpdate(opts *bind.WatchOpts, sink chan<- *RegistryMetadataUpdate) (event.Subscription, error) {

	logs, sub, err := _Registry.contract.WatchLogs(opts, "MetadataUpdate")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RegistryMetadataUpdate)
				if err := _Registry.contract.UnpackLog(event, "MetadataUpdate", log); err != nil {
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

// ParseMetadataUpdate is a log parse operation binding the contract event 0xf8e1a15aba9398e019f0b49df1a4fde98ee17ae345cb5f6b5e2c27f5033e8ce7.
//
// Solidity: event MetadataUpdate(uint256 _tokenId)
func (_Registry *RegistryFilterer) ParseMetadataUpdate(log types.Log) (*RegistryMetadataUpdate, error) {
	event := new(RegistryMetadataUpdate)
	if err := _Registry.contract.UnpackLog(event, "MetadataUpdate", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// RegistryOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the Registry contract.
type RegistryOwnershipTransferredIterator struct {
	Event *RegistryOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *RegistryOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RegistryOwnershipTransferred)
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
		it.Event = new(RegistryOwnershipTransferred)
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
func (it *RegistryOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RegistryOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RegistryOwnershipTransferred represents a OwnershipTransferred event raised by the Registry contract.
type RegistryOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Registry *RegistryFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*RegistryOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Registry.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &RegistryOwnershipTransferredIterator{contract: _Registry.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Registry *RegistryFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *RegistryOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Registry.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RegistryOwnershipTransferred)
				if err := _Registry.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
func (_Registry *RegistryFilterer) ParseOwnershipTransferred(log types.Log) (*RegistryOwnershipTransferred, error) {
	event := new(RegistryOwnershipTransferred)
	if err := _Registry.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// RegistryRegisteredIterator is returned from FilterRegistered and is used to iterate over the raw logs and unpacked data for Registered events raised by the Registry contract.
type RegistryRegisteredIterator struct {
	Event *RegistryRegistered // Event containing the contract specifics and raw log

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
func (it *RegistryRegisteredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RegistryRegistered)
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
		it.Event = new(RegistryRegistered)
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
func (it *RegistryRegisteredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RegistryRegisteredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RegistryRegistered represents a Registered event raised by the Registry contract.
type RegistryRegistered struct {
	AgentId  *big.Int
	AgentURI string
	Owner    common.Address
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterRegistered is a free log retrieval operation binding the contract event 0xca52e62c367d81bb2e328eb795f7c7ba24afb478408a26c0e201d155c449bc4a.
//
// Solidity: event Registered(uint256 indexed agentId, string agentURI, address indexed owner)
func (_Registry *RegistryFilterer) FilterRegistered(opts *bind.FilterOpts, agentId []*big.Int, owner []common.Address) (*RegistryRegisteredIterator, error) {

	var agentIdRule []interface{}
	for _, agentIdItem := range agentId {
		agentIdRule = append(agentIdRule, agentIdItem)
	}

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}

	logs, sub, err := _Registry.contract.FilterLogs(opts, "Registered", agentIdRule, ownerRule)
	if err != nil {
		return nil, err
	}
	return &RegistryRegisteredIterator{contract: _Registry.contract, event: "Registered", logs: logs, sub: sub}, nil
}

// WatchRegistered is a free log subscription operation binding the contract event 0xca52e62c367d81bb2e328eb795f7c7ba24afb478408a26c0e201d155c449bc4a.
//
// Solidity: event Registered(uint256 indexed agentId, string agentURI, address indexed owner)
func (_Registry *RegistryFilterer) WatchRegistered(opts *bind.WatchOpts, sink chan<- *RegistryRegistered, agentId []*big.Int, owner []common.Address) (event.Subscription, error) {

	var agentIdRule []interface{}
	for _, agentIdItem := range agentId {
		agentIdRule = append(agentIdRule, agentIdItem)
	}

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}

	logs, sub, err := _Registry.contract.WatchLogs(opts, "Registered", agentIdRule, ownerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RegistryRegistered)
				if err := _Registry.contract.UnpackLog(event, "Registered", log); err != nil {
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

// ParseRegistered is a log parse operation binding the contract event 0xca52e62c367d81bb2e328eb795f7c7ba24afb478408a26c0e201d155c449bc4a.
//
// Solidity: event Registered(uint256 indexed agentId, string agentURI, address indexed owner)
func (_Registry *RegistryFilterer) ParseRegistered(log types.Log) (*RegistryRegistered, error) {
	event := new(RegistryRegistered)
	if err := _Registry.contract.UnpackLog(event, "Registered", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// RegistryTransferIterator is returned from FilterTransfer and is used to iterate over the raw logs and unpacked data for Transfer events raised by the Registry contract.
type RegistryTransferIterator struct {
	Event *RegistryTransfer // Event containing the contract specifics and raw log

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
func (it *RegistryTransferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RegistryTransfer)
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
		it.Event = new(RegistryTransfer)
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
func (it *RegistryTransferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RegistryTransferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RegistryTransfer represents a Transfer event raised by the Registry contract.
type RegistryTransfer struct {
	From    common.Address
	To      common.Address
	TokenId *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterTransfer is a free log retrieval operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 indexed tokenId)
func (_Registry *RegistryFilterer) FilterTransfer(opts *bind.FilterOpts, from []common.Address, to []common.Address, tokenId []*big.Int) (*RegistryTransferIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}
	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}

	logs, sub, err := _Registry.contract.FilterLogs(opts, "Transfer", fromRule, toRule, tokenIdRule)
	if err != nil {
		return nil, err
	}
	return &RegistryTransferIterator{contract: _Registry.contract, event: "Transfer", logs: logs, sub: sub}, nil
}

// WatchTransfer is a free log subscription operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 indexed tokenId)
func (_Registry *RegistryFilterer) WatchTransfer(opts *bind.WatchOpts, sink chan<- *RegistryTransfer, from []common.Address, to []common.Address, tokenId []*big.Int) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}
	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}

	logs, sub, err := _Registry.contract.WatchLogs(opts, "Transfer", fromRule, toRule, tokenIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RegistryTransfer)
				if err := _Registry.contract.UnpackLog(event, "Transfer", log); err != nil {
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

// ParseTransfer is a log parse operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 indexed tokenId)
func (_Registry *RegistryFilterer) ParseTransfer(log types.Log) (*RegistryTransfer, error) {
	event := new(RegistryTransfer)
	if err := _Registry.contract.UnpackLog(event, "Transfer", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// RegistryURIUpdatedIterator is returned from FilterURIUpdated and is used to iterate over the raw logs and unpacked data for URIUpdated events raised by the Registry contract.
type RegistryURIUpdatedIterator struct {
	Event *RegistryURIUpdated // Event containing the contract specifics and raw log

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
func (it *RegistryURIUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RegistryURIUpdated)
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
		it.Event = new(RegistryURIUpdated)
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
func (it *RegistryURIUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RegistryURIUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RegistryURIUpdated represents a URIUpdated event raised by the Registry contract.
type RegistryURIUpdated struct {
	AgentId   *big.Int
	NewURI    string
	UpdatedBy common.Address
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterURIUpdated is a free log retrieval operation binding the contract event 0x3a2c7fffc2cba7582c690e3b82c453ea02a308326a98a3ad7576c606336409fb.
//
// Solidity: event URIUpdated(uint256 indexed agentId, string newURI, address indexed updatedBy)
func (_Registry *RegistryFilterer) FilterURIUpdated(opts *bind.FilterOpts, agentId []*big.Int, updatedBy []common.Address) (*RegistryURIUpdatedIterator, error) {

	var agentIdRule []interface{}
	for _, agentIdItem := range agentId {
		agentIdRule = append(agentIdRule, agentIdItem)
	}

	var updatedByRule []interface{}
	for _, updatedByItem := range updatedBy {
		updatedByRule = append(updatedByRule, updatedByItem)
	}

	logs, sub, err := _Registry.contract.FilterLogs(opts, "URIUpdated", agentIdRule, updatedByRule)
	if err != nil {
		return nil, err
	}
	return &RegistryURIUpdatedIterator{contract: _Registry.contract, event: "URIUpdated", logs: logs, sub: sub}, nil
}

// WatchURIUpdated is a free log subscription operation binding the contract event 0x3a2c7fffc2cba7582c690e3b82c453ea02a308326a98a3ad7576c606336409fb.
//
// Solidity: event URIUpdated(uint256 indexed agentId, string newURI, address indexed updatedBy)
func (_Registry *RegistryFilterer) WatchURIUpdated(opts *bind.WatchOpts, sink chan<- *RegistryURIUpdated, agentId []*big.Int, updatedBy []common.Address) (event.Subscription, error) {

	var agentIdRule []interface{}
	for _, agentIdItem := range agentId {
		agentIdRule = append(agentIdRule, agentIdItem)
	}

	var updatedByRule []interface{}
	for _, updatedByItem := range updatedBy {
		updatedByRule = append(updatedByRule, updatedByItem)
	}

	logs, sub, err := _Registry.contract.WatchLogs(opts, "URIUpdated", agentIdRule, updatedByRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RegistryURIUpdated)
				if err := _Registry.contract.UnpackLog(event, "URIUpdated", log); err != nil {
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

// ParseURIUpdated is a log parse operation binding the contract event 0x3a2c7fffc2cba7582c690e3b82c453ea02a308326a98a3ad7576c606336409fb.
//
// Solidity: event URIUpdated(uint256 indexed agentId, string newURI, address indexed updatedBy)
func (_Registry *RegistryFilterer) ParseURIUpdated(log types.Log) (*RegistryURIUpdated, error) {
	event := new(RegistryURIUpdated)
	if err := _Registry.contract.UnpackLog(event, "URIUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// RegistryUpgradedIterator is returned from FilterUpgraded and is used to iterate over the raw logs and unpacked data for Upgraded events raised by the Registry contract.
type RegistryUpgradedIterator struct {
	Event *RegistryUpgraded // Event containing the contract specifics and raw log

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
func (it *RegistryUpgradedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RegistryUpgraded)
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
		it.Event = new(RegistryUpgraded)
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
func (it *RegistryUpgradedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RegistryUpgradedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RegistryUpgraded represents a Upgraded event raised by the Registry contract.
type RegistryUpgraded struct {
	Implementation common.Address
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterUpgraded is a free log retrieval operation binding the contract event 0xbc7cd75a20ee27fd9adebab32041f755214dbc6bffa90cc0225b39da2e5c2d3b.
//
// Solidity: event Upgraded(address indexed implementation)
func (_Registry *RegistryFilterer) FilterUpgraded(opts *bind.FilterOpts, implementation []common.Address) (*RegistryUpgradedIterator, error) {

	var implementationRule []interface{}
	for _, implementationItem := range implementation {
		implementationRule = append(implementationRule, implementationItem)
	}

	logs, sub, err := _Registry.contract.FilterLogs(opts, "Upgraded", implementationRule)
	if err != nil {
		return nil, err
	}
	return &RegistryUpgradedIterator{contract: _Registry.contract, event: "Upgraded", logs: logs, sub: sub}, nil
}

// WatchUpgraded is a free log subscription operation binding the contract event 0xbc7cd75a20ee27fd9adebab32041f755214dbc6bffa90cc0225b39da2e5c2d3b.
//
// Solidity: event Upgraded(address indexed implementation)
func (_Registry *RegistryFilterer) WatchUpgraded(opts *bind.WatchOpts, sink chan<- *RegistryUpgraded, implementation []common.Address) (event.Subscription, error) {

	var implementationRule []interface{}
	for _, implementationItem := range implementation {
		implementationRule = append(implementationRule, implementationItem)
	}

	logs, sub, err := _Registry.contract.WatchLogs(opts, "Upgraded", implementationRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RegistryUpgraded)
				if err := _Registry.contract.UnpackLog(event, "Upgraded", log); err != nil {
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
func (_Registry *RegistryFilterer) ParseUpgraded(log types.Log) (*RegistryUpgraded, error) {
	event := new(RegistryUpgraded)
	if err := _Registry.contract.UnpackLog(event, "Upgraded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
