package networks

import (
	// "fmt"
	"math/big"
	"sync"
	"time"

	"chain_interaction/utils"
)

const SUSHISWAP_FACTORY_ADDRESS_POLYGON string = "0xc35DADB65012eC5796536bD9864eD8773aBc74C4"
const QUICKSWAP_FACTORY_ADDRESS_POLYGON string = "0x5757371414417b8C6CAad45bAeF941aBc7d3Ab32"
const UNISWAP_QUERY_ADDRESS_POLYGON string = "0xBc37182dA7E1f99f5Bd75196736BB2ae804Cbf6A"

func Polygon(uniswapMarkets *utils.UniswapV2Markets, i *int, wg *sync.WaitGroup) {
	polygonFactories := []string{QUICKSWAP_FACTORY_ADDRESS_POLYGON, SUSHISWAP_FACTORY_ADDRESS_POLYGON}

	client := utils.GetClient("polygon")

	// the tokens we care about on this network
	var base, power = big.NewInt(10), big.NewInt(18)
	base.Exp(base, power, nil)
	tokens := []utils.Token{
		utils.Token{
			Symbol:       "WMATIC",
			Address:      "0x0d500B1d8E8eF31E21C99d1Db9A6444d3ADf1270",
			Protocol:     "polygon",
			MinLiquidity: base, // 1 WMATIC
		},
	}

	uniswapMarkets.UpdateMarkets(
		client, polygonFactories, UNISWAP_QUERY_ADDRESS_POLYGON, tokens,
	)

	// fmt.Printf("all markets on polygon: %d\n", len(uniswapMarkets.Asset["WMATIC"]["polygon"].AllMarkets))

	uniswapMarkets.UpdateReserves(client, UNISWAP_QUERY_ADDRESS_POLYGON, tokens)
	// fmt.Println("initial reserve update on polygon.")

	uniswapMarkets.EvaluateCrossMarkets(tokens)
  y := *i
  *i++
	for {
		uniswapMarkets.UpdateReserves(client, UNISWAP_QUERY_ADDRESS_POLYGON, tokens)
		uniswapMarkets.UpdateScreen("WMATIC", "polygon", y)
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
