package main

import (
	"sync"
  "os"
  "os/exec"
  "fmt"

  "chain_interaction/interface"
	"chain_interaction/networks"
	"chain_interaction/utils"
)

// bsc
const PANCAKESWAP_FACTORY_ADDRESS_BSC string = "0xcA143Ce32Fe78f1f7019d7d551a6402fC5350c73"
const SUSHISWAP_FACTORY_ADDRESS_BSC string = "0xc35DADB65012eC5796536bD9864eD8773aBc74C4"
const WBNB_ADDRESS string = "0xbb4CdB9CBd36B01bD1cBaEBF2De08d9173bc095c"

func main() {
  cmd := exec.Command("tput", "civis")
  cmd.Stdout = os.Stdout
  cmd.Run()

  fmt.Printf("%s",ui.CLEAR)

	uniswapMarkets := utils.UniswapV2Markets{}

	uniswapMarkets.Setup()

	wg := new(sync.WaitGroup)
	wg.Add(2)

  i := 0;
	go networks.Binance(&uniswapMarkets, &i, wg)
	go networks.Polygon(&uniswapMarkets, &i, wg)
	go networks.Avalanche(&uniswapMarkets, &i, wg)
	// go ui.Update_screen(&uniswapMarkets, e, wg)

	// for ethPairs := range ch1 {
	// 	for k, v := range ethPairs {
	// 		fmt.Printf("key token: %v, pair: %v \n", k, v[0].PairAddress.String())
	// 	}
	// }

	wg.Wait()

}

// abigen --abi ./builds/token.abi --pkg generatedContracts --type Token --out ./generatedContracts/token.go --bin ./builds/token.bin
