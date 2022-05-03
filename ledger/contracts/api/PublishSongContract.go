// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package api

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

// PublishSongContractSong is an auto generated low-level Go binding around an user-defined struct.
type PublishSongContractSong struct {
	Cid                string
	Name               string
	Artists            []string
	Description        string
	PublishedTimestamp *big.Int
	Publisher          common.Address
	IsValid            bool
}

// ApiMetaData contains all meta data concerning the Api contract.
var ApiMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"_cid\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"_name\",\"type\":\"string\"},{\"internalType\":\"string[]\",\"name\":\"_artists\",\"type\":\"string[]\"},{\"internalType\":\"string\",\"name\":\"_description\",\"type\":\"string\"}],\"name\":\"addSong\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"cids\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getAllCid\",\"outputs\":[{\"internalType\":\"string[]\",\"name\":\"\",\"type\":\"string[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"_cid\",\"type\":\"string\"}],\"name\":\"getSong\",\"outputs\":[{\"components\":[{\"internalType\":\"string\",\"name\":\"cid\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"internalType\":\"string[]\",\"name\":\"artists\",\"type\":\"string[]\"},{\"internalType\":\"string\",\"name\":\"description\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"publishedTimestamp\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"publisher\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"isValid\",\"type\":\"bool\"}],\"internalType\":\"structPublishSongContract.Song\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totalSongs\",\"outputs\":[{\"internalType\":\"uint16\",\"name\":\"\",\"type\":\"uint16\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
	Bin: "0x60806040526002805461ffff1916905534801561001b57600080fd5b50600080546001600160a01b031916331790556110048061003d6000396000f3fe608060405234801561001057600080fd5b50600436106100625760003560e01c806350ef10c214610067578063704fe710146100905780637c522a8b146100b05780637ecaa061146100c55780638da5cb5b146100e6578063ea06c8c114610111575b600080fd5b61007a610075366004610b2c565b610126565b6040516100879190610bbb565b60405180910390f35b6100a361009e366004610caa565b6104e3565b6040516100879190610cc3565b6100b861058f565b6040516100879190610cdd565b6002546100d39061ffff1681565b60405161ffff9091168152602001610087565b6000546100f9906001600160a01b031681565b6040516001600160a01b039091168152602001610087565b61012461011f366004610d3f565b610668565b005b6101716040518060e00160405280606081526020016060815260200160608152602001606081526020016000815260200160006001600160a01b031681526020016000151581525090565b60018383604051610183929190610e34565b9081526040519081900360200190206005015460ff600160a01b909104166101f25760405162461bcd60e51b815260206004820152601760248201527f4e6f20736f6e67207769746820676976656e204349442e00000000000000000060448201526064015b60405180910390fd5b60018383604051610204929190610e34565b90815260200160405180910390206040518060e001604052908160008201805461022d90610e44565b80601f016020809104026020016040519081016040528092919081815260200182805461025990610e44565b80156102a65780601f1061027b576101008083540402835291602001916102a6565b820191906000526020600020905b81548152906001019060200180831161028957829003601f168201915b505050505081526020016001820180546102bf90610e44565b80601f01602080910402602001604051908101604052809291908181526020018280546102eb90610e44565b80156103385780601f1061030d57610100808354040283529160200191610338565b820191906000526020600020905b81548152906001019060200180831161031b57829003601f168201915b5050505050815260200160028201805480602002602001604051908101604052809291908181526020016000905b8282101561041257838290600052602060002001805461038590610e44565b80601f01602080910402602001604051908101604052809291908181526020018280546103b190610e44565b80156103fe5780601f106103d3576101008083540402835291602001916103fe565b820191906000526020600020905b8154815290600101906020018083116103e157829003601f168201915b505050505081526020019060010190610366565b50505050815260200160038201805461042a90610e44565b80601f016020809104026020016040519081016040528092919081815260200182805461045690610e44565b80156104a35780601f10610478576101008083540402835291602001916104a3565b820191906000526020600020905b81548152906001019060200180831161048657829003601f168201915b5050509183525050600482015460208201526005909101546001600160a01b0381166040830152600160a01b900460ff1615156060909101529392505050565b600381815481106104f357600080fd5b90600052602060002001600091509050805461050e90610e44565b80601f016020809104026020016040519081016040528092919081815260200182805461053a90610e44565b80156105875780601f1061055c57610100808354040283529160200191610587565b820191906000526020600020905b81548152906001019060200180831161056a57829003601f168201915b505050505081565b60606003805480602002602001604051908101604052809291908181526020016000905b8282101561065f5783829060005260206000200180546105d290610e44565b80601f01602080910402602001604051908101604052809291908181526020018280546105fe90610e44565b801561064b5780601f106106205761010080835404028352916020019161064b565b820191906000526020600020905b81548152906001019060200180831161062e57829003601f168201915b5050505050815260200190600101906105b3565b50505050905090565b6001888860405161067a929190610e34565b9081526040519081900360200190206005015460ff600160a01b90910416156106f05760405162461bcd60e51b815260206004820152602260248201527f536f6e67207769746820676976656e2043494420616c72656164792061646465604482015261321760f11b60648201526084016101e9565b600280546001919060009061070a90849061ffff16610e7e565b825461ffff9182166101009390930a92830291909202199091161790555060038054600181018255600091909152610765907fc2575a0e9e593c00f959f8c92f12db2869c3395a3b0502d05e2516446f71f85b018989610923565b506040518060e0016040528089898080601f016020809104026020016040519081016040528093929190818152602001838380828437600092019190915250505090825250604080516020601f8a01819004810282018101909252888152918101919089908990819084018382808284376000920191909152505050908252506020016107f28587610ef9565b815260200183838080601f016020809104026020016040519081016040528093929190818152602001838380828437600092019190915250505090825250426020820152336040808301919091526001606090920182905251610858908b908b90610e34565b908152602001604051809103902060008201518160000190805190602001906108829291906109a7565b50602082810151805161089b92600185019201906109a7565b50604082015180516108b7916002840191602090910190610a1b565b50606082015180516108d39160038401916020909101906109a7565b506080820151600482015560a08201516005909101805460c0909301511515600160a01b026001600160a81b03199093166001600160a01b03909216919091179190911790555050505050505050565b82805461092f90610e44565b90600052602060002090601f0160209004810192826109515760008555610997565b82601f1061096a5782800160ff19823516178555610997565b82800160010185558215610997579182015b8281111561099757823582559160200191906001019061097c565b506109a3929150610a74565b5090565b8280546109b390610e44565b90600052602060002090601f0160209004810192826109d55760008555610997565b82601f106109ee57805160ff1916838001178555610997565b82800160010185558215610997579182015b82811115610997578251825591602001919060010190610a00565b828054828255906000526020600020908101928215610a68579160200282015b82811115610a685782518051610a589184916020909101906109a7565b5091602001919060010190610a3b565b506109a3929150610a89565b5b808211156109a35760008155600101610a75565b808211156109a3576000610a9d8282610aa6565b50600101610a89565b508054610ab290610e44565b6000825580601f10610ac2575050565b601f016020900490600052602060002090810190610ae09190610a74565b50565b60008083601f840112610af557600080fd5b50813567ffffffffffffffff811115610b0d57600080fd5b602083019150836020828501011115610b2557600080fd5b9250929050565b60008060208385031215610b3f57600080fd5b823567ffffffffffffffff811115610b5657600080fd5b610b6285828601610ae3565b90969095509350505050565b6000815180845260005b81811015610b9457602081850181015186830182015201610b78565b81811115610ba6576000602083870101525b50601f01601f19169290920160200192915050565b60006020808352835160e082850152610bd8610100850182610b6e565b905081850151601f1980868403016040870152610bf58383610b6e565b60408801518782038301606089015280518083529194508501925084840190600581901b8501860160005b82811015610c4c5784878303018452610c3a828751610b6e565b95880195938801939150600101610c20565b5060608a01519650838982030160808a0152610c688188610b6e565b9650505050505050608084015160a084015260a0840151610c9460c08501826001600160a01b03169052565b5060c084015180151560e0850152509392505050565b600060208284031215610cbc57600080fd5b5035919050565b602081526000610cd66020830184610b6e565b9392505050565b6000602080830181845280855180835260408601915060408160051b870101925083870160005b82811015610d3257603f19888603018452610d20858351610b6e565b94509285019290850190600101610d04565b5092979650505050505050565b6000806000806000806000806080898b031215610d5b57600080fd5b883567ffffffffffffffff80821115610d7357600080fd5b610d7f8c838d01610ae3565b909a50985060208b0135915080821115610d9857600080fd5b610da48c838d01610ae3565b909850965060408b0135915080821115610dbd57600080fd5b818b0191508b601f830112610dd157600080fd5b813581811115610de057600080fd5b8c60208260051b8501011115610df557600080fd5b6020830196508095505060608b0135915080821115610e1357600080fd5b50610e208b828c01610ae3565b999c989b5096995094979396929594505050565b8183823760009101908152919050565b600181811c90821680610e5857607f821691505b602082108103610e7857634e487b7160e01b600052602260045260246000fd5b50919050565b600061ffff808316818516808303821115610ea957634e487b7160e01b600052601160045260246000fd5b01949350505050565b634e487b7160e01b600052604160045260246000fd5b604051601f8201601f1916810167ffffffffffffffff81118282101715610ef157610ef1610eb2565b604052919050565b600067ffffffffffffffff80841115610f1457610f14610eb2565b8360051b6020610f25818301610ec8565b868152918501918181019036841115610f3d57600080fd5b865b84811015610fc257803586811115610f575760008081fd5b8801601f3681830112610f6a5760008081fd5b813588811115610f7c57610f7c610eb2565b610f8d818301601f19168801610ec8565b91508082523687828501011115610fa45760008081fd5b80878401888401376000908201870152845250918301918301610f3f565b5097965050505050505056fea264697066735822122054bac9273fb0a95288ae3cfdb50fa8583e2f753a888c4d5d4d12020f57fc41d464736f6c634300080d0033",
}

