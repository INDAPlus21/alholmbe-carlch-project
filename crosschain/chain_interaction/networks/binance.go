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

func Binance(uniswapMarkets *utils.UniswapV2Markets, ch chan map[string][]utils.UniswapV2EthPair, wg *sync.WaitGroup) {
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

	tokens := []utils.Token{
		utils.Token{"WBNB", "0xbb4CdB9CBd36B01bD1cBaEBF2De08d9173bc095c", "bsc"},
		// utils.Token{"WETH", "0x2170Ed0880ac9A755fd29B2688956BD959F933F8", "bsc"},
	}

	// get all markets
	// instead of returning markets, mutate the market object itself
	uniswapMarkets.UpdateMarkets(
		client, bscFactories, UNISWAP_QUERY_ADDRESS_BSC, tokens,
	)

	fmt.Printf("allMarkets: %d\n", len(uniswapMarkets.Asset["WBNB"]["bsc"].AllMarkets))
	fmt.Printf("crossMarkets: %d\n", len(uniswapMarkets.Asset["WBNB"]["bsc"].CrossMarkets))

	uniswapMarkets.UpdateReserves(client, UNISWAP_QUERY_ADDRESS_BSC, tokens)

	// for _, market := range uniswapMarkets.Asset["WBNB"]["bsc"].AllMarkets {
	// 	fmt.Println(*market)
	// }

	// // evaluate for atomic arbs
	uniswapMarkets.EvaluateCrossMarkets()

	for tokenAddress, market := range uniswapMarkets.Asset["WBNB"]["bsc"].CrossMarketsByToken {
		fmt.Printf("%s: %f\n", tokenAddress, market.CurrentArbitrageOpp)
	}

	wg.Done()
	// crossMarkets := []

	// set up listener for new block

	// for every block
	// update reserves
	// communicate the markets back to the main goroutine
	// evaluate for atomic arbs
	// if found, try to execute

}
