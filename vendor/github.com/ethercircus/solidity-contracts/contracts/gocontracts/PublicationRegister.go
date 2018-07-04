// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package gocontracts

import (
	"math/big"
	"strings"

	ethereum "github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/event"
)

// PublicationRegisterABI is the input ABI used to generate the binding from.
const PublicationRegisterABI = "[{\"constant\":false,\"inputs\":[{\"name\":\"whichPublication\",\"type\":\"uint256\"},{\"name\":\"newAdminPaymentPercentage\",\"type\":\"uint8\"}],\"name\":\"decreaseAdminPaymentPercentage\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"whichPublication\",\"type\":\"uint256\"}],\"name\":\"getNumPublished\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"whichPublication\",\"type\":\"uint256\"}],\"name\":\"getAdmin\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"publicationIndex\",\"outputs\":[{\"name\":\"name\",\"type\":\"string\"},{\"name\":\"metaData\",\"type\":\"string\"},{\"name\":\"admin\",\"type\":\"address\"},{\"name\":\"open\",\"type\":\"bool\"},{\"name\":\"numPublished\",\"type\":\"uint256\"},{\"name\":\"minSupportCostWei\",\"type\":\"uint256\"},{\"name\":\"adminPaymentPercentage\",\"type\":\"uint8\"},{\"name\":\"uniqueSupporters\",\"type\":\"uint256\"},{\"name\":\"adminClaimOwedWei\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"whichPublication\",\"type\":\"uint256\"},{\"name\":\"contentIndex\",\"type\":\"uint256\"}],\"name\":\"getContentAuthor\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"name\",\"type\":\"string\"},{\"name\":\"metaData\",\"type\":\"string\"},{\"name\":\"minUpVoteCostWei\",\"type\":\"uint256\"},{\"name\":\"adminPaymentPercentage\",\"type\":\"uint8\"}],\"name\":\"createPublication\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"whichPublication\",\"type\":\"uint256\"},{\"name\":\"contentIndex\",\"type\":\"uint256\"}],\"name\":\"getContent\",\"outputs\":[{\"name\":\"\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"name\",\"type\":\"string\"}],\"name\":\"checkNameTaken\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"whichPublication\",\"type\":\"uint256\"},{\"name\":\"contentIndex\",\"type\":\"uint256\"}],\"name\":\"getContentUniqueSupporters\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"whichPublication\",\"type\":\"uint256\"},{\"name\":\"contentIndex\",\"type\":\"uint256\"},{\"name\":\"commentIndex\",\"type\":\"uint256\"}],\"name\":\"getContentCommentByIndex\",\"outputs\":[{\"name\":\"\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"whichPublication\",\"type\":\"uint256\"}],\"name\":\"withdrawAuthorClaim\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"whichPublication\",\"type\":\"uint256\"},{\"name\":\"author\",\"type\":\"address\"}],\"name\":\"checkAuthorClaim\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"whichPublication\",\"type\":\"uint256\"},{\"name\":\"contentIndex\",\"type\":\"uint256\"}],\"name\":\"getContentBytes\",\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\"},{\"name\":\"\",\"type\":\"bytes32\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"whichPublication\",\"type\":\"uint256\"},{\"name\":\"contentIndex\",\"type\":\"uint256\"}],\"name\":\"getContentNumComments\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"whichPublication\",\"type\":\"uint256\"},{\"name\":\"index\",\"type\":\"uint256\"}],\"name\":\"publishContent\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"whichPublication\",\"type\":\"uint256\"},{\"name\":\"author\",\"type\":\"address\"},{\"name\":\"giveAccess\",\"type\":\"bool\"}],\"name\":\"permissionAuthor\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"whichPublication\",\"type\":\"uint256\"}],\"name\":\"withdrawAdminClaim\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"whichPublication\",\"type\":\"uint256\"},{\"name\":\"postIndex\",\"type\":\"uint256\"},{\"name\":\"optionalComment\",\"type\":\"string\"}],\"name\":\"supportPost\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":true,\"stateMutability\":\"payable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"userContentRegister\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"whichPublication\",\"type\":\"uint256\"},{\"name\":\"metaData\",\"type\":\"string\"}],\"name\":\"updateMetaData\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"whichPublication\",\"type\":\"uint256\"},{\"name\":\"contentIndex\",\"type\":\"uint256\"}],\"name\":\"getContentRevenue\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"numPublications\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"name\",\"type\":\"string\"}],\"name\":\"getPublicationIndex\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"name\":\"userContentRegisterAddress\",\"type\":\"address\"},{\"name\":\"stringBytes32UtilAddress\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"data\",\"type\":\"string\"},{\"indexed\":false,\"name\":\"author\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"whichPublication\",\"type\":\"uint256\"}],\"name\":\"StoreData\",\"type\":\"event\"}]"

// PublicationRegister is an auto generated Go binding around an Ethereum contract.
type PublicationRegister struct {
	PublicationRegisterCaller     // Read-only binding to the contract
	PublicationRegisterTransactor // Write-only binding to the contract
	PublicationRegisterFilterer   // Log filterer for contract events
}

// PublicationRegisterCaller is an auto generated read-only Go binding around an Ethereum contract.
type PublicationRegisterCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PublicationRegisterTransactor is an auto generated write-only Go binding around an Ethereum contract.
type PublicationRegisterTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PublicationRegisterFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type PublicationRegisterFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PublicationRegisterSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type PublicationRegisterSession struct {
	Contract     *PublicationRegister // Generic contract binding to set the session for
	CallOpts     bind.CallOpts        // Call options to use throughout this session
	TransactOpts bind.TransactOpts    // Transaction auth options to use throughout this session
}

// PublicationRegisterCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type PublicationRegisterCallerSession struct {
	Contract *PublicationRegisterCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts              // Call options to use throughout this session
}

// PublicationRegisterTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type PublicationRegisterTransactorSession struct {
	Contract     *PublicationRegisterTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts              // Transaction auth options to use throughout this session
}

// PublicationRegisterRaw is an auto generated low-level Go binding around an Ethereum contract.
type PublicationRegisterRaw struct {
	Contract *PublicationRegister // Generic contract binding to access the raw methods on
}

// PublicationRegisterCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type PublicationRegisterCallerRaw struct {
	Contract *PublicationRegisterCaller // Generic read-only contract binding to access the raw methods on
}

// PublicationRegisterTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type PublicationRegisterTransactorRaw struct {
	Contract *PublicationRegisterTransactor // Generic write-only contract binding to access the raw methods on
}

// NewPublicationRegister creates a new instance of PublicationRegister, bound to a specific deployed contract.
func NewPublicationRegister(address common.Address, backend bind.ContractBackend) (*PublicationRegister, error) {
	contract, err := bindPublicationRegister(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &PublicationRegister{PublicationRegisterCaller: PublicationRegisterCaller{contract: contract}, PublicationRegisterTransactor: PublicationRegisterTransactor{contract: contract}, PublicationRegisterFilterer: PublicationRegisterFilterer{contract: contract}}, nil
}

// NewPublicationRegisterCaller creates a new read-only instance of PublicationRegister, bound to a specific deployed contract.
func NewPublicationRegisterCaller(address common.Address, caller bind.ContractCaller) (*PublicationRegisterCaller, error) {
	contract, err := bindPublicationRegister(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &PublicationRegisterCaller{contract: contract}, nil
}

// NewPublicationRegisterTransactor creates a new write-only instance of PublicationRegister, bound to a specific deployed contract.
func NewPublicationRegisterTransactor(address common.Address, transactor bind.ContractTransactor) (*PublicationRegisterTransactor, error) {
	contract, err := bindPublicationRegister(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &PublicationRegisterTransactor{contract: contract}, nil
}

// NewPublicationRegisterFilterer creates a new log filterer instance of PublicationRegister, bound to a specific deployed contract.
func NewPublicationRegisterFilterer(address common.Address, filterer bind.ContractFilterer) (*PublicationRegisterFilterer, error) {
	contract, err := bindPublicationRegister(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &PublicationRegisterFilterer{contract: contract}, nil
}

// bindPublicationRegister binds a generic wrapper to an already deployed contract.
func bindPublicationRegister(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(PublicationRegisterABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_PublicationRegister *PublicationRegisterRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _PublicationRegister.Contract.PublicationRegisterCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_PublicationRegister *PublicationRegisterRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _PublicationRegister.Contract.PublicationRegisterTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_PublicationRegister *PublicationRegisterRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _PublicationRegister.Contract.PublicationRegisterTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_PublicationRegister *PublicationRegisterCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _PublicationRegister.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_PublicationRegister *PublicationRegisterTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _PublicationRegister.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_PublicationRegister *PublicationRegisterTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _PublicationRegister.Contract.contract.Transact(opts, method, params...)
}

// CheckAuthorClaim is a free data retrieval call binding the contract method 0x81230c06.
//
// Solidity: function checkAuthorClaim(whichPublication uint256, author address) constant returns(uint256)
func (_PublicationRegister *PublicationRegisterCaller) CheckAuthorClaim(opts *bind.CallOpts, whichPublication *big.Int, author common.Address) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _PublicationRegister.contract.Call(opts, out, "checkAuthorClaim", whichPublication, author)
	return *ret0, err
}

// CheckAuthorClaim is a free data retrieval call binding the contract method 0x81230c06.
//
// Solidity: function checkAuthorClaim(whichPublication uint256, author address) constant returns(uint256)
func (_PublicationRegister *PublicationRegisterSession) CheckAuthorClaim(whichPublication *big.Int, author common.Address) (*big.Int, error) {
	return _PublicationRegister.Contract.CheckAuthorClaim(&_PublicationRegister.CallOpts, whichPublication, author)
}

// CheckAuthorClaim is a free data retrieval call binding the contract method 0x81230c06.
//
// Solidity: function checkAuthorClaim(whichPublication uint256, author address) constant returns(uint256)
func (_PublicationRegister *PublicationRegisterCallerSession) CheckAuthorClaim(whichPublication *big.Int, author common.Address) (*big.Int, error) {
	return _PublicationRegister.Contract.CheckAuthorClaim(&_PublicationRegister.CallOpts, whichPublication, author)
}

// CheckNameTaken is a free data retrieval call binding the contract method 0x57612c31.
//
// Solidity: function checkNameTaken(name string) constant returns(bool)
func (_PublicationRegister *PublicationRegisterCaller) CheckNameTaken(opts *bind.CallOpts, name string) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _PublicationRegister.contract.Call(opts, out, "checkNameTaken", name)
	return *ret0, err
}

// CheckNameTaken is a free data retrieval call binding the contract method 0x57612c31.
//
// Solidity: function checkNameTaken(name string) constant returns(bool)
func (_PublicationRegister *PublicationRegisterSession) CheckNameTaken(name string) (bool, error) {
	return _PublicationRegister.Contract.CheckNameTaken(&_PublicationRegister.CallOpts, name)
}

// CheckNameTaken is a free data retrieval call binding the contract method 0x57612c31.
//
// Solidity: function checkNameTaken(name string) constant returns(bool)
func (_PublicationRegister *PublicationRegisterCallerSession) CheckNameTaken(name string) (bool, error) {
	return _PublicationRegister.Contract.CheckNameTaken(&_PublicationRegister.CallOpts, name)
}

// GetAdmin is a free data retrieval call binding the contract method 0x111fd88b.
//
// Solidity: function getAdmin(whichPublication uint256) constant returns(address)
func (_PublicationRegister *PublicationRegisterCaller) GetAdmin(opts *bind.CallOpts, whichPublication *big.Int) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _PublicationRegister.contract.Call(opts, out, "getAdmin", whichPublication)
	return *ret0, err
}

// GetAdmin is a free data retrieval call binding the contract method 0x111fd88b.
//
// Solidity: function getAdmin(whichPublication uint256) constant returns(address)
func (_PublicationRegister *PublicationRegisterSession) GetAdmin(whichPublication *big.Int) (common.Address, error) {
	return _PublicationRegister.Contract.GetAdmin(&_PublicationRegister.CallOpts, whichPublication)
}

// GetAdmin is a free data retrieval call binding the contract method 0x111fd88b.
//
// Solidity: function getAdmin(whichPublication uint256) constant returns(address)
func (_PublicationRegister *PublicationRegisterCallerSession) GetAdmin(whichPublication *big.Int) (common.Address, error) {
	return _PublicationRegister.Contract.GetAdmin(&_PublicationRegister.CallOpts, whichPublication)
}

// GetContent is a free data retrieval call binding the contract method 0x31730a1d.
//
// Solidity: function getContent(whichPublication uint256, contentIndex uint256) constant returns(string)
func (_PublicationRegister *PublicationRegisterCaller) GetContent(opts *bind.CallOpts, whichPublication *big.Int, contentIndex *big.Int) (string, error) {
	var (
		ret0 = new(string)
	)
	out := ret0
	err := _PublicationRegister.contract.Call(opts, out, "getContent", whichPublication, contentIndex)
	return *ret0, err
}

// GetContent is a free data retrieval call binding the contract method 0x31730a1d.
//
// Solidity: function getContent(whichPublication uint256, contentIndex uint256) constant returns(string)
func (_PublicationRegister *PublicationRegisterSession) GetContent(whichPublication *big.Int, contentIndex *big.Int) (string, error) {
	return _PublicationRegister.Contract.GetContent(&_PublicationRegister.CallOpts, whichPublication, contentIndex)
}

// GetContent is a free data retrieval call binding the contract method 0x31730a1d.
//
// Solidity: function getContent(whichPublication uint256, contentIndex uint256) constant returns(string)
func (_PublicationRegister *PublicationRegisterCallerSession) GetContent(whichPublication *big.Int, contentIndex *big.Int) (string, error) {
	return _PublicationRegister.Contract.GetContent(&_PublicationRegister.CallOpts, whichPublication, contentIndex)
}

// GetContentAuthor is a free data retrieval call binding the contract method 0x23f1811d.
//
// Solidity: function getContentAuthor(whichPublication uint256, contentIndex uint256) constant returns(address)
func (_PublicationRegister *PublicationRegisterCaller) GetContentAuthor(opts *bind.CallOpts, whichPublication *big.Int, contentIndex *big.Int) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _PublicationRegister.contract.Call(opts, out, "getContentAuthor", whichPublication, contentIndex)
	return *ret0, err
}

// GetContentAuthor is a free data retrieval call binding the contract method 0x23f1811d.
//
// Solidity: function getContentAuthor(whichPublication uint256, contentIndex uint256) constant returns(address)
func (_PublicationRegister *PublicationRegisterSession) GetContentAuthor(whichPublication *big.Int, contentIndex *big.Int) (common.Address, error) {
	return _PublicationRegister.Contract.GetContentAuthor(&_PublicationRegister.CallOpts, whichPublication, contentIndex)
}

// GetContentAuthor is a free data retrieval call binding the contract method 0x23f1811d.
//
// Solidity: function getContentAuthor(whichPublication uint256, contentIndex uint256) constant returns(address)
func (_PublicationRegister *PublicationRegisterCallerSession) GetContentAuthor(whichPublication *big.Int, contentIndex *big.Int) (common.Address, error) {
	return _PublicationRegister.Contract.GetContentAuthor(&_PublicationRegister.CallOpts, whichPublication, contentIndex)
}

// GetContentBytes is a free data retrieval call binding the contract method 0x89557910.
//
// Solidity: function getContentBytes(whichPublication uint256, contentIndex uint256) constant returns(bytes32, bytes32)
func (_PublicationRegister *PublicationRegisterCaller) GetContentBytes(opts *bind.CallOpts, whichPublication *big.Int, contentIndex *big.Int) ([32]byte, [32]byte, error) {
	var (
		ret0 = new([32]byte)
		ret1 = new([32]byte)
	)
	out := &[]interface{}{
		ret0,
		ret1,
	}
	err := _PublicationRegister.contract.Call(opts, out, "getContentBytes", whichPublication, contentIndex)
	return *ret0, *ret1, err
}

// GetContentBytes is a free data retrieval call binding the contract method 0x89557910.
//
// Solidity: function getContentBytes(whichPublication uint256, contentIndex uint256) constant returns(bytes32, bytes32)
func (_PublicationRegister *PublicationRegisterSession) GetContentBytes(whichPublication *big.Int, contentIndex *big.Int) ([32]byte, [32]byte, error) {
	return _PublicationRegister.Contract.GetContentBytes(&_PublicationRegister.CallOpts, whichPublication, contentIndex)
}

// GetContentBytes is a free data retrieval call binding the contract method 0x89557910.
//
// Solidity: function getContentBytes(whichPublication uint256, contentIndex uint256) constant returns(bytes32, bytes32)
func (_PublicationRegister *PublicationRegisterCallerSession) GetContentBytes(whichPublication *big.Int, contentIndex *big.Int) ([32]byte, [32]byte, error) {
	return _PublicationRegister.Contract.GetContentBytes(&_PublicationRegister.CallOpts, whichPublication, contentIndex)
}

// GetContentCommentByIndex is a free data retrieval call binding the contract method 0x64a5c57e.
//
// Solidity: function getContentCommentByIndex(whichPublication uint256, contentIndex uint256, commentIndex uint256) constant returns(string)
func (_PublicationRegister *PublicationRegisterCaller) GetContentCommentByIndex(opts *bind.CallOpts, whichPublication *big.Int, contentIndex *big.Int, commentIndex *big.Int) (string, error) {
	var (
		ret0 = new(string)
	)
	out := ret0
	err := _PublicationRegister.contract.Call(opts, out, "getContentCommentByIndex", whichPublication, contentIndex, commentIndex)
	return *ret0, err
}

// GetContentCommentByIndex is a free data retrieval call binding the contract method 0x64a5c57e.
//
// Solidity: function getContentCommentByIndex(whichPublication uint256, contentIndex uint256, commentIndex uint256) constant returns(string)
func (_PublicationRegister *PublicationRegisterSession) GetContentCommentByIndex(whichPublication *big.Int, contentIndex *big.Int, commentIndex *big.Int) (string, error) {
	return _PublicationRegister.Contract.GetContentCommentByIndex(&_PublicationRegister.CallOpts, whichPublication, contentIndex, commentIndex)
}

// GetContentCommentByIndex is a free data retrieval call binding the contract method 0x64a5c57e.
//
// Solidity: function getContentCommentByIndex(whichPublication uint256, contentIndex uint256, commentIndex uint256) constant returns(string)
func (_PublicationRegister *PublicationRegisterCallerSession) GetContentCommentByIndex(whichPublication *big.Int, contentIndex *big.Int, commentIndex *big.Int) (string, error) {
	return _PublicationRegister.Contract.GetContentCommentByIndex(&_PublicationRegister.CallOpts, whichPublication, contentIndex, commentIndex)
}

// GetContentNumComments is a free data retrieval call binding the contract method 0x89b3491d.
//
// Solidity: function getContentNumComments(whichPublication uint256, contentIndex uint256) constant returns(uint256)
func (_PublicationRegister *PublicationRegisterCaller) GetContentNumComments(opts *bind.CallOpts, whichPublication *big.Int, contentIndex *big.Int) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _PublicationRegister.contract.Call(opts, out, "getContentNumComments", whichPublication, contentIndex)
	return *ret0, err
}

