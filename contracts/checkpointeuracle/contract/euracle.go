// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package contract

import (
	"math/big"
	"strings"

	eutherum "github.com/eutherum/eutherum"
	"github.com/eutherum/eutherum/accounts/abi"
	"github.com/eutherum/eutherum/accounts/abi/bind"
	"github.com/eutherum/eutherum/common"
	"github.com/eutherum/eutherum/core/types"
	"github.com/eutherum/eutherum/event"
)

// Reference imports to suppress errors if they are not otherwise used.
var (
	_ = big.NewInt
	_ = strings.NewReader
	_ = eutherum.NotFound
	_ = bind.Bind
	_ = common.Big1
	_ = types.BloomLookup
	_ = event.NewSubscription
)

// CheckpointEuracleABI is the input ABI used to generate the binding from.
const CheckpointEuracleABI = "[{\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"_adminlist\",\"type\":\"address[]\"},{\"internalType\":\"uint256\",\"name\":\"_sectionSize\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_processConfirms\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_threshold\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint64\",\"name\":\"index\",\"type\":\"uint64\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"checkpointHash\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"uint8\",\"name\":\"v\",\"type\":\"uint8\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"r\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"s\",\"type\":\"bytes32\"}],\"name\":\"NewCheckpointVote\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"GetAllAdmin\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"\",\"type\":\"address[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"GetLatestCheckpoint\",\"outputs\":[{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"},{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_recentNumber\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"_recentHash\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"_hash\",\"type\":\"bytes32\"},{\"internalType\":\"uint64\",\"name\":\"_sectionIndex\",\"type\":\"uint64\"},{\"internalType\":\"uint8[]\",\"name\":\"v\",\"type\":\"uint8[]\"},{\"internalType\":\"bytes32[]\",\"name\":\"r\",\"type\":\"bytes32[]\"},{\"internalType\":\"bytes32[]\",\"name\":\"s\",\"type\":\"bytes32[]\"}],\"name\":\"SetCheckpoint\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"

// CheckpointEuracleFuncSigs maps the 4-byte function signature to its string representation.
var CheckpointEuracleFuncSigs = map[string]string{
	"45848dfc": "GetAllAdmin()",
	"4d6a304c": "GetLatestCheckpoint()",
	"d459fc46": "SetCheckpoint(uint256,bytes32,bytes32,uint64,uint8[],bytes32[],bytes32[])",
}

// CheckpointEuracleBin is the compiled bytecode used for deploying new contracts.
var CheckpointEuracleBin = "0x608060405234801561001057600080fd5b506040516108703803806108708339818101604052608081101561003357600080fd5b810190808051604051939291908464010000000082111561005357600080fd5b90830190602082018581111561006857600080fd5b825186602082028301116401000000008211171561008557600080fd5b82525081516020918201928201910280838360005b838110156100b257818101518382015260200161009a565b50505050919091016040908152602083015190830151606090930151909450919250600090505b84518110156101855760016000808784815181106100f357fe5b60200260200101516001600160a01b03166001600160a01b0316815260200190815260200160002060006101000a81548160ff021916908315150217905550600185828151811061014057fe5b60209081029190910181015182546001808201855560009485529290932090920180546001600160a01b0319166001600160a01b0390931692909217909155016100d9565b50600592909255600655600755506106ce806101a26000396000f3fe608060405234801561001057600080fd5b50600436106100415760003560e01c806345848dfc146100465780634d6a304c1461009e578063d459fc46146100cf575b600080fd5b61004e6102b0565b60408051602080825283518183015283519192839290830191858101910280838360005b8381101561008a578181015183820152602001610072565b505050509050019250505060405180910390f35b6100a6610365565b6040805167ffffffffffffffff9094168452602084019290925282820152519081900360600190f35b61029c600480360360e08110156100e557600080fd5b81359160208101359160408201359167ffffffffffffffff6060820135169181019060a08101608082013564010000000081111561012257600080fd5b82018360208201111561013457600080fd5b8035906020019184602083028401116401000000008311171561015657600080fd5b91908080602002602001604051908101604052809392919081815260200183836020028082843760009201919091525092959493602081019350359150506401000000008111156101a657600080fd5b8201836020820111156101b857600080fd5b803590602001918460208302840111640100000000831117156101da57600080fd5b919080806020026020016040519081016040528093929190818152602001838360200280828437600092019190915250929594936020810193503591505064010000000081111561022a57600080fd5b82018360208201111561023c57600080fd5b8035906020019184602083028401116401000000008311171561025e57600080fd5b919080806020026020016040519081016040528093929190818152602001838360200280828437600092019190915250929550610380945050505050565b604080519115158252519081900360200190f35b600154606090819067ffffffffffffffff811180156102ce57600080fd5b506040519080825280602002602001820160405280156102f8578160200160208202803683370190505b50905060005b60015481101561035f576001818154811061031557fe5b9060005260206000200160009054906101000a90046001600160a01b031682828151811061033f57fe5b6001600160a01b03909216602092830291909101909101526001016102fe565b50905090565b60025460045460035467ffffffffffffffff90921691909192565b3360009081526020819052604081205460ff1661039c57600080fd5b868840146103a957600080fd5b82518451146103b757600080fd5b81518451146103c557600080fd5b6006546005548660010167ffffffffffffffff1602014310156103ea5750600061068d565b60025467ffffffffffffffff908116908616101561040a5750600061068d565b60025467ffffffffffffffff868116911614801561043c575067ffffffffffffffff851615158061043c575060035415155b156104495750600061068d565b856104565750600061068d565b60408051601960f81b6020808301919091526000602183018190523060601b60228401526001600160c01b031960c08a901b166036840152603e8084018b905284518085039091018152605e909301909352815191012090805b86518110156106875760006001848984815181106104ca57fe5b60200260200101518985815181106104de57fe5b60200260200101518986815181106104f257fe5b602002602001015160405160008152602001604052604051808581526020018460ff1660ff1681526020018381526020018281526020019450505050506020604051602081039080840390855afa158015610551573d6000803e3d6000fd5b505060408051601f1901516001600160a01b03811660009081526020819052919091205490925060ff16905061058657600080fd5b826001600160a01b0316816001600160a01b0316116105a457600080fd5b8092508867ffffffffffffffff167fce51ffa16246bcaf0899f6504f473cd0114f430f566cef71ab7e03d3dde42a418b8a85815181106105e057fe5b60200260200101518a86815181106105f457fe5b60200260200101518a878151811061060857fe5b6020026020010151604051808581526020018460ff1660ff16815260200183815260200182815260200194505050505060405180910390a2600754826001011061067e5750505060048790555050436003556002805467ffffffffffffffff191667ffffffffffffffff8616179055600161068d565b506001016104b0565b50600080fd5b97965050505050505056fea26469706673582212202ddf9eda76bf59c0fc65584c0b22d84ecef2c703765de60439596d6ac34c2b7264736f6c634300060b0033"

// DeployCheckpointEuracle deploys a new Eutherum contract, binding an instance of CheckpointEuracle to it.
func DeployCheckpointEuracle(auth *bind.TransactOpts, backend bind.ContractBackend, _adminlist []common.Address, _sectionSize *big.Int, _processConfirms *big.Int, _threshold *big.Int) (common.Address, *types.Transaction, *CheckpointEuracle, error) {
	parsed, err := abi.JSON(strings.NewReader(CheckpointEuracleABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}

	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(CheckpointEuracleBin), backend, _adminlist, _sectionSize, _processConfirms, _threshold)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &CheckpointEuracle{CheckpointEuracleCaller: CheckpointEuracleCaller{contract: contract}, CheckpointEuracleTransactor: CheckpointEuracleTransactor{contract: contract}, CheckpointEuracleFilterer: CheckpointEuracleFilterer{contract: contract}}, nil
}

// CheckpointEuracle is an auto generated Go binding around an Eutherum contract.
type CheckpointEuracle struct {
	CheckpointEuracleCaller     // Read-only binding to the contract
	CheckpointEuracleTransactor // Write-only binding to the contract
	CheckpointEuracleFilterer   // Log filterer for contract events
}

// CheckpointEuracleCaller is an auto generated read-only Go binding around an Eutherum contract.
type CheckpointEuracleCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// CheckpointEuracleTransactor is an auto generated write-only Go binding around an Eutherum contract.
type CheckpointEuracleTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// CheckpointEuracleFilterer is an auto generated log filtering Go binding around an Eutherum contract events.
type CheckpointEuracleFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// CheckpointEuracleSession is an auto generated Go binding around an Eutherum contract,
// with pre-set call and transact options.
type CheckpointEuracleSession struct {
	Contract     *CheckpointEuracle // Generic contract binding to set the session for
	CallOpts     bind.CallOpts      // Call options to use throughout this session
	TransactOpts bind.TransactOpts  // Transaction auth options to use throughout this session
}

// CheckpointEuracleCallerSession is an auto generated read-only Go binding around an Eutherum contract,
// with pre-set call options.
type CheckpointEuracleCallerSession struct {
	Contract *CheckpointEuracleCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts            // Call options to use throughout this session
}

// CheckpointEuracleTransactorSession is an auto generated write-only Go binding around an Eutherum contract,
// with pre-set transact options.
type CheckpointEuracleTransactorSession struct {
	Contract     *CheckpointEuracleTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts            // Transaction auth options to use throughout this session
}

// CheckpointEuracleRaw is an auto generated low-level Go binding around an Eutherum contract.
type CheckpointEuracleRaw struct {
	Contract *CheckpointEuracle // Generic contract binding to access the raw methods on
}

// CheckpointEuracleCallerRaw is an auto generated low-level read-only Go binding around an Eutherum contract.
type CheckpointEuracleCallerRaw struct {
	Contract *CheckpointEuracleCaller // Generic read-only contract binding to access the raw methods on
}

// CheckpointEuracleTransactorRaw is an auto generated low-level write-only Go binding around an Eutherum contract.
type CheckpointEuracleTransactorRaw struct {
	Contract *CheckpointEuracleTransactor // Generic write-only contract binding to access the raw methods on
}

// NewCheckpointEuracle creates a new instance of CheckpointEuracle, bound to a specific deployed contract.
func NewCheckpointEuracle(address common.Address, backend bind.ContractBackend) (*CheckpointEuracle, error) {
	contract, err := bindCheckpointEuracle(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &CheckpointEuracle{CheckpointEuracleCaller: CheckpointEuracleCaller{contract: contract}, CheckpointEuracleTransactor: CheckpointEuracleTransactor{contract: contract}, CheckpointEuracleFilterer: CheckpointEuracleFilterer{contract: contract}}, nil
}

// NewCheckpointEuracleCaller creates a new read-only instance of CheckpointEuracle, bound to a specific deployed contract.
func NewCheckpointEuracleCaller(address common.Address, caller bind.ContractCaller) (*CheckpointEuracleCaller, error) {
	contract, err := bindCheckpointEuracle(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &CheckpointEuracleCaller{contract: contract}, nil
}

// NewCheckpointEuracleTransactor creates a new write-only instance of CheckpointEuracle, bound to a specific deployed contract.
func NewCheckpointEuracleTransactor(address common.Address, transactor bind.ContractTransactor) (*CheckpointEuracleTransactor, error) {
	contract, err := bindCheckpointEuracle(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &CheckpointEuracleTransactor{contract: contract}, nil
}

// NewCheckpointEuracleFilterer creates a new log filterer instance of CheckpointEuracle, bound to a specific deployed contract.
func NewCheckpointEuracleFilterer(address common.Address, filterer bind.ContractFilterer) (*CheckpointEuracleFilterer, error) {
	contract, err := bindCheckpointEuracle(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &CheckpointEuracleFilterer{contract: contract}, nil
}

// bindCheckpointEuracle binds a generic wrapper to an already deployed contract.
func bindCheckpointEuracle(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(CheckpointEuracleABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_CheckpointEuracle *CheckpointEuracleRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _CheckpointEuracle.Contract.CheckpointEuracleCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_CheckpointEuracle *CheckpointEuracleRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _CheckpointEuracle.Contract.CheckpointEuracleTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_CheckpointEuracle *CheckpointEuracleRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _CheckpointEuracle.Contract.CheckpointEuracleTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_CheckpointEuracle *CheckpointEuracleCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _CheckpointEuracle.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_CheckpointEuracle *CheckpointEuracleTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _CheckpointEuracle.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_CheckpointEuracle *CheckpointEuracleTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _CheckpointEuracle.Contract.contract.Transact(opts, method, params...)
}

// GetAllAdmin is a free data retrieval call binding the contract method 0x45848dfc.
//
// Solidity: function GetAllAdmin() view returns(address[])
func (_CheckpointEuracle *CheckpointEuracleCaller) GetAllAdmin(opts *bind.CallOpts) ([]common.Address, error) {
	var out []interface{}
	err := _CheckpointEuracle.contract.Call(opts, &out, "GetAllAdmin")

	if err != nil {
		return *new([]common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)

	return out0, err

}

// GetAllAdmin is a free data retrieval call binding the contract method 0x45848dfc.
//
// Solidity: function GetAllAdmin() view returns(address[])
func (_CheckpointEuracle *CheckpointEuracleSession) GetAllAdmin() ([]common.Address, error) {
	return _CheckpointEuracle.Contract.GetAllAdmin(&_CheckpointEuracle.CallOpts)
}

// GetAllAdmin is a free data retrieval call binding the contract method 0x45848dfc.
//
// Solidity: function GetAllAdmin() view returns(address[])
func (_CheckpointEuracle *CheckpointEuracleCallerSession) GetAllAdmin() ([]common.Address, error) {
	return _CheckpointEuracle.Contract.GetAllAdmin(&_CheckpointEuracle.CallOpts)
}

// GetLatestCheckpoint is a free data retrieval call binding the contract method 0x4d6a304c.
//
// Solidity: function GetLatestCheckpoint() view returns(uint64, bytes32, uint256)
func (_CheckpointEuracle *CheckpointEuracleCaller) GetLatestCheckpoint(opts *bind.CallOpts) (uint64, [32]byte, *big.Int, error) {
	var out []interface{}
	err := _CheckpointEuracle.contract.Call(opts, &out, "GetLatestCheckpoint")

	if err != nil {
		return *new(uint64), *new([32]byte), *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(uint64)).(*uint64)
	out1 := *abi.ConvertType(out[1], new([32]byte)).(*[32]byte)
	out2 := *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)

	return out0, out1, out2, err

}

// GetLatestCheckpoint is a free data retrieval call binding the contract method 0x4d6a304c.
//
// Solidity: function GetLatestCheckpoint() view returns(uint64, bytes32, uint256)
func (_CheckpointEuracle *CheckpointEuracleSession) GetLatestCheckpoint() (uint64, [32]byte, *big.Int, error) {
	return _CheckpointEuracle.Contract.GetLatestCheckpoint(&_CheckpointEuracle.CallOpts)
}

// GetLatestCheckpoint is a free data retrieval call binding the contract method 0x4d6a304c.
//
// Solidity: function GetLatestCheckpoint() view returns(uint64, bytes32, uint256)
func (_CheckpointEuracle *CheckpointEuracleCallerSession) GetLatestCheckpoint() (uint64, [32]byte, *big.Int, error) {
	return _CheckpointEuracle.Contract.GetLatestCheckpoint(&_CheckpointEuracle.CallOpts)
}

// SetCheckpoint is a paid mutator transaction binding the contract method 0xd459fc46.
//
// Solidity: function SetCheckpoint(uint256 _recentNumber, bytes32 _recentHash, bytes32 _hash, uint64 _sectionIndex, uint8[] v, bytes32[] r, bytes32[] s) returns(bool)
func (_CheckpointEuracle *CheckpointEuracleTransactor) SetCheckpoint(opts *bind.TransactOpts, _recentNumber *big.Int, _recentHash [32]byte, _hash [32]byte, _sectionIndex uint64, v []uint8, r [][32]byte, s [][32]byte) (*types.Transaction, error) {
	return _CheckpointEuracle.contract.Transact(opts, "SetCheckpoint", _recentNumber, _recentHash, _hash, _sectionIndex, v, r, s)
}

// SetCheckpoint is a paid mutator transaction binding the contract method 0xd459fc46.
//
// Solidity: function SetCheckpoint(uint256 _recentNumber, bytes32 _recentHash, bytes32 _hash, uint64 _sectionIndex, uint8[] v, bytes32[] r, bytes32[] s) returns(bool)
func (_CheckpointEuracle *CheckpointEuracleSession) SetCheckpoint(_recentNumber *big.Int, _recentHash [32]byte, _hash [32]byte, _sectionIndex uint64, v []uint8, r [][32]byte, s [][32]byte) (*types.Transaction, error) {
	return _CheckpointEuracle.Contract.SetCheckpoint(&_CheckpointEuracle.TransactOpts, _recentNumber, _recentHash, _hash, _sectionIndex, v, r, s)
}

// SetCheckpoint is a paid mutator transaction binding the contract method 0xd459fc46.
//
// Solidity: function SetCheckpoint(uint256 _recentNumber, bytes32 _recentHash, bytes32 _hash, uint64 _sectionIndex, uint8[] v, bytes32[] r, bytes32[] s) returns(bool)
func (_CheckpointEuracle *CheckpointEuracleTransactorSession) SetCheckpoint(_recentNumber *big.Int, _recentHash [32]byte, _hash [32]byte, _sectionIndex uint64, v []uint8, r [][32]byte, s [][32]byte) (*types.Transaction, error) {
	return _CheckpointEuracle.Contract.SetCheckpoint(&_CheckpointEuracle.TransactOpts, _recentNumber, _recentHash, _hash, _sectionIndex, v, r, s)
}

// CheckpointEuracleNewCheckpointVoteIterator is returned from FilterNewCheckpointVote and is used to iterate over the raw logs and unpacked data for NewCheckpointVote events raised by the CheckpointEuracle contract.
type CheckpointEuracleNewCheckpointVoteIterator struct {
	Event *CheckpointEuracleNewCheckpointVote // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  eutherum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *CheckpointEuracleNewCheckpointVoteIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CheckpointEuracleNewCheckpointVote)
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
		it.Event = new(CheckpointEuracleNewCheckpointVote)
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
func (it *CheckpointEuracleNewCheckpointVoteIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CheckpointEuracleNewCheckpointVoteIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CheckpointEuracleNewCheckpointVote represents a NewCheckpointVote event raised by the CheckpointEuracle contract.
type CheckpointEuracleNewCheckpointVote struct {
	Index          uint64
	CheckpointHash [32]byte
	V              uint8
	R              [32]byte
	S              [32]byte
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterNewCheckpointVote is a free log retrieval operation binding the contract event 0xce51ffa16246bcaf0899f6504f473cd0114f430f566cef71ab7e03d3dde42a41.
//
// Solidity: event NewCheckpointVote(uint64 indexed index, bytes32 checkpointHash, uint8 v, bytes32 r, bytes32 s)
func (_CheckpointEuracle *CheckpointEuracleFilterer) FilterNewCheckpointVote(opts *bind.FilterOpts, index []uint64) (*CheckpointEuracleNewCheckpointVoteIterator, error) {

	var indexRule []interface{}
	for _, indexItem := range index {
		indexRule = append(indexRule, indexItem)
	}

	logs, sub, err := _CheckpointEuracle.contract.FilterLogs(opts, "NewCheckpointVote", indexRule)
	if err != nil {
		return nil, err
	}
	return &CheckpointEuracleNewCheckpointVoteIterator{contract: _CheckpointEuracle.contract, event: "NewCheckpointVote", logs: logs, sub: sub}, nil
}

// WatchNewCheckpointVote is a free log subscription operation binding the contract event 0xce51ffa16246bcaf0899f6504f473cd0114f430f566cef71ab7e03d3dde42a41.
//
// Solidity: event NewCheckpointVote(uint64 indexed index, bytes32 checkpointHash, uint8 v, bytes32 r, bytes32 s)
func (_CheckpointEuracle *CheckpointEuracleFilterer) WatchNewCheckpointVote(opts *bind.WatchOpts, sink chan<- *CheckpointEuracleNewCheckpointVote, index []uint64) (event.Subscription, error) {

	var indexRule []interface{}
	for _, indexItem := range index {
		indexRule = append(indexRule, indexItem)
	}

	logs, sub, err := _CheckpointEuracle.contract.WatchLogs(opts, "NewCheckpointVote", indexRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CheckpointEuracleNewCheckpointVote)
				if err := _CheckpointEuracle.contract.UnpackLog(event, "NewCheckpointVote", log); err != nil {
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

// ParseNewCheckpointVote is a log parse operation binding the contract event 0xce51ffa16246bcaf0899f6504f473cd0114f430f566cef71ab7e03d3dde42a41.
//
// Solidity: event NewCheckpointVote(uint64 indexed index, bytes32 checkpointHash, uint8 v, bytes32 r, bytes32 s)
func (_CheckpointEuracle *CheckpointEuracleFilterer) ParseNewCheckpointVote(log types.Log) (*CheckpointEuracleNewCheckpointVote, error) {
	event := new(CheckpointEuracleNewCheckpointVote)
	if err := _CheckpointEuracle.contract.UnpackLog(event, "NewCheckpointVote", log); err != nil {
		return nil, err
	}
	return event, nil
}
