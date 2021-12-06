// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package core

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

// CoreABI is the input ABI used to generate the binding from.
const CoreABI = "[{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"approved\",\"type\":\"bool\"}],\"name\":\"ApprovalForAll\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256[]\",\"name\":\"ids\",\"type\":\"uint256[]\"},{\"indexed\":false,\"internalType\":\"uint256[]\",\"name\":\"values\",\"type\":\"uint256[]\"}],\"name\":\"TransferBatch\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"TransferSingle\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"string\",\"name\":\"value\",\"type\":\"string\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"}],\"name\":\"URI\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"}],\"name\":\"balanceOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"accounts\",\"type\":\"address[]\"},{\"internalType\":\"uint256[]\",\"name\":\"ids\",\"type\":\"uint256[]\"}],\"name\":\"balanceOfBatch\",\"outputs\":[{\"internalType\":\"uint256[]\",\"name\":\"\",\"type\":\"uint256[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"id\",\"type\":\"bytes\"},{\"internalType\":\"uint64\",\"name\":\"amount\",\"type\":\"uint64\"},{\"internalType\":\"uint16\",\"name\":\"__royaltyFee\",\"type\":\"uint16\"}],\"name\":\"create\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"id\",\"type\":\"bytes\"},{\"internalType\":\"uint64\",\"name\":\"amount\",\"type\":\"uint64\"},{\"internalType\":\"uint16\",\"name\":\"__royaltyFee\",\"type\":\"uint16\"},{\"internalType\":\"address\",\"name\":\"__creator\",\"type\":\"address\"}],\"name\":\"createFromContract\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_tokenId\",\"type\":\"uint256\"}],\"name\":\"creator\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"}],\"name\":\"exists\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"}],\"name\":\"isApprovedForAll\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_tokenId\",\"type\":\"uint256\"}],\"name\":\"royaltyFee\",\"outputs\":[{\"internalType\":\"uint16\",\"name\":\"\",\"type\":\"uint16\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256[]\",\"name\":\"ids\",\"type\":\"uint256[]\"},{\"internalType\":\"uint256[]\",\"name\":\"amounts\",\"type\":\"uint256[]\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"safeBatchTransferFrom\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"safeTransferFrom\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"approved\",\"type\":\"bool\"}],\"name\":\"setApprovalForAll\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes4\",\"name\":\"interfaceId\",\"type\":\"bytes4\"}],\"name\":\"supportsInterface\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"}],\"name\":\"totalSupply\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"uri\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]"

