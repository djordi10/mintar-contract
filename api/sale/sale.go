// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package sale

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

// Reference imports to suppress errors if they are not otherwise used.
var (
	_ = big.NewInt
	_ = strings.NewReader
	_ = ethereum.NotFound
	_ = bind.Bind
	_ = common.Big1
	_ = types.BloomLookup
	_ = event.NewSubscription
)

// SaleABI is the input ABI used to generate the binding from.
const SaleABI = "[{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"_contractAddress\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_tokenId\",\"type\":\"uint256\"}],\"name\":\"SaleCanceled\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"_contractAddress\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"_tokenId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint128\",\"name\":\"_price\",\"type\":\"uint128\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"_seller\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"SaleCreated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"_contractAddress\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"_tokenId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint128\",\"name\":\"_price\",\"type\":\"uint128\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"_buyer\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"_seller\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"SaleSuccessful\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_contractAddress\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_tokenId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"cancelSale\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_contractAddress\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"id\",\"type\":\"bytes\"},{\"internalType\":\"uint64\",\"name\":\"_amount\",\"type\":\"uint64\"},{\"internalType\":\"uint16\",\"name\":\"__royaltyFee\",\"type\":\"uint16\"},{\"internalType\":\"uint128\",\"name\":\"_price\",\"type\":\"uint128\"}],\"name\":\"createAndSell\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_contractAddress\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_tokenId\",\"type\":\"uint256\"},{\"internalType\":\"uint128\",\"name\":\"_price\",\"type\":\"uint128\"},{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"createSale\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_contractAddress\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_seller\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_tokenId\",\"type\":\"uint256\"},{\"internalType\":\"uint128\",\"name\":\"_price\",\"type\":\"uint128\"},{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"createSaleFromContract\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_contractAddress\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_tokenId\",\"type\":\"uint256\"}],\"name\":\"isOnSale\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"uint256[]\",\"name\":\"\",\"type\":\"uint256[]\"},{\"internalType\":\"uint256[]\",\"name\":\"\",\"type\":\"uint256[]\"},{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"name\":\"onERC1155BatchReceived\",\"outputs\":[{\"internalType\":\"bytes4\",\"name\":\"\",\"type\":\"bytes4\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"name\":\"onERC1155Received\",\"outputs\":[{\"internalType\":\"bytes4\",\"name\":\"\",\"type\":\"bytes4\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_contractAddress\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_tokenId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"purchase\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"sales\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"seller\",\"type\":\"address\"},{\"internalType\":\"uint128\",\"name\":\"price\",\"type\":\"uint128\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"serviceFee\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes4\",\"name\":\"interfaceId\",\"type\":\"bytes4\"}],\"name\":\"supportsInterface\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]"

