package main

import (
	"context"
	"fmt"
	"github.com/ethereum/go-ethereum/accounts/keystore"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"math/big"
	"os"
)

var infraEtheriumURL = "https://mainnet.infura.io/v3/5d4822d03f3940748e09a54592655fd5"

func main() {
	//cryptoF()
	//go getBalance()
	if err := getFreeEther(); err != nil {
		fmt.Printf("err getFreeEther", err)
		return
	}
}

const (
	keyDir         = "./wallet"
	password       = "skr123"
	pathToKeyStore = "./wallet/kystore"
	polygonURL     = "https://polygon-mainnet.infura.io/v3/5d4822d03f3940748e09a54592655fd5"
	address        = "3d533ee41b976dc626d381f8932f4bcf9bd3f3c1"
)

func transferCoin() error {
	//Miner have percent - each transaction
	return nil
}

func getFreeEther() error {
	//1.account - make transaction many network; 1 network - have own;
	//ex: Eth -> goerli - address23 = balance 2.4; Mainnet = address23 = 45.1
	//2.Miner have percent - each transaction

	ctx := context.Background()

	cl, err := ethclient.DialContext(ctx, polygonURL)
	if err != nil {
		return err
	}
	defer cl.Close()
	address1 := common.HexToAddress(address)

	balance, err := cl.BalanceAt(ctx, address1, nil)
	if err != nil {
		return err
	}
	fmt.Printf("balance %+v \n", balance)

	nonce, err := cl.PendingNonceAt(ctx, address1)
	if err != nil {
		return err
	}
	gasPrice, err := cl.SuggestGasPrice(ctx)
	if err != nil {
		return err
	}
	amount := big.NewInt(1000000000000000000)
	trx := types.NewTransaction(nonce, address1, amount, 21000, gasPrice, nil)
	//todo: address2; balance - add
	chaindID, err := cl.NetworkID(ctx)
	data, err := os.ReadFile(pathToKeyStore)
	if err != nil {
		return err
	}
	k, err := keystore.DecryptKey(data, password)
	if err != nil {
		return err
	}
	t, err := types.SignTx(trx, types.NewEIP155Signer(chaindID), k.PrivateKey)
	if err != nil {
		return err
	}
	if err = cl.SendTransaction(ctx, t); err != nil {
		return err
	}
	fmt.Println("trans v", t.Hash().Hex())
	return nil
}

func readKeyStore() error {
	data, err := os.ReadFile(pathToKeyStore)
	if err != nil {
		return err
	}
	//fmt.Printf("file content%+v \n", string(data))
	//keystore.EncryptKey()
	decryptedKey, err := keystore.DecryptKey(data, password)
	if err != nil {
		return err
	}

	pData := crypto.FromECDSA(decryptedKey.PrivateKey)
	fmt.Printf("decrypted from File, privateKey %+v \n", hexutil.Encode(pData))
	pubData := crypto.FromECDSAPub(&decryptedKey.PrivateKey.PublicKey)
	fmt.Printf("decrypted from File, privateKey %+v \n", hexutil.Encode(pData))
	fmt.Printf("decrypted from File, publicKey %+v \n", hexutil.Encode(pubData))

	return nil
}

func createKeyStore() error {
	key := keystore.NewKeyStore(keyDir, keystore.StandardScryptN, keystore.StandardScryptP)
	account, err := key.NewAccount(password)
	if err != nil {
		return err
	}
	fmt.Printf("account %+v", account)

	return nil
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