// ApiABI is the input ABI used to generate the binding from.
// Deprecated: Use ApiMetaData.ABI instead.
var ApiABI = ApiMetaData.ABI

// ApiBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use ApiMetaData.Bin instead.
var ApiBin = ApiMetaData.Bin

// DeployApi deploys a new Ethereum contract, binding an instance of Api to it.
func DeployApi(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *Api, error) {
	parsed, err := ApiMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(ApiBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Api{ApiCaller: ApiCaller{contract: contract}, ApiTransactor: ApiTransactor{contract: contract}, ApiFilterer: ApiFilterer{contract: contract}}, nil
}

// Api is an auto generated Go binding around an Ethereum contract.
type Api struct {
	ApiCaller     // Read-only binding to the contract
	ApiTransactor // Write-only binding to the contract
	ApiFilterer   // Log filterer for contract events
}

// ApiCaller is an auto generated read-only Go binding around an Ethereum contract.
type ApiCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ApiTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ApiTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ApiFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ApiFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ApiSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ApiSession struct {
	Contract     *Api              // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ApiCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ApiCallerSession struct {
	Contract *ApiCaller    // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// ApiTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ApiTransactorSession struct {
	Contract     *ApiTransactor    // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ApiRaw is an auto generated low-level Go binding around an Ethereum contract.
type ApiRaw struct {
	Contract *Api // Generic contract binding to access the raw methods on
}

// ApiCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ApiCallerRaw struct {
	Contract *ApiCaller // Generic read-only contract binding to access the raw methods on
}

// ApiTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ApiTransactorRaw struct {
	Contract *ApiTransactor // Generic write-only contract binding to access the raw methods on
}

