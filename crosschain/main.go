package main

import (
	"fmt"
	"os"
	"os/exec"
	"sync"

	ui "crosschain/interface"
	"crosschain/networks"
	"crosschain/utils"
)

func main() {
	uiChoice := os.Args[1]
	if uiChoice == "tui" {
		cmd := exec.Command("tput", "civis")
		cmd.Stdout = os.Stdout
		cmd.Run()
		fmt.Printf("%s", ui.CLEAR)
	}

	uniswapMarkets := utils.UniswapV2Markets{}

	uniswapMarkets.Setup()

	wg := new(sync.WaitGroup)
	wg.Add(1)

	i := 0

	go networks.Network(&uniswapMarkets, &i, uiChoice, "ethereum")
	go networks.Network(&uniswapMarkets, &i, uiChoice, "polygon")
	go networks.Network(&uniswapMarkets, &i, uiChoice, "avalanche")
	go networks.Network(&uniswapMarkets, &i, uiChoice, "bsc")
	go networks.Network(&uniswapMarkets, &i, uiChoice, "aurora")
	go networks.Network(&uniswapMarkets, &i, uiChoice, "fantom")

	wg.Wait()

}

// abigen --abi ./builds/token.abi --pkg generatedContracts --type Token --out ./generatedContracts/token.go --bin ./builds/token.bin
