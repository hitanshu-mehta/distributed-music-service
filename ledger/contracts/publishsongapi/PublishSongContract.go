// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package publishsongapi

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

// Struct0 is an auto generated low-level Go binding around an user-defined struct.
type Struct0 struct {
	Cid                string
	Name               string
	Artists            []string
	Description        string
	PublishedTimestamp *big.Int
	Publisher          common.Address
	IsValid            bool
}

// PublishsongapiMetaData contains all meta data concerning the Publishsongapi contract.
var PublishsongapiMetaData = &bind.MetaData{
	ABI: "[{\"constant\":true,\"inputs\":[{\"name\":\"_cid\",\"type\":\"string\"}],\"name\":\"getSong\",\"outputs\":[{\"components\":[{\"name\":\"cid\",\"type\":\"string\"},{\"name\":\"name\",\"type\":\"string\"},{\"name\":\"artists\",\"type\":\"string[]\"},{\"name\":\"description\",\"type\":\"string\"},{\"name\":\"publishedTimestamp\",\"type\":\"uint256\"},{\"name\":\"publisher\",\"type\":\"address\"},{\"name\":\"isValid\",\"type\":\"bool\"}],\"name\":\"\",\"type\":\"tuple\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"cids\",\"outputs\":[{\"name\":\"\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getAllCid\",\"outputs\":[{\"name\":\"\",\"type\":\"string[]\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"totalSongs\",\"outputs\":[{\"name\":\"\",\"type\":\"uint16\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_cid\",\"type\":\"string\"},{\"name\":\"_name\",\"type\":\"string\"},{\"name\":\"_artists\",\"type\":\"string[]\"},{\"name\":\"_description\",\"type\":\"string\"}],\"name\":\"addSong\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"}]",
	Bin: "0x60806040526002805461ffff1916905534801561001b57600080fd5b50600080546001600160a01b03191633179055610f9d8061003d6000396000f3fe608060405234801561001057600080fd5b50600436106100625760003560e01c806350ef10c214610067578063704fe710146100905780637c522a8b146100b05780637ecaa061146100c55780638da5cb5b146100da578063ea06c8c1146100ef575b600080fd5b61007a610075366004610a48565b610104565b6040516100879190610e53565b60405180910390f35b6100a361009e366004610b45565b61045b565b6040516100879190610e22565b6100b8610502565b6040516100879190610e11565b6100cd6105db565b6040516100879190610e64565b6100e26105e5565b6040516100879190610dfd565b6101026100fd366004610a85565b6105f4565b005b61010c6107d7565b60018260405161011c9190610df1565b9081526040519081900360200190206005015460ff600160a01b90910416151561016457604051600160e51b62461bcd02815260040161015b90610e43565b60405180910390fd5b6001826040516101749190610df1565b90815260408051602092819003830181208054600260018216156101009081026000190190921604601f81018690049095028301810190935260e0820184815291939092849291849184018282801561020e5780601f106101e35761010080835404028352916020019161020e565b820191906000526020600020905b8154815290600101906020018083116101f157829003601f168201915b50505050508152602001600182018054600181600116156101000203166002900480601f0160208091040260200160405190810160405280929190818152602001828054600181600116156101000203166002900480156102b05780601f10610285576101008083540402835291602001916102b0565b820191906000526020600020905b81548152906001019060200180831161029357829003601f168201915b5050505050815260200160028201805480602002602001604051908101604052809291908181526020016000905b828210156103895760008481526020908190208301805460408051601f60026000196101006001871615020190941693909304928301859004850281018501909152818152928301828280156103755780601f1061034a57610100808354040283529160200191610375565b820191906000526020600020905b81548152906001019060200180831161035857829003601f168201915b5050505050815260200190600101906102de565b5050509082525060038201805460408051602060026001851615610100026000190190941693909304601f810184900484028201840190925281815293820193929183018282801561041c5780601f106103f15761010080835404028352916020019161041c565b820191906000526020600020905b8154815290600101906020018083116103ff57829003601f168201915b5050509183525050600482015460208201526005909101546001600160a01b0381166040830152600160a01b900460ff16151560609091015292915050565b600380548290811061046957fe5b600091825260209182902001805460408051601f60026000196101006001871615020190941693909304928301859004850281018501909152818152935090918301828280156104fa5780601f106104cf576101008083540402835291602001916104fa565b820191906000526020600020905b8154815290600101906020018083116104dd57829003601f168201915b505050505081565b60606003805480602002602001604051908101604052809291908181526020016000905b828210156105d15760008481526020908190208301805460408051601f60026000196101006001871615020190941693909304928301859004850281018501909152818152928301828280156105bd5780601f10610592576101008083540402835291602001916105bd565b820191906000526020600020905b8154815290600101906020018083116105a057829003601f168201915b505050505081526020019060010190610526565b5050505090505b90565b60025461ffff1681565b6000546001600160a01b031681565b6001846040516106049190610df1565b9081526040519081900360200190206005015460ff600160a01b909104161561064257604051600160e51b62461bcd02815260040161015b90610e33565b6002805461ffff198116600161ffff92831681019092161790915560038054918201808255600091909152855190916106a4917fc2575a0e9e593c00f959f8c92f12db2869c3395a3b0502d05e2516446f71f85b90910190602088019061081f565b50506040518060e00160405280858152602001848152602001838152602001828152602001428152602001336001600160a01b03168152602001600115158152506001856040516106f59190610df1565b9081526020016040518091039020600082015181600001908051906020019061071f92919061081f565b506020828101518051610738926001850192019061081f565b506040820151805161075491600284019160209091019061089d565b506060820151805161077091600384019160209091019061081f565b506080820151600482015560a08201516005909101805460c0909301511515600160a01b0274ff0000000000000000000000000000000000000000196001600160a01b039093166001600160a01b0319909416939093179190911691909117905550505050565b6040518060e00160405280606081526020016060815260200160608152602001606081526020016000815260200160006001600160a01b031681526020016000151581525090565b828054600181600116156101000203166002900490600052602060002090601f016020900481019282601f1061086057805160ff191683800117855561088d565b8280016001018555821561088d579182015b8281111561088d578251825591602001919060010190610872565b506108999291506108f6565b5090565b8280548282559060005260206000209081019282156108ea579160200282015b828111156108ea57825180516108da91849160209091019061081f565b50916020019190600101906108bd565b50610899929150610910565b6105d891905b8082111561089957600081556001016108fc565b6105d891905b8082111561089957600061092a8282610933565b50600101610916565b50805460018160011615610100020316600290046000825580601f106109595750610977565b601f01602090049060005260206000209081019061097791906108f6565b50565b6000601f8201831361098b57600080fd5b813561099e61099982610e99565b610e72565b81815260209384019390925082018360005b838110156109dc57813586016109c688826109e6565b84525060209283019291909101906001016109b0565b5050505092915050565b6000601f820183136109f757600080fd5b8135610a0561099982610eba565b91508082526020830160208301858383011115610a2157600080fd5b610a2c838284610f1d565b50505092915050565b6000610a4182356105d8565b9392505050565b600060208284031215610a5a57600080fd5b813567ffffffffffffffff811115610a7157600080fd5b610a7d848285016109e6565b949350505050565b60008060008060808587031215610a9b57600080fd5b843567ffffffffffffffff811115610ab257600080fd5b610abe878288016109e6565b945050602085013567ffffffffffffffff811115610adb57600080fd5b610ae7878288016109e6565b935050604085013567ffffffffffffffff811115610b0457600080fd5b610b108782880161097a565b925050606085013567ffffffffffffffff811115610b2d57600080fd5b610b39878288016109e6565b91505092959194509250565b600060208284031215610b5757600080fd5b6000610a7d8484610a35565b6000610a418383610c7e565b610b7881610efa565b82525050565b6000610b8982610ee8565b610b938185610eec565b935083602082028501610ba585610ee2565b60005b84811015610bdc578383038852610bc0838351610b63565b9250610bcb82610ee2565b602098909801979150600101610ba8565b50909695505050505050565b6000610bf382610ee8565b610bfd8185610eec565b935083602082028501610c0f85610ee2565b60005b84811015610bdc578383038852610c2a838351610b63565b9250610c3582610ee2565b602098909801979150600101610c12565b610b7881610f05565b6000610c5a82610ee8565b610c648185610ef5565b9350610c74818560208601610f29565b9290920192915050565b6000610c8982610ee8565b610c938185610eec565b9350610ca3818560208601610f29565b610cac81610f59565b9093019392505050565b6000610cc3602283610eec565b7f536f6e67207769746820676976656e2043494420616c726561647920616464658152600160f11b61321702602082015260400192915050565b6000610d0a601783610eec565b7f4e6f20736f6e67207769746820676976656e204349442e000000000000000000815260200192915050565b805160e080845260009190840190610d4e8282610c7e565b91505060208301518482036020860152610d688282610c7e565b91505060408301518482036040860152610d828282610be8565b91505060608301518482036060860152610d9c8282610c7e565b9150506080830151610db16080860182610de8565b5060a0830151610dc460a0860182610b6f565b5060c0830151610dd760c0860182610c46565b509392505050565b610b7881610f0a565b610b78816105d8565b6000610a418284610c4f565b60208101610e0b8284610b6f565b92915050565b60208082528101610a418184610b7e565b60208082528101610a418184610c7e565b60208082528101610e0b81610cb6565b60208082528101610e0b81610cfd565b60208082528101610a418184610d36565b60208101610e0b8284610ddf565b60405181810167ffffffffffffffff81118282101715610e9157600080fd5b604052919050565b600067ffffffffffffffff821115610eb057600080fd5b5060209081020190565b600067ffffffffffffffff821115610ed157600080fd5b506020601f91909101601f19160190565b60200190565b5190565b90815260200190565b919050565b6000610e0b82610f11565b151590565b61ffff1690565b6001600160a01b031690565b82818337506000910152565b60005b83811015610f44578181015183820152602001610f2c565b83811115610f53576000848401525b50505050565b601f01601f19169056fea265627a7a723058206fb3c94b462190daa6c252f1bdb6c6a07df49246ce1394ee8fb5f9b9c32d95496c6578706572696d656e74616cf50037",
}

