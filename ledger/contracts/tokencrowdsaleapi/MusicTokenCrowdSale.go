// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package tokencrowdsaleapi

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

// TokencrowdsaleapiMetaData contains all meta data concerning the Tokencrowdsaleapi contract.
var TokencrowdsaleapiMetaData = &bind.MetaData{
	ABI: "[{\"constant\":true,\"inputs\":[],\"name\":\"rate\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"weiRaised\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"wallet\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"beneficiary\",\"type\":\"address\"}],\"name\":\"buyTokens\",\"outputs\":[],\"payable\":true,\"stateMutability\":\"payable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"token\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"name\":\"_rate\",\"type\":\"uint256\"},{\"name\":\"_wallet\",\"type\":\"address\"},{\"name\":\"_token\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"payable\":true,\"stateMutability\":\"payable\",\"type\":\"fallback\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"purchaser\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"beneficiary\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"value\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"TokensPurchased\",\"type\":\"event\"}]",
	Bin: "0x60806040523480156200001157600080fd5b5060405160608062000a4e83398101806040526200003391908101906200018d565b6000805460ff1916600117905582828282151562000088576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016200007f90620002f6565b60405180910390fd5b6001600160a01b0382161515620000cd576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016200007f9062000308565b6001600160a01b038116151562000112576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016200007f90620002de565b600292909255600180546001600160a01b039283166001600160a01b0319909116179055600080549190921661010002610100600160a81b0319909116179055506200034c915050565b60006200016a825162000323565b9392505050565b60006200016a825162000330565b60006200016a825162000349565b600080600060608486031215620001a357600080fd5b6000620001b186866200017f565b9350506020620001c4868287016200015c565b9250506040620001d78682870162000171565b9150509250925092565b6000620001f06024836200031a565b7f43726f776473616c653a20746f6b656e20697320746865207a65726f2061646481527f7265737300000000000000000000000000000000000000000000000000000000602082015260400192915050565b6000620002516014836200031a565b7f43726f776473616c653a20726174652069732030000000000000000000000000815260200192915050565b60006200028c6025836200031a565b7f43726f776473616c653a2077616c6c657420697320746865207a65726f20616481527f6472657373000000000000000000000000000000000000000000000000000000602082015260400192915050565b60208082528101620002f081620001e1565b92915050565b60208082528101620002f08162000242565b60208082528101620002f0816200027d565b90815260200190565b6000620002f0826200033d565b6000620002f08262000323565b6001600160a01b031690565b90565b6106f2806200035c6000396000f3fe60806040526004361061004a5760003560e01c80632c4e722e1461005c5780634042b66f14610087578063521eb2731461009c578063ec8ac4d8146100be578063fc0c546a146100cc575b61005a6100556100ee565b6100f2565b005b34801561006857600080fd5b506100716101e7565b60405161007e919061066f565b60405180910390f35b34801561009357600080fd5b506100716101ed565b3480156100a857600080fd5b506100b16101f3565b60405161007e91906105d8565b61005a6100553660046103fa565b3480156100d857600080fd5b506100e1610202565b60405161007e9190610601565b3390565b60005460ff16151561012257604051600160e51b62461bcd0281526004016101199061064f565b60405180910390fd5b6000805460ff19169055346101378282610216565b600061014282610267565b600354909150610158908363ffffffff61028416565b60035561016583826102b3565b826001600160a01b03166101776100ee565b6001600160a01b03167f6faf93231a456e552dbc9961f58d9713ee4f2e69d15f1975b050ef0911053a7b84846040516101b192919061067d565b60405180910390a36101c38383610263565b6101cb6102bd565b6101d58383610263565b50506000805460ff1916600117905550565b60025490565b60035490565b6001546001600160a01b031690565b60005461010090046001600160a01b031690565b6001600160a01b038216151561024157604051600160e51b62461bcd0281526004016101199061062f565b80151561026357604051600160e51b62461bcd0281526004016101199061065f565b5050565b600061027e600254836102f990919063ffffffff16565b92915050565b6000828201838110156102ac57604051600160e51b62461bcd0281526004016101199061060f565b9392505050565b610263828261033a565b6001546040516001600160a01b03909116903480156108fc02916000818181858888f193505050501580156102f6573d6000803e3d6000fd5b50565b600082151561030a5750600061027e565b82820282848281151561031957fe5b04146102ac57604051600160e51b62461bcd0281526004016101199061061f565b610342610202565b6001600160a01b03166340c10f1983836040518363ffffffff1660e01b815260040161036f9291906105e6565b602060405180830381600087803b15801561038957600080fd5b505af115801561039d573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052506103c19190810190610420565b151561026357604051600160e51b62461bcd0281526004016101199061063f565b60006102ac8235610694565b60006102ac82516106a8565b60006020828403121561040c57600080fd5b600061041884846103e2565b949350505050565b60006020828403121561043257600080fd5b600061041884846103ee565b61044781610694565b82525050565b610447816106ad565b6000610463601b8361068b565b7f536166654d6174683a206164646974696f6e206f766572666c6f770000000000815260200192915050565b600061049c60218361068b565b7f536166654d6174683a206d756c7469706c69636174696f6e206f766572666c6f8152600160f81b607702602082015260400192915050565b60006104e2602a8361068b565b7f43726f776473616c653a2062656e656669636961727920697320746865207a658152600160b01b69726f206164647265737302602082015260400192915050565b6000610531601f8361068b565b7f4d696e74656443726f776473616c653a206d696e74696e67206661696c656400815260200192915050565b600061056a601f8361068b565b7f5265656e7472616e637947756172643a207265656e7472616e742063616c6c00815260200192915050565b60006105a360198361068b565b7f43726f776473616c653a20776569416d6f756e74206973203000000000000000815260200192915050565b610447816106a5565b6020810161027e828461043e565b604081016105f4828561043e565b6102ac60208301846105cf565b6020810161027e828461044d565b6020808252810161027e81610456565b6020808252810161027e8161048f565b6020808252810161027e816104d5565b6020808252810161027e81610524565b6020808252810161027e8161055d565b6020808252810161027e81610596565b6020810161027e82846105cf565b604081016105f482856105cf565b90815260200190565b60006001600160a01b03821661027e565b90565b151590565b600061027e8261069456fea265627a7a72305820990d75e67c36fbdb43f12f20d5a022ccb9ceac3c31393c9332e0dea0cab33a976c6578706572696d656e74616cf50037",
}

