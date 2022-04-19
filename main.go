package main

import (

  "fmt"
  "log"
  // "os"
  // "bufio"
  "context"

  "math/big"
  // "encoding/json"

  // "github.com/ethereum/go-ethereum/common"
  "github.com/ethereum/go-ethereum/ethclient"
)

type Pair struct {
  Market    string    `json:"_marketAddress"`
  Tokens    []string  `json:"_tokens"`
}

const marketPairs = "src/randomjson/bsc_allMarketPairs.json"

func main() {
  chains := [7]string{
    "https://rpc3.fantom.network",
    "https://bsc-dataseed.binance.org/",
    "https://cloudflare-eth.com",
    "https://polygon-rpc.com/",
    "https://rpc.heavenswail.one",
    "https://mainnet.aurora.dev",
    "https://api.avax.network/ext/bc/C/rpc",
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
  for i := 0; i < 7; i++ {
    go fetchAPI(chains[i])
  }
  for { }
}

func fetchAPI(chain string) {
  /* Dail to chain */
  for {
    client, err := ethclient.Dial(chain)
    if err != nil { log.Fatal(err) }
    fmt.Printf("Successfully connected to %s\n", chain)

    /* Do stuff */
    blockNumber := big.NewInt(0)

    for {
      header, err := client.HeaderByNumber(context.Background(), nil)
      if err != nil { break }

      if blockNumber.Cmp(header.Number) == 0 { continue }

      blockNumber = header.Number
      block, err := client.BlockByNumber(context.Background(), blockNumber)
      if err != nil { break }

      fmt.Printf("%s : %d\t : %s\n", blockNumber, len(block.Transactions()), chain)
    }
  }
}