// GetContentNumComments is a free data retrieval call binding the contract method 0x89b3491d.
//
// Solidity: function getContentNumComments(whichPublication uint256, contentIndex uint256) constant returns(uint256)
func (_PublicationRegister *PublicationRegisterSession) GetContentNumComments(whichPublication *big.Int, contentIndex *big.Int) (*big.Int, error) {
	return _PublicationRegister.Contract.GetContentNumComments(&_PublicationRegister.CallOpts, whichPublication, contentIndex)
}

// GetContentNumComments is a free data retrieval call binding the contract method 0x89b3491d.
//
// Solidity: function getContentNumComments(whichPublication uint256, contentIndex uint256) constant returns(uint256)
func (_PublicationRegister *PublicationRegisterCallerSession) GetContentNumComments(whichPublication *big.Int, contentIndex *big.Int) (*big.Int, error) {
	return _PublicationRegister.Contract.GetContentNumComments(&_PublicationRegister.CallOpts, whichPublication, contentIndex)
}

// GetContentRevenue is a free data retrieval call binding the contract method 0xeb404954.
//
// Solidity: function getContentRevenue(whichPublication uint256, contentIndex uint256) constant returns(uint256)
func (_PublicationRegister *PublicationRegisterCaller) GetContentRevenue(opts *bind.CallOpts, whichPublication *big.Int, contentIndex *big.Int) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _PublicationRegister.contract.Call(opts, out, "getContentRevenue", whichPublication, contentIndex)
	return *ret0, err
}

