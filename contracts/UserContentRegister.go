// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package main

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

// UserContentRegisterABI is the input ABI used to generate the binding from.
const UserContentRegisterABI = "[{\"constant\":true,\"inputs\":[{\"name\":\"whichUser\",\"type\":\"address\"},{\"name\":\"index\",\"type\":\"uint256\"}],\"name\":\"getUserContentBytes\",\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\"},{\"name\":\"\",\"type\":\"bytes32\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"numUsers\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"userName\",\"type\":\"string\"},{\"name\":\"metaData\",\"type\":\"string\"}],\"name\":\"registerNewUser\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"whichUser\",\"type\":\"address\"},{\"name\":\"index\",\"type\":\"uint256\"}],\"name\":\"getUserContent\",\"outputs\":[{\"name\":\"content\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"content\",\"type\":\"string\"}],\"name\":\"publishContent\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"newUsername\",\"type\":\"string\"}],\"name\":\"updateMyUserName\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"userName\",\"type\":\"string\"}],\"name\":\"getUserAddress\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_metaData\",\"type\":\"string\"}],\"name\":\"updateMetaData\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"username\",\"type\":\"string\"}],\"name\":\"checkUserNameTaken\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"contentIndex\",\"type\":\"uint256\"},{\"name\":\"links\",\"type\":\"string\"}],\"name\":\"updateContentLinks\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"whichUser\",\"type\":\"address\"},{\"name\":\"index\",\"type\":\"uint256\"}],\"name\":\"getContentLinks\",\"outputs\":[{\"name\":\"\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"data\",\"type\":\"string\"}],\"name\":\"signalSaveIPFSData\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"address\"}],\"name\":\"registered\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"address\"}],\"name\":\"userIndex\",\"outputs\":[{\"name\":\"userName\",\"type\":\"string\"},{\"name\":\"profileMetaData\",\"type\":\"string\"},{\"name\":\"numContent\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"whichUser\",\"type\":\"address\"}],\"name\":\"getNumContent\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"name\":\"stringBytes32UtilAddress\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"data\",\"type\":\"string\"},{\"indexed\":false,\"name\":\"author\",\"type\":\"address\"}],\"name\":\"StoreData\",\"type\":\"event\"}]"

// UserContentRegister is an auto generated Go binding around an Ethereum contract.
type UserContentRegister struct {
	UserContentRegisterCaller     // Read-only binding to the contract
	UserContentRegisterTransactor // Write-only binding to the contract
	UserContentRegisterFilterer   // Log filterer for contract events
}

// UserContentRegisterCaller is an auto generated read-only Go binding around an Ethereum contract.
type UserContentRegisterCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// UserContentRegisterTransactor is an auto generated write-only Go binding around an Ethereum contract.
type UserContentRegisterTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// UserContentRegisterFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type UserContentRegisterFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// UserContentRegisterSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type UserContentRegisterSession struct {
	Contract     *UserContentRegister // Generic contract binding to set the session for
	CallOpts     bind.CallOpts        // Call options to use throughout this session
	TransactOpts bind.TransactOpts    // Transaction auth options to use throughout this session
}

// UserContentRegisterCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type UserContentRegisterCallerSession struct {
	Contract *UserContentRegisterCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts              // Call options to use throughout this session
}

// UserContentRegisterTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type UserContentRegisterTransactorSession struct {
	Contract     *UserContentRegisterTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts              // Transaction auth options to use throughout this session
}

// UserContentRegisterRaw is an auto generated low-level Go binding around an Ethereum contract.
type UserContentRegisterRaw struct {
	Contract *UserContentRegister // Generic contract binding to access the raw methods on
}

// UserContentRegisterCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type UserContentRegisterCallerRaw struct {
	Contract *UserContentRegisterCaller // Generic read-only contract binding to access the raw methods on
}

// UserContentRegisterTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type UserContentRegisterTransactorRaw struct {
	Contract *UserContentRegisterTransactor // Generic write-only contract binding to access the raw methods on
}

