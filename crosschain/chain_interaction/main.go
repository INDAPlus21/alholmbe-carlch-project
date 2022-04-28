package main

import (
	"fmt"
	"log"
	"os"

	"chain_interaction/utils"

	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/joho/godotenv"
)

// ethereum
const UNISWAP_FACTORY_ADDRESS string = "0x5C69bEe701ef814a2B6a3EDD4B1652CB9cc5aA6f"
const SUSHISWAP_FACTORY_ADDRESS string = "0xC0AEe478e3658e2610c5F7A4A2E1777cE9e4f2Ac"
const WETH_ADDRESS string = "0xC02aaA39b223FE8D0A0e5C4F27eAD9083C756Cc2"

// ropsten
const SUSHISWAP_FACTORY_ADDRESS_ROPSTEN string = "0xc35DADB65012eC5796536bD9864eD8773aBc74C4"
const UNISWAP_FACTORY_ADDRESS_ROPSTEN string = "0x5C69bEe701ef814a2B6a3EDD4B1652CB9cc5aA6f"
const WETH_ADDRESS_ROPSTEN string = "0xC02aaA39b223FE8D0A0e5C4F27eAD9083C756Cc2"

// bsc
const PANCAKESWAP_FACTORY_ADDRESS_BSC string = "0xcA143Ce32Fe78f1f7019d7d551a6402fC5350c73"
const SUSHISWAP_FACTORY_ADDRESS_BSC string = "0xc35DADB65012eC5796536bD9864eD8773aBc74C4"
const WBNB_ADDRESS string = "0xbb4CdB9CBd36B01bD1cBaEBF2De08d9173bc095c"

func main() {
	// bscFactories := []string{PANCAKESWAP_FACTORY_ADDRESS_BSC, SUSHISWAP_FACTORY_ADDRESS_BSC}
	ethFactories := []string{UNISWAP_FACTORY_ADDRESS_ROPSTEN}

	// get a provider
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	rpc := os.Getenv("rpc_ropsten")

	client, err := ethclient.Dial(rpc)
	if err != nil {
		log.Fatal(err)
	}

	markets := utils.UniswapV2Markets(client, ethFactories)

	// get all markets
	for k, v := range markets {
		fmt.Printf("key token: %v, pair: %v \n", k, v[0].PairAddress.String())
	}

	// communicate the markets back to the main goroutine

	// set up listener for new block

	// for every block
	// update reserves
	// evaluate markets
	// print the opportunity
	// execute

}

// abigen --abi ./builds/token.abi --pkg generatedContracts --type Token --out ./generatedContracts/token.go --bin ./builds/token.bin

// market := map[string](map[string]string){
// 	"0xb4e16d0168e52d35cacd2c6185b44281ec28c9dc": map[string]string{
// 		"description": "WETH/USDC",
// 		"network":     "ethereum",
// 		"dex":         "uniswap",
// 	},
// 	"b": map[string]string{
// 		"d": "d",
// 	},
// }
