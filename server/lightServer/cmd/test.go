package main

import (
	"context"
	"crypto/ecdsa"
	"encoding/json"
	"fmt"
	"github.com/ethereum/go-ethereum/accounts"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/accounts/keystore"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"lightServer/mig/migc"
	"math/big"
	"math/rand"
	"time"
)

// 0xdac155b75872349f5fe720df16396bf385d244d5dc5097b8408bc1aa6bb734b8
//0x76cf1bd281c187a859e425d0ed7970af3e7f3212
//
const KeyStoreDir = `E:\Dev\localEth\testnet01\keystore`
const ContractAddress = "0x9e7495ba2dd6Cc064D6315511cCDA0aF32d3F966"
const Owner = "0xc825f9cd2dc8db9648ab2772170bfac6a022301a"
const OwnerPW = "i"

var ETH = big.NewInt(1000000000000000000)

func randAddr() common.Address {
	var bts = make([]byte, common.AddressLength)
	rand.New(rand.NewSource(time.Now().Unix())).Read(bts)
	return common.BytesToAddress(bts)
}
func ntx(key *ecdsa.PrivateKey, backend bind.ContractBackend, addr common.Address) *bind.TransactOpts {
	nonce, err := backend.PendingNonceAt(context.Background(), addr)
	if err != nil {
		panic(err)
	}

	gasPrice, err := backend.SuggestGasPrice(context.Background())
	if err != nil {
		panic(err)
	}
	txopt := bind.NewKeyedTransactor(key)
	txopt.From = addr
	txopt.Nonce = big.NewInt(0).SetUint64(nonce)
	txopt.GasPrice = gasPrice
	txopt.GasLimit = 30000000
	//txopt.GasLimit = 0
	return txopt
}
func jsonView(i interface{}) {
	j, _ := json.MarshalIndent(i, "", "    ")
	fmt.Println(string(j))
}

var (
	ks   *keystore.KeyStore
	ethc *ethclient.Client
	me   accounts.Account
)

func main() {
	install()
	defer uninstall()
	//
	//deploy(func(contract *migc.Migc) {
	//	count, err := contract.ListURLCount(&bind.CallOpts{})
	//	if err != nil {
	//		panic(err)
	//	}
	//	fmt.Println(count)
	//})
	//use(func(contract *migc.Migc) {
	//	tx, err := contract.CreateBusiness(newTxo(me), "http://ddiggo.iptime.org", "음식점")
	//	if err != nil {
	//		panic(err)
	//	}
	//	fmt.Println(tx.Hash().Hex())
	//})
	//use(func(contract *migc.Migc) {
	//	rx, err := contract.Datamap(&bind.CallOpts{}, "http://ddiggo.iptime.org")
	//	if err != nil {
	//		panic(err)
	//	}
	//	fmt.Println(time.Unix(int64(rx.Expired), 0).Format(time.RFC3339))
	//})
	use(func(contract *migc.Migc) {
		count, err := contract.ListURLCount(&bind.CallOpts{})
		if err != nil {
			panic(err)
		}
		var tmp = make([]string, int(count))
		for i := uint64(0); i < count; i++ {
			tmp[i], _ = contract.Urls(&bind.CallOpts{}, new(big.Int).SetUint64(i))
		}
		for i, s := range tmp {
			fmt.Println(i, ":", s)
		}
	})
}
func install() {
	var err error
	ks = keystore.NewKeyStore(KeyStoreDir, keystore.LightScryptN, keystore.LightScryptP)
	me, err = ks.Find(accounts.Account{Address: common.HexToAddress(Owner)})
	if err != nil {
		panic(err)
	}
	err = ks.Unlock(me, OwnerPW)
	if err != nil {
		panic(err)
	}
	ethc, err = ethclient.Dial("http://localhost:8545")
	if err != nil {
		panic(err)
	}
}
func uninstall() {
	ethc.Close()
}
func deploy(f func(contract *migc.Migc)) {
	addr, deploytx, contract, err := migc.DeployMigc(newTxo(me), ethc)
	if err != nil {
		panic(err)
	}
	jsonView(deploytx)
	fmt.Println("ADDRESS : ", addr.Hex())
	fmt.Println("Contract : ", contract)
	f(contract)
}
func newTxo(who accounts.Account) *bind.TransactOpts {
	var (
		err  error
		ksto *bind.TransactOpts
		n    uint64
	)
	ksto, err = bind.NewKeyStoreTransactor(ks, who)
	if err != nil {
		panic(err)
	}
	// 가스 추정, nonce 셋업
	ksto.GasPrice, err = ethc.SuggestGasPrice(context.Background())
	if err != nil {
		panic(err)
	}
	n, err = ethc.NonceAt(context.Background(), ksto.From, nil)
	if err != nil {
		panic(err)
	}
	ksto.Nonce = new(big.Int).SetUint64(n)
	block, err := ethc.BlockByNumber(context.Background(), nil)
	if err != nil {
		panic(err)
	}
	ksto.GasLimit = block.GasLimit()
	return ksto
}
func use(f func(contract *migc.Migc)) {
	c, err := migc.NewMigc(common.HexToAddress(ContractAddress), ethc)
	if err != nil {
		panic(err)
	}
	f(c)
}
