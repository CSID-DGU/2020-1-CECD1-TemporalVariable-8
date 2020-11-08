// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package migc

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

// MigcABI is the input ABI used to generate the binding from.
const MigcABI = "[{\"inputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"string\",\"name\":\"url\",\"type\":\"string\"},{\"indexed\":true,\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"}],\"name\":\"BusinessCreated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"fromOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"toOwner\",\"type\":\"address\"}],\"name\":\"BusinessTransfer\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"string\",\"name\":\"url\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"expired\",\"type\":\"uint256\"}],\"name\":\"BusinessUpdate\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"previousAdminRole\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"newAdminRole\",\"type\":\"bytes32\"}],\"name\":\"RoleAdminChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"RoleGranted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"RoleRevoked\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"DEFAULT_ADMIN_ROLE\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"url\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"}],\"name\":\"createBusiness\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"name\":\"datamap\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"businessOwner\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"url\",\"type\":\"string\"},{\"internalType\":\"uint64\",\"name\":\"expired\",\"type\":\"uint64\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"}],\"name\":\"getRoleAdmin\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"}],\"name\":\"getRoleMember\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"}],\"name\":\"getRoleMemberCount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"grantRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"hasRole\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"listURLCount\",\"outputs\":[{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"renounceRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"revokeRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"url\",\"type\":\"string\"},{\"internalType\":\"address\",\"name\":\"toOwner\",\"type\":\"address\"}],\"name\":\"transferBusiness\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"urls\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"url\",\"type\":\"string\"}],\"name\":\"weeklyUpdate\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"

