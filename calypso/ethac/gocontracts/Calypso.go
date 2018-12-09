// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package gocontracts

import (
	"math/big"
	"strings"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
)

// CalypsoABI is the input ABI used to generate the binding from.
const CalypsoABI = "[{\"constant\":false,\"inputs\":[{\"name\":\"a\",\"type\":\"address\"}],\"name\":\"addWriteRequest\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"address\"}],\"name\":\"ownersMap\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"a\",\"type\":\"address\"}],\"name\":\"canWrite\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"a\",\"type\":\"address\"}],\"name\":\"checkIfRequestIsValid\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"a\",\"type\":\"address\"}],\"name\":\"isValidReadRequest\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"a\",\"type\":\"address\"}],\"name\":\"AddOwner\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"a\",\"type\":\"address\"}],\"name\":\"AddReadRequest\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"name\":\"holder\",\"type\":\"address\"},{\"name\":\"rrholder\",\"type\":\"address\"},{\"name\":\"o\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"}]"

// CalypsoBin is the compiled bytecode used for deploying new contracts.
const CalypsoBin = `0x608060405234801561001057600080fd5b506040516060806109118339810180604052606081101561003057600080fd5b508051602082015160409092015160008054600160a060020a03928316600160a060020a0319918216179091556001805493831693821693909317909255600380549190931691161790556108878061008a6000396000f3fe6080604052600436106100635760e060020a60003504631029a017811461006857806319a444b71461009d5780633239b71b146100e45780634a4e19d314610117578063778ad09c1461014a578063ac1e9ef41461017d578063bdae9b89146101b0575b600080fd5b34801561007457600080fd5b5061009b6004803603602081101561008b57600080fd5b5035600160a060020a03166101e3565b005b3480156100a957600080fd5b506100d0600480360360208110156100c057600080fd5b5035600160a060020a03166102db565b604080519115158252519081900360200190f35b3480156100f057600080fd5b506100d06004803603602081101561010757600080fd5b5035600160a060020a03166102f0565b34801561012357600080fd5b506100d06004803603602081101561013a57600080fd5b5035600160a060020a0316610389565b34801561015657600080fd5b506100d06004803603602081101561016d57600080fd5b5035600160a060020a0316610408565b34801561018957600080fd5b5061009b600480360360208110156101a057600080fd5b5035600160a060020a031661061b565b3480156101bc57600080fd5b5061009b600480360360208110156101d357600080fd5b5035600160a060020a031661063f565b6101ec336102f0565b151561025957604080517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601e60248201527f596f7520617265206e6f7420612072656769737465726564206f776e65720000604482015290519081900360640190fd5b600154604080517f1029a017000000000000000000000000000000000000000000000000000000008152600160a060020a03848116600483015291519190921691631029a01791602480830192600092919082900301818387803b1580156102c057600080fd5b505af11580156102d4573d6000803e3d6000fd5b5050505050565b60046020526000908152604090205460ff1681565b60008054604080517f3239b71b000000000000000000000000000000000000000000000000000000008152600160a060020a03858116600483015291519190921691633239b71b916024808301926020929190829003018186803b15801561035757600080fd5b505afa15801561036b573d6000803e3d6000fd5b505050506040513d602081101561038157600080fd5b505192915050565b600154604080517f42087d4f000000000000000000000000000000000000000000000000000000008152600160a060020a038481166004830152915160009392909216916342087d4f9160248082019260209290919082900301818787803b1580156103f457600080fd5b505af115801561036b573d6000803e3d6000fd5b600080829050600081600160a060020a03166364e94e676040518163ffffffff1660e060020a02815260040160206040518083038186803b15801561044c57600080fd5b505afa158015610460573d6000803e3d6000fd5b505050506040513d602081101561047657600080fd5b5051604080517f8da5cb5b00000000000000000000000000000000000000000000000000000000815290519192508291600160a060020a038084169263565542569291871691638da5cb5b91600480820192602092909190829003018186803b1580156104e257600080fd5b505afa1580156104f6573d6000803e3d6000fd5b505050506040513d602081101561050c57600080fd5b50516040805160e060020a63ffffffff8516028152600160a060020a039092166004830152516024808301926020929190829003018186803b15801561055157600080fd5b505afa158015610565573d6000803e3d6000fd5b505050506040513d602081101561057b57600080fd5b5051151561061057604080517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152603a60248201527f497320746869732061206d656d626572206f662074686520706f6c696379206f60448201527f662074686520706f696e74656420577269746552657175657374000000000000606482015290519081900360840190fd5b506001949350505050565b600160a060020a03166000908152600460205260409020805460ff19166001179055565b6000819050600081600160a060020a03166364e94e676040518163ffffffff1660e060020a02815260040160206040518083038186803b15801561068257600080fd5b505afa158015610696573d6000803e3d6000fd5b505050506040513d60208110156106ac57600080fd5b5051600154604080517f42087d4f000000000000000000000000000000000000000000000000000000008152600160a060020a03808516600483015291519394509116916342087d4f916024808201926020929091908290030181600087803b15801561071857600080fd5b505af115801561072c573d6000803e3d6000fd5b505050506040513d602081101561074257600080fd5b505115156107d757604080517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602760248201527f5468657265206973206e6f20636f72726573706f6e64696e672077726974652060448201527f7265717565737400000000000000000000000000000000000000000000000000606482015290519081900360840190fd5b600354604080517f31419b50000000000000000000000000000000000000000000000000000000008152600160a060020a038681166004830152915191909216916331419b5091602480830192600092919082900301818387803b15801561083e57600080fd5b505af1158015610852573d6000803e3d6000fd5b5050505050505056fea165627a7a723058209feeef35d18f248ba7d654502295ea7dd9d36c456cb88af3ba8303995600c3cf0029`

// DeployCalypso deploys a new Ethereum contract, binding an instance of Calypso to it.
func DeployCalypso(auth *bind.TransactOpts, backend bind.ContractBackend, holder common.Address, rrholder common.Address, o common.Address) (common.Address, *types.Transaction, *Calypso, error) {
	parsed, err := abi.JSON(strings.NewReader(CalypsoABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(CalypsoBin), backend, holder, rrholder, o)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Calypso{CalypsoCaller: CalypsoCaller{contract: contract}, CalypsoTransactor: CalypsoTransactor{contract: contract}, CalypsoFilterer: CalypsoFilterer{contract: contract}}, nil
}

// Calypso is an auto generated Go binding around an Ethereum contract.
type Calypso struct {
	CalypsoCaller     // Read-only binding to the contract
	CalypsoTransactor // Write-only binding to the contract
	CalypsoFilterer   // Log filterer for contract events
}

// CalypsoCaller is an auto generated read-only Go binding around an Ethereum contract.
type CalypsoCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// CalypsoTransactor is an auto generated write-only Go binding around an Ethereum contract.
type CalypsoTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// CalypsoFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type CalypsoFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// CalypsoSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type CalypsoSession struct {
	Contract     *Calypso          // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// CalypsoCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type CalypsoCallerSession struct {
	Contract *CalypsoCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts  // Call options to use throughout this session
}

// CalypsoTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type CalypsoTransactorSession struct {
	Contract     *CalypsoTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts  // Transaction auth options to use throughout this session
}

// CalypsoRaw is an auto generated low-level Go binding around an Ethereum contract.
type CalypsoRaw struct {
	Contract *Calypso // Generic contract binding to access the raw methods on
}

// CalypsoCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type CalypsoCallerRaw struct {
	Contract *CalypsoCaller // Generic read-only contract binding to access the raw methods on
}

// CalypsoTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type CalypsoTransactorRaw struct {
	Contract *CalypsoTransactor // Generic write-only contract binding to access the raw methods on
}