// NewApi creates a new instance of Api, bound to a specific deployed contract.
func NewApi(address common.Address, backend bind.ContractBackend) (*Api, error) {
	contract, err := bindApi(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Api{ApiCaller: ApiCaller{contract: contract}, ApiTransactor: ApiTransactor{contract: contract}, ApiFilterer: ApiFilterer{contract: contract}}, nil
}

// NewApiCaller creates a new read-only instance of Api, bound to a specific deployed contract.
func NewApiCaller(address common.Address, caller bind.ContractCaller) (*ApiCaller, error) {
	contract, err := bindApi(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ApiCaller{contract: contract}, nil
}

// NewApiTransactor creates a new write-only instance of Api, bound to a specific deployed contract.
func NewApiTransactor(address common.Address, transactor bind.ContractTransactor) (*ApiTransactor, error) {
	contract, err := bindApi(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ApiTransactor{contract: contract}, nil
}

// NewApiFilterer creates a new log filterer instance of Api, bound to a specific deployed contract.
func NewApiFilterer(address common.Address, filterer bind.ContractFilterer) (*ApiFilterer, error) {
	contract, err := bindApi(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ApiFilterer{contract: contract}, nil
}

// bindApi binds a generic wrapper to an already deployed contract.
func bindApi(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(ApiABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Api *ApiRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Api.Contract.ApiCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Api *ApiRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Api.Contract.ApiTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Api *ApiRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Api.Contract.ApiTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Api *ApiCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Api.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Api *ApiTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Api.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Api *ApiTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Api.Contract.contract.Transact(opts, method, params...)
}

// Cids is a free data retrieval call binding the contract method 0x704fe710.
//
// Solidity: function cids(uint256 ) view returns(string)
func (_Api *ApiCaller) Cids(opts *bind.CallOpts, arg0 *big.Int) (string, error) {
	var out []interface{}
	err := _Api.contract.Call(opts, &out, "cids", arg0)

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Cids is a free data retrieval call binding the contract method 0x704fe710.
//
// Solidity: function cids(uint256 ) view returns(string)
func (_Api *ApiSession) Cids(arg0 *big.Int) (string, error) {
	return _Api.Contract.Cids(&_Api.CallOpts, arg0)
}

// Cids is a free data retrieval call binding the contract method 0x704fe710.
//
// Solidity: function cids(uint256 ) view returns(string)
func (_Api *ApiCallerSession) Cids(arg0 *big.Int) (string, error) {
	return _Api.Contract.Cids(&_Api.CallOpts, arg0)
}

// GetAllCid is a free data retrieval call binding the contract method 0x7c522a8b.
//
// Solidity: function getAllCid() view returns(string[])
func (_Api *ApiCaller) GetAllCid(opts *bind.CallOpts) ([]string, error) {
	var out []interface{}
	err := _Api.contract.Call(opts, &out, "getAllCid")

	if err != nil {
		return *new([]string), err
	}

	out0 := *abi.ConvertType(out[0], new([]string)).(*[]string)

	return out0, err

}

// GetAllCid is a free data retrieval call binding the contract method 0x7c522a8b.
//
// Solidity: function getAllCid() view returns(string[])
func (_Api *ApiSession) GetAllCid() ([]string, error) {
	return _Api.Contract.GetAllCid(&_Api.CallOpts)
}

// GetAllCid is a free data retrieval call binding the contract method 0x7c522a8b.
//
// Solidity: function getAllCid() view returns(string[])
func (_Api *ApiCallerSession) GetAllCid() ([]string, error) {
	return _Api.Contract.GetAllCid(&_Api.CallOpts)
}

// GetSong is a free data retrieval call binding the contract method 0x50ef10c2.
//
// Solidity: function getSong(string _cid) view returns((string,string,string[],string,uint256,address,bool))
func (_Api *ApiCaller) GetSong(opts *bind.CallOpts, _cid string) (PublishSongContractSong, error) {
	var out []interface{}
	err := _Api.contract.Call(opts, &out, "getSong", _cid)

	if err != nil {
		return *new(PublishSongContractSong), err
	}

	out0 := *abi.ConvertType(out[0], new(PublishSongContractSong)).(*PublishSongContractSong)

	return out0, err

}

// GetSong is a free data retrieval call binding the contract method 0x50ef10c2.
//
// Solidity: function getSong(string _cid) view returns((string,string,string[],string,uint256,address,bool))
func (_Api *ApiSession) GetSong(_cid string) (PublishSongContractSong, error) {
	return _Api.Contract.GetSong(&_Api.CallOpts, _cid)
}

// GetSong is a free data retrieval call binding the contract method 0x50ef10c2.
//
// Solidity: function getSong(string _cid) view returns((string,string,string[],string,uint256,address,bool))
func (_Api *ApiCallerSession) GetSong(_cid string) (PublishSongContractSong, error) {
	return _Api.Contract.GetSong(&_Api.CallOpts, _cid)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Api *ApiCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Api.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Api *ApiSession) Owner() (common.Address, error) {
	return _Api.Contract.Owner(&_Api.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Api *ApiCallerSession) Owner() (common.Address, error) {
	return _Api.Contract.Owner(&_Api.CallOpts)
}

// TotalSongs is a free data retrieval call binding the contract method 0x7ecaa061.
//
// Solidity: function totalSongs() view returns(uint16)
func (_Api *ApiCaller) TotalSongs(opts *bind.CallOpts) (uint16, error) {
	var out []interface{}
	err := _Api.contract.Call(opts, &out, "totalSongs")

	if err != nil {
		return *new(uint16), err
	}

	out0 := *abi.ConvertType(out[0], new(uint16)).(*uint16)

	return out0, err

}

// TotalSongs is a free data retrieval call binding the contract method 0x7ecaa061.
//
// Solidity: function totalSongs() view returns(uint16)
func (_Api *ApiSession) TotalSongs() (uint16, error) {
	return _Api.Contract.TotalSongs(&_Api.CallOpts)
}

// TotalSongs is a free data retrieval call binding the contract method 0x7ecaa061.
//
// Solidity: function totalSongs() view returns(uint16)
func (_Api *ApiCallerSession) TotalSongs() (uint16, error) {
	return _Api.Contract.TotalSongs(&_Api.CallOpts)
}

// AddSong is a paid mutator transaction binding the contract method 0xea06c8c1.
//
// Solidity: function addSong(string _cid, string _name, string[] _artists, string _description) returns()
func (_Api *ApiTransactor) AddSong(opts *bind.TransactOpts, _cid string, _name string, _artists []string, _description string) (*types.Transaction, error) {
	return _Api.contract.Transact(opts, "addSong", _cid, _name, _artists, _description)
}

// AddSong is a paid mutator transaction binding the contract method 0xea06c8c1.
//
// Solidity: function addSong(string _cid, string _name, string[] _artists, string _description) returns()
func (_Api *ApiSession) AddSong(_cid string, _name string, _artists []string, _description string) (*types.Transaction, error) {
	return _Api.Contract.AddSong(&_Api.TransactOpts, _cid, _name, _artists, _description)
}

// AddSong is a paid mutator transaction binding the contract method 0xea06c8c1.
//
// Solidity: function addSong(string _cid, string _name, string[] _artists, string _description) returns()
func (_Api *ApiTransactorSession) AddSong(_cid string, _name string, _artists []string, _description string) (*types.Transaction, error) {
	return _Api.Contract.AddSong(&_Api.TransactOpts, _cid, _name, _artists, _description)
}