// GetContentRevenue is a free data retrieval call binding the contract method 0xeb404954.
//
// Solidity: function getContentRevenue(whichPublication uint256, contentIndex uint256) constant returns(uint256)
func (_PublicationRegister *PublicationRegisterSession) GetContentRevenue(whichPublication *big.Int, contentIndex *big.Int) (*big.Int, error) {
	return _PublicationRegister.Contract.GetContentRevenue(&_PublicationRegister.CallOpts, whichPublication, contentIndex)
}

// GetContentRevenue is a free data retrieval call binding the contract method 0xeb404954.
//
// Solidity: function getContentRevenue(whichPublication uint256, contentIndex uint256) constant returns(uint256)
func (_PublicationRegister *PublicationRegisterCallerSession) GetContentRevenue(whichPublication *big.Int, contentIndex *big.Int) (*big.Int, error) {
	return _PublicationRegister.Contract.GetContentRevenue(&_PublicationRegister.CallOpts, whichPublication, contentIndex)
}

// GetContentUniqueSupporters is a free data retrieval call binding the contract method 0x60cf71b5.
//
// Solidity: function getContentUniqueSupporters(whichPublication uint256, contentIndex uint256) constant returns(uint256)
func (_PublicationRegister *PublicationRegisterCaller) GetContentUniqueSupporters(opts *bind.CallOpts, whichPublication *big.Int, contentIndex *big.Int) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _PublicationRegister.contract.Call(opts, out, "getContentUniqueSupporters", whichPublication, contentIndex)
	return *ret0, err
}

