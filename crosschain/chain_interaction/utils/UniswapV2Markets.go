package utils

import (
	"fmt"
	"log"
	"math/big"

	UniswapQuery "chain_interaction/UniswapQuery"
	UniswapV2Factory "chain_interaction/generatedContracts"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
)

// useful for testing purposes, so we don't have to load all markets (takes time)
const BATCH_COUNT_LIMIT int = 50
const UNISWAP_BATCH_SIZE int = 100

const MY_ADDRESS string = "0x30429A2FfAE3bE74032B6ADD7ac4A971AbAd4d02"

const WETH_ADDRESS string = "0xC02aaA39b223FE8D0A0e5C4F27eAD9083C756Cc2"

type UniswapV2EthPair struct {
	PairAddress   common.Address
	Token0Address common.Address
	Token1Address common.Address
}

// mapping from non weth-token to it's markets
type MarketByTokenAddress struct {
}

func uniswapV2MarketByFactory(client *ethclient.Client, address string, queryContractAddress string, baseCurrencyAddress string) []UniswapV2EthPair {
	baseCurrency := common.HexToAddress(baseCurrencyAddress)
	uniswapQueryAddress := common.HexToAddress(queryContractAddress)
	factoryAddress := common.HexToAddress(address)

	// list with all marketPairs
	marketPairs := []UniswapV2EthPair{}
	uniswapV2Factory, err := UniswapV2Factory.NewGeneratedContracts(factoryAddress, client)
	if err != nil {
		log.Fatal(err)
	}
	uniswapQuery, err := UniswapQuery.NewUniswapQuery(uniswapQueryAddress, client)
	if err != nil {
		log.Fatal(err)
	}

	bigNum, err := uniswapV2Factory.AllPairsLength(nil)
	numberOfPairs := int(bigNum.Int64())
	fmt.Println(numberOfPairs)

	var x int
	if numberOfPairs > UNISWAP_BATCH_SIZE*BATCH_COUNT_LIMIT {
		x = UNISWAP_BATCH_SIZE * BATCH_COUNT_LIMIT
	} else {
		x = numberOfPairs
	}

	for i := 0; i < x; i += UNISWAP_BATCH_SIZE {
		// get pairs from the network
		pairs, err := uniswapQuery.GetPairsByRange(nil, factoryAddress, big.NewInt(int64(i)), big.NewInt(int64(i+UNISWAP_BATCH_SIZE)))
		if err != nil {
			// this happens because something is weird in the sushiswap factory contract
			fmt.Printf("revert at i = %d, factoryAddress = %s\n", i, factoryAddress)
			log.Fatal(err)
		}

		for j := 0; j < len(pairs); j++ {
			pair := pairs[j]
			pairAddress := pair[2]

			if pair[0] != baseCurrency && pair[1] != baseCurrency {
				// we don't care if none of the tokens in the pair is weth
				continue
			}

			uniswapV2EthPair := UniswapV2EthPair{pairAddress, pair[0], pair[1]}

			marketPairs = append(marketPairs, uniswapV2EthPair)
		}
	}

	return marketPairs

}

func UniswapV2Markets(client *ethclient.Client, addresses []string, queryContractAddress string, baseCurrencyAddress string) (map[string][]UniswapV2EthPair, map[string][]UniswapV2EthPair) {
	WETH := common.HexToAddress(WETH_ADDRESS)

	// for every address, get markets
	// markets is a list with all pairs on this network
	markets := [][]UniswapV2EthPair{}
	for i := 0; i < len(addresses); i++ {
		marketPairs := uniswapV2MarketByFactory(client, addresses[i], queryContractAddress, baseCurrencyAddress)
		markets = append(markets, marketPairs)
	}

	// group markets by non weth token address
	// mapped from the non weth token address
	marketsByToken := map[string][]UniswapV2EthPair{}

	// groups all pairs into a dictionary with the non weth token as the key
	// O(n^2) doesn't matter, this function will only run on startup
	for i := 0; i < len(markets); i++ {
		for j := 0; j < len(markets[i]); j++ {
			if markets[i][j].Token0Address == WETH {
				if _, ok := marketsByToken[markets[i][j].Token1Address.String()]; ok {
					marketsByToken[markets[i][j].Token1Address.String()] =
						append(marketsByToken[markets[i][j].Token1Address.String()], markets[i][j])
				} else {
					marketsByToken[markets[i][j].Token1Address.String()] = []UniswapV2EthPair{markets[i][j]}
				}
			} else {
				if _, ok := marketsByToken[markets[i][j].Token0Address.String()]; ok {
					marketsByToken[markets[i][j].Token0Address.String()] =
						append(marketsByToken[markets[i][j].Token0Address.String()], markets[i][j])
				} else {
					marketsByToken[markets[i][j].Token0Address.String()] = []UniswapV2EthPair{markets[i][j]}
				}
			}
		}
	}

	// a cross markets exists if the same market exists on 2+ places on 1 network
	crossMarkets := make(map[string][]UniswapV2EthPair)
	for k, v := range marketsByToken {
		if len(v) > 1 {
			crossMarkets[k] = v
		}
	}
	return marketsByToken, crossMarkets
}

// abigen --bin=./builds/UniswapQuery.bin --abi=./builds/UniswapQuery.abi --pkg=generatedContracts --out=./generatedContracts/UniswapQuery.go
