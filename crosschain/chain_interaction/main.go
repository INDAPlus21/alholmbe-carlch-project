package main

import (
	"sync"

	"chain_interaction/networks"
	"chain_interaction/utils"
)

// bsc
const PANCAKESWAP_FACTORY_ADDRESS_BSC string = "0xcA143Ce32Fe78f1f7019d7d551a6402fC5350c73"
const SUSHISWAP_FACTORY_ADDRESS_BSC string = "0xc35DADB65012eC5796536bD9864eD8773aBc74C4"
const WBNB_ADDRESS string = "0xbb4CdB9CBd36B01bD1cBaEBF2De08d9173bc095c"

func main() {

	// setup
	ethereum_WETH := utils.Network{Asset: "WETH", Protocol: "ethereum"}
	ethereum_WBNB := utils.Network{Asset: "WBNB", Protocol: "ethereum"}
	ethereum_WMATIC := utils.Network{Asset: "WMATIC", Protocol: "ethereum"}

	bsc_WETH := utils.Network{Asset: "WETH", Protocol: "bsc"}
	bsc_WBNB := utils.Network{Asset: "WBNB", Protocol: "bsc"}
	bsc_WMATIC := utils.Network{Asset: "WMATIC", Protocol: "bsc"}

	polygon_WETH := utils.Network{Asset: "WETH", Protocol: "polygon"}
	polygon_WBNB := utils.Network{Asset: "WBNB", Protocol: "polygon"}
	polygon_WMATIC := utils.Network{Asset: "WMATIC", Protocol: "polygon"}

	uniswapMarkets := utils.UniswapV2Markets{}
	uniswapMarkets.Asset = make(map[string]map[string]*utils.Network)
	uniswapMarkets.Asset["WETH"] = make(map[string]*utils.Network)
	uniswapMarkets.Asset["WETH"]["ethereum"] = &ethereum_WETH
	uniswapMarkets.Asset["WETH"]["bsc"] = &bsc_WETH
	uniswapMarkets.Asset["WETH"]["polygon"] = &polygon_WETH

	uniswapMarkets.Asset["WBNB"] = make(map[string]*utils.Network)
	uniswapMarkets.Asset["WBNB"]["ethereum"] = &ethereum_WBNB
	uniswapMarkets.Asset["WBNB"]["bsc"] = &bsc_WBNB
	uniswapMarkets.Asset["WBNB"]["polygon"] = &polygon_WBNB

	uniswapMarkets.Asset["WMATIC"] = make(map[string]*utils.Network)
	uniswapMarkets.Asset["WMATIC"]["ethereum"] = &ethereum_WMATIC
	uniswapMarkets.Asset["WMATIC"]["bsc"] = &bsc_WMATIC
	uniswapMarkets.Asset["WMATIC"]["polygon"] = &polygon_WMATIC

	ch1 := make(chan map[string][]utils.UniswapV2EthPair)
	wg := new(sync.WaitGroup)
	wg.Add(1)

	go networks.Binance(&uniswapMarkets, ch1, wg)

	// for ethPairs := range ch1 {
	// 	for k, v := range ethPairs {
	// 		fmt.Printf("key token: %v, pair: %v \n", k, v[0].PairAddress.String())
	// 	}
	// }

	wg.Wait()

}

// abigen --abi ./builds/token.abi --pkg generatedContracts --type Token --out ./generatedContracts/token.go --bin ./builds/token.bin
