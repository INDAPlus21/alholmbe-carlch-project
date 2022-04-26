// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package UniswapQuery

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

// UniswapQueryMetaData contains all meta data concerning the UniswapQuery contract.
var UniswapQueryMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"name\":\"BadIndex\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"contractUniswapV2Factory\",\"name\":\"_uniswapFactory\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_startIndex\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_stopIndex\",\"type\":\"uint256\"}],\"name\":\"getPairsByRange\",\"outputs\":[{\"internalType\":\"address[3][]\",\"name\":\"\",\"type\":\"address[3][]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractIUniswapV2Pair[]\",\"name\":\"_pairs\",\"type\":\"address[]\"}],\"name\":\"getReservesByPairs\",\"outputs\":[{\"internalType\":\"uint256[3][]\",\"name\":\"\",\"type\":\"uint256[3][]\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
	Bin: "0x608060405234801561001057600080fd5b50600436106100365760003560e01c80634dbf0f391461003b578063f0a71cd614610064575b600080fd5b61004e610049366004610596565b610084565b60405161005b919061060b565b60405180910390f35b61007761007236600461068d565b610243565b60405161005b91906106c2565b606060008267ffffffffffffffff8111156100a1576100a1610728565b6040519080825280602002602001820160405280156100da57816020015b6100c7610578565b8152602001906001900390816100bf5790505b50905060005b8381101561023b578484828181106100fa576100fa61073e565b905060200201602081019061010f9190610754565b6001600160a01b0316630902f1ac6040518163ffffffff1660e01b8152600401606060405180830381865afa15801561014c573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906101709190610794565b826001600160701b03169250816001600160701b031691508063ffffffff1690508484815181106101a3576101a361073e565b60200260200101516000600381106101bd576101bd61073e565b602002018585815181106101d3576101d361073e565b60200260200101516001600381106101ed576101ed61073e565b602002018686815181106102035761020361073e565b602002602001015160026003811061021d5761021d61073e565b60200201929092529190525280610233816107fa565b9150506100e0565b509392505050565b60606000846001600160a01b031663574f2ba36040518163ffffffff1660e01b8152600401602060405180830381865afa158015610285573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906102a99190610813565b9050808311156102b7578092505b838310156102d85760405163779ffa5760e11b815260040160405180910390fd5b60006102e4858561082c565b905060008167ffffffffffffffff81111561030157610301610728565b60405190808252806020026020018201604052801561033a57816020015b610327610578565b81526020019060019003908161031f5790505b50905060005b8281101561056d5760006001600160a01b038916631e3dd18b610363848b610843565b6040518263ffffffff1660e01b815260040161038191815260200190565b602060405180830381865afa15801561039e573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906103c2919061085b565b9050806001600160a01b0316630dfe16816040518163ffffffff1660e01b8152600401602060405180830381865afa158015610402573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610426919061085b565b8383815181106104385761043861073e565b60200260200101516000600381106104525761045261073e565b60200201906001600160a01b031690816001600160a01b031681525050806001600160a01b031663d21220a76040518163ffffffff1660e01b8152600401602060405180830381865afa1580156104ad573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906104d1919061085b565b8383815181106104e3576104e361073e565b60200260200101516001600381106104fd576104fd61073e565b60200201906001600160a01b031690816001600160a01b0316815250508083838151811061052d5761052d61073e565b60200260200101516002600381106105475761054761073e565b6001600160a01b0390921660209290920201525080610565816107fa565b915050610340565b509695505050505050565b60405180606001604052806003906020820280368337509192915050565b600080602083850312156105a957600080fd5b823567ffffffffffffffff808211156105c157600080fd5b818501915085601f8301126105d557600080fd5b8135818111156105e457600080fd5b8660208260051b85010111156105f957600080fd5b60209290920196919550909350505050565b602080825282518282018190526000919084820190604085019084805b8281101561066857845184835b600381101561065257825182529188019190880190600101610635565b5050509385019360609390930192600101610628565b5091979650505050505050565b6001600160a01b038116811461068a57600080fd5b50565b6000806000606084860312156106a257600080fd5b83356106ad81610675565b95602085013595506040909401359392505050565b602080825282518282018190526000919084820190604085019084805b8281101561066857845184835b60038110156107125782516001600160a01b0316825291880191908801906001016106ec565b50505093850193606093909301926001016106df565b634e487b7160e01b600052604160045260246000fd5b634e487b7160e01b600052603260045260246000fd5b60006020828403121561076657600080fd5b813561077181610675565b9392505050565b80516001600160701b038116811461078f57600080fd5b919050565b6000806000606084860312156107a957600080fd5b6107b284610778565b92506107c060208501610778565b9150604084015163ffffffff811681146107d957600080fd5b809150509250925092565b634e487b7160e01b600052601160045260246000fd5b60006001820161080c5761080c6107e4565b5060010190565b60006020828403121561082557600080fd5b5051919050565b60008282101561083e5761083e6107e4565b500390565b60008219821115610856576108566107e4565b500190565b60006020828403121561086d57600080fd5b81516107718161067556fea2646970667358221220ef13f421ba0b735216d0ae246022a3eb082f7e364fc1681a5e1ef9cb0a9dcde964736f6c634300080d0033",
}

