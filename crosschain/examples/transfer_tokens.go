package utils

import (
	"context"
	"crypto/ecdsa"
	"fmt"
	"log"
	"math/big"

	"os"

	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/joho/godotenv"
	"golang.org/x/crypto/sha3"
)

func TransferTokens() {

	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	rpc := os.Getenv("rpc_rinkeby")
	pk := os.Getenv("pk2")

	client, err := ethclient.Dial(rpc)
	if err != nil {
		log.Fatal(err)
	}

	privateKey, err := crypto.HexToECDSA(pk)
	if err != nil {
		log.Fatal(err)
	}

	publicKey := privateKey.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		log.Fatal("cannot assert type: publicKey is not of type *ecdsa.PublicKey")
	}

	fromAddress := crypto.PubkeyToAddress(*publicKeyECDSA)
	nonce, err := client.PendingNonceAt(context.Background(), fromAddress)
	if err != nil {
		log.Fatal(err)
	}

	value := big.NewInt(0) // in wei (0 eth)
	gasPrice, err := client.SuggestGasPrice(context.Background())
	if err != nil {
		log.Fatal(err)
	}

	WETH := "0xc778417E063141139Fce010982780140Aa0cD5Ab"

	toAddress := common.HexToAddress("0xB06903728e09748E3d941b83f1657B147bA045d2")
	tokenAddress := common.HexToAddress(WETH)

	// we are calling the transfer function on the token contract
	transferFnSignature := []byte("transfer(address,uint256)")
	// generate the keccak256 hash of the signature
	hash := sha3.NewLegacyKeccak256()
	hash.Write(transferFnSignature)
	// the first 4 bytes of the hash is the method id
	methodID := hash.Sum(nil)[:4]
	fmt.Println(hexutil.Encode(methodID)) // 0xa9059cbb

	// need to zero-pad to the left, byte slice must be 32 bytes long
	paddedAddress := common.LeftPadBytes(toAddress.Bytes(), 32)
	fmt.Println(hexutil.Encode(paddedAddress)) // 0x0000000000000000000000004592d8f8d7b001e72cb26a73e4fa1806a51ac79d

	amount := new(big.Int)
	amount.SetString("10000000000000000", 10) // sets the value to 0.01 tokens, in the token denomination

	paddedAmount := common.LeftPadBytes(amount.Bytes(), 32)
	fmt.Println(hexutil.Encode(paddedAmount)) // 0x00000000000000000000000000000000000000000000003635c9adc5dea00000

	// collect everything into a byte slice
	var data []byte
	data = append(data, methodID...)
	data = append(data, paddedAddress...)
	data = append(data, paddedAmount...)

	// estimate the gas needed for the tx
	gasLimit, err := client.EstimateGas(context.Background(), ethereum.CallMsg{
		To:   &tokenAddress,
		Data: data,
	})
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(gasLimit) // 23256

	// value = 0, we are NOT transferring ETH
	tx := types.NewTransaction(nonce, tokenAddress, value, gasLimit, gasPrice, data)

	chainID, err := client.NetworkID(context.Background())
	if err != nil {
		log.Fatal(err)
	}

	signedTx, err := types.SignTx(tx, types.NewEIP155Signer(chainID), privateKey)
	if err != nil {
		log.Fatal(err)
	}

	err = client.SendTransaction(context.Background(), signedTx)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("tx sent: %s\n", signedTx.Hash().Hex())
}
