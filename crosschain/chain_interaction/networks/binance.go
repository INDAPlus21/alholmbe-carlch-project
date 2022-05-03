package networks

import (
	"fmt"
	"log"
	"os"
	"sync"

	"chain_interaction/utils"

	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/joho/godotenv"
)

const SUSHISWAP_FACTORY_ADDRESS_BSC string = "0xc35DADB65012eC5796536bD9864eD8773aBc74C4"
const PANCAKESWAP_FACTORY_ADDRESS_BSC string = "0xca143ce32fe78f1f7019d7d551a6402fc5350c73"
const WBNB_ADDRESS_BSC string = "0xbb4cdb9cbd36b01bd1cbaebf2de08d9173bc095c"
const UNISWAP_QUERY_ADDRESS_BSC string = "0xBc37182dA7E1f99f5Bd75196736BB2ae804Cbf6A"

func Binance(ch chan map[string][]utils.UniswapV2EthPair, wg *sync.WaitGroup) {
	bscFactories := []string{PANCAKESWAP_FACTORY_ADDRESS_BSC, SUSHISWAP_FACTORY_ADDRESS_BSC}

	// get a provider
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	rpc := os.Getenv("rpc_bsc")

	client, err := ethclient.Dial(rpc)
	if err != nil {
		log.Fatal(err)
	}

	// get all eth markets
	allMarkets, crossMarkets, _, _ := utils.UniswapV2Markets(
		client, bscFactories, UNISWAP_QUERY_ADDRESS_BSC, WBNB_ADDRESS_BSC,
	)
	// communicate the markets back to the main goroutine
	// ch <- allMarkets

	fmt.Printf("allMarkets: %d\n", len(allMarkets))
	fmt.Printf("crossMarkets: %d\n", len(crossMarkets))

	// evaluate for atomic arbs
	utils.UpdateReserves(client, &crossMarkets, UNISWAP_QUERY_ADDRESS_BSC)

	fmt.Println(crossMarkets)

	wg.Done()
	// crossMarkets := []

	// set up listener for new block

	// for every block
	// update reserves
	// communicate the markets back to the main goroutine
	// evaluate for atomic arbs
	// if found, try to execute

}