// GetContentUniqueSupporters is a free data retrieval call binding the contract method 0x60cf71b5.
//
// Solidity: function getContentUniqueSupporters(whichPublication uint256, contentIndex uint256) constant returns(uint256)
func (_PublicationRegister *PublicationRegisterSession) GetContentUniqueSupporters(whichPublication *big.Int, contentIndex *big.Int) (*big.Int, error) {
	return _PublicationRegister.Contract.GetContentUniqueSupporters(&_PublicationRegister.CallOpts, whichPublication, contentIndex)
}

// GetContentUniqueSupporters is a free data retrieval call binding the contract method 0x60cf71b5.
//
// Solidity: function getContentUniqueSupporters(whichPublication uint256, contentIndex uint256) constant returns(uint256)
func (_PublicationRegister *PublicationRegisterCallerSession) GetContentUniqueSupporters(whichPublication *big.Int, contentIndex *big.Int) (*big.Int, error) {
	return _PublicationRegister.Contract.GetContentUniqueSupporters(&_PublicationRegister.CallOpts, whichPublication, contentIndex)
}

// GetNumPublished is a free data retrieval call binding the contract method 0x075fdd1c.
//
// Solidity: function getNumPublished(whichPublication uint256) constant returns(uint256)
func (_PublicationRegister *PublicationRegisterCaller) GetNumPublished(opts *bind.CallOpts, whichPublication *big.Int) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _PublicationRegister.contract.Call(opts, out, "getNumPublished", whichPublication)
	return *ret0, err
}

// GetNumPublished is a free data retrieval call binding the contract method 0x075fdd1c.
//
// Solidity: function getNumPublished(whichPublication uint256) constant returns(uint256)
func (_PublicationRegister *PublicationRegisterSession) GetNumPublished(whichPublication *big.Int) (*big.Int, error) {
	return _PublicationRegister.Contract.GetNumPublished(&_PublicationRegister.CallOpts, whichPublication)
}

// GetNumPublished is a free data retrieval call binding the contract method 0x075fdd1c.
//
// Solidity: function getNumPublished(whichPublication uint256) constant returns(uint256)
func (_PublicationRegister *PublicationRegisterCallerSession) GetNumPublished(whichPublication *big.Int) (*big.Int, error) {
	return _PublicationRegister.Contract.GetNumPublished(&_PublicationRegister.CallOpts, whichPublication)
}

// GetPublicationIndex is a free data retrieval call binding the contract method 0xf6b19981.
//
// Solidity: function getPublicationIndex(name string) constant returns(uint256)
func (_PublicationRegister *PublicationRegisterCaller) GetPublicationIndex(opts *bind.CallOpts, name string) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _PublicationRegister.contract.Call(opts, out, "getPublicationIndex", name)
	return *ret0, err
}

// GetPublicationIndex is a free data retrieval call binding the contract method 0xf6b19981.
//
// Solidity: function getPublicationIndex(name string) constant returns(uint256)
func (_PublicationRegister *PublicationRegisterSession) GetPublicationIndex(name string) (*big.Int, error) {
	return _PublicationRegister.Contract.GetPublicationIndex(&_PublicationRegister.CallOpts, name)
}

// GetPublicationIndex is a free data retrieval call binding the contract method 0xf6b19981.
//
// Solidity: function getPublicationIndex(name string) constant returns(uint256)
func (_PublicationRegister *PublicationRegisterCallerSession) GetPublicationIndex(name string) (*big.Int, error) {
	return _PublicationRegister.Contract.GetPublicationIndex(&_PublicationRegister.CallOpts, name)
}

// NumPublications is a free data retrieval call binding the contract method 0xf3acebc8.
//
// Solidity: function numPublications() constant returns(uint256)
func (_PublicationRegister *PublicationRegisterCaller) NumPublications(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _PublicationRegister.contract.Call(opts, out, "numPublications")
	return *ret0, err
}

// NumPublications is a free data retrieval call binding the contract method 0xf3acebc8.
//
// Solidity: function numPublications() constant returns(uint256)
func (_PublicationRegister *PublicationRegisterSession) NumPublications() (*big.Int, error) {
	return _PublicationRegister.Contract.NumPublications(&_PublicationRegister.CallOpts)
}

// NumPublications is a free data retrieval call binding the contract method 0xf3acebc8.
//
// Solidity: function numPublications() constant returns(uint256)
func (_PublicationRegister *PublicationRegisterCallerSession) NumPublications() (*big.Int, error) {
	return _PublicationRegister.Contract.NumPublications(&_PublicationRegister.CallOpts)
}

// PublicationIndex is a free data retrieval call binding the contract method 0x1d51c73c.
//
// Solidity: function publicationIndex( uint256) constant returns(name string, metaData string, admin address, open bool, numPublished uint256, minSupportCostWei uint256, adminPaymentPercentage uint8, uniqueSupporters uint256, adminClaimOwedWei uint256)
func (_PublicationRegister *PublicationRegisterCaller) PublicationIndex(opts *bind.CallOpts, arg0 *big.Int) (struct {
	Name                   string
	MetaData               string
	Admin                  common.Address
	Open                   bool
	NumPublished           *big.Int
	MinSupportCostWei      *big.Int
	AdminPaymentPercentage uint8
	UniqueSupporters       *big.Int
	AdminClaimOwedWei      *big.Int
}, error) {
	ret := new(struct {
		Name                   string
		MetaData               string
		Admin                  common.Address
		Open                   bool
		NumPublished           *big.Int
		MinSupportCostWei      *big.Int
		AdminPaymentPercentage uint8
		UniqueSupporters       *big.Int
		AdminClaimOwedWei      *big.Int
	})
	out := ret
	err := _PublicationRegister.contract.Call(opts, out, "publicationIndex", arg0)
	return *ret, err
}

