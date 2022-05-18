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
	wg.Add(3)

	go networks.Binance(&uniswapMarkets, wg)
	go networks.Polygon(&uniswapMarkets, wg)

	wg.Wait()

}

// abigen --abi ./builds/token.abi --pkg generatedContracts --type Token --out ./generatedContracts/token.go --bin ./builds/token.bin
