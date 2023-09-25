package main

import (
	"context"
	"fmt"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"

	"math/big"
)

var infraEtheriumURL = "https://mainnet.infura.io/v3/5d4822d03f3940748e09a54592655fd5"

func main() {
	cryptoF()
	go getBalance()
}

func cryptoF() {
	pKey, err := crypto.GenerateKey()
	if err != nil {
		fmt.Println(err)
		return
	}
	ecdsaBytes := crypto.FromECDSA(pKey)
	pKeyStr := hexutil.Encode(ecdsaBytes)
	fmt.Println("privatKey generated", pKeyStr)

	public := crypto.FromECDSAPub(&pKey.PublicKey)
	fmt.Println("publicKey generated", hexutil.Encode(public))

	//crypto.PubkeyToAddress(pKey.PublicKey).Hex() //return public-like Address
	fmt.Println("publicKey -> like Address", crypto.PubkeyToAddress(pKey.PublicKey).Hex())

}

func getBalance() {
	ctx := context.Background()
	cl, err := ethclient.DialContext(ctx, infraEtheriumURL)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer cl.Close()

	block, err := cl.BlockByNumber(ctx, nil)

	bNumber, err := cl.BlockNumber(ctx)
	if err != nil {
		fmt.Println(err)
		return
	}
	for _, b := range block.Body().Transactions {
		fmt.Println("trans", b.Value().String())
	}

	fmt.Println(bNumber, "last trans number by Etherium") //las

	addrStr := "0x352EF1dC71438d3736319D50Ea2dE297C843E092"
	address := common.HexToAddress(addrStr)

	balance, err := cl.BalanceAt(ctx, address, nil)
	if err != nil {
		fmt.Println(err)
		return
	}
	fBalance := new(big.Float)
	fBalance.SetString(balance.String())

	fmt.Println("balance", balance.String(), fBalance)
}
