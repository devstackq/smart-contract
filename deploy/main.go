package main

import (
	"context"
	"fmt"
	todo "github.com/devstackq/smart-contract/gen"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"

	"github.com/ethereum/go-ethereum/accounts/keystore"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"math/big"
	"os"
)

const (
	pathToKeyStore = "../wallet/keystore"
	polygonURL     = "https://polygon-mainnet.infura.io/v3/5d4822d03f3940748e09a54592655fd5"
)

func main() {
	ctx := context.Background()
	//read file
	keyStoreFile, err := os.ReadFile(pathToKeyStore)
	if err != nil {
		fmt.Printf("err read file %v", err)
		return
	}

	//decrypt file, by password
	key, err := keystore.DecryptKey(keyStoreFile, "skr123")
	if err != nil {
		fmt.Printf("err DecryptKey %v ", err)
		return
	}
	//connect to
	cl, err := ethclient.DialContext(ctx, polygonURL)
	if err != nil {
		fmt.Printf("err DialContext %v ", err)
		return
	}
	defer cl.Close()
	//pKey := "7bd50df4554e459af04e94576fc0e4b7d21a16b96abb857baa322f1b97f64ea0"
	//pubKey := "0x72b08e6514399239348978D1c101d71A154945A3"
	//get address by pubKey
	addr := crypto.PubkeyToAddress(key.PrivateKey.PublicKey)

	//uniq number for create uniq transaction
	nonce, err := cl.PendingNonceAt(ctx, addr)
	if err != nil {
		fmt.Printf("err PendingNonceAt %v ", err)
		return
	}

	//get actual gasPrice
	gasPrice, err := cl.SuggestGasPrice(ctx)
	if err != nil {
		fmt.Printf("err SuggestGasPrice %v ", err)
		return
	}
	//get networkID for current client
	chainID, err := cl.NetworkID(ctx)
	if err != nil {
		fmt.Printf("err NetworkID %v ", err)
		return
	}
	//create trans signer
	trnOpts, err := bind.NewKeyedTransactorWithChainID(key.PrivateKey, chainID)
	if err != nil {
		fmt.Printf("err NewKeyedTransactorWithChainID %v ", err)
		return
	}
	//prepare trans
	trnOpts.GasPrice = gasPrice //verification Fee, eth equaivilanet
	trnOpts.Nonce = big.NewInt(int64(nonce))
	trnOpts.GasLimit = uint64(300000)
	trnOpts.Value = big.NewInt(1) //send 1 eth - send 1eth contract;

	// DeployTodo deploys a new Ethereum contract
	//resp, tx, _, err := todo.DeployTodo(trnOpts, cl)
	//if err != nil {
	//	fmt.Printf("err DeployTodo %v ", err)
	//	return
	//}
	//fmt.Printf(" resp %+v tx %+v ", resp, tx)
	// do insufficient funds for gas * price + value ; ->  not enough Balance

	//next step - get created Contract address
	contractAddress := common.HexToAddress("0x48B63160d3Bbc7B566aa596768a19523278C1525")

	solidInterface, err := todo.NewTodo(contractAddress, cl)
	if err != nil {
		fmt.Printf("err NewTodo %v ", err)
		return
	}

	//Caller -> layer -> abigen -> golang -> solidity bin -> eth network

	//call method - solc - generated abigen - solidity methods
	//work with 1 client, 1 contract - any methods
	trx, err := solidInterface.Add(trnOpts, "hello world !", "bekhzna")
	if err != nil {
		fmt.Printf("err solidInterface Add %v ", err)
		return
	}

	tasks, err := solidInterface.List(&bind.CallOpts{
		From: addr,
	})
	if err != nil {
		fmt.Printf("err solidInterface List %v ", err)
		return
	}
	fmt.Printf("trx Add() %+v list %+v ", trx, tasks)

	//solidInterface.TodoTransactor.Update()
}
