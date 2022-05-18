package networks

import (
	"fmt"
	"math/big"
	"time"

	"chain_interaction/utils"
)

const SUSHISWAP_FACTORY_ADDRESS_BSC string = "0xc35DADB65012eC5796536bD9864eD8773aBc74C4"
const PANCAKESWAP_FACTORY_ADDRESS_BSC string = "0xca143ce32fe78f1f7019d7d551a6402fc5350c73"
const WBNB_ADDRESS_BSC string = "0xbb4cdb9cbd36b01bd1cbaebf2de08d9173bc095c"
const UNISWAP_QUERY_ADDRESS_BSC string = "0xBc37182dA7E1f99f5Bd75196736BB2ae804Cbf6A"

func Binance(uniswapMarkets *utils.UniswapV2Markets) {
	bscFactories := []string{PANCAKESWAP_FACTORY_ADDRESS_BSC, SUSHISWAP_FACTORY_ADDRESS_BSC}

	client := utils.GetClient("bsc")

	var minLiq, power = big.NewInt(10), big.NewInt(18)
	minLiq.Exp(minLiq, power, nil)
	// the tokens we care about on this network
	tokens := []utils.Token{
		utils.Token{
			Symbol:       "WBNB",
			Address:      "0xbb4CdB9CBd36B01bD1cBaEBF2De08d9173bc095c",
			Protocol:     "bsc",
			MinLiquidity: minLiq, // 1 WBNB
		},
		utils.Token{
			Symbol:       "WETH",
			Address:      "0x2170Ed0880ac9A755fd29B2688956BD959F933F8",
			Protocol:     "bsc",
			MinLiquidity: minLiq, // 1 WETH
		},
	}

	uniswapMarkets.UpdateMarkets(
		client, bscFactories, UNISWAP_QUERY_ADDRESS_BSC, tokens,
	)

	fmt.Printf("all markets on bsc: %d\n", len(uniswapMarkets.Asset["WBNB"]["bsc"].AllMarkets))

	uniswapMarkets.UpdateReserves(client, UNISWAP_QUERY_ADDRESS_BSC, tokens)
	fmt.Println("initial reserve update on binance.")

	uniswapMarkets.EvaluateCrossMarkets(tokens)

	for {
		uniswapMarkets.UpdateReserves(client, UNISWAP_QUERY_ADDRESS_BSC, tokens)
		for _, token := range tokens {
			uniswapMarkets.UpdateScreen(token.Symbol, token.Protocol)
		}
		time.Sleep(10 * time.Second)
	}

}