// MigcBin is the compiled bytecode used for deploying new contracts.
var MigcBin = "0x60806040523480156200001157600080fd5b5060006200001e62000088565b600180546001600160a01b0319166001600160a01b038316908117909155604051919250906000907f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0908290a3506200008260006200007c6200008c565b6200009b565b620001ab565b3390565b6001546001600160a01b031690565b620000a78282620000ab565b5050565b600082815260208181526040909120620000d09183906200136762000124821b17901c565b15620000a757620000e062000088565b6001600160a01b0316816001600160a01b0316837f2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d60405160405180910390a45050565b60006200013b836001600160a01b03841662000144565b90505b92915050565b600062000152838362000193565b6200018a575081546001818101845560008481526020808220909301849055845484825282860190935260409020919091556200013e565b5060006200013e565b60009081526001919091016020526040902054151590565b61181c80620001bb6000396000f3fe608060405234801561001057600080fd5b506004361061010b5760003560e01c80638da5cb5b116100a2578063c609d81611610071578063c609d8161461063d578063ca15c873146106e1578063d547741f146106fe578063f2fde38b1461072a578063faffb97c146107505761010b565b80638da5cb5b146105ae5780639010d07c146105d257806391d14854146105f5578063a217fddf146106355761010b565b806360d41a26116100de57806360d41a26146102c25780636930941214610371578063715018a614610514578063796676be1461051c5761010b565b8063248a9ca3146101105780632f2ff15d1461013f57806333d104721461016d57806336568abe14610296575b600080fd5b61012d6004803603602081101561012657600080fd5b5035610775565b60408051918252519081900360200190f35b61016b6004803603604081101561015557600080fd5b50803590602001356001600160a01b031661078a565b005b61016b6004803603604081101561018357600080fd5b810190602081018135600160201b81111561019d57600080fd5b8201836020820111156101af57600080fd5b803590602001918460018302840111600160201b831117156101d057600080fd5b91908080601f0160208091040260200160405190810160405280939291908181526020018383808284376000920191909152509295949360208101935035915050600160201b81111561022257600080fd5b82018360208201111561023457600080fd5b803590602001918460018302840111600160201b8311171561025557600080fd5b91908080601f0160208091040260200160405190810160405280939291908181526020018383808284376000920191909152509295506107f6945050505050565b61016b600480360360408110156102ac57600080fd5b50803590602001356001600160a01b03166109ff565b61016b600480360360408110156102d857600080fd5b810190602081018135600160201b8111156102f257600080fd5b82018360208201111561030457600080fd5b803590602001918460018302840111600160201b8311171561032557600080fd5b91908080601f016020809104026020016040519081016040528093929190818152602001838380828437600092019190915250929550505090356001600160a01b03169150610a609050565b6104156004803603602081101561038757600080fd5b810190602081018135600160201b8111156103a157600080fd5b8201836020820111156103b357600080fd5b803590602001918460018302840111600160201b831117156103d457600080fd5b91908080601f016020809104026020016040519081016040528093929190818152602001838380828437600092019190915250929550610ce7945050505050565b60405180856001600160a01b0316815260200180602001806020018467ffffffffffffffff168152602001838103835286818151815260200191508051906020019080838360005b8381101561047557818101518382015260200161045d565b50505050905090810190601f1680156104a25780820380516001836020036101000a031916815260200191505b50838103825285518152855160209182019187019080838360005b838110156104d55781810151838201526020016104bd565b50505050905090810190601f1680156105025780820380516001836020036101000a031916815260200191505b50965050505050505060405180910390f35b61016b610e46565b6105396004803603602081101561053257600080fd5b5035610efa565b6040805160208082528351818301528351919283929083019185019080838360005b8381101561057357818101518382015260200161055b565b50505050905090810190601f1680156105a05780820380516001836020036101000a031916815260200191505b509250505060405180910390f35b6105b6610fa3565b604080516001600160a01b039092168252519081900360200190f35b6105b6600480360360408110156105e857600080fd5b5080359060200135610fb2565b6106216004803603604081101561060b57600080fd5b50803590602001356001600160a01b0316610fd3565b604080519115158252519081900360200190f35b61012d610feb565b61016b6004803603602081101561065357600080fd5b810190602081018135600160201b81111561066d57600080fd5b82018360208201111561067f57600080fd5b803590602001918460018302840111600160201b831117156106a057600080fd5b91908080601f016020809104026020016040519081016040528093929190818152602001838380828437600092019190915250929550610ff0945050505050565b61012d600480360360208110156106f757600080fd5b50356111e6565b61016b6004803603604081101561071457600080fd5b50803590602001356001600160a01b03166111fd565b61016b6004803603602081101561074057600080fd5b50356001600160a01b0316611256565b610758611361565b6040805167ffffffffffffffff9092168252519081900360200190f35b60009081526020819052604090206002015490565b6000828152602081905260409020600201546107ad906107a861137c565b610fd3565b6107e85760405162461bcd60e51b815260040180806020018281038252602f8152602001806116e7602f913960400191505060405180910390fd5b6107f28282611380565b5050565b6002826040518082805190602001908083835b602083106108285780518252601f199092019160209182019101610809565b51815160209384036101000a600019018019909216911617905292019485525060405193849003019092206003015467ffffffffffffffff161591506108a99050576040805162461bcd60e51b815260206004820152600f60248201526e1d5c9b081a5cc8185cdcda59db9959608a1b604482015290519081900360640190fd5b6040518060800160405280336001600160a01b031681526020018281526020018381526020014267ffffffffffffffff168152506002836040518082805190602001908083835b6020831061090f5780518252601f1990920191602091820191016108f0565b51815160209384036101000a6000190180199092169116179052920194855250604051938490038101909320845181546001600160a01b0319166001600160a01b03909116178155848401518051919461097194506001860193500190611623565b506040820151805161098d916002840191602090910190611623565b50606091909101516003918201805467ffffffffffffffff191667ffffffffffffffff909216919091179055805460018101825560009190915282516109fa917fc2575a0e9e593c00f959f8c92f12db2869c3395a3b0502d05e2516446f71f85b01906020850190611623565b505050565b610a0761137c565b6001600160a01b0316816001600160a01b031614610a565760405162461bcd60e51b815260040180806020018281038252602f8152602001806117b8602f913960400191505060405180910390fd5b6107f282826113e9565b6002826040518082805190602001908083835b60208310610a925780518252601f199092019160209182019101610a73565b51815160209384036101000a600019018019909216911617905292019485525060405193849003019092206003015467ffffffffffffffff1615159150610b169050576040805162461bcd60e51b81526020600482015260116024820152701d5c9b081a5cc81d5b985cdcda59db9959607a1b604482015290519081900360640190fd5b336001600160a01b03166002836040518082805190602001908083835b60208310610b525780518252601f199092019160209182019101610b33565b51815160209384036101000a60001901801990921691161790529201948552506040519384900301909220546001600160a01b0316929092149150610bca90505760405162461bcd60e51b815260040180806020018281038252602781526020018061173c6027913960400191505060405180910390fd5b60006002836040518082805190602001908083835b60208310610bfe5780518252601f199092019160209182019101610bdf565b51815160209384036101000a600019018019909216911617905292019485525060405193849003018320546001600160a01b0390811694508516928492507f54bdaa308d34406b09b20101de9b6467c26e5901ee017b79c618c70cbe117cf99150600090a3816002846040518082805190602001908083835b60208310610c965780518252601f199092019160209182019101610c77565b51815160209384036101000a6000190180199092169116179052920194855250604051938490030190922080546001600160a01b0319166001600160a01b0394909416939093179092555050505050565b80516020818301810180516002808352938301948301949094209390528254600180850180546040805161010094831615949094026000190190911695909504601f81018590048502830185019095528482526001600160a01b0390921694939092830182828015610d9a5780601f10610d6f57610100808354040283529160200191610d9a565b820191906000526020600020905b815481529060010190602001808311610d7d57829003601f168201915b50505060028085018054604080516020601f6000196101006001871615020190941695909504928301859004850281018501909152818152959695945090925090830182828015610e2c5780601f10610e0157610100808354040283529160200191610e2c565b820191906000526020600020905b815481529060010190602001808311610e0f57829003601f168201915b5050506003909301549192505067ffffffffffffffff1684565b610e4e61137c565b6001546001600160a01b03908116911614610eb0576040805162461bcd60e51b815260206004820181905260248201527f4f776e61626c653a2063616c6c6572206973206e6f7420746865206f776e6572604482015290519081900360640190fd5b6001546040516000916001600160a01b0316907f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0908390a3600180546001600160a01b0319169055565b60038181548110610f0a57600080fd5b600091825260209182902001805460408051601f6002600019610100600187161502019094169390930492830185900485028101850190915281815293509091830182828015610f9b5780601f10610f7057610100808354040283529160200191610f9b565b820191906000526020600020905b815481529060010190602001808311610f7e57829003601f168201915b505050505081565b6001546001600160a01b031690565b6000828152602081905260408120610fca9083611452565b90505b92915050565b6000828152602081905260408120610fca908361145e565b600081565b6002816040518082805190602001908083835b602083106110225780518252601f199092019160209182019101611003565b51815160209384036101000a600019018019909216911617905292019485525060405193849003019092206003015467ffffffffffffffff16151591506110a69050576040805162461bcd60e51b81526020600482015260116024820152701d5c9b081a5cc81d5b985cdcda59db9959607a1b604482015290519081900360640190fd5b336001600160a01b03166002826040518082805190602001908083835b602083106110e25780518252601f1990920191602091820191016110c3565b51815160209384036101000a60001901801990921691161790529201948552506040519384900301909220546001600160a01b031692909214915061115a90505760405162461bcd60e51b81526004018080602001828103825260258152602001806117936025913960400191505060405180910390fd5b4262093a80016002826040518082805190602001908083835b602083106111925780518252601f199092019160209182019101611173565b51815160209384036101000a60001901801990921691161790529201948552506040519384900301909220600301805467ffffffffffffffff191667ffffffffffffffff9490941693909317909255505050565b6000818152602081905260408120610fcd90611473565b60008281526020819052604090206002015461121b906107a861137c565b610a565760405162461bcd60e51b81526004018080602001828103825260308152602001806117636030913960400191505060405180910390fd5b61125e61137c565b6001546001600160a01b039081169116146112c0576040805162461bcd60e51b815260206004820181905260248201527f4f776e61626c653a2063616c6c6572206973206e6f7420746865206f776e6572604482015290519081900360640190fd5b6001600160a01b0381166113055760405162461bcd60e51b81526004018080602001828103825260268152602001806117166026913960400191505060405180910390fd5b6001546040516001600160a01b038084169216907f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e090600090a3600180546001600160a01b0319166001600160a01b0392909216919091179055565b60035490565b6000610fca836001600160a01b03841661147e565b3390565b60008281526020819052604090206113989082611367565b156107f2576113a561137c565b6001600160a01b0316816001600160a01b0316837f2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d60405160405180910390a45050565b600082815260208190526040902061140190826114c8565b156107f25761140e61137c565b6001600160a01b0316816001600160a01b0316837ff6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b60405160405180910390a45050565b6000610fca83836114dd565b6000610fca836001600160a01b038416611541565b6000610fcd82611559565b600061148a8383611541565b6114c057508154600181810184556000848152602080822090930184905584548482528286019093526040902091909155610fcd565b506000610fcd565b6000610fca836001600160a01b03841661155d565b8154600090821061151f5760405162461bcd60e51b81526004018080602001828103825260228152602001806116c56022913960400191505060405180910390fd5b82600001828154811061152e57fe5b9060005260206000200154905092915050565b60009081526001919091016020526040902054151590565b5490565b60008181526001830160205260408120548015611619578354600019808301919081019060009087908390811061159057fe5b90600052602060002001549050808760000184815481106115ad57fe5b6000918252602080832090910192909255828152600189810190925260409020908401905586548790806115dd57fe5b60019003818190600052602060002001600090559055866001016000878152602001908152602001600020600090556001945050505050610fcd565b6000915050610fcd565b828054600181600116156101000203166002900490600052602060002090601f016020900481019282611659576000855561169f565b82601f1061167257805160ff191683800117855561169f565b8280016001018555821561169f579182015b8281111561169f578251825591602001919060010190611684565b506116ab9291506116af565b5090565b5b808211156116ab57600081556001016116b056fe456e756d657261626c655365743a20696e646578206f7574206f6620626f756e6473416363657373436f6e74726f6c3a2073656e646572206d75737420626520616e2061646d696e20746f206772616e744f776e61626c653a206e6577206f776e657220697320746865207a65726f20616464726573737472616e73666572206f6e6c7920616c6c6f77656420666f7220627573696e6573734f776e6572416363657373436f6e74726f6c3a2073656e646572206d75737420626520616e2061646d696e20746f207265766f6b65757064617465206f6e6c7920616c6c6f77656420666f7220627573696e6573734f776e6572416363657373436f6e74726f6c3a2063616e206f6e6c792072656e6f756e636520726f6c657320666f722073656c66a26469706673582212200ea9a604042c8ab800d4c8f0382a41f47c331669dd01ab871db74da956bec40364736f6c63430007040033"

