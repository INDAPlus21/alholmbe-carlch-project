package utils

import (
	"fmt"
	"log"
	"math/big"

	UniswapQuery "chain_interaction/UniswapQuery"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
)

// useful for testing purposes, so we don't have to load all markets (takes time)
const BATCH_COUNT_LIMIT int = 5
const UNISWAP_BATCH_SIZE int = 1000

const UNISWAP_QUERY_ADDRESS_ROPSTEN string = "0x00016943476b76256b31dd90aa9d0ecc7f2c4d38"
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

func uniswapV2MarketByFactory(client *ethclient.Client, address string) []UniswapV2EthPair {
	wethAddress := common.HexToAddress(WETH_ADDRESS)
	uniswapQueryAddress := common.HexToAddress(UNISWAP_QUERY_ADDRESS_ROPSTEN)
	factoryAddress := common.HexToAddress(address)

	// list with all marketPairs
	marketPairs := []UniswapV2EthPair{}

	uniswapQuery, err := UniswapQuery.NewUniswapQuery(uniswapQueryAddress, client)
	if err != nil {
		log.Fatal(err)
	}

	for i := 0; i < BATCH_COUNT_LIMIT*UNISWAP_BATCH_SIZE; i += UNISWAP_BATCH_SIZE {
		// get pairs from the network
		pairs, err := uniswapQuery.GetPairsByRange(nil, factoryAddress, big.NewInt(int64(i)), big.NewInt(int64(i+UNISWAP_BATCH_SIZE)))
		if err != nil {
			// this happens because something is weird in the sushiswap factory contract
			fmt.Printf("revert at i = %d, factoryAddress = %s", i, factoryAddress)
			log.Fatal(err)
		}

		for j := 0; j < len(pairs); j++ {
			pair := pairs[j]
			pairAddress := pair[2]

			if pair[0] != wethAddress && pair[1] != wethAddress {
				// we don't care if none of the tokens in the pair is weth
				continue
			}

			uniswapV2EthPair := UniswapV2EthPair{pairAddress, pair[0], pair[1]}

			marketPairs = append(marketPairs, uniswapV2EthPair)
		}
	}

	return marketPairs

}

func UniswapV2Markets(client *ethclient.Client, addresses []string) map[string][]UniswapV2EthPair {
	WETH := common.HexToAddress(WETH_ADDRESS)
	// for every address, get markets
	// markets is a list with all pairs on this network
	markets := [][]UniswapV2EthPair{}
	for i := 0; i < len(addresses); i++ {
		marketPairs := uniswapV2MarketByFactory(client, addresses[i])
		markets = append(markets, marketPairs)
	}

	// group markets by non weth token address
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
	return marketsByToken
}

// abigen --bin=./builds/UniswapQuery.bin --abi=./builds/UniswapQuery.abi --pkg=generatedContracts --out=./generatedContracts/UniswapQuery.go
