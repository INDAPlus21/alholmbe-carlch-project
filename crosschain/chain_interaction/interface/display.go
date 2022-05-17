package ui

import (
  "fmt"
  "sync"
  "math/big"

  "chain_interaction/utils"
)


func update_asset(asset map[string]*utils.Network) {
  a := new(big.Float)
  a.SetFloat64(0)
  for _,t := range(asset) {
    for _,market := range(t.CrossMarketsByToken) {
      for i := 0; i < len(market.Pairs); i++ {
        if utils.FetchType(market.Pairs[i].Token0Address.String()) != nil &&
           utils.FetchType(market.Pairs[i].Token1Address.String()) != nil { // &&
           //market.CurrentArbitrageOpp.Cmp(a)>1 {
          fmt.Printf("%s %s %s%v%s\n",
            *utils.FetchType(market.Pairs[i].Token0Address.String()),
            *utils.FetchType(market.Pairs[i].Token1Address.String()),
            RED,
            *market.CurrentArbitrageOpp,
            WHITE)
        }
      }
    }
  }
}

var a big.Float
func Update_screen(markets utils.UniswapV2Markets, wg *sync.WaitGroup) {
  utils.InitSets()

  // time.Sleep(10*time.Second)

  // var pairs map[string]string
  // pairs = make(map[string]string)
  for {
    for _,t1 := range(markets.Asset) {
      go update_asset(t1)
    }
  }

  wg.Done()
  /*

  for {
    for _,token1 := range(markets.Asset) {
      for _,token2 := range(token1) {
        for _,market := range(token2.CrossMarketsByToken) {
          if market.CurrentArbitrageOpp.Cmp(a) < 1 {
            if utils.FetchType(market.Pairs[0].Token0Address.String()) != nil &&
              utils.FetchType(market.Pairs[0].Token1Address.String()) != nil {
              fmt.Printf("%s %s %s%f%s\n",
                *utils.FetchType(market.Pairs[0].Token0Address.String()),
                *utils.FetchType(market.Pairs[0].Token1Address.String()),
                RED,
                market.CurrentArbitrageOpp,
                WHITE)
              }
          }
        }
      }
    }
  }
  */
}