// SaleBin is the compiled bytecode used for deploying new contracts.
var SaleBin = "0x608060405234801561001057600080fd5b506113c0806100206000396000f3fe60806040526004361061009c5760003560e01c806398968f151161006457806398968f15146101c0578063a550e4b7146101d3578063bc197c81146101f3578063bdb818a514610254578063c28d81ae14610274578063f23a6e61146102945761009c565b806301ffc9a7146100a15780632992099f146100d65780633a1c83ac146100f8578063895047cc1461017c5780638abdf5aa146101aa575b600080fd5b3480156100ad57600080fd5b506100c16100bc3660046111a0565b6102da565b60405190151581526020015b60405180910390f35b3480156100e257600080fd5b506100f66100f136600461116c565b6102e2565b005b34801561010457600080fd5b5061014f6101133660046110fc565b60006020818152928152604080822090935290815220805460018201546002909201546001600160a01b03909116916001600160801b03169083565b604080516001600160a01b0390941684526001600160801b039092166020840152908201526060016100cd565b34801561018857600080fd5b5061019c61019736600461100c565b610445565b6040519081526020016100cd565b3480156101b657600080fd5b5061019c60015481565b6100f66101ce36600461116c565b61062d565b3480156101df57600080fd5b506100c16101ee3660046110fc565b6109dd565b3480156101ff57600080fd5b5061023b61020e366004610e7f565b7fbc197c819b3e337a6f9652dd10becd7eef83032af3b9d958d3d42f669414662198975050505050505050565b6040516001600160e01b031990911681526020016100cd565b34801561026057600080fd5b506100f661026f366004611127565b610a37565b34801561028057600080fd5b506100f661028f366004610f3a565b610b5e565b3480156102a057600080fd5b5061023b6102af366004610f92565b7ff23a6e612e1ff4830e658fe43f4e3cb4a5f8170bd5d9e69fb5d7a7fa9e4fdf979695505050505050565b60015b919050565b6001600160a01b0383811660009081526020818152604080832086845290915290208054909116331461035c5760405162461bcd60e51b815260206004820152601860248201527f726571756573746572206e6f7420617574686f72697a6564000000000000000060448201526064015b60405180910390fd5b8054604051637921219560e11b815285916001600160a01b038084169263f242432a92610394923092911690899089906004016111fc565b600060405180830381600087803b1580156103ae57600080fd5b505af11580156103c2573d6000803e3d6000fd5b505050506001600160a01b038516600081815260208181526040808320888452825280832080546001600160a01b03191681556001810180546001600160801b03191690556002019290925590518681527f9aab328f565aee990ccdda0fad461403cefd35ef76e04e7c1df569a5dd633caf910160405180910390a25050505050565b60405163622077af60e11b8152600090869082906001600160a01b0383169063c440ef5e9061047e908a908a908a903390600401611234565b602060405180830381600087803b15801561049857600080fd5b505af11580156104ac573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906104d091906111e4565b905060006040518060600160405280336001600160a01b03168152602001866001600160801b031681526020018867ffffffffffffffff168152509050806000808b6001600160a01b03166001600160a01b03168152602001908152602001600020600084815260200190815260200160002060008201518160000160006101000a8154816001600160a01b0302191690836001600160a01b0316021790555060208201518160010160006101000a8154816001600160801b0302191690836001600160801b031602179055506040820151816002015590505081896001600160a01b03167fe23abbfa2abc833198e1ff7fed44fad6e2a1c2218def757f8655458c0f2c4fd487338b604051610617939291906001600160801b039390931683526001600160a01b0391909116602083015267ffffffffffffffff16604082015260600190565b60405180910390a3509150505b95945050505050565b6001600160a01b038084166000908152602081815260408083208684529091529020805460028201549192169083111561069f5760405162461bcd60e51b81526020600482015260136024820152721a5b9cdd59999a58da595b9d08185b5bdd5b9d606a1b6044820152606401610353565b60018201546001600160801b03163410156106f15760405162461bcd60e51b815260206004820152601260248201527134b739bab33334b1b4b2b73a1032ba3432b960711b6044820152606401610353565b604051637921219560e11b8152859081906001600160a01b0382169063f242432a9061072790309033908b908b906004016111fc565b600060405180830381600087803b15801561074157600080fd5b505af1158015610755573d6000803e3d6000fd5b505050508484600201546107699190611309565b600285018190556107ba576001600160a01b038716600090815260208181526040808320898452909152812080546001600160a01b03191681556001810180546001600160801b0319169055600201555b604051630a216a2b60e31b8152600481018790526000906001600160a01b0383169063510b51589060240160206040518083038186803b1580156107fd57600080fd5b505afa158015610811573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906108359190610e5c565b60405163c57dc23560e01b8152600481018990529091506000906001600160a01b0384169063c57dc2359060240160206040518083038186803b15801561087b57600080fd5b505afa15801561088f573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906108b391906111c8565b905060006108c3612710836112b5565b6108d19061ffff16346112ea565b905060006127106001546108e591906112d6565b6108ef90346112ea565b90506000816108fe8434611309565b6109089190611309565b6040519091506001600160a01b0386169084156108fc029085906000818181858888f19350505050158015610941573d6000803e3d6000fd5b506040516001600160a01b0389169082156108fc029083906000818181858888f19350505050158015610978573d6000803e3d6000fd5b5060408051600081523360208201526001600160a01b038a811682840152606082018d905291518d928f16917f23df4888276ddbd130e02bf1325720ca1f4cb8a8101465fb800345716f01c790919081900360800190a3505050505050505050505050565b6001600160a01b0380831660009081526020818152604080832085845290915281208054919290911615801590610a20575060018101546001600160801b031615155b8015610a2f5750600281015415155b949350505050565b610a4384338584610c8a565b610a895760405162461bcd60e51b81526020600482015260176024820152760696e76616c696420746f6b656e206f776e65727368697604c1b6044820152606401610353565b610a9584338584610d2c565b6040805160608082018352338083526001600160801b0386811660208086018281528688018981526001600160a01b038d811660008181528086528b81208f825286528b90208a5181546001600160a01b031916931692909217825592516001820180546001600160801b03191691909716179095555160029094019390935586519182528101929092529381018590529192869290917fe23abbfa2abc833198e1ff7fed44fad6e2a1c2218def757f8655458c0f2c4fd4910160405180910390a35050505050565b610b6a85858584610c8a565b610bb05760405162461bcd60e51b81526020600482015260176024820152760696e76616c696420746f6b656e206f776e65727368697604c1b6044820152606401610353565b610bbc85858584610d2c565b60408051606080820183526001600160a01b038781168084526001600160801b0387811660208087018281528789018a81528e871660008181528085528b81208f825285528b90208a5181546001600160a01b031916991698909817885591516001880180546001600160801b031916919096161790945592516002909501949094558651908152928301919091529381018590529192869290917fe23abbfa2abc833198e1ff7fed44fad6e2a1c2218def757f8655458c0f2c4fd4910160405180910390a3505050505050565b604051627eeac760e11b81526001600160a01b03848116600483015260248201849052600091869183919083169062fdd58e9060440160206040518083038186803b158015610cd857600080fd5b505afa158015610cec573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610d1091906111e4565b9050838110158015610d2157508015155b979650505050505050565b604051637921219560e11b815284906001600160a01b0382169063f242432a90610d609087903090889088906004016111fc565b600060405180830381600087803b158015610d7a57600080fd5b505af1158015610d8e573d6000803e3d6000fd5b505050505050505050565b60008083601f840112610daa578081fd5b50813567ffffffffffffffff811115610dc1578182fd5b6020830191508360208083028501011115610ddb57600080fd5b9250929050565b60008083601f840112610df3578182fd5b50813567ffffffffffffffff811115610e0a578182fd5b602083019150836020828501011115610ddb57600080fd5b80356001600160801b03811681146102dd57600080fd5b80356102dd8161137a565b803567ffffffffffffffff811681146102dd57600080fd5b600060208284031215610e6d578081fd5b8151610e7881611362565b9392505050565b60008060008060008060008060a0898b031215610e9a578384fd5b8835610ea581611362565b97506020890135610eb581611362565b9650604089013567ffffffffffffffff80821115610ed1578586fd5b610edd8c838d01610d99565b909850965060608b0135915080821115610ef5578586fd5b610f018c838d01610d99565b909650945060808b0135915080821115610f19578384fd5b50610f268b828c01610de2565b999c989b5096995094979396929594505050565b600080600080600060a08688031215610f51578081fd5b8535610f5c81611362565b94506020860135610f6c81611362565b935060408601359250610f8160608701610e22565b949793965091946080013592915050565b60008060008060008060a08789031215610faa578182fd5b8635610fb581611362565b95506020870135610fc581611362565b94506040870135935060608701359250608087013567ffffffffffffffff811115610fee578283fd5b610ffa89828a01610de2565b979a9699509497509295939492505050565b600080600080600060a08688031215611023578081fd5b853561102e81611362565b9450602086013567ffffffffffffffff8082111561104a578283fd5b818801915088601f83011261105d578283fd5b81358181111561106f5761106f61134c565b604051601f8201601f19908116603f011681019083821181831017156110975761109761134c565b816040528281528b60208487010111156110af578586fd5b826020860160208301379182016020018590525095506110d491505060408701610e44565b92506110e260608701610e39565b91506110f060808701610e22565b90509295509295909350565b6000806040838503121561110e578182fd5b823561111981611362565b946020939093013593505050565b6000806000806080858703121561113c578384fd5b843561114781611362565b93506020850135925061115c60408601610e22565b9396929550929360600135925050565b600080600060608486031215611180578081fd5b833561118b81611362565b95602085013595506040909401359392505050565b6000602082840312156111b1578081fd5b81356001600160e01b031981168114610e78578182fd5b6000602082840312156111d9578081fd5b8151610e788161137a565b6000602082840312156111f5578081fd5b5051919050565b6001600160a01b0394851681529290931660208301526040820152606081019190915260a06080820181905260009082015260c00190565b6000608082528551806080840152815b8181101561126157602081890181015160a0868401015201611244565b81811115611272578260a083860101525b50601f01601f1916820160a0019050611297602083018667ffffffffffffffff169052565b61ffff841660408301526001600160a01b0383166060830152610624565b600061ffff808416806112ca576112ca611336565b92169190910492915050565b6000826112e5576112e5611336565b500490565b600081600019048311821515161561130457611304611320565b500290565b60008282101561131b5761131b611320565b500390565b634e487b7160e01b600052601160045260246000fd5b634e487b7160e01b600052601260045260246000fd5b634e487b7160e01b600052604160045260246000fd5b6001600160a01b038116811461137757600080fd5b50565b61ffff8116811461137757600080fdfea264697066735822122060db1d9d5979a0c0133512c83f7e4e6443a980040ddcb438b08ae9a83bf6687a64736f6c63430008020033"

