// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package todo

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

// TodoTask is an auto generated low-level Go binding around an user-defined struct.
type TodoTask struct {
	Content string
	Statusl bool
	Author  string
}

// TodoMetaData contains all meta data concerning the Todo contract.
var TodoMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"_content\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"_author\",\"type\":\"string\"}],\"name\":\"add\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_id\",\"type\":\"uint256\"}],\"name\":\"get\",\"outputs\":[{\"components\":[{\"internalType\":\"string\",\"name\":\"content\",\"type\":\"string\"},{\"internalType\":\"bool\",\"name\":\"statusl\",\"type\":\"bool\"},{\"internalType\":\"string\",\"name\":\"author\",\"type\":\"string\"}],\"internalType\":\"structTodo.Task\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"list\",\"outputs\":[{\"components\":[{\"internalType\":\"string\",\"name\":\"content\",\"type\":\"string\"},{\"internalType\":\"bool\",\"name\":\"statusl\",\"type\":\"bool\"},{\"internalType\":\"string\",\"name\":\"author\",\"type\":\"string\"}],\"internalType\":\"structTodo.Task[]\",\"name\":\"\",\"type\":\"tuple[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_id\",\"type\":\"uint256\"}],\"name\":\"remove\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_id\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"_content\",\"type\":\"string\"}],\"name\":\"update\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Bin: "0x608060405234801561000f575f80fd5b503360015f6101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff160217905550610e728061005d5f395ff3fe608060405234801561000f575f80fd5b5060043610610060575f3560e01c80630f560cd7146100645780634cc82215146100825780638da5cb5b1461009e5780639507d39a146100bc578063ebdf86ca146100ec578063f745630f14610108575b5f80fd5b61006c610124565b60405161007991906107ef565b60405180910390f35b61009c60048036038101906100979190610853565b6102b8565b005b6100a661030d565b6040516100b391906108bd565b60405180910390f35b6100d660048036038101906100d19190610853565b610332565b6040516100e3919061092a565b60405180910390f35b61010660048036038101906101019190610a76565b6104a4565b005b610122600480360381019061011d9190610aec565b61058f565b005b60605f805480602002602001604051908101604052809291908181526020015f905b828210156102af578382905f5260205f2090600302016040518060600160405290815f8201805461017690610b73565b80601f01602080910402602001604051908101604052809291908181526020018280546101a290610b73565b80156101ed5780601f106101c4576101008083540402835291602001916101ed565b820191905f5260205f20905b8154815290600101906020018083116101d057829003601f168201915b50505050508152602001600182015f9054906101000a900460ff1615151515815260200160028201805461022090610b73565b80601f016020809104026020016040519081016040528092919081815260200182805461024c90610b73565b80156102975780601f1061026e57610100808354040283529160200191610297565b820191905f5260205f20905b81548152906001019060200180831161027a57829003601f168201915b50505050508152505081526020019060010190610146565b50505050905090565b5f81815481106102cb576102ca610ba3565b5b905f5260205f2090600302015f8082015f6102e691906105c2565b600182015f6101000a81549060ff0219169055600282015f61030891906105c2565b505050565b60015f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff1681565b61033a6105ff565b5f828154811061034d5761034c610ba3565b5b905f5260205f2090600302016040518060600160405290815f8201805461037390610b73565b80601f016020809104026020016040519081016040528092919081815260200182805461039f90610b73565b80156103ea5780601f106103c1576101008083540402835291602001916103ea565b820191905f5260205f20905b8154815290600101906020018083116103cd57829003601f168201915b50505050508152602001600182015f9054906101000a900460ff1615151515815260200160028201805461041d90610b73565b80601f016020809104026020016040519081016040528092919081815260200182805461044990610b73565b80156104945780601f1061046b57610100808354040283529160200191610494565b820191905f5260205f20905b81548152906001019060200180831161047757829003601f168201915b5050505050815250509050919050565b3373ffffffffffffffffffffffffffffffffffffffff1660015f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16146104fc575f80fd5b5f60405180606001604052808481526020015f1515815260200183815250908060018154018082558091505060019003905f5260205f2090600302015f909190919091505f820151815f0190816105539190610d6d565b506020820151816001015f6101000a81548160ff02191690831515021790555060408201518160020190816105889190610d6d565b5050505050565b805f83815481106105a3576105a2610ba3565b5b905f5260205f2090600302015f0190816105bd9190610d6d565b505050565b5080546105ce90610b73565b5f825580601f106105df57506105fc565b601f0160209004905f5260205f20908101906105fb9190610621565b5b50565b6040518060600160405280606081526020015f15158152602001606081525090565b5b80821115610638575f815f905550600101610622565b5090565b5f81519050919050565b5f82825260208201905092915050565b5f819050602082019050919050565b5f81519050919050565b5f82825260208201905092915050565b5f5b8381101561069c578082015181840152602081019050610681565b5f8484015250505050565b5f601f19601f8301169050919050565b5f6106c182610665565b6106cb818561066f565b93506106db81856020860161067f565b6106e4816106a7565b840191505092915050565b5f8115159050919050565b610703816106ef565b82525050565b5f606083015f8301518482035f86015261072382826106b7565b915050602083015161073860208601826106fa565b506040830151848203604086015261075082826106b7565b9150508091505092915050565b5f6107688383610709565b905092915050565b5f602082019050919050565b5f6107868261063c565b6107908185610646565b9350836020820285016107a285610656565b805f5b858110156107dd57848403895281516107be858261075d565b94506107c983610770565b925060208a019950506001810190506107a5565b50829750879550505050505092915050565b5f6020820190508181035f830152610807818461077c565b905092915050565b5f604051905090565b5f80fd5b5f80fd5b5f819050919050565b61083281610820565b811461083c575f80fd5b50565b5f8135905061084d81610829565b92915050565b5f6020828403121561086857610867610818565b5b5f6108758482850161083f565b91505092915050565b5f73ffffffffffffffffffffffffffffffffffffffff82169050919050565b5f6108a78261087e565b9050919050565b6108b78161089d565b82525050565b5f6020820190506108d05f8301846108ae565b92915050565b5f606083015f8301518482035f8601526108f082826106b7565b915050602083015161090560208601826106fa565b506040830151848203604086015261091d82826106b7565b9150508091505092915050565b5f6020820190508181035f83015261094281846108d6565b905092915050565b5f80fd5b5f80fd5b7f4e487b71000000000000000000000000000000000000000000000000000000005f52604160045260245ffd5b610988826106a7565b810181811067ffffffffffffffff821117156109a7576109a6610952565b5b80604052505050565b5f6109b961080f565b90506109c5828261097f565b919050565b5f67ffffffffffffffff8211156109e4576109e3610952565b5b6109ed826106a7565b9050602081019050919050565b828183375f83830152505050565b5f610a1a610a15846109ca565b6109b0565b905082815260208101848484011115610a3657610a3561094e565b5b610a418482856109fa565b509392505050565b5f82601f830112610a5d57610a5c61094a565b5b8135610a6d848260208601610a08565b91505092915050565b5f8060408385031215610a8c57610a8b610818565b5b5f83013567ffffffffffffffff811115610aa957610aa861081c565b5b610ab585828601610a49565b925050602083013567ffffffffffffffff811115610ad657610ad561081c565b5b610ae285828601610a49565b9150509250929050565b5f8060408385031215610b0257610b01610818565b5b5f610b0f8582860161083f565b925050602083013567ffffffffffffffff811115610b3057610b2f61081c565b5b610b3c85828601610a49565b9150509250929050565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52602260045260245ffd5b5f6002820490506001821680610b8a57607f821691505b602082108103610b9d57610b9c610b46565b5b50919050565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52603260045260245ffd5b5f819050815f5260205f209050919050565b5f6020601f8301049050919050565b5f82821b905092915050565b5f60088302610c2c7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff82610bf1565b610c368683610bf1565b95508019841693508086168417925050509392505050565b5f819050919050565b5f610c71610c6c610c6784610820565b610c4e565b610820565b9050919050565b5f819050919050565b610c8a83610c57565b610c9e610c9682610c78565b848454610bfd565b825550505050565b5f90565b610cb2610ca6565b610cbd818484610c81565b505050565b5b81811015610ce057610cd55f82610caa565b600181019050610cc3565b5050565b601f821115610d2557610cf681610bd0565b610cff84610be2565b81016020851015610d0e578190505b610d22610d1a85610be2565b830182610cc2565b50505b505050565b5f82821c905092915050565b5f610d455f1984600802610d2a565b1980831691505092915050565b5f610d5d8383610d36565b9150826002028217905092915050565b610d7682610665565b67ffffffffffffffff811115610d8f57610d8e610952565b5b610d998254610b73565b610da4828285610ce4565b5f60209050601f831160018114610dd5575f8415610dc3578287015190505b610dcd8582610d52565b865550610e34565b601f198416610de386610bd0565b5f5b82811015610e0a57848901518255600182019150602085019450602081019050610de5565b86831015610e275784890151610e23601f891682610d36565b8355505b6001600288020188555050505b50505050505056fea26469706673582212208798980fc029ef04142862aa5b4e6ef854f0a09ff117a828e1e938ceefcbd3e864736f6c63430008150033",
}

