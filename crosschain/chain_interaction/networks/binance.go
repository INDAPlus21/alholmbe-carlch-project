package networks

import (
	"fmt"
	"log"
	"math/big"
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

func Binance(uniswapMarkets *utils.UniswapV2Markets, wg *sync.WaitGroup) {
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

	// the tokens we care about on this network
	var base, power = big.NewInt(10), big.NewInt(18)
	base.Exp(base, power, nil)
	tokens := []utils.Token{
		utils.Token{
			Symbol:       "WBNB",
			Address:      "0xbb4CdB9CBd36B01bD1cBaEBF2De08d9173bc095c",
			Protocol:     "bsc",
			MinLiquidity: base, // 100 WBNB
		},
		// utils.Token{"WETH", "0x2170Ed0880ac9A755fd29B2688956BD959F933F8", "bsc"},
	}

	uniswapMarkets.UpdateMarkets(
		client, bscFactories, UNISWAP_QUERY_ADDRESS_BSC, tokens,
	)

	fmt.Printf("allMarkets: %d\n", len(uniswapMarkets.Asset["WBNB"]["bsc"].AllMarkets))
	fmt.Printf("crossMarkets: %d\n", len(uniswapMarkets.Asset["WBNB"]["bsc"].CrossMarkets))

	uniswapMarkets.UpdateReserves(client, UNISWAP_QUERY_ADDRESS_BSC, tokens)

	// evaluate for atomic arbs
	uniswapMarkets.EvaluateCrossMarkets(tokens)

	for tokenAddress, market := range uniswapMarkets.Asset["WBNB"]["bsc"].CrossMarketsByToken {
		if market.CurrentArbitrageOpp.Cmp(big.NewFloat(0)) == 1 {
			fmt.Printf("%s: %f\n", tokenAddress, market.CurrentArbitrageOpp)
		}
	}

	// for i := 0; i < 50; i++ {
	// 	uniswapMarkets.UpdateReserves(client, UNISWAP_QUERY_ADDRESS_BSC, tokens)
	// 	for tokenAddress, market := range uniswapMarkets.Asset["WBNB"]["bsc"].CrossMarketsByToken {
	// 		if market.CurrentArbitrageOpp.Cmp(big.NewFloat(0)) == 1 {
	// 			fmt.Printf("%s: %f\n", tokenAddress, market.CurrentArbitrageOpp)
	// 		}
	// 	}
	// 	fmt.Println()
	// 	time.Sleep(30 * time.Second)
	// }

	wg.Done()
	// crossMarkets := []

	// set up listener for new block

	// for every block
	// update reserves
	// communicate the markets back to the main goroutine
	// evaluate for atomic arbs
	// if found, try to execute

}