// UniswapQueryABI is the input ABI used to generate the binding from.
// Deprecated: Use UniswapQueryMetaData.ABI instead.
var UniswapQueryABI = UniswapQueryMetaData.ABI

// UniswapQueryBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use UniswapQueryMetaData.Bin instead.
var UniswapQueryBin = UniswapQueryMetaData.Bin

// DeployUniswapQuery deploys a new Ethereum contract, binding an instance of UniswapQuery to it.
func DeployUniswapQuery(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *UniswapQuery, error) {
	parsed, err := UniswapQueryMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(UniswapQueryBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &UniswapQuery{UniswapQueryCaller: UniswapQueryCaller{contract: contract}, UniswapQueryTransactor: UniswapQueryTransactor{contract: contract}, UniswapQueryFilterer: UniswapQueryFilterer{contract: contract}}, nil
}

// UniswapQuery is an auto generated Go binding around an Ethereum contract.
type UniswapQuery struct {
	UniswapQueryCaller     // Read-only binding to the contract
	UniswapQueryTransactor // Write-only binding to the contract
	UniswapQueryFilterer   // Log filterer for contract events
}

// UniswapQueryCaller is an auto generated read-only Go binding around an Ethereum contract.
type UniswapQueryCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// UniswapQueryTransactor is an auto generated write-only Go binding around an Ethereum contract.
type UniswapQueryTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// UniswapQueryFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type UniswapQueryFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// UniswapQuerySession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type UniswapQuerySession struct {
	Contract     *UniswapQuery     // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// UniswapQueryCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type UniswapQueryCallerSession struct {
	Contract *UniswapQueryCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts       // Call options to use throughout this session
}

// UniswapQueryTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type UniswapQueryTransactorSession struct {
	Contract     *UniswapQueryTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts       // Transaction auth options to use throughout this session
}

// UniswapQueryRaw is an auto generated low-level Go binding around an Ethereum contract.
type UniswapQueryRaw struct {
	Contract *UniswapQuery // Generic contract binding to access the raw methods on
}

// UniswapQueryCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type UniswapQueryCallerRaw struct {
	Contract *UniswapQueryCaller // Generic read-only contract binding to access the raw methods on
}

// UniswapQueryTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type UniswapQueryTransactorRaw struct {
	Contract *UniswapQueryTransactor // Generic write-only contract binding to access the raw methods on
}

