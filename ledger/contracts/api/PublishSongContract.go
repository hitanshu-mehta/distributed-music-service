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
	ABI: "[{\"inputs\":[{\"internalType\":\"string\",\"name\":\"_cid\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"_name\",\"type\":\"string\"},{\"internalType\":\"string[]\",\"name\":\"_artists\",\"type\":\"string[]\"},{\"internalType\":\"string\",\"name\":\"_description\",\"type\":\"string\"}],\"name\":\"addSong\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"cids\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getAllCid\",\"outputs\":[{\"internalType\":\"string[]\",\"name\":\"\",\"type\":\"string[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"_cid\",\"type\":\"string\"}],\"name\":\"getSong\",\"outputs\":[{\"components\":[{\"internalType\":\"string\",\"name\":\"cid\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"internalType\":\"string[]\",\"name\":\"artists\",\"type\":\"string[]\"},{\"internalType\":\"string\",\"name\":\"description\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"publishedTimestamp\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"publisher\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"isValid\",\"type\":\"bool\"}],\"internalType\":\"structPublishSongContract.Song\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totalSongs\",\"outputs\":[{\"internalType\":\"uint16\",\"name\":\"\",\"type\":\"uint16\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
	Bin: "0x60806040526001805461ffff1916905534801561001b57600080fd5b50610fcd8061002b6000396000f3fe608060405234801561001057600080fd5b50600436106100575760003560e01c806350ef10c21461005c578063704fe710146100855780637c522a8b146100a55780637ecaa061146100ba578063ea06c8c1146100db575b600080fd5b61006f61006a366004610af5565b6100f0565b60405161007c9190610b84565b60405180910390f35b610098610093366004610c73565b6104ad565b60405161007c9190610c8c565b6100ad610559565b60405161007c9190610ca6565b6001546100c89061ffff1681565b60405161ffff909116815260200161007c565b6100ee6100e9366004610d08565b610632565b005b61013b6040518060e00160405280606081526020016060815260200160608152602001606081526020016000815260200160006001600160a01b031681526020016000151581525090565b6000838360405161014d929190610dfd565b9081526040519081900360200190206005015460ff600160a01b909104166101bc5760405162461bcd60e51b815260206004820152601760248201527f4e6f20736f6e67207769746820676976656e204349442e00000000000000000060448201526064015b60405180910390fd5b600083836040516101ce929190610dfd565b90815260200160405180910390206040518060e00160405290816000820180546101f790610e0d565b80601f016020809104026020016040519081016040528092919081815260200182805461022390610e0d565b80156102705780601f1061024557610100808354040283529160200191610270565b820191906000526020600020905b81548152906001019060200180831161025357829003601f168201915b5050505050815260200160018201805461028990610e0d565b80601f01602080910402602001604051908101604052809291908181526020018280546102b590610e0d565b80156103025780601f106102d757610100808354040283529160200191610302565b820191906000526020600020905b8154815290600101906020018083116102e557829003601f168201915b5050505050815260200160028201805480602002602001604051908101604052809291908181526020016000905b828210156103dc57838290600052602060002001805461034f90610e0d565b80601f016020809104026020016040519081016040528092919081815260200182805461037b90610e0d565b80156103c85780601f1061039d576101008083540402835291602001916103c8565b820191906000526020600020905b8154815290600101906020018083116103ab57829003601f168201915b505050505081526020019060010190610330565b5050505081526020016003820180546103f490610e0d565b80601f016020809104026020016040519081016040528092919081815260200182805461042090610e0d565b801561046d5780601f106104425761010080835404028352916020019161046d565b820191906000526020600020905b81548152906001019060200180831161045057829003601f168201915b5050509183525050600482015460208201526005909101546001600160a01b0381166040830152600160a01b900460ff1615156060909101529392505050565b600281815481106104bd57600080fd5b9060005260206000200160009150905080546104d890610e0d565b80601f016020809104026020016040519081016040528092919081815260200182805461050490610e0d565b80156105515780601f1061052657610100808354040283529160200191610551565b820191906000526020600020905b81548152906001019060200180831161053457829003601f168201915b505050505081565b60606002805480602002602001604051908101604052809291908181526020016000905b8282101561062957838290600052602060002001805461059c90610e0d565b80601f01602080910402602001604051908101604052809291908181526020018280546105c890610e0d565b80156106155780601f106105ea57610100808354040283529160200191610615565b820191906000526020600020905b8154815290600101906020018083116105f857829003601f168201915b50505050508152602001906001019061057d565b50505050905090565b60008888604051610644929190610dfd565b9081526040519081900360200190206005015460ff600160a01b90910416156106ba5760405162461bcd60e51b815260206004820152602260248201527f536f6e67207769746820676976656e2043494420616c72656164792061646465604482015261321760f11b60648201526084016101b3565b6001805481906000906106d290839061ffff16610e47565b825461ffff9182166101009390930a9283029190920219909116179055506002805460018101825560009190915261072d907f405787fa12a823e0f2b7631cc41b3ba8828b3321ca811111fa75cd3aa3bb5ace0189896108ec565b506040518060e0016040528089898080601f016020809104026020016040519081016040528093929190818152602001838380828437600092019190915250505090825250604080516020601f8a01819004810282018101909252888152918101919089908990819084018382808284376000920191909152505050908252506020016107ba8587610ec2565b815260200183838080601f01602080910402602001604051908101604052809392919081815260200183838082843760009201829052509385525050426020840152503360408084019190915260016060909301929092529051610821908b908b90610dfd565b9081526020016040518091039020600082015181600001908051906020019061084b929190610970565b5060208281015180516108649260018501920190610970565b50604082015180516108809160028401916020909101906109e4565b506060820151805161089c916003840191602090910190610970565b506080820151600482015560a08201516005909101805460c0909301511515600160a01b026001600160a81b03199093166001600160a01b03909216919091179190911790555050505050505050565b8280546108f890610e0d565b90600052602060002090601f01602090048101928261091a5760008555610960565b82601f106109335782800160ff19823516178555610960565b82800160010185558215610960579182015b82811115610960578235825591602001919060010190610945565b5061096c929150610a3d565b5090565b82805461097c90610e0d565b90600052602060002090601f01602090048101928261099e5760008555610960565b82601f106109b757805160ff1916838001178555610960565b82800160010185558215610960579182015b828111156109605782518255916020019190600101906109c9565b828054828255906000526020600020908101928215610a31579160200282015b82811115610a315782518051610a21918491602090910190610970565b5091602001919060010190610a04565b5061096c929150610a52565b5b8082111561096c5760008155600101610a3e565b8082111561096c576000610a668282610a6f565b50600101610a52565b508054610a7b90610e0d565b6000825580601f10610a8b575050565b601f016020900490600052602060002090810190610aa99190610a3d565b50565b60008083601f840112610abe57600080fd5b50813567ffffffffffffffff811115610ad657600080fd5b602083019150836020828501011115610aee57600080fd5b9250929050565b60008060208385031215610b0857600080fd5b823567ffffffffffffffff811115610b1f57600080fd5b610b2b85828601610aac565b90969095509350505050565b6000815180845260005b81811015610b5d57602081850181015186830182015201610b41565b81811115610b6f576000602083870101525b50601f01601f19169290920160200192915050565b60006020808352835160e082850152610ba1610100850182610b37565b905081850151601f1980868403016040870152610bbe8383610b37565b60408801518782038301606089015280518083529194508501925084840190600581901b8501860160005b82811015610c155784878303018452610c03828751610b37565b95880195938801939150600101610be9565b5060608a01519650838982030160808a0152610c318188610b37565b9650505050505050608084015160a084015260a0840151610c5d60c08501826001600160a01b03169052565b5060c084015180151560e0850152509392505050565b600060208284031215610c8557600080fd5b5035919050565b602081526000610c9f6020830184610b37565b9392505050565b6000602080830181845280855180835260408601915060408160051b870101925083870160005b82811015610cfb57603f19888603018452610ce9858351610b37565b94509285019290850190600101610ccd565b5092979650505050505050565b6000806000806000806000806080898b031215610d2457600080fd5b883567ffffffffffffffff80821115610d3c57600080fd5b610d488c838d01610aac565b909a50985060208b0135915080821115610d6157600080fd5b610d6d8c838d01610aac565b909850965060408b0135915080821115610d8657600080fd5b818b0191508b601f830112610d9a57600080fd5b813581811115610da957600080fd5b8c60208260051b8501011115610dbe57600080fd5b6020830196508095505060608b0135915080821115610ddc57600080fd5b50610de98b828c01610aac565b999c989b5096995094979396929594505050565b8183823760009101908152919050565b600181811c90821680610e2157607f821691505b602082108103610e4157634e487b7160e01b600052602260045260246000fd5b50919050565b600061ffff808316818516808303821115610e7257634e487b7160e01b600052601160045260246000fd5b01949350505050565b634e487b7160e01b600052604160045260246000fd5b604051601f8201601f1916810167ffffffffffffffff81118282101715610eba57610eba610e7b565b604052919050565b600067ffffffffffffffff80841115610edd57610edd610e7b565b8360051b6020610eee818301610e91565b868152918501918181019036841115610f0657600080fd5b865b84811015610f8b57803586811115610f205760008081fd5b8801601f3681830112610f335760008081fd5b813588811115610f4557610f45610e7b565b610f56818301601f19168801610e91565b91508082523687828501011115610f6d5760008081fd5b80878401888401376000908201870152845250918301918301610f08565b5097965050505050505056fea2646970667358221220823bcf8019753e8d4551e7d7ebc1ee95cee5c7ccc95fac3fc68b139224122a6064736f6c634300080d0033",
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