// DeployMigc deploys a new Ethereum contract, binding an instance of Migc to it.
func DeployMigc(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *Migc, error) {
	parsed, err := abi.JSON(strings.NewReader(MigcABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}

	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(MigcBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Migc{MigcCaller: MigcCaller{contract: contract}, MigcTransactor: MigcTransactor{contract: contract}, MigcFilterer: MigcFilterer{contract: contract}}, nil
}

// Migc is an auto generated Go binding around an Ethereum contract.
type Migc struct {
	MigcCaller     // Read-only binding to the contract
	MigcTransactor // Write-only binding to the contract
	MigcFilterer   // Log filterer for contract events
}

// MigcCaller is an auto generated read-only Go binding around an Ethereum contract.
type MigcCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MigcTransactor is an auto generated write-only Go binding around an Ethereum contract.
type MigcTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MigcFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type MigcFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MigcSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type MigcSession struct {
	Contract     *Migc             // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// MigcCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type MigcCallerSession struct {
	Contract *MigcCaller   // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// MigcTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type MigcTransactorSession struct {
	Contract     *MigcTransactor   // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// MigcRaw is an auto generated low-level Go binding around an Ethereum contract.
type MigcRaw struct {
	Contract *Migc // Generic contract binding to access the raw methods on
}

// MigcCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type MigcCallerRaw struct {
	Contract *MigcCaller // Generic read-only contract binding to access the raw methods on
}

// MigcTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type MigcTransactorRaw struct {
	Contract *MigcTransactor // Generic write-only contract binding to access the raw methods on
}

// NewMigc creates a new instance of Migc, bound to a specific deployed contract.
func NewMigc(address common.Address, backend bind.ContractBackend) (*Migc, error) {
	contract, err := bindMigc(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Migc{MigcCaller: MigcCaller{contract: contract}, MigcTransactor: MigcTransactor{contract: contract}, MigcFilterer: MigcFilterer{contract: contract}}, nil
}

// NewMigcCaller creates a new read-only instance of Migc, bound to a specific deployed contract.
func NewMigcCaller(address common.Address, caller bind.ContractCaller) (*MigcCaller, error) {
	contract, err := bindMigc(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &MigcCaller{contract: contract}, nil
}

// NewMigcTransactor creates a new write-only instance of Migc, bound to a specific deployed contract.
func NewMigcTransactor(address common.Address, transactor bind.ContractTransactor) (*MigcTransactor, error) {
	contract, err := bindMigc(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &MigcTransactor{contract: contract}, nil
}

// NewMigcFilterer creates a new log filterer instance of Migc, bound to a specific deployed contract.
func NewMigcFilterer(address common.Address, filterer bind.ContractFilterer) (*MigcFilterer, error) {
	contract, err := bindMigc(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &MigcFilterer{contract: contract}, nil
}

// bindMigc binds a generic wrapper to an already deployed contract.
func bindMigc(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(MigcABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Migc *MigcRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Migc.Contract.MigcCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Migc *MigcRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Migc.Contract.MigcTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Migc *MigcRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Migc.Contract.MigcTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Migc *MigcCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Migc.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Migc *MigcTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Migc.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Migc *MigcTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Migc.Contract.contract.Transact(opts, method, params...)
}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_Migc *MigcCaller) DEFAULTADMINROLE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _Migc.contract.Call(opts, &out, "DEFAULT_ADMIN_ROLE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_Migc *MigcSession) DEFAULTADMINROLE() ([32]byte, error) {
	return _Migc.Contract.DEFAULTADMINROLE(&_Migc.CallOpts)
}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_Migc *MigcCallerSession) DEFAULTADMINROLE() ([32]byte, error) {
	return _Migc.Contract.DEFAULTADMINROLE(&_Migc.CallOpts)
}

// Datamap is a free data retrieval call binding the contract method 0x69309412.
//
// Solidity: function datamap(string ) view returns(address businessOwner, string name, string url, uint64 expired)
func (_Migc *MigcCaller) Datamap(opts *bind.CallOpts, arg0 string) (struct {
	BusinessOwner common.Address
	Name          string
	Url           string
	Expired       uint64
}, error) {
	var out []interface{}
	err := _Migc.contract.Call(opts, &out, "datamap", arg0)

	outstruct := new(struct {
		BusinessOwner common.Address
		Name          string
		Url           string
		Expired       uint64
	})

	outstruct.BusinessOwner = out[0].(common.Address)
	outstruct.Name = out[1].(string)
	outstruct.Url = out[2].(string)
	outstruct.Expired = out[3].(uint64)

	return *outstruct, err

}

// Datamap is a free data retrieval call binding the contract method 0x69309412.
//
// Solidity: function datamap(string ) view returns(address businessOwner, string name, string url, uint64 expired)
func (_Migc *MigcSession) Datamap(arg0 string) (struct {
	BusinessOwner common.Address
	Name          string
	Url           string
	Expired       uint64
}, error) {
	return _Migc.Contract.Datamap(&_Migc.CallOpts, arg0)
}

// Datamap is a free data retrieval call binding the contract method 0x69309412.
//
// Solidity: function datamap(string ) view returns(address businessOwner, string name, string url, uint64 expired)
func (_Migc *MigcCallerSession) Datamap(arg0 string) (struct {
	BusinessOwner common.Address
	Name          string
	Url           string
	Expired       uint64
}, error) {
	return _Migc.Contract.Datamap(&_Migc.CallOpts, arg0)
}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_Migc *MigcCaller) GetRoleAdmin(opts *bind.CallOpts, role [32]byte) ([32]byte, error) {
	var out []interface{}
	err := _Migc.contract.Call(opts, &out, "getRoleAdmin", role)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_Migc *MigcSession) GetRoleAdmin(role [32]byte) ([32]byte, error) {
	return _Migc.Contract.GetRoleAdmin(&_Migc.CallOpts, role)
}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_Migc *MigcCallerSession) GetRoleAdmin(role [32]byte) ([32]byte, error) {
	return _Migc.Contract.GetRoleAdmin(&_Migc.CallOpts, role)
}

// GetRoleMember is a free data retrieval call binding the contract method 0x9010d07c.
//
// Solidity: function getRoleMember(bytes32 role, uint256 index) view returns(address)
func (_Migc *MigcCaller) GetRoleMember(opts *bind.CallOpts, role [32]byte, index *big.Int) (common.Address, error) {
	var out []interface{}
	err := _Migc.contract.Call(opts, &out, "getRoleMember", role, index)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetRoleMember is a free data retrieval call binding the contract method 0x9010d07c.
//
// Solidity: function getRoleMember(bytes32 role, uint256 index) view returns(address)
func (_Migc *MigcSession) GetRoleMember(role [32]byte, index *big.Int) (common.Address, error) {
	return _Migc.Contract.GetRoleMember(&_Migc.CallOpts, role, index)
}

// GetRoleMember is a free data retrieval call binding the contract method 0x9010d07c.
//
// Solidity: function getRoleMember(bytes32 role, uint256 index) view returns(address)
func (_Migc *MigcCallerSession) GetRoleMember(role [32]byte, index *big.Int) (common.Address, error) {
	return _Migc.Contract.GetRoleMember(&_Migc.CallOpts, role, index)
}

// GetRoleMemberCount is a free data retrieval call binding the contract method 0xca15c873.
//
// Solidity: function getRoleMemberCount(bytes32 role) view returns(uint256)
func (_Migc *MigcCaller) GetRoleMemberCount(opts *bind.CallOpts, role [32]byte) (*big.Int, error) {
	var out []interface{}
	err := _Migc.contract.Call(opts, &out, "getRoleMemberCount", role)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetRoleMemberCount is a free data retrieval call binding the contract method 0xca15c873.
//
// Solidity: function getRoleMemberCount(bytes32 role) view returns(uint256)
func (_Migc *MigcSession) GetRoleMemberCount(role [32]byte) (*big.Int, error) {
	return _Migc.Contract.GetRoleMemberCount(&_Migc.CallOpts, role)
}

// GetRoleMemberCount is a free data retrieval call binding the contract method 0xca15c873.
//
// Solidity: function getRoleMemberCount(bytes32 role) view returns(uint256)
func (_Migc *MigcCallerSession) GetRoleMemberCount(role [32]byte) (*big.Int, error) {
	return _Migc.Contract.GetRoleMemberCount(&_Migc.CallOpts, role)
}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_Migc *MigcCaller) HasRole(opts *bind.CallOpts, role [32]byte, account common.Address) (bool, error) {
	var out []interface{}
	err := _Migc.contract.Call(opts, &out, "hasRole", role, account)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_Migc *MigcSession) HasRole(role [32]byte, account common.Address) (bool, error) {
	return _Migc.Contract.HasRole(&_Migc.CallOpts, role, account)
}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_Migc *MigcCallerSession) HasRole(role [32]byte, account common.Address) (bool, error) {
	return _Migc.Contract.HasRole(&_Migc.CallOpts, role, account)
}

// ListURLCount is a free data retrieval call binding the contract method 0xfaffb97c.
//
// Solidity: function listURLCount() view returns(uint64)
func (_Migc *MigcCaller) ListURLCount(opts *bind.CallOpts) (uint64, error) {
	var out []interface{}
	err := _Migc.contract.Call(opts, &out, "listURLCount")

	if err != nil {
		return *new(uint64), err
	}

	out0 := *abi.ConvertType(out[0], new(uint64)).(*uint64)

	return out0, err

}

// ListURLCount is a free data retrieval call binding the contract method 0xfaffb97c.
//
// Solidity: function listURLCount() view returns(uint64)
func (_Migc *MigcSession) ListURLCount() (uint64, error) {
	return _Migc.Contract.ListURLCount(&_Migc.CallOpts)
}

// ListURLCount is a free data retrieval call binding the contract method 0xfaffb97c.
//
// Solidity: function listURLCount() view returns(uint64)
func (_Migc *MigcCallerSession) ListURLCount() (uint64, error) {
	return _Migc.Contract.ListURLCount(&_Migc.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Migc *MigcCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Migc.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Migc *MigcSession) Owner() (common.Address, error) {
	return _Migc.Contract.Owner(&_Migc.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Migc *MigcCallerSession) Owner() (common.Address, error) {
	return _Migc.Contract.Owner(&_Migc.CallOpts)
}

// Urls is a free data retrieval call binding the contract method 0x796676be.
//
// Solidity: function urls(uint256 ) view returns(string)
func (_Migc *MigcCaller) Urls(opts *bind.CallOpts, arg0 *big.Int) (string, error) {
	var out []interface{}
	err := _Migc.contract.Call(opts, &out, "urls", arg0)

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Urls is a free data retrieval call binding the contract method 0x796676be.
//
// Solidity: function urls(uint256 ) view returns(string)
func (_Migc *MigcSession) Urls(arg0 *big.Int) (string, error) {
	return _Migc.Contract.Urls(&_Migc.CallOpts, arg0)
}

// Urls is a free data retrieval call binding the contract method 0x796676be.
//
// Solidity: function urls(uint256 ) view returns(string)
func (_Migc *MigcCallerSession) Urls(arg0 *big.Int) (string, error) {
	return _Migc.Contract.Urls(&_Migc.CallOpts, arg0)
}

// CreateBusiness is a paid mutator transaction binding the contract method 0x33d10472.
//
// Solidity: function createBusiness(string url, string name) returns()
func (_Migc *MigcTransactor) CreateBusiness(opts *bind.TransactOpts, url string, name string) (*types.Transaction, error) {
	return _Migc.contract.Transact(opts, "createBusiness", url, name)
}

// CreateBusiness is a paid mutator transaction binding the contract method 0x33d10472.
//
// Solidity: function createBusiness(string url, string name) returns()
func (_Migc *MigcSession) CreateBusiness(url string, name string) (*types.Transaction, error) {
	return _Migc.Contract.CreateBusiness(&_Migc.TransactOpts, url, name)
}

// CreateBusiness is a paid mutator transaction binding the contract method 0x33d10472.
//
// Solidity: function createBusiness(string url, string name) returns()
func (_Migc *MigcTransactorSession) CreateBusiness(url string, name string) (*types.Transaction, error) {
	return _Migc.Contract.CreateBusiness(&_Migc.TransactOpts, url, name)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_Migc *MigcTransactor) GrantRole(opts *bind.TransactOpts, role [32]byte, account common.Address) (*types.Transaction, error) {
	return _Migc.contract.Transact(opts, "grantRole", role, account)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_Migc *MigcSession) GrantRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _Migc.Contract.GrantRole(&_Migc.TransactOpts, role, account)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_Migc *MigcTransactorSession) GrantRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _Migc.Contract.GrantRole(&_Migc.TransactOpts, role, account)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Migc *MigcTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Migc.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Migc *MigcSession) RenounceOwnership() (*types.Transaction, error) {
	return _Migc.Contract.RenounceOwnership(&_Migc.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Migc *MigcTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _Migc.Contract.RenounceOwnership(&_Migc.TransactOpts)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address account) returns()
func (_Migc *MigcTransactor) RenounceRole(opts *bind.TransactOpts, role [32]byte, account common.Address) (*types.Transaction, error) {
	return _Migc.contract.Transact(opts, "renounceRole", role, account)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address account) returns()
func (_Migc *MigcSession) RenounceRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _Migc.Contract.RenounceRole(&_Migc.TransactOpts, role, account)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address account) returns()
func (_Migc *MigcTransactorSession) RenounceRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _Migc.Contract.RenounceRole(&_Migc.TransactOpts, role, account)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_Migc *MigcTransactor) RevokeRole(opts *bind.TransactOpts, role [32]byte, account common.Address) (*types.Transaction, error) {
	return _Migc.contract.Transact(opts, "revokeRole", role, account)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_Migc *MigcSession) RevokeRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _Migc.Contract.RevokeRole(&_Migc.TransactOpts, role, account)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_Migc *MigcTransactorSession) RevokeRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _Migc.Contract.RevokeRole(&_Migc.TransactOpts, role, account)
}

