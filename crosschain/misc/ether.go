package misc

import (
	"context"
	"math/big"

	"crosschain/github.com/ethereum/go-ethereum/ethclient"
)

type Network struct {
	name, address string
}

func fetchChain(chain string) int {
	/* Dail to chain */
	for {
		client, err := ethclient.Dial(chain)
		if err != nil {
			return -1
		}

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
			return len(block.Transactions())
		}
	}
}