// NewUserContentRegister creates a new instance of UserContentRegister, bound to a specific deployed contract.
func NewUserContentRegister(address common.Address, backend bind.ContractBackend) (*UserContentRegister, error) {
	contract, err := bindUserContentRegister(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &UserContentRegister{UserContentRegisterCaller: UserContentRegisterCaller{contract: contract}, UserContentRegisterTransactor: UserContentRegisterTransactor{contract: contract}, UserContentRegisterFilterer: UserContentRegisterFilterer{contract: contract}}, nil
}

// NewUserContentRegisterCaller creates a new read-only instance of UserContentRegister, bound to a specific deployed contract.
func NewUserContentRegisterCaller(address common.Address, caller bind.ContractCaller) (*UserContentRegisterCaller, error) {
	contract, err := bindUserContentRegister(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &UserContentRegisterCaller{contract: contract}, nil
}

// NewUserContentRegisterTransactor creates a new write-only instance of UserContentRegister, bound to a specific deployed contract.
func NewUserContentRegisterTransactor(address common.Address, transactor bind.ContractTransactor) (*UserContentRegisterTransactor, error) {
	contract, err := bindUserContentRegister(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &UserContentRegisterTransactor{contract: contract}, nil
}

// NewUserContentRegisterFilterer creates a new log filterer instance of UserContentRegister, bound to a specific deployed contract.
func NewUserContentRegisterFilterer(address common.Address, filterer bind.ContractFilterer) (*UserContentRegisterFilterer, error) {
	contract, err := bindUserContentRegister(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &UserContentRegisterFilterer{contract: contract}, nil
}

// bindUserContentRegister binds a generic wrapper to an already deployed contract.
func bindUserContentRegister(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(UserContentRegisterABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_UserContentRegister *UserContentRegisterRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _UserContentRegister.Contract.UserContentRegisterCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_UserContentRegister *UserContentRegisterRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _UserContentRegister.Contract.UserContentRegisterTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_UserContentRegister *UserContentRegisterRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _UserContentRegister.Contract.UserContentRegisterTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_UserContentRegister *UserContentRegisterCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _UserContentRegister.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_UserContentRegister *UserContentRegisterTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _UserContentRegister.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_UserContentRegister *UserContentRegisterTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _UserContentRegister.Contract.contract.Transact(opts, method, params...)
}

// CheckUserNameTaken is a free data retrieval call binding the contract method 0x88f4cdfb.
//
// Solidity: function checkUserNameTaken(username string) constant returns(bool)
func (_UserContentRegister *UserContentRegisterCaller) CheckUserNameTaken(opts *bind.CallOpts, username string) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _UserContentRegister.contract.Call(opts, out, "checkUserNameTaken", username)
	return *ret0, err
}

// CheckUserNameTaken is a free data retrieval call binding the contract method 0x88f4cdfb.
//
// Solidity: function checkUserNameTaken(username string) constant returns(bool)
func (_UserContentRegister *UserContentRegisterSession) CheckUserNameTaken(username string) (bool, error) {
	return _UserContentRegister.Contract.CheckUserNameTaken(&_UserContentRegister.CallOpts, username)
}

// CheckUserNameTaken is a free data retrieval call binding the contract method 0x88f4cdfb.
//
// Solidity: function checkUserNameTaken(username string) constant returns(bool)
func (_UserContentRegister *UserContentRegisterCallerSession) CheckUserNameTaken(username string) (bool, error) {
	return _UserContentRegister.Contract.CheckUserNameTaken(&_UserContentRegister.CallOpts, username)
}

// GetContentLinks is a free data retrieval call binding the contract method 0x965f5aa6.
//
// Solidity: function getContentLinks(whichUser address, index uint256) constant returns(string)
func (_UserContentRegister *UserContentRegisterCaller) GetContentLinks(opts *bind.CallOpts, whichUser common.Address, index *big.Int) (string, error) {
	var (
		ret0 = new(string)
	)
	out := ret0
	err := _UserContentRegister.contract.Call(opts, out, "getContentLinks", whichUser, index)
	return *ret0, err
}

// GetContentLinks is a free data retrieval call binding the contract method 0x965f5aa6.
//
// Solidity: function getContentLinks(whichUser address, index uint256) constant returns(string)
func (_UserContentRegister *UserContentRegisterSession) GetContentLinks(whichUser common.Address, index *big.Int) (string, error) {
	return _UserContentRegister.Contract.GetContentLinks(&_UserContentRegister.CallOpts, whichUser, index)
}

// GetContentLinks is a free data retrieval call binding the contract method 0x965f5aa6.
//
// Solidity: function getContentLinks(whichUser address, index uint256) constant returns(string)
func (_UserContentRegister *UserContentRegisterCallerSession) GetContentLinks(whichUser common.Address, index *big.Int) (string, error) {
	return _UserContentRegister.Contract.GetContentLinks(&_UserContentRegister.CallOpts, whichUser, index)
}

// GetNumContent is a free data retrieval call binding the contract method 0xfa21dc12.
//
// Solidity: function getNumContent(whichUser address) constant returns(uint256)
func (_UserContentRegister *UserContentRegisterCaller) GetNumContent(opts *bind.CallOpts, whichUser common.Address) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _UserContentRegister.contract.Call(opts, out, "getNumContent", whichUser)
	return *ret0, err
}

// GetNumContent is a free data retrieval call binding the contract method 0xfa21dc12.
//
// Solidity: function getNumContent(whichUser address) constant returns(uint256)
func (_UserContentRegister *UserContentRegisterSession) GetNumContent(whichUser common.Address) (*big.Int, error) {
	return _UserContentRegister.Contract.GetNumContent(&_UserContentRegister.CallOpts, whichUser)
}

// GetNumContent is a free data retrieval call binding the contract method 0xfa21dc12.
//
// Solidity: function getNumContent(whichUser address) constant returns(uint256)
func (_UserContentRegister *UserContentRegisterCallerSession) GetNumContent(whichUser common.Address) (*big.Int, error) {
	return _UserContentRegister.Contract.GetNumContent(&_UserContentRegister.CallOpts, whichUser)
}

// GetUserAddress is a free data retrieval call binding the contract method 0x4985e85c.
//
// Solidity: function getUserAddress(userName string) constant returns(address)
func (_UserContentRegister *UserContentRegisterCaller) GetUserAddress(opts *bind.CallOpts, userName string) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _UserContentRegister.contract.Call(opts, out, "getUserAddress", userName)
	return *ret0, err
}

// GetUserAddress is a free data retrieval call binding the contract method 0x4985e85c.
//
// Solidity: function getUserAddress(userName string) constant returns(address)
func (_UserContentRegister *UserContentRegisterSession) GetUserAddress(userName string) (common.Address, error) {
	return _UserContentRegister.Contract.GetUserAddress(&_UserContentRegister.CallOpts, userName)
}

// GetUserAddress is a free data retrieval call binding the contract method 0x4985e85c.
//
// Solidity: function getUserAddress(userName string) constant returns(address)
func (_UserContentRegister *UserContentRegisterCallerSession) GetUserAddress(userName string) (common.Address, error) {
	return _UserContentRegister.Contract.GetUserAddress(&_UserContentRegister.CallOpts, userName)
}

// GetUserContent is a free data retrieval call binding the contract method 0x368ada4c.
//
// Solidity: function getUserContent(whichUser address, index uint256) constant returns(content string)
func (_UserContentRegister *UserContentRegisterCaller) GetUserContent(opts *bind.CallOpts, whichUser common.Address, index *big.Int) (string, error) {
	var (
		ret0 = new(string)
	)
	out := ret0
	err := _UserContentRegister.contract.Call(opts, out, "getUserContent", whichUser, index)
	return *ret0, err
}

// GetUserContent is a free data retrieval call binding the contract method 0x368ada4c.
//
// Solidity: function getUserContent(whichUser address, index uint256) constant returns(content string)
func (_UserContentRegister *UserContentRegisterSession) GetUserContent(whichUser common.Address, index *big.Int) (string, error) {
	return _UserContentRegister.Contract.GetUserContent(&_UserContentRegister.CallOpts, whichUser, index)
}

// GetUserContent is a free data retrieval call binding the contract method 0x368ada4c.
//
// Solidity: function getUserContent(whichUser address, index uint256) constant returns(content string)
func (_UserContentRegister *UserContentRegisterCallerSession) GetUserContent(whichUser common.Address, index *big.Int) (string, error) {
	return _UserContentRegister.Contract.GetUserContent(&_UserContentRegister.CallOpts, whichUser, index)
}

// GetUserContentBytes is a free data retrieval call binding the contract method 0x174a03f3.
//
// Solidity: function getUserContentBytes(whichUser address, index uint256) constant returns(bytes32, bytes32)
func (_UserContentRegister *UserContentRegisterCaller) GetUserContentBytes(opts *bind.CallOpts, whichUser common.Address, index *big.Int) ([32]byte, [32]byte, error) {
	var (
		ret0 = new([32]byte)
		ret1 = new([32]byte)
	)
	out := &[]interface{}{
		ret0,
		ret1,
	}
	err := _UserContentRegister.contract.Call(opts, out, "getUserContentBytes", whichUser, index)
	return *ret0, *ret1, err
}

// GetUserContentBytes is a free data retrieval call binding the contract method 0x174a03f3.
//
// Solidity: function getUserContentBytes(whichUser address, index uint256) constant returns(bytes32, bytes32)
func (_UserContentRegister *UserContentRegisterSession) GetUserContentBytes(whichUser common.Address, index *big.Int) ([32]byte, [32]byte, error) {
	return _UserContentRegister.Contract.GetUserContentBytes(&_UserContentRegister.CallOpts, whichUser, index)
}

// GetUserContentBytes is a free data retrieval call binding the contract method 0x174a03f3.
//
// Solidity: function getUserContentBytes(whichUser address, index uint256) constant returns(bytes32, bytes32)
func (_UserContentRegister *UserContentRegisterCallerSession) GetUserContentBytes(whichUser common.Address, index *big.Int) ([32]byte, [32]byte, error) {
	return _UserContentRegister.Contract.GetUserContentBytes(&_UserContentRegister.CallOpts, whichUser, index)
}

// NumUsers is a free data retrieval call binding the contract method 0x19a50f49.
//
// Solidity: function numUsers() constant returns(uint256)
func (_UserContentRegister *UserContentRegisterCaller) NumUsers(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _UserContentRegister.contract.Call(opts, out, "numUsers")
	return *ret0, err
}

// NumUsers is a free data retrieval call binding the contract method 0x19a50f49.
//
// Solidity: function numUsers() constant returns(uint256)
func (_UserContentRegister *UserContentRegisterSession) NumUsers() (*big.Int, error) {
	return _UserContentRegister.Contract.NumUsers(&_UserContentRegister.CallOpts)
}

// NumUsers is a free data retrieval call binding the contract method 0x19a50f49.
//
// Solidity: function numUsers() constant returns(uint256)
func (_UserContentRegister *UserContentRegisterCallerSession) NumUsers() (*big.Int, error) {
	return _UserContentRegister.Contract.NumUsers(&_UserContentRegister.CallOpts)
}

// Registered is a free data retrieval call binding the contract method 0xb2dd5c07.
//
// Solidity: function registered( address) constant returns(bool)
func (_UserContentRegister *UserContentRegisterCaller) Registered(opts *bind.CallOpts, arg0 common.Address) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _UserContentRegister.contract.Call(opts, out, "registered", arg0)
	return *ret0, err
}

// Registered is a free data retrieval call binding the contract method 0xb2dd5c07.
//
// Solidity: function registered( address) constant returns(bool)
func (_UserContentRegister *UserContentRegisterSession) Registered(arg0 common.Address) (bool, error) {
	return _UserContentRegister.Contract.Registered(&_UserContentRegister.CallOpts, arg0)
}

// Registered is a free data retrieval call binding the contract method 0xb2dd5c07.
//
// Solidity: function registered( address) constant returns(bool)
func (_UserContentRegister *UserContentRegisterCallerSession) Registered(arg0 common.Address) (bool, error) {
	return _UserContentRegister.Contract.Registered(&_UserContentRegister.CallOpts, arg0)
}

// UserIndex is a free data retrieval call binding the contract method 0xc96679fe.
//
// Solidity: function userIndex( address) constant returns(userName string, profileMetaData string, numContent uint256)
func (_UserContentRegister *UserContentRegisterCaller) UserIndex(opts *bind.CallOpts, arg0 common.Address) (struct {
	UserName        string
	ProfileMetaData string
	NumContent      *big.Int
}, error) {
	ret := new(struct {
		UserName        string
		ProfileMetaData string
		NumContent      *big.Int
	})
	out := ret
	err := _UserContentRegister.contract.Call(opts, out, "userIndex", arg0)
	return *ret, err
}

// UserIndex is a free data retrieval call binding the contract method 0xc96679fe.
//
// Solidity: function userIndex( address) constant returns(userName string, profileMetaData string, numContent uint256)
func (_UserContentRegister *UserContentRegisterSession) UserIndex(arg0 common.Address) (struct {
	UserName        string
	ProfileMetaData string
	NumContent      *big.Int
}, error) {
	return _UserContentRegister.Contract.UserIndex(&_UserContentRegister.CallOpts, arg0)
}

// UserIndex is a free data retrieval call binding the contract method 0xc96679fe.
//
// Solidity: function userIndex( address) constant returns(userName string, profileMetaData string, numContent uint256)
func (_UserContentRegister *UserContentRegisterCallerSession) UserIndex(arg0 common.Address) (struct {
	UserName        string
	ProfileMetaData string
	NumContent      *big.Int
}, error) {
	return _UserContentRegister.Contract.UserIndex(&_UserContentRegister.CallOpts, arg0)
}

// PublishContent is a paid mutator transaction binding the contract method 0x3dccafc4.
//
// Solidity: function publishContent(content string) returns()
func (_UserContentRegister *UserContentRegisterTransactor) PublishContent(opts *bind.TransactOpts, content string) (*types.Transaction, error) {
	return _UserContentRegister.contract.Transact(opts, "publishContent", content)
}

// PublishContent is a paid mutator transaction binding the contract method 0x3dccafc4.
//
// Solidity: function publishContent(content string) returns()
func (_UserContentRegister *UserContentRegisterSession) PublishContent(content string) (*types.Transaction, error) {
	return _UserContentRegister.Contract.PublishContent(&_UserContentRegister.TransactOpts, content)
}

// PublishContent is a paid mutator transaction binding the contract method 0x3dccafc4.
//
// Solidity: function publishContent(content string) returns()
func (_UserContentRegister *UserContentRegisterTransactorSession) PublishContent(content string) (*types.Transaction, error) {
	return _UserContentRegister.Contract.PublishContent(&_UserContentRegister.TransactOpts, content)
}

// RegisterNewUser is a paid mutator transaction binding the contract method 0x2588b0c9.
//
// Solidity: function registerNewUser(userName string, metaData string) returns(bool)
func (_UserContentRegister *UserContentRegisterTransactor) RegisterNewUser(opts *bind.TransactOpts, userName string, metaData string) (*types.Transaction, error) {
	return _UserContentRegister.contract.Transact(opts, "registerNewUser", userName, metaData)
}

// RegisterNewUser is a paid mutator transaction binding the contract method 0x2588b0c9.
//
// Solidity: function registerNewUser(userName string, metaData string) returns(bool)
func (_UserContentRegister *UserContentRegisterSession) RegisterNewUser(userName string, metaData string) (*types.Transaction, error) {
	return _UserContentRegister.Contract.RegisterNewUser(&_UserContentRegister.TransactOpts, userName, metaData)
}

// RegisterNewUser is a paid mutator transaction binding the contract method 0x2588b0c9.
//
// Solidity: function registerNewUser(userName string, metaData string) returns(bool)
func (_UserContentRegister *UserContentRegisterTransactorSession) RegisterNewUser(userName string, metaData string) (*types.Transaction, error) {
	return _UserContentRegister.Contract.RegisterNewUser(&_UserContentRegister.TransactOpts, userName, metaData)
}

// SignalSaveIPFSData is a paid mutator transaction binding the contract method 0x9bcdd999.
//
// Solidity: function signalSaveIPFSData(data string) returns()
func (_UserContentRegister *UserContentRegisterTransactor) SignalSaveIPFSData(opts *bind.TransactOpts, data string) (*types.Transaction, error) {
	return _UserContentRegister.contract.Transact(opts, "signalSaveIPFSData", data)
}

// SignalSaveIPFSData is a paid mutator transaction binding the contract method 0x9bcdd999.
//
// Solidity: function signalSaveIPFSData(data string) returns()
func (_UserContentRegister *UserContentRegisterSession) SignalSaveIPFSData(data string) (*types.Transaction, error) {
	return _UserContentRegister.Contract.SignalSaveIPFSData(&_UserContentRegister.TransactOpts, data)
}

// SignalSaveIPFSData is a paid mutator transaction binding the contract method 0x9bcdd999.
//
// Solidity: function signalSaveIPFSData(data string) returns()
func (_UserContentRegister *UserContentRegisterTransactorSession) SignalSaveIPFSData(data string) (*types.Transaction, error) {
	return _UserContentRegister.Contract.SignalSaveIPFSData(&_UserContentRegister.TransactOpts, data)
}

// UpdateContentLinks is a paid mutator transaction binding the contract method 0x8e897be4.
//
// Solidity: function updateContentLinks(contentIndex uint256, links string) returns()
func (_UserContentRegister *UserContentRegisterTransactor) UpdateContentLinks(opts *bind.TransactOpts, contentIndex *big.Int, links string) (*types.Transaction, error) {
	return _UserContentRegister.contract.Transact(opts, "updateContentLinks", contentIndex, links)
}

// UpdateContentLinks is a paid mutator transaction binding the contract method 0x8e897be4.
//
// Solidity: function updateContentLinks(contentIndex uint256, links string) returns()
func (_UserContentRegister *UserContentRegisterSession) UpdateContentLinks(contentIndex *big.Int, links string) (*types.Transaction, error) {
	return _UserContentRegister.Contract.UpdateContentLinks(&_UserContentRegister.TransactOpts, contentIndex, links)
}

// UpdateContentLinks is a paid mutator transaction binding the contract method 0x8e897be4.
//
// Solidity: function updateContentLinks(contentIndex uint256, links string) returns()
func (_UserContentRegister *UserContentRegisterTransactorSession) UpdateContentLinks(contentIndex *big.Int, links string) (*types.Transaction, error) {
	return _UserContentRegister.Contract.UpdateContentLinks(&_UserContentRegister.TransactOpts, contentIndex, links)
}

// UpdateMetaData is a paid mutator transaction binding the contract method 0x5cc07505.
//
// Solidity: function updateMetaData(_metaData string) returns()
func (_UserContentRegister *UserContentRegisterTransactor) UpdateMetaData(opts *bind.TransactOpts, _metaData string) (*types.Transaction, error) {
	return _UserContentRegister.contract.Transact(opts, "updateMetaData", _metaData)
}

// UpdateMetaData is a paid mutator transaction binding the contract method 0x5cc07505.
//
// Solidity: function updateMetaData(_metaData string) returns()
func (_UserContentRegister *UserContentRegisterSession) UpdateMetaData(_metaData string) (*types.Transaction, error) {
	return _UserContentRegister.Contract.UpdateMetaData(&_UserContentRegister.TransactOpts, _metaData)
}

// UpdateMetaData is a paid mutator transaction binding the contract method 0x5cc07505.
//
// Solidity: function updateMetaData(_metaData string) returns()
func (_UserContentRegister *UserContentRegisterTransactorSession) UpdateMetaData(_metaData string) (*types.Transaction, error) {
	return _UserContentRegister.Contract.UpdateMetaData(&_UserContentRegister.TransactOpts, _metaData)
}

// UpdateMyUserName is a paid mutator transaction binding the contract method 0x46ca7371.
//
// Solidity: function updateMyUserName(newUsername string) returns()
func (_UserContentRegister *UserContentRegisterTransactor) UpdateMyUserName(opts *bind.TransactOpts, newUsername string) (*types.Transaction, error) {
	return _UserContentRegister.contract.Transact(opts, "updateMyUserName", newUsername)
}

// UpdateMyUserName is a paid mutator transaction binding the contract method 0x46ca7371.
//
// Solidity: function updateMyUserName(newUsername string) returns()
func (_UserContentRegister *UserContentRegisterSession) UpdateMyUserName(newUsername string) (*types.Transaction, error) {
	return _UserContentRegister.Contract.UpdateMyUserName(&_UserContentRegister.TransactOpts, newUsername)
}

// UpdateMyUserName is a paid mutator transaction binding the contract method 0x46ca7371.
//
// Solidity: function updateMyUserName(newUsername string) returns()
func (_UserContentRegister *UserContentRegisterTransactorSession) UpdateMyUserName(newUsername string) (*types.Transaction, error) {
	return _UserContentRegister.Contract.UpdateMyUserName(&_UserContentRegister.TransactOpts, newUsername)
}

// UserContentRegisterStoreDataIterator is returned from FilterStoreData and is used to iterate over the raw logs and unpacked data for StoreData events raised by the UserContentRegister contract.
type UserContentRegisterStoreDataIterator struct {
	Event *UserContentRegisterStoreData // Event containing the contract specifics and raw log

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
func (it *UserContentRegisterStoreDataIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(UserContentRegisterStoreData)
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
		it.Event = new(UserContentRegisterStoreData)
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
func (it *UserContentRegisterStoreDataIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *UserContentRegisterStoreDataIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// UserContentRegisterStoreData represents a StoreData event raised by the UserContentRegister contract.
type UserContentRegisterStoreData struct {
	Data   string
	Author common.Address
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterStoreData is a free log retrieval operation binding the contract event 0x85a3aedb44223db6881824597296dbd3f7dea41c2528c1795b04794bfc00a130.
//
// Solidity: e StoreData(data string, author address)
func (_UserContentRegister *UserContentRegisterFilterer) FilterStoreData(opts *bind.FilterOpts) (*UserContentRegisterStoreDataIterator, error) {

	logs, sub, err := _UserContentRegister.contract.FilterLogs(opts, "StoreData")
	if err != nil {
		return nil, err
	}
	return &UserContentRegisterStoreDataIterator{contract: _UserContentRegister.contract, event: "StoreData", logs: logs, sub: sub}, nil
}

// WatchStoreData is a free log subscription operation binding the contract event 0x85a3aedb44223db6881824597296dbd3f7dea41c2528c1795b04794bfc00a130.
//
// Solidity: e StoreData(data string, author address)
func (_UserContentRegister *UserContentRegisterFilterer) WatchStoreData(opts *bind.WatchOpts, sink chan<- *UserContentRegisterStoreData) (event.Subscription, error) {

	logs, sub, err := _UserContentRegister.contract.WatchLogs(opts, "StoreData")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(UserContentRegisterStoreData)
				if err := _UserContentRegister.contract.UnpackLog(event, "StoreData", log); err != nil {
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