// CoreBin is the compiled bytecode used for deploying new contracts.
var CoreBin = "0x608060405234801561001057600080fd5b50611b16806100206000396000f3fe608060405234801561001057600080fd5b50600436106100f45760003560e01c80637d36030f11610097578063c440ef5e11610066578063c440ef5e14610248578063c57dc2351461025b578063e985e9c514610292578063f242432a146102ce576100f4565b80637d36030f146101fa5780638129fc1c1461020d578063a22cb46514610215578063bd85b03914610228576100f4565b80632eb2c2d6116100d35780632eb2c2d6146101625780634e1273f4146101775780634f558e7914610197578063510b5158146101b9576100f4565b8062fdd58e146100f957806301ffc9a71461011f5780630e89341c14610142575b600080fd5b61010c61010736600461147a565b6102e1565b6040519081526020015b60405180910390f35b61013261012d36600461156b565b61037a565b6040519015158152602001610116565b610155610150366004611671565b6103ce565b604051610116919061180a565b610175610170366004611339565b610462565b005b61018a6101853660046114a3565b61050b565b60405161011691906117c9565b6101326101a5366004611671565b600090815260976020526040902054151590565b6101e26101c7366004611671565b600090815260c960205260409020546001600160a01b031690565b6040516001600160a01b039091168152602001610116565b61010c6102083660046115aa565b61066c565b6101756106f2565b610175610223366004611440565b6107d7565b61010c610236366004611671565b60009081526097602052604090205490565b61010c610256366004611605565b6108eb565b61027f610269366004611671565b600090815260ca602052604090205461ffff1690565b60405161ffff9091168152602001610116565b6101326102a0366004611307565b6001600160a01b03918216600090815260666020908152604080832093909416825291909152205460ff1690565b6101756102dc3660046113de565b61098c565b60006001600160a01b0383166103525760405162461bcd60e51b815260206004820152602b60248201527f455243313135353a2062616c616e636520717565727920666f7220746865207a60448201526a65726f206164647265737360a81b60648201526084015b60405180910390fd5b5060009081526065602090815260408083206001600160a01b03949094168352929052205490565b60006001600160e01b03198216636cdb3d1360e11b14806103ab57506001600160e01b031982166303a24d0760e21b145b806103c657506301ffc9a760e01b6001600160e01b03198316145b90505b919050565b6060606780546103dd90611974565b80601f016020809104026020016040519081016040528092919081815260200182805461040990611974565b80156104565780601f1061042b57610100808354040283529160200191610456565b820191906000526020600020905b81548152906001019060200180831161043957829003601f168201915b50505050509050919050565b61046a610a25565b6001600160a01b0316856001600160a01b031614806104905750610490856102a0610a25565b6104f75760405162461bcd60e51b815260206004820152603260248201527f455243313135353a207472616e736665722063616c6c6572206973206e6f74206044820152711bdddb995c881b9bdc88185c1c1c9bdd995960721b6064820152608401610349565b6105048585858585610a2a565b5050505050565b606081518351146105705760405162461bcd60e51b815260206004820152602960248201527f455243313135353a206163636f756e747320616e6420696473206c656e677468604482015268040dad2e6dac2e8c6d60bb1b6064820152608401610349565b600083516001600160401b0381111561059957634e487b7160e01b600052604160045260246000fd5b6040519080825280602002602001820160405280156105c2578160200160208202803683370190505b50905060005b8451811015610664576106298582815181106105f457634e487b7160e01b600052603260045260246000fd5b602002602001015185838151811061061c57634e487b7160e01b600052603260045260246000fd5b60200260200101516102e1565b82828151811061064957634e487b7160e01b600052603260045260246000fd5b602090810291909101015261065d816119db565b90506105c8565b509392505050565b60405163622077af60e11b8152600090309063c440ef5e9061069890879087908790339060040161181d565b602060405180830381600087803b1580156106b257600080fd5b505af11580156106c6573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906106ea9190611689565b949350505050565b600054610100900460ff168061070b575060005460ff16155b61076e5760405162461bcd60e51b815260206004820152602e60248201527f496e697469616c697a61626c653a20636f6e747261637420697320616c72656160448201526d191e481a5b9a5d1a585b1a5e995960921b6064820152608401610349565b600054610100900460ff16158015610799576000805460ff1961ff0019909116610100171660011790555b6107c26040518060400160405280600881526020016775726c2f7b69647d60c01b815250610c31565b80156107d4576000805461ff00191690555b50565b816001600160a01b03166107e9610a25565b6001600160a01b031614156108525760405162461bcd60e51b815260206004820152602960248201527f455243313135353a2073657474696e6720617070726f76616c20737461747573604482015268103337b91039b2b63360b91b6064820152608401610349565b806066600061085f610a25565b6001600160a01b03908116825260208083019390935260409182016000908120918716808252919093529120805460ff1916921515929092179091556108a3610a25565b6001600160a01b03167f17307eab39ab6107e8899845ad3d59bd9653f200f220920489ca2b5937696c31836040516108df911515815260200190565b60405180910390a35050565b600080856040516020016108ff919061180a565b6040516020818303038152906040528051906020012060001c905061093e3382876001600160401b031660405180602001604052806000815250610c48565b600081815260c96020908152604080832080546001600160a01b0388166001600160a01b031990911617905560ca9091529020805461ffff861661ffff199091161790559050949350505050565b610994610a25565b6001600160a01b0316856001600160a01b031614806109ba57506109ba856102a0610a25565b610a185760405162461bcd60e51b815260206004820152602960248201527f455243313135353a2063616c6c6572206973206e6f74206f776e6572206e6f7260448201526808185c1c1c9bdd995960ba1b6064820152608401610349565b6105048585858585610c7d565b335b90565b8151835114610a8c5760405162461bcd60e51b815260206004820152602860248201527f455243313135353a2069647320616e6420616d6f756e7473206c656e677468206044820152670dad2e6dac2e8c6d60c31b6064820152608401610349565b6001600160a01b038416610ab25760405162461bcd60e51b8152600401610349906118aa565b6000610abc610a25565b905060005b8451811015610bc3576000858281518110610aec57634e487b7160e01b600052603260045260246000fd5b602002602001015190506000858381518110610b1857634e487b7160e01b600052603260045260246000fd5b60209081029190910181015160008481526065835260408082206001600160a01b038e168352909352919091205490915081811015610b695760405162461bcd60e51b8152600401610349906118ef565b60008381526065602090815260408083206001600160a01b038e8116855292528083208585039055908b16825281208054849290610ba890849061195c565b9250508190555050505080610bbc906119db565b9050610ac1565b50846001600160a01b0316866001600160a01b0316826001600160a01b03167f4a39dc06d4c0dbc64b70af90fd698a233a518aa5d07e595d983b8c0526c8f7fb8787604051610c139291906117dc565b60405180910390a4610c29818787878787610db2565b505050505050565b8051610c4490606790602084019061114e565b5050565b610c5484848484610f1d565b60008381526097602052604081208054849290610c7290849061195c565b909155505050505050565b6001600160a01b038416610ca35760405162461bcd60e51b8152600401610349906118aa565b6000610cad610a25565b9050610cc7818787610cbe8861102b565b6105048861102b565b60008481526065602090815260408083206001600160a01b038a16845290915290205483811015610d0a5760405162461bcd60e51b8152600401610349906118ef565b60008581526065602090815260408083206001600160a01b038b8116855292528083208785039055908816825281208054869290610d4990849061195c565b909155505060408051868152602081018690526001600160a01b03808916928a821692918616917fc3d58168c5ae7397731d063d5bbf3d657854427343f4c083240f7aacaa2d0f62910160405180910390a4610da9828888888888611084565b50505050505050565b6001600160a01b0384163b15610c295760405163bc197c8160e01b81526001600160a01b0385169063bc197c8190610df69089908990889088908890600401611726565b602060405180830381600087803b158015610e1057600080fd5b505af1925050508015610e40575060408051601f3d908101601f19168201909252610e3d9181019061158e565b60015b610eed57610e4c611a22565b806308c379a01415610e865750610e61611a39565b80610e6c5750610e88565b8060405162461bcd60e51b8152600401610349919061180a565b505b60405162461bcd60e51b815260206004820152603460248201527f455243313135353a207472616e7366657220746f206e6f6e20455243313135356044820152732932b1b2b4bb32b91034b6b83632b6b2b73a32b960611b6064820152608401610349565b6001600160e01b0319811663bc197c8160e01b14610da95760405162461bcd60e51b815260040161034990611862565b6001600160a01b038416610f7d5760405162461bcd60e51b815260206004820152602160248201527f455243313135353a206d696e7420746f20746865207a65726f206164647265736044820152607360f81b6064820152608401610349565b6000610f87610a25565b9050610f9981600087610cbe8861102b565b60008481526065602090815260408083206001600160a01b038916845290915281208054859290610fcb90849061195c565b909155505060408051858152602081018590526001600160a01b0380881692600092918516917fc3d58168c5ae7397731d063d5bbf3d657854427343f4c083240f7aacaa2d0f62910160405180910390a461050481600087878787611084565b6040805160018082528183019092526060916000919060208083019080368337019050509050828160008151811061107357634e487b7160e01b600052603260045260246000fd5b602090810291909101015292915050565b6001600160a01b0384163b15610c295760405163f23a6e6160e01b81526001600160a01b0385169063f23a6e61906110c89089908990889088908890600401611784565b602060405180830381600087803b1580156110e257600080fd5b505af1925050508015611112575060408051601f3d908101601f1916820190925261110f9181019061158e565b60015b61111e57610e4c611a22565b6001600160e01b0319811663f23a6e6160e01b14610da95760405162461bcd60e51b815260040161034990611862565b82805461115a90611974565b90600052602060002090601f01602090048101928261117c57600085556111c2565b82601f1061119557805160ff19168380011785556111c2565b828001600101855582156111c2579182015b828111156111c25782518255916020019190600101906111a7565b506111ce9291506111d2565b5090565b5b808211156111ce57600081556001016111d3565b80356001600160a01b03811681146103c957600080fd5b600082601f83011261120e578081fd5b8135602061121b82611939565b60405161122882826119af565b838152828101915085830183850287018401881015611245578586fd5b855b8581101561126357813584529284019290840190600101611247565b5090979650505050505050565b600082601f830112611280578081fd5b81356001600160401b0381111561129957611299611a0c565b6040516112b0601f8301601f1916602001826119af565b8181528460208386010111156112c4578283fd5b816020850160208301379081016020019190915292915050565b803561ffff811681146103c957600080fd5b80356001600160401b03811681146103c957600080fd5b60008060408385031215611319578182fd5b611322836111e7565b9150611330602084016111e7565b90509250929050565b600080600080600060a08688031215611350578081fd5b611359866111e7565b9450611367602087016111e7565b935060408601356001600160401b0380821115611382578283fd5b61138e89838a016111fe565b945060608801359150808211156113a3578283fd5b6113af89838a016111fe565b935060808801359150808211156113c4578283fd5b506113d188828901611270565b9150509295509295909350565b600080600080600060a086880312156113f5578081fd5b6113fe866111e7565b945061140c602087016111e7565b9350604086013592506060860135915060808601356001600160401b03811115611434578182fd5b6113d188828901611270565b60008060408385031215611452578182fd5b61145b836111e7565b91506020830135801515811461146f578182fd5b809150509250929050565b6000806040838503121561148c578182fd5b611495836111e7565b946020939093013593505050565b600080604083850312156114b5578182fd5b82356001600160401b03808211156114cb578384fd5b818501915085601f8301126114de578384fd5b813560206114eb82611939565b6040516114f882826119af565b8381528281019150858301838502870184018b1015611515578889fd5b8896505b8487101561153e5761152a816111e7565b835260019690960195918301918301611519565b5096505086013592505080821115611554578283fd5b50611561858286016111fe565b9150509250929050565b60006020828403121561157c578081fd5b813561158781611aca565b9392505050565b60006020828403121561159f578081fd5b815161158781611aca565b6000806000606084860312156115be578081fd5b83356001600160401b038111156115d3578182fd5b6115df86828701611270565b9350506115ee602085016112f0565b91506115fc604085016112de565b90509250925092565b6000806000806080858703121561161a578182fd5b84356001600160401b0381111561162f578283fd5b61163b87828801611270565b94505061164a602086016112f0565b9250611658604086016112de565b9150611666606086016111e7565b905092959194509250565b600060208284031215611682578081fd5b5035919050565b60006020828403121561169a578081fd5b5051919050565b6000815180845260208085019450808401835b838110156116d0578151875295820195908201906001016116b4565b509495945050505050565b60008151808452815b81811015611700576020818501810151868301820152016116e4565b818111156117115782602083870101525b50601f01601f19169290920160200192915050565b6001600160a01b0386811682528516602082015260a060408201819052600090611752908301866116a1565b828103606084015261176481866116a1565b9050828103608084015261177881856116db565b98975050505050505050565b6001600160a01b03868116825285166020820152604081018490526060810183905260a0608082018190526000906117be908301846116db565b979650505050505050565b60006020825261158760208301846116a1565b6000604082526117ef60408301856116a1565b828103602084015261180181856116a1565b95945050505050565b60006020825261158760208301846116db565b60006080825261183060808301876116db565b6001600160401b039590951660208301525061ffff9290921660408301526001600160a01b0316606090910152919050565b60208082526028908201527f455243313135353a204552433131353552656365697665722072656a656374656040820152676420746f6b656e7360c01b606082015260800190565b60208082526025908201527f455243313135353a207472616e7366657220746f20746865207a65726f206164604082015264647265737360d81b606082015260800190565b6020808252602a908201527f455243313135353a20696e73756666696369656e742062616c616e636520666f60408201526939103a3930b739b332b960b11b606082015260800190565b60006001600160401b0382111561195257611952611a0c565b5060209081020190565b6000821982111561196f5761196f6119f6565b500190565b60028104600182168061198857607f821691505b602082108114156119a957634e487b7160e01b600052602260045260246000fd5b50919050565b601f8201601f191681016001600160401b03811182821017156119d4576119d4611a0c565b6040525050565b60006000198214156119ef576119ef6119f6565b5060010190565b634e487b7160e01b600052601160045260246000fd5b634e487b7160e01b600052604160045260246000fd5b600060033d1115610a2757600481823e5160e01c90565b600060443d1015611a4957610a27565b6040516003193d81016004833e81513d6001600160401b038160248401118184111715611a7a575050505050610a27565b8285019150815181811115611a9457505050505050610a27565b843d8701016020828501011115611ab057505050505050610a27565b611abf602082860101876119af565b509094505050505090565b6001600160e01b0319811681146107d457600080fdfea264697066735822122003ab2a5c30cfd29aa0653e6adbc54ab1407d71c92cc7c4ade5241036739e73d864736f6c63430008020033"