// PublishsongapiABI is the input ABI used to generate the binding from.
// Deprecated: Use PublishsongapiMetaData.ABI instead.
var PublishsongapiABI = PublishsongapiMetaData.ABI

// PublishsongapiBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use PublishsongapiMetaData.Bin instead.
var PublishsongapiBin = PublishsongapiMetaData.Bin

// DeployPublishsongapi deploys a new Ethereum contract, binding an instance of Publishsongapi to it.
func DeployPublishsongapi(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *Publishsongapi, error) {
	parsed, err := PublishsongapiMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(PublishsongapiBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Publishsongapi{PublishsongapiCaller: PublishsongapiCaller{contract: contract}, PublishsongapiTransactor: PublishsongapiTransactor{contract: contract}, PublishsongapiFilterer: PublishsongapiFilterer{contract: contract}}, nil
}

// Publishsongapi is an auto generated Go binding around an Ethereum contract.
type Publishsongapi struct {
	PublishsongapiCaller     // Read-only binding to the contract
	PublishsongapiTransactor // Write-only binding to the contract
	PublishsongapiFilterer   // Log filterer for contract events
}

// PublishsongapiCaller is an auto generated read-only Go binding around an Ethereum contract.
type PublishsongapiCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PublishsongapiTransactor is an auto generated write-only Go binding around an Ethereum contract.
type PublishsongapiTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PublishsongapiFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type PublishsongapiFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PublishsongapiSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type PublishsongapiSession struct {
	Contract     *Publishsongapi   // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// PublishsongapiCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type PublishsongapiCallerSession struct {
	Contract *PublishsongapiCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts         // Call options to use throughout this session
}

// PublishsongapiTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type PublishsongapiTransactorSession struct {
	Contract     *PublishsongapiTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts         // Transaction auth options to use throughout this session
}