// PublicationIndex is a free data retrieval call binding the contract method 0x1d51c73c.
//
// Solidity: function publicationIndex( uint256) constant returns(name string, metaData string, admin address, open bool, numPublished uint256, minSupportCostWei uint256, adminPaymentPercentage uint8, uniqueSupporters uint256, adminClaimOwedWei uint256)
func (_PublicationRegister *PublicationRegisterSession) PublicationIndex(arg0 *big.Int) (struct {
	Name                   string
	MetaData               string
	Admin                  common.Address
	Open                   bool
	NumPublished           *big.Int
	MinSupportCostWei      *big.Int
	AdminPaymentPercentage uint8
	UniqueSupporters       *big.Int
	AdminClaimOwedWei      *big.Int
}, error) {
	return _PublicationRegister.Contract.PublicationIndex(&_PublicationRegister.CallOpts, arg0)
}

// PublicationIndex is a free data retrieval call binding the contract method 0x1d51c73c.
//
// Solidity: function publicationIndex( uint256) constant returns(name string, metaData string, admin address, open bool, numPublished uint256, minSupportCostWei uint256, adminPaymentPercentage uint8, uniqueSupporters uint256, adminClaimOwedWei uint256)
func (_PublicationRegister *PublicationRegisterCallerSession) PublicationIndex(arg0 *big.Int) (struct {
	Name                   string
	MetaData               string
	Admin                  common.Address
	Open                   bool
	NumPublished           *big.Int
	MinSupportCostWei      *big.Int
	AdminPaymentPercentage uint8
	UniqueSupporters       *big.Int
	AdminClaimOwedWei      *big.Int
}, error) {
	return _PublicationRegister.Contract.PublicationIndex(&_PublicationRegister.CallOpts, arg0)
}

// UserContentRegister is a free data retrieval call binding the contract method 0xdc5bd997.
//
// Solidity: function userContentRegister() constant returns(address)
func (_PublicationRegister *PublicationRegisterCaller) UserContentRegister(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _PublicationRegister.contract.Call(opts, out, "userContentRegister")
	return *ret0, err
}

// UserContentRegister is a free data retrieval call binding the contract method 0xdc5bd997.
//
// Solidity: function userContentRegister() constant returns(address)
func (_PublicationRegister *PublicationRegisterSession) UserContentRegister() (common.Address, error) {
	return _PublicationRegister.Contract.UserContentRegister(&_PublicationRegister.CallOpts)
}

// UserContentRegister is a free data retrieval call binding the contract method 0xdc5bd997.
//
// Solidity: function userContentRegister() constant returns(address)
func (_PublicationRegister *PublicationRegisterCallerSession) UserContentRegister() (common.Address, error) {
	return _PublicationRegister.Contract.UserContentRegister(&_PublicationRegister.CallOpts)
}

// CreatePublication is a paid mutator transaction binding the contract method 0x2ed24c92.
//
// Solidity: function createPublication(name string, metaData string, minUpVoteCostWei uint256, adminPaymentPercentage uint8) returns()
func (_PublicationRegister *PublicationRegisterTransactor) CreatePublication(opts *bind.TransactOpts, name string, metaData string, minUpVoteCostWei *big.Int, adminPaymentPercentage uint8) (*types.Transaction, error) {
	return _PublicationRegister.contract.Transact(opts, "createPublication", name, metaData, minUpVoteCostWei, adminPaymentPercentage)
}

// CreatePublication is a paid mutator transaction binding the contract method 0x2ed24c92.
//
// Solidity: function createPublication(name string, metaData string, minUpVoteCostWei uint256, adminPaymentPercentage uint8) returns()
func (_PublicationRegister *PublicationRegisterSession) CreatePublication(name string, metaData string, minUpVoteCostWei *big.Int, adminPaymentPercentage uint8) (*types.Transaction, error) {
	return _PublicationRegister.Contract.CreatePublication(&_PublicationRegister.TransactOpts, name, metaData, minUpVoteCostWei, adminPaymentPercentage)
}

// CreatePublication is a paid mutator transaction binding the contract method 0x2ed24c92.
//
// Solidity: function createPublication(name string, metaData string, minUpVoteCostWei uint256, adminPaymentPercentage uint8) returns()
func (_PublicationRegister *PublicationRegisterTransactorSession) CreatePublication(name string, metaData string, minUpVoteCostWei *big.Int, adminPaymentPercentage uint8) (*types.Transaction, error) {
	return _PublicationRegister.Contract.CreatePublication(&_PublicationRegister.TransactOpts, name, metaData, minUpVoteCostWei, adminPaymentPercentage)
}

// DecreaseAdminPaymentPercentage is a paid mutator transaction binding the contract method 0x01065a30.
//
// Solidity: function decreaseAdminPaymentPercentage(whichPublication uint256, newAdminPaymentPercentage uint8) returns()
func (_PublicationRegister *PublicationRegisterTransactor) DecreaseAdminPaymentPercentage(opts *bind.TransactOpts, whichPublication *big.Int, newAdminPaymentPercentage uint8) (*types.Transaction, error) {
	return _PublicationRegister.contract.Transact(opts, "decreaseAdminPaymentPercentage", whichPublication, newAdminPaymentPercentage)
}

// DecreaseAdminPaymentPercentage is a paid mutator transaction binding the contract method 0x01065a30.
//
// Solidity: function decreaseAdminPaymentPercentage(whichPublication uint256, newAdminPaymentPercentage uint8) returns()
func (_PublicationRegister *PublicationRegisterSession) DecreaseAdminPaymentPercentage(whichPublication *big.Int, newAdminPaymentPercentage uint8) (*types.Transaction, error) {
	return _PublicationRegister.Contract.DecreaseAdminPaymentPercentage(&_PublicationRegister.TransactOpts, whichPublication, newAdminPaymentPercentage)
}

