package networks

import (
	"log"
	"os"
	"sync"

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

func Ethereum(ch chan map[string][]utils.UniswapV2EthPair, wg *sync.WaitGroup) {
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

	// get all eth markets
	markets := utils.UniswapV2Markets(client, ethFactories)
	ch <- markets
	wg.Done()

	// communicate the markets back to the main goroutine

	// evaluate for atomic arbs

	// set up listener for new block

	// for every block
	// update reserves
	// communicate the markets back to the main goroutine
	// evaluate for atomic arbs
	// if found, try to execute

}
