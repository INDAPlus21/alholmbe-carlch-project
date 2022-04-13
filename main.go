package main

import (

  "fmt"
  "log"
  "os"
  "bufio"
  "context"

  "encoding/json"

  // "github.com/ethereum/go-ethereum/common"
  "github.com/ethereum/go-ethereum/ethclient"
)

type Pair struct {
  Market    string    `json:"_marketAddress"`
  Tokens    []string  `json:"_tokens"`
}

const market_pairs = "src/randomjson/bsc_allMarketPairs.json"
const network = "https://bsc-dataseed.binance.org/"

func main() {

  file_json, err := os.Open(market_pairs)
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

  client, err := ethclient.Dial(network)
  if err != nil {
    log.Fatal(err)
  }
  fmt.Printf("Successfully connected to %s\n", network)

  header, err := client.HeaderByNumber(context.Background(), nil)
  if err != nil {
    log.Fatal(err)
  }
  fmt.Println(header.Number.String())
}
