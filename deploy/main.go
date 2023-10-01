package main

import (
	"context"
	"fmt"
	todo "github.com/devstackq/smart-contract/gen"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"

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
	trnOpts.GasPrice = gasPrice
	trnOpts.Nonce = big.NewInt(int64(nonce))
	trnOpts.GasLimit = uint64(300000)

	// DeployTodo deploys a new Ethereum contract
	resp, tx, _, err := todo.DeployTodo(trnOpts, cl)
	if err != nil {
		fmt.Printf("err DeployTodo %v ", err)
		return
	}
	fmt.Printf(" resp %v tx %v ", resp, tx)

}