// DeployCore deploys a new Ethereum contract, binding an instance of Core to it.
func DeployCore(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *Core, error) {
	parsed, err := abi.JSON(strings.NewReader(CoreABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}

	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(CoreBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Core{CoreCaller: CoreCaller{contract: contract}, CoreTransactor: CoreTransactor{contract: contract}, CoreFilterer: CoreFilterer{contract: contract}}, nil
}

// Core is an auto generated Go binding around an Ethereum contract.
type Core struct {
	CoreCaller     // Read-only binding to the contract
	CoreTransactor // Write-only binding to the contract
	CoreFilterer   // Log filterer for contract events
}

// CoreCaller is an auto generated read-only Go binding around an Ethereum contract.
type CoreCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// CoreTransactor is an auto generated write-only Go binding around an Ethereum contract.
type CoreTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// CoreFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type CoreFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// CoreSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type CoreSession struct {
	Contract     *Core             // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// CoreCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type CoreCallerSession struct {
	Contract *CoreCaller   // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// CoreTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type CoreTransactorSession struct {
	Contract     *CoreTransactor   // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// CoreRaw is an auto generated low-level Go binding around an Ethereum contract.
type CoreRaw struct {
	Contract *Core // Generic contract binding to access the raw methods on
}

// CoreCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type CoreCallerRaw struct {
	Contract *CoreCaller // Generic read-only contract binding to access the raw methods on
}

// CoreTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type CoreTransactorRaw struct {
	Contract *CoreTransactor // Generic write-only contract binding to access the raw methods on
}

