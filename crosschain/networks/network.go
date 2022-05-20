package networks

import (
	// "fmt"

	"time"

	"crosschain/utils"
)

func Network(uniswapMarkets *utils.UniswapV2Markets, i *int, uiChoice string, network string) {
	factories := utils.GetFactories(network)

	client := utils.GetClient(network)

	tokens := utils.GetTokens(network)

	uniswapQueryAddress := utils.DeployedUniswapQueryContracts[network]

	uniswapMarkets.UpdateMarkets(
		client, factories, uniswapQueryAddress, tokens,
	)

	uniswapMarkets.UpdateReserves(client, uniswapQueryAddress, tokens)

	uniswapMarkets.EvaluateCrossMarkets(tokens)
	if uiChoice == "tui" {
		y := *i
		*i++
		for {
			uniswapMarkets.UpdateReserves(client, uniswapQueryAddress, tokens)
			uniswapMarkets.EvaluateCrossMarkets(tokens)
			uniswapMarkets.UpdateScreen(tokens[0].Symbol, tokens[0].Protocol, y)
			time.Sleep(2 * time.Second)
		}
	} else {
		for {
			uniswapMarkets.UpdateReserves(client, uniswapQueryAddress, tokens)
			uniswapMarkets.EvaluateCrossMarkets(tokens)
			uniswapMarkets.PrintOpportunities(tokens[0].Symbol, tokens[0].Protocol)
			time.Sleep(2 * time.Second)
		}
	}

}