// DecreaseAdminPaymentPercentage is a paid mutator transaction binding the contract method 0x01065a30.
//
// Solidity: function decreaseAdminPaymentPercentage(whichPublication uint256, newAdminPaymentPercentage uint8) returns()
func (_PublicationRegister *PublicationRegisterTransactorSession) DecreaseAdminPaymentPercentage(whichPublication *big.Int, newAdminPaymentPercentage uint8) (*types.Transaction, error) {
	return _PublicationRegister.Contract.DecreaseAdminPaymentPercentage(&_PublicationRegister.TransactOpts, whichPublication, newAdminPaymentPercentage)
}

// PermissionAuthor is a paid mutator transaction binding the contract method 0xc9eee140.
//
// Solidity: function permissionAuthor(whichPublication uint256, author address, giveAccess bool) returns()
func (_PublicationRegister *PublicationRegisterTransactor) PermissionAuthor(opts *bind.TransactOpts, whichPublication *big.Int, author common.Address, giveAccess bool) (*types.Transaction, error) {
	return _PublicationRegister.contract.Transact(opts, "permissionAuthor", whichPublication, author, giveAccess)
}

// PermissionAuthor is a paid mutator transaction binding the contract method 0xc9eee140.
//
// Solidity: function permissionAuthor(whichPublication uint256, author address, giveAccess bool) returns()
func (_PublicationRegister *PublicationRegisterSession) PermissionAuthor(whichPublication *big.Int, author common.Address, giveAccess bool) (*types.Transaction, error) {
	return _PublicationRegister.Contract.PermissionAuthor(&_PublicationRegister.TransactOpts, whichPublication, author, giveAccess)
}

// PermissionAuthor is a paid mutator transaction binding the contract method 0xc9eee140.
//
// Solidity: function permissionAuthor(whichPublication uint256, author address, giveAccess bool) returns()
func (_PublicationRegister *PublicationRegisterTransactorSession) PermissionAuthor(whichPublication *big.Int, author common.Address, giveAccess bool) (*types.Transaction, error) {
	return _PublicationRegister.Contract.PermissionAuthor(&_PublicationRegister.TransactOpts, whichPublication, author, giveAccess)
}

// PublishContent is a paid mutator transaction binding the contract method 0xc7071348.
//
// Solidity: function publishContent(whichPublication uint256, index uint256) returns(uint256)
func (_PublicationRegister *PublicationRegisterTransactor) PublishContent(opts *bind.TransactOpts, whichPublication *big.Int, index *big.Int) (*types.Transaction, error) {
	return _PublicationRegister.contract.Transact(opts, "publishContent", whichPublication, index)
}

// PublishContent is a paid mutator transaction binding the contract method 0xc7071348.
//
// Solidity: function publishContent(whichPublication uint256, index uint256) returns(uint256)
func (_PublicationRegister *PublicationRegisterSession) PublishContent(whichPublication *big.Int, index *big.Int) (*types.Transaction, error) {
	return _PublicationRegister.Contract.PublishContent(&_PublicationRegister.TransactOpts, whichPublication, index)
}

// PublishContent is a paid mutator transaction binding the contract method 0xc7071348.
//
// Solidity: function publishContent(whichPublication uint256, index uint256) returns(uint256)
func (_PublicationRegister *PublicationRegisterTransactorSession) PublishContent(whichPublication *big.Int, index *big.Int) (*types.Transaction, error) {
	return _PublicationRegister.Contract.PublishContent(&_PublicationRegister.TransactOpts, whichPublication, index)
}

// SupportPost is a paid mutator transaction binding the contract method 0xd1441d39.
//
// Solidity: function supportPost(whichPublication uint256, postIndex uint256, optionalComment string) returns(bool)
func (_PublicationRegister *PublicationRegisterTransactor) SupportPost(opts *bind.TransactOpts, whichPublication *big.Int, postIndex *big.Int, optionalComment string) (*types.Transaction, error) {
	return _PublicationRegister.contract.Transact(opts, "supportPost", whichPublication, postIndex, optionalComment)
}

// SupportPost is a paid mutator transaction binding the contract method 0xd1441d39.
//
// Solidity: function supportPost(whichPublication uint256, postIndex uint256, optionalComment string) returns(bool)
func (_PublicationRegister *PublicationRegisterSession) SupportPost(whichPublication *big.Int, postIndex *big.Int, optionalComment string) (*types.Transaction, error) {
	return _PublicationRegister.Contract.SupportPost(&_PublicationRegister.TransactOpts, whichPublication, postIndex, optionalComment)
}

// SupportPost is a paid mutator transaction binding the contract method 0xd1441d39.
//
// Solidity: function supportPost(whichPublication uint256, postIndex uint256, optionalComment string) returns(bool)
func (_PublicationRegister *PublicationRegisterTransactorSession) SupportPost(whichPublication *big.Int, postIndex *big.Int, optionalComment string) (*types.Transaction, error) {
	return _PublicationRegister.Contract.SupportPost(&_PublicationRegister.TransactOpts, whichPublication, postIndex, optionalComment)
}

// UpdateMetaData is a paid mutator transaction binding the contract method 0xde189d88.
//
// Solidity: function updateMetaData(whichPublication uint256, metaData string) returns()
func (_PublicationRegister *PublicationRegisterTransactor) UpdateMetaData(opts *bind.TransactOpts, whichPublication *big.Int, metaData string) (*types.Transaction, error) {
	return _PublicationRegister.contract.Transact(opts, "updateMetaData", whichPublication, metaData)
}