// TodoABI is the input ABI used to generate the binding from.
// Deprecated: Use TodoMetaData.ABI instead.
var TodoABI = TodoMetaData.ABI

// TodoBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use TodoMetaData.Bin instead.
var TodoBin = TodoMetaData.Bin

// DeployTodo deploys a new Ethereum contract, binding an instance of Todo to it.
func DeployTodo(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *Todo, error) {
	parsed, err := TodoMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(TodoBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Todo{TodoCaller: TodoCaller{contract: contract}, TodoTransactor: TodoTransactor{contract: contract}, TodoFilterer: TodoFilterer{contract: contract}}, nil
}

// Todo is an auto generated Go binding around an Ethereum contract.
type Todo struct {
	TodoCaller     // Read-only binding to the contract
	TodoTransactor // Write-only binding to the contract
	TodoFilterer   // Log filterer for contract events
}

// TodoCaller is an auto generated read-only Go binding around an Ethereum contract.
type TodoCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TodoTransactor is an auto generated write-only Go binding around an Ethereum contract.
type TodoTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TodoFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type TodoFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TodoSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type TodoSession struct {
	Contract     *Todo             // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// TodoCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type TodoCallerSession struct {
	Contract *TodoCaller   // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// TodoTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type TodoTransactorSession struct {
	Contract     *TodoTransactor   // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// TodoRaw is an auto generated low-level Go binding around an Ethereum contract.
type TodoRaw struct {
	Contract *Todo // Generic contract binding to access the raw methods on
}

// TodoCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type TodoCallerRaw struct {
	Contract *TodoCaller // Generic read-only contract binding to access the raw methods on
}

// TodoTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type TodoTransactorRaw struct {
	Contract *TodoTransactor // Generic write-only contract binding to access the raw methods on
}