// TokencrowdsaleapiABI is the input ABI used to generate the binding from.
// Deprecated: Use TokencrowdsaleapiMetaData.ABI instead.
var TokencrowdsaleapiABI = TokencrowdsaleapiMetaData.ABI

// TokencrowdsaleapiBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use TokencrowdsaleapiMetaData.Bin instead.
var TokencrowdsaleapiBin = TokencrowdsaleapiMetaData.Bin

// DeployTokencrowdsaleapi deploys a new Ethereum contract, binding an instance of Tokencrowdsaleapi to it.
func DeployTokencrowdsaleapi(auth *bind.TransactOpts, backend bind.ContractBackend, _rate *big.Int, _wallet common.Address, _token common.Address) (common.Address, *types.Transaction, *Tokencrowdsaleapi, error) {
	parsed, err := TokencrowdsaleapiMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(TokencrowdsaleapiBin), backend, _rate, _wallet, _token)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Tokencrowdsaleapi{TokencrowdsaleapiCaller: TokencrowdsaleapiCaller{contract: contract}, TokencrowdsaleapiTransactor: TokencrowdsaleapiTransactor{contract: contract}, TokencrowdsaleapiFilterer: TokencrowdsaleapiFilterer{contract: contract}}, nil
}

// Tokencrowdsaleapi is an auto generated Go binding around an Ethereum contract.
type Tokencrowdsaleapi struct {
	TokencrowdsaleapiCaller     // Read-only binding to the contract
	TokencrowdsaleapiTransactor // Write-only binding to the contract
	TokencrowdsaleapiFilterer   // Log filterer for contract events
}