// NewCore creates a new instance of Core, bound to a specific deployed contract.
func NewCore(address common.Address, backend bind.ContractBackend) (*Core, error) {
	contract, err := bindCore(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Core{CoreCaller: CoreCaller{contract: contract}, CoreTransactor: CoreTransactor{contract: contract}, CoreFilterer: CoreFilterer{contract: contract}}, nil
}

// NewCoreCaller creates a new read-only instance of Core, bound to a specific deployed contract.
func NewCoreCaller(address common.Address, caller bind.ContractCaller) (*CoreCaller, error) {
	contract, err := bindCore(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &CoreCaller{contract: contract}, nil
}

// NewCoreTransactor creates a new write-only instance of Core, bound to a specific deployed contract.
func NewCoreTransactor(address common.Address, transactor bind.ContractTransactor) (*CoreTransactor, error) {
	contract, err := bindCore(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &CoreTransactor{contract: contract}, nil
}

// NewCoreFilterer creates a new log filterer instance of Core, bound to a specific deployed contract.
func NewCoreFilterer(address common.Address, filterer bind.ContractFilterer) (*CoreFilterer, error) {
	contract, err := bindCore(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &CoreFilterer{contract: contract}, nil
}

// bindCore binds a generic wrapper to an already deployed contract.
func bindCore(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(CoreABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Core *CoreRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Core.Contract.CoreCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Core *CoreRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Core.Contract.CoreTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Core *CoreRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Core.Contract.CoreTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Core *CoreCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Core.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Core *CoreTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Core.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Core *CoreTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Core.Contract.contract.Transact(opts, method, params...)
}

// BalanceOf is a free data retrieval call binding the contract method 0x00fdd58e.
//
// Solidity: function balanceOf(address account, uint256 id) view returns(uint256)
func (_Core *CoreCaller) BalanceOf(opts *bind.CallOpts, account common.Address, id *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _Core.contract.Call(opts, &out, "balanceOf", account, id)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BalanceOf is a free data retrieval call binding the contract method 0x00fdd58e.
//
// Solidity: function balanceOf(address account, uint256 id) view returns(uint256)
func (_Core *CoreSession) BalanceOf(account common.Address, id *big.Int) (*big.Int, error) {
	return _Core.Contract.BalanceOf(&_Core.CallOpts, account, id)
}

// BalanceOf is a free data retrieval call binding the contract method 0x00fdd58e.
//
// Solidity: function balanceOf(address account, uint256 id) view returns(uint256)
func (_Core *CoreCallerSession) BalanceOf(account common.Address, id *big.Int) (*big.Int, error) {
	return _Core.Contract.BalanceOf(&_Core.CallOpts, account, id)
}

// BalanceOfBatch is a free data retrieval call binding the contract method 0x4e1273f4.
//
// Solidity: function balanceOfBatch(address[] accounts, uint256[] ids) view returns(uint256[])
func (_Core *CoreCaller) BalanceOfBatch(opts *bind.CallOpts, accounts []common.Address, ids []*big.Int) ([]*big.Int, error) {
	var out []interface{}
	err := _Core.contract.Call(opts, &out, "balanceOfBatch", accounts, ids)

	if err != nil {
		return *new([]*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new([]*big.Int)).(*[]*big.Int)

	return out0, err

}

// BalanceOfBatch is a free data retrieval call binding the contract method 0x4e1273f4.
//
// Solidity: function balanceOfBatch(address[] accounts, uint256[] ids) view returns(uint256[])
func (_Core *CoreSession) BalanceOfBatch(accounts []common.Address, ids []*big.Int) ([]*big.Int, error) {
	return _Core.Contract.BalanceOfBatch(&_Core.CallOpts, accounts, ids)
}

// BalanceOfBatch is a free data retrieval call binding the contract method 0x4e1273f4.
//
// Solidity: function balanceOfBatch(address[] accounts, uint256[] ids) view returns(uint256[])
func (_Core *CoreCallerSession) BalanceOfBatch(accounts []common.Address, ids []*big.Int) ([]*big.Int, error) {
	return _Core.Contract.BalanceOfBatch(&_Core.CallOpts, accounts, ids)
}

// Creator is a free data retrieval call binding the contract method 0x510b5158.
//
// Solidity: function creator(uint256 _tokenId) view returns(address)
func (_Core *CoreCaller) Creator(opts *bind.CallOpts, _tokenId *big.Int) (common.Address, error) {
	var out []interface{}
	err := _Core.contract.Call(opts, &out, "creator", _tokenId)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Creator is a free data retrieval call binding the contract method 0x510b5158.
//
// Solidity: function creator(uint256 _tokenId) view returns(address)
func (_Core *CoreSession) Creator(_tokenId *big.Int) (common.Address, error) {
	return _Core.Contract.Creator(&_Core.CallOpts, _tokenId)
}

// Creator is a free data retrieval call binding the contract method 0x510b5158.
//
// Solidity: function creator(uint256 _tokenId) view returns(address)
func (_Core *CoreCallerSession) Creator(_tokenId *big.Int) (common.Address, error) {
	return _Core.Contract.Creator(&_Core.CallOpts, _tokenId)
}

// Exists is a free data retrieval call binding the contract method 0x4f558e79.
//
// Solidity: function exists(uint256 id) view returns(bool)
func (_Core *CoreCaller) Exists(opts *bind.CallOpts, id *big.Int) (bool, error) {
	var out []interface{}
	err := _Core.contract.Call(opts, &out, "exists", id)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// Exists is a free data retrieval call binding the contract method 0x4f558e79.
//
// Solidity: function exists(uint256 id) view returns(bool)
func (_Core *CoreSession) Exists(id *big.Int) (bool, error) {
	return _Core.Contract.Exists(&_Core.CallOpts, id)
}

// Exists is a free data retrieval call binding the contract method 0x4f558e79.
//
// Solidity: function exists(uint256 id) view returns(bool)
func (_Core *CoreCallerSession) Exists(id *big.Int) (bool, error) {
	return _Core.Contract.Exists(&_Core.CallOpts, id)
}

// IsApprovedForAll is a free data retrieval call binding the contract method 0xe985e9c5.
//
// Solidity: function isApprovedForAll(address account, address operator) view returns(bool)
func (_Core *CoreCaller) IsApprovedForAll(opts *bind.CallOpts, account common.Address, operator common.Address) (bool, error) {
	var out []interface{}
	err := _Core.contract.Call(opts, &out, "isApprovedForAll", account, operator)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsApprovedForAll is a free data retrieval call binding the contract method 0xe985e9c5.
//
// Solidity: function isApprovedForAll(address account, address operator) view returns(bool)
func (_Core *CoreSession) IsApprovedForAll(account common.Address, operator common.Address) (bool, error) {
	return _Core.Contract.IsApprovedForAll(&_Core.CallOpts, account, operator)
}

// IsApprovedForAll is a free data retrieval call binding the contract method 0xe985e9c5.
//
// Solidity: function isApprovedForAll(address account, address operator) view returns(bool)
func (_Core *CoreCallerSession) IsApprovedForAll(account common.Address, operator common.Address) (bool, error) {
	return _Core.Contract.IsApprovedForAll(&_Core.CallOpts, account, operator)
}

// RoyaltyFee is a free data retrieval call binding the contract method 0xc57dc235.
//
// Solidity: function royaltyFee(uint256 _tokenId) view returns(uint16)
func (_Core *CoreCaller) RoyaltyFee(opts *bind.CallOpts, _tokenId *big.Int) (uint16, error) {
	var out []interface{}
	err := _Core.contract.Call(opts, &out, "royaltyFee", _tokenId)

	if err != nil {
		return *new(uint16), err
	}

	out0 := *abi.ConvertType(out[0], new(uint16)).(*uint16)

	return out0, err

}

// RoyaltyFee is a free data retrieval call binding the contract method 0xc57dc235.
//
// Solidity: function royaltyFee(uint256 _tokenId) view returns(uint16)
func (_Core *CoreSession) RoyaltyFee(_tokenId *big.Int) (uint16, error) {
	return _Core.Contract.RoyaltyFee(&_Core.CallOpts, _tokenId)
}

// RoyaltyFee is a free data retrieval call binding the contract method 0xc57dc235.
//
// Solidity: function royaltyFee(uint256 _tokenId) view returns(uint16)
func (_Core *CoreCallerSession) RoyaltyFee(_tokenId *big.Int) (uint16, error) {
	return _Core.Contract.RoyaltyFee(&_Core.CallOpts, _tokenId)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_Core *CoreCaller) SupportsInterface(opts *bind.CallOpts, interfaceId [4]byte) (bool, error) {
	var out []interface{}
	err := _Core.contract.Call(opts, &out, "supportsInterface", interfaceId)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_Core *CoreSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _Core.Contract.SupportsInterface(&_Core.CallOpts, interfaceId)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_Core *CoreCallerSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _Core.Contract.SupportsInterface(&_Core.CallOpts, interfaceId)
}

// TotalSupply is a free data retrieval call binding the contract method 0xbd85b039.
//
// Solidity: function totalSupply(uint256 id) view returns(uint256)
func (_Core *CoreCaller) TotalSupply(opts *bind.CallOpts, id *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _Core.contract.Call(opts, &out, "totalSupply", id)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalSupply is a free data retrieval call binding the contract method 0xbd85b039.
//
// Solidity: function totalSupply(uint256 id) view returns(uint256)
func (_Core *CoreSession) TotalSupply(id *big.Int) (*big.Int, error) {
	return _Core.Contract.TotalSupply(&_Core.CallOpts, id)
}

// TotalSupply is a free data retrieval call binding the contract method 0xbd85b039.
//
// Solidity: function totalSupply(uint256 id) view returns(uint256)
func (_Core *CoreCallerSession) TotalSupply(id *big.Int) (*big.Int, error) {
	return _Core.Contract.TotalSupply(&_Core.CallOpts, id)
}

// Uri is a free data retrieval call binding the contract method 0x0e89341c.
//
// Solidity: function uri(uint256 ) view returns(string)
func (_Core *CoreCaller) Uri(opts *bind.CallOpts, arg0 *big.Int) (string, error) {
	var out []interface{}
	err := _Core.contract.Call(opts, &out, "uri", arg0)

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Uri is a free data retrieval call binding the contract method 0x0e89341c.
//
// Solidity: function uri(uint256 ) view returns(string)
func (_Core *CoreSession) Uri(arg0 *big.Int) (string, error) {
	return _Core.Contract.Uri(&_Core.CallOpts, arg0)
}

// Uri is a free data retrieval call binding the contract method 0x0e89341c.
//
// Solidity: function uri(uint256 ) view returns(string)
func (_Core *CoreCallerSession) Uri(arg0 *big.Int) (string, error) {
	return _Core.Contract.Uri(&_Core.CallOpts, arg0)
}

// Create is a paid mutator transaction binding the contract method 0x7d36030f.
//
// Solidity: function create(bytes id, uint64 amount, uint16 __royaltyFee) returns(uint256)
func (_Core *CoreTransactor) Create(opts *bind.TransactOpts, id []byte, amount uint64, __royaltyFee uint16) (*types.Transaction, error) {
	return _Core.contract.Transact(opts, "create", id, amount, __royaltyFee)
}

// Create is a paid mutator transaction binding the contract method 0x7d36030f.
//
// Solidity: function create(bytes id, uint64 amount, uint16 __royaltyFee) returns(uint256)
func (_Core *CoreSession) Create(id []byte, amount uint64, __royaltyFee uint16) (*types.Transaction, error) {
	return _Core.Contract.Create(&_Core.TransactOpts, id, amount, __royaltyFee)
}

// Create is a paid mutator transaction binding the contract method 0x7d36030f.
//
// Solidity: function create(bytes id, uint64 amount, uint16 __royaltyFee) returns(uint256)
func (_Core *CoreTransactorSession) Create(id []byte, amount uint64, __royaltyFee uint16) (*types.Transaction, error) {
	return _Core.Contract.Create(&_Core.TransactOpts, id, amount, __royaltyFee)
}

// CreateFromContract is a paid mutator transaction binding the contract method 0xc440ef5e.
//
// Solidity: function createFromContract(bytes id, uint64 amount, uint16 __royaltyFee, address __creator) returns(uint256)
func (_Core *CoreTransactor) CreateFromContract(opts *bind.TransactOpts, id []byte, amount uint64, __royaltyFee uint16, __creator common.Address) (*types.Transaction, error) {
	return _Core.contract.Transact(opts, "createFromContract", id, amount, __royaltyFee, __creator)
}

// CreateFromContract is a paid mutator transaction binding the contract method 0xc440ef5e.
//
// Solidity: function createFromContract(bytes id, uint64 amount, uint16 __royaltyFee, address __creator) returns(uint256)
func (_Core *CoreSession) CreateFromContract(id []byte, amount uint64, __royaltyFee uint16, __creator common.Address) (*types.Transaction, error) {
	return _Core.Contract.CreateFromContract(&_Core.TransactOpts, id, amount, __royaltyFee, __creator)
}

// CreateFromContract is a paid mutator transaction binding the contract method 0xc440ef5e.
//
// Solidity: function createFromContract(bytes id, uint64 amount, uint16 __royaltyFee, address __creator) returns(uint256)
func (_Core *CoreTransactorSession) CreateFromContract(id []byte, amount uint64, __royaltyFee uint16, __creator common.Address) (*types.Transaction, error) {
	return _Core.Contract.CreateFromContract(&_Core.TransactOpts, id, amount, __royaltyFee, __creator)
}

// Initialize is a paid mutator transaction binding the contract method 0x8129fc1c.
//
// Solidity: function initialize() returns()
func (_Core *CoreTransactor) Initialize(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Core.contract.Transact(opts, "initialize")
}

// Initialize is a paid mutator transaction binding the contract method 0x8129fc1c.
//
// Solidity: function initialize() returns()
func (_Core *CoreSession) Initialize() (*types.Transaction, error) {
	return _Core.Contract.Initialize(&_Core.TransactOpts)
}

// Initialize is a paid mutator transaction binding the contract method 0x8129fc1c.
//
// Solidity: function initialize() returns()
func (_Core *CoreTransactorSession) Initialize() (*types.Transaction, error) {
	return _Core.Contract.Initialize(&_Core.TransactOpts)
}

// SafeBatchTransferFrom is a paid mutator transaction binding the contract method 0x2eb2c2d6.
//
// Solidity: function safeBatchTransferFrom(address from, address to, uint256[] ids, uint256[] amounts, bytes data) returns()
func (_Core *CoreTransactor) SafeBatchTransferFrom(opts *bind.TransactOpts, from common.Address, to common.Address, ids []*big.Int, amounts []*big.Int, data []byte) (*types.Transaction, error) {
	return _Core.contract.Transact(opts, "safeBatchTransferFrom", from, to, ids, amounts, data)
}

// SafeBatchTransferFrom is a paid mutator transaction binding the contract method 0x2eb2c2d6.
//
// Solidity: function safeBatchTransferFrom(address from, address to, uint256[] ids, uint256[] amounts, bytes data) returns()
func (_Core *CoreSession) SafeBatchTransferFrom(from common.Address, to common.Address, ids []*big.Int, amounts []*big.Int, data []byte) (*types.Transaction, error) {
	return _Core.Contract.SafeBatchTransferFrom(&_Core.TransactOpts, from, to, ids, amounts, data)
}

// SafeBatchTransferFrom is a paid mutator transaction binding the contract method 0x2eb2c2d6.
//
// Solidity: function safeBatchTransferFrom(address from, address to, uint256[] ids, uint256[] amounts, bytes data) returns()
func (_Core *CoreTransactorSession) SafeBatchTransferFrom(from common.Address, to common.Address, ids []*big.Int, amounts []*big.Int, data []byte) (*types.Transaction, error) {
	return _Core.Contract.SafeBatchTransferFrom(&_Core.TransactOpts, from, to, ids, amounts, data)
}

// SafeTransferFrom is a paid mutator transaction binding the contract method 0xf242432a.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 id, uint256 amount, bytes data) returns()
func (_Core *CoreTransactor) SafeTransferFrom(opts *bind.TransactOpts, from common.Address, to common.Address, id *big.Int, amount *big.Int, data []byte) (*types.Transaction, error) {
	return _Core.contract.Transact(opts, "safeTransferFrom", from, to, id, amount, data)
}

// SafeTransferFrom is a paid mutator transaction binding the contract method 0xf242432a.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 id, uint256 amount, bytes data) returns()
func (_Core *CoreSession) SafeTransferFrom(from common.Address, to common.Address, id *big.Int, amount *big.Int, data []byte) (*types.Transaction, error) {
	return _Core.Contract.SafeTransferFrom(&_Core.TransactOpts, from, to, id, amount, data)
}

// SafeTransferFrom is a paid mutator transaction binding the contract method 0xf242432a.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 id, uint256 amount, bytes data) returns()
func (_Core *CoreTransactorSession) SafeTransferFrom(from common.Address, to common.Address, id *big.Int, amount *big.Int, data []byte) (*types.Transaction, error) {
	return _Core.Contract.SafeTransferFrom(&_Core.TransactOpts, from, to, id, amount, data)
}

// SetApprovalForAll is a paid mutator transaction binding the contract method 0xa22cb465.
//
// Solidity: function setApprovalForAll(address operator, bool approved) returns()
func (_Core *CoreTransactor) SetApprovalForAll(opts *bind.TransactOpts, operator common.Address, approved bool) (*types.Transaction, error) {
	return _Core.contract.Transact(opts, "setApprovalForAll", operator, approved)
}

// SetApprovalForAll is a paid mutator transaction binding the contract method 0xa22cb465.
//
// Solidity: function setApprovalForAll(address operator, bool approved) returns()
func (_Core *CoreSession) SetApprovalForAll(operator common.Address, approved bool) (*types.Transaction, error) {
	return _Core.Contract.SetApprovalForAll(&_Core.TransactOpts, operator, approved)
}

// SetApprovalForAll is a paid mutator transaction binding the contract method 0xa22cb465.
//
// Solidity: function setApprovalForAll(address operator, bool approved) returns()
func (_Core *CoreTransactorSession) SetApprovalForAll(operator common.Address, approved bool) (*types.Transaction, error) {
	return _Core.Contract.SetApprovalForAll(&_Core.TransactOpts, operator, approved)
}

// CoreApprovalForAllIterator is returned from FilterApprovalForAll and is used to iterate over the raw logs and unpacked data for ApprovalForAll events raised by the Core contract.
type CoreApprovalForAllIterator struct {
	Event *CoreApprovalForAll // Event containing the contract specifics and raw log

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
func (it *CoreApprovalForAllIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CoreApprovalForAll)
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
		it.Event = new(CoreApprovalForAll)
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
func (it *CoreApprovalForAllIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CoreApprovalForAllIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CoreApprovalForAll represents a ApprovalForAll event raised by the Core contract.
type CoreApprovalForAll struct {
	Account  common.Address
	Operator common.Address
	Approved bool
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterApprovalForAll is a free log retrieval operation binding the contract event 0x17307eab39ab6107e8899845ad3d59bd9653f200f220920489ca2b5937696c31.
//
// Solidity: event ApprovalForAll(address indexed account, address indexed operator, bool approved)
func (_Core *CoreFilterer) FilterApprovalForAll(opts *bind.FilterOpts, account []common.Address, operator []common.Address) (*CoreApprovalForAllIterator, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}
	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}

	logs, sub, err := _Core.contract.FilterLogs(opts, "ApprovalForAll", accountRule, operatorRule)
	if err != nil {
		return nil, err
	}
	return &CoreApprovalForAllIterator{contract: _Core.contract, event: "ApprovalForAll", logs: logs, sub: sub}, nil
}

// WatchApprovalForAll is a free log subscription operation binding the contract event 0x17307eab39ab6107e8899845ad3d59bd9653f200f220920489ca2b5937696c31.
//
// Solidity: event ApprovalForAll(address indexed account, address indexed operator, bool approved)
func (_Core *CoreFilterer) WatchApprovalForAll(opts *bind.WatchOpts, sink chan<- *CoreApprovalForAll, account []common.Address, operator []common.Address) (event.Subscription, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}
	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}

	logs, sub, err := _Core.contract.WatchLogs(opts, "ApprovalForAll", accountRule, operatorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CoreApprovalForAll)
				if err := _Core.contract.UnpackLog(event, "ApprovalForAll", log); err != nil {
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

// ParseApprovalForAll is a log parse operation binding the contract event 0x17307eab39ab6107e8899845ad3d59bd9653f200f220920489ca2b5937696c31.
//
// Solidity: event ApprovalForAll(address indexed account, address indexed operator, bool approved)
func (_Core *CoreFilterer) ParseApprovalForAll(log types.Log) (*CoreApprovalForAll, error) {
	event := new(CoreApprovalForAll)
	if err := _Core.contract.UnpackLog(event, "ApprovalForAll", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CoreTransferBatchIterator is returned from FilterTransferBatch and is used to iterate over the raw logs and unpacked data for TransferBatch events raised by the Core contract.
type CoreTransferBatchIterator struct {
	Event *CoreTransferBatch // Event containing the contract specifics and raw log

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
func (it *CoreTransferBatchIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CoreTransferBatch)
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
		it.Event = new(CoreTransferBatch)
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
func (it *CoreTransferBatchIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CoreTransferBatchIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CoreTransferBatch represents a TransferBatch event raised by the Core contract.
type CoreTransferBatch struct {
	Operator common.Address
	From     common.Address
	To       common.Address
	Ids      []*big.Int
	Values   []*big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterTransferBatch is a free log retrieval operation binding the contract event 0x4a39dc06d4c0dbc64b70af90fd698a233a518aa5d07e595d983b8c0526c8f7fb.
//
// Solidity: event TransferBatch(address indexed operator, address indexed from, address indexed to, uint256[] ids, uint256[] values)
func (_Core *CoreFilterer) FilterTransferBatch(opts *bind.FilterOpts, operator []common.Address, from []common.Address, to []common.Address) (*CoreTransferBatchIterator, error) {

	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}
	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _Core.contract.FilterLogs(opts, "TransferBatch", operatorRule, fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &CoreTransferBatchIterator{contract: _Core.contract, event: "TransferBatch", logs: logs, sub: sub}, nil
}

// WatchTransferBatch is a free log subscription operation binding the contract event 0x4a39dc06d4c0dbc64b70af90fd698a233a518aa5d07e595d983b8c0526c8f7fb.
//
// Solidity: event TransferBatch(address indexed operator, address indexed from, address indexed to, uint256[] ids, uint256[] values)
func (_Core *CoreFilterer) WatchTransferBatch(opts *bind.WatchOpts, sink chan<- *CoreTransferBatch, operator []common.Address, from []common.Address, to []common.Address) (event.Subscription, error) {

	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}
	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _Core.contract.WatchLogs(opts, "TransferBatch", operatorRule, fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CoreTransferBatch)
				if err := _Core.contract.UnpackLog(event, "TransferBatch", log); err != nil {
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

// ParseTransferBatch is a log parse operation binding the contract event 0x4a39dc06d4c0dbc64b70af90fd698a233a518aa5d07e595d983b8c0526c8f7fb.
//
// Solidity: event TransferBatch(address indexed operator, address indexed from, address indexed to, uint256[] ids, uint256[] values)
func (_Core *CoreFilterer) ParseTransferBatch(log types.Log) (*CoreTransferBatch, error) {
	event := new(CoreTransferBatch)
	if err := _Core.contract.UnpackLog(event, "TransferBatch", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CoreTransferSingleIterator is returned from FilterTransferSingle and is used to iterate over the raw logs and unpacked data for TransferSingle events raised by the Core contract.
type CoreTransferSingleIterator struct {
	Event *CoreTransferSingle // Event containing the contract specifics and raw log

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
func (it *CoreTransferSingleIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CoreTransferSingle)
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
		it.Event = new(CoreTransferSingle)
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
func (it *CoreTransferSingleIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CoreTransferSingleIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CoreTransferSingle represents a TransferSingle event raised by the Core contract.
type CoreTransferSingle struct {
	Operator common.Address
	From     common.Address
	To       common.Address
	Id       *big.Int
	Value    *big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterTransferSingle is a free log retrieval operation binding the contract event 0xc3d58168c5ae7397731d063d5bbf3d657854427343f4c083240f7aacaa2d0f62.
//
// Solidity: event TransferSingle(address indexed operator, address indexed from, address indexed to, uint256 id, uint256 value)
func (_Core *CoreFilterer) FilterTransferSingle(opts *bind.FilterOpts, operator []common.Address, from []common.Address, to []common.Address) (*CoreTransferSingleIterator, error) {

	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}
	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _Core.contract.FilterLogs(opts, "TransferSingle", operatorRule, fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &CoreTransferSingleIterator{contract: _Core.contract, event: "TransferSingle", logs: logs, sub: sub}, nil
}

// WatchTransferSingle is a free log subscription operation binding the contract event 0xc3d58168c5ae7397731d063d5bbf3d657854427343f4c083240f7aacaa2d0f62.
//
// Solidity: event TransferSingle(address indexed operator, address indexed from, address indexed to, uint256 id, uint256 value)
func (_Core *CoreFilterer) WatchTransferSingle(opts *bind.WatchOpts, sink chan<- *CoreTransferSingle, operator []common.Address, from []common.Address, to []common.Address) (event.Subscription, error) {

	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}
	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _Core.contract.WatchLogs(opts, "TransferSingle", operatorRule, fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CoreTransferSingle)
				if err := _Core.contract.UnpackLog(event, "TransferSingle", log); err != nil {
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

// ParseTransferSingle is a log parse operation binding the contract event 0xc3d58168c5ae7397731d063d5bbf3d657854427343f4c083240f7aacaa2d0f62.
//
// Solidity: event TransferSingle(address indexed operator, address indexed from, address indexed to, uint256 id, uint256 value)
func (_Core *CoreFilterer) ParseTransferSingle(log types.Log) (*CoreTransferSingle, error) {
	event := new(CoreTransferSingle)
	if err := _Core.contract.UnpackLog(event, "TransferSingle", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CoreURIIterator is returned from FilterURI and is used to iterate over the raw logs and unpacked data for URI events raised by the Core contract.
type CoreURIIterator struct {
	Event *CoreURI // Event containing the contract specifics and raw log

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
func (it *CoreURIIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CoreURI)
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
		it.Event = new(CoreURI)
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
func (it *CoreURIIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CoreURIIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CoreURI represents a URI event raised by the Core contract.
type CoreURI struct {
	Value string
	Id    *big.Int
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterURI is a free log retrieval operation binding the contract event 0x6bb7ff708619ba0610cba295a58592e0451dee2622938c8755667688daf3529b.
//
// Solidity: event URI(string value, uint256 indexed id)
func (_Core *CoreFilterer) FilterURI(opts *bind.FilterOpts, id []*big.Int) (*CoreURIIterator, error) {

	var idRule []interface{}
	for _, idItem := range id {
		idRule = append(idRule, idItem)
	}

	logs, sub, err := _Core.contract.FilterLogs(opts, "URI", idRule)
	if err != nil {
		return nil, err
	}
	return &CoreURIIterator{contract: _Core.contract, event: "URI", logs: logs, sub: sub}, nil
}

// WatchURI is a free log subscription operation binding the contract event 0x6bb7ff708619ba0610cba295a58592e0451dee2622938c8755667688daf3529b.
//
// Solidity: event URI(string value, uint256 indexed id)
func (_Core *CoreFilterer) WatchURI(opts *bind.WatchOpts, sink chan<- *CoreURI, id []*big.Int) (event.Subscription, error) {

	var idRule []interface{}
	for _, idItem := range id {
		idRule = append(idRule, idItem)
	}

	logs, sub, err := _Core.contract.WatchLogs(opts, "URI", idRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CoreURI)
				if err := _Core.contract.UnpackLog(event, "URI", log); err != nil {
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

// ParseURI is a log parse operation binding the contract event 0x6bb7ff708619ba0610cba295a58592e0451dee2622938c8755667688daf3529b.
//
// Solidity: event URI(string value, uint256 indexed id)
func (_Core *CoreFilterer) ParseURI(log types.Log) (*CoreURI, error) {
	event := new(CoreURI)
	if err := _Core.contract.UnpackLog(event, "URI", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