// NewUniswapQuery creates a new instance of UniswapQuery, bound to a specific deployed contract.
func NewUniswapQuery(address common.Address, backend bind.ContractBackend) (*UniswapQuery, error) {
	contract, err := bindUniswapQuery(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &UniswapQuery{UniswapQueryCaller: UniswapQueryCaller{contract: contract}, UniswapQueryTransactor: UniswapQueryTransactor{contract: contract}, UniswapQueryFilterer: UniswapQueryFilterer{contract: contract}}, nil
}

// NewUniswapQueryCaller creates a new read-only instance of UniswapQuery, bound to a specific deployed contract.
func NewUniswapQueryCaller(address common.Address, caller bind.ContractCaller) (*UniswapQueryCaller, error) {
	contract, err := bindUniswapQuery(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &UniswapQueryCaller{contract: contract}, nil
}

// NewUniswapQueryTransactor creates a new write-only instance of UniswapQuery, bound to a specific deployed contract.
func NewUniswapQueryTransactor(address common.Address, transactor bind.ContractTransactor) (*UniswapQueryTransactor, error) {
	contract, err := bindUniswapQuery(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &UniswapQueryTransactor{contract: contract}, nil
}

// NewUniswapQueryFilterer creates a new log filterer instance of UniswapQuery, bound to a specific deployed contract.
func NewUniswapQueryFilterer(address common.Address, filterer bind.ContractFilterer) (*UniswapQueryFilterer, error) {
	contract, err := bindUniswapQuery(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &UniswapQueryFilterer{contract: contract}, nil
}

// bindUniswapQuery binds a generic wrapper to an already deployed contract.
func bindUniswapQuery(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(UniswapQueryABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_UniswapQuery *UniswapQueryRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _UniswapQuery.Contract.UniswapQueryCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_UniswapQuery *UniswapQueryRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _UniswapQuery.Contract.UniswapQueryTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_UniswapQuery *UniswapQueryRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _UniswapQuery.Contract.UniswapQueryTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_UniswapQuery *UniswapQueryCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _UniswapQuery.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_UniswapQuery *UniswapQueryTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _UniswapQuery.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_UniswapQuery *UniswapQueryTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _UniswapQuery.Contract.contract.Transact(opts, method, params...)
}

// GetPairsByRange is a free data retrieval call binding the contract method 0xf0a71cd6.
//
// Solidity: function getPairsByRange(address _uniswapFactory, uint256 _startIndex, uint256 _stopIndex) view returns(address[3][])
func (_UniswapQuery *UniswapQueryCaller) GetPairsByRange(opts *bind.CallOpts, _uniswapFactory common.Address, _startIndex *big.Int, _stopIndex *big.Int) ([][3]common.Address, error) {
	var out []interface{}
	err := _UniswapQuery.contract.Call(opts, &out, "getPairsByRange", _uniswapFactory, _startIndex, _stopIndex)

	if err != nil {
		return *new([][3]common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new([][3]common.Address)).(*[][3]common.Address)

	return out0, err

}

// GetPairsByRange is a free data retrieval call binding the contract method 0xf0a71cd6.
//
// Solidity: function getPairsByRange(address _uniswapFactory, uint256 _startIndex, uint256 _stopIndex) view returns(address[3][])
func (_UniswapQuery *UniswapQuerySession) GetPairsByRange(_uniswapFactory common.Address, _startIndex *big.Int, _stopIndex *big.Int) ([][3]common.Address, error) {
	return _UniswapQuery.Contract.GetPairsByRange(&_UniswapQuery.CallOpts, _uniswapFactory, _startIndex, _stopIndex)
}

// GetPairsByRange is a free data retrieval call binding the contract method 0xf0a71cd6.
//
// Solidity: function getPairsByRange(address _uniswapFactory, uint256 _startIndex, uint256 _stopIndex) view returns(address[3][])
func (_UniswapQuery *UniswapQueryCallerSession) GetPairsByRange(_uniswapFactory common.Address, _startIndex *big.Int, _stopIndex *big.Int) ([][3]common.Address, error) {
	return _UniswapQuery.Contract.GetPairsByRange(&_UniswapQuery.CallOpts, _uniswapFactory, _startIndex, _stopIndex)
}

// GetReservesByPairs is a free data retrieval call binding the contract method 0x4dbf0f39.
//
// Solidity: function getReservesByPairs(address[] _pairs) view returns(uint256[3][])
func (_UniswapQuery *UniswapQueryCaller) GetReservesByPairs(opts *bind.CallOpts, _pairs []common.Address) ([][3]*big.Int, error) {
	var out []interface{}
	err := _UniswapQuery.contract.Call(opts, &out, "getReservesByPairs", _pairs)

	if err != nil {
		return *new([][3]*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new([][3]*big.Int)).(*[][3]*big.Int)

	return out0, err

}

// GetReservesByPairs is a free data retrieval call binding the contract method 0x4dbf0f39.
//
// Solidity: function getReservesByPairs(address[] _pairs) view returns(uint256[3][])
func (_UniswapQuery *UniswapQuerySession) GetReservesByPairs(_pairs []common.Address) ([][3]*big.Int, error) {
	return _UniswapQuery.Contract.GetReservesByPairs(&_UniswapQuery.CallOpts, _pairs)
}

// GetReservesByPairs is a free data retrieval call binding the contract method 0x4dbf0f39.
//
// Solidity: function getReservesByPairs(address[] _pairs) view returns(uint256[3][])
func (_UniswapQuery *UniswapQueryCallerSession) GetReservesByPairs(_pairs []common.Address) ([][3]*big.Int, error) {
	return _UniswapQuery.Contract.GetReservesByPairs(&_UniswapQuery.CallOpts, _pairs)
}