// NewCalypso creates a new instance of Calypso, bound to a specific deployed contract.
func NewCalypso(address common.Address, backend bind.ContractBackend) (*Calypso, error) {
	contract, err := bindCalypso(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Calypso{CalypsoCaller: CalypsoCaller{contract: contract}, CalypsoTransactor: CalypsoTransactor{contract: contract}, CalypsoFilterer: CalypsoFilterer{contract: contract}}, nil
}

// NewCalypsoCaller creates a new read-only instance of Calypso, bound to a specific deployed contract.
func NewCalypsoCaller(address common.Address, caller bind.ContractCaller) (*CalypsoCaller, error) {
	contract, err := bindCalypso(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &CalypsoCaller{contract: contract}, nil
}

// NewCalypsoTransactor creates a new write-only instance of Calypso, bound to a specific deployed contract.
func NewCalypsoTransactor(address common.Address, transactor bind.ContractTransactor) (*CalypsoTransactor, error) {
	contract, err := bindCalypso(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &CalypsoTransactor{contract: contract}, nil
}

// NewCalypsoFilterer creates a new log filterer instance of Calypso, bound to a specific deployed contract.
func NewCalypsoFilterer(address common.Address, filterer bind.ContractFilterer) (*CalypsoFilterer, error) {
	contract, err := bindCalypso(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &CalypsoFilterer{contract: contract}, nil
}

// bindCalypso binds a generic wrapper to an already deployed contract.
func bindCalypso(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(CalypsoABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Calypso *CalypsoRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Calypso.Contract.CalypsoCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Calypso *CalypsoRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Calypso.Contract.CalypsoTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Calypso *CalypsoRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Calypso.Contract.CalypsoTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Calypso *CalypsoCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Calypso.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Calypso *CalypsoTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Calypso.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Calypso *CalypsoTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Calypso.Contract.contract.Transact(opts, method, params...)
}

// CanWrite is a free data retrieval call binding the contract method 0x3239b71b.
//
// Solidity: function canWrite(a address) constant returns(bool)
func (_Calypso *CalypsoCaller) CanWrite(opts *bind.CallOpts, a common.Address) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _Calypso.contract.Call(opts, out, "canWrite", a)
	return *ret0, err
}

// CanWrite is a free data retrieval call binding the contract method 0x3239b71b.
//
// Solidity: function canWrite(a address) constant returns(bool)
func (_Calypso *CalypsoSession) CanWrite(a common.Address) (bool, error) {
	return _Calypso.Contract.CanWrite(&_Calypso.CallOpts, a)
}

// CanWrite is a free data retrieval call binding the contract method 0x3239b71b.
//
// Solidity: function canWrite(a address) constant returns(bool)
func (_Calypso *CalypsoCallerSession) CanWrite(a common.Address) (bool, error) {
	return _Calypso.Contract.CanWrite(&_Calypso.CallOpts, a)
}

// IsValidReadRequest is a free data retrieval call binding the contract method 0x778ad09c.
//
// Solidity: function isValidReadRequest(a address) constant returns(bool)
func (_Calypso *CalypsoCaller) IsValidReadRequest(opts *bind.CallOpts, a common.Address) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _Calypso.contract.Call(opts, out, "isValidReadRequest", a)
	return *ret0, err
}

// IsValidReadRequest is a free data retrieval call binding the contract method 0x778ad09c.
//
// Solidity: function isValidReadRequest(a address) constant returns(bool)
func (_Calypso *CalypsoSession) IsValidReadRequest(a common.Address) (bool, error) {
	return _Calypso.Contract.IsValidReadRequest(&_Calypso.CallOpts, a)
}

// IsValidReadRequest is a free data retrieval call binding the contract method 0x778ad09c.
//
// Solidity: function isValidReadRequest(a address) constant returns(bool)
func (_Calypso *CalypsoCallerSession) IsValidReadRequest(a common.Address) (bool, error) {
	return _Calypso.Contract.IsValidReadRequest(&_Calypso.CallOpts, a)
}

// OwnersMap is a free data retrieval call binding the contract method 0x19a444b7.
//
// Solidity: function ownersMap( address) constant returns(bool)
func (_Calypso *CalypsoCaller) OwnersMap(opts *bind.CallOpts, arg0 common.Address) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _Calypso.contract.Call(opts, out, "ownersMap", arg0)
	return *ret0, err
}

// OwnersMap is a free data retrieval call binding the contract method 0x19a444b7.
//
// Solidity: function ownersMap( address) constant returns(bool)
func (_Calypso *CalypsoSession) OwnersMap(arg0 common.Address) (bool, error) {
	return _Calypso.Contract.OwnersMap(&_Calypso.CallOpts, arg0)
}

// OwnersMap is a free data retrieval call binding the contract method 0x19a444b7.
//
// Solidity: function ownersMap( address) constant returns(bool)
func (_Calypso *CalypsoCallerSession) OwnersMap(arg0 common.Address) (bool, error) {
	return _Calypso.Contract.OwnersMap(&_Calypso.CallOpts, arg0)
}

// AddOwner is a paid mutator transaction binding the contract method 0xac1e9ef4.
//
// Solidity: function AddOwner(a address) returns()
func (_Calypso *CalypsoTransactor) AddOwner(opts *bind.TransactOpts, a common.Address) (*types.Transaction, error) {
	return _Calypso.contract.Transact(opts, "AddOwner", a)
}

// AddOwner is a paid mutator transaction binding the contract method 0xac1e9ef4.
//
// Solidity: function AddOwner(a address) returns()
func (_Calypso *CalypsoSession) AddOwner(a common.Address) (*types.Transaction, error) {
	return _Calypso.Contract.AddOwner(&_Calypso.TransactOpts, a)
}

// AddOwner is a paid mutator transaction binding the contract method 0xac1e9ef4.
//
// Solidity: function AddOwner(a address) returns()
func (_Calypso *CalypsoTransactorSession) AddOwner(a common.Address) (*types.Transaction, error) {
	return _Calypso.Contract.AddOwner(&_Calypso.TransactOpts, a)
}

// AddReadRequest is a paid mutator transaction binding the contract method 0xbdae9b89.
//
// Solidity: function AddReadRequest(a address) returns()
func (_Calypso *CalypsoTransactor) AddReadRequest(opts *bind.TransactOpts, a common.Address) (*types.Transaction, error) {
	return _Calypso.contract.Transact(opts, "AddReadRequest", a)
}

// AddReadRequest is a paid mutator transaction binding the contract method 0xbdae9b89.
//
// Solidity: function AddReadRequest(a address) returns()
func (_Calypso *CalypsoSession) AddReadRequest(a common.Address) (*types.Transaction, error) {
	return _Calypso.Contract.AddReadRequest(&_Calypso.TransactOpts, a)
}

// AddReadRequest is a paid mutator transaction binding the contract method 0xbdae9b89.
//
// Solidity: function AddReadRequest(a address) returns()
func (_Calypso *CalypsoTransactorSession) AddReadRequest(a common.Address) (*types.Transaction, error) {
	return _Calypso.Contract.AddReadRequest(&_Calypso.TransactOpts, a)
}

// AddWriteRequest is a paid mutator transaction binding the contract method 0x1029a017.
//
// Solidity: function addWriteRequest(a address) returns()
func (_Calypso *CalypsoTransactor) AddWriteRequest(opts *bind.TransactOpts, a common.Address) (*types.Transaction, error) {
	return _Calypso.contract.Transact(opts, "addWriteRequest", a)
}

// AddWriteRequest is a paid mutator transaction binding the contract method 0x1029a017.
//
// Solidity: function addWriteRequest(a address) returns()
func (_Calypso *CalypsoSession) AddWriteRequest(a common.Address) (*types.Transaction, error) {
	return _Calypso.Contract.AddWriteRequest(&_Calypso.TransactOpts, a)
}

// AddWriteRequest is a paid mutator transaction binding the contract method 0x1029a017.
//
// Solidity: function addWriteRequest(a address) returns()
func (_Calypso *CalypsoTransactorSession) AddWriteRequest(a common.Address) (*types.Transaction, error) {
	return _Calypso.Contract.AddWriteRequest(&_Calypso.TransactOpts, a)
}

// CheckIfRequestIsValid is a paid mutator transaction binding the contract method 0x4a4e19d3.
//
// Solidity: function checkIfRequestIsValid(a address) returns(bool)
func (_Calypso *CalypsoTransactor) CheckIfRequestIsValid(opts *bind.TransactOpts, a common.Address) (*types.Transaction, error) {
	return _Calypso.contract.Transact(opts, "checkIfRequestIsValid", a)
}

// CheckIfRequestIsValid is a paid mutator transaction binding the contract method 0x4a4e19d3.
//
// Solidity: function checkIfRequestIsValid(a address) returns(bool)
func (_Calypso *CalypsoSession) CheckIfRequestIsValid(a common.Address) (*types.Transaction, error) {
	return _Calypso.Contract.CheckIfRequestIsValid(&_Calypso.TransactOpts, a)
}

// CheckIfRequestIsValid is a paid mutator transaction binding the contract method 0x4a4e19d3.
//
// Solidity: function checkIfRequestIsValid(a address) returns(bool)
func (_Calypso *CalypsoTransactorSession) CheckIfRequestIsValid(a common.Address) (*types.Transaction, error) {
	return _Calypso.Contract.CheckIfRequestIsValid(&_Calypso.TransactOpts, a)
}

// OwnersABI is the input ABI used to generate the binding from.
const OwnersABI = "[{\"constant\":true,\"inputs\":[{\"name\":\"a\",\"type\":\"address\"}],\"name\":\"canWrite\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getNrOwners\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"a\",\"type\":\"address\"}],\"name\":\"AddOwner\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"name\":\"a\",\"type\":\"address[]\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"}]"

// OwnersBin is the compiled bytecode used for deploying new contracts.
const OwnersBin = `0x608060405234801561001057600080fd5b506040516103403803806103408339810180604052602081101561003357600080fd5b81019080805164010000000081111561004b57600080fd5b8201602081018481111561005e57600080fd5b815185602082028301116401000000008211171561007b57600080fd5b505080519093506100959250600091506020840190610102565b5060005b81518163ffffffff1610156100fb576001806000848463ffffffff168151811015156100c157fe5b602090810291909101810151600160a060020a03168252810191909152604001600020805460ff1916911515919091179055600101610099565b505061018e565b828054828255906000526020600020908101928215610157579160200282015b828111156101575782518254600160a060020a031916600160a060020a03909116178255602090920191600190910190610122565b50610163929150610167565b5090565b61018b91905b80821115610163578054600160a060020a031916815560010161016d565b90565b6101a38061019d6000396000f3fe608060405260043610610050577c010000000000000000000000000000000000000000000000000000000060003504633239b71b8114610055578063502bb734146100a9578063ac1e9ef4146100d0575b600080fd5b34801561006157600080fd5b506100956004803603602081101561007857600080fd5b503573ffffffffffffffffffffffffffffffffffffffff16610112565b604080519115158252519081900360200190f35b3480156100b557600080fd5b506100be61013d565b60408051918252519081900360200190f35b3480156100dc57600080fd5b50610110600480360360208110156100f357600080fd5b503573ffffffffffffffffffffffffffffffffffffffff16610143565b005b73ffffffffffffffffffffffffffffffffffffffff1660009081526001602052604090205460ff1690565b60005490565b73ffffffffffffffffffffffffffffffffffffffff166000908152600160208190526040909120805460ff1916909117905556fea165627a7a72305820eee30b3e680a8f66e353c7e5262835e81b1092377f5a211066bf33b88fef9c060029`

// DeployOwners deploys a new Ethereum contract, binding an instance of Owners to it.
func DeployOwners(auth *bind.TransactOpts, backend bind.ContractBackend, a []common.Address) (common.Address, *types.Transaction, *Owners, error) {
	parsed, err := abi.JSON(strings.NewReader(OwnersABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(OwnersBin), backend, a)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Owners{OwnersCaller: OwnersCaller{contract: contract}, OwnersTransactor: OwnersTransactor{contract: contract}, OwnersFilterer: OwnersFilterer{contract: contract}}, nil
}

// Owners is an auto generated Go binding around an Ethereum contract.
type Owners struct {
	OwnersCaller     // Read-only binding to the contract
	OwnersTransactor // Write-only binding to the contract
	OwnersFilterer   // Log filterer for contract events
}

// OwnersCaller is an auto generated read-only Go binding around an Ethereum contract.
type OwnersCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// OwnersTransactor is an auto generated write-only Go binding around an Ethereum contract.
type OwnersTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// OwnersFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type OwnersFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// OwnersSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type OwnersSession struct {
	Contract     *Owners           // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// OwnersCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type OwnersCallerSession struct {
	Contract *OwnersCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// OwnersTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type OwnersTransactorSession struct {
	Contract     *OwnersTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// OwnersRaw is an auto generated low-level Go binding around an Ethereum contract.
type OwnersRaw struct {
	Contract *Owners // Generic contract binding to access the raw methods on
}

// OwnersCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type OwnersCallerRaw struct {
	Contract *OwnersCaller // Generic read-only contract binding to access the raw methods on
}

// OwnersTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type OwnersTransactorRaw struct {
	Contract *OwnersTransactor // Generic write-only contract binding to access the raw methods on
}