// DeploySale deploys a new Ethereum contract, binding an instance of Sale to it.
func DeploySale(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *Sale, error) {
	parsed, err := abi.JSON(strings.NewReader(SaleABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}

	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(SaleBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Sale{SaleCaller: SaleCaller{contract: contract}, SaleTransactor: SaleTransactor{contract: contract}, SaleFilterer: SaleFilterer{contract: contract}}, nil
}

// Sale is an auto generated Go binding around an Ethereum contract.
type Sale struct {
	SaleCaller     // Read-only binding to the contract
	SaleTransactor // Write-only binding to the contract
	SaleFilterer   // Log filterer for contract events
}

// SaleCaller is an auto generated read-only Go binding around an Ethereum contract.
type SaleCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SaleTransactor is an auto generated write-only Go binding around an Ethereum contract.
type SaleTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SaleFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type SaleFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SaleSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type SaleSession struct {
	Contract     *Sale             // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// SaleCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type SaleCallerSession struct {
	Contract *SaleCaller   // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// SaleTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type SaleTransactorSession struct {
	Contract     *SaleTransactor   // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// SaleRaw is an auto generated low-level Go binding around an Ethereum contract.
type SaleRaw struct {
	Contract *Sale // Generic contract binding to access the raw methods on
}

// SaleCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type SaleCallerRaw struct {
	Contract *SaleCaller // Generic read-only contract binding to access the raw methods on
}

// SaleTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type SaleTransactorRaw struct {
	Contract *SaleTransactor // Generic write-only contract binding to access the raw methods on
}

// NewSale creates a new instance of Sale, bound to a specific deployed contract.
func NewSale(address common.Address, backend bind.ContractBackend) (*Sale, error) {
	contract, err := bindSale(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Sale{SaleCaller: SaleCaller{contract: contract}, SaleTransactor: SaleTransactor{contract: contract}, SaleFilterer: SaleFilterer{contract: contract}}, nil
}

// NewSaleCaller creates a new read-only instance of Sale, bound to a specific deployed contract.
func NewSaleCaller(address common.Address, caller bind.ContractCaller) (*SaleCaller, error) {
	contract, err := bindSale(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &SaleCaller{contract: contract}, nil
}

// NewSaleTransactor creates a new write-only instance of Sale, bound to a specific deployed contract.
func NewSaleTransactor(address common.Address, transactor bind.ContractTransactor) (*SaleTransactor, error) {
	contract, err := bindSale(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &SaleTransactor{contract: contract}, nil
}

// NewSaleFilterer creates a new log filterer instance of Sale, bound to a specific deployed contract.
func NewSaleFilterer(address common.Address, filterer bind.ContractFilterer) (*SaleFilterer, error) {
	contract, err := bindSale(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &SaleFilterer{contract: contract}, nil
}

// bindSale binds a generic wrapper to an already deployed contract.
func bindSale(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(SaleABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Sale *SaleRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Sale.Contract.SaleCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Sale *SaleRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Sale.Contract.SaleTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Sale *SaleRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Sale.Contract.SaleTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Sale *SaleCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Sale.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Sale *SaleTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Sale.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Sale *SaleTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Sale.Contract.contract.Transact(opts, method, params...)
}

// IsOnSale is a free data retrieval call binding the contract method 0xa550e4b7.
//
// Solidity: function isOnSale(address _contractAddress, uint256 _tokenId) view returns(bool)
func (_Sale *SaleCaller) IsOnSale(opts *bind.CallOpts, _contractAddress common.Address, _tokenId *big.Int) (bool, error) {
	var out []interface{}
	err := _Sale.contract.Call(opts, &out, "isOnSale", _contractAddress, _tokenId)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsOnSale is a free data retrieval call binding the contract method 0xa550e4b7.
//
// Solidity: function isOnSale(address _contractAddress, uint256 _tokenId) view returns(bool)
func (_Sale *SaleSession) IsOnSale(_contractAddress common.Address, _tokenId *big.Int) (bool, error) {
	return _Sale.Contract.IsOnSale(&_Sale.CallOpts, _contractAddress, _tokenId)
}

// IsOnSale is a free data retrieval call binding the contract method 0xa550e4b7.
//
// Solidity: function isOnSale(address _contractAddress, uint256 _tokenId) view returns(bool)
func (_Sale *SaleCallerSession) IsOnSale(_contractAddress common.Address, _tokenId *big.Int) (bool, error) {
	return _Sale.Contract.IsOnSale(&_Sale.CallOpts, _contractAddress, _tokenId)
}

// OnERC1155BatchReceived is a free data retrieval call binding the contract method 0xbc197c81.
//
// Solidity: function onERC1155BatchReceived(address , address , uint256[] , uint256[] , bytes ) pure returns(bytes4)
func (_Sale *SaleCaller) OnERC1155BatchReceived(opts *bind.CallOpts, arg0 common.Address, arg1 common.Address, arg2 []*big.Int, arg3 []*big.Int, arg4 []byte) ([4]byte, error) {
	var out []interface{}
	err := _Sale.contract.Call(opts, &out, "onERC1155BatchReceived", arg0, arg1, arg2, arg3, arg4)

	if err != nil {
		return *new([4]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([4]byte)).(*[4]byte)

	return out0, err

}

// OnERC1155BatchReceived is a free data retrieval call binding the contract method 0xbc197c81.
//
// Solidity: function onERC1155BatchReceived(address , address , uint256[] , uint256[] , bytes ) pure returns(bytes4)
func (_Sale *SaleSession) OnERC1155BatchReceived(arg0 common.Address, arg1 common.Address, arg2 []*big.Int, arg3 []*big.Int, arg4 []byte) ([4]byte, error) {
	return _Sale.Contract.OnERC1155BatchReceived(&_Sale.CallOpts, arg0, arg1, arg2, arg3, arg4)
}

// OnERC1155BatchReceived is a free data retrieval call binding the contract method 0xbc197c81.
//
// Solidity: function onERC1155BatchReceived(address , address , uint256[] , uint256[] , bytes ) pure returns(bytes4)
func (_Sale *SaleCallerSession) OnERC1155BatchReceived(arg0 common.Address, arg1 common.Address, arg2 []*big.Int, arg3 []*big.Int, arg4 []byte) ([4]byte, error) {
	return _Sale.Contract.OnERC1155BatchReceived(&_Sale.CallOpts, arg0, arg1, arg2, arg3, arg4)
}

// OnERC1155Received is a free data retrieval call binding the contract method 0xf23a6e61.
//
// Solidity: function onERC1155Received(address , address , uint256 , uint256 , bytes ) pure returns(bytes4)
func (_Sale *SaleCaller) OnERC1155Received(opts *bind.CallOpts, arg0 common.Address, arg1 common.Address, arg2 *big.Int, arg3 *big.Int, arg4 []byte) ([4]byte, error) {
	var out []interface{}
	err := _Sale.contract.Call(opts, &out, "onERC1155Received", arg0, arg1, arg2, arg3, arg4)

	if err != nil {
		return *new([4]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([4]byte)).(*[4]byte)

	return out0, err

}

// OnERC1155Received is a free data retrieval call binding the contract method 0xf23a6e61.
//
// Solidity: function onERC1155Received(address , address , uint256 , uint256 , bytes ) pure returns(bytes4)
func (_Sale *SaleSession) OnERC1155Received(arg0 common.Address, arg1 common.Address, arg2 *big.Int, arg3 *big.Int, arg4 []byte) ([4]byte, error) {
	return _Sale.Contract.OnERC1155Received(&_Sale.CallOpts, arg0, arg1, arg2, arg3, arg4)
}

// OnERC1155Received is a free data retrieval call binding the contract method 0xf23a6e61.
//
// Solidity: function onERC1155Received(address , address , uint256 , uint256 , bytes ) pure returns(bytes4)
func (_Sale *SaleCallerSession) OnERC1155Received(arg0 common.Address, arg1 common.Address, arg2 *big.Int, arg3 *big.Int, arg4 []byte) ([4]byte, error) {
	return _Sale.Contract.OnERC1155Received(&_Sale.CallOpts, arg0, arg1, arg2, arg3, arg4)
}

// Sales is a free data retrieval call binding the contract method 0x3a1c83ac.
//
// Solidity: function sales(address , uint256 ) view returns(address seller, uint128 price, uint256 amount)
func (_Sale *SaleCaller) Sales(opts *bind.CallOpts, arg0 common.Address, arg1 *big.Int) (struct {
	Seller common.Address
	Price  *big.Int
	Amount *big.Int
}, error) {
	var out []interface{}
	err := _Sale.contract.Call(opts, &out, "sales", arg0, arg1)

	outstruct := new(struct {
		Seller common.Address
		Price  *big.Int
		Amount *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Seller = *abi.ConvertType(out[0], new(common.Address)).(*common.Address)
	outstruct.Price = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)
	outstruct.Amount = *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// Sales is a free data retrieval call binding the contract method 0x3a1c83ac.
//
// Solidity: function sales(address , uint256 ) view returns(address seller, uint128 price, uint256 amount)
func (_Sale *SaleSession) Sales(arg0 common.Address, arg1 *big.Int) (struct {
	Seller common.Address
	Price  *big.Int
	Amount *big.Int
}, error) {
	return _Sale.Contract.Sales(&_Sale.CallOpts, arg0, arg1)
}

// Sales is a free data retrieval call binding the contract method 0x3a1c83ac.
//
// Solidity: function sales(address , uint256 ) view returns(address seller, uint128 price, uint256 amount)
func (_Sale *SaleCallerSession) Sales(arg0 common.Address, arg1 *big.Int) (struct {
	Seller common.Address
	Price  *big.Int
	Amount *big.Int
}, error) {
	return _Sale.Contract.Sales(&_Sale.CallOpts, arg0, arg1)
}

// ServiceFee is a free data retrieval call binding the contract method 0x8abdf5aa.
//
// Solidity: function serviceFee() view returns(uint256)
func (_Sale *SaleCaller) ServiceFee(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Sale.contract.Call(opts, &out, "serviceFee")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// ServiceFee is a free data retrieval call binding the contract method 0x8abdf5aa.
//
// Solidity: function serviceFee() view returns(uint256)
func (_Sale *SaleSession) ServiceFee() (*big.Int, error) {
	return _Sale.Contract.ServiceFee(&_Sale.CallOpts)
}

// ServiceFee is a free data retrieval call binding the contract method 0x8abdf5aa.
//
// Solidity: function serviceFee() view returns(uint256)
func (_Sale *SaleCallerSession) ServiceFee() (*big.Int, error) {
	return _Sale.Contract.ServiceFee(&_Sale.CallOpts)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_Sale *SaleCaller) SupportsInterface(opts *bind.CallOpts, interfaceId [4]byte) (bool, error) {
	var out []interface{}
	err := _Sale.contract.Call(opts, &out, "supportsInterface", interfaceId)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_Sale *SaleSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _Sale.Contract.SupportsInterface(&_Sale.CallOpts, interfaceId)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_Sale *SaleCallerSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _Sale.Contract.SupportsInterface(&_Sale.CallOpts, interfaceId)
}

// CancelSale is a paid mutator transaction binding the contract method 0x2992099f.
//
// Solidity: function cancelSale(address _contractAddress, uint256 _tokenId, uint256 _amount) returns()
func (_Sale *SaleTransactor) CancelSale(opts *bind.TransactOpts, _contractAddress common.Address, _tokenId *big.Int, _amount *big.Int) (*types.Transaction, error) {
	return _Sale.contract.Transact(opts, "cancelSale", _contractAddress, _tokenId, _amount)
}

// CancelSale is a paid mutator transaction binding the contract method 0x2992099f.
//
// Solidity: function cancelSale(address _contractAddress, uint256 _tokenId, uint256 _amount) returns()
func (_Sale *SaleSession) CancelSale(_contractAddress common.Address, _tokenId *big.Int, _amount *big.Int) (*types.Transaction, error) {
	return _Sale.Contract.CancelSale(&_Sale.TransactOpts, _contractAddress, _tokenId, _amount)
}

// CancelSale is a paid mutator transaction binding the contract method 0x2992099f.
//
// Solidity: function cancelSale(address _contractAddress, uint256 _tokenId, uint256 _amount) returns()
func (_Sale *SaleTransactorSession) CancelSale(_contractAddress common.Address, _tokenId *big.Int, _amount *big.Int) (*types.Transaction, error) {
	return _Sale.Contract.CancelSale(&_Sale.TransactOpts, _contractAddress, _tokenId, _amount)
}

// CreateAndSell is a paid mutator transaction binding the contract method 0x895047cc.
//
// Solidity: function createAndSell(address _contractAddress, bytes id, uint64 _amount, uint16 __royaltyFee, uint128 _price) returns(uint256)
func (_Sale *SaleTransactor) CreateAndSell(opts *bind.TransactOpts, _contractAddress common.Address, id []byte, _amount uint64, __royaltyFee uint16, _price *big.Int) (*types.Transaction, error) {
	return _Sale.contract.Transact(opts, "createAndSell", _contractAddress, id, _amount, __royaltyFee, _price)
}

// CreateAndSell is a paid mutator transaction binding the contract method 0x895047cc.
//
// Solidity: function createAndSell(address _contractAddress, bytes id, uint64 _amount, uint16 __royaltyFee, uint128 _price) returns(uint256)
func (_Sale *SaleSession) CreateAndSell(_contractAddress common.Address, id []byte, _amount uint64, __royaltyFee uint16, _price *big.Int) (*types.Transaction, error) {
	return _Sale.Contract.CreateAndSell(&_Sale.TransactOpts, _contractAddress, id, _amount, __royaltyFee, _price)
}

// CreateAndSell is a paid mutator transaction binding the contract method 0x895047cc.
//
// Solidity: function createAndSell(address _contractAddress, bytes id, uint64 _amount, uint16 __royaltyFee, uint128 _price) returns(uint256)
func (_Sale *SaleTransactorSession) CreateAndSell(_contractAddress common.Address, id []byte, _amount uint64, __royaltyFee uint16, _price *big.Int) (*types.Transaction, error) {
	return _Sale.Contract.CreateAndSell(&_Sale.TransactOpts, _contractAddress, id, _amount, __royaltyFee, _price)
}

// CreateSale is a paid mutator transaction binding the contract method 0xbdb818a5.
//
// Solidity: function createSale(address _contractAddress, uint256 _tokenId, uint128 _price, uint256 _amount) returns()
func (_Sale *SaleTransactor) CreateSale(opts *bind.TransactOpts, _contractAddress common.Address, _tokenId *big.Int, _price *big.Int, _amount *big.Int) (*types.Transaction, error) {
	return _Sale.contract.Transact(opts, "createSale", _contractAddress, _tokenId, _price, _amount)
}

// CreateSale is a paid mutator transaction binding the contract method 0xbdb818a5.
//
// Solidity: function createSale(address _contractAddress, uint256 _tokenId, uint128 _price, uint256 _amount) returns()
func (_Sale *SaleSession) CreateSale(_contractAddress common.Address, _tokenId *big.Int, _price *big.Int, _amount *big.Int) (*types.Transaction, error) {
	return _Sale.Contract.CreateSale(&_Sale.TransactOpts, _contractAddress, _tokenId, _price, _amount)
}

// CreateSale is a paid mutator transaction binding the contract method 0xbdb818a5.
//
// Solidity: function createSale(address _contractAddress, uint256 _tokenId, uint128 _price, uint256 _amount) returns()
func (_Sale *SaleTransactorSession) CreateSale(_contractAddress common.Address, _tokenId *big.Int, _price *big.Int, _amount *big.Int) (*types.Transaction, error) {
	return _Sale.Contract.CreateSale(&_Sale.TransactOpts, _contractAddress, _tokenId, _price, _amount)
}

// CreateSaleFromContract is a paid mutator transaction binding the contract method 0xc28d81ae.
//
// Solidity: function createSaleFromContract(address _contractAddress, address _seller, uint256 _tokenId, uint128 _price, uint256 _amount) returns()
func (_Sale *SaleTransactor) CreateSaleFromContract(opts *bind.TransactOpts, _contractAddress common.Address, _seller common.Address, _tokenId *big.Int, _price *big.Int, _amount *big.Int) (*types.Transaction, error) {
	return _Sale.contract.Transact(opts, "createSaleFromContract", _contractAddress, _seller, _tokenId, _price, _amount)
}

// CreateSaleFromContract is a paid mutator transaction binding the contract method 0xc28d81ae.
//
// Solidity: function createSaleFromContract(address _contractAddress, address _seller, uint256 _tokenId, uint128 _price, uint256 _amount) returns()
func (_Sale *SaleSession) CreateSaleFromContract(_contractAddress common.Address, _seller common.Address, _tokenId *big.Int, _price *big.Int, _amount *big.Int) (*types.Transaction, error) {
	return _Sale.Contract.CreateSaleFromContract(&_Sale.TransactOpts, _contractAddress, _seller, _tokenId, _price, _amount)
}

// CreateSaleFromContract is a paid mutator transaction binding the contract method 0xc28d81ae.
//
// Solidity: function createSaleFromContract(address _contractAddress, address _seller, uint256 _tokenId, uint128 _price, uint256 _amount) returns()
func (_Sale *SaleTransactorSession) CreateSaleFromContract(_contractAddress common.Address, _seller common.Address, _tokenId *big.Int, _price *big.Int, _amount *big.Int) (*types.Transaction, error) {
	return _Sale.Contract.CreateSaleFromContract(&_Sale.TransactOpts, _contractAddress, _seller, _tokenId, _price, _amount)
}

// Purchase is a paid mutator transaction binding the contract method 0x98968f15.
//
// Solidity: function purchase(address _contractAddress, uint256 _tokenId, uint256 _amount) payable returns()
func (_Sale *SaleTransactor) Purchase(opts *bind.TransactOpts, _contractAddress common.Address, _tokenId *big.Int, _amount *big.Int) (*types.Transaction, error) {
	return _Sale.contract.Transact(opts, "purchase", _contractAddress, _tokenId, _amount)
}

// Purchase is a paid mutator transaction binding the contract method 0x98968f15.
//
// Solidity: function purchase(address _contractAddress, uint256 _tokenId, uint256 _amount) payable returns()
func (_Sale *SaleSession) Purchase(_contractAddress common.Address, _tokenId *big.Int, _amount *big.Int) (*types.Transaction, error) {
	return _Sale.Contract.Purchase(&_Sale.TransactOpts, _contractAddress, _tokenId, _amount)
}

// Purchase is a paid mutator transaction binding the contract method 0x98968f15.
//
// Solidity: function purchase(address _contractAddress, uint256 _tokenId, uint256 _amount) payable returns()
func (_Sale *SaleTransactorSession) Purchase(_contractAddress common.Address, _tokenId *big.Int, _amount *big.Int) (*types.Transaction, error) {
	return _Sale.Contract.Purchase(&_Sale.TransactOpts, _contractAddress, _tokenId, _amount)
}

// SaleSaleCanceledIterator is returned from FilterSaleCanceled and is used to iterate over the raw logs and unpacked data for SaleCanceled events raised by the Sale contract.
type SaleSaleCanceledIterator struct {
	Event *SaleSaleCanceled // Event containing the contract specifics and raw log

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
func (it *SaleSaleCanceledIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SaleSaleCanceled)
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
		it.Event = new(SaleSaleCanceled)
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
func (it *SaleSaleCanceledIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SaleSaleCanceledIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SaleSaleCanceled represents a SaleCanceled event raised by the Sale contract.
type SaleSaleCanceled struct {
	ContractAddress common.Address
	TokenId         *big.Int
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterSaleCanceled is a free log retrieval operation binding the contract event 0x9aab328f565aee990ccdda0fad461403cefd35ef76e04e7c1df569a5dd633caf.
//
// Solidity: event SaleCanceled(address indexed _contractAddress, uint256 _tokenId)
func (_Sale *SaleFilterer) FilterSaleCanceled(opts *bind.FilterOpts, _contractAddress []common.Address) (*SaleSaleCanceledIterator, error) {

	var _contractAddressRule []interface{}
	for _, _contractAddressItem := range _contractAddress {
		_contractAddressRule = append(_contractAddressRule, _contractAddressItem)
	}

	logs, sub, err := _Sale.contract.FilterLogs(opts, "SaleCanceled", _contractAddressRule)
	if err != nil {
		return nil, err
	}
	return &SaleSaleCanceledIterator{contract: _Sale.contract, event: "SaleCanceled", logs: logs, sub: sub}, nil
}

// WatchSaleCanceled is a free log subscription operation binding the contract event 0x9aab328f565aee990ccdda0fad461403cefd35ef76e04e7c1df569a5dd633caf.
//
// Solidity: event SaleCanceled(address indexed _contractAddress, uint256 _tokenId)
func (_Sale *SaleFilterer) WatchSaleCanceled(opts *bind.WatchOpts, sink chan<- *SaleSaleCanceled, _contractAddress []common.Address) (event.Subscription, error) {

	var _contractAddressRule []interface{}
	for _, _contractAddressItem := range _contractAddress {
		_contractAddressRule = append(_contractAddressRule, _contractAddressItem)
	}

	logs, sub, err := _Sale.contract.WatchLogs(opts, "SaleCanceled", _contractAddressRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SaleSaleCanceled)
				if err := _Sale.contract.UnpackLog(event, "SaleCanceled", log); err != nil {
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

// ParseSaleCanceled is a log parse operation binding the contract event 0x9aab328f565aee990ccdda0fad461403cefd35ef76e04e7c1df569a5dd633caf.
//
// Solidity: event SaleCanceled(address indexed _contractAddress, uint256 _tokenId)
func (_Sale *SaleFilterer) ParseSaleCanceled(log types.Log) (*SaleSaleCanceled, error) {
	event := new(SaleSaleCanceled)
	if err := _Sale.contract.UnpackLog(event, "SaleCanceled", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SaleSaleCreatedIterator is returned from FilterSaleCreated and is used to iterate over the raw logs and unpacked data for SaleCreated events raised by the Sale contract.
type SaleSaleCreatedIterator struct {
	Event *SaleSaleCreated // Event containing the contract specifics and raw log

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
func (it *SaleSaleCreatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SaleSaleCreated)
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
		it.Event = new(SaleSaleCreated)
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
func (it *SaleSaleCreatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SaleSaleCreatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SaleSaleCreated represents a SaleCreated event raised by the Sale contract.
type SaleSaleCreated struct {
	ContractAddress common.Address
	TokenId         *big.Int
	Price           *big.Int
	Seller          common.Address
	Amount          *big.Int
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterSaleCreated is a free log retrieval operation binding the contract event 0xe23abbfa2abc833198e1ff7fed44fad6e2a1c2218def757f8655458c0f2c4fd4.
//
// Solidity: event SaleCreated(address indexed _contractAddress, uint256 indexed _tokenId, uint128 _price, address _seller, uint256 _amount)
func (_Sale *SaleFilterer) FilterSaleCreated(opts *bind.FilterOpts, _contractAddress []common.Address, _tokenId []*big.Int) (*SaleSaleCreatedIterator, error) {

	var _contractAddressRule []interface{}
	for _, _contractAddressItem := range _contractAddress {
		_contractAddressRule = append(_contractAddressRule, _contractAddressItem)
	}
	var _tokenIdRule []interface{}
	for _, _tokenIdItem := range _tokenId {
		_tokenIdRule = append(_tokenIdRule, _tokenIdItem)
	}

	logs, sub, err := _Sale.contract.FilterLogs(opts, "SaleCreated", _contractAddressRule, _tokenIdRule)
	if err != nil {
		return nil, err
	}
	return &SaleSaleCreatedIterator{contract: _Sale.contract, event: "SaleCreated", logs: logs, sub: sub}, nil
}

// WatchSaleCreated is a free log subscription operation binding the contract event 0xe23abbfa2abc833198e1ff7fed44fad6e2a1c2218def757f8655458c0f2c4fd4.
//
// Solidity: event SaleCreated(address indexed _contractAddress, uint256 indexed _tokenId, uint128 _price, address _seller, uint256 _amount)
func (_Sale *SaleFilterer) WatchSaleCreated(opts *bind.WatchOpts, sink chan<- *SaleSaleCreated, _contractAddress []common.Address, _tokenId []*big.Int) (event.Subscription, error) {

	var _contractAddressRule []interface{}
	for _, _contractAddressItem := range _contractAddress {
		_contractAddressRule = append(_contractAddressRule, _contractAddressItem)
	}
	var _tokenIdRule []interface{}
	for _, _tokenIdItem := range _tokenId {
		_tokenIdRule = append(_tokenIdRule, _tokenIdItem)
	}

	logs, sub, err := _Sale.contract.WatchLogs(opts, "SaleCreated", _contractAddressRule, _tokenIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SaleSaleCreated)
				if err := _Sale.contract.UnpackLog(event, "SaleCreated", log); err != nil {
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

// ParseSaleCreated is a log parse operation binding the contract event 0xe23abbfa2abc833198e1ff7fed44fad6e2a1c2218def757f8655458c0f2c4fd4.
//
// Solidity: event SaleCreated(address indexed _contractAddress, uint256 indexed _tokenId, uint128 _price, address _seller, uint256 _amount)
func (_Sale *SaleFilterer) ParseSaleCreated(log types.Log) (*SaleSaleCreated, error) {
	event := new(SaleSaleCreated)
	if err := _Sale.contract.UnpackLog(event, "SaleCreated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SaleSaleSuccessfulIterator is returned from FilterSaleSuccessful and is used to iterate over the raw logs and unpacked data for SaleSuccessful events raised by the Sale contract.
type SaleSaleSuccessfulIterator struct {
	Event *SaleSaleSuccessful // Event containing the contract specifics and raw log

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
func (it *SaleSaleSuccessfulIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SaleSaleSuccessful)
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
		it.Event = new(SaleSaleSuccessful)
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
func (it *SaleSaleSuccessfulIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SaleSaleSuccessfulIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SaleSaleSuccessful represents a SaleSuccessful event raised by the Sale contract.
type SaleSaleSuccessful struct {
	ContractAddress common.Address
	TokenId         *big.Int
	Price           *big.Int
	Buyer           common.Address
	Seller          common.Address
	Amount          *big.Int
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterSaleSuccessful is a free log retrieval operation binding the contract event 0x23df4888276ddbd130e02bf1325720ca1f4cb8a8101465fb800345716f01c790.
//
// Solidity: event SaleSuccessful(address indexed _contractAddress, uint256 indexed _tokenId, uint128 _price, address _buyer, address _seller, uint256 _amount)
func (_Sale *SaleFilterer) FilterSaleSuccessful(opts *bind.FilterOpts, _contractAddress []common.Address, _tokenId []*big.Int) (*SaleSaleSuccessfulIterator, error) {

	var _contractAddressRule []interface{}
	for _, _contractAddressItem := range _contractAddress {
		_contractAddressRule = append(_contractAddressRule, _contractAddressItem)
	}
	var _tokenIdRule []interface{}
	for _, _tokenIdItem := range _tokenId {
		_tokenIdRule = append(_tokenIdRule, _tokenIdItem)
	}

	logs, sub, err := _Sale.contract.FilterLogs(opts, "SaleSuccessful", _contractAddressRule, _tokenIdRule)
	if err != nil {
		return nil, err
	}
	return &SaleSaleSuccessfulIterator{contract: _Sale.contract, event: "SaleSuccessful", logs: logs, sub: sub}, nil
}

// WatchSaleSuccessful is a free log subscription operation binding the contract event 0x23df4888276ddbd130e02bf1325720ca1f4cb8a8101465fb800345716f01c790.
//
// Solidity: event SaleSuccessful(address indexed _contractAddress, uint256 indexed _tokenId, uint128 _price, address _buyer, address _seller, uint256 _amount)
func (_Sale *SaleFilterer) WatchSaleSuccessful(opts *bind.WatchOpts, sink chan<- *SaleSaleSuccessful, _contractAddress []common.Address, _tokenId []*big.Int) (event.Subscription, error) {

	var _contractAddressRule []interface{}
	for _, _contractAddressItem := range _contractAddress {
		_contractAddressRule = append(_contractAddressRule, _contractAddressItem)
	}
	var _tokenIdRule []interface{}
	for _, _tokenIdItem := range _tokenId {
		_tokenIdRule = append(_tokenIdRule, _tokenIdItem)
	}

	logs, sub, err := _Sale.contract.WatchLogs(opts, "SaleSuccessful", _contractAddressRule, _tokenIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SaleSaleSuccessful)
				if err := _Sale.contract.UnpackLog(event, "SaleSuccessful", log); err != nil {
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

// ParseSaleSuccessful is a log parse operation binding the contract event 0x23df4888276ddbd130e02bf1325720ca1f4cb8a8101465fb800345716f01c790.
//
// Solidity: event SaleSuccessful(address indexed _contractAddress, uint256 indexed _tokenId, uint128 _price, address _buyer, address _seller, uint256 _amount)
func (_Sale *SaleFilterer) ParseSaleSuccessful(log types.Log) (*SaleSaleSuccessful, error) {
	event := new(SaleSaleSuccessful)
	if err := _Sale.contract.UnpackLog(event, "SaleSuccessful", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