// UpdateMetaData is a paid mutator transaction binding the contract method 0xde189d88.
//
// Solidity: function updateMetaData(whichPublication uint256, metaData string) returns()
func (_PublicationRegister *PublicationRegisterSession) UpdateMetaData(whichPublication *big.Int, metaData string) (*types.Transaction, error) {
	return _PublicationRegister.Contract.UpdateMetaData(&_PublicationRegister.TransactOpts, whichPublication, metaData)
}

// UpdateMetaData is a paid mutator transaction binding the contract method 0xde189d88.
//
// Solidity: function updateMetaData(whichPublication uint256, metaData string) returns()
func (_PublicationRegister *PublicationRegisterTransactorSession) UpdateMetaData(whichPublication *big.Int, metaData string) (*types.Transaction, error) {
	return _PublicationRegister.Contract.UpdateMetaData(&_PublicationRegister.TransactOpts, whichPublication, metaData)
}

// WithdrawAdminClaim is a paid mutator transaction binding the contract method 0xcd44ee7f.
//
// Solidity: function withdrawAdminClaim(whichPublication uint256) returns()
func (_PublicationRegister *PublicationRegisterTransactor) WithdrawAdminClaim(opts *bind.TransactOpts, whichPublication *big.Int) (*types.Transaction, error) {
	return _PublicationRegister.contract.Transact(opts, "withdrawAdminClaim", whichPublication)
}

// WithdrawAdminClaim is a paid mutator transaction binding the contract method 0xcd44ee7f.
//
// Solidity: function withdrawAdminClaim(whichPublication uint256) returns()
func (_PublicationRegister *PublicationRegisterSession) WithdrawAdminClaim(whichPublication *big.Int) (*types.Transaction, error) {
	return _PublicationRegister.Contract.WithdrawAdminClaim(&_PublicationRegister.TransactOpts, whichPublication)
}

// WithdrawAdminClaim is a paid mutator transaction binding the contract method 0xcd44ee7f.
//
// Solidity: function withdrawAdminClaim(whichPublication uint256) returns()
func (_PublicationRegister *PublicationRegisterTransactorSession) WithdrawAdminClaim(whichPublication *big.Int) (*types.Transaction, error) {
	return _PublicationRegister.Contract.WithdrawAdminClaim(&_PublicationRegister.TransactOpts, whichPublication)
}

// WithdrawAuthorClaim is a paid mutator transaction binding the contract method 0x6be07202.
//
// Solidity: function withdrawAuthorClaim(whichPublication uint256) returns()
func (_PublicationRegister *PublicationRegisterTransactor) WithdrawAuthorClaim(opts *bind.TransactOpts, whichPublication *big.Int) (*types.Transaction, error) {
	return _PublicationRegister.contract.Transact(opts, "withdrawAuthorClaim", whichPublication)
}

// WithdrawAuthorClaim is a paid mutator transaction binding the contract method 0x6be07202.
//
// Solidity: function withdrawAuthorClaim(whichPublication uint256) returns()
func (_PublicationRegister *PublicationRegisterSession) WithdrawAuthorClaim(whichPublication *big.Int) (*types.Transaction, error) {
	return _PublicationRegister.Contract.WithdrawAuthorClaim(&_PublicationRegister.TransactOpts, whichPublication)
}

// WithdrawAuthorClaim is a paid mutator transaction binding the contract method 0x6be07202.
//
// Solidity: function withdrawAuthorClaim(whichPublication uint256) returns()
func (_PublicationRegister *PublicationRegisterTransactorSession) WithdrawAuthorClaim(whichPublication *big.Int) (*types.Transaction, error) {
	return _PublicationRegister.Contract.WithdrawAuthorClaim(&_PublicationRegister.TransactOpts, whichPublication)
}

// PublicationRegisterStoreDataIterator is returned from FilterStoreData and is used to iterate over the raw logs and unpacked data for StoreData events raised by the PublicationRegister contract.
type PublicationRegisterStoreDataIterator struct {
	Event *PublicationRegisterStoreData // Event containing the contract specifics and raw log

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
func (it *PublicationRegisterStoreDataIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PublicationRegisterStoreData)
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
		it.Event = new(PublicationRegisterStoreData)
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
func (it *PublicationRegisterStoreDataIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PublicationRegisterStoreDataIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PublicationRegisterStoreData represents a StoreData event raised by the PublicationRegister contract.
type PublicationRegisterStoreData struct {
	Data             string
	Author           common.Address
	WhichPublication *big.Int
	Raw              types.Log // Blockchain specific contextual infos
}

// FilterStoreData is a free log retrieval operation binding the contract event 0x5032259159ce8a5de47aa590e04810a4de48d59b6ae1100cc652123768ea9002.
//
// Solidity: e StoreData(data string, author address, whichPublication uint256)
func (_PublicationRegister *PublicationRegisterFilterer) FilterStoreData(opts *bind.FilterOpts) (*PublicationRegisterStoreDataIterator, error) {

	logs, sub, err := _PublicationRegister.contract.FilterLogs(opts, "StoreData")
	if err != nil {
		return nil, err
	}
	return &PublicationRegisterStoreDataIterator{contract: _PublicationRegister.contract, event: "StoreData", logs: logs, sub: sub}, nil
}

// WatchStoreData is a free log subscription operation binding the contract event 0x5032259159ce8a5de47aa590e04810a4de48d59b6ae1100cc652123768ea9002.
//
// Solidity: e StoreData(data string, author address, whichPublication uint256)
func (_PublicationRegister *PublicationRegisterFilterer) WatchStoreData(opts *bind.WatchOpts, sink chan<- *PublicationRegisterStoreData) (event.Subscription, error) {

	logs, sub, err := _PublicationRegister.contract.WatchLogs(opts, "StoreData")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PublicationRegisterStoreData)
				if err := _PublicationRegister.contract.UnpackLog(event, "StoreData", log); err != nil {
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
