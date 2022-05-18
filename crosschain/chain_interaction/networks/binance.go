package networks

import (
	//"fmt"
	"math/big"
	"sync"
	"time"

	"chain_interaction/utils"
)

const SUSHISWAP_FACTORY_ADDRESS_BSC string = "0xc35DADB65012eC5796536bD9864eD8773aBc74C4"
const PANCAKESWAP_FACTORY_ADDRESS_BSC string = "0xca143ce32fe78f1f7019d7d551a6402fc5350c73"
const WBNB_ADDRESS_BSC string = "0xbb4cdb9cbd36b01bd1cbaebf2de08d9173bc095c"
const UNISWAP_QUERY_ADDRESS_BSC string = "0xBc37182dA7E1f99f5Bd75196736BB2ae804Cbf6A"

func Binance(uniswapMarkets *utils.UniswapV2Markets, i *int, wg *sync.WaitGroup) {
	bscFactories := []string{PANCAKESWAP_FACTORY_ADDRESS_BSC, SUSHISWAP_FACTORY_ADDRESS_BSC}

	client := utils.GetClient("bsc")

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
		utils.Token{
			Symbol:       "WETH",
			Address:      "0x2170Ed0880ac9A755fd29B2688956BD959F933F8",
			Protocol:     "bsc",
			MinLiquidity: base, // 1 WETH
		},
	}

	uniswapMarkets.UpdateMarkets(
		client, bscFactories, UNISWAP_QUERY_ADDRESS_BSC, tokens,
	)

	// fmt.Printf("all markets on bsc: %d\n", len(uniswapMarkets.Asset["WBNB"]["bsc"].AllMarkets))

	uniswapMarkets.UpdateReserves(client, UNISWAP_QUERY_ADDRESS_BSC, tokens)
	// fmt.Println("initial reserve update on binance.")

	uniswapMarkets.EvaluateCrossMarkets(tokens)
  y := *i
  *i += 2
	for {
		uniswapMarkets.UpdateReserves(client, UNISWAP_QUERY_ADDRESS_BSC, tokens)
		for _, token := range tokens {
		  uniswapMarkets.UpdateScreen(token.Symbol, token.Protocol, y)
      y++
		}
    y -= 2
		time.Sleep(2 * time.Second)
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
