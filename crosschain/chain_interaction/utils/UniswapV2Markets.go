package utils

import (
	"log"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"

	// uniswapV2FactoryFile "chain_interaction/generatedContracts"
	UniswapQuery "chain_interaction/UniswapQuery"
)

// useful for testing purposes, so we don't have to load all markets (takes time)
const BATCH_COUNT_LIMIT int = 100
const UNISWAP_BATCH_SIZE int = 1000

const UNISWAP_QUERY_ADDRESS_ROPSTEN string = "0x00016943476b76256b31dd90aa9d0ecc7f2c4d38"
const MY_ADDRESS string = "0x30429A2FfAE3bE74032B6ADD7ac4A971AbAd4d02"

func uniswapV2MarketByFactory(client *ethclient.Client, address string) {
	myAddress := common.HexToAddress(MY_ADDRESS)
	factoryAddress := common.HexToAddress(address)

	// slice with all marketPairs
	marketPairs := []int{}

	uniswapQuery, err := UniswapQuery.NewUniswapQuery(myAddress, client)
	if err != nil {
		log.Fatal(err)
	}

	for i := 0; i < BATCH_COUNT_LIMIT*UNISWAP_BATCH_SIZE; i += UNISWAP_BATCH_SIZE {
		// interact with contract
	}

}

func UniswapV2Markets(client *ethclient.Client, addresses []string) string {
	// for every address, get markets
	for i := 0; i < len(addresses); i++ {
		uniswapV2MarketByFactory(client, addresses[i])

	}
	return "hey"
}

// abigen --bin=./builds/UniswapQuery.bin --abi=./builds/UniswapQuery.abi --pkg=generatedContracts --out=./generatedContracts/UniswapQuery.go
