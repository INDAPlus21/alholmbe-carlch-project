package main

import (
	"context"
	"fmt"
	"log"
	"math/big"
	"sync"

	"github.com/ethereum/go-ethereum/ethclient"
)

type Pair struct {
	Market string   `json:"_marketAddress"`
	Tokens []string `json:"_tokens"`
}

// const ethereum = "https://cloudflare-eth.com"
const ethereum_archival = "https://eth-mainnet.alchemyapi.io/v2/u-P4WTWXjk4duuq6e4VO8aygIYiSJl4H"
const avalanche = "https://rpc.ankr.com/avalanche"
const bsc = "https://bsc-dataseed.binance.org/"
const polygon = "https://polygon-rpc.com/"
const aurora = "https://mainnet.aurora.dev"

type Market struct {
	protocol      string         // ex "ethereum"
	marketAddress string         // ex "0xB4e16d0168e52d35CaCD2c6185b44281Ec28C9Dc"
	tokenBalances map[string]int // maps address to balance
}

func main() {
	chains := [5]string{
		ethereum_archival,
		avalanche,
		bsc,
		polygon,
		aurora,
	}

	wg := new(sync.WaitGroup)
	wg.Add(len(chains))

	// this should run until we terminate
	for i := 0; i < 7; i++ {
		go fetchAPI(chains[i], wg)
	}

	wg.Wait()
}

func fetchAPI(chain string, wg *sync.WaitGroup) {

	/* Dial to chain */
	// get markets

	// listener that waits on the next block and updates the reserves

	for {
		client, err := ethclient.Dial(chain)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Printf("Successfully connected to %s\n", chain)

		/* Do stuff */
		blockNumber := big.NewInt(0)

		for {
			header, err := client.HeaderByNumber(context.Background(), nil)
			if err != nil {
				break
			}

			if blockNumber.Cmp(header.Number) == 0 {
				continue
			}

			blockNumber = header.Number
			block, err := client.BlockByNumber(context.Background(), blockNumber)
			if err != nil {
				break
			}

			fmt.Printf("%s\t : %d\t : %s\n", blockNumber, len(block.Transactions()), chain)
		}
	}

	wg.Done()
}
