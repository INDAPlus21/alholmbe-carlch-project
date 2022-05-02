package main

import (
  "fmt"
)

/*
 * Datastruktur för Token
 * Kontrakt av en token typ ska ge alla andra tokens av samma typ
 */

const WBNB_ETHEREUM = "0xB8c77482e45F1F44dE1745F52C74426C631bDD52"
const WBNB_BINANCE  = "0xbb4CdB9CBd36B01bD1cBaEBF2De08d9173bc095c"
const WBNB_POLYGON  = "0x3BA4c387f786bFEE076A58914F5Bd38d668B42c3"

const WETH_ETHEREUM = "0xc02aaa39b223fe8d0a0e5c4f27ead9083c756cc2"
const WETH_BINANCE  = "0x2170ed0880ac9a755fd29b2688956bd959f933f8"
const WETH_FANTOM   = "0x658b0c7613e890ee50b8c4bc6a3f41ef411208ad"

const USDT_ETHEREUM = "0xdac17f958d2ee523a2206206994597c13d831ec7"
const USDT_POLYGON  = "0xc2132D05D31c914a87C6611C10748AEb04B58e8F"
const USDT_AVALANCE = "0xc7198437980c041c805A1EDcbA50c1Ce5db95118"

const USDC_ETHEREUM = "0xA0b86991c6218b36c1d19D4a2e9Eb0cE3606eB48"
const USDC_BINANCE  = "0x8AC76a51cc950d9822D68b83fE1Ad97B32Cd580d"
const USDC_POLYGON  = "0x2791Bca1f2de4661ED88A30C99A7a9449Aa84174"
const USDC_AVALANCE = "0xB97EF9Ef8734C71904D8002F8b6Bc66Dd9c48a6E"
const USDC_FANTOM   = "0x04068DA6C83AFCFA0e13ba15A6696662335D5B75"

const WBTC_ETHEREUM = "0x2260FAC5E5542a773Aa44fBCfeDf7C193bc2C599"
const WBTC_POLYGON  = "0x1BFD67037B42Cf73acF2047067bd4F2C47D9BfD6"
const WBTC_AVALANCE = "0x50b7545627a5162F82A992c33b87aDc75187B218"
const WBTC_FANTOM   = "0x321162Cd933E2Be498Cd2267a90534A804051b11"

const WUST_ETHEREUM = "0xa47c8bf37f92aBed4A126BDA807A7b7498661acD"
const WUST_BINANCE  = "0x23396cF899Ca06c4472205fC903bDB4de249D6fC"
const WUST_POLYGON  = "0x692597b009d13C4049a947CAB2239b7d6517875F"

const AVAX_BINANCE  = "0x1CE0c2827e2eF14D5C4f29a091d735A204794041"
const AVAX_AVALANCE = "0xB31f66AA3C1e785363F0875A1B74E27b85FD66c7"
const AVAX_FANTOM   = "0x511D35c52a3C244E7b8bd92c0C297755FbD89212"

func initSet() *TokenSet {
  return &TokenSet {
    set: make(map[string]struct{}),
  }
}

type TokenSet struct {
  set map[string]struct{}
}

/* Exists(int64) bool */
func (ts *TokenSet) exists(token string) bool {
  _, exists := ts.set[token]
  return exists
}

/* GetSet(int64) []int64
 * Error: Returns nil if argument does not exist
 */
func (ts *TokenSet) getSet() []string {
  tokens := []string{}
  for i := range ts.set {
    tokens = append(tokens, i)
  }
  return tokens
}

/* Add(int64) error
 * Error: Throws error if token already exists
 */
func (ts *TokenSet) add(token string) error {
  _, exists := ts.set[token]
  if exists { return fmt.Errorf("ERROR: Add(int64)") }
  ts.set[token] = struct{}{}
  return nil
}

/* Remove(int64) error
 * Error: Throws error if token does not exist
 */
func (ts *TokenSet) remove(token string) error {
  _, exists := ts.set[token]
  if !exists { return fmt.Errorf("ERROR: Remove(int64)") }
  delete(ts.set, token)
  return nil
}

var WETH *TokenSet
var WBNB *TokenSet
var WUST *TokenSet
var WBTC *TokenSet
var USDT *TokenSet
var USDC *TokenSet
var AVAX *TokenSet

