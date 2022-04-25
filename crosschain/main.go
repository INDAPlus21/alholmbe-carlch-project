package main

import (
	"fmt"
	"os"
	"os/exec"
	// "os"
	// "bufio"
	// "encoding/json"
	// "github.com/ethereum/go-ethereum/common"
)

type Pair struct {
	Market string   `json:"_marketAddress"`
	Tokens []string `json:"_tokens"`
}

const clearLine = "\033[H\033[2J"

const marketPairs = "src/randomjson/bsc_allMarketPairs.json"

func main() {
	cmd := exec.Command("tput", "civis")
	cmd.Stdout = os.Stdout
	cmd.Run()
	chains := []Network{
		{"Fantom", "https://rpc3.fantom.network"},
		{"Binance", "https://bsc-dataseed.binance.org/"},
		{"Ethereum", "https://cloudflare-eth.com"},
		{"Polygon", "https://polygon-rpc.com/"},
		{"Harmony", "https://rpc.heavenswail.one"},
		{"Aurora", "https://mainnet.aurora.dev"},
		{"Avalanche", "https://api.avax.network/ext/bc/C/rpc"},
	}
	/*
	  file_json, err := os.Open(marketPairs)
	  if err != nil {
	    log.Fatal(err)
	  }
	  fmt.Printf("Successfully opened file\n")

	  r := bufio.NewReader(file_json)
	  d := json.NewDecoder(r)
	  i := 0

	  d.Token()
	  for d.More() {
	    pair := &Pair{}
	    d.Decode(pair)
	    fmt.Printf("%s\n", pair.Tokens[0])
	    i++
	  }
	  fmt.Printf("Successfully read %d objects\n", i)

	  defer file_json.Close()
	*/
	fmt.Printf("\033[H\033[2JNetwork\t\tTransactions")

	printTable(chains, 2)
	for {
	}

}
