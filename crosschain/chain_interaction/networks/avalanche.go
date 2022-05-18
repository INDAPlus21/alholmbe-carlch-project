package networks

import (
	"fmt"
	"math/big"
	"time"

	"chain_interaction/utils"
)

const PANGOLIN_FACTORY_ADDRESS_AVALANCHE string = "0xefa94de7a4656d787667c749f7e1223d71e9fd88"
const TRADERJOE_FACTORY_ADDRESS_AVALANCHE string = "0x9ad6c38be94206ca50bb0d90783181662f0cfa10"
const WAVAX_ADDRESS_AVALANCHE string = "0xB31f66AA3C1e785363F0875A1B74E27b85FD66c7"
const UNISWAP_QUERY_ADDRESS_AVALANCHE string = "0xbc37182da7e1f99f5bd75196736bb2ae804cbf6a"

func Avalanche(uniswapMarkets *utils.UniswapV2Markets) {
	avalancheFactories := []string{TRADERJOE_FACTORY_ADDRESS_AVALANCHE, PANGOLIN_FACTORY_ADDRESS_AVALANCHE}

	client := utils.GetClient("avalanche")

	var minLiq, power = big.NewInt(10), big.NewInt(18)
	minLiq.Exp(minLiq, power, nil)
	// the tokens we care about on this network
	tokens := []utils.Token{
		utils.Token{
			Symbol:       "WAVAX",
			Address:      WAVAX_ADDRESS_AVALANCHE,
			Protocol:     "avalanche",
			MinLiquidity: minLiq, // 1 WAVAX
		},
	}

	uniswapMarkets.UpdateMarkets(
		client, avalancheFactories, UNISWAP_QUERY_ADDRESS_AVALANCHE, tokens,
	)

	fmt.Printf("all markets on avalanche: %d\n", len(uniswapMarkets.Asset["WAVAX"]["avalanche"].AllMarkets))

	uniswapMarkets.UpdateReserves(client, UNISWAP_QUERY_ADDRESS_AVALANCHE, tokens)
	fmt.Println("initial reserve update on avalanche.")

	uniswapMarkets.EvaluateCrossMarkets(tokens)

	for {
		uniswapMarkets.UpdateReserves(client, UNISWAP_QUERY_ADDRESS_AVALANCHE, tokens)
		uniswapMarkets.EvaluateCrossMarkets(tokens)
		for _, token := range tokens {
			uniswapMarkets.PrintOpportunities(token.Symbol, token.Protocol)
		}
		time.Sleep(10 * time.Second)
	}

}