var sets [7]*TokenSet

func initSets() {
  WETH = initSet()
  WETH.add(WETH_ETHEREUM)
  WETH.add(WETH_FANTOM)
  WETH.add(WETH_BINANCE)
  sets[0] = WETH


  WBNB = initSet()
  WBNB.add(WBNB_BINANCE)
  WBNB.add(WBNB_ETHEREUM)
  WBNB.add(WBNB_POLYGON)
  sets[1] = WBNB

  WUST = initSet()
  WUST.add(WUST_BINANCE)
  WUST.add(WUST_POLYGON)
  WUST.add(WUST_ETHEREUM)
  sets[2] = WUST

  WBTC = initSet()
  WBTC.add(WBTC_FANTOM)
  WBTC.add(WBTC_POLYGON)
  WBTC.add(WBTC_ETHEREUM)
  WBTC.add(WBTC_AVALANCE)
  sets[3] = WBTC

  USDT = initSet()
  USDT.add(USDT_POLYGON)
  USDT.add(USDT_ETHEREUM)
  USDT.add(USDT_AVALANCE)
  sets[4] = USDT

  USDC = initSet()
  USDC.add(USDC_FANTOM)
  USDC.add(USDC_AVALANCE)
  USDC.add(USDC_ETHEREUM)
  USDC.add(USDC_POLYGON)
  USDC.add(USDC_BINANCE)
  sets[5] = USDC

  AVAX = initSet()
  AVAX.add(AVAX_FANTOM)
  AVAX.add(AVAX_BINANCE)
  AVAX.add(AVAX_AVALANCE)
  sets[6] = AVAX
}

func fetchSet(token string) []string {
  for i := 0; i < 7; i++ {
    if !sets[i].exists(token) { continue }
    return sets[i].getSet()
  }
  return nil
}

// Token: BNB
// ==========
// Ether:     0xB8c77482e45F1F44dE1745F52C74426C631bDD52
// Binance:   0xbb4CdB9CBd36B01bD1cBaEBF2De08d9173bc095c
// Polygon:   0x3BA4c387f786bFEE076A58914F5Bd38d668B42c3

// Token: ETH
// ==========
// Ether:     0xc02aaa39b223fe8d0a0e5c4f27ead9083c756cc2
// Binance:   0x2170ed0880ac9a755fd29b2688956bd959f933f8
// Fantom:    0x658b0c7613e890ee50b8c4bc6a3f41ef411208ad

// Token: USDT
// ===========
// Ether:     0xdac17f958d2ee523a2206206994597c13d831ec7
// Polygon:   0xc2132D05D31c914a87C6611C10748AEb04B58e8F
// Avalanche: 0xc7198437980c041c805A1EDcbA50c1Ce5db95118

// Token: USDC
// ===========
// Ether:     0xA0b86991c6218b36c1d19D4a2e9Eb0cE3606eB48
// Binance:   0x8AC76a51cc950d9822D68b83fE1Ad97B32Cd580d
// Polygon:   0x2791Bca1f2de4661ED88A30C99A7a9449Aa84174
// Avalanche: 0xB97EF9Ef8734C71904D8002F8b6Bc66Dd9c48a6E
// Fantom:    0x04068DA6C83AFCFA0e13ba15A6696662335D5B75

// Token: BTC
// ==========
// Ether:     0x2260FAC5E5542a773Aa44fBCfeDf7C193bc2C599
// Polygon:   0x1BFD67037B42Cf73acF2047067bd4F2C47D9BfD6
// Avalanche: 0x50b7545627a5162F82A992c33b87aDc75187B218
// Fantom:    0x321162Cd933E2Be498Cd2267a90534A804051b11

// Token: UST
// ==========
// Ether:     0xa47c8bf37f92aBed4A126BDA807A7b7498661acD
// Binance:   0x23396cF899Ca06c4472205fC903bDB4de249D6fC
// Polygon:   0x692597b009d13C4049a947CAB2239b7d6517875F

// Token: AVAX
// ===========
// Binance:   0x1CE0c2827e2eF14D5C4f29a091d735A204794041
// Avalanche: 0xB31f66AA3C1e785363F0875A1B74E27b85FD66c7
// Fantom:    0x511D35c52a3C244E7b8bd92c0C297755FbD89212