// TokencrowdsaleapiCaller is an auto generated read-only Go binding around an Ethereum contract.
type TokencrowdsaleapiCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TokencrowdsaleapiTransactor is an auto generated write-only Go binding around an Ethereum contract.
type TokencrowdsaleapiTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TokencrowdsaleapiFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type TokencrowdsaleapiFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TokencrowdsaleapiSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type TokencrowdsaleapiSession struct {
	Contract     *Tokencrowdsaleapi // Generic contract binding to set the session for
	CallOpts     bind.CallOpts      // Call options to use throughout this session
	TransactOpts bind.TransactOpts  // Transaction auth options to use throughout this session
}

// TokencrowdsaleapiCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type TokencrowdsaleapiCallerSession struct {
	Contract *TokencrowdsaleapiCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts            // Call options to use throughout this session
}

// TokencrowdsaleapiTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type TokencrowdsaleapiTransactorSession struct {
	Contract     *TokencrowdsaleapiTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts            // Transaction auth options to use throughout this session
}

// TokencrowdsaleapiRaw is an auto generated low-level Go binding around an Ethereum contract.
type TokencrowdsaleapiRaw struct {
	Contract *Tokencrowdsaleapi // Generic contract binding to access the raw methods on
}

// TokencrowdsaleapiCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type TokencrowdsaleapiCallerRaw struct {
	Contract *TokencrowdsaleapiCaller // Generic read-only contract binding to access the raw methods on
}

// TokencrowdsaleapiTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type TokencrowdsaleapiTransactorRaw struct {
	Contract *TokencrowdsaleapiTransactor // Generic write-only contract binding to access the raw methods on
}