// TransferBusiness is a paid mutator transaction binding the contract method 0x60d41a26.
//
// Solidity: function transferBusiness(string url, address toOwner) returns()
func (_Migc *MigcTransactor) TransferBusiness(opts *bind.TransactOpts, url string, toOwner common.Address) (*types.Transaction, error) {
	return _Migc.contract.Transact(opts, "transferBusiness", url, toOwner)
}

// TransferBusiness is a paid mutator transaction binding the contract method 0x60d41a26.
//
// Solidity: function transferBusiness(string url, address toOwner) returns()
func (_Migc *MigcSession) TransferBusiness(url string, toOwner common.Address) (*types.Transaction, error) {
	return _Migc.Contract.TransferBusiness(&_Migc.TransactOpts, url, toOwner)
}

// TransferBusiness is a paid mutator transaction binding the contract method 0x60d41a26.
//
// Solidity: function transferBusiness(string url, address toOwner) returns()
func (_Migc *MigcTransactorSession) TransferBusiness(url string, toOwner common.Address) (*types.Transaction, error) {
	return _Migc.Contract.TransferBusiness(&_Migc.TransactOpts, url, toOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Migc *MigcTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _Migc.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Migc *MigcSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Migc.Contract.TransferOwnership(&_Migc.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Migc *MigcTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Migc.Contract.TransferOwnership(&_Migc.TransactOpts, newOwner)
}

// WeeklyUpdate is a paid mutator transaction binding the contract method 0xc609d816.
//
// Solidity: function weeklyUpdate(string url) returns()
func (_Migc *MigcTransactor) WeeklyUpdate(opts *bind.TransactOpts, url string) (*types.Transaction, error) {
	return _Migc.contract.Transact(opts, "weeklyUpdate", url)
}

// WeeklyUpdate is a paid mutator transaction binding the contract method 0xc609d816.
//
// Solidity: function weeklyUpdate(string url) returns()
func (_Migc *MigcSession) WeeklyUpdate(url string) (*types.Transaction, error) {
	return _Migc.Contract.WeeklyUpdate(&_Migc.TransactOpts, url)
}

// WeeklyUpdate is a paid mutator transaction binding the contract method 0xc609d816.
//
// Solidity: function weeklyUpdate(string url) returns()
func (_Migc *MigcTransactorSession) WeeklyUpdate(url string) (*types.Transaction, error) {
	return _Migc.Contract.WeeklyUpdate(&_Migc.TransactOpts, url)
}

// MigcBusinessCreatedIterator is returned from FilterBusinessCreated and is used to iterate over the raw logs and unpacked data for BusinessCreated events raised by the Migc contract.
type MigcBusinessCreatedIterator struct {
	Event *MigcBusinessCreated // Event containing the contract specifics and raw log

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
func (it *MigcBusinessCreatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MigcBusinessCreated)
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
		it.Event = new(MigcBusinessCreated)
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
func (it *MigcBusinessCreatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MigcBusinessCreatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MigcBusinessCreated represents a BusinessCreated event raised by the Migc contract.
type MigcBusinessCreated struct {
	Url  common.Hash
	Name common.Hash
	Raw  types.Log // Blockchain specific contextual infos
}

// FilterBusinessCreated is a free log retrieval operation binding the contract event 0x6fa45dba9b54e84243956c5756b3d261bd8028ac70962f42e3bdb1b839e45c41.
//
// Solidity: event BusinessCreated(string indexed url, string indexed name)
func (_Migc *MigcFilterer) FilterBusinessCreated(opts *bind.FilterOpts, url []string, name []string) (*MigcBusinessCreatedIterator, error) {

	var urlRule []interface{}
	for _, urlItem := range url {
		urlRule = append(urlRule, urlItem)
	}
	var nameRule []interface{}
	for _, nameItem := range name {
		nameRule = append(nameRule, nameItem)
	}

	logs, sub, err := _Migc.contract.FilterLogs(opts, "BusinessCreated", urlRule, nameRule)
	if err != nil {
		return nil, err
	}
	return &MigcBusinessCreatedIterator{contract: _Migc.contract, event: "BusinessCreated", logs: logs, sub: sub}, nil
}

// WatchBusinessCreated is a free log subscription operation binding the contract event 0x6fa45dba9b54e84243956c5756b3d261bd8028ac70962f42e3bdb1b839e45c41.
//
// Solidity: event BusinessCreated(string indexed url, string indexed name)
func (_Migc *MigcFilterer) WatchBusinessCreated(opts *bind.WatchOpts, sink chan<- *MigcBusinessCreated, url []string, name []string) (event.Subscription, error) {

	var urlRule []interface{}
	for _, urlItem := range url {
		urlRule = append(urlRule, urlItem)
	}
	var nameRule []interface{}
	for _, nameItem := range name {
		nameRule = append(nameRule, nameItem)
	}

	logs, sub, err := _Migc.contract.WatchLogs(opts, "BusinessCreated", urlRule, nameRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MigcBusinessCreated)
				if err := _Migc.contract.UnpackLog(event, "BusinessCreated", log); err != nil {
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

// ParseBusinessCreated is a log parse operation binding the contract event 0x6fa45dba9b54e84243956c5756b3d261bd8028ac70962f42e3bdb1b839e45c41.
//
// Solidity: event BusinessCreated(string indexed url, string indexed name)
func (_Migc *MigcFilterer) ParseBusinessCreated(log types.Log) (*MigcBusinessCreated, error) {
	event := new(MigcBusinessCreated)
	if err := _Migc.contract.UnpackLog(event, "BusinessCreated", log); err != nil {
		return nil, err
	}
	return event, nil
}

// MigcBusinessTransferIterator is returned from FilterBusinessTransfer and is used to iterate over the raw logs and unpacked data for BusinessTransfer events raised by the Migc contract.
type MigcBusinessTransferIterator struct {
	Event *MigcBusinessTransfer // Event containing the contract specifics and raw log

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
func (it *MigcBusinessTransferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MigcBusinessTransfer)
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
		it.Event = new(MigcBusinessTransfer)
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
func (it *MigcBusinessTransferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MigcBusinessTransferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MigcBusinessTransfer represents a BusinessTransfer event raised by the Migc contract.
type MigcBusinessTransfer struct {
	FromOwner common.Address
	ToOwner   common.Address
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterBusinessTransfer is a free log retrieval operation binding the contract event 0x54bdaa308d34406b09b20101de9b6467c26e5901ee017b79c618c70cbe117cf9.
//
// Solidity: event BusinessTransfer(address indexed fromOwner, address indexed toOwner)
func (_Migc *MigcFilterer) FilterBusinessTransfer(opts *bind.FilterOpts, fromOwner []common.Address, toOwner []common.Address) (*MigcBusinessTransferIterator, error) {

	var fromOwnerRule []interface{}
	for _, fromOwnerItem := range fromOwner {
		fromOwnerRule = append(fromOwnerRule, fromOwnerItem)
	}
	var toOwnerRule []interface{}
	for _, toOwnerItem := range toOwner {
		toOwnerRule = append(toOwnerRule, toOwnerItem)
	}

	logs, sub, err := _Migc.contract.FilterLogs(opts, "BusinessTransfer", fromOwnerRule, toOwnerRule)
	if err != nil {
		return nil, err
	}
	return &MigcBusinessTransferIterator{contract: _Migc.contract, event: "BusinessTransfer", logs: logs, sub: sub}, nil
}

// WatchBusinessTransfer is a free log subscription operation binding the contract event 0x54bdaa308d34406b09b20101de9b6467c26e5901ee017b79c618c70cbe117cf9.
//
// Solidity: event BusinessTransfer(address indexed fromOwner, address indexed toOwner)
func (_Migc *MigcFilterer) WatchBusinessTransfer(opts *bind.WatchOpts, sink chan<- *MigcBusinessTransfer, fromOwner []common.Address, toOwner []common.Address) (event.Subscription, error) {

	var fromOwnerRule []interface{}
	for _, fromOwnerItem := range fromOwner {
		fromOwnerRule = append(fromOwnerRule, fromOwnerItem)
	}
	var toOwnerRule []interface{}
	for _, toOwnerItem := range toOwner {
		toOwnerRule = append(toOwnerRule, toOwnerItem)
	}

	logs, sub, err := _Migc.contract.WatchLogs(opts, "BusinessTransfer", fromOwnerRule, toOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MigcBusinessTransfer)
				if err := _Migc.contract.UnpackLog(event, "BusinessTransfer", log); err != nil {
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

// ParseBusinessTransfer is a log parse operation binding the contract event 0x54bdaa308d34406b09b20101de9b6467c26e5901ee017b79c618c70cbe117cf9.
//
// Solidity: event BusinessTransfer(address indexed fromOwner, address indexed toOwner)
func (_Migc *MigcFilterer) ParseBusinessTransfer(log types.Log) (*MigcBusinessTransfer, error) {
	event := new(MigcBusinessTransfer)
	if err := _Migc.contract.UnpackLog(event, "BusinessTransfer", log); err != nil {
		return nil, err
	}
	return event, nil
}

// MigcBusinessUpdateIterator is returned from FilterBusinessUpdate and is used to iterate over the raw logs and unpacked data for BusinessUpdate events raised by the Migc contract.
type MigcBusinessUpdateIterator struct {
	Event *MigcBusinessUpdate // Event containing the contract specifics and raw log

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
func (it *MigcBusinessUpdateIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MigcBusinessUpdate)
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
		it.Event = new(MigcBusinessUpdate)
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
func (it *MigcBusinessUpdateIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MigcBusinessUpdateIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MigcBusinessUpdate represents a BusinessUpdate event raised by the Migc contract.
type MigcBusinessUpdate struct {
	Url     common.Hash
	Expired *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterBusinessUpdate is a free log retrieval operation binding the contract event 0x6db6455c715c1587e7c21e5b3ed51aa8f6bd2dc9d646376e843daf1ea59c2c5f.
//
// Solidity: event BusinessUpdate(string indexed url, uint256 expired)
func (_Migc *MigcFilterer) FilterBusinessUpdate(opts *bind.FilterOpts, url []string) (*MigcBusinessUpdateIterator, error) {

	var urlRule []interface{}
	for _, urlItem := range url {
		urlRule = append(urlRule, urlItem)
	}

	logs, sub, err := _Migc.contract.FilterLogs(opts, "BusinessUpdate", urlRule)
	if err != nil {
		return nil, err
	}
	return &MigcBusinessUpdateIterator{contract: _Migc.contract, event: "BusinessUpdate", logs: logs, sub: sub}, nil
}

// WatchBusinessUpdate is a free log subscription operation binding the contract event 0x6db6455c715c1587e7c21e5b3ed51aa8f6bd2dc9d646376e843daf1ea59c2c5f.
//
// Solidity: event BusinessUpdate(string indexed url, uint256 expired)
func (_Migc *MigcFilterer) WatchBusinessUpdate(opts *bind.WatchOpts, sink chan<- *MigcBusinessUpdate, url []string) (event.Subscription, error) {

	var urlRule []interface{}
	for _, urlItem := range url {
		urlRule = append(urlRule, urlItem)
	}

	logs, sub, err := _Migc.contract.WatchLogs(opts, "BusinessUpdate", urlRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MigcBusinessUpdate)
				if err := _Migc.contract.UnpackLog(event, "BusinessUpdate", log); err != nil {
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

// ParseBusinessUpdate is a log parse operation binding the contract event 0x6db6455c715c1587e7c21e5b3ed51aa8f6bd2dc9d646376e843daf1ea59c2c5f.
//
// Solidity: event BusinessUpdate(string indexed url, uint256 expired)
func (_Migc *MigcFilterer) ParseBusinessUpdate(log types.Log) (*MigcBusinessUpdate, error) {
	event := new(MigcBusinessUpdate)
	if err := _Migc.contract.UnpackLog(event, "BusinessUpdate", log); err != nil {
		return nil, err
	}
	return event, nil
}

// MigcOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the Migc contract.
type MigcOwnershipTransferredIterator struct {
	Event *MigcOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *MigcOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MigcOwnershipTransferred)
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
		it.Event = new(MigcOwnershipTransferred)
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
func (it *MigcOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MigcOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MigcOwnershipTransferred represents a OwnershipTransferred event raised by the Migc contract.
type MigcOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Migc *MigcFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*MigcOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Migc.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &MigcOwnershipTransferredIterator{contract: _Migc.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Migc *MigcFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *MigcOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Migc.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MigcOwnershipTransferred)
				if err := _Migc.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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

// ParseOwnershipTransferred is a log parse operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Migc *MigcFilterer) ParseOwnershipTransferred(log types.Log) (*MigcOwnershipTransferred, error) {
	event := new(MigcOwnershipTransferred)
	if err := _Migc.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	return event, nil
}

// MigcRoleAdminChangedIterator is returned from FilterRoleAdminChanged and is used to iterate over the raw logs and unpacked data for RoleAdminChanged events raised by the Migc contract.
type MigcRoleAdminChangedIterator struct {
	Event *MigcRoleAdminChanged // Event containing the contract specifics and raw log

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
func (it *MigcRoleAdminChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MigcRoleAdminChanged)
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
		it.Event = new(MigcRoleAdminChanged)
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
func (it *MigcRoleAdminChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MigcRoleAdminChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MigcRoleAdminChanged represents a RoleAdminChanged event raised by the Migc contract.
type MigcRoleAdminChanged struct {
	Role              [32]byte
	PreviousAdminRole [32]byte
	NewAdminRole      [32]byte
	Raw               types.Log // Blockchain specific contextual infos
}

// FilterRoleAdminChanged is a free log retrieval operation binding the contract event 0xbd79b86ffe0ab8e8776151514217cd7cacd52c909f66475c3af44e129f0b00ff.
//
// Solidity: event RoleAdminChanged(bytes32 indexed role, bytes32 indexed previousAdminRole, bytes32 indexed newAdminRole)
func (_Migc *MigcFilterer) FilterRoleAdminChanged(opts *bind.FilterOpts, role [][32]byte, previousAdminRole [][32]byte, newAdminRole [][32]byte) (*MigcRoleAdminChangedIterator, error) {

	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}
	var previousAdminRoleRule []interface{}
	for _, previousAdminRoleItem := range previousAdminRole {
		previousAdminRoleRule = append(previousAdminRoleRule, previousAdminRoleItem)
	}
	var newAdminRoleRule []interface{}
	for _, newAdminRoleItem := range newAdminRole {
		newAdminRoleRule = append(newAdminRoleRule, newAdminRoleItem)
	}

	logs, sub, err := _Migc.contract.FilterLogs(opts, "RoleAdminChanged", roleRule, previousAdminRoleRule, newAdminRoleRule)
	if err != nil {
		return nil, err
	}
	return &MigcRoleAdminChangedIterator{contract: _Migc.contract, event: "RoleAdminChanged", logs: logs, sub: sub}, nil
}

// WatchRoleAdminChanged is a free log subscription operation binding the contract event 0xbd79b86ffe0ab8e8776151514217cd7cacd52c909f66475c3af44e129f0b00ff.
//
// Solidity: event RoleAdminChanged(bytes32 indexed role, bytes32 indexed previousAdminRole, bytes32 indexed newAdminRole)
func (_Migc *MigcFilterer) WatchRoleAdminChanged(opts *bind.WatchOpts, sink chan<- *MigcRoleAdminChanged, role [][32]byte, previousAdminRole [][32]byte, newAdminRole [][32]byte) (event.Subscription, error) {

	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}
	var previousAdminRoleRule []interface{}
	for _, previousAdminRoleItem := range previousAdminRole {
		previousAdminRoleRule = append(previousAdminRoleRule, previousAdminRoleItem)
	}
	var newAdminRoleRule []interface{}
	for _, newAdminRoleItem := range newAdminRole {
		newAdminRoleRule = append(newAdminRoleRule, newAdminRoleItem)
	}

	logs, sub, err := _Migc.contract.WatchLogs(opts, "RoleAdminChanged", roleRule, previousAdminRoleRule, newAdminRoleRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MigcRoleAdminChanged)
				if err := _Migc.contract.UnpackLog(event, "RoleAdminChanged", log); err != nil {
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

// ParseRoleAdminChanged is a log parse operation binding the contract event 0xbd79b86ffe0ab8e8776151514217cd7cacd52c909f66475c3af44e129f0b00ff.
//
// Solidity: event RoleAdminChanged(bytes32 indexed role, bytes32 indexed previousAdminRole, bytes32 indexed newAdminRole)
func (_Migc *MigcFilterer) ParseRoleAdminChanged(log types.Log) (*MigcRoleAdminChanged, error) {
	event := new(MigcRoleAdminChanged)
	if err := _Migc.contract.UnpackLog(event, "RoleAdminChanged", log); err != nil {
		return nil, err
	}
	return event, nil
}

// MigcRoleGrantedIterator is returned from FilterRoleGranted and is used to iterate over the raw logs and unpacked data for RoleGranted events raised by the Migc contract.
type MigcRoleGrantedIterator struct {
	Event *MigcRoleGranted // Event containing the contract specifics and raw log

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
func (it *MigcRoleGrantedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MigcRoleGranted)
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
		it.Event = new(MigcRoleGranted)
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
func (it *MigcRoleGrantedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MigcRoleGrantedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MigcRoleGranted represents a RoleGranted event raised by the Migc contract.
type MigcRoleGranted struct {
	Role    [32]byte
	Account common.Address
	Sender  common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterRoleGranted is a free log retrieval operation binding the contract event 0x2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d.
//
// Solidity: event RoleGranted(bytes32 indexed role, address indexed account, address indexed sender)
func (_Migc *MigcFilterer) FilterRoleGranted(opts *bind.FilterOpts, role [][32]byte, account []common.Address, sender []common.Address) (*MigcRoleGrantedIterator, error) {

	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}
	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}
	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _Migc.contract.FilterLogs(opts, "RoleGranted", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return &MigcRoleGrantedIterator{contract: _Migc.contract, event: "RoleGranted", logs: logs, sub: sub}, nil
}

// WatchRoleGranted is a free log subscription operation binding the contract event 0x2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d.
//
// Solidity: event RoleGranted(bytes32 indexed role, address indexed account, address indexed sender)
func (_Migc *MigcFilterer) WatchRoleGranted(opts *bind.WatchOpts, sink chan<- *MigcRoleGranted, role [][32]byte, account []common.Address, sender []common.Address) (event.Subscription, error) {

	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}
	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}
	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _Migc.contract.WatchLogs(opts, "RoleGranted", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MigcRoleGranted)
				if err := _Migc.contract.UnpackLog(event, "RoleGranted", log); err != nil {
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

// ParseRoleGranted is a log parse operation binding the contract event 0x2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d.
//
// Solidity: event RoleGranted(bytes32 indexed role, address indexed account, address indexed sender)
func (_Migc *MigcFilterer) ParseRoleGranted(log types.Log) (*MigcRoleGranted, error) {
	event := new(MigcRoleGranted)
	if err := _Migc.contract.UnpackLog(event, "RoleGranted", log); err != nil {
		return nil, err
	}
	return event, nil
}

// MigcRoleRevokedIterator is returned from FilterRoleRevoked and is used to iterate over the raw logs and unpacked data for RoleRevoked events raised by the Migc contract.
type MigcRoleRevokedIterator struct {
	Event *MigcRoleRevoked // Event containing the contract specifics and raw log

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
func (it *MigcRoleRevokedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MigcRoleRevoked)
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
		it.Event = new(MigcRoleRevoked)
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
func (it *MigcRoleRevokedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MigcRoleRevokedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MigcRoleRevoked represents a RoleRevoked event raised by the Migc contract.
type MigcRoleRevoked struct {
	Role    [32]byte
	Account common.Address
	Sender  common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterRoleRevoked is a free log retrieval operation binding the contract event 0xf6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b.
//
// Solidity: event RoleRevoked(bytes32 indexed role, address indexed account, address indexed sender)
func (_Migc *MigcFilterer) FilterRoleRevoked(opts *bind.FilterOpts, role [][32]byte, account []common.Address, sender []common.Address) (*MigcRoleRevokedIterator, error) {

	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}
	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}
	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _Migc.contract.FilterLogs(opts, "RoleRevoked", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return &MigcRoleRevokedIterator{contract: _Migc.contract, event: "RoleRevoked", logs: logs, sub: sub}, nil
}

// WatchRoleRevoked is a free log subscription operation binding the contract event 0xf6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b.
//
// Solidity: event RoleRevoked(bytes32 indexed role, address indexed account, address indexed sender)
func (_Migc *MigcFilterer) WatchRoleRevoked(opts *bind.WatchOpts, sink chan<- *MigcRoleRevoked, role [][32]byte, account []common.Address, sender []common.Address) (event.Subscription, error) {

	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}
	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}
	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _Migc.contract.WatchLogs(opts, "RoleRevoked", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MigcRoleRevoked)
				if err := _Migc.contract.UnpackLog(event, "RoleRevoked", log); err != nil {
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

// ParseRoleRevoked is a log parse operation binding the contract event 0xf6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b.
//
// Solidity: event RoleRevoked(bytes32 indexed role, address indexed account, address indexed sender)
func (_Migc *MigcFilterer) ParseRoleRevoked(log types.Log) (*MigcRoleRevoked, error) {
	event := new(MigcRoleRevoked)
	if err := _Migc.contract.UnpackLog(event, "RoleRevoked", log); err != nil {
		return nil, err
	}
	return event, nil
}