// NewOwners creates a new instance of Owners, bound to a specific deployed contract.
func NewOwners(address common.Address, backend bind.ContractBackend) (*Owners, error) {
	contract, err := bindOwners(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Owners{OwnersCaller: OwnersCaller{contract: contract}, OwnersTransactor: OwnersTransactor{contract: contract}, OwnersFilterer: OwnersFilterer{contract: contract}}, nil
}

// NewOwnersCaller creates a new read-only instance of Owners, bound to a specific deployed contract.
func NewOwnersCaller(address common.Address, caller bind.ContractCaller) (*OwnersCaller, error) {
	contract, err := bindOwners(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &OwnersCaller{contract: contract}, nil
}

// NewOwnersTransactor creates a new write-only instance of Owners, bound to a specific deployed contract.
func NewOwnersTransactor(address common.Address, transactor bind.ContractTransactor) (*OwnersTransactor, error) {
	contract, err := bindOwners(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &OwnersTransactor{contract: contract}, nil
}

// NewOwnersFilterer creates a new log filterer instance of Owners, bound to a specific deployed contract.
func NewOwnersFilterer(address common.Address, filterer bind.ContractFilterer) (*OwnersFilterer, error) {
	contract, err := bindOwners(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &OwnersFilterer{contract: contract}, nil
}

// bindOwners binds a generic wrapper to an already deployed contract.
func bindOwners(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(OwnersABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Owners *OwnersRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Owners.Contract.OwnersCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Owners *OwnersRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Owners.Contract.OwnersTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Owners *OwnersRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Owners.Contract.OwnersTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Owners *OwnersCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Owners.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Owners *OwnersTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Owners.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Owners *OwnersTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Owners.Contract.contract.Transact(opts, method, params...)
}

// CanWrite is a free data retrieval call binding the contract method 0x3239b71b.
//
// Solidity: function canWrite(a address) constant returns(bool)
func (_Owners *OwnersCaller) CanWrite(opts *bind.CallOpts, a common.Address) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _Owners.contract.Call(opts, out, "canWrite", a)
	return *ret0, err
}

// CanWrite is a free data retrieval call binding the contract method 0x3239b71b.
//
// Solidity: function canWrite(a address) constant returns(bool)
func (_Owners *OwnersSession) CanWrite(a common.Address) (bool, error) {
	return _Owners.Contract.CanWrite(&_Owners.CallOpts, a)
}

// CanWrite is a free data retrieval call binding the contract method 0x3239b71b.
//
// Solidity: function canWrite(a address) constant returns(bool)
func (_Owners *OwnersCallerSession) CanWrite(a common.Address) (bool, error) {
	return _Owners.Contract.CanWrite(&_Owners.CallOpts, a)
}

// GetNrOwners is a free data retrieval call binding the contract method 0x502bb734.
//
// Solidity: function getNrOwners() constant returns(uint256)
func (_Owners *OwnersCaller) GetNrOwners(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Owners.contract.Call(opts, out, "getNrOwners")
	return *ret0, err
}

// GetNrOwners is a free data retrieval call binding the contract method 0x502bb734.
//
// Solidity: function getNrOwners() constant returns(uint256)
func (_Owners *OwnersSession) GetNrOwners() (*big.Int, error) {
	return _Owners.Contract.GetNrOwners(&_Owners.CallOpts)
}

// GetNrOwners is a free data retrieval call binding the contract method 0x502bb734.
//
// Solidity: function getNrOwners() constant returns(uint256)
func (_Owners *OwnersCallerSession) GetNrOwners() (*big.Int, error) {
	return _Owners.Contract.GetNrOwners(&_Owners.CallOpts)
}

// AddOwner is a paid mutator transaction binding the contract method 0xac1e9ef4.
//
// Solidity: function AddOwner(a address) returns()
func (_Owners *OwnersTransactor) AddOwner(opts *bind.TransactOpts, a common.Address) (*types.Transaction, error) {
	return _Owners.contract.Transact(opts, "AddOwner", a)
}

// AddOwner is a paid mutator transaction binding the contract method 0xac1e9ef4.
//
// Solidity: function AddOwner(a address) returns()
func (_Owners *OwnersSession) AddOwner(a common.Address) (*types.Transaction, error) {
	return _Owners.Contract.AddOwner(&_Owners.TransactOpts, a)
}

// AddOwner is a paid mutator transaction binding the contract method 0xac1e9ef4.
//
// Solidity: function AddOwner(a address) returns()
func (_Owners *OwnersTransactorSession) AddOwner(a common.Address) (*types.Transaction, error) {
	return _Owners.Contract.AddOwner(&_Owners.TransactOpts, a)
}

// PolicyABI is the input ABI used to generate the binding from.
const PolicyABI = "[{\"constant\":false,\"inputs\":[{\"name\":\"a\",\"type\":\"address\"}],\"name\":\"RemoveFromPolicy\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"a\",\"type\":\"address\"}],\"name\":\"CanRead\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"a\",\"type\":\"address\"}],\"name\":\"AddOwner\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"name\":\"a\",\"type\":\"address[]\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"}]"

// PolicyBin is the compiled bytecode used for deploying new contracts.
const PolicyBin = `0x608060405234801561001057600080fd5b506040516104873803806104878339810180604052602081101561003357600080fd5b81019080805164010000000081111561004b57600080fd5b8201602081018481111561005e57600080fd5b815185602082028301116401000000008211171561007b57600080fd5b505060008054600160a060020a0319163317905580519093506100a7925060019150602084019061011b565b5060005b8151816001604060020a031610156101145760016002600084846001604060020a03168151811015156100da57fe5b602090810291909101810151600160a060020a03168252810191909152604001600020805460ff19169115159190911790556001016100ab565b50506101a7565b828054828255906000526020600020908101928215610170579160200282015b828111156101705782518254600160a060020a031916600160a060020a0390911617825560209092019160019091019061013b565b5061017c929150610180565b5090565b6101a491905b8082111561017c578054600160a060020a0319168155600101610186565b90565b6102d1806101b66000396000f3fe608060405260043610610050577c010000000000000000000000000000000000000000000000000000000060003504631809a39e8114610055578063565542561461008a578063ac1e9ef4146100d1575b600080fd5b34801561006157600080fd5b506100886004803603602081101561007857600080fd5b5035600160a060020a0316610104565b005b34801561009657600080fd5b506100bd600480360360208110156100ad57600080fd5b5035600160a060020a03166101c4565b604080519115158252519081900360200190f35b3480156100dd57600080fd5b50610088600480360360208110156100f457600080fd5b5035600160a060020a03166101e2565b600054600160a060020a031633146101a357604080517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602f60248201527f596f7520617265206e6f742070617274206f6620746865206f776e657273686960448201527f70206f66207468697320706f6963790000000000000000000000000000000000606482015290519081900360840190fd5b600160a060020a03166000908152600260205260409020805460ff19169055565b600160a060020a031660009081526002602052604090205460ff1690565b600054600160a060020a0316331461028157604080517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602f60248201527f596f7520617265206e6f742070617274206f6620746865206f776e657273686960448201527f70206f66207468697320706f6963790000000000000000000000000000000000606482015290519081900360840190fd5b600160a060020a03166000908152600260205260409020805460ff1916600117905556fea165627a7a72305820a0ed47e9eab40eddacde74f4eadba88672ab6b210f5529deae653ee9b9efa25e0029`

// DeployPolicy deploys a new Ethereum contract, binding an instance of Policy to it.
func DeployPolicy(auth *bind.TransactOpts, backend bind.ContractBackend, a []common.Address) (common.Address, *types.Transaction, *Policy, error) {
	parsed, err := abi.JSON(strings.NewReader(PolicyABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(PolicyBin), backend, a)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Policy{PolicyCaller: PolicyCaller{contract: contract}, PolicyTransactor: PolicyTransactor{contract: contract}, PolicyFilterer: PolicyFilterer{contract: contract}}, nil
}

// Policy is an auto generated Go binding around an Ethereum contract.
type Policy struct {
	PolicyCaller     // Read-only binding to the contract
	PolicyTransactor // Write-only binding to the contract
	PolicyFilterer   // Log filterer for contract events
}

// PolicyCaller is an auto generated read-only Go binding around an Ethereum contract.
type PolicyCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PolicyTransactor is an auto generated write-only Go binding around an Ethereum contract.
type PolicyTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PolicyFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type PolicyFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PolicySession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type PolicySession struct {
	Contract     *Policy           // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// PolicyCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type PolicyCallerSession struct {
	Contract *PolicyCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// PolicyTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type PolicyTransactorSession struct {
	Contract     *PolicyTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// PolicyRaw is an auto generated low-level Go binding around an Ethereum contract.
type PolicyRaw struct {
	Contract *Policy // Generic contract binding to access the raw methods on
}

// PolicyCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type PolicyCallerRaw struct {
	Contract *PolicyCaller // Generic read-only contract binding to access the raw methods on
}

// PolicyTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type PolicyTransactorRaw struct {
	Contract *PolicyTransactor // Generic write-only contract binding to access the raw methods on
}

// NewPolicy creates a new instance of Policy, bound to a specific deployed contract.
func NewPolicy(address common.Address, backend bind.ContractBackend) (*Policy, error) {
	contract, err := bindPolicy(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Policy{PolicyCaller: PolicyCaller{contract: contract}, PolicyTransactor: PolicyTransactor{contract: contract}, PolicyFilterer: PolicyFilterer{contract: contract}}, nil
}

// NewPolicyCaller creates a new read-only instance of Policy, bound to a specific deployed contract.
func NewPolicyCaller(address common.Address, caller bind.ContractCaller) (*PolicyCaller, error) {
	contract, err := bindPolicy(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &PolicyCaller{contract: contract}, nil
}

// NewPolicyTransactor creates a new write-only instance of Policy, bound to a specific deployed contract.
func NewPolicyTransactor(address common.Address, transactor bind.ContractTransactor) (*PolicyTransactor, error) {
	contract, err := bindPolicy(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &PolicyTransactor{contract: contract}, nil
}

// NewPolicyFilterer creates a new log filterer instance of Policy, bound to a specific deployed contract.
func NewPolicyFilterer(address common.Address, filterer bind.ContractFilterer) (*PolicyFilterer, error) {
	contract, err := bindPolicy(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &PolicyFilterer{contract: contract}, nil
}

// bindPolicy binds a generic wrapper to an already deployed contract.
func bindPolicy(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(PolicyABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Policy *PolicyRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Policy.Contract.PolicyCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Policy *PolicyRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Policy.Contract.PolicyTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Policy *PolicyRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Policy.Contract.PolicyTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Policy *PolicyCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Policy.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Policy *PolicyTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Policy.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Policy *PolicyTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Policy.Contract.contract.Transact(opts, method, params...)
}

// CanRead is a free data retrieval call binding the contract method 0x56554256.
//
// Solidity: function CanRead(a address) constant returns(bool)
func (_Policy *PolicyCaller) CanRead(opts *bind.CallOpts, a common.Address) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _Policy.contract.Call(opts, out, "CanRead", a)
	return *ret0, err
}

// CanRead is a free data retrieval call binding the contract method 0x56554256.
//
// Solidity: function CanRead(a address) constant returns(bool)
func (_Policy *PolicySession) CanRead(a common.Address) (bool, error) {
	return _Policy.Contract.CanRead(&_Policy.CallOpts, a)
}

// CanRead is a free data retrieval call binding the contract method 0x56554256.
//
// Solidity: function CanRead(a address) constant returns(bool)
func (_Policy *PolicyCallerSession) CanRead(a common.Address) (bool, error) {
	return _Policy.Contract.CanRead(&_Policy.CallOpts, a)
}

// AddOwner is a paid mutator transaction binding the contract method 0xac1e9ef4.
//
// Solidity: function AddOwner(a address) returns()
func (_Policy *PolicyTransactor) AddOwner(opts *bind.TransactOpts, a common.Address) (*types.Transaction, error) {
	return _Policy.contract.Transact(opts, "AddOwner", a)
}

// AddOwner is a paid mutator transaction binding the contract method 0xac1e9ef4.
//
// Solidity: function AddOwner(a address) returns()
func (_Policy *PolicySession) AddOwner(a common.Address) (*types.Transaction, error) {
	return _Policy.Contract.AddOwner(&_Policy.TransactOpts, a)
}

// AddOwner is a paid mutator transaction binding the contract method 0xac1e9ef4.
//
// Solidity: function AddOwner(a address) returns()
func (_Policy *PolicyTransactorSession) AddOwner(a common.Address) (*types.Transaction, error) {
	return _Policy.Contract.AddOwner(&_Policy.TransactOpts, a)
}

// RemoveFromPolicy is a paid mutator transaction binding the contract method 0x1809a39e.
//
// Solidity: function RemoveFromPolicy(a address) returns()
func (_Policy *PolicyTransactor) RemoveFromPolicy(opts *bind.TransactOpts, a common.Address) (*types.Transaction, error) {
	return _Policy.contract.Transact(opts, "RemoveFromPolicy", a)
}

// RemoveFromPolicy is a paid mutator transaction binding the contract method 0x1809a39e.
//
// Solidity: function RemoveFromPolicy(a address) returns()
func (_Policy *PolicySession) RemoveFromPolicy(a common.Address) (*types.Transaction, error) {
	return _Policy.Contract.RemoveFromPolicy(&_Policy.TransactOpts, a)
}

// RemoveFromPolicy is a paid mutator transaction binding the contract method 0x1809a39e.
//
// Solidity: function RemoveFromPolicy(a address) returns()
func (_Policy *PolicyTransactorSession) RemoveFromPolicy(a common.Address) (*types.Transaction, error) {
	return _Policy.Contract.RemoveFromPolicy(&_Policy.TransactOpts, a)
}

// ReadRequestABI is the input ABI used to generate the binding from.
const ReadRequestABI = "[{\"constant\":true,\"inputs\":[],\"name\":\"xc\",\"outputs\":[{\"name\":\"\",\"type\":\"bytes\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"writeRequest\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"wr\",\"type\":\"address\"}],\"name\":\"CompareReadWrite\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"WriteHash\",\"outputs\":[{\"name\":\"\",\"type\":\"bytes\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"name\":\"a\",\"type\":\"address\"},{\"name\":\"x\",\"type\":\"bytes\"},{\"name\":\"wrHash\",\"type\":\"bytes\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"}]"

// ReadRequestBin is the compiled bytecode used for deploying new contracts.
const ReadRequestBin = `0x608060405234801561001057600080fd5b506040516104f23803806104f28339810180604052606081101561003357600080fd5b81516020830180519193928301929164010000000081111561005457600080fd5b8201602081018481111561006757600080fd5b815164010000000081118282018710171561008157600080fd5b5050929190602001805164010000000081111561009d57600080fd5b820160208101848111156100b057600080fd5b81516401000000008111828201871017156100ca57600080fd5b50506002805433600160a060020a03199182161790915560008054909116600160a060020a038816179055845190935061010d925060019150602085019061012a565b50805161012190600390602084019061012a565b505050506101c5565b828054600181600116156101000203166002900490600052602060002090601f016020900481019282601f1061016b57805160ff1916838001178555610198565b82800160010185558215610198579182015b8281111561019857825182559160200191906001019061017d565b506101a49291506101a8565b5090565b6101c291905b808211156101a457600081556001016101ae565b90565b61031e806101d46000396000f3fe608060405260043610610066577c010000000000000000000000000000000000000000000000000000000060003504635baf4e9f811461006b57806364e94e67146100f55780638da5cb5b146101335780639d0089b814610148578063d27d13db1461019c575b600080fd5b34801561007757600080fd5b506100806101b1565b6040805160208082528351818301528351919283929083019185019080838360005b838110156100ba5781810151838201526020016100a2565b50505050905090810190601f1680156100e75780820380516001836020036101000a031916815260200191505b509250505060405180910390f35b34801561010157600080fd5b5061010a61023e565b6040805173ffffffffffffffffffffffffffffffffffffffff9092168252519081900360200190f35b34801561013f57600080fd5b5061010a61025a565b34801561015457600080fd5b506101886004803603602081101561016b57600080fd5b503573ffffffffffffffffffffffffffffffffffffffff16610276565b604080519115158252519081900360200190f35b3480156101a857600080fd5b50610080610297565b60018054604080516020600284861615610100026000190190941693909304601f810184900484028201840190925281815292918301828280156102365780601f1061020b57610100808354040283529160200191610236565b820191906000526020600020905b81548152906001019060200180831161021957829003601f168201915b505050505081565b60005473ffffffffffffffffffffffffffffffffffffffff1681565b60025473ffffffffffffffffffffffffffffffffffffffff1681565b60005473ffffffffffffffffffffffffffffffffffffffff90811691161490565b6003805460408051602060026001851615610100026000190190941693909304601f810184900484028201840190925281815292918301828280156102365780601f1061020b5761010080835404028352916020019161023656fea165627a7a7230582046dd76491d4a5fdf6d143543ade82522b484af2f71bcf6d748c1248b990c67720029`

// DeployReadRequest deploys a new Ethereum contract, binding an instance of ReadRequest to it.
func DeployReadRequest(auth *bind.TransactOpts, backend bind.ContractBackend, a common.Address, x []byte, wrHash []byte) (common.Address, *types.Transaction, *ReadRequest, error) {
	parsed, err := abi.JSON(strings.NewReader(ReadRequestABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(ReadRequestBin), backend, a, x, wrHash)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &ReadRequest{ReadRequestCaller: ReadRequestCaller{contract: contract}, ReadRequestTransactor: ReadRequestTransactor{contract: contract}, ReadRequestFilterer: ReadRequestFilterer{contract: contract}}, nil
}

// ReadRequest is an auto generated Go binding around an Ethereum contract.
type ReadRequest struct {
	ReadRequestCaller     // Read-only binding to the contract
	ReadRequestTransactor // Write-only binding to the contract
	ReadRequestFilterer   // Log filterer for contract events
}

// ReadRequestCaller is an auto generated read-only Go binding around an Ethereum contract.
type ReadRequestCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ReadRequestTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ReadRequestTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ReadRequestFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ReadRequestFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ReadRequestSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ReadRequestSession struct {
	Contract     *ReadRequest      // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ReadRequestCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ReadRequestCallerSession struct {
	Contract *ReadRequestCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts      // Call options to use throughout this session
}

// ReadRequestTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ReadRequestTransactorSession struct {
	Contract     *ReadRequestTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts      // Transaction auth options to use throughout this session
}

// ReadRequestRaw is an auto generated low-level Go binding around an Ethereum contract.
type ReadRequestRaw struct {
	Contract *ReadRequest // Generic contract binding to access the raw methods on
}

// ReadRequestCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ReadRequestCallerRaw struct {
	Contract *ReadRequestCaller // Generic read-only contract binding to access the raw methods on
}

// ReadRequestTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ReadRequestTransactorRaw struct {
	Contract *ReadRequestTransactor // Generic write-only contract binding to access the raw methods on
}

// NewReadRequest creates a new instance of ReadRequest, bound to a specific deployed contract.
func NewReadRequest(address common.Address, backend bind.ContractBackend) (*ReadRequest, error) {
	contract, err := bindReadRequest(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &ReadRequest{ReadRequestCaller: ReadRequestCaller{contract: contract}, ReadRequestTransactor: ReadRequestTransactor{contract: contract}, ReadRequestFilterer: ReadRequestFilterer{contract: contract}}, nil
}

// NewReadRequestCaller creates a new read-only instance of ReadRequest, bound to a specific deployed contract.
func NewReadRequestCaller(address common.Address, caller bind.ContractCaller) (*ReadRequestCaller, error) {
	contract, err := bindReadRequest(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ReadRequestCaller{contract: contract}, nil
}

// NewReadRequestTransactor creates a new write-only instance of ReadRequest, bound to a specific deployed contract.
func NewReadRequestTransactor(address common.Address, transactor bind.ContractTransactor) (*ReadRequestTransactor, error) {
	contract, err := bindReadRequest(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ReadRequestTransactor{contract: contract}, nil
}

// NewReadRequestFilterer creates a new log filterer instance of ReadRequest, bound to a specific deployed contract.
func NewReadRequestFilterer(address common.Address, filterer bind.ContractFilterer) (*ReadRequestFilterer, error) {
	contract, err := bindReadRequest(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ReadRequestFilterer{contract: contract}, nil
}

// bindReadRequest binds a generic wrapper to an already deployed contract.
func bindReadRequest(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(ReadRequestABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ReadRequest *ReadRequestRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _ReadRequest.Contract.ReadRequestCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ReadRequest *ReadRequestRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ReadRequest.Contract.ReadRequestTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ReadRequest *ReadRequestRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ReadRequest.Contract.ReadRequestTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ReadRequest *ReadRequestCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _ReadRequest.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ReadRequest *ReadRequestTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ReadRequest.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ReadRequest *ReadRequestTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ReadRequest.Contract.contract.Transact(opts, method, params...)
}

// CompareReadWrite is a free data retrieval call binding the contract method 0x9d0089b8.
//
// Solidity: function CompareReadWrite(wr address) constant returns(bool)
func (_ReadRequest *ReadRequestCaller) CompareReadWrite(opts *bind.CallOpts, wr common.Address) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _ReadRequest.contract.Call(opts, out, "CompareReadWrite", wr)
	return *ret0, err
}

// CompareReadWrite is a free data retrieval call binding the contract method 0x9d0089b8.
//
// Solidity: function CompareReadWrite(wr address) constant returns(bool)
func (_ReadRequest *ReadRequestSession) CompareReadWrite(wr common.Address) (bool, error) {
	return _ReadRequest.Contract.CompareReadWrite(&_ReadRequest.CallOpts, wr)
}

// CompareReadWrite is a free data retrieval call binding the contract method 0x9d0089b8.
//
// Solidity: function CompareReadWrite(wr address) constant returns(bool)
func (_ReadRequest *ReadRequestCallerSession) CompareReadWrite(wr common.Address) (bool, error) {
	return _ReadRequest.Contract.CompareReadWrite(&_ReadRequest.CallOpts, wr)
}

// WriteHash is a free data retrieval call binding the contract method 0xd27d13db.
//
// Solidity: function WriteHash() constant returns(bytes)
func (_ReadRequest *ReadRequestCaller) WriteHash(opts *bind.CallOpts) ([]byte, error) {
	var (
		ret0 = new([]byte)
	)
	out := ret0
	err := _ReadRequest.contract.Call(opts, out, "WriteHash")
	return *ret0, err
}

// WriteHash is a free data retrieval call binding the contract method 0xd27d13db.
//
// Solidity: function WriteHash() constant returns(bytes)
func (_ReadRequest *ReadRequestSession) WriteHash() ([]byte, error) {
	return _ReadRequest.Contract.WriteHash(&_ReadRequest.CallOpts)
}

// WriteHash is a free data retrieval call binding the contract method 0xd27d13db.
//
// Solidity: function WriteHash() constant returns(bytes)
func (_ReadRequest *ReadRequestCallerSession) WriteHash() ([]byte, error) {
	return _ReadRequest.Contract.WriteHash(&_ReadRequest.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_ReadRequest *ReadRequestCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _ReadRequest.contract.Call(opts, out, "owner")
	return *ret0, err
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_ReadRequest *ReadRequestSession) Owner() (common.Address, error) {
	return _ReadRequest.Contract.Owner(&_ReadRequest.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_ReadRequest *ReadRequestCallerSession) Owner() (common.Address, error) {
	return _ReadRequest.Contract.Owner(&_ReadRequest.CallOpts)
}

// WriteRequest is a free data retrieval call binding the contract method 0x64e94e67.
//
// Solidity: function writeRequest() constant returns(address)
func (_ReadRequest *ReadRequestCaller) WriteRequest(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _ReadRequest.contract.Call(opts, out, "writeRequest")
	return *ret0, err
}

// WriteRequest is a free data retrieval call binding the contract method 0x64e94e67.
//
// Solidity: function writeRequest() constant returns(address)
func (_ReadRequest *ReadRequestSession) WriteRequest() (common.Address, error) {
	return _ReadRequest.Contract.WriteRequest(&_ReadRequest.CallOpts)
}

// WriteRequest is a free data retrieval call binding the contract method 0x64e94e67.
//
// Solidity: function writeRequest() constant returns(address)
func (_ReadRequest *ReadRequestCallerSession) WriteRequest() (common.Address, error) {
	return _ReadRequest.Contract.WriteRequest(&_ReadRequest.CallOpts)
}

// Xc is a free data retrieval call binding the contract method 0x5baf4e9f.
//
// Solidity: function xc() constant returns(bytes)
func (_ReadRequest *ReadRequestCaller) Xc(opts *bind.CallOpts) ([]byte, error) {
	var (
		ret0 = new([]byte)
	)
	out := ret0
	err := _ReadRequest.contract.Call(opts, out, "xc")
	return *ret0, err
}

// Xc is a free data retrieval call binding the contract method 0x5baf4e9f.
//
// Solidity: function xc() constant returns(bytes)
func (_ReadRequest *ReadRequestSession) Xc() ([]byte, error) {
	return _ReadRequest.Contract.Xc(&_ReadRequest.CallOpts)
}

// Xc is a free data retrieval call binding the contract method 0x5baf4e9f.
//
// Solidity: function xc() constant returns(bytes)
func (_ReadRequest *ReadRequestCallerSession) Xc() ([]byte, error) {
	return _ReadRequest.Contract.Xc(&_ReadRequest.CallOpts)
}

// ReadRequestHolderABI is the input ABI used to generate the binding from.
const ReadRequestHolderABI = "[{\"constant\":false,\"inputs\":[{\"name\":\"a\",\"type\":\"address\"}],\"name\":\"addReadRequest\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"address\"}],\"name\":\"requests\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"address\"}],\"name\":\"hasReq\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"id\",\"type\":\"address\"}],\"name\":\"hasReadRequest\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"}]"

// ReadRequestHolderBin is the compiled bytecode used for deploying new contracts.
const ReadRequestHolderBin = `0x608060405234801561001057600080fd5b50610223806100206000396000f3fe60806040526004361061005b577c0100000000000000000000000000000000000000000000000000000000600035046331419b50811461006057806374adad1d146100955780638eb81102146100e4578063fd6336dc1461012b575b600080fd5b34801561006c57600080fd5b506100936004803603602081101561008357600080fd5b5035600160a060020a031661015e565b005b3480156100a157600080fd5b506100c8600480360360208110156100b857600080fd5b5035600160a060020a03166101a9565b60408051600160a060020a039092168252519081900360200190f35b3480156100f057600080fd5b506101176004803603602081101561010757600080fd5b5035600160a060020a03166101c4565b604080519115158252519081900360200190f35b34801561013757600080fd5b506101176004803603602081101561014e57600080fd5b5035600160a060020a03166101d9565b600160a060020a0316600081815260208181526040808320805460ff191660019081179091559091529020805473ffffffffffffffffffffffffffffffffffffffff19169091179055565b600160205260009081526040902054600160a060020a031681565b60006020819052908152604090205460ff1681565b600160a060020a031660009081526020819052604090205460ff169056fea165627a7a72305820c18b958e0b24b6957a74ae2dd9fc0175b2421ca8eedc9e779d4944da29f479c70029`

// DeployReadRequestHolder deploys a new Ethereum contract, binding an instance of ReadRequestHolder to it.
func DeployReadRequestHolder(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *ReadRequestHolder, error) {
	parsed, err := abi.JSON(strings.NewReader(ReadRequestHolderABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(ReadRequestHolderBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &ReadRequestHolder{ReadRequestHolderCaller: ReadRequestHolderCaller{contract: contract}, ReadRequestHolderTransactor: ReadRequestHolderTransactor{contract: contract}, ReadRequestHolderFilterer: ReadRequestHolderFilterer{contract: contract}}, nil
}

// ReadRequestHolder is an auto generated Go binding around an Ethereum contract.
type ReadRequestHolder struct {
	ReadRequestHolderCaller     // Read-only binding to the contract
	ReadRequestHolderTransactor // Write-only binding to the contract
	ReadRequestHolderFilterer   // Log filterer for contract events
}

// ReadRequestHolderCaller is an auto generated read-only Go binding around an Ethereum contract.
type ReadRequestHolderCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ReadRequestHolderTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ReadRequestHolderTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ReadRequestHolderFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ReadRequestHolderFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ReadRequestHolderSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ReadRequestHolderSession struct {
	Contract     *ReadRequestHolder // Generic contract binding to set the session for
	CallOpts     bind.CallOpts      // Call options to use throughout this session
	TransactOpts bind.TransactOpts  // Transaction auth options to use throughout this session
}

// ReadRequestHolderCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ReadRequestHolderCallerSession struct {
	Contract *ReadRequestHolderCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts            // Call options to use throughout this session
}

// ReadRequestHolderTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ReadRequestHolderTransactorSession struct {
	Contract     *ReadRequestHolderTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts            // Transaction auth options to use throughout this session
}

// ReadRequestHolderRaw is an auto generated low-level Go binding around an Ethereum contract.
type ReadRequestHolderRaw struct {
	Contract *ReadRequestHolder // Generic contract binding to access the raw methods on
}

// ReadRequestHolderCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ReadRequestHolderCallerRaw struct {
	Contract *ReadRequestHolderCaller // Generic read-only contract binding to access the raw methods on
}

// ReadRequestHolderTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ReadRequestHolderTransactorRaw struct {
	Contract *ReadRequestHolderTransactor // Generic write-only contract binding to access the raw methods on
}

// NewReadRequestHolder creates a new instance of ReadRequestHolder, bound to a specific deployed contract.
func NewReadRequestHolder(address common.Address, backend bind.ContractBackend) (*ReadRequestHolder, error) {
	contract, err := bindReadRequestHolder(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &ReadRequestHolder{ReadRequestHolderCaller: ReadRequestHolderCaller{contract: contract}, ReadRequestHolderTransactor: ReadRequestHolderTransactor{contract: contract}, ReadRequestHolderFilterer: ReadRequestHolderFilterer{contract: contract}}, nil
}

// NewReadRequestHolderCaller creates a new read-only instance of ReadRequestHolder, bound to a specific deployed contract.
func NewReadRequestHolderCaller(address common.Address, caller bind.ContractCaller) (*ReadRequestHolderCaller, error) {
	contract, err := bindReadRequestHolder(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ReadRequestHolderCaller{contract: contract}, nil
}

// NewReadRequestHolderTransactor creates a new write-only instance of ReadRequestHolder, bound to a specific deployed contract.
func NewReadRequestHolderTransactor(address common.Address, transactor bind.ContractTransactor) (*ReadRequestHolderTransactor, error) {
	contract, err := bindReadRequestHolder(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ReadRequestHolderTransactor{contract: contract}, nil
}

// NewReadRequestHolderFilterer creates a new log filterer instance of ReadRequestHolder, bound to a specific deployed contract.
func NewReadRequestHolderFilterer(address common.Address, filterer bind.ContractFilterer) (*ReadRequestHolderFilterer, error) {
	contract, err := bindReadRequestHolder(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ReadRequestHolderFilterer{contract: contract}, nil
}

// bindReadRequestHolder binds a generic wrapper to an already deployed contract.
func bindReadRequestHolder(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(ReadRequestHolderABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ReadRequestHolder *ReadRequestHolderRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _ReadRequestHolder.Contract.ReadRequestHolderCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ReadRequestHolder *ReadRequestHolderRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ReadRequestHolder.Contract.ReadRequestHolderTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ReadRequestHolder *ReadRequestHolderRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ReadRequestHolder.Contract.ReadRequestHolderTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ReadRequestHolder *ReadRequestHolderCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _ReadRequestHolder.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ReadRequestHolder *ReadRequestHolderTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ReadRequestHolder.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ReadRequestHolder *ReadRequestHolderTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ReadRequestHolder.Contract.contract.Transact(opts, method, params...)
}

// HasReadRequest is a free data retrieval call binding the contract method 0xfd6336dc.
//
// Solidity: function hasReadRequest(id address) constant returns(bool)
func (_ReadRequestHolder *ReadRequestHolderCaller) HasReadRequest(opts *bind.CallOpts, id common.Address) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _ReadRequestHolder.contract.Call(opts, out, "hasReadRequest", id)
	return *ret0, err
}

// HasReadRequest is a free data retrieval call binding the contract method 0xfd6336dc.
//
// Solidity: function hasReadRequest(id address) constant returns(bool)
func (_ReadRequestHolder *ReadRequestHolderSession) HasReadRequest(id common.Address) (bool, error) {
	return _ReadRequestHolder.Contract.HasReadRequest(&_ReadRequestHolder.CallOpts, id)
}

// HasReadRequest is a free data retrieval call binding the contract method 0xfd6336dc.
//
// Solidity: function hasReadRequest(id address) constant returns(bool)
func (_ReadRequestHolder *ReadRequestHolderCallerSession) HasReadRequest(id common.Address) (bool, error) {
	return _ReadRequestHolder.Contract.HasReadRequest(&_ReadRequestHolder.CallOpts, id)
}

// HasReq is a free data retrieval call binding the contract method 0x8eb81102.
//
// Solidity: function hasReq( address) constant returns(bool)
func (_ReadRequestHolder *ReadRequestHolderCaller) HasReq(opts *bind.CallOpts, arg0 common.Address) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _ReadRequestHolder.contract.Call(opts, out, "hasReq", arg0)
	return *ret0, err
}

// HasReq is a free data retrieval call binding the contract method 0x8eb81102.
//
// Solidity: function hasReq( address) constant returns(bool)
func (_ReadRequestHolder *ReadRequestHolderSession) HasReq(arg0 common.Address) (bool, error) {
	return _ReadRequestHolder.Contract.HasReq(&_ReadRequestHolder.CallOpts, arg0)
}

// HasReq is a free data retrieval call binding the contract method 0x8eb81102.
//
// Solidity: function hasReq( address) constant returns(bool)
func (_ReadRequestHolder *ReadRequestHolderCallerSession) HasReq(arg0 common.Address) (bool, error) {
	return _ReadRequestHolder.Contract.HasReq(&_ReadRequestHolder.CallOpts, arg0)
}

// Requests is a free data retrieval call binding the contract method 0x74adad1d.
//
// Solidity: function requests( address) constant returns(address)
func (_ReadRequestHolder *ReadRequestHolderCaller) Requests(opts *bind.CallOpts, arg0 common.Address) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _ReadRequestHolder.contract.Call(opts, out, "requests", arg0)
	return *ret0, err
}

// Requests is a free data retrieval call binding the contract method 0x74adad1d.
//
// Solidity: function requests( address) constant returns(address)
func (_ReadRequestHolder *ReadRequestHolderSession) Requests(arg0 common.Address) (common.Address, error) {
	return _ReadRequestHolder.Contract.Requests(&_ReadRequestHolder.CallOpts, arg0)
}

// Requests is a free data retrieval call binding the contract method 0x74adad1d.
//
// Solidity: function requests( address) constant returns(address)
func (_ReadRequestHolder *ReadRequestHolderCallerSession) Requests(arg0 common.Address) (common.Address, error) {
	return _ReadRequestHolder.Contract.Requests(&_ReadRequestHolder.CallOpts, arg0)
}

// AddReadRequest is a paid mutator transaction binding the contract method 0x31419b50.
//
// Solidity: function addReadRequest(a address) returns()
func (_ReadRequestHolder *ReadRequestHolderTransactor) AddReadRequest(opts *bind.TransactOpts, a common.Address) (*types.Transaction, error) {
	return _ReadRequestHolder.contract.Transact(opts, "addReadRequest", a)
}

// AddReadRequest is a paid mutator transaction binding the contract method 0x31419b50.
//
// Solidity: function addReadRequest(a address) returns()
func (_ReadRequestHolder *ReadRequestHolderSession) AddReadRequest(a common.Address) (*types.Transaction, error) {
	return _ReadRequestHolder.Contract.AddReadRequest(&_ReadRequestHolder.TransactOpts, a)
}

// AddReadRequest is a paid mutator transaction binding the contract method 0x31419b50.
//
// Solidity: function addReadRequest(a address) returns()
func (_ReadRequestHolder *ReadRequestHolderTransactorSession) AddReadRequest(a common.Address) (*types.Transaction, error) {
	return _ReadRequestHolder.Contract.AddReadRequest(&_ReadRequestHolder.TransactOpts, a)
}

// WriteRequestABI is the input ABI used to generate the binding from.
const WriteRequestABI = "[{\"constant\":true,\"inputs\":[],\"name\":\"LTSID\",\"outputs\":[{\"name\":\"\",\"type\":\"bytes\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"U\",\"outputs\":[{\"name\":\"\",\"type\":\"bytes\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"F\",\"outputs\":[{\"name\":\"\",\"type\":\"bytes\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"a\",\"type\":\"address\"}],\"name\":\"CanRead\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"GetSlice\",\"outputs\":[{\"name\":\"\",\"type\":\"int64[]\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"extraData\",\"outputs\":[{\"name\":\"\",\"type\":\"bytes\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"data\",\"outputs\":[{\"name\":\"\",\"type\":\"bytes\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"E\",\"outputs\":[{\"name\":\"\",\"type\":\"bytes\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"Cs\",\"outputs\":[{\"name\":\"\",\"type\":\"bytes\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"X\",\"outputs\":[{\"name\":\"\",\"type\":\"bytes\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"slice\",\"outputs\":[{\"name\":\"\",\"type\":\"int64\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"Ubar\",\"outputs\":[{\"name\":\"\",\"type\":\"bytes\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"name\":\"d\",\"type\":\"bytes\"},{\"name\":\"ed\",\"type\":\"bytes\"},{\"name\":\"l\",\"type\":\"bytes\"},{\"name\":\"p\",\"type\":\"address\"},{\"name\":\"u\",\"type\":\"bytes\"},{\"name\":\"cs\",\"type\":\"bytes\"},{\"name\":\"sl\",\"type\":\"int64[]\"},{\"name\":\"f\",\"type\":\"bytes\"},{\"name\":\"e\",\"type\":\"bytes\"},{\"name\":\"ub\",\"type\":\"bytes\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"}]"

// WriteRequestBin is the compiled bytecode used for deploying new contracts.
const WriteRequestBin = `0x60806040523480156200001157600080fd5b5060405162000d4f38038062000d4f83398101806040526101408110156200003857600080fd5b8101908080516401000000008111156200005157600080fd5b820160208101848111156200006557600080fd5b81516401000000008111828201871017156200008057600080fd5b505092919060200180516401000000008111156200009d57600080fd5b82016020810184811115620000b157600080fd5b8151640100000000811182820187101715620000cc57600080fd5b50509291906020018051640100000000811115620000e957600080fd5b82016020810184811115620000fd57600080fd5b81516401000000008111828201871017156200011857600080fd5b505060208201516040909201805191949293916401000000008111156200013e57600080fd5b820160208101848111156200015257600080fd5b81516401000000008111828201871017156200016d57600080fd5b505092919060200180516401000000008111156200018a57600080fd5b820160208101848111156200019e57600080fd5b8151640100000000811182820187101715620001b957600080fd5b50509291906020018051640100000000811115620001d657600080fd5b82016020810184811115620001ea57600080fd5b81518560208202830111640100000000821117156200020857600080fd5b505092919060200180516401000000008111156200022557600080fd5b820160208101848111156200023957600080fd5b81516401000000008111828201871017156200025457600080fd5b505092919060200180516401000000008111156200027157600080fd5b820160208101848111156200028557600080fd5b8151640100000000811182820187101715620002a057600080fd5b50509291906020018051640100000000811115620002bd57600080fd5b82016020810184811115620002d157600080fd5b8151640100000000811182820187101715620002ec57600080fd5b50508c519093506200030892506000915060208d0190620003e4565b5088516200031e9060019060208c0190620003e4565b508751620003349060029060208b0190620003e4565b5060098054600160a060020a031916600160a060020a038916179055855162000365906003906020890190620003e4565b5084516200037b906008906020880190620003e4565b5083516200039190600a90602087019062000469565b508251620003a7906006906020860190620003e4565b508151620003bd906005906020850190620003e4565b508051620003d3906004906020840190620003e4565b50505050505050505050506200056f565b828054600181600116156101000203166002900490600052602060002090601f016020900481019282601f106200042757805160ff191683800117855562000457565b8280016001018555821562000457579182015b82811115620004575782518255916020019190600101906200043a565b506200046592915062000527565b5090565b82805482825590600052602060002090600301600490048101928215620005195791602002820160005b83821115620004e257835183826101000a8154816001604060020a03021916908360070b6001604060020a03160217905550926020019260080160208160070104928301926001030262000493565b8015620005175782816101000a8154906001604060020a030219169055600801602081600701049283019260010302620004e2565b505b506200046592915062000547565b6200054491905b808211156200046557600081556001016200052e565b90565b6200054491905b808211156200046557805467ffffffffffffffff191681556001016200054e565b6107d0806200057f6000396000f3fe6080604052600436106100b3577c010000000000000000000000000000000000000000000000000000000060003504630fdcdbfd81146100b85780632688454a1461014257806328811f5914610157578063565542561461016c5780635bf2abcd146101c0578063609d33341461022557806373d4a13a1461023a57806392bbf6e81461024f578063b6f9cf6a14610264578063c1599bd914610279578063e6a01b5a1461028e578063eeeee14d146102d1575b600080fd5b3480156100c457600080fd5b506100cd6102e6565b6040805160208082528351818301528351919283929083019185019080838360005b838110156101075781810151838201526020016100ef565b50505050905090810190601f1680156101345780820380516001836020036101000a031916815260200191505b509250505060405180910390f35b34801561014e57600080fd5b506100cd610371565b34801561016357600080fd5b506100cd6103cc565b34801561017857600080fd5b506101ac6004803603602081101561018f57600080fd5b503573ffffffffffffffffffffffffffffffffffffffff16610427565b604080519115158252519081900360200190f35b3480156101cc57600080fd5b506101d56104d0565b60408051602080825283518183015283519192839290830191858101910280838360005b838110156102115781810151838201526020016101f9565b505050509050019250505060405180910390f35b34801561023157600080fd5b506100cd61054e565b34801561024657600080fd5b506100cd6105a8565b34801561025b57600080fd5b506100cd610603565b34801561027057600080fd5b506100cd61065e565b34801561028557600080fd5b506100cd6106b9565b34801561029a57600080fd5b506102b8600480360360208110156102b157600080fd5b5035610714565b60408051600792830b90920b8252519081900360200190f35b3480156102dd57600080fd5b506100cd610749565b6002805460408051602060018416156101000260001901909316849004601f810184900484028201840190925281815292918301828280156103695780601f1061033e57610100808354040283529160200191610369565b820191906000526020600020905b81548152906001019060200180831161034c57829003601f168201915b505050505081565b6003805460408051602060026001851615610100026000190190941693909304601f810184900484028201840190925281815292918301828280156103695780601f1061033e57610100808354040283529160200191610369565b6006805460408051602060026001851615610100026000190190941693909304601f810184900484028201840190925281815292918301828280156103695780601f1061033e57610100808354040283529160200191610369565b600954604080517f5655425600000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff848116600483015291516000939290921691635655425691602480820192602092909190829003018186803b15801561049e57600080fd5b505afa1580156104b2573d6000803e3d6000fd5b505050506040513d60208110156104c857600080fd5b505192915050565b6060600a80548060200260200160405190810160405280929190818152602001828054801561054457602002820191906000526020600020906000905b82829054906101000a900460070b60070b8152602001906008019060208260070104928301926001038202915080841161050d5790505b5050505050905090565b60018054604080516020600284861615610100026000190190941693909304601f810184900484028201840190925281815292918301828280156103695780601f1061033e57610100808354040283529160200191610369565b6000805460408051602060026001851615610100026000190190941693909304601f810184900484028201840190925281815292918301828280156103695780601f1061033e57610100808354040283529160200191610369565b6005805460408051602060026001851615610100026000190190941693909304601f810184900484028201840190925281815292918301828280156103695780601f1061033e57610100808354040283529160200191610369565b6008805460408051602060026001851615610100026000190190941693909304601f810184900484028201840190925281815292918301828280156103695780601f1061033e57610100808354040283529160200191610369565b6007805460408051602060026001851615610100026000190190941693909304601f810184900484028201840190925281815292918301828280156103695780601f1061033e57610100808354040283529160200191610369565b600a80548290811061072257fe5b9060005260206000209060049182820401919006600802915054906101000a900460070b81565b6004805460408051602060026001851615610100026000190190941693909304601f810184900484028201840190925281815292918301828280156103695780601f1061033e5761010080835404028352916020019161036956fea165627a7a7230582068bbf7bc61fe7f9e0bd52bd87dfb1bb908db980e18654c467b4d3bdeb3bd7e480029`

// DeployWriteRequest deploys a new Ethereum contract, binding an instance of WriteRequest to it.
func DeployWriteRequest(auth *bind.TransactOpts, backend bind.ContractBackend, d []byte, ed []byte, l []byte, p common.Address, u []byte, cs []byte, sl []int64, f []byte, e []byte, ub []byte) (common.Address, *types.Transaction, *WriteRequest, error) {
	parsed, err := abi.JSON(strings.NewReader(WriteRequestABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(WriteRequestBin), backend, d, ed, l, p, u, cs, sl, f, e, ub)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &WriteRequest{WriteRequestCaller: WriteRequestCaller{contract: contract}, WriteRequestTransactor: WriteRequestTransactor{contract: contract}, WriteRequestFilterer: WriteRequestFilterer{contract: contract}}, nil
}

// WriteRequest is an auto generated Go binding around an Ethereum contract.
type WriteRequest struct {
	WriteRequestCaller     // Read-only binding to the contract
	WriteRequestTransactor // Write-only binding to the contract
	WriteRequestFilterer   // Log filterer for contract events
}

// WriteRequestCaller is an auto generated read-only Go binding around an Ethereum contract.
type WriteRequestCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// WriteRequestTransactor is an auto generated write-only Go binding around an Ethereum contract.
type WriteRequestTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// WriteRequestFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type WriteRequestFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// WriteRequestSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type WriteRequestSession struct {
	Contract     *WriteRequest     // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// WriteRequestCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type WriteRequestCallerSession struct {
	Contract *WriteRequestCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts       // Call options to use throughout this session
}

// WriteRequestTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type WriteRequestTransactorSession struct {
	Contract     *WriteRequestTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts       // Transaction auth options to use throughout this session
}

// WriteRequestRaw is an auto generated low-level Go binding around an Ethereum contract.
type WriteRequestRaw struct {
	Contract *WriteRequest // Generic contract binding to access the raw methods on
}

// WriteRequestCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type WriteRequestCallerRaw struct {
	Contract *WriteRequestCaller // Generic read-only contract binding to access the raw methods on
}

// WriteRequestTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type WriteRequestTransactorRaw struct {
	Contract *WriteRequestTransactor // Generic write-only contract binding to access the raw methods on
}

// NewWriteRequest creates a new instance of WriteRequest, bound to a specific deployed contract.
func NewWriteRequest(address common.Address, backend bind.ContractBackend) (*WriteRequest, error) {
	contract, err := bindWriteRequest(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &WriteRequest{WriteRequestCaller: WriteRequestCaller{contract: contract}, WriteRequestTransactor: WriteRequestTransactor{contract: contract}, WriteRequestFilterer: WriteRequestFilterer{contract: contract}}, nil
}

// NewWriteRequestCaller creates a new read-only instance of WriteRequest, bound to a specific deployed contract.
func NewWriteRequestCaller(address common.Address, caller bind.ContractCaller) (*WriteRequestCaller, error) {
	contract, err := bindWriteRequest(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &WriteRequestCaller{contract: contract}, nil
}

// NewWriteRequestTransactor creates a new write-only instance of WriteRequest, bound to a specific deployed contract.
func NewWriteRequestTransactor(address common.Address, transactor bind.ContractTransactor) (*WriteRequestTransactor, error) {
	contract, err := bindWriteRequest(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &WriteRequestTransactor{contract: contract}, nil
}

// NewWriteRequestFilterer creates a new log filterer instance of WriteRequest, bound to a specific deployed contract.
func NewWriteRequestFilterer(address common.Address, filterer bind.ContractFilterer) (*WriteRequestFilterer, error) {
	contract, err := bindWriteRequest(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &WriteRequestFilterer{contract: contract}, nil
}

// bindWriteRequest binds a generic wrapper to an already deployed contract.
func bindWriteRequest(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(WriteRequestABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_WriteRequest *WriteRequestRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _WriteRequest.Contract.WriteRequestCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_WriteRequest *WriteRequestRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _WriteRequest.Contract.WriteRequestTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_WriteRequest *WriteRequestRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _WriteRequest.Contract.WriteRequestTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_WriteRequest *WriteRequestCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _WriteRequest.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_WriteRequest *WriteRequestTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _WriteRequest.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_WriteRequest *WriteRequestTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _WriteRequest.Contract.contract.Transact(opts, method, params...)
}

// CanRead is a free data retrieval call binding the contract method 0x56554256.
//
// Solidity: function CanRead(a address) constant returns(bool)
func (_WriteRequest *WriteRequestCaller) CanRead(opts *bind.CallOpts, a common.Address) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _WriteRequest.contract.Call(opts, out, "CanRead", a)
	return *ret0, err
}

// CanRead is a free data retrieval call binding the contract method 0x56554256.
//
// Solidity: function CanRead(a address) constant returns(bool)
func (_WriteRequest *WriteRequestSession) CanRead(a common.Address) (bool, error) {
	return _WriteRequest.Contract.CanRead(&_WriteRequest.CallOpts, a)
}

// CanRead is a free data retrieval call binding the contract method 0x56554256.
//
// Solidity: function CanRead(a address) constant returns(bool)
func (_WriteRequest *WriteRequestCallerSession) CanRead(a common.Address) (bool, error) {
	return _WriteRequest.Contract.CanRead(&_WriteRequest.CallOpts, a)
}

// Cs is a free data retrieval call binding the contract method 0xb6f9cf6a.
//
// Solidity: function Cs() constant returns(bytes)
func (_WriteRequest *WriteRequestCaller) Cs(opts *bind.CallOpts) ([]byte, error) {
	var (
		ret0 = new([]byte)
	)
	out := ret0
	err := _WriteRequest.contract.Call(opts, out, "Cs")
	return *ret0, err
}

// Cs is a free data retrieval call binding the contract method 0xb6f9cf6a.
//
// Solidity: function Cs() constant returns(bytes)
func (_WriteRequest *WriteRequestSession) Cs() ([]byte, error) {
	return _WriteRequest.Contract.Cs(&_WriteRequest.CallOpts)
}

// Cs is a free data retrieval call binding the contract method 0xb6f9cf6a.
//
// Solidity: function Cs() constant returns(bytes)
func (_WriteRequest *WriteRequestCallerSession) Cs() ([]byte, error) {
	return _WriteRequest.Contract.Cs(&_WriteRequest.CallOpts)
}

// E is a free data retrieval call binding the contract method 0x92bbf6e8.
//
// Solidity: function E() constant returns(bytes)
func (_WriteRequest *WriteRequestCaller) E(opts *bind.CallOpts) ([]byte, error) {
	var (
		ret0 = new([]byte)
	)
	out := ret0
	err := _WriteRequest.contract.Call(opts, out, "E")
	return *ret0, err
}

// E is a free data retrieval call binding the contract method 0x92bbf6e8.
//
// Solidity: function E() constant returns(bytes)
func (_WriteRequest *WriteRequestSession) E() ([]byte, error) {
	return _WriteRequest.Contract.E(&_WriteRequest.CallOpts)
}

// E is a free data retrieval call binding the contract method 0x92bbf6e8.
//
// Solidity: function E() constant returns(bytes)
func (_WriteRequest *WriteRequestCallerSession) E() ([]byte, error) {
	return _WriteRequest.Contract.E(&_WriteRequest.CallOpts)
}

// F is a free data retrieval call binding the contract method 0x28811f59.
//
// Solidity: function F() constant returns(bytes)
func (_WriteRequest *WriteRequestCaller) F(opts *bind.CallOpts) ([]byte, error) {
	var (
		ret0 = new([]byte)
	)
	out := ret0
	err := _WriteRequest.contract.Call(opts, out, "F")
	return *ret0, err
}

// F is a free data retrieval call binding the contract method 0x28811f59.
//
// Solidity: function F() constant returns(bytes)
func (_WriteRequest *WriteRequestSession) F() ([]byte, error) {
	return _WriteRequest.Contract.F(&_WriteRequest.CallOpts)
}

// F is a free data retrieval call binding the contract method 0x28811f59.
//
// Solidity: function F() constant returns(bytes)
func (_WriteRequest *WriteRequestCallerSession) F() ([]byte, error) {
	return _WriteRequest.Contract.F(&_WriteRequest.CallOpts)
}

// GetSlice is a free data retrieval call binding the contract method 0x5bf2abcd.
//
// Solidity: function GetSlice() constant returns(int64[])
func (_WriteRequest *WriteRequestCaller) GetSlice(opts *bind.CallOpts) ([]int64, error) {
	var (
		ret0 = new([]int64)
	)
	out := ret0
	err := _WriteRequest.contract.Call(opts, out, "GetSlice")
	return *ret0, err
}

// GetSlice is a free data retrieval call binding the contract method 0x5bf2abcd.
//
// Solidity: function GetSlice() constant returns(int64[])
func (_WriteRequest *WriteRequestSession) GetSlice() ([]int64, error) {
	return _WriteRequest.Contract.GetSlice(&_WriteRequest.CallOpts)
}

// GetSlice is a free data retrieval call binding the contract method 0x5bf2abcd.
//
// Solidity: function GetSlice() constant returns(int64[])
func (_WriteRequest *WriteRequestCallerSession) GetSlice() ([]int64, error) {
	return _WriteRequest.Contract.GetSlice(&_WriteRequest.CallOpts)
}

// LTSID is a free data retrieval call binding the contract method 0x0fdcdbfd.
//
// Solidity: function LTSID() constant returns(bytes)
func (_WriteRequest *WriteRequestCaller) LTSID(opts *bind.CallOpts) ([]byte, error) {
	var (
		ret0 = new([]byte)
	)
	out := ret0
	err := _WriteRequest.contract.Call(opts, out, "LTSID")
	return *ret0, err
}

// LTSID is a free data retrieval call binding the contract method 0x0fdcdbfd.
//
// Solidity: function LTSID() constant returns(bytes)
func (_WriteRequest *WriteRequestSession) LTSID() ([]byte, error) {
	return _WriteRequest.Contract.LTSID(&_WriteRequest.CallOpts)
}

// LTSID is a free data retrieval call binding the contract method 0x0fdcdbfd.
//
// Solidity: function LTSID() constant returns(bytes)
func (_WriteRequest *WriteRequestCallerSession) LTSID() ([]byte, error) {
	return _WriteRequest.Contract.LTSID(&_WriteRequest.CallOpts)
}

// U is a free data retrieval call binding the contract method 0x2688454a.
//
// Solidity: function U() constant returns(bytes)
func (_WriteRequest *WriteRequestCaller) U(opts *bind.CallOpts) ([]byte, error) {
	var (
		ret0 = new([]byte)
	)
	out := ret0
	err := _WriteRequest.contract.Call(opts, out, "U")
	return *ret0, err
}

// U is a free data retrieval call binding the contract method 0x2688454a.
//
// Solidity: function U() constant returns(bytes)
func (_WriteRequest *WriteRequestSession) U() ([]byte, error) {
	return _WriteRequest.Contract.U(&_WriteRequest.CallOpts)
}

// U is a free data retrieval call binding the contract method 0x2688454a.
//
// Solidity: function U() constant returns(bytes)
func (_WriteRequest *WriteRequestCallerSession) U() ([]byte, error) {
	return _WriteRequest.Contract.U(&_WriteRequest.CallOpts)
}

// Ubar is a free data retrieval call binding the contract method 0xeeeee14d.
//
// Solidity: function Ubar() constant returns(bytes)
func (_WriteRequest *WriteRequestCaller) Ubar(opts *bind.CallOpts) ([]byte, error) {
	var (
		ret0 = new([]byte)
	)
	out := ret0
	err := _WriteRequest.contract.Call(opts, out, "Ubar")
	return *ret0, err
}

// Ubar is a free data retrieval call binding the contract method 0xeeeee14d.
//
// Solidity: function Ubar() constant returns(bytes)
func (_WriteRequest *WriteRequestSession) Ubar() ([]byte, error) {
	return _WriteRequest.Contract.Ubar(&_WriteRequest.CallOpts)
}

// Ubar is a free data retrieval call binding the contract method 0xeeeee14d.
//
// Solidity: function Ubar() constant returns(bytes)
func (_WriteRequest *WriteRequestCallerSession) Ubar() ([]byte, error) {
	return _WriteRequest.Contract.Ubar(&_WriteRequest.CallOpts)
}

// X is a free data retrieval call binding the contract method 0xc1599bd9.
//
// Solidity: function X() constant returns(bytes)
func (_WriteRequest *WriteRequestCaller) X(opts *bind.CallOpts) ([]byte, error) {
	var (
		ret0 = new([]byte)
	)
	out := ret0
	err := _WriteRequest.contract.Call(opts, out, "X")
	return *ret0, err
}

// X is a free data retrieval call binding the contract method 0xc1599bd9.
//
// Solidity: function X() constant returns(bytes)
func (_WriteRequest *WriteRequestSession) X() ([]byte, error) {
	return _WriteRequest.Contract.X(&_WriteRequest.CallOpts)
}

// X is a free data retrieval call binding the contract method 0xc1599bd9.
//
// Solidity: function X() constant returns(bytes)
func (_WriteRequest *WriteRequestCallerSession) X() ([]byte, error) {
	return _WriteRequest.Contract.X(&_WriteRequest.CallOpts)
}

// Data is a free data retrieval call binding the contract method 0x73d4a13a.
//
// Solidity: function data() constant returns(bytes)
func (_WriteRequest *WriteRequestCaller) Data(opts *bind.CallOpts) ([]byte, error) {
	var (
		ret0 = new([]byte)
	)
	out := ret0
	err := _WriteRequest.contract.Call(opts, out, "data")
	return *ret0, err
}

// Data is a free data retrieval call binding the contract method 0x73d4a13a.
//
// Solidity: function data() constant returns(bytes)
func (_WriteRequest *WriteRequestSession) Data() ([]byte, error) {
	return _WriteRequest.Contract.Data(&_WriteRequest.CallOpts)
}

// Data is a free data retrieval call binding the contract method 0x73d4a13a.
//
// Solidity: function data() constant returns(bytes)
func (_WriteRequest *WriteRequestCallerSession) Data() ([]byte, error) {
	return _WriteRequest.Contract.Data(&_WriteRequest.CallOpts)
}

// ExtraData is a free data retrieval call binding the contract method 0x609d3334.
//
// Solidity: function extraData() constant returns(bytes)
func (_WriteRequest *WriteRequestCaller) ExtraData(opts *bind.CallOpts) ([]byte, error) {
	var (
		ret0 = new([]byte)
	)
	out := ret0
	err := _WriteRequest.contract.Call(opts, out, "extraData")
	return *ret0, err
}

// ExtraData is a free data retrieval call binding the contract method 0x609d3334.
//
// Solidity: function extraData() constant returns(bytes)
func (_WriteRequest *WriteRequestSession) ExtraData() ([]byte, error) {
	return _WriteRequest.Contract.ExtraData(&_WriteRequest.CallOpts)
}

// ExtraData is a free data retrieval call binding the contract method 0x609d3334.
//
// Solidity: function extraData() constant returns(bytes)
func (_WriteRequest *WriteRequestCallerSession) ExtraData() ([]byte, error) {
	return _WriteRequest.Contract.ExtraData(&_WriteRequest.CallOpts)
}

// Slice is a free data retrieval call binding the contract method 0xe6a01b5a.
//
// Solidity: function slice( uint256) constant returns(int64)
func (_WriteRequest *WriteRequestCaller) Slice(opts *bind.CallOpts, arg0 *big.Int) (int64, error) {
	var (
		ret0 = new(int64)
	)
	out := ret0
	err := _WriteRequest.contract.Call(opts, out, "slice", arg0)
	return *ret0, err
}

// Slice is a free data retrieval call binding the contract method 0xe6a01b5a.
//
// Solidity: function slice( uint256) constant returns(int64)
func (_WriteRequest *WriteRequestSession) Slice(arg0 *big.Int) (int64, error) {
	return _WriteRequest.Contract.Slice(&_WriteRequest.CallOpts, arg0)
}

// Slice is a free data retrieval call binding the contract method 0xe6a01b5a.
//
// Solidity: function slice( uint256) constant returns(int64)
func (_WriteRequest *WriteRequestCallerSession) Slice(arg0 *big.Int) (int64, error) {
	return _WriteRequest.Contract.Slice(&_WriteRequest.CallOpts, arg0)
}

// WriteRequestHolderABI is the input ABI used to generate the binding from.
const WriteRequestHolderABI = "[{\"constant\":false,\"inputs\":[{\"name\":\"a\",\"type\":\"address\"}],\"name\":\"addWriteRequest\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"a\",\"type\":\"address\"}],\"name\":\"canRead\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"a\",\"type\":\"address\"}],\"name\":\"getRequestByID\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"address\"}],\"name\":\"hasRequest\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"address\"}],\"name\":\"wrmap\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"}]"

// WriteRequestHolderBin is the compiled bytecode used for deploying new contracts.
const WriteRequestHolderBin = `0x608060405234801561001057600080fd5b50610281806100206000396000f3fe608060405260043610610066577c010000000000000000000000000000000000000000000000000000000060003504631029a017811461006b57806342087d4f146100a0578063b7312591146100e7578063e216ea2b14610136578063ed85db2314610169575b600080fd5b34801561007757600080fd5b5061009e6004803603602081101561008e57600080fd5b5035600160a060020a031661019c565b005b3480156100ac57600080fd5b506100d3600480360360208110156100c357600080fd5b5035600160a060020a03166101e9565b604080519115158252519081900360200190f35b3480156100f357600080fd5b5061011a6004803603602081101561010a57600080fd5b5035600160a060020a0316610207565b60408051600160a060020a039092168252519081900360200190f35b34801561014257600080fd5b506100d36004803603602081101561015957600080fd5b5035600160a060020a0316610225565b34801561017557600080fd5b5061011a6004803603602081101561018c57600080fd5b5035600160a060020a031661023a565b600160a060020a0316600081815260208181526040808320805473ffffffffffffffffffffffffffffffffffffffff1916909417909355600190819052919020805460ff19169091179055565b600160a060020a031660009081526001602052604090205460ff1690565b600160a060020a039081166000908152602081905260409020541690565b60016020526000908152604090205460ff1681565b600060208190529081526040902054600160a060020a03168156fea165627a7a723058203b9df46251225a5f4dc7da4ac5584c6e551532fcf505f47a682c49c7a8469e8b0029`

// DeployWriteRequestHolder deploys a new Ethereum contract, binding an instance of WriteRequestHolder to it.
func DeployWriteRequestHolder(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *WriteRequestHolder, error) {
	parsed, err := abi.JSON(strings.NewReader(WriteRequestHolderABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(WriteRequestHolderBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &WriteRequestHolder{WriteRequestHolderCaller: WriteRequestHolderCaller{contract: contract}, WriteRequestHolderTransactor: WriteRequestHolderTransactor{contract: contract}, WriteRequestHolderFilterer: WriteRequestHolderFilterer{contract: contract}}, nil
}

// WriteRequestHolder is an auto generated Go binding around an Ethereum contract.
type WriteRequestHolder struct {
	WriteRequestHolderCaller     // Read-only binding to the contract
	WriteRequestHolderTransactor // Write-only binding to the contract
	WriteRequestHolderFilterer   // Log filterer for contract events
}

// WriteRequestHolderCaller is an auto generated read-only Go binding around an Ethereum contract.
type WriteRequestHolderCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// WriteRequestHolderTransactor is an auto generated write-only Go binding around an Ethereum contract.
type WriteRequestHolderTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// WriteRequestHolderFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type WriteRequestHolderFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// WriteRequestHolderSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type WriteRequestHolderSession struct {
	Contract     *WriteRequestHolder // Generic contract binding to set the session for
	CallOpts     bind.CallOpts       // Call options to use throughout this session
	TransactOpts bind.TransactOpts   // Transaction auth options to use throughout this session
}

// WriteRequestHolderCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type WriteRequestHolderCallerSession struct {
	Contract *WriteRequestHolderCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts             // Call options to use throughout this session
}

// WriteRequestHolderTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type WriteRequestHolderTransactorSession struct {
	Contract     *WriteRequestHolderTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts             // Transaction auth options to use throughout this session
}

// WriteRequestHolderRaw is an auto generated low-level Go binding around an Ethereum contract.
type WriteRequestHolderRaw struct {
	Contract *WriteRequestHolder // Generic contract binding to access the raw methods on
}

// WriteRequestHolderCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type WriteRequestHolderCallerRaw struct {
	Contract *WriteRequestHolderCaller // Generic read-only contract binding to access the raw methods on
}

// WriteRequestHolderTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type WriteRequestHolderTransactorRaw struct {
	Contract *WriteRequestHolderTransactor // Generic write-only contract binding to access the raw methods on
}

// NewWriteRequestHolder creates a new instance of WriteRequestHolder, bound to a specific deployed contract.
func NewWriteRequestHolder(address common.Address, backend bind.ContractBackend) (*WriteRequestHolder, error) {
	contract, err := bindWriteRequestHolder(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &WriteRequestHolder{WriteRequestHolderCaller: WriteRequestHolderCaller{contract: contract}, WriteRequestHolderTransactor: WriteRequestHolderTransactor{contract: contract}, WriteRequestHolderFilterer: WriteRequestHolderFilterer{contract: contract}}, nil
}

// NewWriteRequestHolderCaller creates a new read-only instance of WriteRequestHolder, bound to a specific deployed contract.
func NewWriteRequestHolderCaller(address common.Address, caller bind.ContractCaller) (*WriteRequestHolderCaller, error) {
	contract, err := bindWriteRequestHolder(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &WriteRequestHolderCaller{contract: contract}, nil
}

// NewWriteRequestHolderTransactor creates a new write-only instance of WriteRequestHolder, bound to a specific deployed contract.
func NewWriteRequestHolderTransactor(address common.Address, transactor bind.ContractTransactor) (*WriteRequestHolderTransactor, error) {
	contract, err := bindWriteRequestHolder(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &WriteRequestHolderTransactor{contract: contract}, nil
}

// NewWriteRequestHolderFilterer creates a new log filterer instance of WriteRequestHolder, bound to a specific deployed contract.
func NewWriteRequestHolderFilterer(address common.Address, filterer bind.ContractFilterer) (*WriteRequestHolderFilterer, error) {
	contract, err := bindWriteRequestHolder(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &WriteRequestHolderFilterer{contract: contract}, nil
}

// bindWriteRequestHolder binds a generic wrapper to an already deployed contract.
func bindWriteRequestHolder(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(WriteRequestHolderABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_WriteRequestHolder *WriteRequestHolderRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _WriteRequestHolder.Contract.WriteRequestHolderCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_WriteRequestHolder *WriteRequestHolderRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _WriteRequestHolder.Contract.WriteRequestHolderTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_WriteRequestHolder *WriteRequestHolderRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _WriteRequestHolder.Contract.WriteRequestHolderTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_WriteRequestHolder *WriteRequestHolderCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _WriteRequestHolder.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_WriteRequestHolder *WriteRequestHolderTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _WriteRequestHolder.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_WriteRequestHolder *WriteRequestHolderTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _WriteRequestHolder.Contract.contract.Transact(opts, method, params...)
}

// GetRequestByID is a free data retrieval call binding the contract method 0xb7312591.
//
// Solidity: function getRequestByID(a address) constant returns(address)
func (_WriteRequestHolder *WriteRequestHolderCaller) GetRequestByID(opts *bind.CallOpts, a common.Address) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _WriteRequestHolder.contract.Call(opts, out, "getRequestByID", a)
	return *ret0, err
}

// GetRequestByID is a free data retrieval call binding the contract method 0xb7312591.
//
// Solidity: function getRequestByID(a address) constant returns(address)
func (_WriteRequestHolder *WriteRequestHolderSession) GetRequestByID(a common.Address) (common.Address, error) {
	return _WriteRequestHolder.Contract.GetRequestByID(&_WriteRequestHolder.CallOpts, a)
}

// GetRequestByID is a free data retrieval call binding the contract method 0xb7312591.
//
// Solidity: function getRequestByID(a address) constant returns(address)
func (_WriteRequestHolder *WriteRequestHolderCallerSession) GetRequestByID(a common.Address) (common.Address, error) {
	return _WriteRequestHolder.Contract.GetRequestByID(&_WriteRequestHolder.CallOpts, a)
}

// HasRequest is a free data retrieval call binding the contract method 0xe216ea2b.
//
// Solidity: function hasRequest( address) constant returns(bool)
func (_WriteRequestHolder *WriteRequestHolderCaller) HasRequest(opts *bind.CallOpts, arg0 common.Address) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _WriteRequestHolder.contract.Call(opts, out, "hasRequest", arg0)
	return *ret0, err
}

// HasRequest is a free data retrieval call binding the contract method 0xe216ea2b.
//
// Solidity: function hasRequest( address) constant returns(bool)
func (_WriteRequestHolder *WriteRequestHolderSession) HasRequest(arg0 common.Address) (bool, error) {
	return _WriteRequestHolder.Contract.HasRequest(&_WriteRequestHolder.CallOpts, arg0)
}

// HasRequest is a free data retrieval call binding the contract method 0xe216ea2b.
//
// Solidity: function hasRequest( address) constant returns(bool)
func (_WriteRequestHolder *WriteRequestHolderCallerSession) HasRequest(arg0 common.Address) (bool, error) {
	return _WriteRequestHolder.Contract.HasRequest(&_WriteRequestHolder.CallOpts, arg0)
}

// Wrmap is a free data retrieval call binding the contract method 0xed85db23.
//
// Solidity: function wrmap( address) constant returns(address)
func (_WriteRequestHolder *WriteRequestHolderCaller) Wrmap(opts *bind.CallOpts, arg0 common.Address) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _WriteRequestHolder.contract.Call(opts, out, "wrmap", arg0)
	return *ret0, err
}

// Wrmap is a free data retrieval call binding the contract method 0xed85db23.
//
// Solidity: function wrmap( address) constant returns(address)
func (_WriteRequestHolder *WriteRequestHolderSession) Wrmap(arg0 common.Address) (common.Address, error) {
	return _WriteRequestHolder.Contract.Wrmap(&_WriteRequestHolder.CallOpts, arg0)
}

// Wrmap is a free data retrieval call binding the contract method 0xed85db23.
//
// Solidity: function wrmap( address) constant returns(address)
func (_WriteRequestHolder *WriteRequestHolderCallerSession) Wrmap(arg0 common.Address) (common.Address, error) {
	return _WriteRequestHolder.Contract.Wrmap(&_WriteRequestHolder.CallOpts, arg0)
}

// AddWriteRequest is a paid mutator transaction binding the contract method 0x1029a017.
//
// Solidity: function addWriteRequest(a address) returns()
func (_WriteRequestHolder *WriteRequestHolderTransactor) AddWriteRequest(opts *bind.TransactOpts, a common.Address) (*types.Transaction, error) {
	return _WriteRequestHolder.contract.Transact(opts, "addWriteRequest", a)
}

// AddWriteRequest is a paid mutator transaction binding the contract method 0x1029a017.
//
// Solidity: function addWriteRequest(a address) returns()
func (_WriteRequestHolder *WriteRequestHolderSession) AddWriteRequest(a common.Address) (*types.Transaction, error) {
	return _WriteRequestHolder.Contract.AddWriteRequest(&_WriteRequestHolder.TransactOpts, a)
}

// AddWriteRequest is a paid mutator transaction binding the contract method 0x1029a017.
//
// Solidity: function addWriteRequest(a address) returns()
func (_WriteRequestHolder *WriteRequestHolderTransactorSession) AddWriteRequest(a common.Address) (*types.Transaction, error) {
	return _WriteRequestHolder.Contract.AddWriteRequest(&_WriteRequestHolder.TransactOpts, a)
}

// CanRead is a paid mutator transaction binding the contract method 0x42087d4f.
//
// Solidity: function canRead(a address) returns(bool)
func (_WriteRequestHolder *WriteRequestHolderTransactor) CanRead(opts *bind.TransactOpts, a common.Address) (*types.Transaction, error) {
	return _WriteRequestHolder.contract.Transact(opts, "canRead", a)
}

// CanRead is a paid mutator transaction binding the contract method 0x42087d4f.
//
// Solidity: function canRead(a address) returns(bool)
func (_WriteRequestHolder *WriteRequestHolderSession) CanRead(a common.Address) (*types.Transaction, error) {
	return _WriteRequestHolder.Contract.CanRead(&_WriteRequestHolder.TransactOpts, a)
}

// CanRead is a paid mutator transaction binding the contract method 0x42087d4f.
//
// Solidity: function canRead(a address) returns(bool)
func (_WriteRequestHolder *WriteRequestHolderTransactorSession) CanRead(a common.Address) (*types.Transaction, error) {
	return _WriteRequestHolder.Contract.CanRead(&_WriteRequestHolder.TransactOpts, a)
}