// NewTokencrowdsaleapi creates a new instance of Tokencrowdsaleapi, bound to a specific deployed contract.
func NewTokencrowdsaleapi(address common.Address, backend bind.ContractBackend) (*Tokencrowdsaleapi, error) {
	contract, err := bindTokencrowdsaleapi(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Tokencrowdsaleapi{TokencrowdsaleapiCaller: TokencrowdsaleapiCaller{contract: contract}, TokencrowdsaleapiTransactor: TokencrowdsaleapiTransactor{contract: contract}, TokencrowdsaleapiFilterer: TokencrowdsaleapiFilterer{contract: contract}}, nil
}

// NewTokencrowdsaleapiCaller creates a new read-only instance of Tokencrowdsaleapi, bound to a specific deployed contract.
func NewTokencrowdsaleapiCaller(address common.Address, caller bind.ContractCaller) (*TokencrowdsaleapiCaller, error) {
	contract, err := bindTokencrowdsaleapi(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &TokencrowdsaleapiCaller{contract: contract}, nil
}

// NewTokencrowdsaleapiTransactor creates a new write-only instance of Tokencrowdsaleapi, bound to a specific deployed contract.
func NewTokencrowdsaleapiTransactor(address common.Address, transactor bind.ContractTransactor) (*TokencrowdsaleapiTransactor, error) {
	contract, err := bindTokencrowdsaleapi(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &TokencrowdsaleapiTransactor{contract: contract}, nil
}

// NewTokencrowdsaleapiFilterer creates a new log filterer instance of Tokencrowdsaleapi, bound to a specific deployed contract.
func NewTokencrowdsaleapiFilterer(address common.Address, filterer bind.ContractFilterer) (*TokencrowdsaleapiFilterer, error) {
	contract, err := bindTokencrowdsaleapi(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &TokencrowdsaleapiFilterer{contract: contract}, nil
}

// bindTokencrowdsaleapi binds a generic wrapper to an already deployed contract.
func bindTokencrowdsaleapi(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(TokencrowdsaleapiABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Tokencrowdsaleapi *TokencrowdsaleapiRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Tokencrowdsaleapi.Contract.TokencrowdsaleapiCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Tokencrowdsaleapi *TokencrowdsaleapiRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Tokencrowdsaleapi.Contract.TokencrowdsaleapiTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Tokencrowdsaleapi *TokencrowdsaleapiRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Tokencrowdsaleapi.Contract.TokencrowdsaleapiTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Tokencrowdsaleapi *TokencrowdsaleapiCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Tokencrowdsaleapi.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Tokencrowdsaleapi *TokencrowdsaleapiTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Tokencrowdsaleapi.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Tokencrowdsaleapi *TokencrowdsaleapiTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Tokencrowdsaleapi.Contract.contract.Transact(opts, method, params...)
}

// Rate is a free data retrieval call binding the contract method 0x2c4e722e.
//
// Solidity: function rate() view returns(uint256)
func (_Tokencrowdsaleapi *TokencrowdsaleapiCaller) Rate(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Tokencrowdsaleapi.contract.Call(opts, &out, "rate")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Rate is a free data retrieval call binding the contract method 0x2c4e722e.
//
// Solidity: function rate() view returns(uint256)
func (_Tokencrowdsaleapi *TokencrowdsaleapiSession) Rate() (*big.Int, error) {
	return _Tokencrowdsaleapi.Contract.Rate(&_Tokencrowdsaleapi.CallOpts)
}

// Rate is a free data retrieval call binding the contract method 0x2c4e722e.
//
// Solidity: function rate() view returns(uint256)
func (_Tokencrowdsaleapi *TokencrowdsaleapiCallerSession) Rate() (*big.Int, error) {
	return _Tokencrowdsaleapi.Contract.Rate(&_Tokencrowdsaleapi.CallOpts)
}

// Token is a free data retrieval call binding the contract method 0xfc0c546a.
//
// Solidity: function token() view returns(address)
func (_Tokencrowdsaleapi *TokencrowdsaleapiCaller) Token(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Tokencrowdsaleapi.contract.Call(opts, &out, "token")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Token is a free data retrieval call binding the contract method 0xfc0c546a.
//
// Solidity: function token() view returns(address)
func (_Tokencrowdsaleapi *TokencrowdsaleapiSession) Token() (common.Address, error) {
	return _Tokencrowdsaleapi.Contract.Token(&_Tokencrowdsaleapi.CallOpts)
}

// Token is a free data retrieval call binding the contract method 0xfc0c546a.
//
// Solidity: function token() view returns(address)
func (_Tokencrowdsaleapi *TokencrowdsaleapiCallerSession) Token() (common.Address, error) {
	return _Tokencrowdsaleapi.Contract.Token(&_Tokencrowdsaleapi.CallOpts)
}

// Wallet is a free data retrieval call binding the contract method 0x521eb273.
//
// Solidity: function wallet() view returns(address)
func (_Tokencrowdsaleapi *TokencrowdsaleapiCaller) Wallet(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Tokencrowdsaleapi.contract.Call(opts, &out, "wallet")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Wallet is a free data retrieval call binding the contract method 0x521eb273.
//
// Solidity: function wallet() view returns(address)
func (_Tokencrowdsaleapi *TokencrowdsaleapiSession) Wallet() (common.Address, error) {
	return _Tokencrowdsaleapi.Contract.Wallet(&_Tokencrowdsaleapi.CallOpts)
}

// Wallet is a free data retrieval call binding the contract method 0x521eb273.
//
// Solidity: function wallet() view returns(address)
func (_Tokencrowdsaleapi *TokencrowdsaleapiCallerSession) Wallet() (common.Address, error) {
	return _Tokencrowdsaleapi.Contract.Wallet(&_Tokencrowdsaleapi.CallOpts)
}

// WeiRaised is a free data retrieval call binding the contract method 0x4042b66f.
//
// Solidity: function weiRaised() view returns(uint256)
func (_Tokencrowdsaleapi *TokencrowdsaleapiCaller) WeiRaised(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Tokencrowdsaleapi.contract.Call(opts, &out, "weiRaised")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// WeiRaised is a free data retrieval call binding the contract method 0x4042b66f.
//
// Solidity: function weiRaised() view returns(uint256)
func (_Tokencrowdsaleapi *TokencrowdsaleapiSession) WeiRaised() (*big.Int, error) {
	return _Tokencrowdsaleapi.Contract.WeiRaised(&_Tokencrowdsaleapi.CallOpts)
}

// WeiRaised is a free data retrieval call binding the contract method 0x4042b66f.
//
// Solidity: function weiRaised() view returns(uint256)
func (_Tokencrowdsaleapi *TokencrowdsaleapiCallerSession) WeiRaised() (*big.Int, error) {
	return _Tokencrowdsaleapi.Contract.WeiRaised(&_Tokencrowdsaleapi.CallOpts)
}

// BuyTokens is a paid mutator transaction binding the contract method 0xec8ac4d8.
//
// Solidity: function buyTokens(address beneficiary) payable returns()
func (_Tokencrowdsaleapi *TokencrowdsaleapiTransactor) BuyTokens(opts *bind.TransactOpts, beneficiary common.Address) (*types.Transaction, error) {
	return _Tokencrowdsaleapi.contract.Transact(opts, "buyTokens", beneficiary)
}

// BuyTokens is a paid mutator transaction binding the contract method 0xec8ac4d8.
//
// Solidity: function buyTokens(address beneficiary) payable returns()
func (_Tokencrowdsaleapi *TokencrowdsaleapiSession) BuyTokens(beneficiary common.Address) (*types.Transaction, error) {
	return _Tokencrowdsaleapi.Contract.BuyTokens(&_Tokencrowdsaleapi.TransactOpts, beneficiary)
}

// BuyTokens is a paid mutator transaction binding the contract method 0xec8ac4d8.
//
// Solidity: function buyTokens(address beneficiary) payable returns()
func (_Tokencrowdsaleapi *TokencrowdsaleapiTransactorSession) BuyTokens(beneficiary common.Address) (*types.Transaction, error) {
	return _Tokencrowdsaleapi.Contract.BuyTokens(&_Tokencrowdsaleapi.TransactOpts, beneficiary)
}

// Fallback is a paid mutator transaction binding the contract fallback function.
//
// Solidity: fallback() payable returns()
func (_Tokencrowdsaleapi *TokencrowdsaleapiTransactor) Fallback(opts *bind.TransactOpts, calldata []byte) (*types.Transaction, error) {
	return _Tokencrowdsaleapi.contract.RawTransact(opts, calldata)
}

// Fallback is a paid mutator transaction binding the contract fallback function.
//
// Solidity: fallback() payable returns()
func (_Tokencrowdsaleapi *TokencrowdsaleapiSession) Fallback(calldata []byte) (*types.Transaction, error) {
	return _Tokencrowdsaleapi.Contract.Fallback(&_Tokencrowdsaleapi.TransactOpts, calldata)
}

// Fallback is a paid mutator transaction binding the contract fallback function.
//
// Solidity: fallback() payable returns()
func (_Tokencrowdsaleapi *TokencrowdsaleapiTransactorSession) Fallback(calldata []byte) (*types.Transaction, error) {
	return _Tokencrowdsaleapi.Contract.Fallback(&_Tokencrowdsaleapi.TransactOpts, calldata)
}

// TokencrowdsaleapiTokensPurchasedIterator is returned from FilterTokensPurchased and is used to iterate over the raw logs and unpacked data for TokensPurchased events raised by the Tokencrowdsaleapi contract.
type TokencrowdsaleapiTokensPurchasedIterator struct {
	Event *TokencrowdsaleapiTokensPurchased // Event containing the contract specifics and raw log

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
func (it *TokencrowdsaleapiTokensPurchasedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TokencrowdsaleapiTokensPurchased)
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
		it.Event = new(TokencrowdsaleapiTokensPurchased)
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
func (it *TokencrowdsaleapiTokensPurchasedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TokencrowdsaleapiTokensPurchasedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TokencrowdsaleapiTokensPurchased represents a TokensPurchased event raised by the Tokencrowdsaleapi contract.
type TokencrowdsaleapiTokensPurchased struct {
	Purchaser   common.Address
	Beneficiary common.Address
	Value       *big.Int
	Amount      *big.Int
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterTokensPurchased is a free log retrieval operation binding the contract event 0x6faf93231a456e552dbc9961f58d9713ee4f2e69d15f1975b050ef0911053a7b.
//
// Solidity: event TokensPurchased(address indexed purchaser, address indexed beneficiary, uint256 value, uint256 amount)
func (_Tokencrowdsaleapi *TokencrowdsaleapiFilterer) FilterTokensPurchased(opts *bind.FilterOpts, purchaser []common.Address, beneficiary []common.Address) (*TokencrowdsaleapiTokensPurchasedIterator, error) {

	var purchaserRule []interface{}
	for _, purchaserItem := range purchaser {
		purchaserRule = append(purchaserRule, purchaserItem)
	}
	var beneficiaryRule []interface{}
	for _, beneficiaryItem := range beneficiary {
		beneficiaryRule = append(beneficiaryRule, beneficiaryItem)
	}

	logs, sub, err := _Tokencrowdsaleapi.contract.FilterLogs(opts, "TokensPurchased", purchaserRule, beneficiaryRule)
	if err != nil {
		return nil, err
	}
	return &TokencrowdsaleapiTokensPurchasedIterator{contract: _Tokencrowdsaleapi.contract, event: "TokensPurchased", logs: logs, sub: sub}, nil
}

// WatchTokensPurchased is a free log subscription operation binding the contract event 0x6faf93231a456e552dbc9961f58d9713ee4f2e69d15f1975b050ef0911053a7b.
//
// Solidity: event TokensPurchased(address indexed purchaser, address indexed beneficiary, uint256 value, uint256 amount)
func (_Tokencrowdsaleapi *TokencrowdsaleapiFilterer) WatchTokensPurchased(opts *bind.WatchOpts, sink chan<- *TokencrowdsaleapiTokensPurchased, purchaser []common.Address, beneficiary []common.Address) (event.Subscription, error) {

	var purchaserRule []interface{}
	for _, purchaserItem := range purchaser {
		purchaserRule = append(purchaserRule, purchaserItem)
	}
	var beneficiaryRule []interface{}
	for _, beneficiaryItem := range beneficiary {
		beneficiaryRule = append(beneficiaryRule, beneficiaryItem)
	}

	logs, sub, err := _Tokencrowdsaleapi.contract.WatchLogs(opts, "TokensPurchased", purchaserRule, beneficiaryRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TokencrowdsaleapiTokensPurchased)
				if err := _Tokencrowdsaleapi.contract.UnpackLog(event, "TokensPurchased", log); err != nil {
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

// ParseTokensPurchased is a log parse operation binding the contract event 0x6faf93231a456e552dbc9961f58d9713ee4f2e69d15f1975b050ef0911053a7b.
//
// Solidity: event TokensPurchased(address indexed purchaser, address indexed beneficiary, uint256 value, uint256 amount)
func (_Tokencrowdsaleapi *TokencrowdsaleapiFilterer) ParseTokensPurchased(log types.Log) (*TokencrowdsaleapiTokensPurchased, error) {
	event := new(TokencrowdsaleapiTokensPurchased)
	if err := _Tokencrowdsaleapi.contract.UnpackLog(event, "TokensPurchased", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