// PublishsongapiRaw is an auto generated low-level Go binding around an Ethereum contract.
type PublishsongapiRaw struct {
	Contract *Publishsongapi // Generic contract binding to access the raw methods on
}

// PublishsongapiCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type PublishsongapiCallerRaw struct {
	Contract *PublishsongapiCaller // Generic read-only contract binding to access the raw methods on
}

// PublishsongapiTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type PublishsongapiTransactorRaw struct {
	Contract *PublishsongapiTransactor // Generic write-only contract binding to access the raw methods on
}

// NewPublishsongapi creates a new instance of Publishsongapi, bound to a specific deployed contract.
func NewPublishsongapi(address common.Address, backend bind.ContractBackend) (*Publishsongapi, error) {
	contract, err := bindPublishsongapi(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Publishsongapi{PublishsongapiCaller: PublishsongapiCaller{contract: contract}, PublishsongapiTransactor: PublishsongapiTransactor{contract: contract}, PublishsongapiFilterer: PublishsongapiFilterer{contract: contract}}, nil
}

// NewPublishsongapiCaller creates a new read-only instance of Publishsongapi, bound to a specific deployed contract.
func NewPublishsongapiCaller(address common.Address, caller bind.ContractCaller) (*PublishsongapiCaller, error) {
	contract, err := bindPublishsongapi(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &PublishsongapiCaller{contract: contract}, nil
}

// NewPublishsongapiTransactor creates a new write-only instance of Publishsongapi, bound to a specific deployed contract.
func NewPublishsongapiTransactor(address common.Address, transactor bind.ContractTransactor) (*PublishsongapiTransactor, error) {
	contract, err := bindPublishsongapi(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &PublishsongapiTransactor{contract: contract}, nil
}

// NewPublishsongapiFilterer creates a new log filterer instance of Publishsongapi, bound to a specific deployed contract.
func NewPublishsongapiFilterer(address common.Address, filterer bind.ContractFilterer) (*PublishsongapiFilterer, error) {
	contract, err := bindPublishsongapi(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &PublishsongapiFilterer{contract: contract}, nil
}

// bindPublishsongapi binds a generic wrapper to an already deployed contract.
func bindPublishsongapi(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(PublishsongapiABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Publishsongapi *PublishsongapiRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Publishsongapi.Contract.PublishsongapiCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Publishsongapi *PublishsongapiRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Publishsongapi.Contract.PublishsongapiTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Publishsongapi *PublishsongapiRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Publishsongapi.Contract.PublishsongapiTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Publishsongapi *PublishsongapiCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Publishsongapi.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Publishsongapi *PublishsongapiTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Publishsongapi.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Publishsongapi *PublishsongapiTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Publishsongapi.Contract.contract.Transact(opts, method, params...)
}

// Cids is a free data retrieval call binding the contract method 0x704fe710.
//
// Solidity: function cids(uint256 ) view returns(string)
func (_Publishsongapi *PublishsongapiCaller) Cids(opts *bind.CallOpts, arg0 *big.Int) (string, error) {
	var out []interface{}
	err := _Publishsongapi.contract.Call(opts, &out, "cids", arg0)

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Cids is a free data retrieval call binding the contract method 0x704fe710.
//
// Solidity: function cids(uint256 ) view returns(string)
func (_Publishsongapi *PublishsongapiSession) Cids(arg0 *big.Int) (string, error) {
	return _Publishsongapi.Contract.Cids(&_Publishsongapi.CallOpts, arg0)
}

// Cids is a free data retrieval call binding the contract method 0x704fe710.
//
// Solidity: function cids(uint256 ) view returns(string)
func (_Publishsongapi *PublishsongapiCallerSession) Cids(arg0 *big.Int) (string, error) {
	return _Publishsongapi.Contract.Cids(&_Publishsongapi.CallOpts, arg0)
}

// GetAllCid is a free data retrieval call binding the contract method 0x7c522a8b.
//
// Solidity: function getAllCid() view returns(string[])
func (_Publishsongapi *PublishsongapiCaller) GetAllCid(opts *bind.CallOpts) ([]string, error) {
	var out []interface{}
	err := _Publishsongapi.contract.Call(opts, &out, "getAllCid")

	if err != nil {
		return *new([]string), err
	}

	out0 := *abi.ConvertType(out[0], new([]string)).(*[]string)

	return out0, err

}

// GetAllCid is a free data retrieval call binding the contract method 0x7c522a8b.
//
// Solidity: function getAllCid() view returns(string[])
func (_Publishsongapi *PublishsongapiSession) GetAllCid() ([]string, error) {
	return _Publishsongapi.Contract.GetAllCid(&_Publishsongapi.CallOpts)
}

// GetAllCid is a free data retrieval call binding the contract method 0x7c522a8b.
//
// Solidity: function getAllCid() view returns(string[])
func (_Publishsongapi *PublishsongapiCallerSession) GetAllCid() ([]string, error) {
	return _Publishsongapi.Contract.GetAllCid(&_Publishsongapi.CallOpts)
}

// GetSong is a free data retrieval call binding the contract method 0x50ef10c2.
//
// Solidity: function getSong(string _cid) view returns((string,string,string[],string,uint256,address,bool))
func (_Publishsongapi *PublishsongapiCaller) GetSong(opts *bind.CallOpts, _cid string) (Struct0, error) {
	var out []interface{}
	err := _Publishsongapi.contract.Call(opts, &out, "getSong", _cid)

	if err != nil {
		return *new(Struct0), err
	}

	out0 := *abi.ConvertType(out[0], new(Struct0)).(*Struct0)

	return out0, err

}

// GetSong is a free data retrieval call binding the contract method 0x50ef10c2.
//
// Solidity: function getSong(string _cid) view returns((string,string,string[],string,uint256,address,bool))
func (_Publishsongapi *PublishsongapiSession) GetSong(_cid string) (Struct0, error) {
	return _Publishsongapi.Contract.GetSong(&_Publishsongapi.CallOpts, _cid)
}

// GetSong is a free data retrieval call binding the contract method 0x50ef10c2.
//
// Solidity: function getSong(string _cid) view returns((string,string,string[],string,uint256,address,bool))
func (_Publishsongapi *PublishsongapiCallerSession) GetSong(_cid string) (Struct0, error) {
	return _Publishsongapi.Contract.GetSong(&_Publishsongapi.CallOpts, _cid)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Publishsongapi *PublishsongapiCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Publishsongapi.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Publishsongapi *PublishsongapiSession) Owner() (common.Address, error) {
	return _Publishsongapi.Contract.Owner(&_Publishsongapi.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Publishsongapi *PublishsongapiCallerSession) Owner() (common.Address, error) {
	return _Publishsongapi.Contract.Owner(&_Publishsongapi.CallOpts)
}

// TotalSongs is a free data retrieval call binding the contract method 0x7ecaa061.
//
// Solidity: function totalSongs() view returns(uint16)
func (_Publishsongapi *PublishsongapiCaller) TotalSongs(opts *bind.CallOpts) (uint16, error) {
	var out []interface{}
	err := _Publishsongapi.contract.Call(opts, &out, "totalSongs")

	if err != nil {
		return *new(uint16), err
	}

	out0 := *abi.ConvertType(out[0], new(uint16)).(*uint16)

	return out0, err

}

// TotalSongs is a free data retrieval call binding the contract method 0x7ecaa061.
//
// Solidity: function totalSongs() view returns(uint16)
func (_Publishsongapi *PublishsongapiSession) TotalSongs() (uint16, error) {
	return _Publishsongapi.Contract.TotalSongs(&_Publishsongapi.CallOpts)
}

// TotalSongs is a free data retrieval call binding the contract method 0x7ecaa061.
//
// Solidity: function totalSongs() view returns(uint16)
func (_Publishsongapi *PublishsongapiCallerSession) TotalSongs() (uint16, error) {
	return _Publishsongapi.Contract.TotalSongs(&_Publishsongapi.CallOpts)
}

// AddSong is a paid mutator transaction binding the contract method 0xea06c8c1.
//
// Solidity: function addSong(string _cid, string _name, string[] _artists, string _description) returns()
func (_Publishsongapi *PublishsongapiTransactor) AddSong(opts *bind.TransactOpts, _cid string, _name string, _artists []string, _description string) (*types.Transaction, error) {
	return _Publishsongapi.contract.Transact(opts, "addSong", _cid, _name, _artists, _description)
}

// AddSong is a paid mutator transaction binding the contract method 0xea06c8c1.
//
// Solidity: function addSong(string _cid, string _name, string[] _artists, string _description) returns()
func (_Publishsongapi *PublishsongapiSession) AddSong(_cid string, _name string, _artists []string, _description string) (*types.Transaction, error) {
	return _Publishsongapi.Contract.AddSong(&_Publishsongapi.TransactOpts, _cid, _name, _artists, _description)
}

// AddSong is a paid mutator transaction binding the contract method 0xea06c8c1.
//
// Solidity: function addSong(string _cid, string _name, string[] _artists, string _description) returns()
func (_Publishsongapi *PublishsongapiTransactorSession) AddSong(_cid string, _name string, _artists []string, _description string) (*types.Transaction, error) {
	return _Publishsongapi.Contract.AddSong(&_Publishsongapi.TransactOpts, _cid, _name, _artists, _description)
}
