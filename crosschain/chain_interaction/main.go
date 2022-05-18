package main

import (
	"sync"

	// ui "chain_interaction/interface"
	"chain_interaction/networks"
	"chain_interaction/utils"
)

func main() {

	uniswapMarkets := utils.UniswapV2Markets{}

	uniswapMarkets.Setup()

	wg := new(sync.WaitGroup)
	wg.Add(1)

	go networks.Binance(&uniswapMarkets)
	go networks.Polygon(&uniswapMarkets)
	go networks.Avalanche(&uniswapMarkets)

	wg.Wait()

}

// abigen --abi ./builds/token.abi --pkg generatedContracts --type Token --out ./generatedContracts/token.go --bin ./builds/token.bin
