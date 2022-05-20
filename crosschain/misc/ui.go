package misc

import (
	"fmt"
)

const clr = "\033[H\033[2J"
const clrln = "\033[2K"

const (
	White  string = "\033[0m"
	Red           = "\033[31m"
	Green         = "\033[32m"
	Yellow        = "\033[33m"
)

func printTable(chains []Network, offset int) {
	for i := 0; i < len(chains); i++ {
		go func(i int) {
			for {
				updateTable(chains[i], (i + offset))
			}
		}(i)
	}
}

func updateTable(chain Network, offset int) {
	transactions := fetchChain(chain.address)
	if transactions == -1 {
		fmt.Printf("\033[%dH\033[2K%s%s%s\033[%d;17H%d", (1 + offset), Red, chain.name, White, (1 + offset), transactions)
		return
	}
	fmt.Printf("\033[%dH\033[2K%s%s%s\033[%d;17H%d", (1 + offset), Green, chain.name, White, (1 + offset), transactions)
}