// NewTodo creates a new instance of Todo, bound to a specific deployed contract.
func NewTodo(address common.Address, backend bind.ContractBackend) (*Todo, error) {
	contract, err := bindTodo(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Todo{TodoCaller: TodoCaller{contract: contract}, TodoTransactor: TodoTransactor{contract: contract}, TodoFilterer: TodoFilterer{contract: contract}}, nil
}

// NewTodoCaller creates a new read-only instance of Todo, bound to a specific deployed contract.
func NewTodoCaller(address common.Address, caller bind.ContractCaller) (*TodoCaller, error) {
	contract, err := bindTodo(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &TodoCaller{contract: contract}, nil
}

// NewTodoTransactor creates a new write-only instance of Todo, bound to a specific deployed contract.
func NewTodoTransactor(address common.Address, transactor bind.ContractTransactor) (*TodoTransactor, error) {
	contract, err := bindTodo(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &TodoTransactor{contract: contract}, nil
}

// NewTodoFilterer creates a new log filterer instance of Todo, bound to a specific deployed contract.
func NewTodoFilterer(address common.Address, filterer bind.ContractFilterer) (*TodoFilterer, error) {
	contract, err := bindTodo(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &TodoFilterer{contract: contract}, nil
}

// bindTodo binds a generic wrapper to an already deployed contract.
func bindTodo(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := TodoMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Todo *TodoRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Todo.Contract.TodoCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Todo *TodoRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Todo.Contract.TodoTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Todo *TodoRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Todo.Contract.TodoTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Todo *TodoCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Todo.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Todo *TodoTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Todo.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Todo *TodoTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Todo.Contract.contract.Transact(opts, method, params...)
}

// Get is a free data retrieval call binding the contract method 0x9507d39a.
//
// Solidity: function get(uint256 _id) view returns((string,bool,string))
func (_Todo *TodoCaller) Get(opts *bind.CallOpts, _id *big.Int) (TodoTask, error) {
	var out []interface{}
	err := _Todo.contract.Call(opts, &out, "get", _id)

	if err != nil {
		return *new(TodoTask), err
	}

	out0 := *abi.ConvertType(out[0], new(TodoTask)).(*TodoTask)

	return out0, err

}

// Get is a free data retrieval call binding the contract method 0x9507d39a.
//
// Solidity: function get(uint256 _id) view returns((string,bool,string))
func (_Todo *TodoSession) Get(_id *big.Int) (TodoTask, error) {
	return _Todo.Contract.Get(&_Todo.CallOpts, _id)
}

// Get is a free data retrieval call binding the contract method 0x9507d39a.
//
// Solidity: function get(uint256 _id) view returns((string,bool,string))
func (_Todo *TodoCallerSession) Get(_id *big.Int) (TodoTask, error) {
	return _Todo.Contract.Get(&_Todo.CallOpts, _id)
}

// List is a free data retrieval call binding the contract method 0x0f560cd7.
//
// Solidity: function list() view returns((string,bool,string)[])
func (_Todo *TodoCaller) List(opts *bind.CallOpts) ([]TodoTask, error) {
	var out []interface{}
	err := _Todo.contract.Call(opts, &out, "list")

	if err != nil {
		return *new([]TodoTask), err
	}

	out0 := *abi.ConvertType(out[0], new([]TodoTask)).(*[]TodoTask)

	return out0, err

}

// List is a free data retrieval call binding the contract method 0x0f560cd7.
//
// Solidity: function list() view returns((string,bool,string)[])
func (_Todo *TodoSession) List() ([]TodoTask, error) {
	return _Todo.Contract.List(&_Todo.CallOpts)
}

// List is a free data retrieval call binding the contract method 0x0f560cd7.
//
// Solidity: function list() view returns((string,bool,string)[])
func (_Todo *TodoCallerSession) List() ([]TodoTask, error) {
	return _Todo.Contract.List(&_Todo.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Todo *TodoCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Todo.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Todo *TodoSession) Owner() (common.Address, error) {
	return _Todo.Contract.Owner(&_Todo.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Todo *TodoCallerSession) Owner() (common.Address, error) {
	return _Todo.Contract.Owner(&_Todo.CallOpts)
}

// Add is a paid mutator transaction binding the contract method 0xebdf86ca.
//
// Solidity: function add(string _content, string _author) returns()
func (_Todo *TodoTransactor) Add(opts *bind.TransactOpts, _content string, _author string) (*types.Transaction, error) {
	return _Todo.contract.Transact(opts, "add", _content, _author)
}

// Add is a paid mutator transaction binding the contract method 0xebdf86ca.
//
// Solidity: function add(string _content, string _author) returns()
func (_Todo *TodoSession) Add(_content string, _author string) (*types.Transaction, error) {
	return _Todo.Contract.Add(&_Todo.TransactOpts, _content, _author)
}

// Add is a paid mutator transaction binding the contract method 0xebdf86ca.
//
// Solidity: function add(string _content, string _author) returns()
func (_Todo *TodoTransactorSession) Add(_content string, _author string) (*types.Transaction, error) {
	return _Todo.Contract.Add(&_Todo.TransactOpts, _content, _author)
}

// Remove is a paid mutator transaction binding the contract method 0x4cc82215.
//
// Solidity: function remove(uint256 _id) returns()
func (_Todo *TodoTransactor) Remove(opts *bind.TransactOpts, _id *big.Int) (*types.Transaction, error) {
	return _Todo.contract.Transact(opts, "remove", _id)
}

// Remove is a paid mutator transaction binding the contract method 0x4cc82215.
//
// Solidity: function remove(uint256 _id) returns()
func (_Todo *TodoSession) Remove(_id *big.Int) (*types.Transaction, error) {
	return _Todo.Contract.Remove(&_Todo.TransactOpts, _id)
}

// Remove is a paid mutator transaction binding the contract method 0x4cc82215.
//
// Solidity: function remove(uint256 _id) returns()
func (_Todo *TodoTransactorSession) Remove(_id *big.Int) (*types.Transaction, error) {
	return _Todo.Contract.Remove(&_Todo.TransactOpts, _id)
}

// Update is a paid mutator transaction binding the contract method 0xf745630f.
//
// Solidity: function update(uint256 _id, string _content) returns()
func (_Todo *TodoTransactor) Update(opts *bind.TransactOpts, _id *big.Int, _content string) (*types.Transaction, error) {
	return _Todo.contract.Transact(opts, "update", _id, _content)
}

// Update is a paid mutator transaction binding the contract method 0xf745630f.
//
// Solidity: function update(uint256 _id, string _content) returns()
func (_Todo *TodoSession) Update(_id *big.Int, _content string) (*types.Transaction, error) {
	return _Todo.Contract.Update(&_Todo.TransactOpts, _id, _content)
}

// Update is a paid mutator transaction binding the contract method 0xf745630f.
//
// Solidity: function update(uint256 _id, string _content) returns()
func (_Todo *TodoTransactorSession) Update(_id *big.Int, _content string) (*types.Transaction, error) {
	return _Todo.Contract.Update(&_Todo.TransactOpts, _id, _content)
}
